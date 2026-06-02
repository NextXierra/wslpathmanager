package backend

import (
	"golang.org/x/sys/windows/registry"
	"os"
	"path/filepath"
)

const autoStartKeyName = "wslpathmanager"

// AutostartService manages the Windows Registry for the auto-startup feature.
type AutostartService struct{}

// NewAutostartService creates a new instance of AutostartService.
func NewAutostartService() *AutostartService {
	return &AutostartService{}
}

func (s *AutostartService) getExecutablePath() string {
	exePath, err := os.Executable()
	if err != nil {
		return ""
	}
	// Return absolute path just in case
	absPath, _ := filepath.Abs(exePath)
	// Enclose in quotes to handle paths with spaces
	return `"` + absPath + `"`
}

// IsAutoStartEnabled checks if the application is set to start on boot.
func (s *AutostartService) IsAutoStartEnabled() bool {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer key.Close()

	val, _, err := key.GetStringValue(autoStartKeyName)
	if err != nil {
		return false
	}

	return val == s.getExecutablePath()
}

// ToggleAutoStart enables or disables the auto-startup feature in the Windows registry.
func (s *AutostartService) ToggleAutoStart(enabled bool) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	if enabled {
		return key.SetStringValue(autoStartKeyName, s.getExecutablePath())
	} else {
		return key.DeleteValue(autoStartKeyName)
	}
}
