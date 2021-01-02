// +build e2e

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/architectv/merchantx-task/pkg/handler"
	"github.com/architectv/merchantx-task/pkg/repository"
	"github.com/architectv/merchantx-task/pkg/service"
	"github.com/architectv/merchantx-task/test"
	"github.com/gofiber/fiber/v2"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_E2E_App(t *testing.T) {
	const prefix = "./"
	db, err := test.PrepareTestDatabase(prefix)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()
	// defer test.ClearTestDatabase(db, prefix)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	app := fiber.New()
	handlers.InitRoutes(app)

	Convey("Given params", t, func() {
		const (
			expectedStatus      = fiber.StatusOK
			expectedCreateCount = 5
			expectedUpdateCount = 1
			expectedDeleteCount = 1
			expectedErrorCount  = 5
		)

		expectedBody := fmt.Sprintf(`{"create_count":%d,"update_count":%d,"delete_count":%d,"error_count":%d}`,
			expectedCreateCount, expectedUpdateCount, expectedDeleteCount, expectedErrorCount)
		const (
			inputId   = 1
			inputLink = test.OkLink
		)
		inputBody := fmt.Sprintf(`{"id": %d, "link": "%s"}`, inputId, inputLink)
		Convey("When put method", func() {
			req := httptest.NewRequest(
				"PUT",
				"/api/offers/",
				bytes.NewBufferString(inputBody),
			)
			req.Header.Set("Content-type", "application/json")

			resp, err := app.Test(req, -1)
			body, bodyErr := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))
			Convey("Then should be Ok", func() {
				So(err, ShouldBeNil)
				So(bodyErr, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, expectedStatus)
				So(string(body), ShouldEqual, expectedBody)
			})
		})
	})
}
