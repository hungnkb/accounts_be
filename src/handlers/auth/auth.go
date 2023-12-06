package auth

import (
	"be/src/common/credentialEnum"
	httpCodeEnum "be/src/common/httpEnum/httpCode"
	httpMessageEnum "be/src/common/httpEnum/httpMessage"
	connection "be/src/database"
	"be/src/handlers/account"
	"be/src/models/accountModel"
	credentialModel "be/src/models/model"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type JwtPayload struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Claims struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
	jwt.RegisteredClaims
}

type JwtToken struct {
	AccessToken string `json:"accessToken"`
}

func Router(api fiber.Router) {
	authApi := api.Group("/auth")
	authApi.Post("/register", func(c *fiber.Ctx) error {
		return Register(c)
	})
	authApi.Post("/login", func(c *fiber.Ctx) error {
		return Login(c)
	})
}

func Register(c *fiber.Ctx) error {
	payload := struct {
		Username string `validate:"required,min=6,max=20" json:"username"`
		Password string `validate:"required,min=6,max=20" json:"password"`
		Name     string `json:"name"`
		Email    string `validate:"required" json:"email"`
	}{}
	c.BodyParser(&payload)
	hashPassword, _ := HashPassword(payload.Password)

	db := connection.Postgres()
	credentialCheck := credentialModel.Credential{}
	errCredential := db.First(&credentialCheck, &credentialModel.Credential{
		Username: payload.Username,
		Key:      credentialEnum.PASSWORD,
	}).Or(&credentialModel.Credential{
		Email: payload.Email,
		Key:   credentialEnum.PASSWORD,
	})
	errUsername := db.First(&accountModel.Account{}, &accountModel.Account{
		Username: payload.Username,
	})
	errEmail := db.First(&accountModel.Account{}, &accountModel.Account{
		Email: payload.Email,
	})
	if errCredential != nil {
		return c.Status(httpCodeEnum.EXIST).JSON(httpMessageEnum.ACCOUNT_EXIST)
	}
	if errUsername != nil {
		return c.Status(httpCodeEnum.EXIST).JSON(httpMessageEnum.USERNAME_EXIST)
	}
	if errEmail != nil {
		return c.Status(httpCodeEnum.EXIST).JSON(httpMessageEnum.EMAIL_EXIST)
	}
	account := accountModel.Account{
		Name:     payload.Name,
		Email:    payload.Email,
		Username: payload.Username,
	}
	db.Create(&account)
	db.Create(&credentialModel.Credential{
		AccountId: account.ID,
		Key:       credentialEnum.PASSWORD,
		Username:  payload.Username,
		Password:  hashPassword,
		Email:     payload.Email,
	})
	return c.Status(httpCodeEnum.OK).JSON(account)
}

func Login(c *fiber.Ctx) error {
	payload := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	c.BodyParser(&payload)
	db := connection.Postgres()
	account := accountModel.Account{}
	checkAccount := db.First(&account, &accountModel.Account{
		Username: payload.Username,
	})
	if checkAccount.Error != nil {
		return fiber.NewError(httpCodeEnum.UNAUTHORIZED, httpMessageEnum.USERNAME_NOT_FOUND)
	}
	credential := credentialModel.Credential{
		ID: account.ID,
	}
	db.Where(&credential).First(&credential)
	checkPassword := bcrypt.CompareHashAndPassword([]byte(credential.Password), []byte(payload.Password))
	if checkPassword != nil {
		return c.Status(httpCodeEnum.UNAUTHORIZED).JSON(httpMessageEnum.WRONG_PASSWORD)
	}
	jwtPayload := &JwtPayload{
		ID: account.ID,
	}
	genAccessToken, _ := GenAccessToken(jwtPayload)
	bearerToken := "Bearer " + genAccessToken
	newAccessToken := JwtToken{
		AccessToken: bearerToken,
	}
	return c.Status(httpCodeEnum.OK).JSON(newAccessToken)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenAccessToken(payload *JwtPayload) (tokenString string, error error) {
	content := jwt.New(jwt.SigningMethodHS256)
	claims := content.Claims.(jwt.MapClaims)
	claims["username"] = payload.Username
	claims["id"] = payload.ID
	claims["name"] = payload.Name
	claims["email"] = payload.Email
	claims["exp"] = time.Now().Add(time.Hour * 1000).Unix()
	tokenString, error = content.SignedString([]byte("hehe"))
	return
}

func VerifyToken(token *jwt.Token) {
	info, ok := token.Method.(*jwt.SigningMethodHMAC)
	fmt.Println(info, ok)
}

func AuthGuard(c *fiber.Ctx) (user account.Account, err error) {
	bearerToken := c.Get("Authorization")
	token := strings.SplitAfter(bearerToken, "Bearer ")
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
		return user, fiber.NewError(httpCodeEnum.NOT_FOUND, httpMessageEnum.USER_NOT_FOUND)
	}
	user, err = account.GetOne(idFloat64)
	if err != nil {
		return user, fiber.NewError(httpCodeEnum.UNAUTHORIZED, httpMessageEnum.UNAUTHORIZED)
	}
	return user, nil
}
