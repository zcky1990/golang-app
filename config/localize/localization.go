package localize

import (
	"fmt"
	c "golang_app/golangApp/constants"
	"os"

	"gopkg.in/yaml.v3"
)

// Translation represents a translation map
type Translation map[string]string

// translations maps language codes to translation maps
type translations map[string]Translation

// I18n represents internationalization settings
type I18n struct {
	Translations *translations
	Locale       string
}

func newI18n(filepath, locale string) *I18n {
	trans, err := loadTranslations(filepath)
	if err != nil {
		return nil
	}
	return &I18n{
		Translations: trans,
		Locale:       locale,
	}
}

// LoadTranslations loads translations from YAML files
func loadTranslations(filepath string) (*translations, error) {
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var translations translations
	err = yaml.Unmarshal(yamlFile, &translations)
	if err != nil {
		return nil, err
	}
	return &translations, nil
}

// GetMessage retrieves translated message for a given message ID
func (i18n *I18n) GetMessage(messageID string) (string, error) {
	if i18n.Translations == nil {
		return "", fmt.Errorf("Translations not loaded")
	}
	translation, ok := (*i18n.Translations)[i18n.Locale]
	if !ok {
		return "", fmt.Errorf("Translation not found for locale: %s", i18n.Locale)
	}
	message, ok := translation[messageID]
	if !ok {
		return "", fmt.Errorf("Message not found with ID: %s", messageID)
	}
	return message, nil
}

func (i18n *I18n) GetMessageWithLocale(messageID string, locale string) (string, error) {
	if i18n.Translations == nil {
		return "", fmt.Errorf("Translations not loaded")
	}
	translation, ok := (*i18n.Translations)[locale]
	if !ok {
		return "", fmt.Errorf("Translation not found for locale: %s", locale)
	}
	message, ok := translation[messageID]
	if !ok {
		return "", fmt.Errorf("Message not found with ID: %s", messageID)
	}
	return message, nil
}

type Localization struct {
	I18n *I18n
}

func NewLocalization() *Localization {
	locale := newI18n("i18n.yml", c.LOCALE_INDONESIA)
	return &Localization{I18n: locale}
}

func (c *Localization) GetLocalizationMessage(messageKey string) string {
	message, _ := c.I18n.GetMessage(messageKey)
	return message
}

func (c *Localization) GetLocalizationMessageWithLocale(messageKey string, locale string) string {
	message, _ := c.I18n.GetMessageWithLocale(messageKey, locale)
	return message
}
