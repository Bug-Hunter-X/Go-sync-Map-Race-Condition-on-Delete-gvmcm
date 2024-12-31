```go
func main() {
    var m sync.Map
    m.Store("key", 1)
    value, ok := m.Load("key")
    fmt.Printf("Value: %v, OK: %t\n", value, ok)
    m.Delete("key")
    value, ok = m.Load("key")
    fmt.Printf("Value: %v, OK: %t\n", value, ok)
}
```