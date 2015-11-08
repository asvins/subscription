package main

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type PaymentManager struct{}

func (p PaymentManager) Pay(s *Subscriber, sub Subscription, db *gorm.DB) error {
	if !p.isValidCreditCard(sub) {
		return errors.New("Invalid Input")
	}
	s.NextPayment = s.LastPayed.AddDate(0, 1, 0)
	s.LastPayed = time.Now()
	return s.Save(db)
}

// mocking credit card validation
func (p PaymentManager) isValidCreditCard(sub Subscription) bool {
	return true
}
