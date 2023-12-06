package group

import (
	httpCodeEnum "be/src/common/httpEnum/httpCode"
	httpMessageEnum "be/src/common/httpEnum/httpMessage"
	connection "be/src/database"
	"be/src/handlers/account"
	"be/src/handlers/auth"
	repository "be/src/models"
	groupModel "be/src/models/groupItemModel"

	"github.com/gofiber/fiber/v2"
)

func Router(api fiber.Router) {
	groupApi := api.Group("/groups")
	groupApi.Post("/", func(c *fiber.Ctx) error {
		user, error := auth.AuthGuard(c)
		if error != nil {
			return fiber.NewError(httpCodeEnum.UNAUTHORIZED, httpMessageEnum.UNAUTHORIZED)
		}
		return Create(c, &user)
	})
	groupApi.Get("/", func(c *fiber.Ctx) error {
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
		CategoryId int    `json:"categoryId"`
	}{}
	c.BodyParser(&payload)
	db := connection.Postgres()
	group := groupModel.Group{
		Name:       payload.Name,
		AccountId:  *user.ID,
		CategoryId: payload.CategoryId,
	}
	db.Create(&group)
	return c.Status(httpCodeEnum.OK).JSON(group)
}

func GetAll(c *fiber.Ctx, user *account.Account) error {
	type Group struct {
		ID         int                 `json:"id"`
		Name       string              `json:"name"`
		AccountId  int                 `json:"accountId"`
		CategoryId int                 `json:"categoryId"`
		Category   repository.Category `json:"category"`
	}
	db := connection.Postgres()
	group := []groupModel.Group{}
	db.Model(&groupModel.Group{}).Find(&group)
	return c.Status(httpCodeEnum.OK).JSON(group)
}
