package database

import "gorm.io/gorm"

type DBHandler interface {
	Find(interface{}, ...interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
}
