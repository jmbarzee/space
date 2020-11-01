package space

import (
	"math"
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
	if !FloatsEqual(a.X, b.X, MinErr) {
		return false
	}
	if !FloatsEqual(a.Y, b.Y, MinErr) {
		return false
	}
	if !FloatsEqual(a.Z, b.Z, MinErr) {
		return false
	}
	return true
}

func TestNewCartesian(t *testing.T) {
	sq2o2 := math.Sqrt2 / 2
	cases := []CartesianTest{
		{
			Operation: func(v Cartesian) Cartesian {
				o := NewOrientation(0, 0)
				return NewCartesian(o, 1)
			},
			Expected: Cartesian{0, 0, 1},
		},
		{
			Operation: func(v Cartesian) Cartesian {
				t := 0.0 / 4.0 * math.Pi
				p := 1.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewCartesian(o, 1)
			},
			Expected: Cartesian{0, 0, 1},
		},
		{
			Operation: func(v Cartesian) Cartesian {
				t := 1.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewCartesian(o, 1)
			},
			Expected: Cartesian{sq2o2, 0, sq2o2},
		},
		{
			Operation: func(v Cartesian) Cartesian {
				t := 2.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewCartesian(o, 1)
			},
			Expected: Cartesian{1, 0, 0},
		},
		{
			Operation: func(v Cartesian) Cartesian {
				t := 3.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewCartesian(o, 1)
			},
			Expected: Cartesian{sq2o2, 0, -sq2o2},
		},
		{
			Operation: func(v Cartesian) Cartesian {
				t := 4.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewCartesian(o, 1)
			},
			Expected: Cartesian{0, 0, -1},
		},
		{
			Operation: func(v Cartesian) Cartesian {
				t := 5.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewCartesian(o, 1)
			},
			Expected: Cartesian{-sq2o2, 0, -sq2o2},
		},
		{
			Operation: func(v Cartesian) Cartesian {
				t := 6.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewCartesian(o, 1)
			},
			Expected: Cartesian{-1, 0, 0},
		},
		{
			Operation: func(v Cartesian) Cartesian {
				t := 7.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewCartesian(o, 1)
			},
			Expected: Cartesian{-sq2o2, 0, sq2o2},
		},
		{
			Operation: func(v Cartesian) Cartesian {
				t := 8.0 / 4.0 * math.Pi
				p := 0.0 / 4.0 * math.Pi
				o := NewOrientation(t, p)
				return NewCartesian(o, 1)
			},
			Expected: Cartesian{0, 0, 1},
		},
	}
	RunCartesianTests(t, cases)
}
func TestCartesianTranslate(t *testing.T) {
	cases := []CartesianTest{
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{2, 3, 5}
				return v.Translate(u)
			},
			Expected: Cartesian{2, 3, 5},
		},
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{-2, -3, -5}
				return v.Translate(u)
			},
			Expected: Cartesian{-2, -3, -5},
		},
		{
			Initial: Cartesian{2, 3, 5},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{0, 0, 0}
				return v.Translate(u)
			},
			Expected: Cartesian{2, 3, 5},
		},
		{
			Initial: Cartesian{-2, -3, -5},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{0, 0, 0}
				return v.Translate(u)
			},
			Expected: Cartesian{-2, -3, -5},
		},
	}
	RunCartesianTests(t, cases)
}

func TestCartesianScale(t *testing.T) {
	cases := []CartesianTest{
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Cartesian) Cartesian {
				i := 0.0
				return v.Scale(i)
			},
			Expected: Cartesian{0, 0, 0},
		},
		{
			Initial: Cartesian{2, 3, 5},
			Operation: func(v Cartesian) Cartesian {
				i := 0.0
				return v.Scale(i)
			},
			Expected: Cartesian{0, 0, 0},
		},
		{
			Initial: Cartesian{-2, -3, -5},
			Operation: func(v Cartesian) Cartesian {
				i := 0.0
				return v.Scale(i)
			},
			Expected: Cartesian{0, 0, 0},
		},
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Cartesian) Cartesian {
				i := 1.1
				return v.Scale(i)
			},
			Expected: Cartesian{0, 0, 0},
		},
		{
			Initial: Cartesian{2, 3, 5},
			Operation: func(v Cartesian) Cartesian {
				i := 1.1
				return v.Scale(i)
			},
			Expected: Cartesian{2.2, 3.3, 5.5},
		},
		{
			Initial: Cartesian{-2, -3, -5},
			Operation: func(v Cartesian) Cartesian {
				i := 1.1
				return v.Scale(i)
			},
			Expected: Cartesian{-2.2, -3.3, -5.5},
		},
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Cartesian) Cartesian {
				i := -1.1
				return v.Scale(i)
			},
			Expected: Cartesian{0, 0, 0},
		},
		{
			Initial: Cartesian{2, 3, 5},
			Operation: func(v Cartesian) Cartesian {
				i := -1.1
				return v.Scale(i)
			},
			Expected: Cartesian{-2.2, -3.3, -5.5},
		},
		{
			Initial: Cartesian{-2, -3, -5},
			Operation: func(v Cartesian) Cartesian {
				i := -1.1
				return v.Scale(i)
			},
			Expected: Cartesian{2.2, 3.3, 5.5},
		},
	}
	RunCartesianTests(t, cases)
}

func TestCartesianTransform(t *testing.T) {
	cases := []CartesianTest{
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Cartesian) Cartesian {
				u := Matrix{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Cartesian{0, 0, 0},
		},
		{
			Initial: Cartesian{2, 3, 5},
			Operation: func(v Cartesian) Cartesian {
				u := Matrix{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Cartesian{2, 3, 5},
		},
		{
			Initial: Cartesian{2, 3, 5},
			Operation: func(v Cartesian) Cartesian {
				u := Matrix{
					{1, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Cartesian{2, 0, 0},
		},
		{
			Initial: Cartesian{2, 3, 5},
			Operation: func(v Cartesian) Cartesian {
				u := Matrix{
					{0, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Cartesian{0, 3, 0},
		},
		{
			Initial: Cartesian{2, 3, 5},
			Operation: func(v Cartesian) Cartesian {
				u := Matrix{
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Cartesian{0, 0, 5},
		},
		{
			Initial: Cartesian{2, 3, 5},
			Operation: func(v Cartesian) Cartesian {
				u := Matrix{
					{2, 0, 0, 0},
					{0, 2, 0, 0},
					{0, 0, 2, 0},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Cartesian{4, 6, 10},
		},
		{
			Initial: Cartesian{2, 3, 5},
			Operation: func(v Cartesian) Cartesian {
				u := Matrix{
					{1, 0, 0, 9},
					{0, 1, 0, 9},
					{0, 0, 1, 9},
					{0, 0, 0, 1},
				}
				return v.Transform(u)
			},
			Expected: Cartesian{11, 12, 14},
		},
	}
	RunCartesianTests(t, cases)
}

func TestCartesianProject(t *testing.T) {
	cases := []CartesianTest{
		{
			Initial: Cartesian{1, 0, 0},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{2, 0, 0}
				return v.Project(u)
			},
			Expected: Cartesian{2, 0, 0},
		},
		{
			Initial: Cartesian{0, 1, 0},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{2, 0, 0}
				return v.Project(u)
			},
			Expected: Cartesian{0, 0, 0},
		},
		{
			Initial: Cartesian{0, 0, 1},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{2, 0, 0}
				return v.Project(u)
			},
			Expected: Cartesian{0, 0, 0},
		},
		{
			Initial: Cartesian{1, 0, 0},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{0, 2, 0}
				return v.Project(u)
			},
			Expected: Cartesian{0, 0, 0},
		},
		{
			Initial: Cartesian{0, 1, 0},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{0, 2, 0}
				return v.Project(u)
			},
			Expected: Cartesian{0, 2, 0},
		},
		{
			Initial: Cartesian{0, 0, 1},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{0, 2, 0}
				return v.Project(u)
			},
			Expected: Cartesian{0, 0, 0},
		},
		{
			Initial: Cartesian{1, 0, 0},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{0, 0, 2}
				return v.Project(u)
			},
			Expected: Cartesian{0, 0, 0},
		},
		{
			Initial: Cartesian{0, 1, 0},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{0, 0, 2}
				return v.Project(u)
			},
			Expected: Cartesian{0, 0, 0},
		},
		{
			Initial: Cartesian{0, 0, 1},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{0, 0, 2}
				return v.Project(u)
			},
			Expected: Cartesian{0, 0, 2},
		},
		{
			Initial: Cartesian{1, 1, 1},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{6, 0, 0}
				return v.Project(u)
			},
			Expected: Cartesian{2, 2, 2},
		},
		{
			Initial: Cartesian{1, 1, 1},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{0, 6, 0}
				return v.Project(u)
			},
			Expected: Cartesian{2, 2, 2},
		},
		{
			Initial: Cartesian{1, 1, 1},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{0, 0, 6}
				return v.Project(u)
			},
			Expected: Cartesian{2, 2, 2},
		},
	}
	RunCartesianTests(t, cases)
}

func TestCartesianNegative(t *testing.T) {
	cases := []CartesianTest{
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Cartesian) Cartesian {
				return v.Negative()
			},
			Expected: Cartesian{0, 0, 0},
		},
		{
			Initial: Cartesian{1, 0, 0},
			Operation: func(v Cartesian) Cartesian {
				return v.Negative()
			},
			Expected: Cartesian{-1, 0, 0},
		},
		{
			Initial: Cartesian{0, 2, 3},
			Operation: func(v Cartesian) Cartesian {
				return v.Negative()
			},
			Expected: Cartesian{0, -2, -3},
		},
		{
			Initial: Cartesian{1, 2, 3},
			Operation: func(v Cartesian) Cartesian {
				return v.Negative()
			},
			Expected: Cartesian{-1, -2, -3},
		},
	}
	RunCartesianTests(t, cases)
}

func TestCartesianTranslationMatrix(t *testing.T) {
	cases := []CartesianTest{
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{2, 3, 5}
				m := u.TranslationMatrix()
				return v.Transform(m)
			},
			Expected: Cartesian{2, 3, 5},
		},
		{
			Initial: Cartesian{0, 0, 0},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{-2, -3, -5}
				m := u.TranslationMatrix()
				return v.Transform(m)
			},
			Expected: Cartesian{-2, -3, -5},
		},
		{
			Initial: Cartesian{2, 3, 5},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{0, 0, 0}
				m := u.TranslationMatrix()
				return v.Transform(m)
			},
			Expected: Cartesian{2, 3, 5},
		},
		{
			Initial: Cartesian{-2, -3, -5},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{0, 0, 0}
				m := u.TranslationMatrix()
				return v.Transform(m)
			},
			Expected: Cartesian{-2, -3, -5},
		},
		{
			Initial: Cartesian{2, 3, 5},
			Operation: func(v Cartesian) Cartesian {
				u := Cartesian{7, 8, 9}
				m := u.TranslationMatrix()
				return v.Transform(m)
			},
			Expected: Cartesian{9, 11, 14},
		},
	}
	RunCartesianTests(t, cases)
}
