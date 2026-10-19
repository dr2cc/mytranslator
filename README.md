# mytranslator
A simple example of using Cobra for a command-line application

mytranslator/
├── cmd/
│   └── translator/
│       └── main.go         # Точка входа (минимальный код)
├── internal/
│   ├── app/
│   │   └── app.go          # Тот самый агрегатор (Main)
│   ├── config/
│   │   └── config.go       # Настройка Viper и структура конфига
│   ├── translate/
│   │   └── service.go      # Бизнес-логика перевода (API)
│   └── ui/
│       └── cli/
│           └── root.go     # Главная команда Cobra
├── config.yaml             # Конфигурационный файл
├── go.mod
└── go.sum
