package handlers

import (
	"time"

	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/backend/config"
	"github.com/fes111/rmm/projects/backend/controller"
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

	exp := time.Now().Add(time.Minute * 5).Unix()
	token, err := createJWTToken(*user, exp)
	if err != nil {
		return err
	}

	refreshToken, err := createJWTToken(*user, 0)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HTTPOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: func() string {
			if config.DEV {
				return "None"
			}
			return "Strict"
		}(),
	})
	return c.JSON(fiber.Map{"user": user, "access_token": token})
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

	exp := time.Now().Add(time.Minute * 5).Unix()
	token, err := createJWTToken(*user, exp)
	if err != nil {
		return err
	}

	refreshToken, err := createJWTToken(*user, 0)
	if err != nil {
		return err
	}
	controller.AddRefreshToken(refreshToken, user.ID)

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HTTPOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: func() string {
			if config.DEV {
				return "None"
			}
			return "Strict"
		}(),
	})
	return c.JSON(fiber.Map{"user": user, "access_token": token})
}

func HandleLogout(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid refresh token")
	}
	token, err := controller.GetRefreshToken(refreshToken)
	if err != nil {
		return err
	}
	controller.ClearRefreshTokens(token.UserID)
	c.ClearCookie("refresh_token")
	return c.SendStatus(fiber.StatusOK)
}

func HandleRefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid refresh token")
	}
	token, err := controller.GetRefreshToken(refreshToken)
	if err != nil {
		return err
	}
	if (token.ExpiresAt - time.Now().Unix()) < 0 {
		return fiber.NewError(fiber.StatusBadRequest, "refresh token expired")
	}
	controller.DeleteRefreshToken(refreshToken)
	c.ClearCookie("refresh_token")

	user, _ := controller.GetUserById(token.UserID)
	exp := time.Now().Add(time.Minute * 5).Unix()
	newToken, err := createJWTToken(user, exp)
	if err != nil {
		return err
	}
	refreshToken, err = createJWTToken(user, 0)
	if err != nil {
		return err
	}
	controller.AddRefreshToken(refreshToken, user.ID)
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HTTPOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: func() string {
			if config.DEV {
				return "None"
			}
			return "Strict"
		}(),
	})
	return c.JSON(fiber.Map{"access_token": newToken})
}

func createJWTToken(user models.User, exp int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user
	claims["iat"] = time.Now().Unix()
	if exp != 0 {
		claims["exp"] = exp
	}
	t, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return t, nil
}
