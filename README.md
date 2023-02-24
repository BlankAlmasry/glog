# Glog Verbosity Wrapper

This is a wrapper for glog that allows you to set verbosity levels.

The following verbosity levels are supported:

* `0` - `Exit`, `Fatal`, `Error`, `Warning`, `Info`
* `1` - `Exit`, `Fatal`, `Error`, `Warning`
* `2` - `Exit`, `Fatal`, `Error`
* `3` - `Exit`, `Fatal`
* `4` - `Exit`
* `5` - `Disable all logging`

To set the verbosity level, use the `flag.Set` function:

```go
flag.Set("v", "0") // Fatal, Error, Warning, Info
flag.Set("v", "1") // Fatal, Error, Warning
flag.Set("v", "2") // Fatal, Error
flag.Set("v", "3") // Fatal
flag.Set("v", "4") // Disable all logging
```

## Example Usage

```go
package main

import (
    "github.com/golang/glog"
    "github.com/example/glog"
)

func main() {
    // Set the verbosity level to 1 (Fatal, Error, Warning)
    flag.Set("v", "1")

    // Log a message at the "Error" level
    glog.Error("This is an error message")

    // Log a message at the "Warning" level
    glog.Warning("This is a warning message")

    // Log a message at the "Fatal" level
    glog.Fatal("This is a fatal message")

    // This message will not be logged since the verbosity level is set to 1
    glog.Info("This is an info message")
}
```
