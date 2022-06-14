package compressor

type CompressType int32

const (
	Raw CompressType = iota
	Gzip
	Snappy
	Zlib
)

type Compressor interface {
	Zip([]byte) ([]byte, error)
	Unzip([]byte) ([]byte, error)
}

// Compressors
var Compressors = map[CompressType]Compressor{}