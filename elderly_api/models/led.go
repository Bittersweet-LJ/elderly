package models

import "time"

// Led 光感烟雾数据结构体
type Led struct {
	Vsmooth   int64     `json:"vsmooth" db:"vsmooth"`
	Vlight    int64     `json:"vlight" db:"vlight"`
	States    string    `json:"states" db:"states"`
	SenseTime time.Time `json:"sense_time" db:"sense_time"`
}
