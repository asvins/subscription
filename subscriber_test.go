package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/asvins/common_db/postgres"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSubscriber(t *testing.T) {
	Convey("When creating a subscriber", t, func() {
		email := "john.doe@example.com"
		lastPayed := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
		nextPayment := time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC)
		s, _ := NewSubscriber(email, lastPayed, nextPayment, PaymentStatusDelayed)
		dbCfg := DBConfig()
		db := postgres.GetDatabase(dbCfg)
		err := s.Create(db)
		Convey("We can create a subscriber successfully", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("We can retrieve a saved subscriber", func() {
			var newSub Subscriber
			GetSubscriber(email, &newSub, db)
			fmt.Println(newSub)
			So(newSub.LastPayed.Unix(), ShouldEqual, s.LastPayed.Unix())
		})
		Convey("We can update an already saved subscriber", func() {
			var newSub Subscriber
			t := time.Date(2015, time.November, 10, 23, 0, 0, 0, time.UTC)
			s.LastPayed = t
			s.NextPayment = t.AddDate(0, 1, 0)
			s.Save(db)
			GetSubscriber(email, &newSub, db)
			So(newSub.LastPayed.Unix(), ShouldEqual, t.Unix())
		})
		db.Exec("TRUNCATE TABLE subscribers")
	})
}
