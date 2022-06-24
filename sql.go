package main

type File struct {
	ID       string `gorm:"primaryKey"`
  Type     string `gorm:"notNull"`
  Rating   uint8  `gorm:"default:0"`
  Path     string `gorm:"notNull"`
  ParentID string
  Tags     []*Tag `gorm:"many2many:file_tags;"`
  Children []File `gorm:"foreignkey:ParentID"`
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
  Files       []*File `gorm:"many2many:file_tags;"`
}
