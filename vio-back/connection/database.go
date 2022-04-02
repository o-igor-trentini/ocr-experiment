package connection

import (
	"fmt"
	"log"
	"os"
	"time"
	"vio-back/appconst"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func OpenVioConnection() {
	loggerConfig := logger.Default.LogMode(logger.Silent)

	if os.Getenv(string(appconst.GIN_MODE)) != "release" {
		loggerConfig = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		)
	}

	host := os.Getenv(string(appconst.PG_HOST))
	user := os.Getenv(string(appconst.PG_USER))
	pass := os.Getenv(string(appconst.PG_PASS))
	name := os.Getenv(string(appconst.PG_NAME))
	port := os.Getenv(string(appconst.PG_PORT))
	schema := os.Getenv(string(appconst.PG_SCH))

	url := "host=%s user=%s password=%s dbname=%s port=%s search_path=%s"
	strConnection := fmt.Sprintf(url, host, user, pass, name, port, schema)

	db, err := gorm.Open(postgres.Open(strConnection),
		&gorm.Config{
			Logger: loggerConfig,
		},
	)
	if err != nil {
		log.Fatal("erro ao carregar as variáveis de ambiente")
	}

	DB = db
}

func OpenLogConnection() *gorm.DB {
	loggerConfig := logger.Default.LogMode(logger.Silent)

	if os.Getenv(string(appconst.GIN_MODE)) != "release" {
		loggerConfig = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		)
	}

	host := os.Getenv(string(appconst.PG_HOST))
	user := os.Getenv(string(appconst.PG_USER))
	pass := os.Getenv(string(appconst.PG_PASS))
	name := os.Getenv(string(appconst.PG_NAME))
	port := os.Getenv(string(appconst.PG_PORT))
	schema := os.Getenv(string(appconst.PG_SCH_LOG))

	url := "host=%s user=%s password=%s dbname=%s port=%s search_path=%s"
	dsn := fmt.Sprintf(url, host, user, pass, name, port, schema)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: loggerConfig})
	if err != nil {
		log.Fatal("erro ao abrir conexão de log")
	}

	return db
}
