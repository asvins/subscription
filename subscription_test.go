package main

import (
	"testing"

	"github.com/asvins/common_db/postgres"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSubscription(t *testing.T) {
	Convey("When creating a subscription", t, func() {
		s, _ := NewSubscription("07051368923", "Rua Luciano Gualberto, 300", "Rua Almeida Prado, 21", "555123459994032", "john.doe@gmail.com", "+5511987726423")
		dbCfg := DBConfig()
		db := postgres.GetDatabase(dbCfg)
		err := s.Create(db)
		Convey("We can create a subscription successfully", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("We can retrieve a saved subscription", func() {
			var newSub Subscription
			GetSubscription("john.doe@gmail.com", &newSub, db)
			So(newSub.CPF, ShouldEqual, s.CPF)
		})
		Convey("We can update an already saved subscription", func() {
			var newSub Subscription
			s.CPF = "newcpf"
			s.Save(db)
			GetSubscription("john.doe@gmail.com", &newSub, db)
			So(newSub.CPF, ShouldEqual, "newcpf")
		})
		db.Exec("TRUNCATE TABLE subscriptions")
	})
}
