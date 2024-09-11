package cmd

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"testing"
	"time"
)

//"v1:eth:block:merkle:start=56753453:count=500:end_timestamp=223423423"
//"v1:eth:block:merkle:start=56753453:count=500:difficulty=230"

func performProofOfWork(input [32]byte) [32]byte {
	difficulty := int64(230)
	target := new(big.Int).Exp(big.NewInt(2), big.NewInt(difficulty), nil) // 调整难度
	var hash [32]byte
	nonce := 0

	for {
		data := append(input[:], []byte(fmt.Sprintf("%d", nonce))...)
		hash = sha256.Sum256(data)
		if new(big.Int).SetBytes(hash[:]).Cmp(target) == -1 {
			break
		}
		nonce++
	}
	return hash
}

func TestPerformProofOfWork(t *testing.T) {
	// block hash
	input := [32]byte{11, 22, 30, 44, 50, 66, 77, 81, 91, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210, 220, 230, 240, 250, 1, 2, 3, 4, 5, 6}

	// 开始计时
	start := time.Now()

	// 运行你的函数
	performProofOfWork(input)

	// 计算执行时间
	elapsed := time.Since(start)

	t.Logf("performProofOfWork took %s", elapsed)
}
