package nft

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {

	url := "https://eth-mainnet.alchemyapi.io/v2/demo"

	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"alchemy_getAssetTransfers\",\"params\":[{\"fromBlock\":\"0x0\",\"toBlock\":\"latest\",\"withMetadata\":false,\"excludeZeroValue\":true,\"maxCount\":\"0x3e8\"}]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
