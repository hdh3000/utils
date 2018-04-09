package masters

import (
	"bytes"
	"compress/gzip"
	"io"
)

// files_leaderboard_tmpl_html returns raw, uncompressed file data.
func files_leaderboard_tmpl_html() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xa4, 0x55,
		0xcd, 0xae, 0x9b, 0x3c, 0x10, 0xdd, 0xf3, 0x14, 0xf3, 0x59, 0x59, 0x7c,
		0x57, 0x2a, 0xa4, 0xdd, 0x5d, 0x5d, 0x41, 0xa4, 0xb6, 0x77, 0x59, 0x55,
		0x57, 0x6a, 0xfa, 0x00, 0x13, 0x3c, 0x09, 0x16, 0xc6, 0xa6, 0xb6, 0xa9,
		0x44, 0x11, 0xef, 0x5e, 0x39, 0x40, 0xc2, 0x5f, 0xd2, 0x54, 0x5d, 0x91,
		0x19, 0x9f, 0x73, 0xcc, 0x39, 0x63, 0x9c, 0xa6, 0x01, 0x4e, 0x47, 0xa1,
		0x08, 0x98, 0x24, 0xe4, 0x64, 0x0e, 0x1a, 0x0d, 0x67, 0xd0, 0xb6, 0x41,
		0xfc, 0x1f, 0xd7, 0xa9, 0xab, 0x4b, 0x82, 0xcc, 0x15, 0x72, 0x17, 0xc4,
		0xfe, 0x01, 0x12, 0xd5, 0x29, 0x61, 0xa4, 0x98, 0x6f, 0x10, 0xf2, 0x5d,
		0x00, 0x00, 0x10, 0x17, 0xe4, 0x10, 0xd2, 0x0c, 0x8d, 0x25, 0x97, 0xb0,
		0xca, 0x1d, 0xc3, 0x67, 0xd6, 0x2f, 0x39, 0xe1, 0x24, 0xed, 0x5e, 0xb5,
		0x42, 0xc9, 0x4b, 0x94, 0x5a, 0xff, 0xc2, 0x78, 0xdb, 0x35, 0x47, 0x5c,
		0x85, 0x05, 0x25, 0x8c, 0x93, 0x4d, 0x8d, 0x28, 0x9d, 0xd0, 0x8a, 0x41,
		0xaa, 0x95, 0x23, 0xe5, 0x12, 0x86, 0xe0, 0x0c, 0x72, 0xe1, 0xbb, 0x50,
		0x29, 0x29, 0x72, 0x02, 0x54, 0x35, 0x68, 0x97, 0x91, 0x61, 0x4b, 0x11,
		0xac, 0x5c, 0xa6, 0xcd, 0x88, 0x6f, 0xb1, 0x28, 0x6a, 0xe0, 0x5a, 0x0d,
		0x60, 0x29, 0x54, 0x0e, 0x86, 0x64, 0xc2, 0xac, 0xab, 0x25, 0xd9, 0x8c,
		0xc8, 0x31, 0xc8, 0x0c, 0x1d, 0x13, 0xb6, 0xed, 0x5a, 0x51, 0x6a, 0xad,
		0xf7, 0xb8, 0xed, 0x4c, 0xc6, 0x07, 0xcd, 0xeb, 0x5d, 0x10, 0x73, 0xf1,
		0x13, 0x04, 0x4f, 0x98, 0xd7, 0x46, 0xa1, 0xae, 0xfb, 0x0f, 0x0b, 0xd9,
		0x39, 0xc5, 0xbe, 0x3b, 0xac, 0x5c, 0xab, 0x4b, 0xe7, 0x9e, 0xa7, 0x78,
		0x3b, 0xe1, 0xf4, 0xe5, 0x9f, 0x14, 0x25, 0x5a, 0x07, 0x55, 0xc9, 0xd1,
		0x11, 0x9f, 0x29, 0x5c, 0x30, 0x4d, 0x13, 0x7d, 0x41, 0xeb, 0xbe, 0x77,
		0xa8, 0xb6, 0x5d, 0xdf, 0x69, 0xfe, 0x73, 0xf0, 0x66, 0x50, 0xe5, 0x42,
		0x9d, 0x6c, 0xef, 0xae, 0x69, 0x0c, 0xaa, 0x13, 0xc1, 0x46, 0xbc, 0xdb,
		0x94, 0xf0, 0x92, 0x40, 0xf4, 0xd9, 0x27, 0x6e, 0x1d, 0x2a, 0x67, 0xdb,
		0x76, 0xf2, 0xbe, 0x90, 0x4a, 0xb4, 0xb6, 0x93, 0x08, 0x53, 0x7f, 0xc4,
		0x96, 0x6f, 0xb7, 0xc0, 0x84, 0xe7, 0x53, 0x32, 0x43, 0xae, 0x27, 0x30,
		0x72, 0xb8, 0x29, 0xa3, 0x57, 0x61, 0x4b, 0x89, 0xf5, 0x9b, 0xb6, 0xe7,
		0x80, 0x17, 0x3e, 0x97, 0x8c, 0xaf, 0x58, 0xd0, 0x6d, 0xd8, 0x1d, 0xf6,
		0xa2, 0x3b, 0xd5, 0xfd, 0x96, 0x6a, 0xf3, 0x57, 0xc2, 0xf3, 0x89, 0xac,
		0xe6, 0xe2, 0x4f, 0x9f, 0x11, 0x07, 0x76, 0x07, 0x77, 0x32, 0xba, 0x2a,
		0x47, 0x80, 0xc9, 0xb0, 0x44, 0x9a, 0xfb, 0x79, 0x6d, 0xca, 0xe8, 0xa3,
		0x94, 0x6f, 0x22, 0xcd, 0xc7, 0xe3, 0x5a, 0x93, 0xf4, 0x69, 0x92, 0x09,
		0xad, 0x77, 0xb3, 0x32, 0x8f, 0x89, 0x67, 0x91, 0xe6, 0xf7, 0xe3, 0x5c,
		0xa0, 0xef, 0x87, 0xb4, 0x80, 0x3f, 0x3a, 0xdd, 0x0b, 0xb1, 0x67, 0x8b,
		0x23, 0x68, 0x03, 0xff, 0xd3, 0x0f, 0xe8, 0x74, 0xf6, 0x99, 0xa9, 0xe0,
		0xc3, 0xf3, 0xd3, 0xbc, 0xf5, 0xfe, 0xa9, 0x6d, 0xc3, 0xa6, 0x21, 0x69,
		0xa9, 0x6d, 0x87, 0x4d, 0xfd, 0x8a, 0xaf, 0x48, 0x5d, 0x3e, 0x9b, 0xc7,
		0x86, 0xda, 0x53, 0xd6, 0xbe, 0xb1, 0x95, 0x72, 0x1e, 0xbc, 0x13, 0x14,
		0x1e, 0x0c, 0x61, 0x3e, 0xb9, 0x53, 0x26, 0xee, 0x6e, 0xfb, 0xde, 0x0b,
		0x82, 0x9e, 0x0c, 0xa7, 0x8a, 0xac, 0x7d, 0x79, 0x24, 0xe2, 0x68, 0x2f,
		0xe8, 0x53, 0xc7, 0xfa, 0x97, 0x83, 0x7b, 0x2d, 0xc7, 0x09, 0xf4, 0xed,
		0x1b, 0x8f, 0xfe, 0x92, 0xdd, 0x76, 0x7f, 0x38, 0x03, 0x2f, 0x08, 0x82,
		0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xda, 0x80, 0x5f, 0xbc, 0xa7, 0x06,
		0x00, 0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}