package client

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/lizongying/go-webdav/internal/config"
	"github.com/lizongying/go-webdav/internal/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Multistatus struct {
	XMLName   xml.Name   `xml:"multistatus"`
	Responses []Response `xml:"response"`
}

type Response struct {
	Href     string `xml:"href"`
	Propstat struct {
		Prop struct {
			Getcontentlength string `xml:"getcontentlength"`
			Getcontenttype   string `xml:"getcontenttype"`
			Displayname      string `xml:"displayname"`
		} `xml:"prop"`
	} `xml:"propstat"`
}

type Node struct {
	mime  string
	url   string
	child []*Node
}

type Client struct {
	addr     string
	username string
	password string
	client   *http.Client
	dirs     []string
	lan      string
}

func NewClient(config *config.Config) (s *Client, err error) {
	s = new(Client)
	s.client = http.DefaultClient

	u, err := url.Parse(config.Server.Host)
	if err != nil {
		log.Panicln(err)
	}
	s.addr = u.Host
	s.username = u.User.Username()
	s.password, _ = u.User.Password()
	scheme := strings.ToLower(u.Scheme)
	s.lan = fmt.Sprintf("%s://%s:%s", scheme, utils.Lan(), u.Port())
	log.Println("lan", s.lan)

	for _, dir := range config.Dirs {
		paths := strings.Split(dir, ":")
		s.dirs = append(s.dirs, paths[0])
	}
	return
}

func (s *Client) List() (err error) {
	for _, dir := range s.dirs {
		log.Println(fmt.Sprintf("%s%s", s.lan, dir))
		req, err := http.NewRequest("PROPFIND", fmt.Sprintf("%s%s", s.lan, dir), nil)
		if err != nil {
			log.Println(err)
			continue
		}

		req.SetBasicAuth(s.username, s.password)

		resp, err := s.client.Do(req)
		if err != nil {
			log.Println(err)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		var multistatus Multistatus

		if err = xml.NewDecoder(bytes.NewReader(body)).Decode(&multistatus); err != nil {
			log.Println(err)
			continue
		}

		for _, r := range multistatus.Responses {
			if strings.Index(r.Propstat.Prop.Getcontenttype, "video/") == 0 || strings.Index(r.Propstat.Prop.Getcontenttype, "audio/") == 0 {
				log.Println(r.Propstat.Prop.Displayname, r.Propstat.Prop.Getcontenttype, r.Propstat.Prop.Getcontentlength, r.Href)
			}
		}

		_ = resp.Body.Close()
	}

	return
}
