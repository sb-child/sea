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
		s, err := service.GetKeyID(&key.PublicKey)
		t.Assert(err, nil)
		endTime := time.Now()
		fmt.Printf("GenerateKey: %s, GetKeyID: %s, Total: %s\n", genTime.Sub(startTime), endTime.Sub(genTime), endTime.Sub(startTime))
		fmt.Printf("KeyID: %s\n", s)
	})
}
