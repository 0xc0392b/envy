package envy

import (
  "os"
  "testing"
)

var (
  // the test inputs
  testStrings = [2]stringTest{
    stringTest{
      Key: "foo",
      Value: "bar",
      Default: "test",
      Expected: "test",
      ShouldSet: false,
    },
    stringTest{
      Key: "bar",
      Value: "baz",
      Default: "test",
      Expected: "baz",
      ShouldSet: true,
    },
  }
)

// type stringTest represents an environment variable test. Contains a key,
// value, default value, expected return value, and should set flag - which
// tells the init loop to/to not set the variable before fetching it.
type stringTest struct {
  Key string
  Value string
  Default string
  Expected string
  ShouldSet bool
}

// TestGetEnvAsStr loops over an array of instances of stringTest and sets each
// key->value environment variable if ShouldSet is true, else is left unset.
// Then attempts to fetch the stringTest by its key and checks that the value
// returned matched Expected.
func TestGetEnvAsStr(t *testing.T) {
  for i := 0; i < len(testStrings); i++ {
    // set environment variables
    if testStrings[i].ShouldSet {
      os.Setenv(
        testStrings[i].Key,
        testStrings[i].Value,
      )
    }

    // get environment variables
    if v := GetEnvAsStr(
      testStrings[i].Key,
      testStrings[i].Default,
    ); v != testStrings[i].Expected {
      t.Errorf(
        "ENV %s is %s; expected %s",
        testStrings[i].Key,
        v,
        testStrings[i].Expected,
      )
    }
  }
}
