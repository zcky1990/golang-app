package config

import (
	"golang_app/golangApp/lib/i18n"
)

var localization *i18n.I18n

func init() {
	localization = i18n.NewI18n("i18n.yml", "ind")
}

func GetInstance() *i18n.I18n {
	return localization
}
