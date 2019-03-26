package reqthing

import "log"
import "net/http"

var Info log.Logger
var Debug log.Logger

// Assert that Reqthing is a roundtripper
var _ http.RoundTripper = new(Reqthing)

type Reqthing struct {
	tr       http.RoundTripper
	name     string
	cacheDir string
	hSeed    []byte
}

type Options struct {
	Name     string
	CacheDir string
	HashSeed string
}

// get cache path for a url
func (r *Reqthing) cacheId(url string) (string, error) {
	// normalize the url.
	// calculate hmac
	// filepath/join to the cache dir
	return "", nil
}

// create a new Reqthing
func New(o Options) *Reqthing {
	return &Reqthing{
		tr:       http.DefaultTransport,
		name:     o.Name,
		cacheDir: o.CacheDir,
		hSeed:    []byte(o.HashSeed),
	}
}

// perform an http request
func (r *Reqthing) RoundTrip(req *http.Request) (*http.Response, error) {
	return r.tr.RoundTrip(req)
}
