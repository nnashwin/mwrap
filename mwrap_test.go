package mwrap

import (
	"gopkg.in/mgo.v2"
	"os"
	"reflect"
	"testing"
)

func TestNewSession(t *testing.T) {
	testSess := mgo.Session{}
	expected := &testSess
	actual, err := NewSession("mongodb://localhost")
	if err != nil {
		os.Exit(1)
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(expected) {
		t.Errorf("Test failed, not retrieving mongo session")
	}
}

func TestGetColl(t *testing.T) {
	testColl := mgo.Collection{}
	expected := &testColl
	actual, err := GetColl("mongodb://localhost", "test", "test")
	if err != nil {
		os.Exit(1)
	}

	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		t.Errorf("Test failed, did not retrieve collection")
	}
}
