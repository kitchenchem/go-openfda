package fda

import "net/http"

const (
  defaultBaseUrl = "https://api.fda.gov/"


)


type Client struct {
  client *http.Client
  BaseUrl string
  Authenticator Authenticator
}


