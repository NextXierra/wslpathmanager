package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"io/fs"
)

//go:embed all:frontend
var assets embed.FS

func getAssets() fs.FS {
	subFS, err := fs.Sub(assets, "frontend")
	if err != nil {
		panic(err)
	}
	return subFS
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "WSL Path Manager",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: getAssets(),
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
