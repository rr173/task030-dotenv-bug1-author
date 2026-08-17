package httpapi

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProbeEmptyVariablesEncodedAsArray(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/parse", bytes.NewBufferString(`{"content":""}`))
	rec := httptest.NewRecorder()
	New().Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	body, err := io.ReadAll(rec.Result().Body)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(body, []byte(`"variables":[]`)) {
		t.Fatalf("empty variables must be encoded as an array, got %s", body)
	}
}
