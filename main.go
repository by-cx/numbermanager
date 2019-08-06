package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var config Config

func init() {
	var err error

	// When test mode do practically nothing
	if flag.Lookup("test.v") != nil {
		fmt.Fprintln(os.Stderr, "Testing mode engaged!")
	} else {
		// Parsing config from environment variables
		err = envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err.Error())
		}

		// Establish the database connection
		if config.DbType == "sqlite" {
			db, err = gorm.Open("sqlite3", "gorm.db")
			if err != nil {
				log.Fatalln("Database connection issue:", err)
			}
		} else if config.DbType == "mysql" {
			db, err = gorm.Open("mysql", config.DbUser+":"+config.DbPassword+"@"+config.DbHost+"/"+config.DbName+"?charset=utf8&parseTime=True&loc=Local")
			if err != nil {
				log.Fatalln("Database connection issue:", err)
			}
		} else {
			log.Fatalln("Unsupported database engine (" + config.DbType + ")")
		}
	}

}

func main() {
	// Close connection to the database at the end of the program
	defer db.Close()

	// Database's tables migration
	db.AutoMigrate(&Number{})

	// API
	e := echo.New()

	e.Static("/files", "files")
	e.Static("/bower_components", "bower_components")

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `
# Number manager

This is a simple API that holds value of numbers. It's great for keeping track of builds in CIs. It can set new values, increment or decrement them. The API is pretty simple and doesn't use any kind of authentication.

## API

### Registration

	POST /api/number/register
	-- no input --

Returns ID of a new number you need to access its value in other API endpoints.

### Set

	POST /api/number/:id
	{"number": 99}

Sets specific value to the number.

### Get

	GET /api/number/:id

Returns value of the number.

### Increment

	POST /api/number/:id/incr
	-- no input --

Increment the number by one.

### Decrement

	POST /api/number/:id/decr
	-- no input --

Decrement the number by one.
		`)
	})

	e.POST("/api/number/register", func(c echo.Context) error {
		number, err := NewNumber()
		if err != nil {
			return c.JSONPretty(http.StatusInternalServerError, map[string]string{"error": err.Error()}, "  ")
		}

		return c.JSONPretty(http.StatusOK, ResponseAlwaysNumber{
			Message: "ok",
			Number:  number.Number,
			ID:      number.ID,
		}, "  ")
	})
	e.POST("/api/number/:id", func(c echo.Context) error {
		number, err := Find(c.Param("id"))
		if err != nil {
			return c.JSONPretty(http.StatusInternalServerError, Response{
				Message: "error occured while searching for your number",
				Error:   err.Error(),
			}, "  ")
		}

		var value = NumberValue{}

		err = c.Bind(&value)
		if err != nil {
			return c.JSONPretty(http.StatusInternalServerError, Response{
				Message: "error occurred while parsing the input value",
				Error:   err.Error(),
			}, "  ")
		}

		err = number.Set(value.Number)
		if err != nil {
			return c.JSONPretty(http.StatusInternalServerError, Response{
				Message: "error occured while setting the new value to the number",
				Error:   err.Error(),
			}, "  ")
		}

		return c.JSONPretty(http.StatusOK, ResponseAlwaysNumber{
			Message: "ok",
			Number:  number.Number,
		}, "  ")
	})
	e.GET("/api/number/:id", func(c echo.Context) error {
		number, err := Find(c.Param("id"))
		if err != nil {
			return c.JSONPretty(http.StatusInternalServerError, Response{
				Message: "error occured while searching for your number",
				Error:   err.Error(),
			}, "  ")
		}
		return c.JSONPretty(http.StatusOK, ResponseAlwaysNumber{
			Message: "ok",
			Number:  number.Number,
		}, "  ")
	})
	e.GET("/api/number/:id/incr", func(c echo.Context) error {
		number, err := Find(c.Param("id"))
		if err != nil {
			return c.JSONPretty(http.StatusInternalServerError, Response{
				Message: "error occured while searching for your number",
				Error:   err.Error(),
			}, "  ")
		}

		err = number.Incr()
		if err != nil {
			return c.JSONPretty(http.StatusInternalServerError, Response{
				Message: "error occured while incrementing the number",
				Error:   err.Error(),
			}, "  ")
		}

		return c.JSONPretty(http.StatusOK, ResponseAlwaysNumber{
			Message: "ok",
			Number:  number.Number,
		}, "  ")
	})
	e.GET("/api/number/:id/decr", func(c echo.Context) error {
		number, err := Find(c.Param("id"))
		if err != nil {
			return c.JSONPretty(http.StatusInternalServerError, Response{
				Message: "error occured while searching for your number",
				Error:   err.Error(),
			}, "  ")
		}

		err = number.Decr()
		if err != nil {
			return c.JSONPretty(http.StatusInternalServerError, Response{
				Message: "error occured while decrementing the number",
				Error:   err.Error(),
			}, "  ")
		}

		return c.JSONPretty(http.StatusOK, ResponseAlwaysNumber{
			Message: "ok",
			Number:  number.Number,
		}, "  ")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
