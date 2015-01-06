GoSS
====

[![Build Status](http://img.shields.io/travis/osamingo/goss.svg?style=flat)](https://travis-ci.org/osamingo/goss)
[![Coverage](http://img.shields.io/codecov/c/github/osamingo/goss.svg?style=flat)](https://codecov.io/github/osamingo/goss)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/osamingo/goss/blob/master/LICENSE)

## Description

GoSS provides sorted-slice for golang.

## Installation

```
$ go get github.com/osamingo/goss
```

## Quick start

```go
package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	mrand "math/rand"
	"sync"

	"github.com/osamingo/goss"
)

type Result struct {
	ID    string
	Score int
}

func (r *Result) Target() int64 {
	return int64(r.Score)
}

func (r *Result) Priority() string {
	return r.ID
}

func randStr() string {
	rb := make([]byte, 32)
	rand.Read(rb)
	return base64.URLEncoding.EncodeToString(rb)
}

func main() {

	rs := []*Result{}
	fmt.Println("# Before")
	for i := 0; i < 10; i++ {
		r := &Result{ID: randStr()}
		fmt.Printf("%2d: %v\n", i+1, r)
		rs = append(rs, r)
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(rs))

	finChan := make(chan bool)
	retChan := make(chan *Result, len(rs))

	go func() {
		wg.Wait()
		finChan <- true
	}()

	for _, r := range rs {
		go func(res *Result) {
			defer wg.Done()
			res.Score = mrand.Intn(1000)
			retChan <- res
		}(r)
	}

	s := &goss.SortedSlice{DESC: true}

LOOP:
	for {
		select {
		case <-finChan:
			break LOOP
		case r := <-retChan:
			s.Add(r)
		}
	}

	fmt.Println("\n# After")
	for i, r := range s.S {
		fmt.Printf("%2d: %v\n", i+1, r)
	}

}
```

## Tips

```go
type Result struct {
	ID    string
	Score int
	
	TargetFunc   func() int64
	PriorityFunc func() string
}

func (r *Result) Target() int64 {
	return r.TargetFunc()
}

func (r *Result) Priority() string {
	return r.PriorityFunc()
}
```

## License

MIT

