package io_test

import (
	"fmt"
	"sea/api/io"
	"testing"
	"time"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/grand"
)

func Test_ConnectionRoute(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		startTime := time.Now()
		cr := io.ConnectionRoute{}
		var current [256]byte
		var sender [256]byte
		var receiver [256]byte
		var r1 [256]byte
		var r2 [256]byte
		// var r3 [256]byte
		copy(current[:], grand.B(256))
		copy(sender[:], grand.B(256))
		copy(receiver[:], grand.B(256))
		copy(r1[:], grand.B(256))
		copy(r2[:], grand.B(256))
		// copy(r3[:], grand.B(256))
		cr.Init().SetCurrentHash(current).SetSenderHash(sender).SetReceiverHash(receiver).AddRelayHash(r1).AddRelayHash(r2)
		t.Assert(cr.CurrentHash, current)
		t.Assert(cr.SenderHash, sender)
		t.Assert(cr.ReceiverHash, receiver)
		t.Assert(cr.RelayHash[0], r1)
		t.Assert(cr.RelayHash[1], r2)
		t.Assert(cr.IsValid(), false)
		cr.AddRelayHash(current)
		t.Assert(cr.IsValid(), true)
		endTime := time.Now()
		fmt.Printf("Total: %s\n", endTime.Sub(startTime))
		fmt.Printf("cr: %+v\n", cr)
	})
}
