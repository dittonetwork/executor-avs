package endpoint

import (
	"context"
	"errors"
	api "github.com/dittonetwork/executor-avs/api/operator"
	"github.com/dittonetwork/executor-avs/pkg/encoding/json"
	"github.com/dittonetwork/executor-avs/pkg/log"
	"github.com/mailru/easyjson"
	"net/http"
	"syscall"
)

type responder struct {
}

func (r responder) RespondError(ctx context.Context, w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, context.Canceled):
		w.WriteHeader(http.StatusRequestTimeout)
	default:
		log.WithContext(ctx).Error("internal server error", log.Err(err))

		w.WriteHeader(http.StatusInternalServerError)

		r.RespondEasyJson(ctx, w, api.ErrorMessageResponse{
			Message: "internal server error",
		})
	}
}

func (r responder) RespondJson(ctx context.Context, w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		r.logError(ctx, err)
	}
}

func (r responder) RespondEasyJson(ctx context.Context, w http.ResponseWriter, v easyjson.Marshaler) {
	w.Header().Set("Content-Type", "application/json")
	_, _, err := easyjson.MarshalToHTTPResponseWriter(v, w)
	if err != nil {
		r.logError(ctx, err)
	}
}

func (r responder) logError(ctx context.Context, err error) {
	if errors.Is(err, syscall.EPIPE) {
		log.WithContext(ctx).With(log.Err(err)).Warn("broken pipe")
	} else if errors.Is(err, syscall.ECONNRESET) {
		log.WithContext(ctx).With(log.Err(err)).Warn("connection reset by peer")
	} else {
		log.WithContext(ctx).With(log.Err(err)).Error("failed to encode json")
	}
}
