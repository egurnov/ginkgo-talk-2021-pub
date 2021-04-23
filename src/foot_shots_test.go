package src_test

import (
	"log"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FootShots", func() {

	Specify("Type conversions", func() {
		// BEGIN TypeConv OMIT
		var f float64 = 5
		Expect(f).ToNot(Equal(5)) // HL

		Expect(5.1).To(BeEquivalentTo(5)) // HL
		Expect(5).ToNot(BeEquivalentTo(5.1))

		Expect(f).To(BeNumerically("==", 5))
		// END TypeConv OMIT
	})

	Specify("ContainElement", func() {
		// BEGIN ContainElement OMIT
		s := []byte("abracadabra")
		Expect(s).ToNot(ContainElement("cad")) // HL

		Expect(s).To(ContainSubstring("cad"))
		// END ContainElement OMIT
	})

	Specify("Eventually", func() {
		// BEGIN Eventually OMIT
		start := time.Now()
		isReady := func() bool {
			return time.Since(start) > 800*time.Millisecond
		}

		Eventually(isReady()).ShouldNot(BeTrue()) // HL

		Eventually(isReady).Should(BeTrue())
		// END Eventually OMIT
	})

	Context("Clojures bad", func() {
		// BEGIN closures1 OMIT
		var v = 5

		It("uses v", func() {
			Expect(v).To(Equal(5))
		})

		It("uses v too", func() {
			v = 6 // HL
			Expect(v).To(Equal(6))
		})
		// END closures1 OMIT
	})

	Context("Clojures good", func() {
		// BEGIN closures2 OMIT
		var v int

		BeforeEach(func() { // HL
			v = 5 // HL
		}) // HL

		It("uses v", func() {
			Expect(v).To(Equal(5))
		})

		It("uses v too", func() {
			v = 6
			Expect(v).To(Equal(6))
		})
		// END closures2 OMIT
	})

	Context("Generating tests", func() {
		// BEGIN Loop1 OMIT
		for i := 0; i < 5; i++ {
			Specify("test #"+strconv.Itoa(i), func() {
				log.Println("Running test #", i)
			})
		}
		// END Loop1 OMIT

		var _ = ` // BEGIN Loop1-res OMIT
2021/04/01 15:23:11 Running test # 5
2021/04/01 15:23:11 Running test # 5
2021/04/01 15:23:11 Running test # 5
2021/04/01 15:23:11 Running test # 5
2021/04/01 15:23:11 Running test # 5
` // END Loop1-res OMIT

		// BEGIN Loop2 OMIT
		for i := 0; i < 5; i++ {
			i := i // create a local copy of the loop variable // HL
			Specify("test #"+strconv.Itoa(i), func() {
				log.Println("Running test #", i)
			})
		}
		// END Loop2 OMIT

		var _ = ` // BEGIN Loop2-res OMIT
2021/04/01 15:23:11 Running test # 0
2021/04/01 15:23:11 Running test # 1
2021/04/01 15:23:11 Running test # 2
2021/04/01 15:23:11 Running test # 3
2021/04/01 15:23:11 Running test # 4
` // END Loop2-res OMIT

	})

	Context("Async functions", func() {
		XContext("Recover", func() {
			// BEGIN recover1 OMIT
			It("panics in a goroutine", func(done Done) {
				go func() {
					Fail("Oh noes!")
					close(done)
				}()
			})
			// END recover1 OMIT

			var _ = ` // BEGIN recover1-res OMIT
panic: 
Your test failed.
Ginkgo panics to prevent subsequent assertions from running.
Normally Ginkgo rescues this panic so you shouldn't see it.

But, if you make an assertion in a goroutine, Ginkgo can't capture the panic.
To circumvent this, you should call

        defer GinkgoRecover()

at the top of the goroutine that caused this panic.
` // END recover1-res OMIT

			// BEGIN recover2 OMIT
			It("fails in a goroutine", func(done Done) {
				go func() {
					defer GinkgoRecover() // HL
					Fail("Oh noes!")
					close(done)
				}()
			})
			// END recover2 OMIT

			var _ = ` // BEGIN recover2-res OMIT
[Fail] FootShots Async functions Recover [It] fails in a goroutine too
` // END recover2-res OMIT

		})

		XContext("in a bad case", func() {
			// BEGIN async1 OMIT
			It("fails in a goroutine", func() {
				go func() {
					defer GinkgoRecover()
					time.Sleep(500 * time.Millisecond)
					Fail("Oh noes!")
				}()
			})

			It("doesn't do anything bad", func() {
				time.Sleep(800 * time.Millisecond)
			})
			// END async1 OMIT

			var _ = ` // BEGIN async1-res OMIT
[Fail] FootShots Async functions in a bad case [It] doesn't do anything bad
` // END async1-res OMIT
		})

		XWhen("done properly", func() {
			// BEGIN async2 OMIT
			It("fails in a goroutine", func(done Done) { // HL
				go func() {
					defer GinkgoRecover()
					time.Sleep(500 * time.Millisecond)
					Fail("Oh noes!")
					close(done) // HL
				}()
			})

			It("doesn't do anything bad", func() {
				time.Sleep(800 * time.Millisecond)
			})
			// END async2 OMIT

			var _ = ` // BEGIN async2-res OMIT
[Fail] FootShots Async functions when done properly [It] fails in a goroutine
` // END async2-res OMIT
		})
	})

})
