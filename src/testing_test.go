package src

import (
	"reflect"
	"testing"
)

type T struct {
	A int
	B string
}

func doSomeStuff(x int) (T, error) {
	return T{
		A: x,
		B: "A",
	}, nil
}

var testCases = []struct {
	name    string
	arg     int
	want    T
	wantErr bool
}{
	{},
}

// BEGIN testing1 OMIT
func TestTesting(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	prepare(t)

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := doSomeStuff(tt.arg)

			if tt.wantErr != (err != nil) {
				t.Fatalf("failed: %v", err)
			}
			if !reflect.DeepEqual(got.A, tt.want.A) {
				t.Errorf("expected = %q, want %q", got, tt.want)
			}
			validateB(t, tt.want.B)
		})
	}
}

// END testing1 OMIT

// BEGIN testing2 OMIT
func validateB(t *testing.T, s string) {
	t.Helper()

	// do some checks
}

func prepare(t *testing.T) {
	// do some setup

	t.Cleanup(func() {
		// do some teardown
	})
}

// END testing2 OMIT
