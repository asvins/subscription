package models

const (
	PaymentStatusOpen = iota
	PaymentStatusDelayed
	PaymentStatusOK
)

type Subscription struct {
	CPF              string `json:"cpf" gorm:"column:cpf"`
	Address          string `json:"address"`
	DeliveryAddress  string `json:"delivery_address"`
	CreditCardNumber string `json:"credit_card" gorm:"column:credit_card"`
	Email            string `json:"email" gorm:"column:email;primary_key"`
	Phone            string `json:"phone"`
}
