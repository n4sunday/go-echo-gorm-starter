package db

import (
	"fmt"
	"main/config"
	"main/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()

	host := configuration.DB_HOST
	password := configuration.DB_PASSWORD
	port := configuration.DB_PORT
	database := configuration.DB_NAME
	user := configuration.DB_USERNAME

	connect_string := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", host, user, password, database, port)
	db, err = gorm.Open(postgres.Open(connect_string), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
}

func DbManager() *gorm.DB {
	return db
}

func Paginate(c echo.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		if page == 0 {
			page = 1
		}

		sort := c.QueryParam("sort")

		if sort == "" {
			sort = "desc"
		}

		order := c.QueryParam("order")
		if order == "" {
			order = "id"
		}

		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit).Order(fmt.Sprintf("%s %s", order, sort))
	}
}
