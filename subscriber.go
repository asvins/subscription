package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func GetSubscriber(email string, s *Subscriber, db *gorm.DB) error {
	return db.Where("email = ?", email).First(&s).Error
}

func GetSubscriberByPatient(email string, s *Subscriber, db *gorm.DB) error {
	return db.Where("owner = ?", email).First(&s).Error
}

func (s *Subscriber) GetSubscription(sub *Subscription, db *gorm.DB) error {
	return GetSubscription(s.Email, sub, db)
}

func (s *Subscriber) Create(db *gorm.DB) error {
	return db.FirstOrCreate(s).Error
}

func (s *Subscriber) Save(db *gorm.DB) error {
	return db.Save(s).Error
}
