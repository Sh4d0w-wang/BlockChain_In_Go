package utils

import (
    "bytes"
    "encoding/binary"
    "log"
)

// IntToHex convert int to hex bytes
func IntToHex(num int64) []byte {
    buff := new(bytes.Buffer)
    err := binary.Write(buff, binary.BigEndian, num)
    if err != nil {
        log.Panic(err)
    }
    return buff.Bytes()
}
