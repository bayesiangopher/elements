// Package distance implements base linear algebra routines
// of numerical measurement of how far apart objects are.

// API for:
// 1) gonum/mat mat.VecDense and mat.Dense structures;
// 2) float64 arrays.

// Reading: https://en.wikipedia.org/wiki/Distance

package distance

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

type dimError struct {
	err       string
	firstDim  int
	secondDim int
}

func (e *dimError) Error() string {
	return fmt.Sprintf("%s: first dimension: %d, second dimension: %d",
		e.err, e.firstDim, e.secondDim)
}

// VecEuclidean calculates Euclidean (L2) distance between
// mat.VecDence vector instances.
// Reading: https://en.wikipedia.org/wiki/Euclidean_distance
// General form: d(p, q) = sqrt(pow(q1-p1, 2) + pow(q2-p2, 2) + ...)
func VecEuclidean(v, u *mat.VecDense) (dist float64) {
	if err := checkDim(v, u); err != nil {
		panic(err)
	}
	vr, _ := v.Dims()
	for i := 0; i < vr; i++ {
		dist += math.Pow(v.AtVec(i)-u.AtVec(i), 2)
	}
	return math.Sqrt(dist)
}

// VecL1 calculates taxicab distance (L1) between
// mat.VecDence vector instances.
// Reading: https://en.wikipedia.org/wiki/Taxicab_geometry
// General form: d(p, q) = ||q - p||1 = sum(|pi - qi|, 1, n)
func VecL1(v, u *mat.VecDense) (dist float64) {
	if err := checkDim(v, u); err != nil {
		panic(err)
	}
	vr, _ := v.Dims()
	for i := 0; i < vr; i++ {
		dist += math.Abs(v.AtVec(i) - u.AtVec(i))
	}
	return
}

// VecCanberra calculates canberra distance between
// mat.VecDence vector instances.
// Reading: https://en.wikipedia.org/wiki/Canberra_distance
// General form: d(p, q) = sum((|pi - qi|/|pi|+|qi|) ,1, n)
func VecCanberra(v, u *mat.VecDense) (dist float64) {
	if err := checkDim(v, u); err != nil {
		panic(err)
	}
	vr, _ := v.Dims()
	for i := 0; i < vr; i++ {
		numerator := math.Abs(v.AtVec(i) - u.AtVec(i))
		denominator := math.Abs(v.AtVec(i)) + math.Abs(u.AtVec(i))
		dist += numerator / denominator
	}
	return
}

// VecChebyshev calculates Chebyshev distance (Linf) between
// mat.VecDence vector instances.
// Reading: https://en.wikipedia.org/wiki/Chebyshev_distance
// General form: d(p, q) = max(|xi - yi|, i)
// whitch equal to: lim(sum(|xi - yi|^p, 1, n)^1/p, p -> inf)
func VecChebyshev(v, u *mat.VecDense) (dist float64) {
	if err := checkDim(v, u); err != nil {
		panic(err)
	}
	vr, _ := v.Dims()
	for i := 0; i < vr; i++ {
		dist = math.Max(dist, math.Abs(v.AtVec(i)-u.AtVec(i)))
	}
	return
}

func checkDim(N, M mat.Matrix) (err error){
	Nr, Nc := N.Dims()
	Mr, Mc := M.Dims()
	switch {
	case Nc == 1 && Mc == 1:
		if Nr != Mr {
			return &dimError{"Vectors dimension mismatch", Nr, Mr}
		}
	default:
		if Nc != Mr {
			return &dimError{"Matrices dimension mismatch", Nc, Mr}
		}
	}
	return nil
}
