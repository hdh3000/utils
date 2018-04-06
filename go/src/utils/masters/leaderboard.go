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
		0xc1, 0xae, 0x9b, 0x30, 0x10, 0xbc, 0xf3, 0x15, 0x5b, 0x2b, 0xc7, 0x86,
		0x5c, 0xab, 0x27, 0x88, 0xd4, 0xf6, 0x1d, 0xab, 0xea, 0x49, 0x7d, 0xfd,
		0x80, 0x0d, 0xde, 0x04, 0x0b, 0x63, 0x23, 0x7b, 0xa9, 0x44, 0x11, 0xff,
		0x5e, 0x39, 0x40, 0x02, 0x81, 0xa4, 0xa9, 0x7a, 0x22, 0x5e, 0xcf, 0x0c,
		0x9e, 0xd9, 0x8d, 0x69, 0x5b, 0x90, 0x74, 0x54, 0x86, 0x40, 0x68, 0x42,
		0x49, 0xee, 0x60, 0xd1, 0x49, 0x01, 0x5d, 0x17, 0x25, 0x1f, 0xa4, 0xcd,
		0xb8, 0xa9, 0x08, 0x72, 0x2e, 0xf5, 0x3e, 0x4a, 0xc2, 0x03, 0x34, 0x9a,
		0x53, 0x2a, 0xc8, 0x88, 0x50, 0x20, 0x94, 0xfb, 0x08, 0x00, 0x20, 0x29,
		0x89, 0x11, 0xb2, 0x1c, 0x9d, 0x27, 0x4e, 0x45, 0xcd, 0xc7, 0xed, 0x27,
		0x31, 0x6c, 0xb1, 0x62, 0x4d, 0xfb, 0x57, 0x6b, 0x50, 0xcb, 0x0a, 0xb5,
		0xb5, 0xbf, 0x31, 0xd9, 0xf5, 0xc5, 0x09, 0xd7, 0x60, 0x49, 0xa9, 0x90,
		0xe4, 0x33, 0xa7, 0x2a, 0x56, 0xd6, 0x08, 0xc8, 0xac, 0x61, 0x32, 0x9c,
		0x0a, 0x04, 0x76, 0x28, 0x55, 0xa8, 0x42, 0x6d, 0xb4, 0x2a, 0x08, 0xd0,
		0x34, 0x60, 0x39, 0x27, 0x27, 0x96, 0x22, 0x58, 0x73, 0x6e, 0xdd, 0x84,
		0xef, 0xb1, 0x2c, 0x1b, 0x90, 0xd6, 0x8c, 0x60, 0xad, 0x4c, 0x01, 0x8e,
		0x74, 0x2a, 0x3c, 0x37, 0x9a, 0x7c, 0x4e, 0xc4, 0x02, 0x72, 0x47, 0xc7,
		0x54, 0xec, 0xfa, 0x52, 0x9c, 0x79, 0x1f, 0x3c, 0xee, 0x7a, 0x93, 0xc9,
		0xc1, 0xca, 0x66, 0x1f, 0x25, 0x52, 0xfd, 0x02, 0x25, 0x53, 0x11, 0xb4,
		0x51, 0x99, 0xeb, 0xfb, 0xc7, 0x8d, 0xfc, 0x9c, 0xe2, 0x50, 0x1d, 0x77,
		0xae, 0xab, 0x4b, 0xe5, 0x91, 0xa7, 0x64, 0x37, 0xe3, 0x0c, 0xcb, 0xbf,
		0x29, 0x6a, 0xf4, 0x0c, 0x75, 0x25, 0x91, 0x49, 0xde, 0x28, 0x5c, 0x30,
		0x6d, 0x1b, 0x7f, 0x43, 0xcf, 0x3f, 0x7b, 0x54, 0xd7, 0xad, 0xbf, 0xe9,
		0xf6, 0xe7, 0xe8, 0xcd, 0xa1, 0x29, 0x94, 0x39, 0xf9, 0xc1, 0x5d, 0xdb,
		0x3a, 0x34, 0x27, 0x82, 0x8d, 0xfa, 0xb8, 0xa9, 0xe0, 0x25, 0x85, 0xf8,
		0x6b, 0x48, 0xdc, 0x33, 0x1a, 0xf6, 0x5d, 0x37, 0x3b, 0x2f, 0x64, 0x1a,
		0xbd, 0xef, 0x25, 0xb6, 0x59, 0x18, 0xb1, 0xe5, 0xe9, 0x16, 0x98, 0xed,
		0x79, 0x4a, 0x6e, 0x90, 0xeb, 0x09, 0x4c, 0x1c, 0x6e, 0xaa, 0xf8, 0x55,
		0xf9, 0x4a, 0x63, 0xf3, 0x66, 0xfd, 0x39, 0xe0, 0x85, 0xcf, 0x25, 0xe3,
		0x3b, 0x96, 0x74, 0x1f, 0xf6, 0x80, 0xbd, 0xa8, 0xce, 0x75, 0x7f, 0x64,
		0xd6, 0xfd, 0x93, 0xf0, 0x6d, 0x47, 0x56, 0x73, 0x09, 0xd3, 0xe7, 0xd4,
		0x41, 0x3c, 0xc0, 0x9d, 0x9c, 0xad, 0xab, 0x09, 0x60, 0xd6, 0x2c, 0x95,
		0x15, 0xa1, 0x5f, 0x9b, 0x2a, 0xfe, 0xac, 0xf5, 0x9b, 0xca, 0x8a, 0x69,
		0xbb, 0xd6, 0x24, 0x43, 0x9a, 0xe4, 0xb6, 0x3e, 0xb8, 0x59, 0xe9, 0xc7,
		0xcc, 0xb3, 0xca, 0x8a, 0xc7, 0x71, 0x2e, 0xd0, 0x8f, 0x43, 0x5a, 0xc0,
		0x9f, 0xed, 0xee, 0x85, 0x38, 0x63, 0xbf, 0xe7, 0xae, 0x1e, 0x29, 0xcf,
		0xb5, 0xa5, 0x6d, 0xc9, 0xc8, 0xe9, 0x3c, 0xdf, 0xfd, 0xd3, 0xac, 0x45,
		0xc7, 0x8a, 0xb6, 0x07, 0x47, 0x58, 0xcc, 0x6e, 0x85, 0xd9, 0xf9, 0xee,
		0x9f, 0xfc, 0x5d, 0x11, 0x0c, 0x64, 0x38, 0xd5, 0xe4, 0xfd, 0xcb, 0x33,
		0x21, 0xc5, 0xef, 0x8a, 0xbe, 0xf4, 0xac, 0xff, 0x19, 0xbd, 0xeb, 0x72,
		0x9a, 0xc0, 0x50, 0xbe, 0xf3, 0x18, 0xae, 0xc9, 0x5d, 0xff, 0xc9, 0x18,
		0x79, 0x51, 0x14, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xb3, 0x0f,
		0xdd, 0x69, 0x06, 0x00, 0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}
