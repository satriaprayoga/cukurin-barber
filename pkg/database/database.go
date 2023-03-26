package database

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/satriaprayoga/cukurin-barber/models"
	"github.com/satriaprayoga/cukurin-barber/pkg/settings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Conn *gorm.DB

func Setup() {
	now := time.Now()
	var err error

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		settings.AppConfigSetting.Database.Host,
		settings.AppConfigSetting.Database.User,
		settings.AppConfigSetting.Database.Password,
		settings.AppConfigSetting.Database.Name,
		settings.AppConfigSetting.Database.Port)
	fmt.Printf("%s", connectionString)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)

	Conn, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   settings.AppConfigSetting.Database.TablePrefix,
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		log.Printf("connection.setup err : %v", err)
		panic(err)
	}
	sqlDB, err := Conn.DB()
	if err != nil {
		log.Printf("connection.setup DB err : %v", err)
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	migrate()

	timeSpent := time.Since(now)
	log.Printf("Config database is ready in %v", timeSpent)
}

func migrate() {
	log.Println("START AUTO MIGRATE")
	Conn.AutoMigrate(
		//models.KUser{},
		//models.KSession{},
		models.FileUpload{})
	log.Println("FINISHING AUTO MIGRATE ")
}

// GetWhereLikeStruct :
func GetWhereLikeStruct(v reflect.Value, t reflect.Type, searchParam string, fieldLst string) string {
	result := ""
	vt := v.Type()
	if fieldLst == "" {
		for i := 0; i < vt.NumField(); i++ {
			varName := fmt.Sprintf("%v", v.Type().Field(i).Name) //field Name
			varType := v.Type().Field(i).Type                    //fmt.Sprintf("%v", v.Type().Field(i).Type) // field type data
			// varValue := fmt.Sprintf("%v", v.Field(i).Interface()) //
			field, _ := t.Elem().FieldByName(fmt.Sprintf("%v", varName)) // getTag json
			varTagJSON := fmt.Sprintf("%v", field.Tag)                   //get value json

			i1 := strings.Index(varTagJSON, `"`)
			str1 := varTagJSON[i1+1:]

			i2 := strings.Index(str1, `"`)
			str2 := str1[:i2]
			varFieldtable := fmt.Sprintf(str2)
			fmt.Printf("%v\n", varType)
			sType := fmt.Sprintf("%v\n", varType)
			fmt.Printf(sType)
			if strings.Contains(sType, "models") {
				continue
			}
			if strings.Index(varFieldtable, ",") > 0 {
				varFieldtable = strings.Split(varFieldtable, ",")[0]
			}
			// switch varType {
			// 	case int16
			// }
			result += fmt.Sprintf("OR lower(%s::varchar) LIKE '%%%s%%' ", varFieldtable, strings.ToLower(searchParam))
		}
	} else {
		arrField := strings.Split(fieldLst, ",")
		for i := 0; i < len(arrField); i++ {
			varName := arrField[i]
			result += fmt.Sprintf("OR lower(%s::varchar) LIKE '%%%s%%' ", varName, strings.ToLower(searchParam))
		}
	}

	i1 := strings.Index(result, `OR`)
	str1 := result[i1+2:]

	// fmt.Printf("\n%s\n", str1)
	result = "( " + str1 + " )"
	fmt.Printf("%s", result)
	return result
}
