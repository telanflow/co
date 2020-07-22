package co

import "testing"

func TestMistGenerate(t *testing.T) {
	atomId := MistGenerate()
	t.Log(atomId)
}

func BenchmarkMistGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MistGenerate()
	}
}
