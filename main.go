package main

// const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func main(){

	// value := gjson.Get(json, "name.last")
	// print(value.String())

	// value2, _ := sjson.Set(json, "name.last", "Pimchanagul")
	// println(value2)

	r := gin.Default();
	r.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	r.Run(":8080")
}