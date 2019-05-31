package reqthing

import stdlog "log"
import "net/http"

var Log stdlog.Logger

// Assert that Reqthing is a roundtripper
var _ http.RoundTripper = new(Reqthing)

type Reqthing struct {
	tr       http.RoundTripper
	name     string
	cacheDir string
	hSeed    []byte
}

type Options struct {
	CacheDir string // cache storage location
	HashSeed []byte // hash seed for cache keys
}

// get cache path for a url
func (r *Reqthing) cacheId(url string) (string, error) {
	// normalize the url.
	// calculate hmac
	// filepath/join to the cache dir
	return "", nil
}

type key struct {
	part1 uint64
	part2 uint64
}

// create a new Reqthing
func New(o Options) *Reqthing {
	return &Reqthing{
		tr:       http.DefaultTransport,
		cacheDir: o.CacheDir,
		hSeed:    []byte(o.HashSeed),
	}
}

// cache key for request.
func getKey(req *http.Request) key {
	return key{}
}

// perform an http request
func (r *Reqthing) RoundTrip(req *http.Request) (*http.Response, error) {
	Log.Printf("key: %v", getKey(req))
	return r.tr.RoundTrip(req)
}
