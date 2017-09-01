package main_test

import (
	"io/ioutil"
	"os/exec"

	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {

	It("validates the reponse of the provided url", func() {

		cmd := exec.Command(binaryPath, "http://backend-challenge-winter-2017.herokuapp.com/customers.json")
		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		expectedResponse, err := ioutil.ReadFile("fixtures/expected_response_integration.json")
		Expect(err).NotTo(HaveOccurred())
		time.Sleep(time.Second)
		Eventually(string(session.Out.Contents())).Should(MatchJSON(expectedResponse))

		gexec.CleanupBuildArtifacts()
	})
})
