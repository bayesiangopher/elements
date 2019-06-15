package distance

import (
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestCheckDimVec(t *testing.T) {
	vec := mat.NewVecDense(3, []float64{0,1,2})
	vecOk := mat.NewVecDense(3, []float64{0,1,2})
	vecFail := mat.NewVecDense(2, []float64{0,1})

	if err := checkDim(vec, vecOk); err != nil {
		t.Error(err)
	}

	if err := checkDim(vec, vecFail); err == nil {
		t.Error("test failed: must be error")
	}
}

func TestCheckDimMat(t *testing.T) {
	matx := mat.NewDense(2, 2, []float64{0,1,2,3})
	matOk := mat.NewDense(2, 1, []float64{0,1})
	matFail := mat.NewDense(3, 1, []float64{0,1,2})

	if err := checkDim(matx, matOk); err != nil {
		t.Error(err)
	}

	if err := checkDim(matx, matFail); err == nil {
		t.Error("test failed: must be error")
	}
}

func TestVecEuclidean(t *testing.T) {
	vecFirst := mat.NewVecDense(3, []float64{0,1,2})
	vecSecond := mat.NewVecDense(3, []float64{3,4,5})

	if dist := VecEuclidean(vecFirst, vecSecond); dist != 5.196152422706632 {
		t.Errorf("dist: %g, but must be 5.196152422706632", dist)
	}
}

func TestVecL1(t *testing.T) {
	vecFirst := mat.NewVecDense(3, []float64{0,1,2})
	vecSecond := mat.NewVecDense(3, []float64{5,4,3})

	if dist := VecL1(vecFirst, vecSecond); dist != 9.0 {
		t.Errorf("dist: %g, but must be 9.0", dist)
	}
}

func TestVecCanberra(t *testing.T) {
	vecFirst := mat.NewVecDense(3, []float64{0,1,2})
	vecSecond := mat.NewVecDense(3, []float64{5,4,3})

	if dist := VecCanberra(vecFirst, vecSecond); dist != 1.8 {
		t.Errorf("dist: %g, but must be 1.8", dist)
	}
}

func TestVecChebyshev(t *testing.T) {
	vecFirst := mat.NewVecDense(3, []float64{0,1,2})
	vecSecond := mat.NewVecDense(3, []float64{5,4,3})

	if dist := VecChebyshev(vecFirst, vecSecond); dist != 5.0 {
		t.Errorf("dist: %g, but must be 5.0", dist)
	}
}
