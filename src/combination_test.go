package src

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

var _ = Describe("Combining matchers", func() {

	It("can create custom matchers", func() {
		// BEGIN transform OMIT
		type T struct {
			name, id string
		}

		getName := func(t T) string {
			return t.name
		}

		withName := func(name string) types.GomegaMatcher {
			return WithTransform(getName, Equal(name))
		}

		arr := []T{{"a", "1"}, {"b", "2"}, {"c", "3"}}
		Expect(arr).To(ContainElement(withName("a")))
		// END transform OMIT
	})

	It("can combine matchers", func() {
		// BEGIN combine OMIT
		Expect(5).To(
			And(
				BeNumerically(">", 4),
				BeNumerically("<", 6),
			),
		)

		BeInRange := func(from, to interface{}) types.GomegaMatcher {
			return SatisfyAll(
				BeNumerically(">", from),
				BeNumerically("<", to),
			)
		}
		Expect(5).To(BeInRange(3, 6))

		Expect([]int{4, 8, 15, 16, 23, 42}).To(
			ContainElements(
				BeInRange(10, 20),
				BeInRange(40, 50),
			),
		)
		// END combine OMIT
	})

})
