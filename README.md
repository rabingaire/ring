# Ring Buffer

[![Go Version](https://img.shields.io/github/go-mod/go-version/rabingaire/ring)](https://tip.golang.org/doc/go1.17)
[![Go Report Card](https://goreportcard.com/badge/github.com/rabingaire/ring)](https://goreportcard.com/report/github.com/rabingaire/ring)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

```
type Ring interface {
	Put(value interface{})
	Get() interface{}
	Size() int64
	Capacity() int64
}
```

## Further Implementations

- [ ] Provide both thread safe and thread unsafe API
- [ ] Thread safe implementation using sync.Mutex
