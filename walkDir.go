package main

import (
	"errors"
	"io/fs"
	"path/filepath"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func walkDir() {
	var entries []File

	filepath.WalkDir(PATH, func(path string, file fs.DirEntry,  err error) error {
		if err != nil {
			return err
		}
		
		var fileType string
		var hash string
		var parentDir = filepath.Dir(path)

		if file.IsDir() {
			fileType = "folder"
			hash = uuid.NewString()
		} else {
			if (!matchInArray(path, allowedFileTypes)) {
				return nil
			}
			fileType = "file"
			hash = hashFile(path)
		}

		var parent FilePath
		if err := db.Where(&FilePath{Path: parentDir}).First(&parent).Error; err != nil {
			// Look for FilePath in entries slice bacause it has not yet been saved to the db
			if (errors.Is(err, gorm.ErrRecordNotFound)){
				for _, entry := range entries {
					if (entry.FilePath.Path == parentDir) {
						parent.FileID = entry.ID
						break
					}
				}
			}
		}

		entry := File{ID: hash, Type: fileType, FilePath: FilePath{Path: path}, ParentID: parent.FileID}
		entries = append(entries, entry)

		if (len(entries) > BATCH_INSERT_SIZE) {
			db.Clauses(clause.OnConflict{DoNothing: true}).Create(entries)
			entries = nil
		}

		return nil;
	});
	
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(entries)
}