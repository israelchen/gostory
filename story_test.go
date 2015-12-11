package gostory

import (
	"testing"
)

func TestNew(t *testing.T) {
	if s := New("test-new-story"); s.Name != "test-new-story" {
		t.Error("name is different.")
	}
}

func TestNewWithEmptyName(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic with empty name.")
		}
	}()

	_ = New("")
	t.Fail()
}

func TestAddData(t *testing.T) {
	s := New("test-add-data")
	s.AddData("userId", 123456789)

	value, ok := s.Data["userId"]

	if !ok {
		t.Error("expected userId key to be present in data.")
	}

	if value.Value.(int) != 123456789 {
		t.Error("value of userId is different than expected.")
	}
}

func TestAddDataWithEmptyKey(t *testing.T) {
	s := New("test-add-data-with-empty-key")

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic with empty data key.")
		}
	}()

	s.AddData("", "value without key")
	t.Fail()
}
