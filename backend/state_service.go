package backend

import (
	"encoding/json"
	"os"
)

type StateService struct {
	profilesFilePath string
}

func NewStateService() *StateService {
	return &StateService{
		profilesFilePath: "profiles.json",
	}
}

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

func (s *StateService) SaveProfiles(profiles map[string]DistroProfile) error {
	content, err := json.MarshalIndent(profiles, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.profilesFilePath, content, 0644)
}
