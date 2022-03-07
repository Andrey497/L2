package server

import (
	"log"
	"net/http"
	"rest/pkg/handler"
)

type Server struct {
	handler *handler.Handler
	router  *http.ServeMux
}

func InitServer(h *handler.Handler, rt *http.ServeMux) *Server {
	return &Server{handler: h, router: rt}
}

func (s *Server) StartServer() error {
	configuredRouter := s.configureRouter()
	return http.ListenAndServe(":"+port, configuredRouter)
}

func (s *Server) configureRouter() http.Handler {
	s.router.HandleFunc("/events_for_day", s.middleware(s.handler.GetEventForDay))
	s.router.HandleFunc("/events_for_week", s.middleware(s.handler.GetEventForWeek))
	s.router.HandleFunc("/events_for_month", s.middleware(s.handler.GetEventForMonth))
	s.router.HandleFunc("/create_event", s.middleware(s.handler.CreateEvent))
	s.router.HandleFunc("/update_event", s.middleware(s.handler.UpdateEvent))
	s.router.HandleFunc("/delete_event", s.middleware(s.handler.DeleteEvent))
	return s.router
}
func (s *Server) middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Print(request.URL)
		next(writer, request)
	}

}
