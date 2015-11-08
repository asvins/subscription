package main

import (
	"net/http"
	"strconv"

	"github.com/asvins/common_db/postgres"
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

	//db := postgres.GetDatabase(DBConfig())

	r.JSON(w, http.StatusCreated, "{}")
}
