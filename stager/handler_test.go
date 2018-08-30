package st8ger_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"

	"code.cloudfoundry.org/eirini/eirinifakes"
	"code.cloudfoundry.org/eirini/opi"
	stager "code.cloudfoundry.org/eirini/st8ger"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"
	"github.com/julienschmidt/httprouter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {

	var (
		ts     *httptest.Server
		logger lager.Logger
		client *http.Client

		backend            *eirinifakes.FakeBackend
		stagingClient      *eirinifakes.FakeSt8ger
		responseRecorder   *httptest.ResponseRecorder
		stagingHandler     *stager.StagingHandler
		stagingRequestJSON string
	)

	BeforeEach(func() {
		logger = lagertest.NewTestLogger("test")
		client = &http.Client{}
		backend = new(eirinifakes.FakeBackend)
		stagingClient = new(eirinifakes.FakeSt8ger)
		stagingHandler = stager.NewStagingHandler(stagingClient, backend, logger)
		stagingRequestJSON = `{"app_id":"myapp", "lifecycle":"kube-backend"}`
	})

	Context("Router", func() {

		JustBeforeEach(func() {
			ts = httptest.NewServer(stager.New(stagingClient, backend, logger))
		})

		Context("When it receives a staging request", func() {

			It("returns an Accepted response", func() {
				req, err := http.NewRequest("PUT", ts.URL+"/stage/myguid", bytes.NewReader([]byte(stagingRequestJSON)))
				Expect(err).ToNot(HaveOccurred())
				res, err := client.Do(req)
				Expect(err).ToNot(HaveOccurred())
				Expect(res.StatusCode).To(Equal(http.StatusAccepted))
			})

		})

		Context("When a wrong request is made", func() {
			It("returns a BadRequest response", func() {
				req, err := http.NewRequest("PUT", ts.URL+"/stage/myguid", bytes.NewReader([]byte{}))
				Expect(err).ToNot(HaveOccurred())
				res, err := client.Do(req)
				Expect(err).ToNot(HaveOccurred())
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
			})
		})
		It("Serves the DELETE /stage/:staging_guid endpoint", func() {
			req, err := http.NewRequest("DELETE", ts.URL+"/stage/myguid", nil)
			Expect(err).ToNot(HaveOccurred())
			res, err := client.Do(req)
			Expect(err).ToNot(HaveOccurred())
			Expect(res.StatusCode).To(Equal(200))
		})

		It("Serves the POST /stage/commpleted endpoint", func() {
			_, err := http.Post(ts.URL+"/stage/myguid/completed", "nothing", nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("Stage", func() {

		var (
			req *http.Request
			err error
		)

		BeforeEach(func() {
			req, err = http.NewRequest("PUT", "/stage/myguid", bytes.NewReader([]byte(stagingRequestJSON)))
			Expect(err).ToNot(HaveOccurred())
		})

		Context("BuildReceipe", func() {

			Context("When the staging task was created successfully", func() {
				JustBeforeEach(func() {
					responseRecorder = httptest.NewRecorder()
					param := httprouter.Param{Key: "staging_guid", Value: "myguid"}
					stagingHandler.Stage(responseRecorder, req, httprouter.Params{param})
				})

				It("is called with the right guid and staging request", func() {
					guid, req := backend.CreateStagingTaskArgsForCall(0)
					Expect(guid).To(Equal("myguid"))
					Expect(req.AppId).To(Equal("myapp"))
				})

			})

			Context("When the staging task creation fails", func() {
				JustBeforeEach(func() {
					responseRecorder = httptest.NewRecorder()
					param := httprouter.Param{Key: "staging_guid", Value: "myguid"}
					backend.CreateStagingTaskReturns(opi.Task{}, errors.New("aargh"))
					stagingHandler.Stage(responseRecorder, req, httprouter.Params{param})
				})

				It("returns an InternalServerError response when it errors", func() {
					Expect(responseRecorder.Code).To(Equal(http.StatusInternalServerError))
				})
			})
		})

		//*****END-BUILD-RECEIPE******

		Context("Stager", func() {
			Context("When the staging runs successfully", func() {

				JustBeforeEach(func() {
					responseRecorder = httptest.NewRecorder()
					param := httprouter.Param{Key: "staging_guid", Value: "myguid"}

					backend.CreateStagingTaskReturns(opi.Task{
						Image: "staging-image",
						Env: map[string]string{
							"UPLOAD_URL":   "http://upload-it.com",
							"DOWNLOAD_URL": "http://download-it.com",
							"STAGING_GUID": "staging-guid",
							"APP_ID":       "app-id",
						},
					}, nil)
					stagingHandler.Stage(responseRecorder, req, httprouter.Params{param})
				})

				It("is called with the staging task created by the backend", func() {
					task := stagingClient.RunArgsForCall(0)
					Expect(task.Env["DOWNLOAD_URL"]).To(Equal("http://download-it.com"))
					Expect(task.Env["UPLOAD_URL"]).To(Equal("http://upload-it.com"))
					Expect(task.Env["STAGING_GUID"]).To(Equal("staging-guid"))
					Expect(task.Env["APP_ID"]).To(Equal("app-id"))
				})
			})

			Context("When staging fails", func() {
				JustBeforeEach(func() {
					responseRecorder = httptest.NewRecorder()
					param := httprouter.Param{Key: "staging_guid", Value: "myguid"}

					backend.CreateStagingTaskReturns(opi.Task{
						Image: "staging-image",
						Env:   map[string]string{}}, nil)

					stagingClient.RunReturns(errors.New("$%}@#! I failed"))
					stagingHandler.Stage(responseRecorder, req, httprouter.Params{param})
				})

				It("returns a InternalServerError code", func() {
					Expect(responseRecorder.Code).To(Equal(http.StatusInternalServerError))
				})
			})
		})

	})
})
