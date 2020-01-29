package envy

import (
  "os"
  "strconv"
)

// GetEnvAsFloat32 takes a string key k and a float32 default value d and
// returns the environment variable with the key k if it is not the empty string
// else returns the default value. Returns -1.0 if cannot parse value as float.
func GetEnvAsFloat32(k string, d float32) float32 {
  if s := os.Getenv(k); s != "" {
    if f, err := strconv.ParseFloat(s, 32); err != nil {
      return -1.0
    } else {
      return float32(f)
    }
  } else {
    return d
  }
}

// GetEnvAsFloat32 takes a string key k and a float64 default value d and
// returns the environment variable with the key k if it is not the empty string
// else returns the default value. Returns -1.0 if cannot parse value as float.
func GetEnvAsFloat64(k string, d float64) float64 {
  if s := os.Getenv(k); s != "" {
    if f, err := strconv.ParseFloat(s, 64); err != nil {
      return -1.0
    } else {
      return f
    }
  } else {
    return d
  }
}
