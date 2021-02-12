package sessionStoreJWT

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/jethro91/fiber-mongo-redis/src/config"
	"github.com/jethro91/fiber-mongo-redis/src/util"
)

func SessionLogin(c *fiber.Ctx, user primitive.M) error {
	now := util.TimeNowUnixEpoch()

	jwtToken := jwt.New(jwt.SigningMethodHS256)

	jwtClaims := jwtToken.Claims.(jwt.MapClaims)
	user["password"] = "******"
	jwtClaims["user"] = user
	jwtClaims["createdAt"] = now
	jwtClaims["refresh"] = now + config.SESSION_REFRESH_TIMEOUT
	jwtClaims["expires"] = now + config.SESSION_EXPIRE_TIMEOUT

	clientSessionId, err := jwtToken.SignedString([]byte(config.APP_SECRET))
	if err != nil {
		return err
	}

	setSessionUser(c, clientSessionId, jwtClaims, user)

	return nil
}

func setSessionUser(
	c *fiber.Ctx,
	clientSessionId string,
	jwtClaims jwt.MapClaims,
	user primitive.M,
) {
	newSession := bson.M{
		"clientSessionId": clientSessionId,
		"user":            jwtClaims["user"],
		"createdAt":       jwtClaims["createdAt"],
		"refresh":         jwtClaims["refresh"],
		"expires":         jwtClaims["expires"],
	}

	Data = newSession
	User = user
	c.Set("authorization", clientSessionId)
	cookie := new(fiber.Cookie)
	cookie.Name = "authorization"
	cookie.Value = clientSessionId
	cookie.Path = "/"
	c.Cookie(cookie)
}
