package gosteam

import (
	"bytes"
	"fmt"
	"github.com/valyala/fasthttp"
)

func (s *session) getRedirect(req *fasthttp.Request, resp *fasthttp.Response, statuscode int, funcname string) error {
	for {
		if err := fasthttp.Do(req, resp); err != nil {
			return errorText(fmt.Sprintf("request error %s", funcname))
		}

		s.takeCookie(resp)

		switch resp.StatusCode() {
		case statuscode:
			return nil
		case 200:
			return nil
		case 302:
			location := resp.Header.Peek("Location")
			req.SetRequestURIBytes(location)
			s.cookieClient.FillRequest(req)
			continue
		default:
			return errorStatusCode(funcname, resp.StatusCode())
		}
	}
}

func (s *session) takeCookie(resp *fasthttp.Response) {
	resp.Header.VisitAllCookie(func(key, value []byte) {
		processingcookie := bytes.Split(value, []byte(";"))
		cookie := bytes.Split(processingcookie[0], []byte("="))
		s.cookieClient.SetBytesKV(cookie[0], cookie[1])
	})
}