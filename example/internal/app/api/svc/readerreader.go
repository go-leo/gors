package svc

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/go-leo/gors"
)

type ReaderReaderService struct {
}

func (svc *ReaderReaderService) GetReaderReader(ctx context.Context, reader io.Reader) (io.Reader, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, gors.NewHttpError(http.StatusBadRequest, err)
	}
	return bytes.NewReader(append([]byte("hello "), data...)), nil
}

func (svc *ReaderReaderService) HeadReaderReader(ctx context.Context, reader io.Reader) (io.Reader, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, gors.NewHttpError(http.StatusBadRequest, err)
	}
	return bytes.NewReader(append([]byte("hello "), data...)), nil
}
