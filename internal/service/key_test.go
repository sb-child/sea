package service_test

import (
	"fmt"
	"sea/internal/service"
	"testing"
	"time"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_GenerateKeyID(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		startTime := time.Now()
		key, err := service.GenerateKey()
		gtest.Assert(err, nil)
		genTime := time.Now()
		var s string
		for i := 0; i < 100000; i++ {
			s1, err := service.GetKeyID(&key.PublicKey)
			t.Assert(err, nil)
			if len(s) != 0 {
				t.Assert(s, s1)
			}
			s = s1
		}
		endTime := time.Now()
		fmt.Printf("GenerateKey: %s, GetKeyID: %s, Total: %s\n", genTime.Sub(startTime), endTime.Sub(genTime)/100000, endTime.Sub(startTime))
		fmt.Printf("KeyID: %s\n", s)
	})
}
