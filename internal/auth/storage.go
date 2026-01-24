package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

)

// StorageBackend defines the interface for credential storage
type StorageBackend interface {
	Save(profile string, data []byte) error
	Load(profile string) ([]byte, error)
	Delete(profile string) error
	Name() string
}

// KeyringStorage uses system keyring for credential storage
type KeyringStorage struct {
	serviceName string
}

// NewKeyringStorage creates a keyring storage backend
func NewKeyringStorage(serviceName string) *KeyringStorage {
	return &KeyringStorage{
		serviceName: serviceName,
	}
}

func (s *KeyringStorage) Save(profile string, data []byte) error {
	// Import keyring here to avoid issues if not available
	// The keyring library is already imported in manager.go
	return saveToKeyring(s.serviceName, profile, string(data))
}

func (s *KeyringStorage) Load(profile string) ([]byte, error) {
	data, err := loadFromKeyring(s.serviceName, profile)
	if err != nil {
		return nil, err
	}
	return []byte(data), nil
}

func (s *KeyringStorage) Delete(profile string) error {
	return deleteFromKeyring(s.serviceName, profile)
}

func (s *KeyringStorage) Name() string {
	return "system-keyring"
}

// EncryptedFileStorage stores credentials in encrypted files
type EncryptedFileStorage struct {
	baseDir string
	key     []byte
}

// NewEncryptedFileStorage creates an encrypted file storage backend
func NewEncryptedFileStorage(baseDir string) (*EncryptedFileStorage, error) {
	// Generate or load encryption key
	key, err := getOrCreateEncryptionKey(baseDir)
	if err != nil {
		return nil, fmt.Errorf("failed to get encryption key: %w", err)
	}

	return &EncryptedFileStorage{
		baseDir: baseDir,
		key:     key,
	}, nil
}

func (s *EncryptedFileStorage) Save(profile string, data []byte) error {
	encrypted, err := s.encrypt(data)
	if err != nil {
		return fmt.Errorf("failed to encrypt credentials: %w", err)
	}

	credFile := s.getCredentialFilePath(profile)
	if err := os.MkdirAll(filepath.Dir(credFile), 0700); err != nil {
		return err
	}

	return os.WriteFile(credFile, encrypted, 0600)
}

func (s *EncryptedFileStorage) Load(profile string) ([]byte, error) {
	credFile := s.getCredentialFilePath(profile)
	encrypted, err := os.ReadFile(credFile)
	if err != nil {
		return nil, fmt.Errorf("credentials not found for profile '%s'", profile)
	}

	return s.decrypt(encrypted)
}

func (s *EncryptedFileStorage) Delete(profile string) error {
	credFile := s.getCredentialFilePath(profile)
	return os.Remove(credFile)
}

func (s *EncryptedFileStorage) Name() string {
	return "encrypted-file"
}

func (s *EncryptedFileStorage) getCredentialFilePath(profile string) string {
	return filepath.Join(s.baseDir, "credentials", profile+".enc")
}

// encrypt encrypts data using AES-GCM
func (s *EncryptedFileStorage) encrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(s.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// decrypt decrypts data using AES-GCM
func (s *EncryptedFileStorage) decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(s.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, fmt.Errorf("invalid ciphertext")
	}

	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt credentials: %w", err)
	}

	return plaintext, nil
}

// PlainFileStorage stores credentials in plain JSON files (legacy/development only)
type PlainFileStorage struct {
	baseDir string
}

// NewPlainFileStorage creates a plain file storage backend
func NewPlainFileStorage(baseDir string) *PlainFileStorage {
	return &PlainFileStorage{
		baseDir: baseDir,
	}
}

func (s *PlainFileStorage) Save(profile string, data []byte) error {
	credFile := s.getCredentialFilePath(profile)
	if err := os.MkdirAll(filepath.Dir(credFile), 0700); err != nil {
		return err
	}
	return os.WriteFile(credFile, data, 0600)
}

func (s *PlainFileStorage) Load(profile string) ([]byte, error) {
	credFile := s.getCredentialFilePath(profile)
	data, err := os.ReadFile(credFile)
	if err != nil {
		return nil, fmt.Errorf("credentials not found for profile '%s'", profile)
	}
	return data, nil
}

func (s *PlainFileStorage) Delete(profile string) error {
	credFile := s.getCredentialFilePath(profile)
	return os.Remove(credFile)
}

func (s *PlainFileStorage) Name() string {
	return "plain-file"
}

func (s *PlainFileStorage) getCredentialFilePath(profile string) string {
	return filepath.Join(s.baseDir, "credentials", profile+".json")
}

// getOrCreateEncryptionKey generates or loads the encryption key
func getOrCreateEncryptionKey(baseDir string) ([]byte, error) {
	keyFile := filepath.Join(baseDir, ".keyfile")

	// Try to load existing key
	if data, err := os.ReadFile(keyFile); err == nil {
		key, err := base64.StdEncoding.DecodeString(string(data))
		if err == nil && len(key) == 32 {
			return key, nil
		}
	}

	// Generate new key
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	// Save key
	if err := os.MkdirAll(baseDir, 0700); err != nil {
		return nil, err
	}

	encoded := base64.StdEncoding.EncodeToString(key)
	if err := os.WriteFile(keyFile, []byte(encoded), 0600); err != nil {
		return nil, err
	}

	return key, nil
}

// ListProfiles lists all stored credential profiles
func (m *Manager) ListProfiles() ([]string, error) {
	var profiles []string

	if m.useKeyring {
		// For keyring storage, we need to track profiles separately
		// Read from a profiles list file
		profilesFile := filepath.Join(m.configDir, "profiles.json")
		data, err := os.ReadFile(profilesFile)
		if err != nil {
			if os.IsNotExist(err) {
				return []string{}, nil
			}
			return nil, err
		}

		if err := json.Unmarshal(data, &profiles); err != nil {
			return nil, err
		}
	} else {
		// For file storage, list files in credentials directory
		credDir := filepath.Join(m.configDir, "credentials")
		entries, err := os.ReadDir(credDir)
		if err != nil {
			if os.IsNotExist(err) {
				return []string{}, nil
			}
			return nil, err
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				name := entry.Name()
				// Remove .json or .enc extension
				if ext := filepath.Ext(name); ext == ".json" || ext == ".enc" {
					profiles = append(profiles, name[:len(name)-len(ext)])
				}
			}
		}
	}

	return profiles, nil
}

// addProfileToList adds a profile to the tracked list (for keyring storage)
func (m *Manager) addProfileToList(profile string) error {
	if !m.useKeyring {
		return nil // Not needed for file storage
	}

	profiles, err := m.ListProfiles()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// Check if already exists
	for _, p := range profiles {
		if p == profile {
			return nil
		}
	}

	profiles = append(profiles, profile)
	data, err := json.Marshal(profiles)
	if err != nil {
		return err
	}

	profilesFile := filepath.Join(m.configDir, "profiles.json")
	if err := os.MkdirAll(m.configDir, 0700); err != nil {
		return err
	}

	return os.WriteFile(profilesFile, data, 0600)
}

// removeProfileFromList removes a profile from the tracked list
func (m *Manager) removeProfileFromList(profile string) error {
	if !m.useKeyring {
		return nil // Not needed for file storage
	}

	profiles, err := m.ListProfiles()
	if err != nil {
		return err
	}

	var updated []string
	for _, p := range profiles {
		if p != profile {
			updated = append(updated, p)
		}
	}

	data, err := json.Marshal(updated)
	if err != nil {
		return err
	}

	profilesFile := filepath.Join(m.configDir, "profiles.json")
	return os.WriteFile(profilesFile, data, 0600)
}
