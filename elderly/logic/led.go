package logic

import (
	"elderly/dao/mysql"
	"elderly/models"
)

func GetLedStateListHandler() ([]*models.Led, error) {
	return mysql.GetLedStateListHandler()
}
