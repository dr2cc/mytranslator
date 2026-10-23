package translate

import (
	"github.com/bregydoc/gtranslate"
)

// • Stateless Service (Сервис без состояния).
// Смысл - предоставление методов.
// Цель - инкапсуляция методов (поведения) отдельно от данных.
// Это объект, который предоставляет определенную бизнес-логику (сервис), но не хранит никаких данных между вызовами.
// Каждый запрос к нему полностью независим.
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
