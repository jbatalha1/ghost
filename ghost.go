package ghost

import (
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"time"
)

var console = color.New(color.FgRed, color.Bold)
var spi = spinner.New(spinner.CharSets[11], 100*time.Millisecond)

type Ghost struct {
	Config map[string]interface{}
	Rest   *Rest
}

type Rest struct {
	Proxy            interface{}
	Auth             interface{}
	Payload          interface{}
	Response         string
	Type             string
	ContentType      string
	RetryCount       int
	RetryWaitTime    time.Duration
	RetryMaxWaitTime time.Duration
	URL              string
	QueryParams      map[string]string
	RedirectPolicy   interface{}
}

type Proxy struct {
	Ip string
}

type ProxyGateway struct {
	Ip     []string
	Random bool
}

type Token struct {
	Key   string
	Value string
}

type BasicAuth struct {
	Username string
	Password string
}

type NoAuth struct {
}

type DataModel struct {
	Data interface{}
}

/*func (g Ghost) New(auth interface{}, contentType string) *Ghost {
	g.Rest.Auth = auth
	g.Rest.ContentType = contentType
	return &g
}
*/
