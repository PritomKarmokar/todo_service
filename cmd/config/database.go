package config

import (
	"fmt"
	gommonLog "github.com/labstack/gommon/log"
	"todo_service/cmd/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// db connection
var db *gorm.DB

func getDsn() string {
	dbName := viper.GetString("MASTER_DB_NAME")
	dbUser := viper.GetString("MASTER_DB_USER")
	dbPassword := viper.GetString("MASTER_DB_PASSWORD")
	dbHost := viper.GetString("MASTER_DB_HOST")
	dbPort := viper.GetString("MASTER_DB_PORT")
	dbSSLMode := viper.GetString("MASTER_SSL_MODE")
	dbTimeZone := viper.GetString("TIME_ZONE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
		dbSSLMode,
		dbTimeZone,
	)

	return dsn
}

func ConnectDB() {
	var err error
	dsn := getDsn()
	slowSQLThreshold := viper.GetDuration("SLOW_SQL_THRESHOLD") * time.Millisecond

	newLogger := logger.New(
		log.New(os.Stdout, "Slow SQL Log: ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             slowSQLThreshold,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Panic(err)
	}

	// apply AutoMigrate here
	db.AutoMigrate(&models.Todo{})

	sqlDb, err := db.DB()
	if err != nil {
		log.Panic(err)
	}
	sqlDb.SetMaxIdleConns(viper.GetInt("SET_MAX_IDLE_CONNECTIONS"))
	sqlDb.SetMaxOpenConns(viper.GetInt("SET_MAX_OPEN_CONNECTIONS"))
	sqlDb.SetConnMaxIdleTime(viper.GetDuration("SET_CONNECTION_MAX_IDLE_TIME"))
	sqlDb.SetConnMaxLifetime(viper.GetDuration("SET_CONNECTION_MAX_LIFE_TIME"))

	err = sqlDb.Ping()
	if err != nil {
		log.Panic(err)
	}

	gommonLog.Info("DB connection established")
}

func GetDatabase() *gorm.DB {
	return db
}
