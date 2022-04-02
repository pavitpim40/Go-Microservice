package main

import (
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main(){



	
	value := gjson.Get(json, "name.last")
	print(value.String())

	value2, _ := sjson.Set(json, "name.last", "Pimchanagul")
	println(value2)
}