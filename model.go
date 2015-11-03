package main

import (
	"errors"
	"time"
)

const (
  PaymentStatusOK = iota
	PaymentStatusDelayed
	PaymentStatusOpen
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

func NewSubscription(userId, name, cpf, address, deliveryAddress, creditCardNumber, email, phone string) (*Subscription, error) {
	if cpf == "" || email == "" || deliveryAddress == "" || creditCardNumber == "" {
    return nil, errors.New("Invalid Input")
	}
	return &Subscription{Name: name, CPF: cpf, Address: address, DeliveryAddress: deliveryAddress, CreditCardNumber: creditCardNumber, Email: email, Phone: phone}, nil
}

type Subscriber struct {
	Email string `json:"email"`
	LastPayed time.Time `json:"last_payed"`
	NextPayment time.Time `json:"next_payment"`
	PaymentStatus int `json:"payment_status"`
}

func NewSubscriber(email string, lastPayed, nextPayment time.Time, paymentStatus int) (*Subscriber, error) {
  if email == "" {
    return nil, errors.New("Invalid Input")
	}
	return &Subscriber{Email: email, LastPayed: lastPayed, NextPayment: nextPayment, PaymentStatus: paymentStatus}, nil
}

type PaymentManager interface {
	Pay(*Subscriber, Subscription) error
}
