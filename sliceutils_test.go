package sliceutils

import (
	"strings"
	"testing"
)

type user struct {
	email    string
	pass     string
	verified bool
}

var alice = user{
	email:    "alice@chains.io",
	pass:     "pass1234",
	verified: false,
}

var bob = user{
	email:    "bob@chains.io",
	pass:     "pass1234",
	verified: false,
}

func TestMapWithStrings(t *testing.T) {
	t.Parallel()
	input := []string{"hello", "world", "bye"}
	want := []string{"HELLO", "WORLD", "BYE"}
	got := Map(input, strings.ToUpper)
	if !equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestMapWithNumbers(t *testing.T) {
	t.Parallel()

	addTax := func(v int) int {
		return v * 21 / 100
	}

	input := []int{100, 200, 300}
	want := []int{21, 42, 63}
	got := Map(input, addTax)

	if !equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestMapWithStructs(t *testing.T) {
	t.Parallel()
	users := []user{alice, bob}

	verify := func(u user) user {
		u.verified = true
		return u
	}

	got := Map(users, verify)
	want := []user{
		{"alice@chains.io", "pass1234", true},
		{"bob@chains.io", "pass1234", true},
	}

	if !equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestFilterWithNumbers(t *testing.T) {
	t.Parallel()

	positives := func(n int) bool {
		return n >= 0
	}

	input := []int{-99, 200, 300}
	want := []int{200, 300}
	got := Filter(input, positives)

	if !equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestFilterWithStructs(t *testing.T) {
	t.Parallel()

	jane := user{
		email:    "jane@chains.io",
		pass:     "pass1234",
		verified: true,
	}

	users := []user{bob, alice, jane}
	want := []user{jane}

	got := Filter(users, func(u user) bool {
		return u.verified
	})

	if !equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestFind(t *testing.T) {
	t.Parallel()
	strs := []string{"hello", "world", "bye"}

	finder := func(term string) func(s string) bool {
		return func(s string) bool {
			return strings.Contains(s, term)
		}
	}

	byeFinder := finder("bye")
	w, ok := Find(strs, byeFinder)
	if !ok || w != "bye" {
		t.Errorf("want to find word bye, got %q", w)
	}
}

func TestFind_NoMatch(t *testing.T) {
	t.Parallel()
	strs := []string{"hello", "world", "bye"}

	finder := func(term string) func(s string) bool {
		return func(s string) bool {
			return strings.Contains(s, term)
		}
	}

	ciaoFinder := finder("ciao")
	w, ok := Find(strs, ciaoFinder)
	if ok || w != "" {
		t.Errorf("want to find nothing, got %q", w)
	}
}

func TestSome(t *testing.T) {
	t.Parallel()
	strs := []string{"hello", "world", "bye"}

	isWorld := func(s string) bool {
		return s == "world"
	}

	got := Some(strs, isWorld)
	if !got {
		t.Error("want to find a match and get true")
	}
}

func TestSome_NoMatch(t *testing.T) {
	t.Parallel()
	strs := []string{"hello", "planet", "bye"}

	isWorld := func(s string) bool {
		return s == "world"
	}

	got := Some(strs, isWorld)
	if got {
		t.Error("want to find no matches and get false")
	}
}

func TestEvery(t *testing.T) {
	users := []user{alice, bob}

	pendingVerification := func(u user) bool {
		return !u.verified
	}

	got := Every(users, pendingVerification)
	if !got {
		t.Error("want all entries to match and get true")
	}
}

func TestEvery_NoneMatch(t *testing.T) {
	users := []user{alice, bob}

	verified := func(u user) bool {
		return u.verified
	}

	got := Every(users, verified)
	if got {
		t.Error("want not all entries to match and get false")
	}
}

func TestEvery_NotAllMatch(t *testing.T) {
	users := []user{alice, bob}

	isAlice := func(u user) bool {
		return u.email == "alice@chains.io"
	}

	got := Every(users, isAlice)
	if got {
		t.Error("want not all entries to match and get false")
	}
}

func TestForEach(t *testing.T) {
	u1 := &user{"A@B.c", "abc", false}
	users := []*user{u1}

	lowercaseEmail := func(u *user) {
		u.email = strings.ToLower(u.email)
	}

	ForEach(users, lowercaseEmail)

	if u1.email != "a@b.c" {
		t.Errorf("want email a@b.c, got %s", u1.email)
	}
}

func TestReduce(t *testing.T) {
	t.Parallel()

	input := []float64{10, 10, 7, 8, 10}
	want := 45.0

	sum := Reduce(input, func(acc, n float64) float64 {
		return acc + n
	}, 0.0)

	if sum != want {
		t.Errorf("want %f, got %f", want, sum)
	}
}

func equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
