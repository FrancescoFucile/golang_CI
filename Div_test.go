package ci_test_1

import (
	log "github.com/romana/rlog"
	"reflect"
	"testing"
)

func TestDiv(t *testing.T) {
	t.Run("non zero", testDiv(input{6, 3}, expected{2, nil}))
	t.Run("zero", testDiv(input{2, 0}, expected{0, reflect.TypeOf(new(MathError))}))
}

type input struct {
	a, b float32
}

type expected struct {
	res     float32
	errType reflect.Type
}

func testDiv(in input, exp expected) func(t *testing.T) {
	return func(t *testing.T) {
		res, err := Div(in.a, in.b)
		log.Debugf("err type: %v", reflect.TypeOf(err))
		if res != exp.res ||
			reflect.TypeOf(err) != exp.errType {
			t.Fail()
		}
	}
}
