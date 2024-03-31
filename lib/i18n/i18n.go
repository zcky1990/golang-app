package i18n

import (
	"fmt"
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

func NewI18n(filepath, locale string) *I18n {
	trans, err := LoadTranslations(filepath)
	if err != nil {
		return nil
	}
	return &I18n{
		Translations: trans,
		Locale:       locale,
	}
}

// LoadTranslations loads translations from YAML files
func LoadTranslations(filepath string) (*translations, error) {
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
		return "", fmt.Errorf("translations not loaded")
	}
	translation, ok := (*i18n.Translations)[i18n.Locale]
	if !ok {
		return "", fmt.Errorf("translation not found for locale: %s", i18n.Locale)
	}
	message, ok := translation[messageID]
	if !ok {
		return "", fmt.Errorf("message not found with ID: %s", messageID)
	}
	return message, nil
}
