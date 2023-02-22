package urlservice

import (
	"encoding/base64"
	"encoding/binary"
	"hash/crc64"
)

// Shortener shortens a URL.
type Shortener interface {
	Shorten(string) string
}

// NewCRC64Shortener creates a new shortener that uses CRC64 as a hash function.
func NewCRC64Shortener(poly ...uint64) Shortener {
	p := uint64(crc64.ECMA)
	if len(poly) > 0 {
		p = poly[0]
	}
	return &crc64Shortener{
		table: crc64.MakeTable(p),
	}
}

var _ Shortener = (*crc64Shortener)(nil)

type crc64Shortener struct {
	table *crc64.Table
}

func (s *crc64Shortener) Shorten(url string) string {
	hash := crc64.Checksum([]byte(url), s.table)

	var b [8]byte
	binary.LittleEndian.PutUint64(b[:], hash)

	return encode(b[:])
}

func encode(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}
