# pkgo

[![Go Reference](https://pkg.go.dev/badge/github.com/starshine-sys/pkgo.svg)](https://pkg.go.dev/github.com/starshine-sys/pkgo)

`pkgo` is a simple wrapper around [PluralKit](https://pluralkit.me/)'s REST API.

## Usage

Import the package into your project:

```go
import "github.com/starshine-sys/pkgo"
```

All API actions are done through a `Session` struct. This can be authenticated or unauthenticated; for unauthenticated sessions, only public information will be returned.

```go
pk := pkgo.New("authentication token")
// or
pk := pkgo.New("")
```

### Example

```go
sysID := "exmpl"

pk := pkgo.New("")

sys, err := pk.System(sysID)
// ID => exmpl
// Name => PluralKit Example System
// Tag => | PluralKit 🦊
// Created => 2020-01-12 02:00:33.387824 +0000 UTC

front, err := pk.Fronters(sysID)
// Timestamp => 2020-01-12 02:21:44.024493 +0000 UTC
// Members => [Myriad Kit Tester T. Testington]

msg, err := s.Message(859157735644069928)
// ID => 859157735644069928
// Original => 859157734331252777
// Sender => 694563574386786314
```

```go
pk = pkgo.New("notARealToken")

sys, err = pk.EditSystem(pkgo.EditSystemData{
    Name:        "Testing System",
    Description: "Hi, we're a system! 👋",
    Tag:         "| Testers",
})
```