package udp

import (
	"fmt"
	"strings"

	"github.com/inneslabs/logd/cmd"
	"github.com/inneslabs/logd/pkg"
	"google.golang.org/protobuf/proto"
)

func (svc *UdpSvc) handleWrite(c *cmd.Cmd, p *pkg.Pkg) {
	if !svc.guard.Good([]byte(svc.secrets.Write), p) {
		return
	}
	msgBytes, err := proto.Marshal(c.Msg)
	if err != nil {
		return
	}
	segments := strings.Split(c.Msg.GetKey(), "/")
	if len(segments) < 3 {
		return
	}
	// IMPORTANT:
	// This is currently how msg keys are mapped to the rings
	storeKey := fmt.Sprintf("/%s/%s", segments[1], segments[2])
	svc.logStore.Write(storeKey, msgBytes)
	svc.forSubs <- &ProtoPair{
		Msg:   c.Msg,
		Bytes: msgBytes,
	}
}
