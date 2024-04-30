package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/dittonetwork/executor-avs/pkg/log"
)

// Recover catches panic and sends a response with status 500.
// Logs error with runtime stack.
func Recover(
	logger log.Logger,
	logParamsFn func(logger log.Logger, r *http.Request) []log.Field,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tee := NewTeeReader(r.Body)
			r.Body = tee

			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)

					_ = tee.Close()
					r.Body = io.NopCloser(bytes.NewBuffer(tee.Tee()))
					logParams := logParamsFn(logger, r)
					logParams = append(logParams, log.String("error", fmt.Sprintf("%v", err)))

					logger.With(log.Stack("stacktrace")).With(logParams...).Error("panic with stacktrace")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

type TeeReader struct {
	io.ReadCloser
	buffer [][]byte
}

func NewTeeReader(r io.ReadCloser) *TeeReader {
	return &TeeReader{r, nil}
}

func (r *TeeReader) Read(p []byte) (int, error) {
	n, err := r.ReadCloser.Read(p)
	if err != nil {
		return n, fmt.Errorf("read: %w", err)
	}

	if n == len(p) {
		r.buffer = append(r.buffer, p)
	} else if n > 0 {
		r.buffer = append(r.buffer, p[:n])
	}

	return n, nil
}

func (r *TeeReader) Close() error {
	return r.ReadCloser.Close()
}

func (r *TeeReader) Tee() []byte {
	bs := make([]byte, 0)
	for _, b := range r.buffer {
		bs = append(bs, b...)
	}
	return bs
}

func LogPanicRequest(logger log.Logger, r *http.Request) []log.Field {
	requestBodyBytes := copyBody(r, logger)

	return []log.Field{
		log.String("request_body", string(requestBodyBytes)),
		log.String("request_query", r.RequestURI),
	}
}
