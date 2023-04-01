package mysql

import (
	"database/sql"
	"elderly/models"

	"go.uber.org/zap"
)

func GetLedStateListHandler() (ledState []*models.Led, err error) {
	sqlStr := `select vsmooth,vlight,states,sense_time from smooth_ledstatus order by sense_time DESC limit 10 `
	if err = db.Select(&ledState, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no led state in db")
			err = nil
		}
	}
	return
}
