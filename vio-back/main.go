package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"vio-back/appconst"
	"vio-back/connection"
	"vio-back/connection/migration"
	"vio-back/routes"
	"vio-back/services/violog"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	initEnv()

	db := connection.OpenLogConnection()
	violog.InitLogger(db)
	initLocalTime()
	initConnections()
}

func main() {
	argumentCommands()
}

func argumentCommands() {
	if len(os.Args) <= 1 {
		runServerHttp()
	}

	for _, arg := range os.Args {
		switch arg {
		case "-m", "--migrate":
			migration.Migrate()

		case "-cm", "--create-migration":
			migration.CreateMigrationFile()

		case "-cs", "--create-structure":
			migration.CreateDatabaseStructure()

		case "--http":
			runServerHttp()

		case "--https":
			runServerHttps()
		}
	}
}

func runServerHttp() {
	migration.Migrate()

	r := gin.New()
	r.Use(gin.Recovery())

	if os.Getenv(string(appconst.GIN_MODE)) != "release" {
		r.Use(gin.Logger())
	}

	routes.Routes(connection.DB, r)

	err := r.Run(
		fmt.Sprintf(
			"%s:%s",
			os.Getenv(string(appconst.GIN_HOST)),
			os.Getenv(string(appconst.GIN_PORT)),
		),
	)
	if err != nil {
		violog.InfoLogger.Printf("erro ao inicilair o GIN; erro: %s", err)
	}
}

func runServerHttps() {
	migration.Migrate()

	r := gin.Default()
	routes.Routes(connection.DB, r)

	s := &http.Server{
		Addr:           ":" + os.Getenv(string(appconst.GIN_PORT)),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServeTLS("localhost.pem", "localhost-key.pem")
	if err != nil {
		fmt.Println(err)
	}
}

func initEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("erro ao carregar as variáveis de ambiente; erro: %s", err)
	}
}

func initConnections() {
	connection.OpenVioConnection()
	connection.OpenLogConnection()
}

func initLocalTime() {
	var err error
	time.Local, err = time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Fatalf("erro ao inicializar o horário local; erro:%s", err)
	}
}
