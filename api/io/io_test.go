package io_test

import (
	"fmt"
	"sea/api/io"
	"testing"
	"time"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_GenerateKeyID(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		startTime := time.Now()
		cr := io.ConnectionRoute{}
		cr.Init().SetCurrentHash().SetReceiverHash().SetCurrentHash().AddRelayHash().AddRelayHash()
		// key, err := service.GenerateKey()
		// gtest.Assert(err, nil)
		// genTime := time.Now()
		// var s string
		// for i := 0; i < 100000; i++ {
		// 	s1, err := service.GetKeyID(&key.PublicKey)
		// 	t.Assert(err, nil)
		// 	if len(s) != 0 {
		// 		t.Assert(s, s1)
		// 	}
		// 	s = s1
		// }
		endTime := time.Now()
		fmt.Printf("Total: %s\n", endTime.Sub(startTime))
		fmt.Printf("cr: %+v\n", cr)
	})
}
