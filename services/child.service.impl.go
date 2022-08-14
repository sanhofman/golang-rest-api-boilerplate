package services

import (
    "errors"
	"context"
	"time"

    "gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"github.com/wpcodevo/golang-mongodb/database/common/dbModels"

	"github.com/wpcodevo/golang-mongodb/models"
)

type ChildServiceImpl struct {
	db             *gorm.DB
	ctx            context.Context
}

func NewChildServiceImpl(db *gorm.DB, ctx context.Context) ChildService {
	return &ChildServiceImpl{db, ctx}
}

func (p *ChildServiceImpl) CreateChild(Child *models.CreateChildRequest, ctx *gin.Context) (dbModels.Child, error) {
	Child.CreateAt = time.Now()
	Child.UpdatedAt = Child.CreateAt

    var child dbModels.Child

    // Check existence.
    p.db.Where("name = ? AND parent = ?", Child.Name, Child.Parent).First(&child)
    if child.Name != "" {
        return child, errors.New("child already exists")
    }

    // @TODO:: bind user id instead of name
	currentUser := ctx.MustGet("currentUser").(*models.DBResponse)

    // Create child.
    child.Name = Child.Name
    child.Parent = Child.Parent
    child.CreatedBy = currentUser.Name

    if result := p.db.Create(&child); result.Error != nil {
        return child, errors.New("child creation failed")
    }

	return child, nil
}

func (p *ChildServiceImpl) UpdateChild(id string, data *models.UpdateChild) (dbModels.Child, error) {
    var child dbModels.Child

    if result := p.db.First(&child, id); result.Error != nil {
        return child, errors.New("Child not found")
    }

    // Map values.
    child.Name = data.Name // required

    if (data.Parent != "") {
        child.Parent = data.Parent
    }

    p.db.Save(&child)

	return child, nil
}

func (p *ChildServiceImpl) FindChildById(id string) (dbModels.Child, error) {
    var child dbModels.Child

    if result := p.db.First(&child, id); result.Error != nil {
        // c.AbortWithError(http.StatusNotFound, result.Error)
        return child, errors.New("Child not found")
    }

    return child, nil
}

func (p *ChildServiceImpl) FindChildren(page int, limit int) ([]dbModels.Child, error) {
    // @TODO:: implement paging?
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

    var children []dbModels.Child

    if result := p.db.Find(&children); result.Error != nil {
        return children, errors.New("No children found")
    }

    // @TODO:: use separate input/output objects?
    return children, nil
}

func (p *ChildServiceImpl) DeleteChild(id string) error {
// 	obId, _ := primitive.ObjectIDFromHex(id)
// 	query := bson.M{"_id": obId}
//
// 	res, err := p.childCollection.DeleteOne(p.ctx, query)
// 	if err != nil {
// 		return err
// 	}
//
// 	if res.DeletedCount == 0 {
// 		return errors.New("no document with that Id exists")
// 	}

	return nil
}