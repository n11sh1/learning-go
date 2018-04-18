package main

import (
	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/postgres"
	"path/filepath"
	"github.com/BurntSushi/toml"
	"log"
)

type DbConfig struct {
	Posgre Posgre `toml:"Posgre"`
}

type Posgre struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
	User string `toml:"user"`
	Dbname string `toml:"dbname"`
	Password string `toml:"password"`
}

type Product struct {
	gorm.Model
	Code string
	Price uint
}

var config DbConfig
var DB_CONFIG string

func init()  {
	abs_path, _ := filepath.Abs(".")
	_, err := toml.DecodeFile(filepath.Join(abs_path, "sandbox/db/gorm_posgre/db_config.toml"), &config)
	if err != nil {
		log.Fatal(err)
	}
	DB_CONFIG = "host=" + config.Posgre.Host +
		" port=" + config.Posgre.Port +
		" user=" + config.Posgre.User +
		" dbname=" + config.Posgre.Dbname +
		" password=" + config.Posgre.Password +
		" sslmode=disable"
}


func main() {
	db, err := gorm.Open("postgres", DB_CONFIG)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// #MIGRATE
	db.AutoMigrate(&Product{})

	// #INSERT
	product1 := Product{Code: "aaa", Price: 100}
	flag := db.NewRecord(product1)
	log.Printf("new record flag: %s", flag)
	db.Create(&product1)

	// #SELECT
	var product Product
	db.First(&product)
	log.Printf("first: %s", product)
}
