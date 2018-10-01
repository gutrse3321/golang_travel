package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	Name string
	Phone string
}

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		fmt.Println("server disconnect!")
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// 链接test数据库，查询people表，如果没有就会创建
	c := session.DB("test").C("people")

	// 插入数据
	//err = c.Insert(
	//	&Person{"tomo", "+086 189 8016 5759"},
	//	&Person{"yuki", "+086 130 8830 4812"},
	//	)
	//if err != nil {
	//	fmt.Println("Insert error!")
	//	log.Fatal(err)
	//}

	result := Person{}
	// 修改
	// 修改的时候，需要设置每个字段的值，不然就是空值，如下只设置了name的值，然而Phone没值
	// 如果只想修改name，这需要在bson加个$set：bson.M{"$set": bson.M{"name": "tomonori"}}
	//err = c.Update(bson.M{"name": "tomo"}, bson.M{"name": "tomonori"})
	//if err != nil {
	//	fmt.Println("Update error!")
	//	log.Fatal(err)
	//}
	//origYuki := bson.M{"name": "yuki"}
	//upYuki := bson.M{"$set": bson.M{"name": "yuki-chan"}}
	//err = c.Update(origYuki, upYuki)

	// 删除
	err = c.Remove(bson.M{"name": "tomonori"})
	if err != nil {
		fmt.Println("Remove error!")
		log.Fatal(err)
	}

	// 查询
	err = c.Find(bson.M{"name": "tomonori"}).One(&result)
	if err != nil {
		fmt.Println("Find error!")
		log.Fatal(err)
	}
	//err = c.Find(bson.M{"name": "yuki-chan"}).One(&result)
	fmt.Println("Phone", result.Phone)
}
