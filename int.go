package envy

import (
  "os"
  "strconv"
)

// GetEnvAsFloat32 takes a string key k and a int32 default value d and returns
// the environment variable with the key k if it is not the empty string
// else returns the default value. Returns -1 if cannot parse value as int.
func GetEnvAsInt32(k string, d int32) int32 {
  if s := os.Getenv(k); s != "" {
    if i, err := strconv.ParseInt(s, 10, 32); err != nil {
      return -1
    } else {
      return int32(i)
    }
  } else {
    return d
  }
}

// GetEnvAsFloat64 takes a string key k and a int64 default value d and returns
// the environment variable with the key k if it is not the empty string
// else returns the default value. Returns -1 if cannot parse value as int.
func GetEnvAsInt64(k string, d int64) int64 {
  if s := os.Getenv(k); s != "" {
    if i, err := strconv.ParseInt(s, 10, 64); err != nil {
      return -1
    } else {
      return i
    }
  } else {
    return d
  }
}
