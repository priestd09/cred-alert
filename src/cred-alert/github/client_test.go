package github_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/ghttp"
	"github.com/pivotal-golang/lager/lagertest"

	"cred-alert/github"
	"cred-alert/metrics"
	"cred-alert/metrics/metricsfakes"
)

var _ = Describe("Client", func() {
	var (
		client              github.Client
		server              *ghttp.Server
		fakeEmitter         *metricsfakes.FakeEmitter
		remainingCallsGauge *metricsfakes.FakeGauge
		logger              *lagertest.TestLogger
		header              http.Header
	)

	BeforeEach(func() {
		server = ghttp.NewServer()
		header = http.Header{
			"X-RateLimit-Limit":     []string{"60"},
			"X-RateLimit-Remaining": []string{"43"},
			"X-RateLimit-Reset":     []string{"1467645800"},
		}
		fakeEmitter = new(metricsfakes.FakeEmitter)
		httpClient := &http.Client{}

		logger = lagertest.NewTestLogger("client")

		remainingCallsGauge = new(metricsfakes.FakeGauge)
		fakeEmitter.GaugeStub = func(name string) metrics.Gauge {
			return remainingCallsGauge
		}
		client = github.NewClient(server.URL(), httpClient, fakeEmitter)
	})

	AfterEach(func() {
		if server != nil {
			server.Close()
		}
	})

	It("sets vnd.github.diff as the accept content-type header, and recieves a diff", func() {
		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/repos/owner/repo/compare/a...b"),
				ghttp.VerifyHeader(http.Header{
					"Accept": []string{"application/vnd.github.diff"},
				}),
				ghttp.RespondWith(http.StatusOK, `THIS IS THE DIFF`, header),
			),
		)

		diff, err := client.CompareRefs(logger, "owner", "repo", "a", "b")
		Expect(err).NotTo(HaveOccurred())
		Expect(diff).To(Equal("THIS IS THE DIFF"))
	})

	It("returns an error if the API returns an error", func() {
		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/repos/owner/repo/compare/a...b"),
				ghttp.VerifyHeader(http.Header{
					"Accept": []string{"application/vnd.github.diff"},
				}),
				ghttp.RespondWith(http.StatusInternalServerError, "", header),
			),
		)

		_, err := client.CompareRefs(logger, "owner", "repo", "a", "b")
		Expect(err).To(HaveOccurred())
	})

	It("returns an error if the API does not respond", func() {
		server.Close()
		server = nil

		_, err := client.CompareRefs(logger, "owner", "repo", "a", "b")
		Expect(err).To(HaveOccurred())
	})

	It("logs remaining api requests", func() {
		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/repos/owner/repo/compare/a...b"),
				ghttp.VerifyHeader(http.Header{
					"Accept": []string{"application/vnd.github.diff"},
				}),
				ghttp.RespondWith(http.StatusOK, "", header),
			),
		)
		_, err := client.CompareRefs(logger, "owner", "repo", "a", "b")
		Expect(err).ToNot(HaveOccurred())
		Expect(remainingCallsGauge.UpdateCallCount()).To(Equal(1))
		_, value := remainingCallsGauge.UpdateArgsForCall(0)
		Expect(value).To(Equal(float32(43)))
	})

})
