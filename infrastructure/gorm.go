package infrastructure

import (
	"emoshu_practice/config"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormHandler struct {
	Db *gorm.DB
}

func InitDB() {
	db := newGorm()
	fmt.Println(db)

}

func newGorm() *gorm.DB {
	driver := config.Mysql{
		Host:     config.GetEnvWithDefautl("DB_HOST", "db"),
		Port:     config.GetEnvWithDefautl("DB_PORT", "3306"),
		UserName: config.GetEnvWithDefautl("DB_USER", "root"),
		Password: config.GetEnvWithDefautl("DB_PASSWORD", "password"),
		DBName:   config.GetEnvWithDefautl("DB_NAME", "emoshu_practice_dev"),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", driver.UserName, driver.Password, driver.Host, driver.Port, driver.DBName)
	db := connectDB(dsn)
	return db
}

func connectDB(dsn string) *gorm.DB {
	db, err := openDBWithTimeLimit(dsn, 5)
	if err != nil {
		panic(err)
	}
	return db
}

func openDBWithTimeLimit(dsn string, count int) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		if count == 0 {
			return nil, fmt.Errorf("Retry count over")
		}
		time.Sleep(time.Second)
		count--
		return openDBWithTimeLimit(dsn, count)
	}
	return db, nil
}
