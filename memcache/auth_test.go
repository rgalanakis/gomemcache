/*
Copyright 2011 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package memcache provides a client for the memcached cache server.
package memcache_test

import (
	"github.com/rgalanakis/gomemcache/memcache"
	"net"
	"os"
	"testing"
	"time"
)

const (
	saslServerKey = "GOMEMCACHE_SASL_SERVER"
	usernameKey = "GOMEMCACHE_USERNAME"
	passwordKey = "GOMEMCACHE_PASSWORD"
)

func TestAuth(t *testing.T) {
	saslServer := os.Getenv(saslServerKey)
	username := os.Getenv(usernameKey)
	password := os.Getenv(passwordKey)
	if saslServer == "" || username == "" || password == "" {
		t.Errorf(
			"Auth envvars not present.\n%s: %d chars, %s: %d chars, %s: %d chars",
			saslServerKey,
			len(saslServer),
			usernameKey,
			len(username),
			passwordKey,
			len(password))
	}
	_, err := net.Dial("tcp", saslServer)
	if err != nil {
		t.Errorf("Failed to dial %s: %v", testServer, err)
	}

	cl := memcache.New(saslServer)
	cl.Timeout = 5 * time.Second
	cl.SetAuthentication(username, []byte(password))
	_, err = cl.Get("tester")
	if err != nil && err != memcache.ErrCacheMiss {
		t.Errorf("Failed to get/connect: %v", err)
	}
}
