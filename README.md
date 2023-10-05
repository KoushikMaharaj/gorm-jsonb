# gorm-jsonb
[gorm](https://gorm.io) jsonb datatype

In this package you can store json as well as json array in database
****************************************************************
# install
```
go get github.com/KoushikMaharaj/gorm-jsonb
```
****************************************************************
# example
```go
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

```
****************************************************************
# Reference
[dariubs](https://github.com/dariubs/gorm-jsonb)

