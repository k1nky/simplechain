package blockchain

import (
	"bytes"
	"encoding/binary"
)

func IntToBytes(i int64) []byte {
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.BigEndian, i); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
