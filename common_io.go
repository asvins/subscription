package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	authModels "github.com/asvins/auth/models"
	"github.com/asvins/common_db/postgres"
	"github.com/asvins/common_io"
	tm "github.com/asvins/core/models"
	"github.com/asvins/notification/mailer"
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
	consumer.HandleTopic("treatment_created", handleTreatmentCreated)

	if err = consumer.StartListening(); err != nil {
		log.Fatal(err)
	}
}

/*
*	Handlers
 */
func handleTreatmentCreated(msg []byte) {
	t := tm.Treatment{}
	err := json.Unmarshal(msg, &t)

	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}
	var s Subscriber
	db := postgres.GetDatabase(DBConfig())

	err = GetSubscriberByPatient(t.PatientId, &s, db)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}
	if s.PaymentStatus == PaymentStatusOpen {
		m := mailer.Mail{
			To:      []string{s.Email},
			Subject: "Dados Atualizados",
			Body:    mailer.TemplateFinishProfile,
		}
		mailMsg, _ := json.Marshal(m)
		producer.Publish("send_mail", mailMsg)
	}

}

func handleUserCreated(msg []byte) {
	usr := authModels.User{}
	err := json.Unmarshal(msg, &usr)

	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}

	if usr.Scope == "patient" {
		subs := Subscription{Owner: strconv.Itoa(usr.ID), Email: usr.Email}

		db := postgres.GetDatabase(DBConfig())
		if err := subs.Create(db); err != nil {
			fmt.Println("[ERROR] ", err.Error())
			return
		}

		subscriber, err := NewSubscriber(usr.ID, usr.Email, time.Now(), time.Now().AddDate(0, 1, 0), PaymentStatusOpen)
		if err != nil {
			fmt.Println("[ERROR]", err.Error())
			return
		}
		if err := subscriber.Create(db); err != nil {
			fmt.Println("[ERROR] ", err.Error())
			return
		}
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
