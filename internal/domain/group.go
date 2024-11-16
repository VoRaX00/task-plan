package domain

type Group struct {
	Name          string   `json:"name"`
	Users         []int64  `json:"users"`
	Tasks         []Task   `json:"tasks" gorm:"foreignkey:GroupId"`
	LevelProgress []string `json:"levelProgress"`
}
