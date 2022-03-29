package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sonr-io/highway-go/models"
	"github.com/stripe/stripe-go/v72"
)

func (ws *Server) CreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var req struct {
		Items []models.SnrItem `json:"items"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewDecoder.Decode: %v", err)
		return
	}

	if len(req.Items) != 1 {
		//throw error saying not enough items
		return
	}

	pi, err := ws.Ctrl.StripeIntent(req.Items[0], name)
	log.Printf("pi.New: %v", pi.ClientSecret)

	// fmt.Println(pi.Status)
	// fmt.Println(stripe.PaymentIntentStatusSucceeded)

	ws.Ctrl.AttachIntent(pi.ID, name)

	//TODO this is bad
	// go func(item models.SnrItem, name string) {
	// 	time.Sleep(30 * time.Second) //this is based on the stripe timeout 80
	// 	fmt.Print(12345)
	// 	newPi, _ := ws.Ctrl.StripeIntent(item)
	// 	fmt.Println(newPi.Status)
	// 	fmt.Println(stripe.PaymentIntentStatusSucceeded)
	// 	if newPi.Status == stripe.PaymentIntentStatusSucceeded {
	// 		//record payment
	// 		ws.Ctrl.RecordPayment(name)
	// 	}
	// }(req.Items[0], name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("pi.New: %v", err)
		return
	}

	writeJSON(w, struct {
		ClientSecret string `json:"clientSecret"`
	}{
		ClientSecret: pi.ClientSecret,
	})
}

func (ws *Server) StripeWebhook(w http.ResponseWriter, req *http.Request) {
	const MaxBodyBytes = int64(65536)
	req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	event := stripe.Event{}

	if err := json.Unmarshal(payload, &event); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse webhook body json: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println("PaymentIntent was successful!")

		ws.Ctrl.UpdatePayment(paymentIntent.ID)

	case "payment_method.attached":
		var paymentMethod stripe.PaymentMethod
		err := json.Unmarshal(event.Data.Raw, &paymentMethod)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println("PaymentMethod was attached to a Customer!")
	// ... handle other event types
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}

	w.WriteHeader(http.StatusOK)
}
