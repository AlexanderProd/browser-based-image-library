package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(dbPath string) (*gorm.DB) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
  if err != nil {
    panic("failed to connect database")
  }

	db.AutoMigrate(&File{}, &Tag{}, &Category{})
	
	return db
}

type File struct {
	ID         string      `gorm:"primaryKey"`
  Type       string      `gorm:"notNull"`
  Rating     uint8       `gorm:"default:0"`
  Path       string      `gorm:"notNull"`
  ParentID   string
  Tags       []*Tag      `gorm:"many2many:file_tags;"`
  Categories []*Category `gorm:"many2many:file_categories;"`
  Children   []File      `gorm:"foreignkey:ParentID"`
}

type Tag struct {
  ID    int     `gorm:"primaryKey,autoIncrement"`
  Name  string  `gorm:"notNull;unique"`
  Color string
  Files []*File `gorm:"many2many:file_tags;"`
}

type Category struct {
  ID          string  `gorm:"primaryKey,autoIncrement"`
  Name        string  `gorm:"notNull;unique"`
  Description string
  Files       []*File `gorm:"many2many:file_categories;"`
}
