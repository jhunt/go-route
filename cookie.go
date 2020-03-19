package route

import (
	"fmt"
	"net/http"
	"strings"
)

func sessionCookie(name string) string {
	return fmt.Sprintf("%s_s", strings.ToLower(name))
}

func sessionHeader(name string) string {
	return fmt.Sprintf("X-%s-Session", name)
}

func SessionID(req *http.Request, name string) string {
	if s := req.Header.Get(sessionHeader(name)); s != "" {
		return s
	}

	if c, err := req.Cookie(sessionCookie(name)); err == nil {
		return c.Value
	}

	return ""
}

func (r *Request) SessionID(name string) string {
	return SessionID(r.Req, name)
}

func (r *Request) SetCookie(name, val, path string) {
	http.SetCookie(r.w, &http.Cookie{
		Name:  name,
		Value: val,
		Path:  path,
	})
}

func (r *Request) ClearCookie(name, path string) {
	http.SetCookie(r.w, &http.Cookie{
		Name:   name,
		Path:   path,
		MaxAge: 0,
	})
}

func (r *Request) SetSession(name, id string) {
	r.SetCookie(sessionCookie(name), id, "/")
	r.w.Header().Set(sessionHeader(name), id)
}

func (r *Request) ClearSession(name string) {
	r.ClearCookie(sessionCookie(name), "/")
}
