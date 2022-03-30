# Ring Buffer

[![Go Reference](https://pkg.go.dev/badge/github.com/rabingaire/ring.svg)](https://pkg.go.dev/github.com/rabingaire/ring)
[![Go Report Card](https://goreportcard.com/badge/github.com/rabingaire/ring)](https://goreportcard.com/report/github.com/rabingaire/ring)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

```
type Ring[T any] interface {
	Put(value T)
	Get() T
	Size() int64
	Capacity() int64
}
```
