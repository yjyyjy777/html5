package main
import(
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
//UserInfo 是学生表单.
type UserInfo struct{
	gorm.Model
	Name string 
	Gender string 
	Hobby string 
}
func main(){
	db,err:=gorm.Open("mysql", "root:wj770826@(192.168.33.10)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err !=nil{
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&UserInfo{})
	// u2 :=UserInfo{Name: "zxy",}
	// db.Create(&u2)
	// u2 :=UserInfo(Name:"zxy",Gender:"female")
	// db.Create(&u2)
	var user UserInfo
	db.Where("name = ?","123").First(&user)
	fmt.Println(user)
	


}