package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	authModels "github.com/asvins/auth/models"
	"github.com/asvins/common_db/postgres"
	"github.com/asvins/common_io"
	"github.com/asvins/utils/config"
)

func setupCommonIo() {
	cfg := common_io.Config{}

	err := config.Load("common_io_config.gcfg", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	/*
	*	Producer
	 */
	producer, err = common_io.NewProducer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	/*
	*	Consumer
	 */
	consumer = common_io.NewConsumer(cfg)

	/*
	*	Topics
	 */
	consumer.HandleTopic("user_created", handleUserCreated)

	if err = consumer.StartListening(); err != nil {
		log.Fatal(err)
	}
}

/*
*	Handlers
 */
func handleUserCreated(msg []byte) {
	usr := authModels.User{}
	err := json.Unmarshal(msg, &usr)

	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
	}

	if usr.Scope == "patient" {
		subs := Subscription{Owner: strconv.Itoa(usr.ID), Email: usr.Email}

		db := postgres.GetDatabase(DBConfig())
		subs.Create(db)
	}
}

func sendSubscriptionPaid(subscription *Subscription) {
	b, err := json.Marshal(subscription)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}

	producer.Publish("subscription_paid", b)
}

func sendSubscriptionUpdated(subscription *Subscription) {
	b, err := json.Marshal(subscription)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}

	producer.Publish("subscription_updated", b)
}
