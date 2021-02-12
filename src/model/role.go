package model

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jethro91/fiber-mongo-redis/src/database/mongoDB"
	sessionStore "github.com/jethro91/fiber-mongo-redis/src/model/sessionStoreJWT"
	"github.com/jethro91/fiber-mongo-redis/src/util"
)

type Role struct {
	ID   string `bson:"_id" json:"_id,omitempty"`
	Name string `bson:"name" json:"name,omitempty"`

	// Audit
	CreatedAt     int64  `bson:"createdAt" json:"createdAt,,omitempty"`
	CreatedById   string `bson:"createdById" json:"createdById,omitempty"`
	CreatedByName string `bson:"createdByName" json:"createdByName,omitempty"`
	UpdatedAt     int64  `bson:"updatedAt" json:"updatedAt,omitempty"`
	UpdatedById   string `bson:"updatedById" json:"updatedById,omitempty"`
	UpdatedByName string `bson:"updatedByName" json:"updatedByName,omitempty"`
}

type RoleList []Role

const roleDatabaseId = ""
const roleCollectionId = "role"

func (listData *RoleList) Find(c *fiber.Ctx, filter primitive.M, fields primitive.M) <-chan error {
	chanError := make(chan error)
	go func() {
		requestQuery, err := AssignRequestQuery(c)
		if err != nil {
			chanError <- err
			return
		}
		if requestQuery.Search != "" {
			filter["search"] = requestQuery.Search
		}
		result, err := mongoDB.Find(
			roleDatabaseId,
			roleCollectionId,
			filter,
			fields,
			requestQuery.Limit,
			requestQuery.Page,
			requestQuery.Sort,
			requestQuery.SortDir,
		)
		if err != nil {
			chanError <- err
			return
		}
		err = util.BsonArrToStruct(result, listData)
		if err != nil {
			chanError <- err
			return
		}
		defer close(chanError)
	}()

	return chanError
}

func (listData *RoleList) Count(c *fiber.Ctx, filter primitive.M, resultPointer *int64) <-chan error {
	chanError := make(chan error)
	go func() {
		requestQuery, err := AssignRequestQuery(c)
		if err != nil {
			chanError <- err
			return
		}
		if requestQuery.Search != "" {
			filter["search"] = requestQuery.Search
		}
		totalItem, err := mongoDB.Count(
			roleDatabaseId,
			roleCollectionId,
			filter,
		)
		if err != nil {
			chanError <- err
			return
		}
		*resultPointer = totalItem
		defer close(chanError)
	}()

	return chanError
}

func (data *Role) FindOne(c *fiber.Ctx, filter primitive.M, fields primitive.M) <-chan error {
	chanError := make(chan error)
	go func() {
		defer close(chanError)
		result, err := mongoDB.FindOne(
			roleDatabaseId,
			roleCollectionId,
			filter,
			fields,
		)
		if err != nil {
			chanError <- err
			return
		}
		err = util.BsonToStruct(result, data)
		if err != nil {
			chanError <- err
			return
		}
	}()
	return chanError

}

func (data *Role) Exists(c *fiber.Ctx, filter primitive.M, resultPointer *bool) <-chan error {
	chanError := make(chan error)
	go func() {
		defer close(chanError)
		isExists, err := mongoDB.Exists(
			roleDatabaseId,
			roleCollectionId,
			filter,
		)
		if err != nil {
			chanError <- err
			return
		}
		*resultPointer = isExists
	}()
	return chanError

}

func (data *Role) InsertOne(formData interface{}, disableTimeStamp bool) <-chan error {
	chanError := make(chan error)
	go func() {
		defer close(chanError)
		newDoc := bson.M{}
		err := util.PrimitiveM(formData, &newDoc)
		if err != nil {
			chanError <- err
			return
		}
		myUser := sessionStore.User
		err = mongoDB.InsertOne(
			roleDatabaseId,
			roleCollectionId,
			newDoc,
			myUser,
			disableTimeStamp,
		)
		if err != nil {
			chanError <- err
			return
		}
	}()

	return chanError
}

func (data *Role) UpdateOne(filter primitive.M, formData interface{}, disableTimeStamp bool) <-chan error {
	chanError := make(chan error)
	go func() {
		defer close(chanError)
		myUser := sessionStore.User
		updateDoc := bson.M{}
		err := util.PrimitiveM(formData, &updateDoc)
		if err != nil {
			chanError <- err
			return
		}

		err = mongoDB.UpdateOne(
			roleDatabaseId,
			roleCollectionId,
			filter,
			updateDoc,
			myUser,
			disableTimeStamp,
		)
		if err != nil {
			chanError <- err
			return
		}
	}()
	return chanError
}

func (data *Role) DeleteOne(filter primitive.M) <-chan error {
	chanError := make(chan error)
	go func() {
		defer close(chanError)
		err := mongoDB.DeleteOne(
			roleDatabaseId,
			roleCollectionId,
			filter,
		)
		if err != nil {
			chanError <- err
			return
		}
	}()
	return chanError
}
