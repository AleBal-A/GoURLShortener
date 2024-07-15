package urldelete

import (
	resp "GoURLShortener/internal/lib/api/response"
	"GoURLShortener/internal/storage"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type URLDeleter interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, urlDeleter URLDeleter) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "handlers.url.delete.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(request.Context())),
		)

		alias := chi.URLParam(request, "alias")

		if alias == "" {
			log.Info("alias is empty")

			render.JSON(writer, request, resp.Error("invalid request"))
			return
		}

		err := urlDeleter.DeleteURL(alias)
		if err != nil {
			if errors.Is(err, storage.ErrURLNotFound) {
				log.Info("alias not found")
				render.JSON(writer, request, http.StatusNotFound)
				return
			}

			log.Error("failed to delete URL")
			render.JSON(writer, request, resp.Error("internal server error"))
			return
		}

		log.Info("successful delete URL by alias")
		// no JSON because sending the body is not recommended
		writer.WriteHeader(http.StatusNoContent)
	}
}
