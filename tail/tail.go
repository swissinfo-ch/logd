/*
Copyright © 2024 JOSEPH INNES <avianpneuma@gmail.com>
*/
package tail

import (
	"fmt"
	"net"
	"time"

	"github.com/swissinfo-ch/logd/auth"
	"github.com/swissinfo-ch/logd/cmd"
	"github.com/swissinfo-ch/logd/transport"
	"google.golang.org/protobuf/proto"
)

type Format string

const (
	Plain = Format("plain")
	Msg   = Format("msg")
)

func Tail(conn net.Conn, readSecret []byte, format Format) (<-chan any, error) {
	payload, err := proto.Marshal(&cmd.Cmd{
		Name: cmd.Name_TAIL,
	})
	if err != nil {
		return nil, fmt.Errorf("marshal ping msg err: %w", err)
	}
	sig, err := auth.Sign(readSecret, payload, time.Now())
	if err != nil {
		return nil, fmt.Errorf("sign tail msg err: %w", err)
	}
	_, err = conn.Write(sig)
	if err != nil {
		return nil, fmt.Errorf("write tail msg err: %w", err)
	}
	out := make(chan any)
	go read(conn, out, format)
	go ping(conn, readSecret)
	return out, nil
}

func read(conn net.Conn, out chan<- any, format Format) {
	for {
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("error reading from conn: %s\r\n", err)
		}
		switch format {
		case Plain:
			out <- buf[:n]
		case Msg:
			m := &cmd.Msg{}
			err := proto.Unmarshal(buf[:n], m)
			if err != nil {
				fmt.Println("unpack msg err:", err)
				continue
			}
			out <- m
		}
	}
}

func ping(conn net.Conn, readSecret []byte) {
	for {
		time.Sleep(transport.PingPeriod)
		payload, err := proto.Marshal(&cmd.Cmd{
			Name: cmd.Name_PING,
		})
		if err != nil {
			fmt.Println("marshal ping msg err:", err)
			continue
		}
		sig, err := auth.Sign(readSecret, payload, time.Now())
		if err != nil {
			fmt.Println("sign ping msg err:", err)
			continue
		}
		_, err = conn.Write(sig)
		if err != nil {
			fmt.Println("write ping msg err:", err)
		}
	}
}
