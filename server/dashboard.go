package server

import (
	"net/http"

	"github.com/sonr-io/highway-go/models"
)

// Index renders the dashboard index page, displaying the created credential
// as well as any other credentials previously registered by the authenticated
// user.
func (ws *Server) Index(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)
	credentials, err := ws.Ctrl.GetCredentialsForUser(&user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	templateData := struct {
		User        models.User
		Credentials []models.Credential
	}{
		user,
		credentials,
	}
	renderTemplate(w, "index.html", templateData)
}

func (ws *Server) PaymentPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "payment.html", nil)
}

func (ws *Server) RoadMapPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "roadmap.html", nil)
}

func (ws *Server) CheckoutPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "success.html", nil)
}

// Login renders the login/registration page.
func (ws *Server) Login(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.html", nil)
}
