package config

import (
	"golang_app/golangApp/lib/i18n"
)

type Localization struct {
	locale *i18n.I18n
}

func NewLocalization() *Localization {
	locale := i18n.NewI18n("i18n.yml", "ind")
	return &Localization{locale: locale}
}

func (c *Localization) Localization(messageKey string) string {
	message, _ := c.locale.GetMessage(messageKey)
	return message
}
