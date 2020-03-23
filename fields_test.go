package log

import (
	"reflect"
	"testing"
)

func TestCombineFields(t *testing.T) {
	// empty fields
	{
		old := map[string]interface{}{
			"ka": "va",
			"kb": "vb",
		}

		have, err := combineFields(old, nil)
		if err != nil {
			t.Error("want nil")
			return
		}
		if !reflect.DeepEqual(have, old) {
			t.Errorf("have:%v, old:%v", have, old)
			return
		}

		// check if it is a new one
		for k := range old {
			delete(old, k)
		}
		if len(have) == len(old) {
			t.Errorf("want not equal, have:%v, old:%v", have, old)
			return
		}
	}
	// ddd number of field
	{
		old := map[string]interface{}{
			"ka": "va",
			"kb": "vb",
		}

		have, err := combineFields(old, []interface{}{"kc", "vc", "kd"})
		if err != _ErrNumberOfFieldsMustNotBeOdd {
			t.Error("want equal")
			return
		}
		if !reflect.DeepEqual(have, old) {
			t.Errorf("have:%v, old:%v", have, old)
			return
		}

		// check if it is a new one
		for k := range old {
			delete(old, k)
		}
		if len(have) == len(old) {
			t.Errorf("want not equal, have:%v, old:%v", have, old)
			return
		}
	}
	// nil old
	{
		have, err := combineFields(nil, []interface{}{"kc", "vc", "kb", "vb2"})
		if err != nil {
			t.Error("want nil")
			return
		}
		want := map[string]interface{}{
			"kb": "vb2",
			"kc": "vc",
		}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	// non-nil old
	{
		old := map[string]interface{}{
			"ka": "va",
			"kb": "vb",
		}

		have, err := combineFields(old, []interface{}{"kc", "vc", "kb", "vb2"})
		if err != nil {
			t.Error("want nil")
			return
		}
		want := map[string]interface{}{
			"ka": "va",
			"kb": "vb2",
			"kc": "vc",
		}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}

		// check if it is a new one
		if reflect.DeepEqual(have, old) {
			t.Errorf("want not equal, have:%v, old:%v", have, old)
			return
		}
	}
	// non-nil old with non-string type of key
	{
		old := map[string]interface{}{
			"ka": "va",
			"kb": "vb",
		}

		have, err := combineFields(old, []interface{}{"kc", "vc", "kb", "vb2", 1, 2, "kd", "vd"})
		if err != _ErrTypeOfFieldKeyMustBeString {
			t.Error("want equal")
			return
		}
		want := map[string]interface{}{
			"ka": "va",
			"kb": "vb2",
			"kc": "vc",
		}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}

		// check if it is a new one
		if reflect.DeepEqual(have, old) {
			t.Errorf("want not equal, have:%v, old:%v", have, old)
			return
		}
	}
	// non-nil old with empty key
	{
		old := map[string]interface{}{
			"ka": "va",
			"kb": "vb",
		}

		have, err := combineFields(old, []interface{}{"kc", "vc", "kb", "vb2", "", "vd", "ke", "ve"})
		if err != _ErrFieldKeyMustNotBeEmpty {
			t.Error("want equal")
			return
		}
		want := map[string]interface{}{
			"ka": "va",
			"kb": "vb2",
			"kc": "vc",
		}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}

		// check if it is a new one
		if reflect.DeepEqual(have, old) {
			t.Errorf("want not equal, have:%v, old:%v", have, old)
			return
		}
	}
}

func TestCloneFields(t *testing.T) {
	// nil source
	{
		var m map[string]interface{}
		m2 := cloneFields(m)
		if m2 != nil {
			t.Error("want nil")
			return
		}
	}
	// non-nil source
	{
		m := map[string]interface{}{
			"a": "va",
			"b": "vb",
		}
		m2 := cloneFields(m)
		if !reflect.DeepEqual(m, m2) {
			t.Errorf("m2:%v, m:%v", m2, m)
			return
		}
		// check if it is a new one
		for k := range m {
			delete(m, k)
		}
		if len(m2) == len(m) {
			t.Errorf("want not equal, m2:%v, m:%v", m2, m)
			return
		}
	}
}
