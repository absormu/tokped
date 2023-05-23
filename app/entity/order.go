package entity

type OrderHistory struct {
	OrderNo       string        `json:"-"`
	Seller        Seller        `json:"-"`
	Buyer         Buyer         `json:"-"`
	Shipping      Shipping      `json:"-"`
	OrderDetails  OrderDetail   `json:"-"`
	Logistic      Logistic      `json:"-"`
	PaymentMethod PaymentMethod `json:"-"`
	Total         Total         `json:"-"`
	CreatedAt     string        `json:"-"`
}

type OrderHistoryData struct {
	OrderNo       string        `json:"order_no,omitempty"`
	Seller        Seller        `json:"seller,omitempty"`
	Buyer         Buyer         `json:"buyer,omitempty"`
	Shipping      Shipping      `json:"shipping,omitempty"`
	OrderDetails  []OrderDetail `json:"order_details,omitempty"`
	Logistic      Logistic      `json:"logistic,omitempty"`
	PaymentMethod PaymentMethod `json:"payment_method,omitempty"`
	Total         Total         `json:"total,omitempty"`
	CreatedAt     string        `json:"created_at,omitempty"`
}

type Seller struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Buyer struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Shipping struct {
	Name    string `json:"name,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Address string `json:"address,omitempty"`
}

type OrderDetail struct {
	ProductID     int    `json:"product_id,omitempty"`
	ProductName   string `json:"product_name,omitempty"`
	ProductWeight int    `json:"product_weight,omitempty"`
	ProductPrice  int    `json:"product_price,omitempty"`
	Quantity      int    `json:"quantity,omitempty"`
	TotalAmount   int    `json:"total_amount,omitempty"`
}

type Logistic struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	ShippingCost int    `json:"shipping_cost,omitempty"`
}

type PaymentMethod struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Total struct {
	Quantity       int `json:"quantity,omitempty"`
	Weight         int `json:"weight,omitempty"`
	ProductAmount  int `json:"product_amount,omitempty"`
	ShoppingAmount int `json:"shopping_amount,omitempty"`
	ServiceCharge  int `json:"service_charge,omitempty"`
	Amount         int `json:"amount,omitempty"`
}
