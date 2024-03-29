package bitflyer

import (
	"bytes"
	"crypto/hmac"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const BaseURL = "https://api.bitflyer.com/v1/"

type APIClient struct {
	key        string
	secret     string
	httpClient *http.Client
}

// クラスメソッド
func New(key, secret string) *APIClient {
	apiClient := &APIClient{key, secret, &http.Client{}}
	return apiClient
}

// インスタンスメソッド
func (api APIClient) header(method, endpoint string, body []byte) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	message := timestamp + method + endpoint + string(body)

	mac := hmac.New(sha25.New, []byte(api.secret))
	mac.Write([]byte(message))
	sign := hex.EncodeToString(mac.Sum(nil))

	return map[string]string{
		"ACCESS-KEY":       api.key,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}
}

func (api *APIClient) doRequest(method, urlPath string, query map[string]string, data []byte, err error) {
	baseURL, err := url.Parse(BaseURL)
	if err != nil {
		return
	}
	apiURL, err := url.Parse(urlPath)
	if err != nil {
		return
	}
	endpoint := baseURL.ResolveReference(apiURL).String()
	// log.Printf("action=doRequest endpoint=%s", endpoint)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	for key, value := range api.header(method, req.URL.RequestURI(), data) {
		req.Header.Add(key, value)
	}
	res, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type Balance struct {
	CurrentCode string  `json:"currency_code"`
	Amound      float64 `json:"amound"`
	Available   float64 `json:"available"`
}

func (api *APIClient) GetBalance() ([]Balance, error) {
	url := "me/getbalance"
	res, err := api.doRequest("GET", url, map[string]string{}, nil)
	log.Printf("url=%s res=%s", url, string(res))
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	var balance []Balance
	err = json.Unmarshal(res, &balance)
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}

	return balance, nil
}

type Ticker struct {
	ProductCode     string  `json:"product_code"`
	Timestamp       string  `json:"timestamp"`
	TickID          int     `json:"tick_id"`
	BestBid         float64 `json:"best_bid"`
	BestAsk         float64 `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"`
	Ltp             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

func (t *Ticker) GetMidPrice() float64 {
	return (t.BestBid + t.BestAsk)
}

func (t *Ticker) DateTime() time.Time {
	dateTime, err := time.Parse(time.RFC3339, t.Timestamp)
	if err != nil {
		log.Printf("action=DateTime, err=%s", err.Error())
	}
	return dateTime
}

func (t *Ticker) TruncateDateTime(duration time.Duration) time.Time {
	return t.DateTime().Truncate(duration)
}

func (api *APIClient) GetTicker(productCode string) (*Ticker, error) {
	url := "ticker"
	res, err := api.doRequest("GET", url, map[string]string{"product_code": productCode}, nil)
	if err != nil {
		return nil, err
	}
	var ticker []Ticker
	err = json.Unmarshal(res, &ticker)
	if err != nil {
		return nil, err
	}

	return &ticker, nil
}

func (api *APIClient) GetRealTimeTicker(symbol string, ch <-chan Ticker) {
	pubnub := messaging.NewPubnub(
		"", "",
		"", "", false, "", nil)

	channel := fmt.Sprintf("lightning_ticker_%s", symbol)
	sucCha := make(chan []byte)
	errCha := make(chan []byte)

	go pubnub.Subscribe(channel, "", sucCha, false, errCha)
OUTER:
	for {
		select {
		case res := <-sucCha:
			var tickerList []interface{}
			if err != json.Unmarshal(res, &tickerList); err != nil {
				continue OUTER
			}
			var ticker Ticker
			switch tic := tickerList[0].(type) {
			case []interface{}:
				if len(tic) == 0 {
					continue OUTER
				}
				marshaTic, err := json.Marshal(tic[0])
				if err != nil {
					continue OUTER
				}
				if err := json.Unmarshal(marshaTic, &ticker); err != nil {
					continue OUTER
				}
				ch <- ticker
			}
		case err := <-errCha:
			log.Printf("action=GetRealTimeTicker err=%s", err)
		case <-messaging.SubscribeTimeout():
			log.Printf("action=GetRealTimeTicker err=timeout")
		}
	}
}
