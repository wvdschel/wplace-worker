package wplace

import (
	"compress/flate"
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/klauspost/compress/zstd"
)

// zstdReadCloser wraps a zstd.Decoder to implement io.ReadCloser
type zstdReadCloser struct {
	decoder *zstd.Decoder
}

func (z *zstdReadCloser) Read(p []byte) (n int, err error) {
	return z.decoder.Read(p)
}

func (z *zstdReadCloser) Close() error {
	// zstd.Decoder doesn't have a Close method that returns an error
	// but we need to implement the io.ReadCloser interface
	z.decoder.Close()
	return nil
}

type CloseFunc func() error

func wrapClose(rc io.ReadCloser, close CloseFunc) CloseFunc {
	return func() error {
		if err := rc.Close(); err != nil {
			return err
		}

		return close()
	}
}

type wrappedReadCloser struct {
	io.Reader

	close CloseFunc
}

func (w *wrappedReadCloser) Close() error {
	return w.close()
}

func respBody(resp *http.Response) io.ReadCloser {
	encoding := resp.Header.Get("Content-Encoding")
	if encoding == "" {
		return resp.Body
	}

	// Handle nested compression by processing from right to left
	encodings := strings.Split(encoding, ",")
	var reader io.Reader = resp.Body
	var close CloseFunc = resp.Body.Close

	// Process encodings in reverse order (right to left) to handle nested compression
	for i := len(encodings) - 1; i >= 0; i-- {
		encoding := strings.TrimSpace(encodings[i])
		switch strings.ToLower(encoding) {
		case "gzip":
			gzReader, err := gzip.NewReader(reader)
			if err != nil {
				// If we can't create the gzip reader, return the original body
				return resp.Body
			}
			close = wrapClose(gzReader, close)
			reader = gzReader
		case "deflate":
			flateReader := flate.NewReader(reader)
			close = wrapClose(flateReader, close)
			reader = flateReader
		case "br", "brotli":
			// Brotli is not supported in the standard library or klauspost/compress
			// For now, treat as uncompressed if we don't know how to handle it
			reader = brotli.NewReader(reader)
		case "zstd":
			zstdReader, err := zstd.NewReader(reader)
			if err != nil {
				// If we can't create the zstd reader, return the original body
				return resp.Body
			}
			zr := &zstdReadCloser{decoder: zstdReader}
			close = wrapClose(zr, close)
			reader = zr
		default:
			// Unknown encoding, treat as uncompressed
			continue
		}
	}

	return &wrappedReadCloser{
		Reader: reader,
		close:  close,
	}
}
