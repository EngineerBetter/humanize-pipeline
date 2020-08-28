package main_test

import (
	"io/ioutil"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("humanize-pipeline", func() {
	var session *gexec.Session
	var binaryPath, inputFilePath, expected string

	BeforeSuite(func() {
		binaryPath = buildBinary()
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	BeforeEach(func() {
		inputFilePath = "fixtures/simple-pipeline-ordered.yml"
		expected = getExpected("fixtures/simple-pipeline-ordered.yml")
	})

	JustBeforeEach(func() {
		session = runBinary(binaryPath, inputFilePath)
	})

	Context("When run with a valid, ordered pipeline", func() {
		It("exits with status code 0", func() {
			Eventually(session).Should(gexec.Exit(0))
		})

		It("doesn't change the input yaml", func() {
			Eventually(session).Should(gbytes.Say(expected))
		})
	})
})

func buildBinary() string {
	binaryPath, err := gexec.Build("github.com/EngineerBetter/humanize-pipeline")
	Expect(err).NotTo(HaveOccurred())

	return binaryPath
}

func runBinary(path, inputFilePath string) *gexec.Session {
	cmd := exec.Command(path, inputFilePath)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session
}

func getExpected(path string) string {
	expectedBytes, err := ioutil.ReadFile(path)
	Expect(err).ToNot(HaveOccurred())
	return string(expectedBytes)
}
