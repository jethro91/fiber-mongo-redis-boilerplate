package model

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/database/mongoDB"
	sessionStore "github.com/jethro91/fiber-mongo-redis-boilerplate/src/model/sessionStoreJWT"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
)

type PasswordReset struct {
	ID        string `bson:"_id" json:"_id,omitempty"`
	UserID    string `bson:"userId" json:"userId,omitempty"`
	Token     string `bson:"token" json:"token,omitempty"`
	ExpiredAt int64  `bson:"expiredAt" json:"expiredAt,,omitempty"`
}

type PasswordResetList []PasswordReset

const passwordResetDatabaseId = ""
const passwordResetCollectionId = "passwordReset"

func (listData *PasswordResetList) Find(c *fiber.Ctx, filter primitive.M, fields primitive.M) <-chan error {
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
			passwordResetDatabaseId,
			passwordResetCollectionId,
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

func (listData *PasswordResetList) Count(c *fiber.Ctx, filter primitive.M, resultPointer *int64) <-chan error {
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
			passwordResetDatabaseId,
			passwordResetCollectionId,
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

func (data *PasswordReset) FindOne(c *fiber.Ctx, filter primitive.M, fields primitive.M) <-chan error {
	chanError := make(chan error)
	go func() {
		defer close(chanError)
		result, err := mongoDB.FindOne(
			passwordResetDatabaseId,
			passwordResetCollectionId,
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

func (data *PasswordReset) Exists(c *fiber.Ctx, filter primitive.M, resultPointer *bool) <-chan error {
	chanError := make(chan error)
	go func() {
		defer close(chanError)
		isExists, err := mongoDB.Exists(
			passwordResetDatabaseId,
			passwordResetCollectionId,
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

func (data *PasswordReset) InsertOne(formData interface{}, disableTimeStamp bool) <-chan error {
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
			passwordResetDatabaseId,
			passwordResetCollectionId,
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

func (data *PasswordReset) UpdateOne(filter primitive.M, formData interface{}, disableTimeStamp bool) <-chan error {
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
			passwordResetDatabaseId,
			passwordResetCollectionId,
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

func (data *PasswordReset) DeleteOne(filter primitive.M) <-chan error {
	chanError := make(chan error)
	go func() {
		defer close(chanError)
		err := mongoDB.DeleteOne(
			passwordResetDatabaseId,
			passwordResetCollectionId,
			filter,
		)
		if err != nil {
			chanError <- err
			return
		}
	}()
	return chanError
}

func (listData *PasswordResetList) DeleteMany(filter primitive.M) <-chan error {
	chanError := make(chan error)
	go func() {
		defer close(chanError)
		err := mongoDB.DeleteMany(
			passwordResetDatabaseId,
			passwordResetCollectionId,
			filter,
		)
		if err != nil {
			chanError <- err
			return
		}
	}()
	return chanError
}
