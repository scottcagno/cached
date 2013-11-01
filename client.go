// *
// * Copyright 2013, Scott Cagno, All rights reserved
// * BSD Licensed - sites.google.com/site/bsdc3license
// *
// * cached :: client.go :: database client
// *

package cached

import (
	"encoding/json"
	"log"
	"net"
)

// Client struct
type Client struct {
	conn net.Conn
	enc  *json.Encoder
	dec  *json.Decoder
}

// InitClient initializes and returns a new client
func InitClient() *Client {
	return &Client{}
}

// Open connects to provided host and initializes dec, and enc
func (self *Client) Open(host string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Printf("*client open: %s\n", err)
	}
	self.conn = conn
	self.enc = json.NewEncoder(self.conn)
	self.dec = json.NewDecoder(self.conn)
}

// Close terminates the connection
func (self *Client) Close() {
	err := self.conn.Close()
	if err != nil {
		log.Printf("*client close: %s\n", err)
	}
}

// Raw sends raw byte data to the server
func (self *Client) Raw(b []byte) bool {
	self.enc.Encode(b)
	var r bool
	self.dec.Decode(&r)
	return r
}

// Add creates a new entry
func (self *Client) Add(k string, v interface{}) bool {
	self.enc.Encode(Request{"add", k, v})
	var r bool
	self.dec.Decode(&r)
	return r
}

// Set updates an existing value or creates a new one
func (self *Client) Set(k string, v interface{}) bool {
	self.enc.Encode(Request{"set", k, v})
	var r bool
	self.dec.Decode(&r)
	return r
}

// Get retuns a value for a matching key
func (self *Client) Get(k string) interface{} {
	self.enc.Encode(Request{"get", k, nil})
	var r interface{}
	self.dec.Decode(&r)
	return r
}

// Del removes a key and value if they exist
func (self *Client) Del(k string) bool {
	self.enc.Encode(Request{"del", k, nil})
	var r bool
	self.dec.Decode(&r)
	return r
}

// Exp sets a key to expire in `n` seconds
func (self *Client) Exp(k string, n int64) bool {
	self.enc.Encode(Request{"exp", k, n})
	var r bool
	self.dec.Decode(&r)
	return r
}

// Ttl checks the time to live for a key
func (self *Client) Ttl(k string) int64 {
	self.enc.Encode(Request{"ttl", k, nil})
	var r int64
	self.dec.Decode(&r)
	return r
}
