package middleware

import (
	"errors"
	"reflect"

	"github.com/gofiber/fiber/v2"
	sessionStore "github.com/jethro91/fiber-mongo-redis-boilerplate/src/model/sessionStoreRedis"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckSessionRedis(c *fiber.Ctx) error {
	sessionStore.Data = bson.M{}
	sessionStore.User = bson.M{}
	// clientSessionId := c.Get("authorization")
	clientSessionId := c.Cookies("authorization")
	var err error

	sessionClaims := bson.M{}
	if clientSessionId != "" {
		sessionClaims, err = sessionStore.GetSession(clientSessionId)
		if err != nil {
			sessionStore.SessionLogout(c)
			return util.HttpUnauthorized(c, err.Error())
		}
		if reflect.DeepEqual(sessionClaims, bson.M{}) {
			return util.HttpUnauthorized(c, "Session Expired: Invalid Session")
		}
		err = checkSessionJWTValidity(c, clientSessionId, sessionClaims)
		if err != nil {
			return util.HttpUnauthorized(c, err.Error())
		}
	}

	return c.Next()
}

func checkSesseionRedisValidity(c *fiber.Ctx, clientSessionId string, sessionClaims primitive.M) error {
	// Parse Timeout
	now := util.TimeNowUnixEpoch()
	var refreshTimeout int64
	var expireTimeout int64
	err := util.PrimitiveFloatTo64(sessionClaims["refresh"], &refreshTimeout)
	if err != nil {
		return err
	}
	err = util.PrimitiveFloatTo64(sessionClaims["expires"], &expireTimeout)
	if err != nil {
		return err
	}
	// Chek if Expires
	if now > expireTimeout {
		err = sessionStore.SessionLogout(c)
		return errors.New("Session Expired")
	}
	// Check if need Update Session
	if now < expireTimeout && now > refreshTimeout {
		user := bson.M{}
		err = util.PrimitiveM(sessionClaims["user"], &user)
		if err != nil {
			return err
		}
		// Update the Session
		err = sessionStore.UpdateSession(c, clientSessionId, user)
		if err != nil {
			return err
		}
	}
	return nil
}
