package cli

import (
	"fmt"

	"github.com/dr2cc/mytranslator/internal/config"
	"github.com/dr2cc/mytranslator/internal/translate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Можно легко менять тип параметра transService с интерфейса (контракта) на конкретный сервис.

// // interface Уже готов к рефакторингу (когда буду делать тесты).
// // Описываем контракт:
// // "мне подойдет любой объект, у которого есть метод Translate"
// type TextTranslator interface {
// 	Translate(text, targetLang string) (string, error)
// }

// // Здесь команда принимает контракт (интерфейс), а не конкретный пакет
// func NewRootCmd(cfg *config.Config, transService TextTranslator) *cobra.Command {

// Здесь команда принимает конкретный пакет, а не интерфейс
func NewRootCmd(cfg *config.Config, transService *translate.Service) *cobra.Command {
	// В библиотеке Cobra вся логика консольного приложения строится вокруг структуры cobra.Command.
	// Когда пишем cmd := &cobra.Command{...}, то создается новый узел (команда или подкоманда) в дереве CLI-приложения.
	// Cobra активно модифицирует объект cobra.Command при работе: привязывает к нему флаги, регистрирует дочерние команды и отслеживает состояние выполнения.
	// По этому используем указатель (&).
	cmd := &cobra.Command{
		Use:   "translator [фраза на английском]",
		Short: "Translator — это простой CLI-переводчик с английского языка",
		// Проверяем, что передан ровно 1 аргумент (фраза для перевода)
		Args: cobra.ExactArgs(1),
		// hooks (?)
		// The *Run functions are executed in the following order:
		//   * PersistentPreRun()
		//   * PreRun()
		//   * Run()
		//   * PostRun()
		//   * PersistentPostRun()
		// RunE- основная бизнес-логика. E на конце- функция возвращает ошибку.
		RunE: func(cmd *cobra.Command, args []string) error {
			textToTranslate := args[0]

			// Viper автоматически выберет: флаг --to, если он есть.
			// Если флага нет — возьмет значение target_lang из YAML.
			targetLang := viper.GetString("to")

			result, err := transService.Translate(textToTranslate, targetLang)
			if err != nil {
				return fmt.Errorf("ошибка перевода: %w", err)
			}

			fmt.Printf("Оригинал (en): %s\n", textToTranslate)
			fmt.Printf("Перевод (%s): %s\n", targetLang, result)
			return nil
		},
	}

	// Флаги не объявляются внутри структуры &cobra.Command{}, они привязываются к ней сразу после её создания через методы:
	// • Локальные флаги (cmd.Flags()): доступны только этой конкретной команде.
	// • Глобальные флаги (cmd.PersistentFlags()): доступны этой команде и всем её подкомандам (например, флаг --verbose).
	// Добавляем локальный флаг --to (или -t). Значение по умолчанию берем из уже загруженного YAML-конфига
	cmd.Flags().StringP("to", "t", cfg.TargetLang, "Целевой язык для перевода (например: ru, es, fr)")

	// Связываем флаг Cobra с системой Viper
	_ = viper.BindPFlag("to", cmd.Flags().Lookup("to"))

	return cmd
}
