package migration

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"time"
	"vio-back/connection"
	"vio-back/helpers"
	"vio-back/models"
	"vio-back/services/violog"

	"gorm.io/gorm"
)

type MigrationExecute struct {
	Name string
	Run  func(*gorm.DB) error
}

func Migrate() {
	migrationsToExec := []interface{}{}

	for _, function := range migrationsToExec {
		f := reflect.ValueOf(function)
		result := f.Call(nil)
		migration := result[0].Interface().(MigrationExecute)

		actualMigration := models.Migration{
			Name: migration.Name,
		}

		err := connection.DB.
			Where("mi_name", migration.Name).
			First(&actualMigration).
			Error
		if err != nil {
			if !helpers.IsGormRecordNotFoundError(err) {
				fmt.Println(err)
				continue
			}
		}

		if actualMigration.ID > 0 {
			fmt.Println("### Migration já executada")
			continue
		}

		if err := migration.Run(connection.DB); err != nil {
			continue
		}

		connection.DB.Create(&actualMigration)
	}
}

// func createEnum(name, types string) string {
// 	return fmt.Sprintf(`
// 		DO $$
// 		BEGIN
// 			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typename = '%s') THEN
// 				CREATE TYPE %s AS ENUM (%s);
// 			END IF;
// 		END
// 		$$
// 	`,
// 		name,
// 		name,
// 		types,
// 	)
// }

func CreateDatabaseStructure() {
	migrateLogs()
	preMigration()
	migrateModels()
}

func preMigration() {
	violog.InfoLogger.Print("criando os enums das tabelas")

	connection.DB.Exec("CREATE SCHEMA IF NOT EXISTS vio")

	// connection.DB.Exec(createEnum())
}

func migrateModels() {
	err := connection.DB.AutoMigrate(
		&models.Migration{},
	)
	if err != nil {
		violog.InfoLogger.Printf("erro ao executar as migrações; erro:%s", err)
	}

	// adicionar unique keys
}

func migrateLogs() {
	db := connection.OpenLogConnection()

	db.Exec("CREATE SCHEMA IF NOT EXISTS log")

	if err := db.AutoMigrate(violog.LogError{}); err != nil {
		fmt.Printf("erro ao executar migração do logger err:%" + err.Error())
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func CreateMigrationFile() {
	timeNow := time.Now().Format("2006121545")

	fileExemple := []byte(fmt.Sprintf(`package migration

import "gorm.io/gorm"

func migration_%s() MigrationExecute {
    return MigrationExecute{
        Name: "migration_%s",
        Run: func(d *gorm.DB) error {

            return nil
        },
    }
}
    `, timeNow, timeNow))

	fileName := fmt.Sprintf("migration_%s.go", timeNow)

	err := ioutil.WriteFile(fmt.Sprintf("./connection/migration/%s", fileName), fileExemple, 0644)
	if err != nil {
		violog.ErrorLogger.Printf("erro ao criar arquivo de migração; erro:%s", err)
	}
}
