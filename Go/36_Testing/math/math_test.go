package math_test

import (
	"testing"

	"github.com/Kumaravi-admin/lesson_testing/math"
)

func TestMinInt(t *testing.T) {
	result := math.MinInt(10, 5)
	if result != 5 {
		t.Errorf("expected: %d, got:%d", 5, result)
	}
}


func TestMaxInt(t *testing.T) {
	result := math.MaxInt(10,100 )
	if result != 100 {
		t.Errorf("expected: %d, got: %d", 100, result)
	}
}

/* terminla output

github.com/Kumaravi-admin/lesson_testing/math   2.539s
PS D:\LearnCode\Go\36_Testing> go test ./...
ok      github.com/Kumaravi-admin/lesson_testing/math   (cached)
PS D:\LearnCode\Go\36_Testing> go test ./...
ok      github.com/Kumaravi-admin/lesson_testing/math   (cached)


*/