package domain

type Group struct {
	Name          string   `json:"name"`
	Users         []User   `json:"users"`
	Tasks         []Task   `json:"tasks" gorm:"foreignkey:GroupId"`
	LevelProgress []string `json:"levelProgress"`
}
