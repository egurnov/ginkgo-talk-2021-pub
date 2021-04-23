package src_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Matchers", func() {
	var (
		err    error
		p1, p2 *int
	)
	d1 := time.Date(2020, time.October, 29, 19, 30, 0, 0, time.UTC)
	d2 := time.Date(2020, time.October, 29, 19, 34, 0, 0, time.UTC)
	f := func() { panic("Nooooooooo!") }
	someTask := func() error { return nil }
	ch := make(chan int, 1)
	var v int

	someCheck := func() error { return nil }
	someFunc := func() int { return 5 }

	Specify("Matchers", func() {
		// BEGIN matchers1 OMIT
		Expect(5).To(Equal(5))
		Expect(5.0).To(BeEquivalentTo(5))
		Expect(p1).To(BeIdenticalTo(p2))

		Expect(err).ToNot(HaveOccurred())
		Expect(someTask()).To(Succeed())
		Expect(f).To(Panic())

		Expect(5).To(BeNumerically("<", 5.1))
		Expect(d1).To(BeTemporally("~", d2, 5*time.Minute))

		Expect("Abracadabra").To(ContainSubstring("cad"))
		Expect("x-y=z").To(ContainSubstring("%v-%v", "x", "y"))
		Expect("me@example.com").To(MatchRegexp("[a-z]+@[a-z]+\\.[a-z]{2,}"))
		Expect("{\"a\": 1, \"b\": 2}").To(MatchJSON("{\"b\": 2, \"a\": 1}"))
		// END matchers1 OMIT

		// BEGIN matchers2 OMIT
		Expect(ch).To(BeSent(7))
		Expect(ch).To(Receive(&v))
		Expect(ch).ToNot(BeClosed())

		theSequence := []int{4, 8, 15, 16, 23, 42}
		Expect(theSequence).ToNot(BeEmpty())
		Expect(theSequence).To(HaveLen(6))
		Expect(theSequence).To(ContainElement(23))
		Expect(15).To(BeElementOf(theSequence))
		Expect(theSequence).To(ConsistOf(8, 16, 42, 23, 15, 4))

		shoppingList := map[string]int{"apples": 4, "tomatoes": 10, "milk": 1}
		Expect(shoppingList).To(HaveKey("apples"))
		Expect(shoppingList).To(HaveKeyWithValue("tomatoes", 10))
		Expect(shoppingList).To(ConsistOf(1, 4, 10))
		// END matchers2 OMIT

		// BEGIN eventually OMIT
		Eventually(someCheck).Should(Succeed())
		Consistently(someFunc).Should(BeNumerically(">", 3))
		// END eventually OMIT
	})
})
