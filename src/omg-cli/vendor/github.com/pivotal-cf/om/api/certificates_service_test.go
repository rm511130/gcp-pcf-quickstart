package api_test

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/api/fakes"
)

var _ = Describe("CertificatesService", func() {
	var (
		client  *fakes.HttpClient
		service api.CertificatesService
	)

	BeforeEach(func() {
		client = &fakes.HttpClient{}
		service = api.NewCertificatesService(client)
	})

	Describe("Generate", func() {
		It("returns a cert and key", func() {
			var path string
			var header http.Header
			var body io.Reader

			requestBody := `{
"certificate": "some-certificate",
"key": "some-key"
}`
			client.DoStub = func(req *http.Request) (*http.Response, error) {
				path = req.URL.Path
				body = req.Body
				header = req.Header

				var resp *http.Response
				if path == "/api/v0/certificates/generate" && req.Method == "POST" {
					return &http.Response{StatusCode: http.StatusOK,
						Body: ioutil.NopCloser(strings.NewReader(requestBody)),
					}, nil
				}
				return resp, nil
			}

			output, err := service.Generate("*.example.com,*.example.org")
			Expect(err).NotTo(HaveOccurred())

			Expect(header.Get("Content-Type")).To(Equal("application/json"))

			bodyBytes, err := ioutil.ReadAll(body)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(bodyBytes)).To(MatchJSON(`{
				"domains": [
				"*.example.com",
				"*.example.org"
				]
			}`))

			Expect(output).To(Equal(requestBody))

			Expect(path).To(Equal("/api/v0/certificates/generate"))
		})

		Context("failure cases", func() {
			Context("when the client cannot make the request", func() {
				It("returns an error", func() {
					client.DoReturns(nil, errors.New("client do errored"))

					_, err := service.Generate("some-domains")
					Expect(err).To(MatchError("client do errored"))
				})
			})

			Context("when Ops Manager returns a non-200 status code", func() {
				BeforeEach(func() {
					client = &fakes.HttpClient{}
					client.DoStub = func(req *http.Request) (*http.Response, error) {
						var resp *http.Response
						if req.URL.Path == "/api/v0/certificates/generate" &&
							req.Method == "POST" {
							return &http.Response{
								StatusCode: http.StatusInternalServerError,
								Body:       ioutil.NopCloser(bytes.NewBufferString(`{}`)),
							}, nil
						}
						return resp, nil
					}
				})

				It("returns an error", func() {
					service := api.NewCertificatesService(client)
					_, err := service.Generate("some-domains")
					Expect(err).To(MatchError(ContainSubstring("request failed: unexpected response")))
				})
			})
		})
	})
})
