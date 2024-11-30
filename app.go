package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	// front "github.com/wailsapp/wails/v2/internal/frontend"
	runtime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.EventsOn(ctx, "game:start", func(args ...interface{}) {
		fmt.Printf("[EVENT] game:start, args: %v\n", args)

		// Проверяем, является ли первый аргумент срезом строк
		if strSlice, ok := args[0].([]string); ok {
			// Успешно привели к []string
			fmt.Println("Successfully converted:", strSlice)
		} else {
			// Приведение не удалось
			fmt.Println("Conversion failed")
		}
	})

	runtime.EventsOn(ctx, "servers:request", func(args ...interface{}) {
		go func() {
			servers, err := LoadServers()
			if err != nil {
				runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
					Title:   "Ошибка",
					Message: "Ошибка получения списка серверов:\n" + err.Error(),
					Type:    runtime.ErrorDialog,
				})
				return
			}
			runtime.EventsEmit(ctx, "servers:update", servers.Arizona)
		}()
	})

	// Config
	runtime.EventsOn(ctx, "config:write", func(args ...interface{}) {

	})
	runtime.EventsOn(ctx, "config:request", func(args ...interface{}) {
		go func() {

		}()
	})

	// Game path
	runtime.EventsOn(ctx, "settings:requestFileDialog", func(args ...interface{}) {
		path, err := runtime.OpenDirectoryDialog(ctx, runtime.OpenDialogOptions{})
		if err != nil {
			runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
				Title:   "Ошибка",
				Message: "Ошибка при открытии диалога: " + err.Error(),
				Type:    runtime.ErrorDialog,
			})
			return
		}
		if len(path) > 0 {
			gamePath := fmt.Sprintf("%s\\gta_sa.exe", path)
			_, err := os.Stat(gamePath)
			if err != nil {
				runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
					Title:   "Error",
					Message: "Файл \"gta_sa.exe\" не найден в \"" + path + "\"",
					Type:    runtime.ErrorDialog,
				})
				return
			}
			runtime.EventsEmit(ctx, "settings:fileDialogPathSelected", path)
		}
	})
}

func (a *App) StartGame(name string, gamePath string, parameters []string) error {
	if len(name) < 3 || len(name) > 22 {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Title:   "Ошибка запуска игры",
			Message: "Некоррекнтный ник-нейм",
			Type:    runtime.ErrorDialog,
		})
		return nil
	}
	// var gamePath string = fmt.Sprintf("%s\\gta_sa.exe", gamePath)
	var batFile = fmt.Sprintf("%s\\%s", gamePath, "alternative-launcher.bat")
	var batText = fmt.Sprintf("@echo off\ncd /D %%~dp0\nstart gta_sa.exe %s\nexit", strings.Join(parameters, " "))

	err := os.WriteFile(batFile, []byte(batText), 0644)
	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Title:   "Ошибка запуска игры",
			Message: fmt.Sprintf("Ошибка: %s", err.Error()),
			Type:    runtime.ErrorDialog,
		})
		return err
	}
	cmd := exec.Command(batFile).Run()
	// proc, err := os.StartProcess(batFile, []string{}, &os.ProcAttr{})
	fmt.Println(cmd)
	return nil
}

func (a *App) ReadConfig() string {
	bytes, err := os.ReadFile(CONFIG_FILE_PATH)
	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Title:   "Ваши настройки не были загружены",
			Message: fmt.Sprintf("Произошла ошибка при чтении файла настроек: %s\nБыли загружены стандартные настройки.", err.Error()),
			Type:    runtime.InfoDialog,
		})
		return ""
	}
	return string(bytes)
}

func (a *App) SaveConfig(json string) {
	fmt.Println("SAVE CONFIG:", json)
	err := os.WriteFile(CONFIG_FILE_PATH, []byte(json), 0644)
	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Title:   "Ваши настройки не были сохранены",
			Message: "Ошибка при сохранении настроек: " + err.Error(),
			Type:    runtime.ErrorDialog,
		})
	}
}