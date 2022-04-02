package main

// const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


func initViper(){
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
	log.Fatalf("cannot read in viper config: %s", err)
	}
}

func init(){
	runtime.GOMAXPROCS(1)
	initViper()
}
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
	fmt.Println("hello : ", viper.GetString("app.name"))
	r.Run(":" + viper.GetString("app.port"))
}