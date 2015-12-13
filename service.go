package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/asvins/common_db/postgres"
	"github.com/asvins/router/errors"
	"github.com/unrolled/render"
)

func SubscriptionShowHandler(w http.ResponseWriter, req *http.Request) {
	r := render.New()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, HEAD")
	w.Header().Add("Access-Control-Allow-Headers", "X-PINGOTHER, Origin, X-Requested-With, Content-Type, Accept")

	email := req.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Invalid Input", 400)
		return
	}

	db := postgres.GetDatabase(DBConfig())
	var s Subscription
	err := GetSubscription(email, &s, db)
	if err != nil {
		http.Error(w, "Not Found", 404)
		return
	}

	r.JSON(w, http.StatusOK, s)
}

func SubscriptionListHandler(w http.ResponseWriter, req *http.Request) {
	r := render.New()
	db := postgres.GetDatabase(DBConfig())

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, HEAD")
	w.Header().Add("Access-Control-Allow-Headers", "X-PINGOTHER, Origin, X-Requested-With, Content-Type, Accept")

	page, err := strconv.Atoi(req.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}

	var subs []Subscription
	err = GetSubscriptions(page, &subs, db)
	if err != nil {
		http.Error(w, "Not found", 404)
		return
	}

	r.JSON(w, http.StatusOK, subs)
}

func SubscriptionNewHandler(w http.ResponseWriter, req *http.Request) {
	r := render.New()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, HEAD")
	w.Header().Add("Access-Control-Allow-Headers", "X-PINGOTHER, Origin, X-Requested-With, Content-Type, Accept")

	db := postgres.GetDatabase(DBConfig())
	var s Subscription

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&s); err != nil {
		http.Error(w, "Invalid parameters", 400)
		return
	}

	sub, err := NewSubscriber(0, s.Email, time.Now(), time.Now().AddDate(0, 1, 0), PaymentStatusOK)
	if s.Create(db) != nil || sub.Create(db) != nil || err != nil {
		http.Error(w, "Service Unavailable", 503)
		return
	}

	b, err := json.Marshal(&s)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		producer.Publish("subscription_updated", b)
	}

	r.JSON(w, http.StatusCreated, "{}")
}

func PayHandler(w http.ResponseWriter, req *http.Request) {
	r := render.New()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, HEAD")
	w.Header().Add("Access-Control-Allow-Headers", "X-PINGOTHER, Origin, X-Requested-With, Content-Type, Accept")

	db := postgres.GetDatabase(DBConfig())

	if req.ParseForm() != nil || req.FormValue("email") == "" {
		http.Error(w, "Invalid Input", 400)
		return
	}

	email := req.FormValue("email")
	var sub Subscription
	var s Subscriber
	var p PaymentManager

	if GetSubscription(email, &sub, db) != nil || GetSubscriber(email, &s, db) != nil || p.Pay(&s, sub, db) != nil {
		http.Error(w, "Not Found", 404)
		return
	}

	sendSubscriptionUpdated(&sub)
	sendSubscriptionPaid(&sub)

	r.JSON(w, http.StatusOK, "{}")
}

func retrievePaymentStatus(w http.ResponseWriter, req *http.Request) {
	r := render.New()
	db := postgres.GetDatabase(DBConfig())
	subs := &Subscriber{}

	patient_id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	subs.PatientId = patient_id
	subs, err = subs.RetrieveSubscriber(db)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if subs.PaymentStatus == PaymentStatusOK {
		r.JSON(w, http.StatusOK, "{}")
	} else {
		r.JSON(w, http.StatusNotFound, "{}")
	}
}

func updateSubscription(w http.ResponseWriter, req *http.Request) errors.Http {
	subs := Subscription{}

	if err := BuildStructFromReqBody(&subs, req.Body); err != nil {
		return errors.BadRequest("[ERROR] Malformed request body")
	}

	email := req.URL.Query().Get("email")
	if email == "" {
		return errors.BadRequest("[ERROR] email is mandatory")
	}

	db := postgres.GetDatabase(DBConfig())
	if err := subs.UpdateByEmail(email, db); err != nil {
		return errors.InternalServerError(err.Error())
	}

	rend := render.New()
	rend.JSON(w, http.StatusOK, subs)
	return nil
}
