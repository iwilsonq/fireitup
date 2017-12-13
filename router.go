package main

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

// Handler - http handler type
type Handler func(*Client, interface{})

// Router - http router
type Router struct {
	rules   map[string]Handler
	session *r.Session
}

// NewRouter - factory for new router instance
func NewRouter(session *r.Session) *Router {
	return &Router{
		rules:   make(map[string]Handler),
		session: session,
	}
}
