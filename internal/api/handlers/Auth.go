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
	"github.com/mr-utzig/gear-stock/internal/api/models"
	"github.com/mr-utzig/gear-stock/internal/api/utils"
	"golang.org/x/crypto/bcrypt"
)

type JwtClaims struct {
	UserID        int64  `json:"id"`
	UserName      string `json:"name"`
	UserEmail     string `json:"email"`
	UserProfileID int    `json:"profile_id"`
	jwt.RegisteredClaims
}

type LoginResponse struct {
	User         models.UserDataResponse `json:"user"`
	FirstAccess  bool                    `json:"first_access"`
	SessionToken *string                 `json:"token"`
}

type AuthHandler struct {
	userModel *models.UserModel
}

func NewAuthHandler(userModel *models.UserModel) *AuthHandler {
	return &AuthHandler{userModel: userModel}
}

func (a *AuthHandler) HandleLogin(c echo.Context) error {
	data := new(models.UserLoginRequest)
	if err := c.Bind(data); err != nil {
		log.Error("handlers -> Auth -> HandleLogin -> c.Bind", err)
		return c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Bad Request", data))
	}

	user, err := a.userModel.ValidateCredentials(data)
	if err == sql.ErrNoRows || err == bcrypt.ErrMismatchedHashAndPassword {
		return c.JSON(http.StatusUnauthorized, utils.NewResponse(false, "Unauthorized", data))
	} else if err != nil {
		log.Error("handlers -> Auth -> HandleLogin -> ValidateCredentials", err, data)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", data))
	}

	token, err := createTokenJWT(user)
	if err != nil {
		log.Error("handlers -> Auth -> HandleLogin -> createTokenJWT", err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", data))
	}

	user.SessionToken = token
	if err := a.userModel.UpdateSessionToken(user.ID, user.SessionToken); err != nil {
		log.Error("handlers -> Auth -> HandleLogin -> UpdateSessionToken", err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", data))
	}

	response := LoginResponse{
		User: models.UserDataResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			ProfileID: user.ProfileID,
			Status:    user.Status,
		},
		SessionToken: user.SessionToken,
		FirstAccess:  user.FirstAccess,
	}

	return c.JSON(http.StatusOK, utils.NewResponse(true, "OK", response))
}

func createTokenJWT(u *models.User) (*string, error) {
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, JwtClaims{
		u.ID,
		u.Name,
		u.Email,
		u.ProfileID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	SigningKey, _ := base64.StdEncoding.DecodeString(os.Getenv("JWT_SIGING_KEY"))
	token, err := rawToken.SignedString(SigningKey)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
