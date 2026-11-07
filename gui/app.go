package main

import (
	"context"
	"fmt"

	"github.com/dr2cc/mytranslator/internal/translate"
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
// Метод startup вызывается при запуске приложения. Контекст сохраняется,
// чтобы можно было вызывать методы среды выполнения (runtime).
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// // Greet returns a greeting for the given name
// func (a *App) Greet(name string) string {
// 	return fmt.Sprintf("Hello %s, It's show time!", name)
// }

func (a *App) Omg(star string) string {
	// Создаем экземпляр Service. Не используем DI (плохо!?)
	//❗Как правильно использовать DI в такой связке?
	translateService := translate.NewService()
	// Обращаемся к Translate (методу Service struct{})
	result, _ := translateService.Translate(star, "ru")
	//❗Не обрабатываем ошибку. Это плохо! Или нет? Уточнить!!
	return fmt.Sprintf("OMG! It's you, %s!", result)
}
