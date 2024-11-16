package domain

type Task struct {
	Name        string `json:"name" gorm:"type:varchar(100)"`
	Description string `json:"description" gorm:"type:text"`
	Level       string `json:"level" gorm:"type:varchar(50)"`
}
