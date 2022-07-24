package repositories

import (
	"context"
	"dareAPI/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Repo interface {
	GetDareByID(id string) (*model.Dare, error)
	GetAllDares() (model.DareContainer, error)
	CreateDare(d *model.Dare) error
	UpdateDare(id, newDareQuestion string) error
	DeleteDare(id string) error
}

type DareRepo struct {
	Ctx        context.Context
	Client     *mongo.Client
	Collection *mongo.Collection
}

// NewDareRepo creates a Dare repository connected to the MongoDB
func NewDareRepo(dbURI string) (Repo, error) {
	var NewDareRepo DareRepo
	// Create a non-nil empty context, with no deadline and timeout
	NewDareRepo.Ctx = context.Background()

	// Create a connection by passing the given URI
	NewDareRepo.Client, _ = mongo.Connect(NewDareRepo.Ctx, options.Client().ApplyURI(dbURI))
	if err := NewDareRepo.Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return &DareRepo{}, err
	}

	// collection pool handle
	NewDareRepo.Collection = NewDareRepo.Client.Database("Dare_Mongo").Collection("Dares")

	return &NewDareRepo, nil
}

// GetDareByID return a pointer to Dare and an error. If ID not found, return an empty dare.
func (r *DareRepo) GetDareByID(id string) (*model.Dare, error) {
	var dare model.Dare
	// Convert the given id string to a form that Mongo can read
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return &model.Dare{}, err
	}

	err = r.Collection.FindOne(r.Ctx, model.Dare{ID: objectID}).Decode(&dare)
	if err != nil {
		return &model.Dare{}, err
	}

	return &dare, nil
}

// GetAllDares returns a container of the fetched dare.
func (r *DareRepo) GetAllDares() (model.DareContainer, error) {
	var dares = make(model.DareContainer, 0)
	cursor, err := r.Collection.Find(r.Ctx, bson.M{})

	if err != nil {
		return dares, err
	}

	// shut the cursor down when there is nothing left in the pool of connection
	// after closing cursor, Next and TryNext won't work anymore
	defer cursor.Close(r.Ctx)

	for cursor.Next(r.Ctx) {
		var dare model.Dare
		// Decode the bson type into the given struct
		cursor.Decode(&dare)
		dares = append(dares, dare)
	}

	return dares, nil
}

// CreateDare inputs the give Dare into the database
func (r *DareRepo) CreateDare(d *model.Dare) error {
	_, err := r.Collection.InsertOne(r.Ctx, d)
	if err != nil {
		return err
	}
	return nil
}

// UpdateDare updates the dare within the database by the given id
func (r *DareRepo) UpdateDare(id, newDareQuestion string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// update by the filter bson.M and update through bson.D
	_, err = r.Collection.UpdateOne(
		r.Ctx,
		bson.M{"_id": objectID},
		bson.D{{"$set", bson.D{{"dare", newDareQuestion}}}})

	if err != nil {
		return err
	}

	return nil
}

func (r *DareRepo) DeleteDare(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(r.Ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	return nil
}
