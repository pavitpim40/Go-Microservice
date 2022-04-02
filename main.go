package main

// const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
import (
	"fmt"
	"golang-project/repository"
	"golang-project/service"
	"log"
	"runtime"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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

func NewMySQL() *gorm.DB {
	cfg := &mysql.Config{
	 Net:                  "tcp",
	 Addr:                 fmt.Sprintf("%v:%v", viper.GetString("database.host"), viper.GetString("database.port")),
	 DBName:               viper.GetString("database.dbname"),
	 User:                 viper.GetString("database.user"),
	 Passwd:               viper.GetString("database.pass"),
	 AllowNativePasswords: true,
	}
	dbClient, err := gorm.Open("mysql", cfg.FormatDSN())
	if err != nil {
	 log.Fatal(err)
	}
	return dbClient
   }

func main(){
	var (
		e = echo.New()
		database = NewMySQL()
	)
	defer database.Close()
	ser := service.NewHandle(service.NewService(repository.New(database)))
	e.GET("/", ser.CallLogin)
 
	e.Logger.Fatal(e.Start(":" + viper.GetString("app.port")))
}