# envy
Environment variables but with float, int, and bool parsing and default values

Envy is a library for fetching and parsing environment variables. Envy makes it
easier to fetch and parse ints, floats, and bools from system env variables, and
set default 'fallback' values.

### strings
```go
import (
  "github.com/0xc0392b/envy"
)

func main() {
  e := envy.GetEnvAsStr("SOME_STR_VARIABLE", "default value")
  // use e
}
```

### ints
```go
import (
  "github.com/0xc0392b/envy"
)

func main() {
  e32 := envy.GetEnvAsInt32("SOME_INT32_VARIABLE", 1)
  // use e32

  e64 := envy.GetEnvAsInt64("SOME_INT64_VARIABLE", 1)
  // use e64
}
```

### floats
```go
import (
  "github.com/0xc0392b/envy"
)

func main() {
  e32 := envy.GetEnvAsFloat32("SOME_FLOAT32_VARIABLE", 3.14)
  // use e32

  e64 := envy.GetEnvAsFloat64("SOME_FLOAT64_VARIABLE", 3.14)
  // use e64
}
```

### bools
```go
import (
  "github.com/0xc0392b/envy"
)

func main() {
  e := envy.GetEnvAsBool("SOME_BOOL_VARIABLE", true)
  // use e
}
```
