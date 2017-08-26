package app

import "gitlab.com/mhyusufibrahim/teahrm/public"

// Initialize routes
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
	a.Router.HandleFunc("/users", a.createUser).Methods("POST")
	a.Router.HandleFunc("/users/{id:[0-9]+}", a.getUser).Methods("GET")
	a.Router.HandleFunc("/users/{id:[0-9]+}", a.updateUser).Methods("PUT")
	a.Router.HandleFunc("/users/{id:[0-9]+}", a.deleteUser).Methods("DELETE")

	a.Router.HandleFunc("/token", a.getApiToken).Methods("POST")
	a.Router.HandleFunc("/token", a.changeApiToken).Methods("PUT")

	a.Router.HandleFunc("/", public.HomeHandler)
	a.Router.HandleFunc("/login", public.LoginHandler)
	a.Router.HandleFunc("/register", public.RegisterHandler)
}
