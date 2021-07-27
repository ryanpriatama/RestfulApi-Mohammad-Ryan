package calculator

import (
	"testing"
)

func TestAddition(t *testing.T) {
	// misal a=2 (positif) dan b=1 (positif)
	if Addition(2, 1) != 3 {
		t.Error("Expected 2 (+) 1 to equal 3")
	}

	// misal a=-2 (negatif) dan b=1 (positif)
	if Addition(-2, 1) != -1 {
		t.Error("Expected -2 (+) 1 to equal -1")
	}

	// misal a=2 (positif) dan b=-1 (negatif)
	if Addition(2, -1) != 1 {
		t.Error("Expected 2 (+) -1 to equal 1")
	}

	// misal a=-2 (negatif) dan b=-1 (negatif)
	if Addition(-2, -1) != -3 {
		t.Error("Expected -2 (+) -1 to equal -3")
	}
}

func TestSubtraction(t *testing.T) {
	// misal a=2 (positif) dan b=1 (positif)
	if Subtraction(2, 1) != 1 {
		t.Error("Expected 2 (-) 1 to equal 1")
	}

	// misal a=-2 (negatif) dan b=1 (positif)
	if Subtraction(-2, 1) != -3 {
		t.Error("Expected -2 (-) 1 to equal -3")
	}

	// misal a=2 (positif) dan b=-1 (negatif)
	if Subtraction(2, -1) != 3 {
		t.Error("Expected 2 (-) -1 to equal 3")
	}

	// misal a=-2 (negatif) dan b=-1 (negatif)
	if Subtraction(-2, -1) != -1 {
		t.Error("Expected -2 (-) -1 to equal -1")
	}
}

func TestDivision(t *testing.T) {
	// misal a=2 (positif) dan b=1 (positif)
	if Division(2, 1) != 2 {
		t.Error("Expected 2 (/) 1 to equal 2")
	}

	// misal a=-2 (negatif) dan b=1 (positif)
	if Division(-2, 1) != -2 {
		t.Error("Expected -2 (/) 1 to equal -2")
	}

	// misal a=2 (positif) dan b=-1 (negatif)
	if Division(2, -1) != -2 {
		t.Error("Expected 2 (/) -1 to equal -2")
	}

	// misal a=-2 (negatif) dan b=-1 (negatif)
	if Division(-2, -1) != 2 {
		t.Error("Expected -2 (/) -1 to equal 2")
	}

	// menambahkan test jika hasil desimal, misal a= +-1 dan b= +-2
	// misal a=1 (positif) dan b=2 (positif)
	if Division(1, 2) != 0.5 {
		t.Error("Expected 1 (/) 2 to equal 0.5")
	}

	// misal a=-1 (negatif) dan b=2 (positif)
	if Division(-1, 2) != -0.5 {
		t.Error("Expected -1 (/) 2 to equal -0.5")
	}

	// misal a=2 (positif) dan b=-1 (negatif)
	if Division(1, -2) != -0.5 {
		t.Error("Expected 1 (/) -2 to equal -0.5")
	}

	// misal a=-2 (negatif) dan b=-1 (negatif)
	if Division(-1, -2) != 0.5 {
		t.Error("Expected -1 (/) -2 to equal 0.5")
	}
}

func TestMultiplication(t *testing.T) {
	// misal a=2 (positif) dan b=1 (positif)
	if Multiplication(2, 1) != 2 {
		t.Error("Expected 2 (*) 1 to equal 2")
	}

	// misal a=-2 (negatif) dan b=1 (positif)
	if Multiplication(-2, 1) != -2 {
		t.Error("Expected -2 (*) 1 to equal -2")
	}

	// misal a=2 (positif) dan b=-1 (negatif)
	if Multiplication(2, -1) != -2 {
		t.Error("Expected 2 (*) -1 to equal -2")
	}

	// misal a=-2 (negatif) dan b=-1 (negatif)
	if Multiplication(-2, -1) != 2 {
		t.Error("Expected -2 (*) -1 to equal 2")
	}
}
