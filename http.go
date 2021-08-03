package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type fastResp struct {
	URL         string
	body        []byte
	contentType string
	TimeCost    int64
	Error       string
}

func fastGet(urls []string) (*fastResp, error) {
	fast := make(chan *fastResp)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	for _, url := range urls {
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		go func(_url string) {
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				safeSend(fast, &fastResp{URL: _url, Error: err.Error()})
				return
			}
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				safeSend(fast, &fastResp{URL: _url, Error: err.Error()})
				return
			}
			safeSend(fast, &fastResp{URL: _url, body: bytes, contentType: resp.Header.Get("content-type")})
		}(url)
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case resp := <-fast:
		close(fast)
		return resp, nil
	case <-time.After(time.Second * 2):
		return nil, errors.New("wait timeout")
	}
}

func fastTest(urls []string) ([]fastResp, error) {
	wg := sync.WaitGroup{}
	wg.Add(len(urls))

	client := http.Client{
		Timeout: time.Second * 15,
	}
	results := make(chan fastResp, len(urls))

	for _, url := range urls {
		go func(_url string, _wg *sync.WaitGroup) {
			start := currentMs()
			resp, err := client.Get(_url)
			if err != nil {
				fmt.Println("Done ", _url)
				results <- fastResp{URL: _url, Error: err.Error(), TimeCost: currentMs() - start}
				_wg.Done()
				return
			}
			_, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Done ", _url)
				results <- fastResp{URL: _url, Error: err.Error(), TimeCost: currentMs() - start}
				_wg.Done()
				return
			}
			fmt.Println("Done ", _url)
			results <- fastResp{URL: _url, TimeCost: currentMs() - start}
			_wg.Done()
		}(url, &wg)
	}

	wg.Wait()
	var respList = make([]fastResp, 0)
	respList = append(respList, <-results)
	respList = append(respList, <-results)
	respList = append(respList, <-results)

	close(results)
	return respList, nil
}

func currentMs() int64 {
	return time.Now().UnixNano() / 1e6
}

func safeSend(ch chan *fastResp, value *fastResp) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	ch <- value
	return false
}
