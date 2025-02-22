package stat

import (
	"net/http"
	"server/configs"
	"server/pkg/middleware"
	"server/pkg/resp"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatHandlerDep struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDep) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))
}

func (s *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid from param", http.StatusBadRequest)
			return
		}
		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid to param", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "Invalid by param", http.StatusBadRequest)
			return
		}
		stats := s.StatRepository.GetStats(by, from, to)

		resp.ResponceJson(w, 200, stats)
	}
}
