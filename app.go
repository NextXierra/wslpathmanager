package main

import (
	"context"
	"wslpathmanager-go/backend"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App provides the application structure and methods exposed to the frontend.
type App struct {
	ctx              context.Context
	wslService       *backend.WslService
	shimService      *backend.ShimService
	pathInjector     *backend.PathInjectorService
	stateService     *backend.StateService
	autostartService *backend.AutostartService
}

// NewApp creates a new App application struct containing all required backend services.
func NewApp() *App {
	return &App{
		wslService:       backend.NewWslService(),
		shimService:      backend.NewShimService(),
		pathInjector:     backend.NewPathInjectorService(),
		stateService:     backend.NewStateService(),
		autostartService: backend.NewAutostartService(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetDistros retrieves all available WSL distributions via the WslService.
func (a *App) GetDistros() ([]backend.WslDistro, error) {
	return a.wslService.GetDistros()
}

// ScanPaths scans a given WSL distribution for a set of predefined and custom tools.
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

// SaveSettings saves the user's selected tools for a distro, generates shims, and updates the Windows PATH.
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

// ToggleAutoStart enables or disables the auto-startup feature using the AutostartService.
func (a *App) ToggleAutoStart(enabled bool) error {
	return a.autostartService.ToggleAutoStart(enabled)
}

// GetAutoStartStatus returns the current status of the auto-startup setting.
func (a *App) GetAutoStartStatus() bool {
	return a.autostartService.IsAutoStartEnabled()
}

// ToggleTraySetting enables or disables the minimize-to-tray on close behavior.
func (a *App) ToggleTraySetting(enabled bool) error {
	settings, _ := a.stateService.GetSettings()
	settings.MinimizeToTray = enabled
	return a.stateService.SaveSettings(settings)
}

// GetTraySetting returns the current status of the minimize-to-tray setting.
func (a *App) GetTraySetting() bool {
	settings, _ := a.stateService.GetSettings()
	return settings.MinimizeToTray
}

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	settings, _ := a.stateService.GetSettings()
	if settings.MinimizeToTray {
		runtime.WindowHide(ctx)
		return true
	}
	return false
}
