package store

import (
	"testing"
)

func TestStore(t *testing.T) {
	s := New()

	// Test Set and Get.
	s.Set("foo", "bar")
	if value, ok := s.Get("foo"); !ok || value != "bar" {
		t.Errorf("Get(foo) = %q, %v, want %q, true", value, ok, "bar")
	}

	// Test Del.
	if ok := s.Del("foo"); !ok {
		t.Errorf("Del(foo) = false, want true")
	}
	if _, ok := s.Get("foo"); ok {
		t.Errorf("Get(foo) after Del(foo) = _, true, want _, false")
	}
}
