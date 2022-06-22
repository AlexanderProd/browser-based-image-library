package main

/* ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
Filename  string `gorm:"->;type:GENERATED ALWAYS AS (fileNameFromPath(path));default:(-);"` */
type File struct {
	ID       string   `gorm:"primaryKey"`
  Type     string   `gorm:"notNull"`
  Rating   uint8    `gorm:"default:0"`
  FilePath FilePath
  ParentID string
  Tags     []*Tag   `gorm:"many2many:file_tags;"`
  Children []File   `gorm:"foreignkey:ParentID"`
}

type FilePath struct {
  FileID string `gorm:"primaryKey"`
  Path   string `gorm:"notNull;unique"`
}

type Tag struct {
  ID    int     `gorm:"primaryKey,autoIncrement"`
  Name  string  `gorm:"notNull;unique"`
  Color string
  Files []*File `gorm:"many2many:file_tags;"`
}
