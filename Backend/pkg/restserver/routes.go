package restserver

import "github.com/go-chi/chi/v5"

func (s Server) SetupEmailRoutes() {
	// s.ServerChi.Mount("/enron", getEmailUsers(s))
	s.ServerChi.Mount("/search", getDocuments(s))
}

// func getEmailUsers(s Server) chi.Router {
// 	r := chi.NewRouter()
// 	r.Route("/users", func(r chi.Router) {
// 		r.Get("/", s.emailControl.GetEmailUsers)
// 		r.Get("/{user-name}", s.emailControl.BulkDocuments)
// 	})
// 	return r
// }

func getDocuments(s Server) chi.Router {
	r := chi.NewRouter()
	r.Get("/{temp}", s.emailControl.SearchDocumentsByfilter)
	return r
}
