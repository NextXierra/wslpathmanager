package backend

type AppConfig struct {
	ToolList    []string `json:"toolList"`
	CustomTools []string `json:"customTools"`
}

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

type WslDistro struct {
	Name string `json:"name"`
}

type SelectableTool struct {
	ToolName   string `json:"toolName"`
	WslPath    string `json:"wslPath"`
	IsSelected bool   `json:"isSelected"`
}

type DistroProfile struct {
	SelectedTools []string `json:"selectedTools"`
}
