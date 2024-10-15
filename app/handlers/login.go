package handlers

import (
	"database/sql"
	"encoding/base64"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mr-utzig/gear-stock/app/model"
	"github.com/mr-utzig/gear-stock/app/utils"
	"github.com/mr-utzig/gear-stock/db"
)

type JwtClaims struct {
	UserID    int    `json:"id"`
	UserName  string `json:"name"`
	UserEmail string `json:"email"`
	jwt.RegisteredClaims
}

func LoginPage(c echo.Context) error {
	token, _ := c.Cookie("tkn")
	if token != nil {
		return c.Redirect(http.StatusPermanentRedirect, "/")
	}

	return c.Render(http.StatusOK, "login", nil)
}

func LoginValidation(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		log.Error("handlers -> login -> Validate -> c.Bind", err)
		return c.HTML(http.StatusBadRequest, utils.MessageBadRequest)
	}

	db := db.Connection()
	defer db.Close()

	err := u.Validate(db)
	if err == sql.ErrNoRows {
		return c.HTML(http.StatusUnauthorized, utils.MessageUserNotFound)
	} else if err != nil {
		log.Error("handlers -> login -> Validate -> u.Validate", err)

		return c.HTML(http.StatusInternalServerError, utils.MessageInternalServerError)
	}

	token, err := createTokenJWT(u)
	if err != nil {
		log.Error("handlers -> login -> Validate -> createTokenJWT", err)

		return c.HTML(http.StatusInternalServerError, utils.MessageInternalServerError)
	}

	c.SetCookie(&http.Cookie{
		Name:     "tkn",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return c.HTML(http.StatusOK, utils.MessageLoginSuccess)
}

func createTokenJWT(u *model.User) (string, error) {
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, JwtClaims{
		u.ID,
		u.Name,
		u.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	SigningKey, _ := base64.StdEncoding.DecodeString(os.Getenv("JWT_SIGING_KEY"))
	token, err := rawToken.SignedString(SigningKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
