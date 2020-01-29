package envy

import (
  "os"
  "testing"
)

var (
  // the test inputs
  testBools = [3]boolTest{
    boolTest{
      Key: "foo",
      Value: "bar",
      Default: true,
      Expected: true,
      ShouldSet: false,
    },
    boolTest{
      Key: "bar",
      Value: "false",
      Default: false,
      Expected: false,
      ShouldSet: true,
    },
    boolTest{
      Key: "baz",
      Value: "test",
      Default: false,
      Expected: false,
      ShouldSet: true,
    },
  }
)

// type boolTest represents an environment variable test. Contains a key,
// value, default value, expected return value, and should set flag - which
// tells the init loop to/to not set the variable before fetching it.
type boolTest struct {
  Key string
  Value string
  Default bool
  Expected bool
  ShouldSet bool
}

// TestGetEnvAsStr loops over an array of instances of boolTest and sets each
// key->value environment variable if ShouldSet is true, else is left unset.
// Then attempts to fetch the boolTest by its key and checks that the value
// returned matched Expected.
func TestGetEnvAsBool(t *testing.T) {
  for i := 0; i < len(testBools); i++ {
    // set environment variables
    if testBools[i].ShouldSet {
      os.Setenv(
        testBools[i].Key,
        testBools[i].Value,
      )
    }

    // get environment variables
    if v := GetEnvAsBool(
      testBools[i].Key,
      testBools[i].Default,
    ); v != testBools[i].Expected {
      t.Errorf(
        "ENV %s is %t; expected %t",
        testBools[i].Key,
        v,
        testBools[i].Expected,
      )
    }
  }
}
