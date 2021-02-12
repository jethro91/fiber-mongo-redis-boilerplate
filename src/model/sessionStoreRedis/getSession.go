package sessionStoreRedis

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/database/redisDB"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
)

func GetSession(clientSessionId string) (primitive.M, error) {
	result := bson.M{}
	user := bson.M{}
	result, err := redisDB.GetSession(clientSessionId)
	if err != nil {
		return result, err
	}
	err = util.PrimitiveM(result["user"], &user)

	Data = result
	User = user

	return result, nil
}
