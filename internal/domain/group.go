package domain

type Group struct {
	Name          string          `json:"name" db:"name"`
	Users         []User          `json:"users" db:"users"`
	Tasks         []Task          `json:"tasks" db:"tasks"`
	LevelProgress []LevelProgress `json:"levelProgress" db:"level_progress"`
}
