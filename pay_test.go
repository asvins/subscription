package main

import (
	"testing"
	"time"

	"github.com/asvins/common_db/postgres"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPayment(t *testing.T) {
	Convey("When paying a montly subscription", t, func() {
		p := PaymentManager{}
		lastPayed := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
		nextPayment := time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC)
		s, _ := NewSubscriber(email, lastPayed, nextPayment, PaymentStatusDelayed)
		dbCfg := DBConfig()
		db := postgres.GetDatabase(dbCfg)
		s.Create(db)
		sub, _ := NewSubscription("07051368923", "Rua Luciano Gualberto, 300", "Rua Almeida Prado, 21", "555123459994032", "john.doe@gmail.com", "+5511987726423")
		sub.Create(db)
		Convey("We can pay a subscription", func() {
			err := p.Pay(s, *sub, db)
			So(err, ShouldEqual, nil)
		})
		db.Exec("TRUNCATE TABLE subscribers;")
		db.Exec("TRUNCATE TABLE subscriptions;")
	})
}
