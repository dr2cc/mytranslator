package app

import (
	"fmt"
	"os"

	"github.com/dr2cc/mytranslator/internal/config"
	"github.com/dr2cc/mytranslator/internal/translate"
	"github.com/dr2cc/mytranslator/internal/ui/cli"
)

// Агрегатор
func Run() {
	// Загружаем конфигурацию (YAML)
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Ошибка загрузки конфигурации: %v\n", err)
		os.Exit(1)
	}

	// ⬇ Dependency Injection
	// Только применяя DI мы делаем Run() агрегатором и именно в нем решаем, какую именно реализацию выдать нашей Cobra-команде.

	// Если мы сделаем translate.NewService() глобальной функцией, а не структурой, то:
	// • Пакет app теперь НЕ управляет. Cobra-команда будет сама решать, откуда брать перевод.
	// • Функцию перевода заменить на другую НЕ возможно. Если завтра понадобится DeepL, придется лезть в код Cobra-команды и менять вызов translate.Translate на deepl.Translate.
	// • Код НЕвозможно протестировать локально. При запуске тестов Cobra-команда всегда будет пытаться отправить настоящий сетевой запрос через эту глобальную функцию.

	// Инициализируем сервисы бизнес-логики
	translateService := translate.NewService()

	// ⬇ (DI) UI/CLI слою нужно то, что делают сервисы
	rootCmd := cli.NewRootCmd(cfg, translateService)

	// Чтобы дерево команд ожило и Cobra распарсила ввод пользователя (какие флаги переданы, какие аргументы написаны),
	// вызываем метод Execute():
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	// Метод Execute() сам находит нужную подкоманду, проверяет её флаги, валидирует аргументы, запускает цепочку хуков и выполняет функцию RunE.
}
