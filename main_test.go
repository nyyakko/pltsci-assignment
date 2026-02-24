package main

import (
	"assignment/api/hoover"
	"assignment/middleware"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cucumber/godog"
)

type State struct {
	response *httptest.ResponseRecorder
}

func (self *State) WhenISendRequestWithBodyTo(method string, endpoint string, body string) error {
	req, err := http.NewRequest(method, endpoint, strings.NewReader(body))
	if err != nil {
		return err
	}

	middleware.ErrorHandlerMiddleware(hoover.Controller.CleaningSessions).ServeHTTP(self.response, req)

	return nil
}

func (self *State) ThenTheResponseCodeShouldBe(expectedCode int) error {
	if self.response.Code != expectedCode {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", expectedCode, self.response.Code)
	}
	return nil
}

func (self *State) AndTheResponseBodyShouldBe(expectedResponse string) error {
	response := self.response.Body.String()

	if strings.TrimSpace(response) != strings.TrimSpace(expectedResponse) {
		return fmt.Errorf("expected response body to be: %s, but actual is: %s", expectedResponse, response)
	}

	return nil
}

func TestMain(t *testing.T) {
	state := State{}

	suite := godog.TestSuite {
		ScenarioInitializer: func (ctx *godog.ScenarioContext) {
			ctx.Before(func (ctx context.Context, scenario *godog.Scenario) (context.Context, error) {
				state.response = httptest.NewRecorder()
				return ctx, nil
			})

			ctx.Step("^I send \"(GET|POST|PUT|DELETE)\" request with body to \"([^\"]*)\":$", state.WhenISendRequestWithBodyTo)
			ctx.Step("^the response code should be (\\d+)$", state.ThenTheResponseCodeShouldBe)
			ctx.Step("^the response body should be:$", state.AndTheResponseBodyShouldBe)
		},
		Options: &godog.Options {
			Format: "pretty",
			Paths: []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("failed to run tests")
	}
}
