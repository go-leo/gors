package svc

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/go-leo/gors"
)

type ReaderBytesService struct {
}

func (svc *ReaderBytesService) GetReaderBytes(ctx context.Context, reader io.Reader) ([]byte, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, gors.NewHttpError(http.StatusBadRequest, err)
	}
	return append([]byte("hello "), data...), nil
}

func (svc *ReaderBytesService) PostReaderBytes(ctx context.Context, reader io.Reader) ([]byte, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, gors.NewHttpError(http.StatusBadRequest, err)
	}
	return append([]byte("hello "), data...), nil
}
