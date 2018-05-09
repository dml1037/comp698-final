import (
 	"net/http"
 	"net/http/httptest"
	"strings"
 	"testing"
 )
 
 func TestHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	request, _ := http.NewRequest("GET", "/home", nil)
 	response := httptest.NewRecorder()
	hello(response, request)
	IndexHandler(response, request)
 
	expected := "Hello, World!"
	if response.Body.String() != expected {
	expected := "COMP698 Final Project with Bootstrap"
	if !strings.Contains(response.Body.String(), expected) {
 		t.Fatalf("Got: %s\nExpected: %s",
 			response.Body.String(), expected)
 	}