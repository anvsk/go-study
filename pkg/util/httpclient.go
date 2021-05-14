package util

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "time"

    "github.com/pieterclaerhout/go-log"
    "github.com/sethgrid/pester"
)

var DefaultClient *pester.Client
var AsyncClient *pester.Client

func initHttpClient() {
    // 普通请求
    DefaultClient = pester.New()
    DefaultClient.Concurrency = 1
    DefaultClient.MaxRetries = 3
    DefaultClient.Backoff = pester.ExponentialJitterBackoff
    DefaultClient.Timeout = 3 * time.Second
    // 并发请求，接收最快的一个响应
    AsyncClient = pester.New()
    AsyncClient.Concurrency = 3
    AsyncClient.MaxRetries = 2
    AsyncClient.Backoff = pester.ExponentialJitterBackoff
    AsyncClient.Timeout = 3 * time.Second

}

type Req struct {
    Url     string
    Params  interface{}
    Res     interface{}
    Headers map[string]string
    Check   func([]byte) error //用于校验结果
    Async   bool
}

func Get(req Req) error {
    log.Debug("GET", req.Url)
    query := ""
    if tmp, flag := req.Params.(map[string][]string); flag {
        query = url.Values(tmp).Encode()
    }
    res, _ := http.NewRequest("GET", req.Url+"?"+query, nil)
    if len(req.Headers) > 0 {
        for k, v := range req.Headers {
            res.Header.Add(k, v)
        }
    }
    response, err := getClient(req).Do(res)
    if err != nil {
        log.ErrorDump(err.Error(), req.Url)
        return err
    }
    body, _ := ioutil.ReadAll(response.Body)
    json.Unmarshal(body, req.Res)
    if req.Check != nil {
        if err := req.Check(body); err != nil {
            log.Error("ResError", req.Url, err.Error())
            return err
        }
    }
    return nil
}

func Post(req Req) error {
    log.Debug("POST", req.Url)
    params, _ := json.Marshal(req.Params)
    res, _ := http.NewRequest("POST", req.Url, strings.NewReader(string(params)))
    if len(req.Headers) > 0 {
        for k, v := range req.Headers {
            res.Header.Add(k, v)
        }
    }
    defer res.Body.Close()
    response, err := getClient(req).Do(res)
    if err != nil {
        log.ErrorDump(err.Error(), req.Url)
        return err
    }
    body, _ := ioutil.ReadAll(response.Body)
    json.Unmarshal(body, req.Res)
    if req.Check != nil {
        if err := req.Check(body); err != nil {
            log.Error("ResError", req.Url, err.Error())
            return err
        }
    }
    return nil
}

func getClient(req Req) *pester.Client {
    if req.Async {
        return AsyncClient
    }
    return DefaultClient
}
