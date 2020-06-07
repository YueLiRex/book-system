package models

type Reservation struct {
	StartDatetime int64  `db:"startDatetime" json:"startDatetime"`
	Duration      int64  `json:"duration"`
	Location      string `json:"location"`
}
