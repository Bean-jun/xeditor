package main

import (
	"embed"
	"log"
	"xeditor/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	App := app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "xeditor",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: App.Startup,
		Bind: []interface{}{
			App,
		},
		Frameless: true,
		Debug:     options.Debug{OpenInspectorOnStartup: true},
	})

	if err != nil {
		log.Println("Error:", err.Error())
	}
}
