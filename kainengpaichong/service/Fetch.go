package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kainengpaichong/model"
	"net/http"
	url2 "net/url"
	"strconv"
	"strings"
)

type Responses struct {
	Data struct {
		Data1  []Name `json:"1"`
		Data4  []Name `json:"4"`
		Data8  []Name `json:"8"`
		Data10 []Name `json:"10"`
	} `json:"data"`
}

type Name struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
type res1 struct {
	Data []Data `json:"data"`
}
type Data struct {
	Id       int    `json:"id"`
	Word     string `json:"word"`
	QsStar   int    `json:"qsStar"`
	SugPrice string `json:"sugPrice"`
}

type res2 struct {
	Data []price `json:"data"`
}
type price struct {
	SuggestPriceStr string `json:"suggestPriceStr"`
	TopPosPriceStr  string `json:"topPosPriceStr"`
	MinPriceStr     string `json:"minPriceStr"`
}
type res3 struct {
	Ok bool `json:"ok"`
}

var Cookies = "t=0f8bb0662b87687ca4239dcf02acac5a; cna=kCF2HFsMABkBASQOA7KUz60l; xman_i=aid=2209288609069; _ga=GA1.1.1887126010.1689557478; recommend_login=email; l=fBasuwDnPMHWjep-BOfwPurza77OSIRAguPzaNbMi9fP93f95oZRB1C3EkTpC3GVFsiXR3S3hc2kBeYBqQd-nxv9kloV96HmnhkSGN5..; _ga_GKMVMVMZNM=GS1.1.1706869548.47.0.1706869548.0.0.0; uns_unc_f=trfc_i=safcps^224eq2j0^51mjaujf^1ho4h9g97; ali_apache_id=11.187.61.173.1712820434875.699713.9; cookie2=151309256286f64221eda73393495ec9; _tb_token_=74ee753b3e535; _samesite_flag_=true; _ym_uid=1713969996196033140; _ym_d=1713969996; cbc=G556FDE370A2A32D3F3C311042FF0F8821179E0CAF54F58490E; _hvn_login=4; ali_apache_tracktmp=W_signed=Y; intl_locale=zh_CN; XSRF-TOKEN=54b2e113-db70-4beb-a437-e2e2c0639ce4; NWG=NNW; _uetvid=741c69e0024911efab1e7bb4cac6d11e; _ga_RVSKK1KF5N=GS1.1.1716965834.1.0.1716965834.60.0.0; buyer_ship_to_info=local_country=CA; sc_g_cfg_f=sc_b_currency=USD&sc_b_locale=en_US; _m_h5_c=12a4e845b743f388e330ac130f1f0934_1717576980523%3Bf547d307c10dc4fc426a714fc142d771; banThirdCookie=flag; xman_us_f=x_locale=zh_CN&last_popup_time=1717655259781&x_l=1&x_user=CN|Chunhong|Yang|cgs|244740850&no_popup_today=n; intl_common_forever=wZc03Y+EWnAQiwRCdtxWWE4fuEgXw5joyX3IJghWkn9rnz2zeM931A==; xman_us_t=l_source=alibaba&sign=y&need_popup=y&x_user=rTwRT8JTtZYB2nUTJiJcr85ybdkU0R7GNnMe7YldhiE=&ctoken=7o0t18u8741d&x_lid=szkn; xman_f=d2GQN1f+ff+c8xJEIWn6askvLjpQ0zblXD2TAGOMY309rgIeMJLR3QsFmylefI1yXV3/Xfv8ddveJ/sCh5XlKexMVTXBVFHPMkkct1ZkR3A02arytwCxLsiph3+pnNFyEP+fTdSQ2yh9LeA5+6uQo80RXuejTmHuvOJBSzmoGSxRbFlHnDoq7nAZgUlQc3nxyD4E5vrjtb/5jJZbM5MjU78OtBcQw2dfqeiiKTglRnk8eqECzJGWIweLhLgDEKg6WdQTA8fAM1DJ0LeMIVvncH4yk9L/3JCVOe07xq5GzlpCXMWGNfh/OwYhvwkbB7ZysP+hlICST6onUew69KB4mgeASUdUW3lDai/5C+grPosBEgbNEbgh0txKgabeex3aKC0c+R00KNw=; acs_usuc_t=acs_rt=cfc4c787726e43dd8f85fcd88527e1c0; havana_lgc2_4=eyJoaWQiOjIyMDkyODg2MDkwNjksInNnIjoiMTkxZDkyOTc4MTI4ZjFjMmNhYjU2MTkxZjY2NWU2MzkiLCJzaXRlIjo0LCJ0b2tlbiI6IjFTTnBaODFpMUpYM3I0VEVHZ1BxU2lRIn0; _hvn_lgc_=4; havana_tgc=eyJjcmVhdGVUaW1lIjoxNzE3NjU1MjU5NzU2LCJsYW5nIjoiemhfQ04iLCJwYXRpYWxUZ2MiOnsiYWNjSW5mb3MiOnsiNCI6eyJhY2Nlc3NUeXBlIjoxLCJtZW1iZXJJZCI6MjIwOTI4ODYwOTA2OSwidGd0SWQiOiIxOEwwb0Fyc25TSFN0Nk9hVmlRdUdGdyJ9fX19; sgcookie=E100I0oohMkEudZY2JhaKmHzbvWlx7x3UNDBtCVrKCurMCfcyOFD3Oyx51H1FCmoz/m3Y31MryBgpQW9AB9cifpXFdiMtK2+bYiq95HRTX/PkWk=; csg=647d4d4d; xman_t=G49n7tq5+Uwuyw+R7a5b7WTKUHYxr/uCQH3t8mx63FGmkDRpCQT++5OHx3tFvq0EEy9s8S5Wwqj6BzHE96czmCpA22043eGaMmyq7fG51Gca5ZeWE+XdmKFT49KVGvIMujIF7yovYatQRSRVQdGMfYD+J8vmbsNT1POdgfyzRUMDROT3eWy7K2qBBGd53TD3WdYZwIOwdwJ6FZbKmv3pWsty74+nw1DU42ZIwhElnv0S7KT9Zpswag/pEfXNOuIi5ebzyjAsPAwFRwKS0wNEo8/KW5CreHov7A1Kf3sP+CUbFEjzrNOTPatOJeqVEhEixoFz69HTP81DKr0CIfUDBNDg64+mOcGZmKkoL9zbOnyO2MLuBau0ZxkBCDDEwQeOU73ZmjyVvWsUBG+juHwLq7HoH1B5ZETWrs59nPyxFe5I+2Bd2EBxJ52gWMG5VX8yT0OXoHX84NbXFV5Yqwqb5P6/my1Z0RieXbJr/qlU0mqm2lHv7gUQCIhs1utB0Y5hihrbL6zTFdn+A1oT4LmdCvpeF9frW1rC9+egR8ceO5QgHKATLcNOOb/XWL9G/jumJvCTwTz8pxQ9Q31723olxUnI9GWiNnj6AZqwZIlQemqenMB71Bj1vH6zy9B5TcYNhSaMu8tvrjCx2381l2r+8aSdH7Ya9tv7Cq3YpCGy+LBx2hVcpeEt0w==; ali_apache_track=ms=|mt=3|mid=szkn; JSESSIONID=6203BA0C28FDAFF8D383002F931C0721; xlly_s=1; _csrf_token=1718468646811; _m_h5_tk=e4d97dc756614665f31503d5e4de34f0_1718615493293; _m_h5_tk_enc=abd5ee3feec08f493ca813d9b46b4ea6; ug_se_c=organic_1718607275686; icbu_s_tag=10_5_11; tfstk=fyvoFuGlKQ55fTrKrKXSONzO6acv3z6CYeedJ9QEgZ7byzeRTIXHlUmIVXGWiwY2oaoSRaLc-eLcJ0OJeBYvkefKVLnWikx6oefIVwQF-eKewmHtB3t5d97nWAH9vYj69ejPTgSqYGsh85yIjJ-5d9uYNE02e3TMlj_q-9oc3MSFY95PzS-VlMFP8e7PgrSGf97e89oc3gsP8wSF4sPVxj_zLNvehKulmSlOD6HpEgfl-JQ0Kmod2_bM4Z0Z7K2fZN-PoJ0a0lJG-Z9ZkjxX8Q8RceDioOR6opfNKY41qI-G3aBZsSXkDU9GUd0ueiChrIJPikyeb37f_BYor7_ywEf10srnCiLOu39yiDa24FQlUi5xQcxVTIp51L34EOR62Tde7qFRzCjF4lP4_uxLdiop3WNCaisc5dD0FarlGQkommVsN_Sf2VnmmWNCaiscWmm01y1Pcg3O.; isg=BJycZEtJFsCvBuJqHM-YJqSabbpOFUA_OnG8xHadKwdqwT9LniH2zog_JSk5yXiXgo"

func Fetch() {
	url := "https://www2.alibaba.com/api/campaign/type/allList"
	client := &http.Client{}
	formData := url2.Values{}
	formData.Set("ads", "{\"productLineId\":110101}")
	formData.Set("type", "normal")
	formData.Set("data", "{\"paging\":true,\"typeList\":[1,2,4,6,8,10,23,24,44],\"page\":1,\"size\":20}")
	formData.Set("_csrf", "54b2e113-db70-4beb-a437-e2e2c0639ce4")
	req, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode()))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", Cookies)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code:", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	var response Responses
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	for _, item := range response.Data.Data1 {
		if item.ID == 255851929 {
			continue
		}
		Fetch1(item.ID, item.Title)
	}
}

// Fetch1 查看词
func Fetch1(planId int, plantitle string) {
	url := fmt.Sprintf("https://www2.alibaba.com/api/campaign/%d/bidword/query?ads=%7B%22productLineId%22%3A110101%7D&type=normal&data=%7B%22queryEffect%22%3A%7B%22interval%22%3A%22all%22%2C%22mode%22%3A%22online%22%7D%7D&_csrf=54b2e113-db70-4beb-a437-e2e2c0639ce4", planId)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Cookie", Cookies)
	req.Header.Set("Referer", "https://alicrm.alibaba.com/?spm=a2700.8443394.onepage_nav.9.97553e5f1hUe0q")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code:", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	var res1 res1
	err = json.Unmarshal(body, &res1)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	for _, item := range res1.Data {
		sugPrice, err := strconv.ParseFloat(item.SugPrice, 64)
		if err != nil {
			fmt.Println("转换失败:", err)
			return
		}
		Fetch2(planId, item.Word, item.Id, sugPrice, plantitle)

	}

}

// Fetch2 查看价格
func Fetch2(planid int, word string, wordid int, nowprice float64, plantitle string) {
	url := fmt.Sprintf("https://www2.alibaba.com/api/recommend/bid/keyword/bid")
	client := &http.Client{}
	formData := url2.Values{}
	formData.Set("type", "normal")
	formData.Set("ads", "{\"productLineId\":110101}")
	data := fmt.Sprintf(`{"campaignId":%d,"keywordList":["%s"]}`, planid, word)
	formData.Set("data", data)
	formData.Set("_csrf", "54b2e113-db70-4beb-a437-e2e2c0639ce4")
	req, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", Cookies)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code:", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	var res2 res2
	err = json.Unmarshal(body, &res2)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	fmt.Println(word, res2.Data[0].SuggestPriceStr, res2.Data[0].MinPriceStr, res2.Data[0].TopPosPriceStr, wordid, planid)
	suggestPriceStr, err := strconv.ParseFloat(res2.Data[0].SuggestPriceStr, 64)
	maxpriceStr, err := strconv.ParseFloat(res2.Data[0].TopPosPriceStr, 64)
	minPriceStr, err := strconv.ParseFloat(res2.Data[0].MinPriceStr, 64)

	if err != nil {
		// 解析失败，处理错误
		fmt.Println("Error parsing suggestPriceStr:", err)
		return
	}
	sqldata := model.Fromplan{
		PlanName:        plantitle,
		Planid:          planid,
		Wordid:          wordid,
		Word:            word,
		SuggestPriceStr: suggestPriceStr,
		Maxprice:        maxpriceStr,
		MinPrice:        minPriceStr,
		NowPrice:        nowprice,
	}
	model.CreateOrUpdateData(&sqldata)
	Fetch3(planid, wordid, suggestPriceStr)
}

// Fetch3 修改价格
func Fetch3(planid int, wordid int, price float64) {
	sqldata := model.Fromplan{
		Wordid: wordid,
	}
	setprice, err := model.SelectMaxPrice(&sqldata)
	if err != nil {
		fmt.Println("查找最大出价失败:", err)
		return
	}
	if setprice == 0 {
		sqldata := model.Fromplan{
			Wordid:   wordid,
			Setprice: 20.0,
		}
		model.CreateOrUpdateData(&sqldata)
	}
	sqldatas := model.Fromplan{
		Wordid: wordid,
	}
	setprices, err := model.SelectMaxPrice(&sqldatas)
	// 比较 price 和 setprice，取较小的值
	if price < setprices {
		sqldata := model.Fromplan{
			NowPrice: price,
		}
		model.CreateOrUpdateData(&sqldata)
		fmt.Println("价格未超过最大值，保持推荐价格:", price)
	} else {
		price = setprices
		sqldata := model.Fromplan{
			NowPrice: price,
			Remarks:  "超过最高价格",
		}
		model.CreateOrUpdateData(&sqldata)
		fmt.Println("价格超过最大值，调整为最大值:", price)
	}

	url := fmt.Sprintf("https://www2.alibaba.com/api/campaign/%d/bidword/update", planid)
	client := &http.Client{}
	formData := url2.Values{}
	formData.Set("type", "normal")
	formData.Set("ads", "{\"productLineId\":110101}")
	data := fmt.Sprintf("{\"updateType\":\"add_price\",\"keywordList\":[{\"id\":%d,\"price\":%f}]}", wordid, price)
	formData.Set("data", data)
	formData.Set("_csrf", "54b2e113-db70-4beb-a437-e2e2c0639ce4")
	req, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", Cookies)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code:", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	var res3 res3
	err = json.Unmarshal(body, &res3)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	if res3.Ok {
		fmt.Println("修改成功")
	} else {
		fmt.Println("修改失败")
	}

}
