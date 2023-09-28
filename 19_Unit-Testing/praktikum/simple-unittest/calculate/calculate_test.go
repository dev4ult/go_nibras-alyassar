package calculate

import "testing"

func TestAddition(t *testing.T) {
	if Addition(1, 2) != 3 {
		t.Error("1 (+) 2 supposed equal to 3")
	}

	if Addition(-1, -2) != -3 {
		t.Error("-1 (+) -2 supposed equal to -3")
	}
}

func TestSubstraction(t *testing.T) {
	if Substraction(1, 2) != -1 {
		t.Error("1 (-) 2 supposed equal to -1")
	}

	if Substraction(5, 3) != 2 {
		t.Error("5 (-) 3 supposed equal to 2")
	}

	if Substraction(3, 3) != 0 {
		t.Error("3 (-) 3 supposed equal to 0")
	}
}

func TestDivision(t *testing.T) {
	if Division(1, 2) != 0.5 {
		t.Error("1 (/) 2 supposed equal to 0.5")
	}

	if Division(5, 2) != 2.5 {
		t.Error("5 (/) 2 supposed equal to 2.5")
	}

	if Division(9, 3) != 3 {
		t.Error("9 (/) 3 supposed equal to 3")
	}

	if Division(3, 3) != 1 {
		t.Error("3 (/) 3 supposed equal to 1")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(3, 2) != 6 {
		t.Error("3 (X) 2 supposed equal to 6")
	}

	if Multiplication(5, 0) != 0 {
		t.Error("5 (X) 0 supposed equal to 0")
	}

	if Multiplication(3, 3) != 9 {
		t.Error("3 (X) 3 supposed equal to 9")
	}
}

