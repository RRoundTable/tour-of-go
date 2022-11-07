package ch03

import (
    "net"
    "testing"
)

func TestListener(t *testing.T) {
    listener, err := net.Listen("tcp", "127.0.0.1:0")
    if err != nil {
        t.Fetal(err)
    }
    defer func() { _ = listener.Close() }()

    t.Logf("Bound to %q", listener.Addr())
}


