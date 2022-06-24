package main

import (
	"errors"
	"io/fs"
	"path/filepath"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func saveEntriesToDB(entries []File) {
	db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"path", "parent_id"}),
	}).Create(entries)
}

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
			hash = hashString(path)
		} else {
			if (!matchInArray(path, allowedFileTypes)) {
				return nil
			}
			fileType = "file"
			hash, _ = hashFile(path)
		}

		var parent File
		if err := db.Where(&File{Path: parentDir}).First(&parent).Error; err != nil {
			// Look for FilePath in entries slice bacause it has not yet been saved to the db
			if (errors.Is(err, gorm.ErrRecordNotFound)){
				for _, entry := range entries {
					if (entry.Path == parentDir) {
						parent.ID = entry.ID
						break
					}
				}
			}
		}
		
		entry := File{ID: hash, Type: fileType, ParentID: parent.ID, Path: path}
		entries = append(entries, entry)

		if (len(entries) >= BATCH_INSERT_SIZE) {
			saveEntriesToDB(entries)
			entries = nil
		}

		return nil;
	});
	
	if (len(entries) != 0) {
		saveEntriesToDB(entries)
	}
}