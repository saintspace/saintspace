package server

import (
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Server struct {
	logger           Logger
	webAppFileSystem fs.FS
}

func New(logger Logger, webAppFileSystem fs.FS) *Server {
	return &Server{
		logger:           logger,
		webAppFileSystem: webAppFileSystem,
	}
}

type Logger interface {
	InfoWithContext(message string, keysAndValues ...interface{})
	ErrorWithContext(message string, keysAndValues ...interface{})
	DebugWithContext(message string, keysAndValues ...interface{})
}

func (s *Server) Run() error {
	fsys, err := fs.Sub(s.webAppFileSystem, "dist")
	if err != nil {
		log.Fatal(err)
	}
	weAppFileSystemHandler := http.FileServer(http.FS(fsys))

	r := gin.Default()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/info", s.GetInstanceInfoHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		path := strings.TrimPrefix(c.Request.URL.Path, "/")
		if path == "" {
			path = "index.html"
		}
		_, err := fsys.Open(path)
		if err != nil {
			c.Request.URL.Path = "/"
		} else {
			// TODO: Cache that this path exists and doesn't not need to be checked again.
		}
		weAppFileSystemHandler.ServeHTTP(c.Writer, c.Request)
	})

	return r.Run(":3000") // Listen and serve on :3000
}
