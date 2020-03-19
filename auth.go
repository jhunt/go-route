package route

type BasicAuth struct {
	Username string
	Password string
}

func (r *Request) BasicAuth(auth BasicAuth) bool {
	u, p, ok := r.Req.BasicAuth()
	if !ok {
		r.Fail(Error{code: 401})
		return false
	}

	if u != auth.Username || p != auth.Password {
		r.Fail(Error{code: 403})
		return false
	}

	return true
}
