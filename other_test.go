package goutil

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitPointer(t *testing.T) {
	type T struct {
		A string
		B int
	}
	var i **********T
	v := reflect.ValueOf(&i)
	done := InitPointer(v)
	assert.True(t, done)
	assert.Equal(t, T{}, **********i)
}

func TestReferenceSlice(t *testing.T) {
	v := reflect.ValueOf([]int{1, 2})
	v = ReferenceSlice(v, 1)
	ret := v.Interface().([]*int)
	t.Logf("%#v", ret)

	v = reflect.ValueOf([]int{})
	v = ReferenceSlice(v, 1)
	ret = v.Interface().([]*int)
	t.Logf("%#v", ret)
}

func TestDereferenceSlice(t *testing.T) {
	one := 1
	two := 2
	v := reflect.ValueOf([]*int{&one, &two})
	v = DereferenceSlice(v)
	ret := v.Interface().([]int)
	t.Logf("%#v", ret)

	v = reflect.ValueOf([]*int{})
	v = DereferenceSlice(v)
	ret = v.Interface().([]int)
	t.Logf("%#v", ret)
}
