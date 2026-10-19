package translate

import (
	"github.com/bregydoc/gtranslate"
)

// Service отвечает за перевод текста.
type Service struct{}

func NewService() *Service {
	return &Service{}
}

// Translate переводит текст с английского ("en") на целевой язык.
func (s *Service) Translate(text, targetLang string) (string, error) {
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: "en",
			To:   targetLang,
		},
	)
	if err != nil {
		return "", err
	}
	return translated, nil
}
