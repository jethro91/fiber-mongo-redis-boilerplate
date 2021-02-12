package sessionStoreJWT

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	Data primitive.M
	User primitive.M
)

func ClearSession() {
	Data = bson.M{}
	User = bson.M{}
}
