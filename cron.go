package main

import (
	"fmt"
	"time"

	"github.com/asvins/common_db/postgres"
)

func startPaymentVerificationCron() {
	fmt.Println("[INFO] Starting payment verification cron")
	go func() {
		systemPaymentVerification()
		for {
			<-time.After(time.Hour * 24)
			systemPaymentVerification()
		}
	}()
}

func systemPaymentVerification() {
	subs := Subscriber{}
	db := postgres.GetDatabase(DBConfig())

	subscribers, err := subs.Retrieve(db)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}
	for _, currSubscriber := range subscribers {
		if time.Now().After(currSubscriber.NextPayment) {
			currSubscriber.PaymentStatus = PaymentStatusDelayed
			if err := currSubscriber.Update(db); err != nil {
				fmt.Println("[ERROR] ", err.Error())
				continue
			}
		}
	}
}
