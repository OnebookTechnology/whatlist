package models

type Biggie struct {
	Id           int           `json:"id"`
	Name         string        `json:"name"`
	Identity     string        `json:"identity"`
	Intro        string        `json:"intro"`
	Sendword     string        `json:"sendword"`
	Weight       uint8         `json:"weight"`
	Signtime     string        `json:"signtime"`
	Image        string        `json:"image"`
	LatestListId int           `json:"latest_list_id"`
	Lists        []*BiggieList `json:"lists,omitempty"`
}
