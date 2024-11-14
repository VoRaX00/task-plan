package domain

type Group struct {
	Id            int64    `json:"id" db:"id"`
	Name          string   `json:"name" db:"name"`
	Users         []int64  `json:"users" db:"users"`
	Tasks         []Task   `json:"tasks" db:"tasks"`
	LevelProgress []string `json:"levelProgress" db:"level_progress"`
}
