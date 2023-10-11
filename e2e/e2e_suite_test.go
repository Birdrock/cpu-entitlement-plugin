package e2e_test

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	SetDefaultEventuallyTimeout(120 * time.Second)
	SetDefaultEventuallyPollingInterval(5 * time.Second)
	RunSpecs(t, "E2e Suite")
}

var config Config

type Config struct {
	AdminPassword string `json:"admin_password"`
	AdminUsername string `json:"admin_username"`
	Api           string `json:"api"`
}

var _ = SynchronizedBeforeSuite(func() []byte {
	_ = GetEnv("CPU_ENTITLEMENT_PLUGIN_BINARY")
	_ = GetEnv("CPU_OVERENTITLEMENT_INSTANCES_PLUGIN_BINARY")

	configPath := GetEnv("CONFIG")
	configFile, err := os.Open(configPath)
	Expect(err).NotTo(HaveOccurred())

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	Expect(err).NotTo(HaveOccurred())
	return []byte{}
}, func(_ []byte) {})

func GetEnv(varName string) string {
	value := os.Getenv(varName)
	ExpectWithOffset(1, value).NotTo(BeEmpty())
	return value
}
