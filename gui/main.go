package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "mytranslator",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		// Привязка методов (method binding) - методы структуры App становятся доступны во фронтенде как JS-функции
		// (см. frontend/wailsjs/go/main/App.js — генерируется Wails).
		Bind: []interface{}{
			app,
			// Можно привязать любое количество структур.
			// Убедитесь, что вы создали экземпляр структуры и передайте его в качестве параметра Bind:
			// &mystruct1{},
			// &mystruct2{},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
