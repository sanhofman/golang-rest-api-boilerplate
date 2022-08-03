package services

import (
	"context"
	"errors"
	"time"

	"github.com/wpcodevo/golang-mongodb/models"
	//"github.com/wpcodevo/golang-mongodb/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChildServiceImpl struct {
	childCollection *mongo.Collection
	ctx            context.Context
}

func NewChildService(childCollection *mongo.Collection, ctx context.Context) ChildService {
	return &ChildServiceImpl{childCollection, ctx}
}

func (p *ChildServiceImpl) CreateChild(Child *models.CreateChildRequest) (*models.DBChild, error) {
	Child.CreateAt = time.Now()
	Child.UpdatedAt = Child.CreateAt
	res, err := p.childCollection.InsertOne(p.ctx, Child)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("Child with that name already exists")
		}
		return nil, err
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"name": 1}, Options: opt}

	if _, err := p.childCollection.Indexes().CreateOne(p.ctx, index); err != nil {
		return nil, errors.New("could not create index for name")
	}

	var newChild *models.DBChild
	query := bson.M{"_id": res.InsertedID}
	if err = p.childCollection.FindOne(p.ctx, query).Decode(&newChild); err != nil {
		return nil, err
	}

	return newChild, nil
}

func (p *ChildServiceImpl) UpdateChild(id string, data *models.UpdateChild) (*models.DBChild, error) {
// 	doc, err := utils.ToDoc(data)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	obId, _ := primitive.ObjectIDFromHex(id)
// 	query := bson.D{{Key: "_id", Value: obId}}
// 	update := bson.D{{Key: "$set", Value: doc}}
// 	res := p.childCollection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))
//
// 	var updatedChild *models.DBChild
//
// 	if err := res.Decode(&updatedChild); err != nil {
// 		return nil, errors.New("no Child with that Id exists")
// 	}

	// return updatedChild, nil

	return nil, errors.New("no Child with that Id exists")
}

func (p *ChildServiceImpl) FindChildById(id string) (*models.DBChild, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var Child *models.DBChild

	if err := p.childCollection.FindOne(p.ctx, query).Decode(&Child); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return Child, nil
}

func (p *ChildServiceImpl) FindChildren(page int, limit int) ([]*models.DBChild, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	skip := (page - 1) * limit

	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))

	query := bson.M{}

	cursor, err := p.childCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(p.ctx)

	var Childs []*models.DBChild

	for cursor.Next(p.ctx) {
		Child := &models.DBChild{}
		err := cursor.Decode(Child)

		if err != nil {
			return nil, err
		}

		Childs = append(Childs, Child)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(Childs) == 0 {
		return []*models.DBChild{}, nil
	}

	return Childs, nil
}

func (p *ChildServiceImpl) DeleteChild(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := p.childCollection.DeleteOne(p.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}