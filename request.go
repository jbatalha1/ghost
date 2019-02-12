package ghost

import (
	"encoding/json"
	"github.com/go-resty/resty"
	"math/rand"
	"time"
)

type RestResponse struct {
	Raw    *resty.Response
	String string
}

var Unmarshal = func(payload string, v interface{}) error {
	return json.Unmarshal([]byte(payload), &v)
}

func (g Ghost) MakeRequest() (r RestResponse, err error) {
	return g.request()
}

func (g Ghost) request() (r RestResponse, err error) {

	resp := resty.SetRetryCount(g.Rest.RetryCount).
		SetRetryWaitTime(g.Rest.RetryWaitTime).
		SetRetryMaxWaitTime(g.Rest.RetryMaxWaitTime)

	if g.Rest.RedirectPolicy != nil {
		resp.SetRedirectPolicy(g.Rest.RedirectPolicy)
	}

	if g.Rest.ContentType != "" {
		resp.SetHeader("Content-Type", g.Rest.ContentType)
	}

	if g.Rest.Proxy != nil {
		switch v := g.Rest.Proxy.(type) {
		case *Proxy:
			resp.SetProxy(g.Rest.Proxy.(*Proxy).Ip)
		case *ProxyGateway:
			rand.Seed(time.Now().Unix())
			n := rand.Int() % len(g.Rest.Proxy.(*ProxyGateway).Ip)
			if g.Rest.Proxy.(*ProxyGateway).Random {
				resp.SetProxy(g.Rest.Proxy.(*ProxyGateway).Ip[n])
			}
		default:
			console.Println("[✖] Unexpected Auth type, why?", v)
			return r, err
		}
	}

	switch v := g.Rest.Auth.(type) {
	case *Token:
		resp.SetHeader(g.Rest.Auth.(*Token).Value, g.Rest.Auth.(*Token).Value)
	case *BasicAuth:
		resp.SetBasicAuth(g.Rest.Auth.(*BasicAuth).Username, g.Rest.Auth.(*BasicAuth).Password)
	case *NoAuth:
	default:
		console.Println("[✖] Unexpected Auth type, why?", v)
		return r, err
	}

	if g.Rest.QueryParams != nil {
		resp.SetQueryParams(g.Rest.QueryParams)
	}

	if g.Rest.Type == "GET" {
		resp, err := resp.R().Get(g.Rest.URL)
		if err != nil {
			console.Println("[✖] Unexpected GET request, why?", err)
			return r, err
		}

		r.Raw = resp
		r.String = resp.String()
		return r, err

	} else {
		resp, err := resp.R().SetBody(g.Rest.Payload).Post(g.Rest.URL)
		if err != nil {
			console.Printf("[✖] Unexpected POST request, why?", err)
			return r, err
		}

		r.Raw = resp
		r.String = resp.String()
		return r, err

	}

}
