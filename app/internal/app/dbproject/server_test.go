package dbproject

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ImbaCow/bd_project/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func Test_server_handleUserCreate(t *testing.T) {
	s := newServer(teststore.New(), nil)
	tests := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"login":    "login",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "no password",
			payload: map[string]string{
				"login": "login",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "small password",
			payload: map[string]string{
				"login":    "login",
				"password": "pas",
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/user/add", b)
			s.router.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
