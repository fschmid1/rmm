package handlers

import (
	"time"

	"github.com/fes111/rmm/projects/rmm/go/backend/config"
	"github.com/fes111/rmm/projects/rmm/go/lib/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleSignUp(c *fiber.Ctx) error {
	req := new(SignupRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid signup credentials")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}

	err = config.Database.Transaction(func(tx *gorm.DB) error {
		if result := tx.Create(&user); result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		return err
	}
	token, exp, err := createJWTToken(*user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
}

func HandleLogin(c *fiber.Ctx) error {
	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	if req.Email == "" || req.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
	}

	user := new(models.User)
	result := config.Database.Where("email = ?", req.Email).First(&user)
	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, exp, err := createJWTToken(*user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
}

func createJWTToken(user models.User) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user
	claims["exp"] = exp
	t, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "", 0, err
	}

	return t, exp, nil
}
