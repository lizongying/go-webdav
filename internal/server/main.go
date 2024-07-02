package server

import (
	"fmt"
	"github.com/lizongying/go-webdav/internal/config"
	"github.com/lizongying/go-webdav/internal/utils"
	"golang.org/x/net/webdav"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func createDir(dir string) {
	fileInfo, err := os.Stat(dir)

	if err == nil {
		if !fileInfo.IsDir() {
			if err = os.MkdirAll(dir, 0755); err != nil {
				log.Panicln(err)
			}
		}
	} else if os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			log.Panicln(err)
		}
	} else {
		log.Panicln(err)
	}
}

type Server struct {
	addr      string
	scheme    string
	basicAuth string
	certFile  string
	keyFile   string
	mux       *http.ServeMux
}

func NewServer(config *config.Config) (s *Server, err error) {
	s = new(Server)

	u, err := url.Parse(config.Server.Host)
	if err != nil {
		log.Panicln(err)
	}
	s.addr = u.Host
	s.basicAuth = u.User.String()
	s.scheme = strings.ToLower(u.Scheme)
	if s.scheme == "https" {
		s.certFile = config.Server.Cert
		s.keyFile = config.Server.Key
	}
	lan := fmt.Sprintf("%s://***:***@%s:%s", s.scheme, utils.Lan(), u.Port())
	log.Println("lan", lan)
	s.mux = http.NewServeMux()

	for _, dir := range config.Dirs {
		paths := strings.Split(dir, ":")
		log.Println("route", fmt.Sprintf("%s%s", lan, paths[0]))
		createDir(paths[1])
		h := s.AuthMiddleware(&webdav.Handler{
			FileSystem: webdav.Dir(paths[1]),
			LockSystem: webdav.NewMemLS(),
			Prefix:     paths[0],
		})
		s.mux.Handle(paths[0], h)
	}
	return
}

func (s *Server) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if s.basicAuth != "" {
			username, password, ok := r.BasicAuth()
			if !ok {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if s.basicAuth != fmt.Sprintf("%s:%s", username, password) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) Run() (err error) {
	if s.scheme == "https" {
		err = http.ListenAndServeTLS(s.addr, s.certFile, s.keyFile, s.mux)
	} else {
		err = http.ListenAndServe(s.addr, s.mux)
	}

	return
}
