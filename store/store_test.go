package store_test

import (
	"os"
	"store"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetReturnsNotOKIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/kvtest.store"
	s, err := store.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	_, ok := s.Get("doesnotexist")
	if ok {
		t.Fatal("unexpected ok")
	}
}

func TestGetReturnsOKIfKeyExists(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/kvtest.store"
	s, err := store.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	s.Set("exists", "value)")

	_, ok := s.Get("exists")
	if !ok {
		t.Fatal("unexpected not ok")
	}
}

func TestSetCorrectlySetsKeyValuePairInStore(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/kvtest.store"
	kv, err := store.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	key := "key"
	want := "value"

	kv.Set(key, want)

	got, ok := kv.Get(key)

	if !ok {
		t.Fatal(ok)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSetCorrectlyUpdatesKeyValuePairInStore(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/kvtest.store"
	kv, err := store.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	key := "exists"
	want := "value"
	originalValue := "original value"

	kv.Set(key, originalValue)
	kv.Set(key, want)

	got, ok := kv.Get(key)

	if !ok {
		t.Fatal(ok)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSaveSavesTheStoreCorrectly(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		key   string
		value string
	}

	path := t.TempDir() + "/kvtest.store"
	kv, err := store.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	testCases := []TestCase{
		{key: "A", value: "1"},
		{key: "B", value: "2"},
		{key: "C", value: "3"},
		{key: "D", value: "4"},
		{key: "E", value: "5"},
	}

	for _, tc := range testCases {
		kv.Set(tc.key, tc.value)
	}

	err = kv.Save()
	if err != nil {
		t.Fatal(err)
	}

	kv2, err := store.OpenStore(path)
	if err != nil {
		t.Fatal()
	}

	for _, tc := range testCases {
		v, err := kv2.Get(tc.key)
		if !err {
			t.Error("Expected OK, got False")
		}
		if v != tc.value {
			t.Errorf("want %q, got %q", tc.value, v)
		}
	}
}

func TestOpenStore_ErrorsWhenPathUnreadable(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/unreadable.store"
	if _, err := os.Create(path); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(path, 0o000); err != nil {
		t.Fatal(err)
	}
	_, err := store.OpenStore(path)
	if err == nil {
		t.Fatal("no error")
	}
}

func TestOpenStore_ReturnsErrorOnInvalidData(t *testing.T) {
	t.Parallel()
	_, err := store.OpenStore("testdata/invalid.store")
	if err == nil {
		t.Fatal("no error")
	}
}

func TestSaveErrorsWhenPathUnwritable(t *testing.T) {
	t.Parallel()
	s, err := store.OpenStore("bogus/unwritable.store")
	if err != nil {
		t.Fatal(err)
	}
	err = s.Save()
	if err == nil {
		t.Fatal("no error")
	}
}
