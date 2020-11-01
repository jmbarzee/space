package space

import "testing"

type OrientationTest struct {
	Initial   Orientation
	Operation func(Orientation) Orientation
	Expected  Orientation
}

func RunOrientationTests(t *testing.T, cases []OrientationTest) {
	for i, c := range cases {
		actual := c.Operation(c.Initial)
		if !OrientationsEqual(c.Expected, actual) {
			t.Fatalf("Test %v failed. Orientations were not equal:\n\tExpected: %v,\n\tActual: %v", i, c.Expected, actual)
		}
	}
}

// OrientationsEqual compares Orientations
func OrientationsEqual(a, b Orientation) bool {
	if !FloatsEqual(a.Theta, b.Theta, MinErr) {
		return false
	}
	if !FloatsEqual(a.Phi, b.Phi, MinErr) {
		return false
	}
	return true
}
