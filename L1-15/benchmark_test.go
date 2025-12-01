package main

import "testing"

func BenchmarkBadVersion(b *testing.B) {
    v := createHugeString(1 << 10)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = v[:100]
    }
}

func BenchmarkGoodVersion(b *testing.B) {
    v := createHugeString(1 << 10)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = string([]rune(v[:100]))
    }
}