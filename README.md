# Go sync.Map Race Condition on Delete

This repository demonstrates a potential race condition in Go's `sync.Map` when concurrently deleting and loading keys.  Under specific timing conditions, a deleted key may still return a value when `Load()` is called.  This can lead to unexpected behavior and hard-to-debug issues in concurrent applications.

The `bug.go` file contains code that reproduces the problem. The `bugSolution.go` file provides a possible solution.