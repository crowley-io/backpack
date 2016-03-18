package engine

import (
	"reflect"
	"testing"
)

func TestParseConfigurationWithRandomFile(t *testing.T) {

	if _, err := ParseConfiguration("../testdata/random.yml"); err == nil {
		t.Errorf("was expecting an error")
	}

}

func TestParseConfigurationWithInvalidFilePath(t *testing.T) {

	if _, err := ParseConfiguration("../testdata/random"); err == nil {
		t.Errorf("was expecting an error")
	}

}

func TestParseConfiguration(t *testing.T) {

	e := &Configuration{}
	e.Backpack.Prehooks = []string{"foo", "bar"}
	e.Backpack.Posthooks = []string{"fuu", "beer"}
	e.Backpack.Execute = []string{"make", "make test"}

	c, err := ParseConfiguration("../testdata/simple.yml")

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if !reflect.DeepEqual(e, c) {
		t.Errorf("was expecting '%+v' but received '%+v'", e, c)
	}

	if !reflect.DeepEqual(c.Prehooks(), e.Backpack.Prehooks) {
		t.Error("Prehooks() doesn't match Prehooks in struct")
	}

	if !reflect.DeepEqual(c.Posthooks(), e.Backpack.Posthooks) {
		t.Error("Posthooks() doesn't match Posthooks in struct")
	}

	if !reflect.DeepEqual(c.Execute(), e.Backpack.Execute) {
		t.Error("Execute() doesn't match Execute in struct")
	}

}
