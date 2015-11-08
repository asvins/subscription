package main

import (
	"errors"
	"time"
)

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

func NewSubscription(cpf, address, deliveryAddress, creditCardNumber, email, phone string) (*Subscription, error) {
	if cpf == "" || email == "" || deliveryAddress == "" || creditCardNumber == "" {
		return nil, errors.New("Invalid Input")
	}
	return &Subscription{CPF: cpf, Address: address, DeliveryAddress: deliveryAddress, CreditCardNumber: creditCardNumber, Email: email, Phone: phone}, nil
}

type Subscriber struct {
	Email         string    `json:"email" gorm:"column:email;primary_key"`
	LastPayed     time.Time `json:"last_payed"`
	NextPayment   time.Time `json:"next_payment"`
	PaymentStatus int       `json:"payment_status" gorm:"column:payment_status;default:0"`
}

func NewSubscriber(email string, lastPayed, nextPayment time.Time, paymentStatus int) (*Subscriber, error) {
	if email == "" {
		return nil, errors.New("Invalid Input")
	}
	return &Subscriber{Email: email, LastPayed: lastPayed, NextPayment: nextPayment, PaymentStatus: paymentStatus}, nil
}

type Payer interface {
	Pay(*Subscriber, Subscription) error
}
