package main

import (
	"fmt"
	"strconv"

	"github.com/yunnan-project/service"
)

func WriteData(i int, serviceSetup service.ServiceSetup) {
	fmt.Println("i_value:" + strconv.Itoa(i))
	edu := service.Education{
		Name:           strconv.Itoa(i),
		Gender:         strconv.Itoa(i),
		Nation:         strconv.Itoa(i),
		EntityID:       strconv.Itoa(i),
		Place:          strconv.Itoa(i),
		BirthDay:       strconv.Itoa(i),
		EnrollDate:     strconv.Itoa(i),
		GraduationDate: strconv.Itoa(i),
		SchoolName:     strconv.Itoa(i),
		Major:          strconv.Itoa(i),
		QuaType:        strconv.Itoa(i),
		Length:         strconv.Itoa(i),
		Mode:           strconv.Itoa(i),
		Level:          strconv.Itoa(i),
		Graduation:     strconv.Itoa(i),
		CertNo:         strconv.Itoa(i),
		Photo:          strconv.Itoa(i),
	}
	_, err := serviceSetup.SaveEdu(edu)
	// msg, err := serviceSetup.SaveEdu(edu)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// fmt.Println("信息发布成功, 交易编号为: " + msg)
	}
	ch <- 0
}
