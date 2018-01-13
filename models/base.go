package models

import (
	"fmt"
	"net/url"
	"github.com/jinzhu/gorm"
	"os"
	"time"
	"log"
	"crypto/md5"
	"encoding/hex"
	"database/sql"

	"go-wyy/service/conf"
)

var DB *gorm.DB



var dbconf *conf.DbConf
var db_host,db_port,db_name,db_user,db_pass string
func init() {

	dbconf, err := dbconf.Load("database.json")
	if err != nil {
		fmt.Println(err)
	}
	db_host = dbconf.DbHost
	db_port = dbconf.DbPort
	db_name = dbconf.DbName
	db_user = dbconf.DbUser
	db_pass = dbconf.DbPass
}

func SyncDB() {
	createDB()
	Connect()
	DB.
		Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
		&AdminUser{},
		&Commentt{},
		&Comments{},
		&HotComments{},
		&Song{},
		&User{},
	)
}
func AddAdmin() {
	var user AdminUser
	fmt.Println("please input username for system administrator")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println("please input password for system administrator")
	var password string
	fmt.Scanf("%s", &password)
	user.Username = name
	h := md5.New()
	h.Write([]byte(password))
	user.Password = hex.EncodeToString(h.Sum(nil))
	if err := DB.Create(&user).Error; err != nil {
		fmt.Println("admin create error,please run this application again")
		os.Exit(0)
	} else {
		fmt.Println("admin create finished")
	}
}

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s&parseTime=true",
		db_user,
		db_pass,
		db_host,
		db_port,
		db_name,
		url.QueryEscape("Asia/Shanghai"))

	var err error

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Print("master detabase connect error:", err)
		os.Exit(0)
	}

	DB.SingularTable(true)
	DB.DB().SetMaxOpenConns(2000)
	DB.DB().SetMaxIdleConns(100)
	DB.DB().SetConnMaxLifetime(100 * time.Nanosecond)
}

func createDB() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&loc=%s&parseTime=true", db_user, db_pass, db_host, db_port, url.QueryEscape("Asia/Shanghai"))
	sqlstring := fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8mb4 COLLATE utf8mb4_general_ci", db_name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		log.Println(r)
	} else {
		log.Println("Database ", db_name, " created")
	}
	defer db.Close()

}
