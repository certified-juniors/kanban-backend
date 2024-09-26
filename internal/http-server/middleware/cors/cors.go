package cors

import (
	"net/http"
	"strconv"
	"strings"
)

// CORSOptions настраивает параметры CORS
type CORSOptions struct {
	AllowedOrigins   []string // Список разрешенных Origin
	AllowedMethods   []string // Список разрешенных HTTP-методов
	AllowedHeaders   []string // Список разрешенных заголовков
	ExposedHeaders   []string // Список заголовков, которые могут быть доступны клиенту
	AllowCredentials bool     // Разрешить передачу куки и заголовков авторизации
	MaxAge           int      // Время жизни Preflight-запроса
}

// New создает новое middleware для обработки CORS
func New(opts CORSOptions) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			if origin == "" || !isOriginAllowed(origin, opts.AllowedOrigins) {
				next.ServeHTTP(w, r)
				return
			}

			// Устанавливаем основные заголовки CORS
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")

			if opts.AllowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			// Если это Preflight-запрос (OPTIONS)
			if r.Method == http.MethodOptions {
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(opts.AllowedMethods, ", "))
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(opts.AllowedHeaders, ", "))
				if opts.MaxAge > 0 {
					w.Header().Set("Access-Control-Max-Age", strconv.Itoa(opts.MaxAge))
				}
				w.WriteHeader(http.StatusNoContent)
				return
			}

			// Устанавливаем Exposed Headers, если они заданы
			if len(opts.ExposedHeaders) > 0 {
				w.Header().Set("Access-Control-Expose-Headers", strings.Join(opts.ExposedHeaders, ", "))
			}

			// Передаем управление следующему обработчику
			next.ServeHTTP(w, r)
		})
	}
}

// isOriginAllowed проверяет, разрешен ли запрашиваемый Origin
func isOriginAllowed(origin string, allowedOrigins []string) bool {
	for _, o := range allowedOrigins {
		if o == origin || o == "*" {
			return true
		}
	}
	return false
}
