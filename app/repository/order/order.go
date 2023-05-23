package report

import (
	entity "github.com/absormu/tokped/app/entity"
	md "github.com/absormu/tokped/app/middleware"
	db "github.com/absormu/tokped/pkg/mariadb"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetOrderHistory(c echo.Context, orderID string) (order entity.OrderHistory, e error) {
	logger := md.GetLogger(c)
	logger.WithFields(logrus.Fields{
		"params": orderID,
	}).Info("repository: GetOrderHistory")

	db := db.MariaDBInit()

	defer db.Close()
	query := "SELECT o.id, o.created_at, o.seller_id, o.seller_name, o.buyer_id, o.buyer_name, " +
		"o.shipping_name, o.shipping_phone, o.shipping_address, " +
		"o.logistic_id, o.logistic_name, o.total_shipping_cost, " +
		"o.payment_method_id, o.payment_method_name, " +
		"o.total_quantity, o.total_weight, " +
		"o.total_product_amount, o.total_shopping_amount, " +
		"o.service_charge, o.total_amount " +
		"FROM orders o " +
		"JOIN order_details od ON o.id = od.order_id WHERE o.id = '" + orderID + "'"

	logger.WithFields(logrus.Fields{"query": query}).Info("repository: GetOrderHistory-query")

	result, e := db.Query(query)
	if e != nil {
		return
	}

	defer result.Close()
	for result.Next() {
		if e = result.Scan(
			&order.OrderNo, &order.CreatedAt, &order.Seller.ID, &order.Seller.Name, &order.Buyer.ID, &order.Buyer.Name,
			&order.Shipping.Name, &order.Shipping.Phone, &order.Shipping.Address,
			&order.Logistic.ID, &order.Logistic.Name, &order.Logistic.ShippingCost,
			&order.PaymentMethod.ID, &order.PaymentMethod.Name,
			&order.Total.Quantity, &order.Total.Weight,
			&order.Total.ProductAmount, &order.Total.ShoppingAmount,
			&order.Total.ServiceCharge, &order.Total.Amount); e != nil {
			return
		}
	}

	return
}
