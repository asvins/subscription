package main

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

const (
	PaymentStatusOpen = iota
	PaymentStatusDelayed
	PaymentStatusOK
)

type Subscription struct {
	CPF              string `json:"cpf" gorm:"column:cpf"`
	Owner            string `json:"owner" gorm:"column:owner"`
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
	PatientId     int       `json:"patient_id"`
	Email         string    `json:"email" gorm:"column:email;primary_key"`
	LastPayed     time.Time `json:"last_payed"`
	NextPayment   time.Time `json:"next_payment"`
	PaymentStatus int       `json:"payment_status" gorm:"column:payment_status;default:0"`
}

func NewSubscriber(patientId int, email string, lastPayed, nextPayment time.Time, paymentStatus int) (*Subscriber, error) {
	if email == "" {
		return nil, errors.New("Invalid Input")
	}
	return &Subscriber{PatientId: patientId, Email: email, LastPayed: lastPayed, NextPayment: nextPayment, PaymentStatus: paymentStatus}, nil
}

func (s *Subscriber) RetrieveSubscriber(db *gorm.DB) (*Subscriber, error) {
	subscribers := []Subscriber{}
	if err := db.Where(s).Find(&subscribers).Error; err != nil {
		return nil, err
	}

	if len(subscribers) != 1 {
		return nil, errors.New("[FATAL] Database is inconsistent. More then one subscriber associated with the same pacient_id")
	}

	return &subscribers[0], nil
}

type Payer interface {
	Pay(*Subscriber, Subscription) error
}
