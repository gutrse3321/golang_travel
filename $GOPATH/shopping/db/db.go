// 创建一个包，注意包名和文件夹名是相同的
package db

import (
	"shopping/models"
)

func LoadItem(id int) *models.Item {
	return &models.Item{
		Price: 9.001,
	}
}
