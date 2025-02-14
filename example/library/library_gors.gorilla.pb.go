// Code generated by protoc-gen-gors-gorilla. DO NOT EDIT.

package library

import (
	context "context"
	fmt "fmt"
	v2 "github.com/go-leo/gors/v2"
	errorx "github.com/go-leo/gox/errorx"
	urlx "github.com/go-leo/gox/netx/urlx"
	mux "github.com/gorilla/mux"
	protojson "google.golang.org/protobuf/encoding/protojson"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http "net/http"
)

type LibraryServiceGorillaService interface {
	CreateShelf(ctx context.Context, request *CreateShelfRequest) (*Shelf, error)
	GetShelf(ctx context.Context, request *GetShelfRequest) (*Shelf, error)
	ListShelves(ctx context.Context, request *ListShelvesRequest) (*ListShelvesResponse, error)
	DeleteShelf(ctx context.Context, request *DeleteShelfRequest) (*emptypb.Empty, error)
	MergeShelves(ctx context.Context, request *MergeShelvesRequest) (*Shelf, error)
	CreateBook(ctx context.Context, request *CreateBookRequest) (*Book, error)
	GetBook(ctx context.Context, request *GetBookRequest) (*Book, error)
	ListBooks(ctx context.Context, request *ListBooksRequest) (*ListBooksResponse, error)
	DeleteBook(ctx context.Context, request *DeleteBookRequest) (*emptypb.Empty, error)
	UpdateBook(ctx context.Context, request *UpdateBookRequest) (*Book, error)
	MoveBook(ctx context.Context, request *MoveBookRequest) (*Book, error)
}

func AppendLibraryServiceGorillaRoute(router *mux.Router, service LibraryServiceGorillaService) *mux.Router {
	handler := LibraryServiceGorillaHandler{
		service: service,
		decoder: LibraryServiceGorillaRequestDecoder{
			unmarshalOptions: protojson.UnmarshalOptions{},
		},
		encoder: LibraryServiceGorillaResponseEncoder{
			marshalOptions:   protojson.MarshalOptions{},
			unmarshalOptions: protojson.UnmarshalOptions{},
		},
		errorEncoder: v2.DefaultErrorEncoder,
	}
	router.NewRoute().Name("/google.example.library.v1.LibraryService/CreateShelf").
		Methods("POST").
		Path("/v1/shelves").
		Handler(handler.CreateShelf())
	router.NewRoute().Name("/google.example.library.v1.LibraryService/GetShelf").
		Methods("GET").
		Path("/v1/shelves/{shelf}").
		Handler(handler.GetShelf())
	router.NewRoute().Name("/google.example.library.v1.LibraryService/ListShelves").
		Methods("GET").
		Path("/v1/shelves").
		Handler(handler.ListShelves())
	router.NewRoute().Name("/google.example.library.v1.LibraryService/DeleteShelf").
		Methods("DELETE").
		Path("/v1/shelves/{shelf}").
		Handler(handler.DeleteShelf())
	router.NewRoute().Name("/google.example.library.v1.LibraryService/MergeShelves").
		Methods("POST").
		Path("/v1/shelves/{shelf}:merge").
		Handler(handler.MergeShelves())
	router.NewRoute().Name("/google.example.library.v1.LibraryService/CreateBook").
		Methods("POST").
		Path("/v1/shelves/{shelf}/books").
		Handler(handler.CreateBook())
	router.NewRoute().Name("/google.example.library.v1.LibraryService/GetBook").
		Methods("GET").
		Path("/v1/shelves/{shelf}/books/{book}").
		Handler(handler.GetBook())
	router.NewRoute().Name("/google.example.library.v1.LibraryService/ListBooks").
		Methods("GET").
		Path("/v1/shelves/{shelf}/books").
		Handler(handler.ListBooks())
	router.NewRoute().Name("/google.example.library.v1.LibraryService/DeleteBook").
		Methods("DELETE").
		Path("/v1/shelves/{shelf}/books/{book}").
		Handler(handler.DeleteBook())
	router.NewRoute().Name("/google.example.library.v1.LibraryService/UpdateBook").
		Methods("PATCH").
		Path("/v1/shelves/{shelf}/books/{book}").
		Handler(handler.UpdateBook())
	router.NewRoute().Name("/google.example.library.v1.LibraryService/MoveBook").
		Methods("POST").
		Path("/v1/shelves/{shelf}/books/{book}:move").
		Handler(handler.MoveBook())
	return router
}

type LibraryServiceGorillaHandler struct {
	service      LibraryServiceGorillaService
	decoder      LibraryServiceGorillaRequestDecoder
	encoder      LibraryServiceGorillaResponseEncoder
	errorEncoder v2.ErrorEncoder
}

func (h LibraryServiceGorillaHandler) CreateShelf() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.CreateShelf(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.CreateShelf(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.CreateShelf(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h LibraryServiceGorillaHandler) GetShelf() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.GetShelf(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.GetShelf(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.GetShelf(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h LibraryServiceGorillaHandler) ListShelves() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.ListShelves(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.ListShelves(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.ListShelves(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h LibraryServiceGorillaHandler) DeleteShelf() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.DeleteShelf(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.DeleteShelf(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.DeleteShelf(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h LibraryServiceGorillaHandler) MergeShelves() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.MergeShelves(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.MergeShelves(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.MergeShelves(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h LibraryServiceGorillaHandler) CreateBook() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.CreateBook(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.CreateBook(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.CreateBook(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h LibraryServiceGorillaHandler) GetBook() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.GetBook(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.GetBook(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.GetBook(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h LibraryServiceGorillaHandler) ListBooks() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.ListBooks(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.ListBooks(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.ListBooks(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h LibraryServiceGorillaHandler) DeleteBook() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.DeleteBook(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.DeleteBook(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.DeleteBook(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h LibraryServiceGorillaHandler) UpdateBook() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.UpdateBook(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.UpdateBook(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.UpdateBook(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h LibraryServiceGorillaHandler) MoveBook() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.MoveBook(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.MoveBook(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.MoveBook(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

type LibraryServiceGorillaRequestDecoder struct {
	unmarshalOptions protojson.UnmarshalOptions
}

func (decoder LibraryServiceGorillaRequestDecoder) CreateShelf(ctx context.Context, r *http.Request) (*CreateShelfRequest, error) {
	req := &CreateShelfRequest{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err := decoder.unmarshalOptions.Unmarshal(data, req.Shelf); err != nil {
		return nil, err
	}
	return req, nil
}
func (decoder LibraryServiceGorillaRequestDecoder) GetShelf(ctx context.Context, r *http.Request) (*GetShelfRequest, error) {
	req := &GetShelfRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	req.Name = fmt.Sprintf("shelves/%s", vars.Get("shelf"))
	if varErr != nil {
		return nil, varErr
	}
	return req, nil
}
func (decoder LibraryServiceGorillaRequestDecoder) ListShelves(ctx context.Context, r *http.Request) (*ListShelvesRequest, error) {
	req := &ListShelvesRequest{}
	queries := r.URL.Query()
	var queryErr error
	req.PageSize, queryErr = errorx.Break[int32](queryErr)(urlx.GetInt[int32](queries, "page_size"))
	req.PageToken = queries.Get("page_token")
	if queryErr != nil {
		return nil, queryErr
	}
	return req, nil
}
func (decoder LibraryServiceGorillaRequestDecoder) DeleteShelf(ctx context.Context, r *http.Request) (*DeleteShelfRequest, error) {
	req := &DeleteShelfRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	req.Name = fmt.Sprintf("shelves/%s", vars.Get("shelf"))
	if varErr != nil {
		return nil, varErr
	}
	return req, nil
}
func (decoder LibraryServiceGorillaRequestDecoder) MergeShelves(ctx context.Context, r *http.Request) (*MergeShelvesRequest, error) {
	req := &MergeShelvesRequest{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err := decoder.unmarshalOptions.Unmarshal(data, req); err != nil {
		return nil, err
	}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	req.Name = fmt.Sprintf("shelves/%s", vars.Get("shelf"))
	if varErr != nil {
		return nil, varErr
	}
	return req, nil
}
func (decoder LibraryServiceGorillaRequestDecoder) CreateBook(ctx context.Context, r *http.Request) (*CreateBookRequest, error) {
	req := &CreateBookRequest{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err := decoder.unmarshalOptions.Unmarshal(data, req.Book); err != nil {
		return nil, err
	}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	req.Parent = fmt.Sprintf("shelves/%s", vars.Get("shelf"))
	if varErr != nil {
		return nil, varErr
	}
	return req, nil
}
func (decoder LibraryServiceGorillaRequestDecoder) GetBook(ctx context.Context, r *http.Request) (*GetBookRequest, error) {
	req := &GetBookRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	req.Name = fmt.Sprintf("shelves/%s/books/%s", vars.Get("shelf"), vars.Get("book"))
	if varErr != nil {
		return nil, varErr
	}
	return req, nil
}
func (decoder LibraryServiceGorillaRequestDecoder) ListBooks(ctx context.Context, r *http.Request) (*ListBooksRequest, error) {
	req := &ListBooksRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	req.Parent = fmt.Sprintf("shelves/%s", vars.Get("shelf"))
	if varErr != nil {
		return nil, varErr
	}
	queries := r.URL.Query()
	var queryErr error
	req.PageSize, queryErr = errorx.Break[int32](queryErr)(urlx.GetInt[int32](queries, "page_size"))
	req.PageToken = queries.Get("page_token")
	if queryErr != nil {
		return nil, queryErr
	}
	return req, nil
}
func (decoder LibraryServiceGorillaRequestDecoder) DeleteBook(ctx context.Context, r *http.Request) (*DeleteBookRequest, error) {
	req := &DeleteBookRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	req.Name = fmt.Sprintf("shelves/%s/books/%s", vars.Get("shelf"), vars.Get("book"))
	if varErr != nil {
		return nil, varErr
	}
	return req, nil
}
func (decoder LibraryServiceGorillaRequestDecoder) UpdateBook(ctx context.Context, r *http.Request) (*UpdateBookRequest, error) {
	req := &UpdateBookRequest{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err := decoder.unmarshalOptions.Unmarshal(data, req.Book); err != nil {
		return nil, err
	}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	if req.Book == nil {
		req.Book = &Book{}
	}
	req.Book.Name = fmt.Sprintf("shelves/%s/books/%s", vars.Get("shelf"), vars.Get("book"))
	if varErr != nil {
		return nil, varErr
	}
	return req, nil
}
func (decoder LibraryServiceGorillaRequestDecoder) MoveBook(ctx context.Context, r *http.Request) (*MoveBookRequest, error) {
	req := &MoveBookRequest{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err := decoder.unmarshalOptions.Unmarshal(data, req); err != nil {
		return nil, err
	}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	req.Name = fmt.Sprintf("shelves/%s/books/%s", vars.Get("shelf"), vars.Get("book"))
	if varErr != nil {
		return nil, varErr
	}
	return req, nil
}

type LibraryServiceGorillaResponseEncoder struct {
	marshalOptions   protojson.MarshalOptions
	unmarshalOptions protojson.UnmarshalOptions
}

func (encoder LibraryServiceGorillaResponseEncoder) CreateShelf(ctx context.Context, w http.ResponseWriter, resp *Shelf) error {
	return v2.ResponseEncoder(ctx, w, resp, encoder.marshalOptions)
}
func (encoder LibraryServiceGorillaResponseEncoder) GetShelf(ctx context.Context, w http.ResponseWriter, resp *Shelf) error {
	return v2.ResponseEncoder(ctx, w, resp, encoder.marshalOptions)
}
func (encoder LibraryServiceGorillaResponseEncoder) ListShelves(ctx context.Context, w http.ResponseWriter, resp *ListShelvesResponse) error {
	return v2.ResponseEncoder(ctx, w, resp, encoder.marshalOptions)
}
func (encoder LibraryServiceGorillaResponseEncoder) DeleteShelf(ctx context.Context, w http.ResponseWriter, resp *emptypb.Empty) error {
	return v2.ResponseEncoder(ctx, w, resp, encoder.marshalOptions)
}
func (encoder LibraryServiceGorillaResponseEncoder) MergeShelves(ctx context.Context, w http.ResponseWriter, resp *Shelf) error {
	return v2.ResponseEncoder(ctx, w, resp, encoder.marshalOptions)
}
func (encoder LibraryServiceGorillaResponseEncoder) CreateBook(ctx context.Context, w http.ResponseWriter, resp *Book) error {
	return v2.ResponseEncoder(ctx, w, resp, encoder.marshalOptions)
}
func (encoder LibraryServiceGorillaResponseEncoder) GetBook(ctx context.Context, w http.ResponseWriter, resp *Book) error {
	return v2.ResponseEncoder(ctx, w, resp, encoder.marshalOptions)
}
func (encoder LibraryServiceGorillaResponseEncoder) ListBooks(ctx context.Context, w http.ResponseWriter, resp *ListBooksResponse) error {
	return v2.ResponseEncoder(ctx, w, resp, encoder.marshalOptions)
}
func (encoder LibraryServiceGorillaResponseEncoder) DeleteBook(ctx context.Context, w http.ResponseWriter, resp *emptypb.Empty) error {
	return v2.ResponseEncoder(ctx, w, resp, encoder.marshalOptions)
}
func (encoder LibraryServiceGorillaResponseEncoder) UpdateBook(ctx context.Context, w http.ResponseWriter, resp *Book) error {
	return v2.ResponseEncoder(ctx, w, resp, encoder.marshalOptions)
}
func (encoder LibraryServiceGorillaResponseEncoder) MoveBook(ctx context.Context, w http.ResponseWriter, resp *Book) error {
	return v2.ResponseEncoder(ctx, w, resp, encoder.marshalOptions)
}
