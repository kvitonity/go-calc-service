package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExpressionHandler(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
		wantBody   string
	}{
		{"valid addition", "op1=2&operator=%2B&op2=4", http.StatusOK, `{"result":6}`},
		{"invalid op1", "op1=abc&operator=%2B&op2=4", http.StatusBadRequest, ""},
		{"division by zero", "op1=10&operator=%2F&op2=0", http.StatusBadRequest, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/expression?"+tt.query, nil)
			w := httptest.NewRecorder()

			ExpressionHandler(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", w.Code, tt.wantStatus)
			}

			if tt.wantBody != "" {
				body := w.Body.String()
				if body != tt.wantBody+"\n" {
					t.Errorf("body = %q, want %q", body, tt.wantBody+"\n")
				}
			}
		})
	}
}
