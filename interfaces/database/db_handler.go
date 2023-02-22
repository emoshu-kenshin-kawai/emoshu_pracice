package database

import (
	"gorm.io/gorm"
)

type DBHandler interface {
	Find(interface{}, ...interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
	Joins(string, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Model(interface{}) *gorm.DB
	Updates(interface{}) *gorm.DB
	Where(interface{}, ...interface{}) *gorm.DB
}
