package main

import (
	"changeme/pkg/ldifde"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) GetGraphData(selectType, selectValue string) (data string) {
	fmt.Printf("selectType: %s\n", selectType)
	fmt.Printf("selectValue: %s\n", selectValue)

	if selectType == "group" {
		data = ldifde.GetGroupGraph(selectValue)
	} else if selectType == "user" {
		data = ldifde.GetUserGraph(selectValue)
	}
	return
}
