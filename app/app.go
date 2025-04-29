package app

import (
	"context"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type FK struct {
	ctx context.Context
}

func (a *FK) init(ctx context.Context) {
	a.ctx = ctx
}

func (a *FK) Xs() string {
	return "???"
}

// App struct
type App struct {
	FK
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.FK.init(a.ctx)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	result, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.WarningDialog,
		Title:         "Question",
		Message:       "Do you want to continue?",
		DefaultButton: "No",
	})
	log.Println(result, err)
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
