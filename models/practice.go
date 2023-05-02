package models

type Practice struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	I18NEs string `json:"i18n_es" gorm:"column:i18n_es"`
}
