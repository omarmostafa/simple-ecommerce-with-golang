package infrastructures

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
	"simple-ecommerce/app/Migrations"
	"sync"
)
type DatabaseConnection struct {
}

func NewDatabaseConnection()  *DatabaseConnection{
	if db == nil {
		ConnectionOnce.Do(func() {
			db = &DatabaseConnection{}
		})
	}
	return  db
}
func init()  {
	conn:=NewDatabaseConnection().GetDB()
	Migrations.CreateCategoriesTable(*conn)
	Migrations.CreateProductTable(*conn)
	Migrations.CreateUserTable(*conn)
}
func (connection *DatabaseConnection) GetDB() *gorm.DB {
	if dbConnection==nil{
		connOnce.Do(func() {
			_ = godotenv.Load()
			url :=fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("db_user"), os.Getenv("db_pass"),os.Getenv("db_name"))
			conn, _ := gorm.Open("mysql", url)
			dbConnection = conn
		})
	}

	return  dbConnection
}

var (
	db *DatabaseConnection
	ConnectionOnce sync.Once
	connOnce sync.Once
	dbConnection *gorm.DB
)

func MySQlConnection() *DatabaseConnection  {
	if db == nil {
		ConnectionOnce.Do(func() {
			db = &DatabaseConnection{}
		})
	}
	return  db
}