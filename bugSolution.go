```go
func main() {
    var m sync.Map
    m.Store("key", 1)
    value, ok := m.Load("key")
    fmt.Printf("Value: %v, OK: %t\n", value, ok)

    // Proper way to handle concurrent delete and load
    m.LoadAndDelete("key")
    value, ok = m.Load("key")
    fmt.Printf("Value: %v, OK: %t\n", value, ok)
}
```