package tail

import (
	"fmt"
	"net"
	"time"

	"github.com/swissinfo-ch/logd/auth"
)

func Tail(conn net.Conn, readSecret []byte) (<-chan []byte, error) {
	sig, err := auth.Sign(readSecret, []byte("tail"), time.Now())
	if err != nil {
		return nil, fmt.Errorf("sign tail msg err: %w", err)
	}
	_, err = conn.Write(sig)
	if err != nil {
		return nil, fmt.Errorf("write tail msg err: %w", err)
	}
	out := make(chan []byte)
	go func(conn net.Conn) {
		for {
			buf := make([]byte, 2048)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Printf("error reading from conn: %s\r\n", err)
			}
			out <- buf[:n]
		}
	}(conn)
	go func(conn net.Conn, readSecret []byte) {
		for {
			time.Sleep(time.Second)
			sig, err := auth.Sign(readSecret, []byte("ping"), time.Now())
			if err != nil {
				fmt.Println("sign ping msg err:", err)
				continue
			}
			_, err = conn.Write(sig)
			if err != nil {
				fmt.Println("write ping msg err:", err)
			}
		}
	}(conn, readSecret)
	return out, nil
}
