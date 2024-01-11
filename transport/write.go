/*
Copyright © 2024 JOSEPH INNES <avianpneuma@gmail.com>
*/
package transport

import (
	"errors"
	"fmt"
	"net"

	"github.com/swissinfo-ch/logd/auth"
	"github.com/swissinfo-ch/logd/cmd"
	"google.golang.org/protobuf/proto"
)

func (t *Transporter) handleWrite(c *cmd.Cmd, raddr *net.UDPAddr, sum, timeBytes, payload []byte) error {
	if c.Msg == nil {
		return errors.New("msg is nil")
	}
	valid, err := auth.Verify(t.writeSecret, sum, timeBytes, payload)
	if !valid || err != nil {
		return fmt.Errorf("%s unauthorised to write: %w", raddr.IP.String(), err)
	}
	// marshal msg
	msgBytes, err := proto.Marshal(c.Msg)
	if err != nil {
		return fmt.Errorf("protobuf marshal msg err: %w", err)
	}
	// write to buffer
	t.buf.Write(&msgBytes)
	// pipe to tails
	t.Out <- &ProtoPair{
		Msg:   c.Msg,
		Bytes: &msgBytes,
	}
	// pipe to alarm svc
	t.alarmSvc.In <- c.Msg
	return nil
}
