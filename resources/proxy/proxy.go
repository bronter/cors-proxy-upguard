package proxy

import (
  "github.com/labstack/echo"
  "github.com/parnurzeal/gorequest"
  "io"
  "io/ioutil"
)

func Proxy(c echo.Context) error {

  r := c.Request()

  request := gorequest.New().CustomMethod(r.Method(), "https://app.upguard.com/" + r.URL().Path())

  h := r.Header()
  for _, header := range h.Keys() {
    request = request.Set(header, h.Get(header))
  }

  q := r.URL().QueryString()
  request = request.Query(q)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body(), 1048576))

  if (err != nil) {
    return err;
  }

  request.SendString(string(body[:]))

  response, rbody, err2 := request.End()

  if (err2 != nil) {
    return err;
  }

  c.String(response.StatusCode, rbody)

  return nil;
}
