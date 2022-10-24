package restserver

import "github.com/go-chi/chi/v5"

func (s Server) SetupEmailRoutes() {
	s.ServerChi.Mount("/enron", getEmailUsers(s))
}

func getEmailUsers(s Server) chi.Router {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Get("/", s.emailControl.GetEmailUsers)
		r.Get("/{user-name}", s.emailControl.GetEmailsByUser)
	})
	return r
}
