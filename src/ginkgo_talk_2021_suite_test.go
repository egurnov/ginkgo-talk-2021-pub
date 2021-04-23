package src_test

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// BEGIN suite OMIT
func TestGinkgoTalk2021(t *testing.T) {
	log.SetOutput(GinkgoWriter) // HLgw

	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoTalk2021 Suite") // HLrun
}

// END suite OMIT

var _ = BeforeSuite(func() {
	log.Println("Before Suite")
})

var _ = AfterSuite(func() {
	log.Println("After Suite")
})
