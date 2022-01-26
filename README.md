# Ring Buffer

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
