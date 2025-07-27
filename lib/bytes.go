package lib

import "encoding/binary"

func Uint64ToBytes(u uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, u)
	return b
}

func Int64ToBytes(i int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func Uint32ToBytes(u uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, u)
	return b
}

func Int32ToBytes(i int32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(i))
	return b
}
