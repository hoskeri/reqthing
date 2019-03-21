package reqthing

import "log"
import "net/http"

var Info log.Logger
var Debug log.Logger

type Reqthing struct {
	tr         *http.Transport
	name       string
	cacheDir   string
	hashSecret []byte
}

type Options struct {
	CacheDir   string
	Name       string
	HashSecret string
}

// get cache path for a url
func (r *Reqthing) cacheId(url string) (string, error) {
	// normalize the url.
	// calculate hmac
	// filepath/join to the cache dir
	return "", nil
}

// create a new Reqthing
func New(options Options) *Reqthing {
	return &Reqthing{}
}

// perform an http request
func (r *Reqthing) Roundtrip(req *http.Request) (*http.Response, error) {
	return nil, nil
}
