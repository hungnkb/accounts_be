package account

import (
	// "be/src/handlers/auth"
	httpCodeEnum "be/src/common/httpEnum/httpCode"
	httpMessageEnum "be/src/common/httpEnum/httpMessage"
	connection "be/src/database"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Account struct {
	ID       *int   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func Router(api fiber.Router) {
	accountApi := api.Group("/accounts")
	accountApi.Get("/me", func(c *fiber.Ctx) error {
		return Me(c)
	})
}

func Me(c *fiber.Ctx) error {
	bearerToken := c.Get("Authorization")
	token := strings.Split(bearerToken, "Bearer ")
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token[1], claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(("erorr"))
		}
		return []byte("hehe"), nil
	})
	id := claims["id"]
	idFloat64, ok := id.(float64)
	if ok == false {
		return fiber.NewError(httpCodeEnum.NOT_FOUND, httpMessageEnum.USER_NOT_FOUND)
	}
	if user, error := GetOne(idFloat64); error != nil {
		return fiber.NewError(httpCodeEnum.NOT_FOUND, httpMessageEnum.USER_NOT_FOUND)
	} else {
		return c.Status(httpCodeEnum.OK).JSON(user)
	}
}

func GetOne(id float64) (user Account, error error) {
	db := connection.Postgres()
	var idParse *int
	idParse = new(int)
	*idParse = int(id)
	if err := db.Where(&Account{ID: idParse}).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
