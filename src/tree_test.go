package src_test

import (
	"log"

	. "github.com/onsi/ginkgo"
)

// BEGIN tree OMIT
var _ = Describe("Test tree", func() {
	Context("Some case", func() {
		Specify("some general case", func() {
			// Test 1
		})
		When("something happens", func() {
			It("does this", func() {
				// Test 2
			})
			Specify("longer test case", func() {
				// Test 3
				By("Step 1")
				// ...
				By("Step 2")
				// ...
			})
		})
	})
})

// END tree OMIT

var _ = Describe("Ginkgo", func() {

	XDescribe("Focus/Skip/Pending", func() {
		someCondition := false
		// BEGIN focus OMIT
		FIt("Focus", func() {
			// ...
		})

		PIt("Pending", func() {
			Fail("never happens")
		})

		XIt("Also pending", func() {
			Fail("never happens")
		})

		It("Skip", func() {
			if someCondition {
				Skip("nope")
			}
		})
		// END focus OMIT
	})

	// BEGIN setup-tree OMIT
	Context("Context 1", func() {
		BeforeEach(func() {})
		JustBeforeEach(func() {})
		JustAfterEach(func() {})
		AfterEach(func() {})

		Specify("test 1", func() {})

		Context("Context 2", func() {
			BeforeEach(func() {})
			JustBeforeEach(func() {})
			JustAfterEach(func() {})
			AfterEach(func() {})

			Specify("test 2", func() {})
		})
	})
	// END setup-tree OMIT

	Context("Setup and Teardown", func() {
		BeforeEach(func() { log.Println("BeforeEach 1") })
		JustBeforeEach(func() { log.Println("JustBeforeEach 1") })
		JustAfterEach(func() { log.Println("JustAfterEach 1") })
		AfterEach(func() { log.Println("AfterEach 1") })

		Specify("outer context test", func() {
			log.Println("Test 1")
		})

		Context("Inner context", func() {
			BeforeEach(func() { log.Println("BeforeEach 2") })
			JustBeforeEach(func() { log.Println("JustBeforeEach 2") })
			JustAfterEach(func() { log.Println("JustAfterEach 2") })
			AfterEach(func() { log.Println("AfterEach 2") })

			It("works", func() {
				log.Println("Test 2")
			})
		})
	})

	var _ = ` // BEGIN setup-res OMIT
Before Suite
----------------
BeforeEach 1
JustBeforeEach 1
Test 1
JustAfterEach 1
AfterEach 1
----------------
BeforeEach 1
BeforeEach 2
JustBeforeEach 1
JustBeforeEach 2
Test 2
JustAfterEach 2
JustAfterEach 1
AfterEach 2
AfterEach 1
----------------
After Suite
` // END setup-res OMIT

})
