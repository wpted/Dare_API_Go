package repositories

import (
	"dareAPI/model"
	"errors"
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
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

// getRandomID queries the database and return a random valid id between the row length
func getRandomID(d *DareRepo) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(10)
}

// GetByID returns the first match to the given ID from the gorm.DB database
func (d *DareRepo) GetByID(dare *model.Dare, id int) error {
	result := d.Db.First(&dare.ID)
	return result.Error()
}

// GetRandomDare implements
func (d *DareRepo) GetRandomDare(dare *model.Dare) error {
	randomId := getRandomID(d)
	err := d.GetByID(dare, randomId)
	return err
}

func (d *DareRepo) GetAllDares(dare *model.Dare) error {
	return errors.New("this is a temporary error for getting all the dares within the model")
}

func (d *DareRepo) AddDare(dare *model.Dare) error {
	result := d.Db.Create(&dare)
	return result.Error()
}

func (d *DareRepo) UpdateDare(dare *model.Dare) error {
	result := d.Db.Save(&dare)
	return result.Error()
}
