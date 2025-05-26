package RybakovMymathDz8

import (
	"fmt"
	"math"
)

func FuctorialFunc(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * FuctorialFunc(n-1)
}

/*
Ограничения для a ∈ R, b ∈ R, b = m/n, где m ∈ R, n ∈ R, но n != 0, соответсвенно a^(m/n).
Ограничения:
Если a = 0,(n < 0, m > 0 || m < 0, n > 0);
Если a = 0, m = n;
Если a < 0, n = 2g, g ∈ N;

для a ∈ R, b ∈ R, соответсвенно a^b.
Если a = 0, b < 0;
Если a = 0, b = 0;

Ограничения для a ∈ R, b ∈ R, соответсвенно log_b a.
Если a <= 0;
Если b = 1, b <= 0;
*/

// a^b = e^(b*ln(a))
func RTailorDegree(a, b float64) (float64, error) {
	if a == 0 && b == 0 {
		return 0, fmt.Errorf("Неопределенность")
	}
	if a == 0 && b < 0 {
		return 0, fmt.Errorf("Неопределенность")
	}

	return (b * math.Log(a)), nil
}

func FuncRTailorDegree(a, n, m float64, RootDegree func(float64, float64) (float64, error)) (float64, error) {
	root, err := RootDegree(n, m)
	if err != nil {
		return 0, err
	}
	if a == 0 && n == 0 && m == 0 {
		return 0, fmt.Errorf("Неопределенность")
	}
	if a == 0 && root < 0 {
		return 0, fmt.Errorf("Неопределенность")
	}
	if -a < 0 && math.Mod(m, 2.0) == 0.0 {
		return 0, fmt.Errorf("Неопределенность")
	}

	return (root * math.Log(a)), nil
}

// ограничения для ln
func LnLimitations(a, b float64) (float64, float64, error) {
	if b <= 0 || b == 1 {
		return 0, 0, fmt.Errorf("Неправильное заданное b")
	}
	if a <= 0 {
		return 0, 0, fmt.Errorf("Неправильное заданное a")
	}

	return a, b, nil
}

// Степень с произвольными покащателями
func RootDegree(n, m float64) (float64, error) {
	return (n / m), nil
}

func RoundIfClose(x float64) float64 {
	result := math.Round(x)
	if math.Abs(x-result) < 0.001 {
		return result
	}
	return x
}
