package repositories

import (
	"dareAPI/model"
	"github.com/jinzhu/gorm"
)

type Repo interface {
	AddDare(dare *model.Dare) error
	UpdateDare(dare *model.Dare) error
}

// DareRepo is a struct with a pointer to the database object gorm.DB that prevents Go from making a copy of this object gorm.DB.
// Here db is a pointer to gorm.DB, a representation of a pool of database connections
type DareRepo struct {
	Db *gorm.DB
}

// NewDareRepo instantiates a new gorm.db database
func NewDareRepo(db *gorm.DB) *DareRepo {
	return &DareRepo{
		Db: db}
}

//
//func (d dareRepo) GetByID() {
//
//}

func (d *DareRepo) AddDare(dare *model.Dare) error {
	result := d.Db.Create(&dare)
	return result.Error()
}

func (d *DareRepo) UpdateDare(dare *model.Dare) error {
	result := d.Db.Save(&dare)
	return result.Error()
}
