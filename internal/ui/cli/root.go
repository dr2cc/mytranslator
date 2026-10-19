package cli

import (
	"fmt"

	"github.com/dr2cc/mytranslator/internal/config"
	"github.com/dr2cc/mytranslator/internal/translate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewRootCmd(cfg *config.Config, transService *translate.Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "translator [фраза на английском]",
		Short: "Translator — это простой CLI-переводчик с английского языка",
		// Проверяем, что передан ровно 1 аргумент (фраза для перевода)
		Args: cobra.ExactArgs(1),
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

	// Добавляем флаг --to (или -t). Значение по умолчанию берем из уже загруженного YAML-конфига
	cmd.Flags().StringP("to", "t", cfg.TargetLang, "Целевой язык для перевода (например: ru, es, fr)")

	// Связываем флаг Cobra с системой Viper
	_ = viper.BindPFlag("to", cmd.Flags().Lookup("to"))

	return cmd
}
