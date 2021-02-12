package sessionStoreJWT

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/jethro91/fiber-mongo-redis/src/config"
	"github.com/jethro91/fiber-mongo-redis/src/util"
)

func GetSession(clientSessionId string) (primitive.M, error) {
	result := bson.M{}
	user := bson.M{}

	jwtClaims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(clientSessionId, jwtClaims, getSessionJwtCb)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	result = bson.M{
		"clientSessionId": clientSessionId,
		"user":            jwtClaims["user"],
		"createdAt":       jwtClaims["createdAt"],
		"refresh":         jwtClaims["refresh"],
		"expires":         jwtClaims["expires"],
	}

	err = util.PrimitiveM(result["user"], &user)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	Data = result
	user["password"] = "******"
	User = user

	return result, nil
}

func getSessionJwtCb(token *jwt.Token) (interface{}, error) {
	return []byte(config.APP_SECRET), nil
}
