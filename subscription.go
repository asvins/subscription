package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func GetSubscription(email string, s *Subscription, db *gorm.DB) error {
	return db.Where("email = ?", email).First(&s).Error
}

func (s *Subscription) Save(db *gorm.DB) error {
	return db.Save(s).Error
}
