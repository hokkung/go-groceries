package domain

type Name struct {
	Primary     string `gorm:"column:primary_name"`
	ThaiName    string `gorm:"column:name_thai"`
	EnglishName string `gorm:"column:name_english`
}
