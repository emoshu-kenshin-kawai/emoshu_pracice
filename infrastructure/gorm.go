package infrastructure

import (
	"emoshu_practice/config"
	"emoshu_practice/domain"
	"emoshu_practice/interfaces/database"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBHandler struct {
	DB *gorm.DB
}

func InitDB() database.DBHandler {
	db := newGorm()
	return db
}

func (handler *DBHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.DB.Find(out, where...)
}

func (handler *DBHandler) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return handler.DB.First(dest, conds...)
}

func (handler *DBHandler) Joins(query string, args ...interface{}) *gorm.DB {
	return handler.DB.Joins(query, args...)
}

func (handler *DBHandler) Create(value interface{}) *gorm.DB {
	return handler.DB.Create(value)
}

func (handler *DBHandler) Save(value interface{}) *gorm.DB {
	return handler.DB.Save(value)
}

func (handler *DBHandler) Model(value interface{}) *gorm.DB {
	return handler.DB.Model(value)
}

func (handler *DBHandler) Updates(values interface{}) *gorm.DB {
	return handler.DB.Updates(values)
}

func (handler *DBHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.DB.Where(query, args...)
}

func (handler *DBHandler) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return handler.DB.Delete(value, conds...)
}

func newGorm() database.DBHandler {
	driver := config.Mysql{
		Host:     config.GetEnvWithDefautl("DB_HOST", "db"),
		Port:     config.GetEnvWithDefautl("DB_PORT", "3306"),
		UserName: config.GetEnvWithDefautl("DB_USER", "emoshu_user"),
		Password: config.GetEnvWithDefautl("DB_PASSWORD", "password"),
		DBName:   config.GetEnvWithDefautl("DB_NAME", "emoshu_practice_dev"),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", driver.UserName, driver.Password, driver.Host, driver.Port, driver.DBName)
	db := connectDB(dsn)
	err := autoMigrate(db)
	if err != nil {
		panic(err)
	}
	dbHandler := new(DBHandler)
	dbHandler.DB = db
	return dbHandler
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
			return nil, fmt.Errorf("retry count over")
		}
		time.Sleep(time.Second)
		count--
		return openDBWithTimeLimit(dsn, count)
	}
	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		new(domain.Department),
		new(domain.EmploymentStatus),
		new(domain.Status),
		new(domain.Member),
		new(domain.MemberDepartment),
		new(domain.MemberRole),
		new(domain.Role),
	)
	return err
}
