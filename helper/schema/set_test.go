package schema

import (
	"reflect"
	"testing"
)

func TestSetAdd(t *testing.T) {
	s := &Set{F: testSetInt}
	s.Add(1)
	s.Add(5)
	s.Add(25)

	expected := []interface{}{1, 5, 25}
	actual := s.List()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("bad: %#v", actual)
	}
}

func TestSetContains(t *testing.T) {
	s := &Set{F: testSetInt}
	s.Add(5)

	if s.Contains(2) {
		t.Fatal("should not contain")
	}
	if !s.Contains(5) {
		t.Fatal("should contain")
	}
}

func TestSetDifference(t *testing.T) {
	s1 := &Set{F: testSetInt}
	s2 := &Set{F: testSetInt}

	s1.Add(1)
	s1.Add(5)

	s2.Add(5)
	s2.Add(25)

	expected := []interface{}{1}
	actual := s1.Difference(s2).List()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("bad: %#v", actual)
	}
}

func TestSetIntersection(t *testing.T) {
	s1 := &Set{F: testSetInt}
	s2 := &Set{F: testSetInt}

	s1.Add(1)
	s1.Add(5)

	s2.Add(5)
	s2.Add(25)

	expected := []interface{}{5}
	actual := s1.Intersection(s2).List()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("bad: %#v", actual)
	}
}

func TestSetUnion(t *testing.T) {
	s1 := &Set{F: testSetInt}
	s2 := &Set{F: testSetInt}

	s1.Add(1)
	s1.Add(5)

	s2.Add(5)
	s2.Add(25)

	expected := []interface{}{1, 5, 25}
	actual := s1.Union(s2).List()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("bad: %#v", actual)
	}
}

func testSetInt(v interface{}) int {
	return v.(int)
}
