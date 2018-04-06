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
		0xcd, 0x6e, 0xdb, 0x3c, 0x10, 0xbc, 0xeb, 0x29, 0xf6, 0x23, 0x7c, 0xfc,
		0x2c, 0x5f, 0x8b, 0x40, 0xf2, 0xa1, 0xc9, 0xb1, 0x28, 0x02, 0x34, 0x7d,
		0x80, 0x35, 0xb9, 0x96, 0x08, 0xd1, 0xa4, 0x40, 0xae, 0x0a, 0xa8, 0x82,
		0xde, 0xbd, 0xa0, 0x25, 0x25, 0xfa, 0x71, 0x9c, 0x14, 0x3d, 0xc9, 0x1c,
		0xce, 0x0e, 0x39, 0xb3, 0x0b, 0xba, 0xeb, 0x40, 0xd1, 0x59, 0x5b, 0x02,
		0x61, 0x08, 0x15, 0xf9, 0x93, 0x43, 0xaf, 0x04, 0xf4, 0x7d, 0x92, 0xfd,
		0xa7, 0x9c, 0xe4, 0xb6, 0x26, 0x28, 0xf9, 0x62, 0x8e, 0x49, 0x16, 0x3f,
		0x60, 0xd0, 0x16, 0xb9, 0x20, 0x2b, 0x22, 0x40, 0xa8, 0x8e, 0x09, 0x00,
		0x40, 0x76, 0x21, 0x46, 0x90, 0x25, 0xfa, 0x40, 0x9c, 0x8b, 0x86, 0xcf,
		0xfb, 0x2f, 0x62, 0xdc, 0x62, 0xcd, 0x86, 0x8e, 0x4f, 0xce, 0xa2, 0x51,
		0x35, 0x1a, 0xe7, 0x7e, 0x63, 0x76, 0x18, 0xc0, 0x59, 0xad, 0xc5, 0x0b,
		0xe5, 0x42, 0x51, 0x90, 0x5e, 0xd7, 0xac, 0x9d, 0x15, 0x20, 0x9d, 0x65,
		0xb2, 0x9c, 0x0b, 0x04, 0xf6, 0xa8, 0x74, 0x44, 0xa1, 0xb1, 0x46, 0x57,
		0x04, 0x68, 0x5b, 0x70, 0x5c, 0x92, 0x17, 0x5b, 0x11, 0x6c, 0xb8, 0x74,
		0x7e, 0x56, 0x1f, 0xf0, 0x72, 0x69, 0x41, 0x39, 0x3b, 0x91, 0x8d, 0xb6,
		0x15, 0x78, 0x32, 0xb9, 0x08, 0xdc, 0x1a, 0x0a, 0x25, 0x11, 0x0b, 0x28,
		0x3d, 0x9d, 0x73, 0x71, 0x18, 0xa0, 0x54, 0x86, 0x10, 0x3d, 0x1e, 0x06,
		0x93, 0xd9, 0xc9, 0xa9, 0xf6, 0x98, 0x64, 0x4a, 0xff, 0x02, 0xad, 0x72,
		0x11, 0xb5, 0x51, 0xdb, 0xb7, 0xf3, 0xa7, 0x8d, 0xf2, 0x9a, 0xe2, 0x88,
		0x4e, 0x3b, 0x6f, 0xab, 0x57, 0xe4, 0x9e, 0xa7, 0xec, 0xb0, 0xa8, 0x19,
		0x97, 0x1f, 0x29, 0x1a, 0x0c, 0x0c, 0x4d, 0xad, 0x90, 0x49, 0xad, 0x14,
		0x5e, 0x39, 0x5d, 0x97, 0x7e, 0xc3, 0xc0, 0x3f, 0x07, 0x56, 0xdf, 0xdf,
		0x3e, 0x69, 0xfd, 0x73, 0xf2, 0xe6, 0xd1, 0x56, 0xda, 0x16, 0x61, 0x74,
		0xd7, 0x75, 0x1e, 0x6d, 0x41, 0xb0, 0xd3, 0xff, 0xef, 0x6a, 0x78, 0xc8,
		0x21, 0x7d, 0x8c, 0x89, 0x07, 0x46, 0xcb, 0xa1, 0xef, 0x17, 0xf7, 0x05,
		0x69, 0x30, 0x84, 0x41, 0x62, 0x2f, 0xe3, 0x88, 0x6d, 0x6f, 0xb7, 0xe1,
		0xec, 0xaf, 0x53, 0xb2, 0x62, 0xde, 0x4e, 0x60, 0xe6, 0x70, 0x57, 0xa7,
		0x4f, 0x3a, 0xd4, 0x06, 0xdb, 0x67, 0x17, 0xae, 0x01, 0x6f, 0x7c, 0x6e,
		0x2b, 0xbe, 0xe3, 0x85, 0xde, 0xa7, 0xdd, 0xa9, 0xde, 0xa0, 0x4b, 0xdd,
		0x1f, 0xd2, 0xf9, 0xbf, 0x12, 0x5e, 0x77, 0xe4, 0x66, 0x2e, 0x71, 0xfa,
		0xbc, 0x3e, 0xcd, 0x92, 0x99, 0xf7, 0xa2, 0x88, 0xbd, 0xd8, 0xd5, 0xe9,
		0x63, 0xe9, 0xb4, 0xa4, 0x79, 0x27, 0x6e, 0x2a, 0x16, 0xde, 0x35, 0xf5,
		0x2a, 0xe4, 0x99, 0x1c, 0xec, 0x6a, 0x2d, 0xab, 0xab, 0x64, 0x91, 0x3e,
		0x6b, 0x59, 0xad, 0x05, 0xd7, 0xa2, 0x31, 0x79, 0xf2, 0xfb, 0x10, 0x9d,
		0xdf, 0xe8, 0xdd, 0x22, 0x1f, 0x2d, 0xab, 0xfb, 0xd1, 0x6f, 0xd8, 0xf7,
		0x03, 0xdd, 0xd0, 0x3f, 0x3d, 0x09, 0x5b, 0xb8, 0xeb, 0xc8, 0xaa, 0xf9,
		0x18, 0x2f, 0x29, 0x1f, 0x6c, 0xaf, 0x53, 0x61, 0x4d, 0xfb, 0x93, 0x27,
		0xac, 0x16, 0x8f, 0xc3, 0xe2, 0xce, 0xef, 0xbb, 0x79, 0xd1, 0x04, 0x63,
		0x31, 0x14, 0x0d, 0x85, 0xf0, 0xf0, 0x19, 0xff, 0xe9, 0x8b, 0xa6, 0xaf,
		0x43, 0xd5, 0xbf, 0x4c, 0xe0, 0xe6, 0x4d, 0x98, 0x9c, 0x8f, 0xcb, 0xf5,
		0x67, 0x7c, 0x29, 0x0f, 0xc3, 0xbf, 0xc6, 0x44, 0x4e, 0x92, 0xe4, 0x4f,
		0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0x1a, 0x26, 0x6d, 0x6c, 0x06, 0x00,
		0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}
