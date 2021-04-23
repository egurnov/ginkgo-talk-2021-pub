package src_test

import (
	"testing"

	. "github.com/onsi/gomega"
)

// go test -count=1 ./gomega

// BEGIN gomega OMIT
func TestGomegaStandalone(t *testing.T) {
	g := NewWithT(t)
	g.Expect(5).To(Equal(5))
}

// END gomega OMIT
