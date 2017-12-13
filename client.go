package main

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

// Client - rethinkdb client driver
type Client struct {
	db      string
	session *r.Session
}

func NewClient(session *r.Session) *Client {

	return &Client{
		db:      "test",
		session: session,
	}
}
