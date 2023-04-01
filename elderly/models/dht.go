package models

import "time"

// Dht 温湿度列表结构体
type Dht struct {
	Humidity    int       `json:"humidity" db:"humidity"`
	Temperature int       `json:"temperature" db:"temperature"`
	SenseTime   time.Time `json:"sense_time" db:"sense_time"`
}
