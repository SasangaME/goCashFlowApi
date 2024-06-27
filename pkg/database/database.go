package database

import (
	"fmt"
	"github.com/SasangaME/goCashFlowApi/pkg/model/entity"
	"github.com/SasangaME/goCashFlowApi/pkg/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func Connect() {
	p, err := util.Config("DB_PORT")
	if err != nil {
		log.Fatal(err)
	}
	port, err := strconv.ParseInt(p, 10, 64)

	host, err := util.Config("DB_HOST")
	if err != nil {
		log.Fatal(err)
	}

	dbUser, err := util.Config("DB_USER")
	if err != nil {
		log.Fatal(err)
	}

	dbPassword, err := util.Config("DB_PASSWORD")
	if err != nil {
		log.Fatal(err)
	}

	dbName, err := util.Config("DB_NAME")
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable "+
		"TimeZone=Asia/Shanghai", host, dbUser,
		dbPassword, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("connected to the database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	err = db.AutoMigrate(&entity.Role{})
	if err != nil {
		return
	}
	Database = DbInstance{
		Db: db,
	}
}
