package main

import (
	"context"
	"wslpathmanager-go/backend"
)

// App struct
type App struct {
	ctx          context.Context
	wslService   *backend.WslService
	shimService  *backend.ShimService
	pathInjector *backend.PathInjectorService
	stateService *backend.StateService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		wslService:   backend.NewWslService(),
		shimService:  backend.NewShimService(),
		pathInjector: backend.NewPathInjectorService(),
		stateService: backend.NewStateService(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetDistros() ([]backend.WslDistro, error) {
	return a.wslService.GetDistros()
}

func (a *App) ScanPaths(distroName string) ([]backend.SelectableTool, error) {
	config := backend.DefaultAppConfig()
	toolsToScan := append(config.ToolList, config.CustomTools...)
	
	paths, err := a.wslService.ScanPaths(distroName, toolsToScan)
	if err != nil {
		return nil, err
	}
	
	profiles, err := a.stateService.GetProfiles()
	if err != nil {
		return nil, err
	}
	
	var savedTools []string
	if profile, ok := profiles[distroName]; ok {
		savedTools = profile.SelectedTools
	}
	
	for i, path := range paths {
		for _, savedTool := range savedTools {
			if path.ToolName == savedTool {
				paths[i].IsSelected = true
				break
			}
		}
	}
	
	return paths, nil
}

func (a *App) SaveSettings(distroName string, selectedTools []string) error {
	// 1. Create or Update Shims
	err := a.shimService.CreateShims(distroName, selectedTools)
	if err != nil {
		return err
	}
	
	// 2. Save to profiles.json
	profiles, _ := a.stateService.GetProfiles()
	if profiles == nil {
		profiles = make(map[string]backend.DistroProfile)
	}
	
	profile := profiles[distroName]
	profile.SelectedTools = selectedTools
	profiles[distroName] = profile
	
	err = a.stateService.SaveProfiles(profiles)
	if err != nil {
		return err
	}
	
	// 3. Inject or Remove from Windows PATH based on selection
	shimDir := a.shimService.GetShimDirectory(distroName)
	if len(selectedTools) > 0 {
		return a.pathInjector.InjectPath([]string{shimDir})
	} else {
		return a.pathInjector.RemovePath([]string{shimDir})
	}
}
