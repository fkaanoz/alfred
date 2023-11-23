package server

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"path"
	"strings"
	"sync"
	"time"
)

const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request)
type Middleware func(h Handler) Handler

type muxEntry struct {
	path    string
	handler Handler
	method  string
}

type Config struct {
	Logger      *zap.SugaredLogger
	Timeout     time.Duration
	Middlewares []Middleware
}

type Server struct {
	mu          sync.RWMutex
	muxEntries  map[string]muxEntry
	logger      *zap.SugaredLogger
	timeout     time.Duration
	Middlewares []Middleware
}

func NewServer(conf Config) *Server {
	return &Server{
		muxEntries:  map[string]muxEntry{},
		logger:      conf.Logger,
		timeout:     conf.Timeout,
		Middlewares: conf.Middlewares,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	reqCtx := context.WithValue(ctx, ctxKey, Values{
		TrackID: uuid.New().String(),
	})

	p := r.URL.Path
	me, ok := s.muxEntries[p]
	if !ok {
		NotFound(reqCtx, w, r)
		return
	}

	if me.method != r.Method {
		NotAllowed(reqCtx, w, r)
		return
	}

	h := wrapMiddlewares(me.handler, s.Middlewares...)
	h(reqCtx, w, r)
}

func (s *Server) GET(path string, h Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	me := muxEntry{
		path:    cleanPath(path),
		method:  MethodGet,
		handler: h,
	}

	s.muxEntries[path] = me
}

func (s *Server) POST(path string, h Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	me := muxEntry{
		path:    cleanPath(path),
		method:  MethodPost,
		handler: h,
	}

	s.muxEntries[path] = me
}

func (s *Server) PUT(path string, h Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	me := muxEntry{
		path:    cleanPath(path),
		method:  MethodPut,
		handler: h,
	}

	s.muxEntries[path] = me
}

func (s *Server) DELETE(path string, h Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	me := muxEntry{
		path:    cleanPath(path),
		method:  MethodDelete,
		handler: h,
	}

	s.muxEntries[path] = me
}

func cleanPath(p string) string {
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}

	p = path.Clean(p)

	if strings.HasSuffix(p, "/") {
		p = p[:(len(p) - 1)]
	}

	return p
}

func wrapMiddlewares(handler Handler, mv ...Middleware) Handler {
	for _, v := range mv {

		handler = v(handler)
	}

	return handler
}

func NotAllowed(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("method not allowed"))
}

func NotFound(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}
