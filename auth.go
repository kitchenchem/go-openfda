package fda

import "net/http"

type Authenticator interface {
  Add(req *http.Request)
}
