package auth

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEncryptedFileStorage(t *testing.T) {
	// Create temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "gdrv-test-*")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Fatalf("failed to remove temp dir: %v", err)
		}
	})

	storage, err := NewEncryptedFileStorage(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create encrypted storage: %v", err)
	}

	testData := []byte(`{"profile":"test","access_token":"test-token"}`)

	// Test Save
	err = storage.Save("test-profile", testData)
	if err != nil {
		t.Errorf("Save failed: %v", err)
	}

	// Verify file exists and is encrypted
	credFile := filepath.Join(tmpDir, "credentials", "test-profile.enc")
	encryptedData, err := os.ReadFile(credFile)
	if err != nil {
		t.Errorf("Failed to read encrypted file: %v", err)
	}

	// Encrypted data should not match original
	if string(encryptedData) == string(testData) {
		t.Error("Data was not encrypted")
	}

	// Test Load
	loaded, err := storage.Load("test-profile")
	if err != nil {
		t.Errorf("Load failed: %v", err)
	}

	if string(loaded) != string(testData) {
		t.Errorf("Loaded data doesn't match original. Got: %s, Want: %s", string(loaded), string(testData))
	}

	// Test Delete
	err = storage.Delete("test-profile")
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	// Verify file is deleted
	if _, err := os.Stat(credFile); !os.IsNotExist(err) {
		t.Error("File was not deleted")
	}
}

func TestPlainFileStorage(t *testing.T) {
	// Create temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "gdrv-test-*")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Fatalf("failed to remove temp dir: %v", err)
		}
	})

	storage := NewPlainFileStorage(tmpDir)
	testData := []byte(`{"profile":"test","access_token":"test-token"}`)

	// Test Save
	err = storage.Save("test-profile", testData)
	if err != nil {
		t.Errorf("Save failed: %v", err)
	}

	// Test Load
	loaded, err := storage.Load("test-profile")
	if err != nil {
		t.Errorf("Load failed: %v", err)
	}

	if string(loaded) != string(testData) {
		t.Errorf("Loaded data doesn't match. Got: %s, Want: %s", string(loaded), string(testData))
	}

	// Test Delete
	err = storage.Delete("test-profile")
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}
}

func TestEncryptionRoundTrip(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gdrv-test-*")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Fatalf("failed to remove temp dir: %v", err)
		}
	})

	storage, err := NewEncryptedFileStorage(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create encrypted storage: %v", err)
	}

	testCases := []string{
		"simple text",
		`{"complex":"json","with":"values"}`,
		"text with special characters: üöä@#$%^&*()",
		"",
	}

	for i, testData := range testCases {
		encrypted, err := storage.encrypt([]byte(testData))
		if err != nil {
			t.Errorf("Test case %d: encrypt failed: %v", i, err)
			continue
		}

		decrypted, err := storage.decrypt(encrypted)
		if err != nil {
			t.Errorf("Test case %d: decrypt failed: %v", i, err)
			continue
		}

		if string(decrypted) != testData {
			t.Errorf("Test case %d: roundtrip failed. Got: %s, Want: %s", i, string(decrypted), testData)
		}
	}
}

func TestGetOrCreateEncryptionKey(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gdrv-test-*")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Fatalf("failed to remove temp dir: %v", err)
		}
	})

	// First call should create a new key
	key1, err := getOrCreateEncryptionKey(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create key: %v", err)
	}

	if len(key1) != 32 {
		t.Errorf("Key length is %d, expected 32", len(key1))
	}

	// Second call should load the same key
	key2, err := getOrCreateEncryptionKey(tmpDir)
	if err != nil {
		t.Fatalf("Failed to load key: %v", err)
	}

	if string(key1) != string(key2) {
		t.Error("Loaded key doesn't match created key")
	}
}

func TestManagerListProfiles(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gdrv-test-*")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Fatalf("failed to remove temp dir: %v", err)
		}
	})

	// Create manager with plain file storage for easier testing
	mgr := NewManagerWithOptions(tmpDir, ManagerOptions{ForcePlainFile: true})

	// Initially should be empty
	profiles, err := mgr.ListProfiles()
	if err != nil {
		t.Fatalf("ListProfiles failed: %v", err)
	}

	if len(profiles) != 0 {
		t.Errorf("Expected 0 profiles, got %d", len(profiles))
	}

	// Save some test credentials
	testData := []byte(`{"profile":"profile1","access_token":"token1"}`)
	err = mgr.storage.Save("profile1", testData)
	if err != nil {
		t.Fatalf("Failed to save credentials: %v", err)
	}

	testData2 := []byte(`{"profile":"profile2","access_token":"token2"}`)
	err = mgr.storage.Save("profile2", testData2)
	if err != nil {
		t.Fatalf("Failed to save credentials: %v", err)
	}

	// List profiles
	profiles, err = mgr.ListProfiles()
	if err != nil {
		t.Fatalf("ListProfiles failed: %v", err)
	}

	if len(profiles) != 2 {
		t.Errorf("Expected 2 profiles, got %d", len(profiles))
	}

	// Check profile names
	profileMap := make(map[string]bool)
	for _, p := range profiles {
		profileMap[p] = true
	}

	if !profileMap["profile1"] || !profileMap["profile2"] {
		t.Errorf("Missing expected profiles. Got: %v", profiles)
	}
}
