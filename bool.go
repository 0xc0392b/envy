package envy

import (
  "os"
  "strconv"
)

// GetEnvAsBool takes a string key k and a bool default value d and return the
// environment variable with the key k if it is not the empty string else
// returns the default value. Returns false if cannot parse value as bool.
func GetEnvAsBool(k string, d bool) bool {
  if s := os.Getenv(k); s != "" {
    if b, err := strconv.ParseBool(s); err != nil {
      return false
    } else {
      return b
    }
  } else {
    return d
  }
}
