package database

import (
	"fmt"
	"net/url"

	"github.com/coronasafe/kerala-map-api/config"
	"github.com/coronasafe/kerala-map-api/model"

	"github.com/jinzhu/gorm"
)

// ConnectDB connect to db
func ConnectDB() {
	var postgresURL string
	var err error
	if config.Config.DB.URL == "" {
		dsn := url.URL{
			User:     url.UserPassword(config.Config.DB.User, config.Config.DB.Password),
			Scheme:   "postgres",
			Host:     fmt.Sprintf("%s:%d", config.Config.DB.Host, config.Config.DB.Port),
			Path:     config.Config.DB.Name,
			RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
		}
		postgresURL = dsn.String()
		fmt.Println("Using PSQL variables")
	} else {
		postgresURL = config.Config.DB.URL
		fmt.Println("Using PSQL url")
	}

	DB, err = gorm.Open("postgres", postgresURL)
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database")

	DB.AutoMigrate(&model.Description{}, &model.User{}, &model.Feature{})
	fmt.Println("Database migrated")
}
