package http

import (
	"net/http"
)

// Redis test redis connection
func (s *Server) Redis(w http.ResponseWriter, r *http.Request) {
	var err error
	testResult := ""

	addr := r.URL.Query().Get("addr")
	if addr != "" {
		testResult, err = s.xtest.TestRedisNewAddr(addr)
	} else {
		testResult, err = s.xtest.TestRedisDefaultAddr()
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(testResult))
}
