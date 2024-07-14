package redirect

import (
	resp "GoURLShortener/internal/lib/api/response"
	"GoURLShortener/internal/lib/logger/sl"
	"GoURLShortener/internal/storage"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type URLGetter interface {
	GetURL(alias string) (string, error)
}

func New(log *slog.Logger, urlGetter URLGetter) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "handlers.url.redirect.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(request.Context())))

		alias := chi.URLParam(request, "alias")
		if alias == "" {
			log.Info("alias is empty")

			render.JSON(writer, request, resp.Error("invalid request"))
		}

		resURL, err := urlGetter.GetURL(alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("url not found", "alias", alias)

			render.JSON(writer, request, resp.Error("not found"))
			return
		}
		if err != nil {
			log.Error("failed to get url", sl.Err(err))

			render.JSON(writer, request, resp.Error("internal error"))
			return
		}

		log.Info("got url", slog.String("url", resURL))

		// redirect to found url
		// 301: redirect permanently (saved in cache)
		// 302: (client does not cache)
		http.Redirect(writer, request, resURL, http.StatusFound)
	}
}
