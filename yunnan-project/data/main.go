package main

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("49.4.80.188:27017")
	if err != nil {
		fmt.Println(111)
		panic(err)
	}
	defer session.Close()

	c := session.DB("ZYC-2018").C("YP_饮片赋码信息表")
	i, _ := c.Count()
	fmt.Println(i)
	var d []MedTagging
	c.Find(nil).All(&d)
	fmt.Println(d)
	fmt.Println(time.Now())

}
