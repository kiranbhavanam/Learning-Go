package main 
import "testing"
func Sum(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}
// Benchmark for Sum function
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1000) // Run with a fixed value
	}
}
