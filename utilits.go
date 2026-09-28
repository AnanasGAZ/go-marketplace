package main

import "math"

// проверяет, не будет ли переполнение при сложении
func checkedAddNonNegative(a, b int64) (int64, bool) {
	if a < 0 || b < 0 || a > math.MaxInt64-b {
		return 0, false
	}
	return a + b, true
}

// проверяет, не будет ли переполнения в списке
func checkedMulPositive(a, b int64) (int64, bool) {
	//если a*b  слишком велко - переполнение
	//a * b <= limit
	//a <= limit / b
	// если a > limit / b - выдадим ошибку
	if a <= 0 || b <= 0 || a > math.MaxInt64/b {
		return 0, false
	}
	return a * b, true
}
