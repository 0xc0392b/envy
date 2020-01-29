package envy

import (
  "os"
  "testing"
)

var (
  // the 32-bit test inputs
  testFloat32s = []float32Test{
    float32Test{
      Key: "foo",
      Value: "1.23",
      Default: 3.56,
      Expected: 3.56,
      ShouldSet: false,
    },
    float32Test{
      Key: "bar",
      Value: "3.14",
      Default: 3.14,
      Expected: 3.14,
      ShouldSet: true,
    },
    float32Test{
      Key: "baz",
      Value: "test",
      Default: 0.0,
      Expected: -1.0,
      ShouldSet: true,
    },
  }

  // the 64-bit test inputs
  testFloat64s = []float64Test{
    float64Test{
      Key: "foo",
      Value: "1.23",
      Default: 3.56,
      Expected: 3.56,
      ShouldSet: false,
    },
    float64Test{
      Key: "bar",
      Value: "3.14",
      Default: 3.14,
      Expected: 3.14,
      ShouldSet: true,
    },
    float64Test{
      Key: "baz",
      Value: "test",
      Default: 0.0,
      Expected: -1.0,
      ShouldSet: true,
    },
  }
)

// type float32Test represents an environment variable test. Contains a key,
// value, default value, expected return value, and should set flag - which
// tells the init loop to/to not set the variable before fetching it.
type float32Test struct {
  Key string
  Value string
  Default float32
  Expected float32
  ShouldSet bool
}

// type float64Test represents an environment variable test. Contains a key,
// value, default value, expected return value, and should set flag - which
// tells the init loop to/to not set the variable before fetching it.
type float64Test struct {
  Key string
  Value string
  Default float64
  Expected float64
  ShouldSet bool
}

// TestGetEnvAsFloat32 loops over an array of instances of float32Test and sets
// each key->value environment variable if ShouldSet is true, else is left
// unset. Then attempts to fetch the float32Test by its key and checks that the
// value returned matched Expected.
func TestGetEnvAsFloat32(t *testing.T) {
  for i := 0; i < len(testFloat32s); i++ {
    // set environment variables
    if testFloat32s[i].ShouldSet {
      os.Setenv(
        testFloat32s[i].Key,
        testFloat32s[i].Value,
      )
    }

    // get environment variables
    if v := GetEnvAsFloat32(
      testFloat32s[i].Key,
      testFloat32s[i].Default,
    ); v != testFloat32s[i].Expected {
      t.Errorf(
        "ENV %s is %f; expected %f",
        testFloat32s[i].Key,
        v,
        testFloat32s[i].Expected,
      )
    }
  }
}

// TestGetEnvAsFloat64 loops over an array of instances of float64Test and sets
// each key->value environment variable if ShouldSet is true, else is left
// unset. Then attempts to fetch the float64Test by its key and checks that the
// value returned matched Expected.
func TestGetEnvAsFloat64(t *testing.T) {
  for i := 0; i < len(testFloat64s); i++ {
    // set environment variables
    if testFloat64s[i].ShouldSet {
      os.Setenv(
        testFloat64s[i].Key,
        testFloat64s[i].Value,
      )
    }

    // get environment variables
    if v := GetEnvAsFloat64(
      testFloat64s[i].Key,
      testFloat64s[i].Default,
    ); v != testFloat64s[i].Expected {
      t.Errorf(
        "ENV %s is %f; expected %f",
        testFloat64s[i].Key,
        v,
        testFloat64s[i].Expected,
      )
    }
  }
}
