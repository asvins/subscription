package main

import (
	"net/http"
	"strings"

	"github.com/asvins/router"
	"github.com/asvins/router/logger"
	"github.com/unrolled/render"
)

func DiscoveryHandler(w http.ResponseWriter, req *http.Request) {
	prefix := strings.Join([]string{ServerConfig.Server.Addr, ServerConfig.Server.Port}, ":")
	r := render.New()

	//add discovery links here
	discoveryMap := map[string]string{"discovery": prefix + "/api/discovery"}

	r.JSON(w, http.StatusOK, discoveryMap)
}

func DefRoutes() *router.Router {
	r := router.NewRouter()

	//subscription
	r.AddRoute("/api/discovery", router.GET, DiscoveryHandler)
	r.AddRoute("/api/subscription/show", router.GET, SubscriptionShowHandler)
	r.AddRoute("/api/subscription/list", router.GET, SubscriptionListHandler)
	r.AddRoute("/api/subscription/new", router.POST, SubscriptionNewHandler)
	r.AddRoute("/api/subscription/pay", router.POST, PayHandler)
	r.Handle("/api/subscription", router.PUT, updateSubscription, []router.Interceptor{})

	//subscriber
	r.AddRoute("/api/subscriber/:id/paymentstatus", router.GET, retrievePaymentStatus)

	// interceptors
	r.AddBaseInterceptor("/", logger.NewLogger())
	return r
}
