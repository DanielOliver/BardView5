package bardlog

import (
	"github.com/docker/distribution/uuid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"regexp"
	"time"
)

const (
	KeyLogger    = "logger"
	KeyRequestId = "request_id"
	KeyLogType   = "type"

	HeaderXRequestId = "X-Request-ID"
)

func GetLogger(c *gin.Context) zerolog.Logger {
	if logger, exists := c.Get(KeyLogger); exists {
		return logger.(zerolog.Logger)
	}
	return log.Logger
}

func UseLoggingWithRequestId(logger zerolog.Logger,
	skipPath []string,
	skipPathRegexp *regexp.Regexp) func(c *gin.Context) {

	var skip map[string]struct{}
	if length := len(skipPath); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, path := range skipPath {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		requestId := uuid.Generate().String()
		c.Set(KeyRequestId, requestId)
		c.Header(HeaderXRequestId, requestId)
		c.Set(KeyLogger, logger.With().Str(KeyRequestId, requestId).Logger())

		c.Next()
		track := true

		if _, ok := skip[path]; ok {
			track = false
		}

		if track &&
			skipPathRegexp != nil &&
			skipPathRegexp.MatchString(path) {
			track = false
		}

		if track {
			end := time.Now().UTC()
			latency := end.Sub(start)
			msg := "Request"
			if len(c.Errors) > 0 {
				msg = c.Errors.String()
			}

			addAttributes := func(event *zerolog.Event) *zerolog.Event {
				return event.
					Int("status", c.Writer.Status()).
					Str("method", c.Request.Method).
					Str("path", c.Request.URL.Path).
					Str("path_common", c.FullPath()).
					Str("ip", c.ClientIP()).
					Dur("latency", latency).
					Str("user_agent", c.Request.UserAgent()).
					Str(KeyLogType, "Request").
					Str(KeyRequestId, requestId).
					Int64("request_content_length", c.Request.ContentLength).
					Int("response_content_length", c.Writer.Size())

			}

			switch {
			case c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError:
				{
					addAttributes(logger.Warn()).
						Msg(msg)
				}
			case c.Writer.Status() >= http.StatusInternalServerError:
				{
					addAttributes(logger.Error()).
						Msg(msg)
				}
			default:
				addAttributes(logger.Info()).
					Msg(msg)
			}
		}
	}
}
