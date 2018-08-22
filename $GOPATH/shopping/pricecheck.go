package shopping

// 导入的地址是$GOPATH/src/shopping/db
import (
	"shopping/db"
)

func PriceCheck(itemId int) (float64, bool) {
	item := db.LoadItem(itemId)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}
