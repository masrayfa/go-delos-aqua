package middleware

import (
	"log"
	"net/http"
	"sync"
)

type StatsResponse struct {
	Count int `json:"count"`
	UniqueUserAgent int `json:"unique_user_agent"`
}

type EndpointStats struct {
	Count int
	UniqueUserAgent map[string]struct{}
}

type Stats struct {
	mu    sync.Mutex
	stats map[string]*EndpointStats
}

func NewStats() *Stats {
	return &Stats{
		stats: make(map[string]*EndpointStats),
	}
}

func (s *Stats) AddEndpointStats(endpoint string, userAgent string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if stat, ok := s.stats[endpoint]; ok {
		stat.Count++
		stat.UniqueUserAgent[userAgent] = struct{}{}
	} else {
		s.stats[endpoint] = &EndpointStats{
			Count: 1,
			UniqueUserAgent: map[string]struct{}{userAgent: {}},
		}
	}
}

func (s *Stats) GetEndpointStats() map[string]interface{} {
	s.mu.Lock()

	defer s.mu.Unlock()

	results := make(map[string]interface{})
	for endpoint, stat := range s.stats {
		results[endpoint] = &StatsResponse{
			Count: stat.Count,
			UniqueUserAgent: len(stat.UniqueUserAgent),
		}
	}

	return results
}


func (s *Stats) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("@Middleware.Stats")

		originalPath := r.Header.Get("X-Original-Path")
		log.Println("@Middleware.Stats:originalPath ->", originalPath)

		userAgent := r.Header.Get("User-Agent")
		s.AddEndpointStats(r.Method + " " + originalPath, userAgent)

		next.ServeHTTP(w, r)
	})
}