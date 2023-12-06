package item

import (
	httpCodeEnum "be/src/common/httpEnum/httpCode"
	httpMessageEnum "be/src/common/httpEnum/httpMessage"
	connection "be/src/database"
	"be/src/handlers/account"
	"be/src/handlers/auth"
	"be/src/models/itemModel"
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
)

type ItemCreate struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	GroupId  int    `json:"groupId"`
}

func Router(api fiber.Router) {
	itemApi := api.Group("items")
	itemApi.Post("/", func(c *fiber.Ctx) error {
		user, error := auth.AuthGuard(c)
		if error != nil {
			return fiber.NewError(httpCodeEnum.UNAUTHORIZED, httpMessageEnum.UNAUTHORIZED)
		}
		return Create(c, &user)
	})
	itemApi.Get("/", func(c *fiber.Ctx) error {
		user, error := auth.AuthGuard(c)
		if error != nil {
			return fiber.NewError(httpCodeEnum.UNAUTHORIZED, httpMessageEnum.UNAUTHORIZED)
		}
		return GetAll(c, &user)
	})
}

func Create(c *fiber.Ctx, user *account.Account) error {
	payload := struct {
		Name       string `json:"name"`
		Username   string `json:"username"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		GroupId    *int   `json:"groupId"`
		AccountId  *int   `json:"accountId"`
		CategoryId *int   `json:"categoryId"`
	}{}
	c.BodyParser(&payload)
	db := connection.Postgres()
	item := itemModel.Item{
		Name:     payload.Name,
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		GroupId:  payload.GroupId,
	}
	if payload.GroupId != nil {
		userId := *&user.ID
		item.AccountId = userId
	}
	result := db.Create(&item).Preload("Group")
	if result.Error != nil {
		return fiber.NewError(httpCodeEnum.WRONG, result.Error.Error())
	}
	return c.Status(httpCodeEnum.OK).JSON(item)
}

func GetAll(c *fiber.Ctx, user *account.Account) error {
	db := connection.Postgres()
	items := []itemModel.Item{}
	db.Model(&[]itemModel.Item{}).Preload("Group").Preload("Group.Category").Find(&items)
	for i := 0; i < len(items); i++ {
		items[i].Password = base64.StdEncoding.EncodeToString([]byte(items[i].Password))
	}
	return c.Status(httpCodeEnum.OK).JSON(items)
}
