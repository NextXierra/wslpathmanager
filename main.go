package main

import (
	"embed"
	"io/fs"
	"os"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend
var assets embed.FS

//go:embed build/windows/icon.ico
var trayIcon []byte

func getAssets() fs.FS {
	subFS, err := fs.Sub(assets, "frontend")
	if err != nil {
		panic(err)
	}
	return subFS
}

func main() {
	app := NewApp()

	onReady := func() {
		systray.SetIcon(trayIcon)
		systray.SetTitle("WSL Path Manager")
		systray.SetTooltip("WSL Path Manager is running")

		mShow := systray.AddMenuItem("Show App", "Open the application window")
		mQuit := systray.AddMenuItem("Quit", "Exit the application")

		go func() {
			for {
				select {
				case <-mShow.ClickedCh:
					if app.ctx != nil {
						runtime.WindowShow(app.ctx)
					}
				case <-mQuit.ClickedCh:
					systray.Quit()
					os.Exit(0)
				}
			}
		}()
	}

	onExit := func() {
		// Clean up logic if needed
	}

	// Start tray in background
	go systray.Run(onReady, onExit)

	err := wails.Run(&options.App{
		Title:  "WSL Path Manager",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: getAssets(),
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnBeforeClose:    app.beforeClose,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
