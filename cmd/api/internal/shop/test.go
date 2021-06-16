package shop

import (
    "encoding/json"
    "fmt"
    "go-ticket/cmd/api/com"
    "io/ioutil"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/pieterclaerhout/go-log"
)

type ReqMenuJson struct {
    Menujson string `form:"menujson"`
}

func TestMenu(c *gin.Context) {
    req := ReqMenuJson{}
    if err := c.Bind(&req); err != nil {
        com.Error(c, err.Error())
        return
    }
    log.InfoDump(req, "params")
    res, err := reqWx(req.Menujson)
    if err != nil {
        com.Error(c, err.Error())
    }
    aa := map[string]interface{}{}
    json.Unmarshal([]byte(res), &aa)
    com.Responce(c, aa)
}

func reqWx(menujson string) (ss string, err error) {
    log.InfoDump(menujson, "menujson")

    url := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=45_3KV2l81kaijP9hoQQwJs6JVXrCMxmWTMeT96CaRxLGvIiHeS_OdaQXrFt2sJqoO2ix3rN72BRN48Pc1eCXaW0jZTDwXOtasRu9NTao7GHWCYG7_cLpqlp5F6R16zy5vLD9YFedMjE0awfD7aEVWeABANEF"
    method := "POST"

    payload := strings.NewReader(menujson)

    client := &http.Client{}
    req, err := http.NewRequest(method, url, payload)

    if err != nil {
        return
        fmt.Println(err)
        return
    }
    req.Header.Add("Content-Type", "application/json")

    res, err := client.Do(req)
    if err != nil {
        return

        fmt.Println(err)
        return
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return

        fmt.Println(err)
        return
    }
    fmt.Println(string(body))
    return string(body), nil
}
