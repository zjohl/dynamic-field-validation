package main_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {

	It("validates the reponse of the provided url", func() {

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			bytes, err := ioutil.ReadFile("fixtures/example_api_response.json")
			Expect(err).NotTo(HaveOccurred())
			w.Write(bytes)
		}))

		cmd := exec.Command(binaryPath, server.URL)
		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		expectedResponse, err := ioutil.ReadFile("fixtures/expected_response.json")
		Expect(err).NotTo(HaveOccurred())
		Eventually(session.Out.Contents()).Should(Equal(expectedResponse))

		gexec.CleanupBuildArtifacts()
	})
})
