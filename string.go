package envy

import (
  "os"
)

// GetEnvAsStr takes a string key k and string default value d and returns the
// environment variable with the key k if it is not the empty string else
// returns the default value.
func GetEnvAsStr(k, d string) string {
  if s := os.Getenv(k); s != "" {
    return s
  } else {
    return d
  }
}
