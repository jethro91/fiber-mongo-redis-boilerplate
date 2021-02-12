package sessionStoreJWT

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/config"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
)

func UpdateSession(c *fiber.Ctx, clientSessionId string, user primitive.M) error {
	now := util.TimeNowUnixEpoch()

	jwtToken := jwt.New(jwt.SigningMethodHS256)

	user["password"] = "******"
	jwtClaims := jwtToken.Claims.(jwt.MapClaims)
	jwtClaims["user"] = user
	jwtClaims["createdAt"] = now
	jwtClaims["refresh"] = now + config.SESSION_REFRESH_TIMEOUT
	jwtClaims["expires"] = now + config.SESSION_EXPIRE_TIMEOUT

	updatedClientSessionId, err := jwtToken.SignedString([]byte(config.APP_SECRET))
	if err != nil {
		return err
	}

	updateSessionUser(c, updatedClientSessionId, jwtClaims, user)

	return nil
}

func updateSessionUser(
	c *fiber.Ctx,
	updatedClientSessionId string,
	jwtClaims jwt.MapClaims,
	user primitive.M,
) {
	updatedSession := bson.M{
		"clientSessionId": updatedClientSessionId,
		"user":            jwtClaims["user"],
		"createdAt":       jwtClaims["createdAt"],
		"refresh":         jwtClaims["refresh"],
		"expires":         jwtClaims["expires"],
	}

	Data = updatedSession
	User = user
	c.Set("authorization", updatedClientSessionId)
	cookie := new(fiber.Cookie)
	cookie.Name = "authorization"
	cookie.Value = updatedClientSessionId
	cookie.Path = "/"
	c.Cookie(cookie)
}
