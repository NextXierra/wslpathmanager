package backend

import (
	"encoding/json"
	"os"
)

// StateService provides read and write access to application state and preferences.
type StateService struct {
	profilesFilePath string
	settingsFilePath string
}

// NewStateService creates a new instance of StateService.
func NewStateService() *StateService {
	return &StateService{
		profilesFilePath: "profiles.json",
		settingsFilePath: "settings.json",
	}
}

// GetProfiles reads the saved distribution profiles from disk.
func (s *StateService) GetProfiles() (map[string]DistroProfile, error) {
	profiles := make(map[string]DistroProfile)

	if _, err := os.Stat(s.profilesFilePath); os.IsNotExist(err) {
		return profiles, nil
	}

	content, err := os.ReadFile(s.profilesFilePath)
	if err != nil {
		return profiles, err
	}

	if len(content) > 0 {
		err = json.Unmarshal(content, &profiles)
		if err != nil {
			return profiles, err
		}
	}

	return profiles, nil
}

// SaveProfiles writes the provided distribution profiles to disk.
func (s *StateService) SaveProfiles(profiles map[string]DistroProfile) error {
	content, err := json.MarshalIndent(profiles, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.profilesFilePath, content, 0644)
}

// GetSettings reads the global application settings from disk.
func (s *StateService) GetSettings() (AppSettings, error) {
	settings := AppSettings{MinimizeToTray: false} // Default off

	if _, err := os.Stat(s.settingsFilePath); os.IsNotExist(err) {
		return settings, nil
	}

	content, err := os.ReadFile(s.settingsFilePath)
	if err != nil {
		return settings, err
	}

	if len(content) > 0 {
		err = json.Unmarshal(content, &settings)
		if err != nil {
			return settings, err
		}
	}

	return settings, nil
}

// SaveSettings writes the global application settings to disk.
func (s *StateService) SaveSettings(settings AppSettings) error {
	content, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.settingsFilePath, content, 0644)
}
