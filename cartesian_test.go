package space

import (
	"testing"
)

type CartesianTest struct {
	Initial   Cartesian
	Operation func(Cartesian) Cartesian
	Expected  Cartesian
}

func RunCartesianTests(t *testing.T, cases []CartesianTest) {
	for i, c := range cases {
		actual := c.Operation(c.Initial)
		if !CartesiansEqual(c.Expected, actual) {
			t.Fatalf("Test %v failed. Cartesians were not equal:\n\tExpected: %v,\n\tActual: %v", i, c.Expected, actual)
		}
	}
}

// CartesiansEqual compares Cartesians
func CartesiansEqual(a, b Cartesian) bool {
	if !near(a.X, b.X) {
		return false
	}
	if !near(a.Y, b.Y) {
		return false
	}
	if !near(a.Z, b.Z) {
		return false
	}
	return true
}

func TestNewCartesian(t *testing.T) {
	cases := []CartesianTest{
		{
			Operation: func(v Cartesian) Cartesian {
				return NewCartesian(0, 0, 0)
			},
			Expected: Cartesian{0, 0, 0},
		},
		{
			Operation: func(v Cartesian) Cartesian {
				return NewCartesian(2, 3, 5)
			},
			Expected: Cartesian{2, 3, 5},
		},
		{
			Operation: func(v Cartesian) Cartesian {
				return NewCartesian(-2, -3, -5)
			},
			Expected: Cartesian{-2, -3, -5},
		},
	}
	RunCartesianTests(t, cases)
}

func TestCartesianSpherical(t *testing.T) {

	cases := make([]SphericalTest, len(AllEquivalencies))
	for i, p := range AllEquivalencies {
		c := p.Cartesian
		cases[i] = SphericalTest{
			Operation: func(v Spherical) Spherical {
				return c.Spherical()
			},
			Expected: p.Spherical,
		}
	}
	RunSphericalTests(t, cases)
}

func TestCartesianTranslationMatrix(t *testing.T) {
	cases := []CartesianTest{
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{2, 3, 5}
				m := u.TranslationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: Cartesian{2, 3, 5},
		},
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{-2, -3, -5}
				m := u.TranslationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: Cartesian{-2, -3, -5},
		},
		{
			Initial: Cartesian{2, 3, 5},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{0, 0, 0}
				m := u.TranslationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: Cartesian{2, 3, 5},
		},
		{
			Initial: Cartesian{-2, -3, -5},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{0, 0, 0}
				m := u.TranslationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: Cartesian{-2, -3, -5},
		},
		{
			Initial: Cartesian{2, 3, 5},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{7, 8, 9}
				m := u.TranslationMatrix()
				return v.Transform(m).Cartesian()
			},
			Expected: Cartesian{9, 11, 14},
		},
	}
	RunCartesianTests(t, cases)
}
