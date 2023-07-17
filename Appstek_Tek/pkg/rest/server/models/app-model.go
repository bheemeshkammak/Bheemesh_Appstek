package models

type App struct {
	Id int64 `json:"id,omitempty"`

	App string `json:"app,omitempty"`

	Employee string `json:"employee,omitempty"`
}
