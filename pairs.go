package space

import (
	"math"
)

const MinErr = 0.000001

type vectorEquivalency struct {
	Cartesian Cartesian
	Spherical Spherical
}

func rad(n, d int) float64 {
	return float64(n) / float64(d) * math.Pi
}

func near(a, b float64) bool {
	return float64(math.Abs(a-b)) < MinErr
}

var sqrt2o2 = math.Sqrt2 / 2.0

var AllEquivalencies = []vectorEquivalency{
	Origin,

	AxisX,
	AxisXN,
	AxisY,
	AxisYN,
	AxisZ,
	AxisZN,

	AxisX3,
	AxisXN3,
	AxisY3,
	AxisYN3,
	AxisZ3,
	AxisZN3,

	OctantXYZ,
	OctantNXYZ,
	OctantXNYZ,
	OctantNXNYZ,
	OctantXYNZ,
	OctantNXYNZ,
	OctantXNYNZ,
	OctantNXNYNZ,

	OctantXYZ3,
	OctantNXYZ3,
	OctantXNYZ3,
	OctantNXNYZ3,
	OctantXYNZ3,
	OctantNXYNZ3,
	OctantXNYNZ3,
	OctantNXNYNZ3,
}

// Origin
var (
	Origin = vectorEquivalency{
		Cartesian: Cartesian{
			X: 0,
			Y: 0,
			Z: 0,
		},
		Spherical: Spherical{
			R: 0,
			T: 0,
			P: 0,
		},
	}
)

// Axes with length of one
var (
	AxisX = vectorEquivalency{
		Cartesian: Cartesian{
			X: 1,
			Y: 0,
			Z: 0,
		},
		Spherical: Spherical{
			R: 1,
			T: 0,
			P: rad(1, 2),
		},
	}

	AxisXN = vectorEquivalency{
		Cartesian: Cartesian{
			X: -1,
			Y: 0,
			Z: 0,
		},
		Spherical: Spherical{
			R: 1,
			T: rad(1, 1),
			P: rad(1, 2),
		},
	}

	AxisY = vectorEquivalency{
		Cartesian: Cartesian{
			X: 0,
			Y: 1,
			Z: 0,
		},
		Spherical: Spherical{
			R: 1,
			T: rad(1, 2),
			P: rad(1, 2),
		},
	}

	AxisYN = vectorEquivalency{
		Cartesian: Cartesian{
			X: 0,
			Y: -1,
			Z: 0,
		},
		Spherical: Spherical{
			R: 1,
			T: rad(3, 2),
			P: rad(1, 2),
		},
	}

	AxisZ = vectorEquivalency{
		Cartesian: Cartesian{
			X: 0,
			Y: 0,
			Z: 1,
		},
		Spherical: Spherical{
			R: 1,
			T: 0,
			P: 0,
		},
	}

	AxisZN = vectorEquivalency{
		Cartesian: Cartesian{
			X: 0,
			Y: 0,
			Z: -1,
		},
		Spherical: Spherical{
			R: 1,
			T: 0,
			P: rad(1, 1),
		},
	}
)

// Axes with length of three
var (
	AxisX3 = vectorEquivalency{
		Cartesian: Cartesian{
			X: 3,
			Y: 0,
			Z: 0,
		},
		Spherical: Spherical{
			R: 3,
			T: 0,
			P: rad(1, 2),
		},
	}

	AxisXN3 = vectorEquivalency{
		Cartesian: Cartesian{
			X: -3,
			Y: 0,
			Z: 0,
		},
		Spherical: Spherical{
			R: 3,
			T: rad(1, 1),
			P: rad(1, 2),
		},
	}

	AxisY3 = vectorEquivalency{
		Cartesian: Cartesian{
			X: 0,
			Y: 3,
			Z: 0,
		},
		Spherical: Spherical{
			R: 3,
			T: rad(1, 2),
			P: rad(1, 2),
		},
	}

	AxisYN3 = vectorEquivalency{
		Cartesian: Cartesian{
			X: 0,
			Y: -3,
			Z: 0,
		},
		Spherical: Spherical{
			R: 3,
			T: rad(3, 2),
			P: rad(1, 2),
		},
	}

	AxisZ3 = vectorEquivalency{
		Cartesian: Cartesian{
			X: 0,
			Y: 0,
			Z: 3,
		},
		Spherical: Spherical{
			R: 3,
			T: 0,
			P: 0,
		},
	}

	AxisZN3 = vectorEquivalency{
		Cartesian: Cartesian{
			X: 0,
			Y: 0,
			Z: -3,
		},
		Spherical: Spherical{
			R: 3,
			T: 0,
			P: rad(1, 1),
		},
	}
)

// Octants with length of one
var (
	OctantXYZ = vectorEquivalency{
		Cartesian: Cartesian{
			X: 0.5773502669,
			Y: 0.5773502669,
			Z: 0.5773502669,
		},
		Spherical: Spherical{
			R: 1,
			T: rad(1, 4),
			P: 0.304086724 * math.Pi,
		},
	}
	OctantNXYZ = vectorEquivalency{
		Cartesian: Cartesian{
			X: -0.5773502669,
			Y: 0.5773502669,
			Z: 0.5773502669,
		},
		Spherical: Spherical{
			R: 1,
			T: rad(3, 4),
			P: 0.304086724 * math.Pi,
		},
	}
	OctantNXNYZ = vectorEquivalency{
		Cartesian: Cartesian{
			X: -0.5773502669,
			Y: -0.5773502669,
			Z: 0.5773502669,
		},
		Spherical: Spherical{
			R: 1,
			T: rad(5, 4),
			P: 0.304086724 * math.Pi,
		},
	}
	OctantXNYZ = vectorEquivalency{
		Cartesian: Cartesian{
			X: 0.5773502669,
			Y: -0.5773502669,
			Z: 0.5773502669,
		},
		Spherical: Spherical{
			R: 1,
			T: rad(7, 4),
			P: 0.304086724 * math.Pi,
		},
	}
	OctantXYNZ = vectorEquivalency{
		Cartesian: Cartesian{
			X: 0.5773502669,
			Y: 0.5773502669,
			Z: -0.5773502669,
		},
		Spherical: Spherical{
			R: 1,
			T: rad(1, 4),
			P: 0.695913276 * math.Pi,
		},
	}
	OctantNXYNZ = vectorEquivalency{
		Cartesian: Cartesian{
			X: -0.5773502669,
			Y: 0.5773502669,
			Z: -0.5773502669,
		},
		Spherical: Spherical{
			R: 1,
			T: rad(3, 4),
			P: 0.695913276 * math.Pi,
		},
	}
	OctantNXNYNZ = vectorEquivalency{
		Cartesian: Cartesian{
			X: -0.5773502669,
			Y: -0.5773502669,
			Z: -0.5773502669,
		},
		Spherical: Spherical{
			R: 1,
			T: rad(5, 4),
			P: 0.695913276 * math.Pi,
		},
	}
	OctantXNYNZ = vectorEquivalency{
		Cartesian: Cartesian{
			X: 0.5773502669,
			Y: -0.5773502669,
			Z: -0.5773502669,
		},
		Spherical: Spherical{
			R: 1,
			T: rad(7, 4),
			P: 0.695913276 * math.Pi,
		},
	}
)

// Octants with length of three
var (
	OctantXYZ3 = vectorEquivalency{
		Cartesian: Cartesian{
			X: 1.7320508007,
			Y: 1.7320508007,
			Z: 1.7320508007,
		},
		Spherical: Spherical{
			R: 3,
			T: rad(1, 4),
			P: 0.304086724 * math.Pi,
		},
	}
	OctantNXYZ3 = vectorEquivalency{
		Cartesian: Cartesian{
			X: -1.7320508007,
			Y: 1.7320508007,
			Z: 1.7320508007,
		},
		Spherical: Spherical{
			R: 3,
			T: rad(3, 4),
			P: 0.304086724 * math.Pi,
		},
	}
	OctantNXNYZ3 = vectorEquivalency{
		Cartesian: Cartesian{
			X: -1.7320508007,
			Y: -1.7320508007,
			Z: 1.7320508007,
		},
		Spherical: Spherical{
			R: 3,
			T: rad(5, 4),
			P: 0.304086724 * math.Pi,
		},
	}
	OctantXNYZ3 = vectorEquivalency{
		Cartesian: Cartesian{
			X: 1.7320508007,
			Y: -1.7320508007,
			Z: 1.7320508007,
		},
		Spherical: Spherical{
			R: 3,
			T: rad(7, 4),
			P: 0.304086724 * math.Pi,
		},
	}
	OctantXYNZ3 = vectorEquivalency{
		Cartesian: Cartesian{
			X: 1.7320508007,
			Y: 1.7320508007,
			Z: -1.7320508007,
		},
		Spherical: Spherical{
			R: 3,
			T: rad(1, 4),
			P: 0.695913276 * math.Pi,
		},
	}
	OctantNXYNZ3 = vectorEquivalency{
		Cartesian: Cartesian{
			X: -1.7320508007,
			Y: 1.7320508007,
			Z: -1.7320508007,
		},
		Spherical: Spherical{
			R: 3,
			T: rad(3, 4),
			P: 0.695913276 * math.Pi,
		},
	}
	OctantNXNYNZ3 = vectorEquivalency{
		Cartesian: Cartesian{
			X: -1.7320508007,
			Y: -1.7320508007,
			Z: -1.7320508007,
		},
		Spherical: Spherical{
			R: 3,
			T: rad(5, 4),
			P: 0.695913276 * math.Pi,
		},
	}
	OctantXNYNZ3 = vectorEquivalency{
		Cartesian: Cartesian{
			X: 1.7320508007,
			Y: -1.7320508007,
			Z: -1.7320508007,
		},
		Spherical: Spherical{
			R: 3,
			T: rad(7, 4),
			P: 0.695913276 * math.Pi,
		},
	}
)
