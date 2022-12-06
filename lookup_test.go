package benchmarklookup_test

import (
	"math/rand"
	"strconv"
	"testing"

	benchmarklookup "github.com/ngicks/benchmark-lookup"
)

func outlet(v any) {

}

func fillSlice(target []string) {
	for i := 0; i < len(target); i++ {
		target[i] = strconv.FormatInt(int64(i), 10)
	}
}

func fillMap(mm map[string]struct{}, size int) {
	for i := 0; i < size; i++ {
		mm[strconv.FormatInt(int64(i), 10)] = struct{}{}
	}
}

func randValue() []string {
	sl := make([]string, 100000)
	for i := 0; i < len(sl); i++ {
		sl[i] = strconv.FormatInt(rand.Int63n(100000), 10)
	}
	return sl
}

var randIndex = randValue()

func repeatNTime(repeat int, testFn func(attempt int)) {
	for i := 0; i < repeat; i++ {
		testFn(i)
	}
}

func BenchmarkLookUpSlice1(b *testing.B) {
	sl := make([]string, 1)
	fillSlice(sl)
	testFn := func(i int) {
		idx := benchmarklookup.LookUpSlice(sl, randIndex[i])
		outlet(idx)
	}
	b.ResetTimer()
	repeatNTime(1000, testFn)
}

func BenchmarkLookUpSlice5(b *testing.B) {
	sl := make([]string, 5)
	fillSlice(sl)
	testFn := func(i int) {
		idx := benchmarklookup.LookUpSlice(sl, randIndex[i])
		outlet(idx)
	}
	b.ResetTimer()
	repeatNTime(1000, testFn)
}
func BenchmarkLookUpSlice10(b *testing.B) {
	sl := make([]string, 10)
	fillSlice(sl)
	testFn := func(i int) {
		idx := benchmarklookup.LookUpSlice(sl, randIndex[i])
		outlet(idx)
	}
	b.ResetTimer()
	repeatNTime(1000, testFn)
}

func BenchmarkLookUpSlice15(b *testing.B) {
	sl := make([]string, 15)
	fillSlice(sl)
	testFn := func(i int) {
		idx := benchmarklookup.LookUpSlice(sl, randIndex[i])
		outlet(idx)
	}
	b.ResetTimer()
	repeatNTime(1000, testFn)
}

func BenchmarkLookUpSlice25(b *testing.B) {
	sl := make([]string, 25)
	fillSlice(sl)
	testFn := func(i int) {
		idx := benchmarklookup.LookUpSlice(sl, randIndex[i])
		outlet(idx)
	}
	b.ResetTimer()
	repeatNTime(1000, testFn)
}
func BenchmarkLookUpSlice50(b *testing.B) {
	sl := make([]string, 50)
	fillSlice(sl)
	testFn := func(i int) {
		idx := benchmarklookup.LookUpSlice(sl, randIndex[i])
		outlet(idx)
	}
	b.ResetTimer()
	repeatNTime(1000, testFn)
}

func BenchmarkLookUpSlice100(b *testing.B) {
	sl := make([]string, 100)
	fillSlice(sl)
	testFn := func(i int) {
		idx := benchmarklookup.LookUpSlice(sl, randIndex[i])
		outlet(idx)
	}
	b.ResetTimer()
	repeatNTime(1000, testFn)
}

func BenchmarkLookUpSlice1000(b *testing.B) {
	sl := make([]string, 1000)
	fillSlice(sl)
	testFn := func(i int) {
		idx := benchmarklookup.LookUpSlice(sl, randIndex[i])
		outlet(idx)
	}
	b.ResetTimer()
	repeatNTime(1000, testFn)
}
func BenchmarkLookUpMap1(b *testing.B) {
	mm := make(map[string]struct{})
	fillMap(mm, 1)

	testFn := func(i int) {
		_, ok := mm[randIndex[i]]
		if ok {
			outlet(i)
		}
	}

	b.ResetTimer()
	repeatNTime(1000, testFn)
}

func BenchmarkLookUpMap5(b *testing.B) {
	mm := make(map[string]struct{})
	fillMap(mm, 5)

	testFn := func(i int) {
		_, ok := mm[randIndex[i]]
		if ok {
			outlet(i)
		}
	}
	b.ResetTimer()

	repeatNTime(1000, testFn)
}
func BenchmarkLookUpMap10(b *testing.B) {
	mm := make(map[string]struct{})
	fillMap(mm, 10)

	testFn := func(i int) {
		_, ok := mm[randIndex[i]]
		if ok {
			outlet(i)
		}
	}
	b.ResetTimer()

	repeatNTime(1000, testFn)
}

func BenchmarkLookUpMap15(b *testing.B) {
	mm := make(map[string]struct{})
	fillMap(mm, 15)

	testFn := func(i int) {
		_, ok := mm[randIndex[i]]
		if ok {
			outlet(i)
		}
	}
	b.ResetTimer()

	repeatNTime(1000, testFn)
}

func BenchmarkLookUpMap50(b *testing.B) {
	mm := make(map[string]struct{})
	fillMap(mm, 50)

	testFn := func(i int) {
		_, ok := mm[randIndex[i]]
		if ok {
			outlet(i)
		}
	}
	b.ResetTimer()

	repeatNTime(1000, testFn)
}
func BenchmarkLookUpMap100(b *testing.B) {
	mm := make(map[string]struct{})
	fillMap(mm, 100)

	testFn := func(i int) {
		_, ok := mm[randIndex[i]]
		if ok {
			outlet(i)
		}
	}
	b.ResetTimer()

	repeatNTime(1000, testFn)
}

func BenchmarkLookUpMap1000(b *testing.B) {
	mm := make(map[string]struct{})
	fillMap(mm, 1000)

	testFn := func(i int) {
		_, ok := mm[randIndex[i]]
		if ok {
			outlet(i)
		}
	}
	b.ResetTimer()
	repeatNTime(1000, testFn)
}
