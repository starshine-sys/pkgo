# pkgo

[![Go Reference](https://pkg.go.dev/badge/github.com/starshine-sys/pkgo.svg)](https://pkg.go.dev/github.com/starshine-sys/pkgo/v2) [![godocs.io](http://godocs.io/github.com/starshine-sys/pkgo?status.svg)](http://godocs.io/github.com/starshine-sys/pkgo/v2)

`pkgo` is a simple wrapper around [PluralKit](https://pluralkit.me/)'s REST API.

## Usage

Import the package into your project:

```go
import "github.com/starshine-sys/pkgo/v2"
```

All API actions are done through a `Session` struct. This can be authenticated or unauthenticated; for unauthenticated sessions, only public information will be returned.  
To get an API token, run `pk;token` on Discord.

```go
pk := pkgo.New("authentication token")
// or
pk := pkgo.New("")
```

### Note on branches/versions

The `main` branch contains pkgo v1, which supports version 1 of PluralKit's API. Like the API version it supports, it's deprecated and won't receive any further updates.  
The `v2` branch contains pkgo v2, which supports version 2 of PluralKit's API (with groups).

## Example

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
    Name:        pkgo.NewNullableString("Testing System"),
    Description: pkgo.NewNullableString("Hi, we're a system! 👋"),
    Tag:         pkgo.NewNullableString("| Testers"),
})
```
