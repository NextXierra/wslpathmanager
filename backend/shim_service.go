package backend

import (
	"fmt"
	"os"
	"path/filepath"
)

// ShimService handles the creation and management of batch script shims for WSL tools.
type ShimService struct{}

// NewShimService creates a new instance of ShimService.
func NewShimService() *ShimService {
	return &ShimService{}
}

// GetShimDirectory returns the absolute path to the directory where shims for a distro are stored.
func (s *ShimService) GetShimDirectory(distroName string) string {
	appData, err := os.UserCacheDir() // This maps to LocalAppData on Windows
	if err != nil {
		appData = os.Getenv("LOCALAPPDATA")
	}
	return filepath.Join(appData, "wslpathmanager", "shims", distroName)
}

// CreateShims generates executable batch scripts (.bat) in the host Windows environment 
// that forward commands to their respective Linux tools inside the specified WSL distribution.
func (s *ShimService) CreateShims(distroName string, tools []string) error {
	shimDir := s.GetShimDirectory(distroName)

	if _, err := os.Stat(shimDir); err == nil {
		os.RemoveAll(shimDir)
	}

	err := os.MkdirAll(shimDir, os.ModePerm)
	if err != nil {
		return err
	}

	for _, tool := range tools {
		shimPath := filepath.Join(shimDir, tool+".cmd")
		scriptContent := fmt.Sprintf("@echo off\r\nwsl.exe -d %s -e %s %%*", distroName, tool)
		err := os.WriteFile(shimPath, []byte(scriptContent), 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
