package bilibili

import (
	"BiliMallRepublish/bilibili/apiStruct"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// userSessdata 存储用户输入的sessdata
var userSessdata string

// 封装一个带sessdata的mall.bilibili.com的api通用请求方法
func callApi(url string, method string, body io.Reader, sessData *string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Host", "mall.bilibili.com")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Connection", "close")
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_3_1 like Mac OS X) AppleWebKit/17613.3.9.1.16 (KHTML, like Gecko) Version/15.3.0 Mobile/15E148 Safari/17613.3.9.1.16")
	req.Header.Add("Accept-Language", "zh-CN,zh-Hans;q=0.9")
	req.Header.Add("Accept-Encoding", "deflate")
	if sessData != nil {
		req.Header.Add("Cookie", "SESSDATA="+*sessData+";")
	}
	if method == "POST" {
		req.Header.Add("Content-Type", "application/json")
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// SetSessdata 对外暴露的sessdata设置方法
func SetSessdata(sessdata string) {
	userSessdata = sessdata
}

// GetNowPublishedList 获取用户已上架商品的api
func GetNowPublishedList() (err error, publishItemList []apiStruct.PublishItem) {
	url := "https://mall.bilibili.com/mall-magic-c/internet/c2c/items/pageQueryMyPublish?pageSize=200&pageNo=1&filterType=1"
	method := "GET"
	resp, err := callApi(url, method, nil, &userSessdata)
	if err != nil {
		return
	}
	publishedListResponse := &apiStruct.PublishedListResponse{}
	err = json.Unmarshal(resp, publishedListResponse)
	if err != nil {
		return
	}
	publishItemList = publishedListResponse.Data.List
	return
}

// DropItem 通过c2cItemsId下架商品的api
func DropItem(c2cItemsId int64) (err error) {
	url := "https://mall.bilibili.com/mall-magic-c/internet/c2c/items/dropC2cItems"
	method := "POST"
	payload := strings.NewReader(fmt.Sprintf(`{"c2cItemsId":"%d"}`, c2cItemsId))
	resp, err := callApi(url, method, payload, &userSessdata)
	if err != nil {
		return
	}
	dropItemResponse := &apiStruct.DropItemResponse{}
	err = json.Unmarshal(resp, dropItemResponse)
	if err != nil {
		return
	}
	if !dropItemResponse.Data.Success {
		err = errors.New(dropItemResponse.Data.ErrMsg)
		return
	}
	return nil
}

// CheckItems 通过blindBoxIds集合检查物品有效性，同时获取用于上架token的api
func CheckItems(items []string) (err error, token string, showTime string) {
	url := "https://mall.bilibili.com/magic-c/c2c/blind-box/check"
	method := "POST"
	payloadStr := fmt.Sprintf(`{"blindBoxIds":[%s]}`, strings.Join(items, ","))
	payload := strings.NewReader(payloadStr)
	resp, err := callApi(url, method, payload, &userSessdata)
	if err != nil {
		return
	}
	checkResponse := &apiStruct.CheckResponse{}
	err = json.Unmarshal(resp, checkResponse)
	if err != nil {
		return
	}
	if !checkResponse.Data.Status {
		err = errors.New(checkResponse.Data.InvalidMsg)
		return
	}
	token = checkResponse.Data.Token
	showTime = checkResponse.Data.ShowTime
	return
}

// PublishItem 通过price价格及上一步获取的token进行商品上架的api，isConfirm为false时接口返回二次确认信息，为true时正式上架
func PublishItem(price string, token string, isConfirm bool) (err error, isPublish bool, discount string) {

	url := "https://mall.bilibili.com/mall-magic-c/internet/c2c/v2/publish"
	method := "POST"
	payloadStr := fmt.Sprintf(`{"price":"%s","token":"%s","isConfirm":%s}`, price, token, strconv.FormatBool(isConfirm))
	payload := strings.NewReader(payloadStr)
	resp, err := callApi(url, method, payload, &userSessdata)
	if err != nil {
		return
	}
	publishResponse := &apiStruct.PublishResponse{}
	err = json.Unmarshal(resp, publishResponse)
	if err != nil {
		return
	}
	if publishResponse.Code != 0 {
		err = errors.New(publishResponse.Message)
		return
	}
	if publishResponse.Data.Status == "CONFIRM" {
		discount = publishResponse.Data.ConfirmInfo.Discount
		return
	} else if publishResponse.Data.Status == "SUCCESS" {
		isPublish = true
		return
	}
	err = errors.New("未知错误！")
	return
}
