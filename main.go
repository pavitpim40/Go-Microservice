package main

// const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
import (
	"golang-project/service"
	"log"
	"runtime"

	"github.com/labstack/echo/v4"
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


	e := echo.New()
	ser := service.NewHandle(service.NewService())
	e.GET("/", ser.CallLogin)
	e.Logger.Fatal(e.Start(":" + viper.GetString("app.port")))
}