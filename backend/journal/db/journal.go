package db

import (
	"journal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// create a constant for the database file name
const DB_FILENAME = "journal.db"

// function to reset and seed the database for development
func ResetDB(db *gorm.DB) error {
	// Drop all tables
	err := db.Migrator().DropTable(&models.Journal{})
	if err != nil {
		return err
	}

	// Reset the auto-increment value
	db.Exec("DELETE FROM sqlite_sequence WHERE name='journals'")

	// Migrate the schema
	err = db.AutoMigrate(&models.Journal{})
	if err != nil {
		return err
	}

	// Seed the database with some initial data
	journals := []models.Journal{
		{Title: "First Entry", Content: "This is my first journal entry."},
		{Title: "Second Entry", Content: "This is my second journal entry."},
	}

	for _, j := range journals {
		if err := db.Create(&j).Error; err != nil {
			return err
		}
	}

	return nil
}

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(DB_FILENAME), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Journal{})
	if err != nil {
		return nil, err
	}

	// Reset the DB for development
	ResetDB(db)
	return db, nil
}

func CreateJournal(db *gorm.DB, j *models.Journal) error {
	return db.Create(j).Error
}

func GetJournal(db *gorm.DB, id uint) (*models.Journal, error) {
	var j models.Journal
	err := db.First(&j, id).Error
	return &j, err
}

func UpdateJournal(db *gorm.DB, j *models.Journal) error {
	return db.Save(j).Error
}

func DeleteJournal(db *gorm.DB, id uint) error {
	return db.Delete(&models.Journal{}, id).Error
}
