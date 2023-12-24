package main

import (
	"encoding/json"
	"fmt"
	"github.com/KoushikMaharaj/gorm-jsonb"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Data struct {
	Info    gorm_jsonb.JSON
	Address gorm_jsonb.JSONArray
}

func main() {
	var data Data
	data.Info = make(map[string]interface{})
	data.Info["first_name"] = "Amitabh"
	data.Info["last_name"] = "Khurana"
	data.Address = []map[string]interface{}{}
	data.Address = append(data.Address, map[string]interface{}{"office_address1": "Pune", "office_address2": "Mumbai"})
	data.Address = append(data.Address, map[string]interface{}{"home_address": "Pune"})
	postgresExample(data)
	mysqlExample(data)
}

func postgresExample(data Data) {
	url := "postgres://test:test@localhost:5432/test"
	fmt.Printf("url %s\n", url)
	dB, err := gorm.Open(postgres.Open(url))
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Successfully made postgres DB connection\n")

	err = dB.Create(&data).Error
	if err != nil {
		log.Fatalf("%v", err)
	}
	var fetchedData []Data

	result := dB.Raw("select * from data").Scan(&fetchedData)
	if result.Error != nil {
		log.Fatalf("%v", result.Error)
	}
	log.Printf("Rows affected: %v\n", result.RowsAffected)
	log.Printf("Fectched data: %v\n", fetchedData)
	marshal, _ := json.Marshal(&fetchedData)
	log.Printf("Fectched data in json format: %v\n", string(marshal))
}

func mysqlExample(data Data) {
	url := "root:test@tcp(localhost:3306)/test"
	dB, err := gorm.Open(mysql.Open(url))
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Successfully made mysql DB connection\n")

	err = dB.Create(&data).Error
	if err != nil {
		log.Fatalf("%v", err)
	}
	var fetchedData []Data
	result := dB.Raw("select * from data").Scan(&fetchedData)
	if result.Error != nil {
		log.Fatalf("%v", result.Error)
	}
	log.Printf("Rows affected: %v\n", result.RowsAffected)
	log.Printf("Fectched data: %v\n", fetchedData)
	marshal, _ := json.Marshal(&fetchedData)
	log.Printf("Fectched data in json format: %v\n", string(marshal))
}
