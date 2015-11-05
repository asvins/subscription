package main

import (
	"fmt"
	"testing"

	"github.com/asvins/common_db/postgres"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSubscription(t *testing.T) {
	Convey("When creating a subscription", t, func() {
		s, _ := NewSubscription("07051368923", "Rua Luciano Gualberto, 300", "Rua Almeida Prado, 21", "555123459994032", "john.doe@gmail.com", "+5511987726423")
		dbCfg := DBConfig()
		fmt.Println("%v", dbCfg)
		db := postgres.GetDatabase(dbCfg)
		err := s.Save(db)
		Convey("Subscription is saved successfully", func() {
			So(err, ShouldEqual, nil)
		})
	})
}
