package model

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/database/mongoDB"
	sessionStore "github.com/jethro91/fiber-mongo-redis-boilerplate/src/model/sessionStoreJWT"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
)

type User struct {
	ID         string `bson:"_id" json:"_id,omitempty"`
	IsArchived bool   `bson:"isArchived" json:"isArchived,omitempty"`
	Email      string `bson:"email" json:"email,omitempty"`
	Name       string `bson:"name" json:"name,omitempty"`
	Password   string `bson:"password" json:"password,omitempty"`
	VerifiedAt int64  `bson:"verifiedAt" json:"verifiedAt,omitempty"`
	// Audit
	CreatedAt     int64  `bson:"createdAt" json:"createdAt,omitempty"`
	CreatedById   string `bson:"createdById" json:"createdById,omitempty"`
	CreatedByName string `bson:"createdByName" json:"createdByName,omitempty"`
	UpdatedAt     int64  `bson:"updatedAt" json:"updatedAt,omitempty"`
	UpdatedById   string `bson:"updatedById" json:"updatedById,omitempty"`
	UpdatedByName string `bson:"updatedByName" json:"updatedByName,omitempty"`
	DeletedAt     int64  `bson:"deletedAt" json:"deletedAt,omitempty"`
	DeletedById   string `bson:"deletedById" json:"deletedById,omitempty"`
	DeletedByName string `bson:"deletedByName" json:"deletedByName,omitempty"`
	// Roles
	Roles       primitive.M `bson:"roles" json:"roles,omitempty"`
	RolesAt     int64       `bson:"rolesAt" json:"rolesAt,omitempty"`
	RolesById   string      `bson:"rolesById" json:"rolesById,omitempty"`
	RolesByName string      `bson:"rolesByName" json:"rolesByName,omitempty"`
}
type UserList []User

const userDatabaseId = ""
const userCollectionId = "user"

func (listData *UserList) Find(c *fiber.Ctx, filter primitive.M, fields primitive.M) <-chan error {
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
			userDatabaseId,
			userCollectionId,
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

func (listData *UserList) Count(c *fiber.Ctx, filter primitive.M, resultPointer *int64) <-chan error {
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
			userDatabaseId,
			userCollectionId,
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

func (data *User) FindOne(c *fiber.Ctx, filter primitive.M, fields primitive.M) <-chan error {
	chanError := make(chan error)
	go func() {
		defer close(chanError)
		result, err := mongoDB.FindOne(
			userDatabaseId,
			userCollectionId,
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

func (data *User) Exists(c *fiber.Ctx, filter primitive.M, resultPointer *bool) <-chan error {
	chanError := make(chan error)
	go func() {
		defer close(chanError)
		isExists, err := mongoDB.Exists(
			userDatabaseId,
			userCollectionId,
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

func (data *User) InsertOne(formData interface{}, disableTimeStamp bool) <-chan error {
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
			userDatabaseId,
			userCollectionId,
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

func (data *User) UpdateOne(filter primitive.M, formData interface{}, disableTimeStamp bool) <-chan error {
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
			userDatabaseId,
			userCollectionId,
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

func (data *User) DeleteOne(filter primitive.M) <-chan error {
	chanError := make(chan error)
	go func() {
		defer close(chanError)
		err := mongoDB.DeleteOne(
			userDatabaseId,
			userCollectionId,
			filter,
		)
		if err != nil {
			chanError <- err
			return
		}
	}()
	return chanError
}
