package handler

import (
	"net/http"

	md "github.com/absormu/tokped/app/middleware"
	usecaseorder "github.com/absormu/tokped/app/usecase/order"
	pkgjwt "github.com/absormu/tokped/pkg/jwt"
	lg "github.com/absormu/tokped/pkg/response"
	resp "github.com/absormu/tokped/pkg/response"
	sdk "github.com/absormu/tokped/pkg/sdk"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func GetOrderHistoryHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: GetOrderHistoryHandler")

	extractToken, e := pkgjwt.ExtractToken(c)
	if e != nil {
		logger.Error("Catch error extractToken")
		e = resp.CustomError(c, http.StatusUnauthorized, sdk.ERR_UNAUTHORIZED,
			lg.Language{Bahasa: "Authorization missing", English: "Authorization missing"}, nil, nil)
		return
	}

	logger.WithFields(log.Fields{
		"extractToken": extractToken,
	}).Info("ExtractToken")

	e = usecaseorder.GetOrderHistory(c, extractToken)

	return
}
