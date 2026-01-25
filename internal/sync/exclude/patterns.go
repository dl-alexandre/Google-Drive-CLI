package exclude

import (
	"path"
	"strings"
)

type Matcher struct {
	patterns []string
}

func DefaultPatterns() []string {
	return []string{
		".git/",
		".DS_Store",
		"._*",
		"node_modules/",
		"vendor/",
		"*.log",
		"*.tmp",
		".env",
		".env.*",
		"*.key",
		"*.pem",
	}
}

func New(patterns []string) *Matcher {
	merged := append([]string{}, DefaultPatterns()...)
	for _, p := range patterns {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		merged = append(merged, p)
	}
	return &Matcher{patterns: merged}
}

func (m *Matcher) IsExcluded(relPath string, isDir bool) bool {
	if m == nil {
		return false
	}
	relPath = strings.TrimPrefix(relPath, "./")
	for _, p := range m.patterns {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		if strings.HasSuffix(p, "/") {
			dirPattern := strings.TrimSuffix(p, "/")
			if relPath == dirPattern || strings.HasPrefix(relPath, dirPattern+"/") {
				return true
			}
			continue
		}
		if strings.ContainsAny(p, "*?[]") {
			if ok, _ := path.Match(p, relPath); ok {
				return true
			}
			base := path.Base(relPath)
			if ok, _ := path.Match(p, base); ok {
				return true
			}
			continue
		}
		if relPath == p || strings.HasPrefix(relPath, p+"/") {
			return true
		}
		if !isDir && path.Base(relPath) == p {
			return true
		}
	}
	return false
}
