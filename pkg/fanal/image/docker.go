package image

import (
	"bufio"
	"compress/gzip"
	"io"
	"os"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/tarball"
	"golang.org/x/xerrors"

	"github.com/khulnasoft/tunnel/pkg/fanal/utils"
)

func tryDockerArchive(fileName string) (v1.Image, error) {
	img, err := tarball.Image(fileOpener(fileName), nil)
	if err != nil {
		return nil, xerrors.Errorf("unable to open %s as a Docker image: %w", fileName, err)
	}
	return img, nil
}

func fileOpener(fileName string) func() (io.ReadCloser, error) {
	return func() (io.ReadCloser, error) {
		f, err := os.Open(fileName)
		if err != nil {
			return nil, xerrors.Errorf("unable to open the file: %w", err)
		}

		var r io.Reader
		br := bufio.NewReader(f)
		r = br

		if utils.IsGzip(br) {
			r, err = gzip.NewReader(br)
			if err != nil {
				_ = f.Close()
				return nil, xerrors.Errorf("failed to open gzip: %w", err)
			}
		}
		return io.NopCloser(r), nil
	}
}
