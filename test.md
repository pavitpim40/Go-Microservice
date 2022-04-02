```go
package main


import (
	"fmt"
	"log"
	"runtime"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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


   func isDoubleArray( arr []int) string{

    dict:= make(map[int]int)
    for _ , num :=  range arr {
		if (dict[num] < 2){
			dict[num] = dict[num]+1
		} else {
			return "N"
		}
        
    }
    return "Y"
}

func sumWithRange( i int , j int,k int) int {
	sum := 0
	inc := 0
	

	for  true {
		if (i + inc < j) {
			sum += i + inc
			inc++
			continue
		} else {
			break
		}
	}
	inc = 0
	for true {
		if(j - inc > k) {
			sum += j - inc
			inc++
			continue
		} else {
			sum += k;
			break
		}
	}
	return sum
}

func caesarDe(strCipher string, step_move byte) string {
   
    // str_cipher := strings.ToLower(strCipher)
    str_slice_src := []byte(strCipher)
    str_slice_dst := make([]byte, len(str_slice_src), len(str_slice_src))

 
    for i := 0; i < len(str_slice_src); i++ {
    
        if str_slice_src[i] >= 97+step_move {
            str_slice_dst[i] = str_slice_src[i] - step_move
        } else {
            str_slice_dst[i] = str_slice_src[i] + 26 - step_move
        }
	}
    fmt.Println ("ciphertext:", strCipher, str_slice_src)
    fmt.Println ("plaintext:", string (str_slice_dst), str_slice_dst)
    return string(str_slice_dst)

}


const lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func rotateText(inputText string, rot int) string {
	rot %= 26
	rotatedText := []byte(inputText)

	for index, byteValue := range rotatedText {
		if byteValue >= 'a' && byteValue <= 'z' {
			rotatedText[index] = lowerCaseAlphabet[(int((26+(byteValue-'a')))-rot)%26]
		} else if byteValue >= 'A' && byteValue <= 'Z' {
			rotatedText[index] = upperCaseAlphabet[(int((26+(byteValue-'A')))-rot)%26]
		}
	}
	return string(rotatedText)
}
func main(){

	// testInput := []int {5, 5, 2, 2, 3, 3,3} 
	// fmt.Println(isDoubleArray(testInput))


	fmt.Println(rotateText("vtaog",2))
	i,j := 42,2701
	p := &j
	p = &i
	fmt.Println(*p)

}
```