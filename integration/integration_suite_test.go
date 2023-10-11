package integration_test

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"code.cloudfoundry.org/cpu-entitlement-plugin/httpclient"
	. "code.cloudfoundry.org/cpu-entitlement-plugin/test_utils"
	logcache "code.cloudfoundry.org/go-log-cache"
	"code.cloudfoundry.org/lager/v3"
	"code.cloudfoundry.org/lager/v3/lagertest"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	SetDefaultEventuallyTimeout(120 * time.Second)
	SetDefaultEventuallyPollingInterval(5 * time.Second)
	RunSpecs(t, "Integration Suite")
}

var (
	cfApi                string
	logEmitterHttpClient *http.Client
	logCacheClient       *logcache.Client
	getToken             func() (string, error)
	logger               lager.Logger
	config               Config
)

type Config struct {
	AdminPassword string `json:"admin_password"`
	AdminUsername string `json:"admin_username"`
	Api           string `json:"api"`
}

var _ = SynchronizedBeforeSuite(func() []byte {
	configPath := GetEnv("CONFIG")
	configFile, err := os.Open(configPath)
	Expect(err).NotTo(HaveOccurred())

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	Expect(err).NotTo(HaveOccurred())

	cfApi = config.Api
	cfUsername := config.AdminUsername
	cfPassword := config.AdminPassword

	Expect(Cmd("cf", "api", cfApi, "--skip-ssl-validation").Run()).To(gexec.Exit(0))
	Expect(Cmd("cf", "login", "-u", cfUsername, "-p", cfPassword).Run()).To(gexec.Exit(0))

	logEmitterHttpClient = createInsecureHttpClient()
	Eventually(pingTestLogEmitter).Should(BeTrue())

	logCacheURL := getLogCacheURL()
	getToken = func() (string, error) {
		return getCmdOutput("cf", "oauth-token"), nil
	}
	logCacheClient = logcache.NewClient(
		logCacheURL,
		logcache.WithHTTPClient(httpclient.NewAuthClient(getToken)),
	)
	logger = lagertest.NewTestLogger("cumulative-usage-fetcher-test")

	return []byte{}
}, func(_ []byte) {})

func createInsecureHttpClient() *http.Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig.InsecureSkipVerify = true
	return &http.Client{Transport: transport}
}

func GetEnv(varName string) string {
	value := os.Getenv(varName)
	ExpectWithOffset(1, value).NotTo(BeEmpty())
	return value
}

func pingTestLogEmitter() bool {
	response, err := logEmitterHttpClient.Get(getTestLogEmitterURL())
	if err != nil {
		return false
	}
	defer response.Body.Close()
	return response.StatusCode == http.StatusOK
}

func getTestLogEmitterURL() string {
	return strings.Replace(cfApi, "api.", "test-log-emitter.", 1)
}

func getLogCacheURL() string {
	return strings.Replace(cfApi, "https://api.", "http://log-cache.", 1)
}
