package reqthing

import "log"
import "net/url"
import "net/http"

var Info log.Logger
var Debug log.Logger

type Reqthing struct {
	name      string
	base      *url.URL
	tr        *http.Transport
	cachePath string
}
