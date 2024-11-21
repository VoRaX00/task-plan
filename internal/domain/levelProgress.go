package domain

type LevelProgress struct {
	Name    string
	GroupID uint `gorm:"many2many:group_level_progress;"`
}
