package integration_test

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
)

const (
	// Attempts connection
	host       = "app:8080"
	healthPath = "http://" + host + "/healthz"
	attempts   = 20

	// HTTP REST
	basePath = "http://" + host + "/v1"
)

func TestMain(m *testing.M) {
	err := healthCheck(attempts)
	if err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", host, err)
	}

	log.Printf("Integration tests: host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck(attempts int) error {
	var err error

	for attempts > 0 {
		err = Do(Get(healthPath), Expect().Status().Equal(http.StatusOK))
		if err == nil {
			return nil
		}

		log.Printf("Integration tests: url %s is not available, attempts left: %d", healthPath, attempts)
		time.Sleep(time.Second)
		attempts--
	}

	return err
}

// HTTP GET: /v1/user/:userId/planning.
func TestHTTPGetOrderedPlanning(t *testing.T) {
	Test(t,
		Description("Get ordered planning by user id"),
		Get(basePath+"/user/"+"30ecc27b-9df7-4dd3-b52f-d001e79bd035"+"/planning"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().String().Contains(`{"course_planning":{`),
	)
}

// HTTP POST: /v1/user/course
func TestHTTPDoCourseOrder(t *testing.T) {
	body := `{
		"userId": "30ecc27b-9df7-4dd3-b52f-d001e79bd035",
		"courses": [
			{
				"desiredCourse": "PortfolioConstruction",
				"requiredCourse": "PortfolioTheories"
			},
			{
				"desiredCourse": "InvestmentManagement",
				"requiredCourse": "Investment"
			},
			{
				"desiredCourse": "Investment",
				"requiredCourse": "Finance"
			},
			{
				"desiredCourse": "PortfolioTheories",
				"requiredCourse": "Investment"
			},
			{
				"desiredCourse": "InvestmentStyle",
				"requiredCourse": "InvestmentManagement"
			}
		]
	}`
	Test(t,
		Description("DoCourseOrder Success"),
		Post(basePath+"/user/course"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
	)

	body = `{
	}`
	Test(t,
		Description("DoCourseOrder Fail"),
		Post(basePath+"/user/course"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".error").Equal("invalid request body"),
	)
}
