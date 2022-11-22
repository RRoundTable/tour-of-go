package ch03

import (
    "net"
    "testing"
)

func TestListener(t *testing.T) {
    listener, err := net.Listen("tcp", "127.0.0.1:20000")
    if err != nil {
        t.Fatal(err)
    }
    defer func() { _ = listener.Close() }()

    t.Logf("Bound to %q", listener.Addr())
    for {
        conn, err := listener.Accept()
        t.Log(conn)
        if err != nil {
            t.Log(err)
            return
        }
        go func(c net.Conn) {
            defer c.Close()
            t.Log("Connected")
        }(conn)
    }
}


