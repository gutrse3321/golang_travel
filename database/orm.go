package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spaolacci/murmur3"
	"strings"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Age          int64 `gorm:"null"`
	Birthday     time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`
	MemberNumber *string `gorm:"unique;not null"`
	Num          int     `gorm:"AUTO_INCREMENT"`
	Address      string  `gorm:"index:addr"`
	IgnoreMe     int     `gorm:"-"`
}

type Like struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64 `gorm:"unique_index:hash_idx"`
	CreatedAt time.Time
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/go_test?charset=utf8")
	catch(err)
	// 启用Logger，显示详细日志
	db.LogMode(true)
	defer db.Close()

	// 全局设置表名不可以为复数形式
	db.SingularTable(true)

	// 判断是否存在表，接收字符串，或地址
	if db.HasTable("user") == false {
		db.CreateTable(&User{})
	}

	// 通过 db.Set 设置一些额外的表属性
	if !db.HasTable(&Like{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}).Error; err != nil {
			catch(err)
		}
	}
	// 插入
	like := &Like{
		Ip: "192.168.1.1",
		Ua: "fucking shit",
		Title: "fuck off",
		Hash: murmur3.Sum64([]byte(strings.Join([]string{"192.168.1.1", "fucking shit", "fuck off"}, "-"))) >> 1,
		CreatedAt: time.Now(),
	}
	//
	//if err := db.Create(like).Error; err != nil {
	//	catch(err)
	//}

	// 删除
	//if err := db.Where(&Like{Hash: 7285552096162104610}).Delete(Like{}).Error; err != nil {
	//	catch(err)
	//}

	//查询
	// db.Model() 选择一个表，再用 db.Where() 构造查询条件，后面可以使用 db.Count() 计算数量
	// 如果要获取对象，可以使用 db.Find(&Likes) 或者只需要查一条记录 db.First(&Like)
	var count int
	selectErr := db.Model(&Like{}).Where(&Like{Ip: "192.168.1.1",Ua: "fucking shit",Title: "fuck off"}).Count(&count).Error
	item := db.First(&Like{})
	if selectErr != nil {
		catch(selectErr)
	}
	fmt.Println(item)

	// 修改
	//// 查询第一条，修改他，提交
	//	//db.First(&like)
	//	//like.Ip = "168.122.14.5"
	//	//db.Save(&like)

	// Model() 选择一张表的方式
	//db.Model(&like).Update("Ua", "yes tomo")
	// 还可以使用Where增加查询条件
	//db.Model(&like).Where("Ip = ?", "168.122.14.5").Update("Title", "update Title")
	//
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
