package tests

import "testing"

type OneWord int64

type FourWords struct {
	a OneWord
	b OneWord
	c OneWord
	d OneWord
}

type EightWords struct {
	a FourWords
	b FourWords
}

type SixteenWords struct {
	a EightWords
	b EightWords
}

type ThirtyTwoWords struct {
	a SixteenWords
	b SixteenWords
}

type SixtyFourWords struct {
	a ThirtyTwoWords
	b ThirtyTwoWords
}

type HundredTwentyEightWords struct {
	a SixtyFourWords
	b SixtyFourWords
}

func BenchmarkPassByReferenceOneWord(b *testing.B) {
	f := func(s *OneWord) {}
	s := OneWord(0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(&s)
	}
}

func BenchmarkPassByValueOneWord(b *testing.B) {
	f := func(s OneWord) {}
	s := OneWord(0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(s)
	}
}

func BenchmarkPassByReferenceFourWords(b *testing.B) {
	f := func(s *FourWords) {}
	s := FourWords{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(&s)
	}
}

func BenchmarkPassByValueFourWords(b *testing.B) {
	f := func(s FourWords) {}
	s := FourWords{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(s)
	}
}

func BenchmarkPassByReferenceEightWords(b *testing.B) {
	f := func(s *EightWords) {}
	s := EightWords{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(&s)
	}
}

func BenchmarkPassByValueEightWords(b *testing.B) {
	f := func(s EightWords) {}
	s := EightWords{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(s)
	}
}

func BenchmarkPassByReferenceSixteenWords(b *testing.B) {
	f := func(s *SixteenWords) {}
	s := SixteenWords{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(&s)
	}
}

func BenchmarkPassByValueSixteenWords(b *testing.B) {
	f := func(s SixteenWords) {}
	s := SixteenWords{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(s)
	}
}

func BenchmarkPassByReferenceThirtyTwoWords(b *testing.B) {
	f := func(s *ThirtyTwoWords) {}
	s := ThirtyTwoWords{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(&s)
	}
}

func BenchmarkPassByValueThirtyTwoWords(b *testing.B) {
	f := func(s ThirtyTwoWords) {}
	s := ThirtyTwoWords{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(s)
	}
}

func BenchmarkPassByReferenceSixtyFourWords(b *testing.B) {
	f := func(s *SixtyFourWords) {}
	s := SixtyFourWords{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(&s)
	}
}

func BenchmarkPassByValueSixtyFourWords(b *testing.B) {
	f := func(s SixtyFourWords) {}
	s := SixtyFourWords{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(s)
	}
}

func BenchmarkPassByReferenceHundredTwentyEightWords(b *testing.B) {
	f := func(s *HundredTwentyEightWords) {}
	s := HundredTwentyEightWords{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(&s)
	}
}

func BenchmarkPassByValueHundredTwentyEightWords(b *testing.B) {
	f := func(s HundredTwentyEightWords) {}
	s := HundredTwentyEightWords{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(s)
	}
}
