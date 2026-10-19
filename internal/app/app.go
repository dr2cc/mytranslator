package app

import (
	"fmt"
	"os"

	"github.com/dr2cc/mytranslator/internal/config"
	"github.com/dr2cc/mytranslator/internal/translate"
	"github.com/dr2cc/mytranslator/internal/ui/cli"
)

// Main — главная функция-агрегатор, запускающая приложение.
func Main() {
	// 1. Загружаем конфигурацию (YAML)
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Ошибка загрузки конфигурации: %v\n", err)
		os.Exit(1)
	}

	// 2. Инициализируем сервисы бизнес-логики
	translateService := translate.NewService()

	// 3. Инициализируем UI/CLI слой
	rootCmd := cli.NewRootCmd(cfg, translateService)

	// 4. Запускаем выполнение командной строки
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
