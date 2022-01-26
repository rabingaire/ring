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

- [ ] Thread safe implementation using sync.Mutex
