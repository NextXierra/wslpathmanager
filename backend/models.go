package backend

// AppConfig holds the default configuration for tools to scan.
type AppConfig struct {
	ToolList    []string `json:"toolList"`
	CustomTools []string `json:"customTools"`
}

// DefaultAppConfig returns the default application configuration.
func DefaultAppConfig() *AppConfig {
	return &AppConfig{
		ToolList: []string{
			"python3", "python", "pip3",
			"node", "npm", "npx",
			"bun",
			"php", "composer",
			"ruby", "gem",
			"go",
			"java", "javac",
			"cargo", "rustc",
			"gcc", "g++", "make",
			"git",
		},
		CustomTools: []string{},
	}
}

// WslDistro represents a Windows Subsystem for Linux distribution.
type WslDistro struct {
	Name string `json:"name"`
}

// SelectableTool represents a tool found in a WSL distribution that can be selected.
type SelectableTool struct {
	ToolName   string `json:"toolName"`
	WslPath    string `json:"wslPath"`
	IsSelected bool   `json:"isSelected"`
}

// DistroProfile holds the user's selected tools for a specific distribution.
type DistroProfile struct {
	SelectedTools []string `json:"selectedTools"`
}

// AppSettings holds global application preferences.
type AppSettings struct {
	MinimizeToTray bool `json:"minimizeToTray"`
}
