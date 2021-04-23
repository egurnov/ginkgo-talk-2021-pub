package src

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ = ` // BEGIN testify-out OMIT
=== RUN   TestTestify
    testify_test.go:10: 
        	Error Trace:	testify_test.go:10
        	Error:      	Not equal: 
        	            	expected: "actual"         // HL
        	            	actual  : "expected"       // HL
        	Test:       	TestTestify
--- FAIL: TestTestify (0.00s)
` // END testify-out OMIT

// BEGIN testify1 OMIT
func TestTestify(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	prepare(t)

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := doSomeStuff(tt.arg)

			require.NoError(t, err)
			assert.Equal(t, got.A, tt.want.A)

			validateB(t, tt.want.B)
		})
	}
}

// END testify1 OMIT

func TestExtensions(t *testing.T) {
	{
		isOkay := func(x int) bool {
			return x >= 42
		}

		tests := []struct {
			name      string
			arg       int
			assertion assert.BoolAssertionFunc
		}{
			{"-1 is bad", -1, assert.False},
			{"42 is good", 42, assert.True},
			{"41 is bad", 41, assert.False},
			{"45 is cool", 45, assert.True},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tt.assertion(t, isOkay(tt.arg))
			})
		}

	}

	{
		adder := func(x, y int) int {
			return x + y
		}

		type args struct {
			x int
			y int
		}
		tests := []struct {
			name      string
			args      args
			expect    int
			assertion assert.ComparisonAssertionFunc
		}{
			{"2+2=4", args{2, 2}, 4, assert.Equal},
			{"2+2!=5", args{2, 2}, 5, assert.NotEqual},
			{"2+3==5", args{2, 3}, 5, assert.Exactly},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tt.assertion(t, tt.expect, adder(tt.args.x, tt.args.y))
			})
		}
	}

	{
		dumbParseNum := func(input string, v interface{}) error {
			return json.Unmarshal([]byte(input), v)
		}

		tests := []struct {
			name      string
			arg       string
			assertion assert.ErrorAssertionFunc
		}{
			{"1.2 is number", "1.2", assert.NoError},
			{"1.2.3 not number", "1.2.3", assert.Error},
			{"true is not number", "true", assert.Error},
			{"3 is number", "3", assert.NoError},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				var x float64
				tt.assertion(t, dumbParseNum(tt.arg, &x))
			})
		}
	}

	{
		dumbParse := func(input string) interface{} {
			var x interface{}
			json.Unmarshal([]byte(input), &x)
			return x
		}

		tests := []struct {
			name      string
			arg       string
			assertion assert.ValueAssertionFunc
		}{
			{"true is not nil", "true", assert.NotNil},
			{"empty string is nil", "", assert.Nil},
			{"zero is not nil", "0", assert.NotNil},
			{"zero is zero", "0", assert.Zero},
			{"false is zero", "false", assert.Zero},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tt.assertion(t, dumbParse(tt.arg))
			})
		}
	}
}
