package main

import (
	"fmt"
	"github.com/KoushikMaharaj/gorm-jsonb"
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

	fmt.Println(data)
}
