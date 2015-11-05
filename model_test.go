package main

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	email = "john.doe@example.com"
)

func TestNewSubscriber(t *testing.T) {
	lastPayed := time.Now()
	nextPayment := lastPayed.Add(24 * time.Hour)
	Convey("When creating a new subscriber", t, func() {
		s, err := NewSubscriber(email, lastPayed, nextPayment, PaymentStatusOpen)
		Convey("When email is not nil, we don't get an error", func() {
			So(err, ShouldEqual, nil)
			So(s.Email, ShouldEqual, email)
		})
		s, err = NewSubscriber("", lastPayed, nextPayment, PaymentStatusOpen)
		Convey("When email is nil, we get an error", func() {
			So(err, ShouldNotEqual, nil)
		})
	})
}
