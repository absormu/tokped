package job

import (
	"net/http"

	"github.com/absormu/tokped/app/entity"
	md "github.com/absormu/tokped/app/middleware"
	repoorder "github.com/absormu/tokped/app/repository/order"
	lg "github.com/absormu/tokped/pkg/response"
	resp "github.com/absormu/tokped/pkg/response"
	sdk "github.com/absormu/tokped/pkg/sdk"
	"github.com/labstack/echo/v4"
)

func GetOrderHistory(c echo.Context, extractToken entity.ExtractToken) (e error) {
	logger := md.GetLogger(c)
	logger.WithField("request", extractToken).Info("usecase: GetOrderHistory")

	idStr := c.Param("id")

	var order entity.OrderHistory
	if order, e = repoorder.GetOrderHistory(c, idStr); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error failure query GetOrderHistory")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
			lg.Language{Bahasa: "Failure query get order ID", English: "Failure query get order ID"}, nil, nil)
		return
	}

	if order.OrderNo == "" {
		logger.Error("Catch error order history not found")
		e = resp.CustomError(c, http.StatusNotFound, sdk.ERR_UNAUTHORIZED,
			lg.Language{Bahasa: "Order history tidak tersedia", English: "Order history not found"}, nil, nil)
		return
	}

	var data entity.OrderHistoryData
	data.OrderNo = order.OrderNo

	data.Buyer.ID = order.Buyer.ID
	data.Buyer.Name = order.Buyer.Name

	data.Seller.ID = order.Seller.ID
	data.Seller.Name = order.Seller.Name

	data.Shipping.Name = order.Shipping.Name
	data.Shipping.Phone = order.Shipping.Phone
	data.Shipping.Address = order.Shipping.Address

	// "order_details": [

	data.Logistic.ID = order.Logistic.ID
	data.Logistic.Name = order.Logistic.Name
	data.Logistic.ShippingCost = order.Logistic.ShippingCost

	data.PaymentMethod.ID = order.PaymentMethod.ID
	data.PaymentMethod.Name = order.PaymentMethod.Name

	data.Total.Quantity = order.Total.Quantity
	data.Total.Weight = order.Total.Weight
	data.Total.ProductAmount = order.Total.ProductAmount
	data.Total.ShoppingAmount = order.Total.ShoppingAmount
	data.Total.ServiceCharge = order.Total.ServiceCharge
	data.Total.Amount = order.Total.Amount

	data.CreatedAt = order.CreatedAt

	e = resp.CustomError(c, http.StatusOK, sdk.ERR_SUCCESS,
		lg.Language{Bahasa: "Sukses", English: "Success"}, nil, data)
	return
}
