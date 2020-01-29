package envy

import (
  "os"
  "testing"
)

var (
  // the 32-bit test inputs
  testInt32s = []int32Test{
    int32Test{
      Key: "foo",
      Value: "1",
      Default: 2,
      Expected: 2,
      ShouldSet: false,
    },
    int32Test{
      Key: "bar",
      Value: "2",
      Default: 3,
      Expected: 2,
      ShouldSet: true,
    },
    int32Test{
      Key: "baz",
      Value: "test",
      Default: 0,
      Expected: -1,
      ShouldSet: true,
    },
  }

  // the 64-bit test inputs
  testInt64s = []int64Test{
    int64Test{
      Key: "foo",
      Value: "1",
      Default: 2,
      Expected: 2,
      ShouldSet: false,
    },
    int64Test{
      Key: "bar",
      Value: "2",
      Default: 3,
      Expected: 2,
      ShouldSet: true,
    },
    int64Test{
      Key: "baz",
      Value: "test",
      Default: 0,
      Expected: -1,
      ShouldSet: true,
    },
  }
)

// type int32Test represents an environment variable test. Contains a key,
// value, default value, expected return value, and should set flag - which
// tells the init loop to/to not set the variable before fetching it.
type int32Test struct {
  Key string
  Value string
  Default int32
  Expected int32
  ShouldSet bool
}

// type int64Test represents an environment variable test. Contains a key,
// value, default value, expected return value, and should set flag - which
// tells the init loop to/to not set the variable before fetching it.
type int64Test struct {
  Key string
  Value string
  Default int64
  Expected int64
  ShouldSet bool
}

// TestGetEnvAsInt32 loops over an array of instances of int32Test and sets
// each key->value environment variable if ShouldSet is true, else is left
// unset. Then attempts to fetch the int32Test by its key and checks that the
// value returned matched Expected.
func TestGetEnvAsInt32(t *testing.T) {
  for i := 0; i < len(testInt32s); i++ {
    // set environment variables
    if testInt32s[i].ShouldSet {
      os.Setenv(
        testInt32s[i].Key,
        testInt32s[i].Value,
      )
    }

    // get environment variables
    if v := GetEnvAsInt32(
      testInt32s[i].Key,
      testInt32s[i].Default,
    ); v != testInt32s[i].Expected {
      t.Errorf(
        "ENV %s is %d; expected %d",
        testInt32s[i].Key,
        v,
        testInt32s[i].Expected,
      )
    }
  }
}

// TestGetEnvAsInt64 loops over an array of instances of int64Test and sets
// each key->value environment variable if ShouldSet is true, else is left
// unset. Then attempts to fetch the int64Test by its key and checks that the
// value returned matched Expected.
func TestGetEnvAsInt64(t *testing.T) {
  for i := 0; i < len(testInt64s); i++ {
    // set environment variables
    if testInt64s[i].ShouldSet {
      os.Setenv(
        testInt64s[i].Key,
        testInt64s[i].Value,
      )
    }

    // get environment variables
    if v := GetEnvAsInt64(
      testInt64s[i].Key,
      testInt64s[i].Default,
    ); v != testInt64s[i].Expected {
      t.Errorf(
        "ENV %s is %d; expected %d",
        testInt64s[i].Key,
        v,
        testInt64s[i].Expected,
      )
    }
  }
}
