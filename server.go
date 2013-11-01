// *
// * Copyright 2013, Scott Cagno, All rights reserved
// * BSD Licensed - sites.google.com/site/bsdc3license
// *
// * cached :: server.go :: database server
// *

package cached

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// ListenAndServe attempts to listen on specified port
func (self *Store) ListenAndServe(port string) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Cached v0.1a\nListening on %v\n", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("*server accept: %s\n", err)
			continue
		}
		go handleConn(conn, self)
	}
}

// handleConn handles client connection
func handleConn(conn net.Conn, self *Store) {
	log.Printf("%v connected\n", conn.RemoteAddr())
	dec, enc := json.NewDecoder(conn), json.NewEncoder(conn)
	for {
		var r Request
		err := dec.Decode(&r)
		if err == io.EOF {
			break
		} else if err != nil {
			conn.Close()
			return
		} else {
			conn.SetDeadline(time.Now().Add(time.Duration(300) * time.Second))
		}
		fmt.Printf("%+v\n", r)
		switch r.Cmd {
		case "add":
			enc.Encode(self.add(&r))
		case "set":
			enc.Encode(self.set(&r))
		case "get":
			enc.Encode(self.get(&r))
		case "del":
			enc.Encode(self.del(&r))
		case "exp":
			enc.Encode(self.exp(&r))
		case "ttl":
			enc.Encode(self.ttl(&r))
		default:
			enc.Encode(false)
		}
	}
}
