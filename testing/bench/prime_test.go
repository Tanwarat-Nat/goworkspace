package bench

import "testing"

func TestIsPrime(t *testing.T) {
	if !isPrime(7867) {
		t.Error("7867 is not prime")
	}
	if !isPrimeV2(7867) {
		t.Error("7867 is not prime")
	}
	if isPrime(10) {
		t.Error("10 is prime")
	}
	if isPrimeV2(10) {
		t.Error("10 is prime")
	}
}

// Benchmark[Name] (*testing.B){}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime(7867)
	}
}

func BenchmarkIsPrimeV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrimeV2(7867)
	}
}

func BenchmarkIsPrimeV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrimeV3(7867)
	}
}

/*
อยากรู้ว่าการ execute prime แต่ละครั้งมันใช่เวลานานขนาดไหน
benchmark มีซิกเนเจอร์ของมันอยู่คล้ายๆเทส


*/
