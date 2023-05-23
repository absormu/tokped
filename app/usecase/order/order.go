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

	var orders []entity.OrderHistory
	if orders, e = repoorder.GetOrderHistory(c, idStr); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error failure query GetOrderHistory")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
			lg.Language{Bahasa: "Failure query get order ID", English: "Failure query get order ID"}, nil, nil)
		return
	}

	if len(orders) == 0 {
		logger.Error("Catch error order history not found")
		e = resp.CustomError(c, http.StatusNotFound, sdk.ERR_UNAUTHORIZED,
			lg.Language{Bahasa: "Order history tidak tersedia", English: "Order history not found"}, nil, nil)
		return
	}

	var orderDetailData []entity.OrderDetail
	var orderDetail entity.OrderDetail
	var orderNo, buyerName, sellerName, shippingName, shippingPhone, shippingAddress string
	var logisticID, logisticName, paymentMethodID, paymentMethodName, createdAt string
	var buyerID, sellerID, logisticShippingCost int
	var totalQuantity, totalWeight, totalProductAmount, totalShoppingAmount, totalServiceCharge, totalAmount int

	for _, order := range orders {
		orderNo = order.OrderNo
		buyerID = order.Buyer.ID
		buyerName = order.Buyer.Name

		sellerID = order.Seller.ID
		sellerName = order.Seller.Name

		shippingName = order.Shipping.Name
		shippingPhone = order.Shipping.Phone
		shippingAddress = order.Shipping.Address

		orderDetail.ProductName = order.OrderDetails.ProductName
		orderDetail.Quantity = order.OrderDetails.Quantity
		orderDetail.ProductWeight = order.OrderDetails.ProductWeight
		orderDetail.ProductPrice = order.OrderDetails.ProductPrice
		orderDetail.TotalAmount = order.OrderDetails.TotalAmount

		logisticID = order.Logistic.ID
		logisticName = order.Logistic.Name
		logisticShippingCost = order.Logistic.ShippingCost

		paymentMethodID = order.PaymentMethod.ID
		paymentMethodName = order.PaymentMethod.Name

		totalQuantity = order.Total.Quantity
		totalWeight = order.Total.Weight
		totalProductAmount = order.Total.ProductAmount
		totalShoppingAmount = order.Total.ShoppingAmount
		totalServiceCharge = order.Total.ServiceCharge
		totalAmount = order.Total.Amount

		createdAt = order.CreatedAt

		orderDetailData = append(orderDetailData, orderDetail)
	}

	var data entity.OrderHistoryData
	data.OrderNo = orderNo
	data.Buyer.ID = buyerID
	data.Buyer.Name = buyerName

	data.Seller.ID = sellerID
	data.Seller.Name = sellerName

	data.Shipping.Name = shippingName
	data.Shipping.Phone = shippingPhone
	data.Shipping.Address = shippingAddress

	data.OrderDetails = orderDetailData

	data.Logistic.ID = logisticID
	data.Logistic.Name = logisticName
	data.Logistic.ShippingCost = logisticShippingCost

	data.PaymentMethod.ID = paymentMethodID
	data.PaymentMethod.Name = paymentMethodName

	data.Total.Quantity = totalQuantity
	data.Total.Weight = totalWeight
	data.Total.ProductAmount = totalProductAmount
	data.Total.ShoppingAmount = totalShoppingAmount
	data.Total.ServiceCharge = totalServiceCharge
	data.Total.Amount = totalAmount

	data.CreatedAt = createdAt

	e = resp.CustomError(c, http.StatusOK, sdk.ERR_SUCCESS,
		lg.Language{Bahasa: "Sukses", English: "Success"}, nil, data)
	return
}
