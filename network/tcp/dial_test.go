package ch03

import (
    "io"
    "net"
    "testing"
)


func TestDial(t *testing.T){
    listener, err := net.Listen("tcp", "127.0.0.1:")
    if err != nil {
        t.Fatal(err)
    }

    done := make(chan struct{})
    go func () {
        // Server side
        defer func() { done <- struct{}{} }()

        for {
            t.Log("Waiting for connection")
            conn, err := listener.Accept()
            if err != nil {
                t.Log(err)
                return
            }
            t.Log("Server side Connection Accepted", conn)
            go func(c net.Conn) {
                defer func() {
                    t.Log("Server close connection")
                    c.Close()
                    done <- struct{}{}
                }()
                buf := make([]byte, 1024)
                for {
                    n, err := c.Read(buf)
                    if err != nil {
                        if err != io.EOF {
                            t.Error(err)
                        }
                        return
                    }
                    t.Logf("received: %q", buf[:n])
                }
            }(conn)
        }
    }()
    // Client side
    t.Log(listener.Addr())
    t.Log("start dial")
    conn, err := net.Dial("tcp", listener.Addr().String())
    t.Log("Client side connection", conn)
    if err != nil {
        t.Fatal(err)
    }
    conn.Close()
    t.Log("Client close connection")
    <- done
    listener.Close()
    <- done

}
