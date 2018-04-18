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
	"net"
	"testing"
	"github.com/rgalanakis/gomemcache/memcache"
	"time"
	"os"
)

var (
	testServerWithAuth = os.Getenv("GOMEMCACHE_SASL_SERVER")
	testServerUsername = os.Getenv("GOMEMCACHE_USERNAME")
	testServerPass = os.Getenv("GOMEMCACHE_PASSWORD")
)

func TestAuth(t *testing.T) {
	if testServerWithAuth == "" || testServerUsername == "" || testServerPass == "" {
		t.Errorf("Auth test environment not setup properly.")
	}
	_, err := net.Dial("tcp", testServerWithAuth)
	if err != nil {
		t.Errorf("Failed to dial %s: %v", testServer, err)
	}

	cl := memcache.New(testServerWithAuth)
	cl.Timeout = 5 * time.Second
	cl.SetAuthentication(testServerUsername, []byte(testServerPass))
	_, err = cl.Get("tester")
	if err != nil && err != memcache.ErrCacheMiss {
		t.Errorf("Failed to get/connect: %v", err)
	}
}
