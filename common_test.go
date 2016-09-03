package main

import (
  "net/http"
  "net/http/httptest"
  "os"
  "testing"
  "github.com/gin-gonic/gin"
)

var tmpUserList []user
var tmpArticleList []article

// This function is used to do setup before executing the test functions
func TestMain(m *testing.M) {
  //Set Gin to Test Mode
  gin.SetMode(gin.TestMode)

  // Run the other tests
  os.Exit(m.Run())
}

// Helper function to create a router during testing
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
  // Create a response recorder
  w := httptest.NewRecord()

  // Create the service and process the above request.
  r.ServeHTTP(w, req)

  if !f(w) {
    t.Fail()
  }
}

// This is a helper function that allows us to reuse some code in the above
// test methods

func testMiddlewareRequest(t *testing.T, r *gin.Engine, expectedHTTPcode int) {
  // Create a request to send to the above route
  req, _ := http.newRequest("GET", "/", nil)

  // Process the request and test the response
  testHTTPResponse(t, r, req, fun(w *httptest.ResponseRecorder) bool {
  return w.Code == expectedHTTPcode
})
}

// This function is used to store the main lists into the temporary one
// for testing
func saveList() {
  tmpUserList = userList
  tmpArticleList = articleList
}

// This function is used to restore the main lists from the temporary one
func restoreList() {
  userList = tmpUserList
  article = tmpArticleList
}
