package logic

import (
	"elderly/dao/mysql"
	"elderly/models"
)

func GetDhtList() ([]*models.Dht, error) {
	//从数据库查询所有的温湿度信息，并返回
	return mysql.GetDhtList()
}
