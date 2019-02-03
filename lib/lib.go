package lib

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/FreshworksStudio/bs-go-utils/apiEntity"
)

/**
 * Generic helpers and utility functions
**/

// Respond - Responds with a json response
func Respond(res http.ResponseWriter, obj interface{}) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(obj)
	res.Write([]byte("\n"))
}

// Dump and marshall an object
func Dump(obj interface{}) {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		log.Printf(string(data))
	}
}

// LoggingHandler prints timing and requests bodies of incoming api requests
func LoggingHandler(next http.Handler) http.Handler {
	var logger = log.New(os.Stderr, "", log.LstdFlags)

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		loggingWriter := &LoggingResponseWriter{res, http.StatusOK}

		startTime := time.Now()
		next.ServeHTTP(loggingWriter, req)
		elapsedTime := time.Now().Sub(startTime)

		logger.Printf(
			`"%s %s %s %d %d" %f`,
			req.Method, req.RequestURI, req.Proto,
			loggingWriter.statusCode, 0, elapsedTime.Seconds(),
		)

	})
}

// LoggingResponseWriter - Middleware Logging Response
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader adds status code to response
func (res *LoggingResponseWriter) WriteHeader(code int) {
	res.statusCode = code
	res.ResponseWriter.WriteHeader(code)
}

// Abs - Returns the Absolute Value of an integer
func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// Distance - manhatten distance
func Distance(c1 apiEntity.Coord, c2 apiEntity.Coord) int {
	return Abs(c1.X-c2.X) + Abs(c1.Y-c2.Y)
}

// DirectionFromCoords - given c1 and c2 return direction between them e.g. up
func DirectionFromCoords(c1 apiEntity.Coord, c2 apiEntity.Coord) string {
	vertical := c1.Y - c2.Y
	horizontal := c1.X - c2.X
	if vertical == 0 {
		if horizontal > 0 {
			return apiEntity.Right
		}
		return apiEntity.Left
	}
	if vertical < 0 {
		return apiEntity.Up
	}
	return apiEntity.Down
}
