package io_test

import (
	"fmt"
	"sea/api/io"
	"sea/internal/consts"
	"testing"
	"time"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/grand"
)

func Test_ConnectionRoute(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		startTime := time.Now()
		cr := io.ConnectionRoute{}
		var current [consts.SERVER_ID_BYTES]byte
		var sender [consts.SERVER_ID_BYTES]byte
		var receiver [consts.SERVER_ID_BYTES]byte
		var r1 [consts.SERVER_ID_BYTES]byte
		var r2 [consts.SERVER_ID_BYTES]byte
		// var r3 [256]byte
		copy(current[:], grand.B(consts.SERVER_ID_BYTES))
		copy(sender[:], grand.B(consts.SERVER_ID_BYTES))
		copy(receiver[:], grand.B(consts.SERVER_ID_BYTES))
		copy(r1[:], grand.B(consts.SERVER_ID_BYTES))
		copy(r2[:], grand.B(consts.SERVER_ID_BYTES))
		// copy(r3[:], grand.B(256))
		cr.Init().SetCurrentID(current).SetSenderID(sender).SetReceiverID(receiver).AddRelayID(r1).AddRelayID(r2)
		t.Assert(cr.CurrentID, current)
		t.Assert(cr.SenderID, sender)
		t.Assert(cr.ReceiverID, receiver)
		t.Assert(cr.RelayID[0], r1)
		t.Assert(cr.RelayID[1], r2)
		t.Assert(cr.IsValid(), false)
		cr.AddRelayID(current)
		t.Assert(cr.IsValid(), true)
		endTime := time.Now()
		fmt.Printf("Total: %s\n", endTime.Sub(startTime))
		fmt.Printf("cr: %+v\n", cr)
	})
}
