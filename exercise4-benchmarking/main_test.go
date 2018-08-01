package main

import "testing"

// Assigning this out of scope var the value from each run of the
// benchmark in order to prevent compiler inlining functions as an optimisation.
var result string

func BenchmarkWithPlus(b *testing.B) {
	var res string
	for n := 0; n < b.N; n++ {
		res = withPlus()
	}
	result = res
}

func BenchmarkWithSprintf(b *testing.B) {
	var res string
	for n := 0; n < b.N; n++ {
		res = withSprintf()
	}
	result = res
}

func BenchmarkWithBuffer(b *testing.B) {
	var res string
	for n := 0; n < b.N; n++ {
		res = withBuffer()
	}
	result = res
}

func BenchmarkWithStringBuilder(b *testing.B) {
	var res string
	for n := 0; n < b.N; n++ {
		result = withStringBuilder()
	}
	result = res
}
