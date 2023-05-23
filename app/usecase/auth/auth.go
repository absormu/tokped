package auth

import (
	"net/http"
	"time"

	"github.com/absormu/tokped/app/entity"
	md "github.com/absormu/tokped/app/middleware"
	repoauth "github.com/absormu/tokped/app/repository/auth"
	cm "github.com/absormu/tokped/pkg/configuration"
	lg "github.com/absormu/tokped/pkg/response"
	resp "github.com/absormu/tokped/pkg/response"
	sdk "github.com/absormu/tokped/pkg/sdk"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context, req entity.Auth) (e error) {
	logger := md.GetLogger(c)
	logger.WithField("request", req).Info("usecase: Login")

	// Throws unauthorized error
	if req.Email == "" || req.Password == "" {
		logger.Error("Catch error missing mandatory parameter")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_MISSING,
			lg.Language{Bahasa: nil, English: "Missing mandatory parameter"}, nil, nil)
		return
	}

	// cek email & get password
	params := make(map[string]string)
	params["email"] = req.Email
	params["active"] = "1"
	params["is_deleted"] = "0"

	var user entity.User

	if user, e = repoauth.GetAuthEmail(c, params); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error failure query GetAuthEmail")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
			lg.Language{Bahasa: nil, English: "Failure query"}, nil, nil)
		return
	}

	// login bycrypt
	passDB := user.Password
	password := req.Password
	match := CheckPasswordHash(password, passDB)

	if !match {
		logger.Error("Catch error user not found")
		e = resp.CustomError(c, http.StatusUnauthorized, sdk.ERR_USER_NOT_FOUND,
			lg.Language{Bahasa: "Email atau kata sandi salah", English: "Email or password is not correct"}, nil, nil)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = xid.New().String()
	claims["user_id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = req.Email
	claims["user_contact_id"] = user.UserContactID
	claims["role_id"] = user.RoleID
	claims["exp"] = time.Now().Add(time.Duration(cm.Config.TokenLifeTime) * time.Second).Unix()

	// Generate encoded token and send it as response.
	t, e := token.SignedString([]byte(cm.Config.ClientSecret))
	if e != nil {
		logger.WithField("error", e.Error()).Error("Catch error generate encoded token")
		e = resp.CustomError(c, http.StatusUnauthorized, sdk.ERR_UNAUTHORIZED,
			lg.Language{Bahasa: nil, English: "Unauthorized"}, nil, nil)
		return
	}

	res := entity.OAuthMessage{
		AccessToken: t,
		TokenType:   "bearer",
		ExpiresIn:   cm.Config.TokenLifeTime,
	}

	e = resp.CustomError(c, http.StatusOK, sdk.ERR_SUCCESS,
		lg.Language{Bahasa: "Sukses", English: "Success"}, nil, res)
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
