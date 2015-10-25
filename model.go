package main

import (
	"errors"
)

type Subscription struct {
	Name             string `json:"name"`
	CPF              string `json:"cpf"`
	Address          string `json:"address"`
	DeliveryAddress  string `json:"delivery_address"`
	CreditCardNumber string `json:"credit_card"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
}

func NewSubscription(userId, name, cpf, address, deliveryAddress, creditCardNumber, email, phone string) *Subscription, error {
	if cpf == "" || email == "" || DeliveryAddress == "" || CreditCardNumber == "" {
    return nil, errors.New("Invalid Input")
	}
	return &Subscription{UserID: userId, Name: name, CPF: cpf, Address: address, DeliveryAddress: deliveryAddress, CreditCardNumber: creditCardNumber, Email: email, Phone: phone}, nil
}
