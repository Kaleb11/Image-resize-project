package db

import (
	"image/internal/constant/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbPlatform interface {
	Open() (*gorm.DB, error)
	Migrate() error
}

type dbPlatform struct {
	dbURL string
}

func Initialize(dbURL string) DbPlatform {
	return &dbPlatform{
		dbURL: dbURL}
}

func (cp *dbPlatform) Open() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(cp.dbURL), &gorm.Config{})
}

func (cp *dbPlatform) Migrate() error {
	db, err := cp.Open()
	if err != nil {
		return err
	}
	dbc, err := db.DB()
	if err != nil {
		return err
	}
	defer dbc.Close()
	if !db.Migrator().HasTable(&model.Image{}) {
		err := db.Migrator().CreateTable(&model.Image{})
		if err != nil {
			return err
		}
	}
	if !db.Migrator().HasTable(&model.Formater{}) {
		err := db.Migrator().CreateTable(&model.Formater{})
		if err != nil {
			return err
		}
	}
	return nil
}

// At system initialization
// Check if the super admin exists by checking if the permission and roles table don't exist,
//
