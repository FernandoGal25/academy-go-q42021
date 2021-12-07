package controller

import "net/url"

// Is based on the Context struct of the Echo framework,
// Allows the Echo Context to be changed for a different
// implementation while mantaining the same code structure.
type Context interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Param(name string) string
	QueryParams() url.Values
}
