package mysql

import (
	"database/sql"
	"elderly/models"

	"go.uber.org/zap"
)

func GetDhtList() (dhtList []*models.Dht, err error) {
	sqlStr := `select humidity,temperature,sense_time from dht22 order by sense_time DESC limit 10`
	if err = db.Select(&dhtList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no dht in db")
			err = nil
		}
	}
	return
}
