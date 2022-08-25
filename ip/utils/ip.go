package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

// @Title  请填写文件名称（需修改）
// @Description  请填写文件描述（需修改）
// @Author  clx  2022/8/24 5:23 下午
// @Update  clx  2022/8/24 5:23 下午

// IPInfo ip地址信息
type IPInfo struct {
	IP string
	Location string
}

// 获取ip信息
func GetIpInfo(ip string) (info *IPInfo) {
	resp:= HttpIPApiResp{}
	res := HttpIPApi(ip)
	err := json.Unmarshal([]byte(res),&resp)
	if err!=nil{
		fmt.Println(err)
		return 
	}
	if len(resp.Result)<=0{
		return
	}
	info = &IPInfo{
		IP:       ip,
		Location: resp.Result[0].DisplayData.ResultData.TplData.Location,
	}
	return 
}

// HttpIPApiResp HttpIPApiResp
type HttpIPApiResp struct {
	Srcid      string `json:"Srcid"`
	ResultCode string `json:"ResultCode"`
	Status     string `json:"status"`
	QueryID    string `json:"QueryID"`
	Result     []struct {
		DisplayData struct {
			Strategy struct {
				TempName  string `json:"tempName"`
				Precharge string `json:"precharge"`
				CtplOrPhp string `json:"ctplOrPhp"`
			} `json:"strategy"`
			ResultData struct {
				TplData struct {
					Srcid         string `json:"srcid"`
					Resourceid    string `json:"resourceid"`
					OriginQuery   string `json:"OriginQuery"`
					Origipquery   string `json:"origipquery"`
					Query         string `json:"query"`
					Origip        string `json:"origip"`
					Location      string `json:"location"`
					Userip        string `json:"userip"`
					Showlamp      string `json:"showlamp"`
					Tplt          string `json:"tplt"`
					Titlecont     string `json:"titlecont"`
					Realurl       string `json:"realurl"`
					ShowLikeShare string `json:"showLikeShare"`
					ShareImage    string `json:"shareImage"`
					DataSource    string `json:"data_source"`
				} `json:"tplData"`
				ExtData struct {
					Tplt        string `json:"tplt"`
					Resourceid  string `json:"resourceid"`
					OriginQuery string `json:"OriginQuery"`
				} `json:"extData"`
			} `json:"resultData"`
		} `json:"DisplayData"`
		ResultURL        string        `json:"ResultURL"`
		Weight           string        `json:"Weight"`
		Sort             string        `json:"Sort"`
		SrcID            string        `json:"SrcID"`
		ClickNeed        string        `json:"ClickNeed"`
		SubResult        []interface{} `json:"SubResult"`
		SubResNum        string        `json:"SubResNum"`
		ArPassthrough    []interface{} `json:"ar_passthrough"`
		RecoverCacheTime string        `json:"RecoverCacheTime"`
	} `json:"Result"`
	Data []struct {
		Srcid         string `json:"srcid"`
		Resourceid    string `json:"resourceid"`
		OriginQuery   string `json:"OriginQuery"`
		Origipquery   string `json:"origipquery"`
		Query         string `json:"query"`
		Origip        string `json:"origip"`
		Location      string `json:"location"`
		Userip        string `json:"userip"`
		Showlamp      string `json:"showlamp"`
		Tplt          string `json:"tplt"`
		Titlecont     string `json:"titlecont"`
		Realurl       string `json:"realurl"`
		ShowLikeShare string `json:"showLikeShare"`
		ShareImage    string `json:"shareImage"`
	} `json:"data"`
	ResultNum string `json:"ResultNum"`
}

func HttpIPApi(ip string)string{
	if len(ip) == 0 {
		return ""
	}

	url := fmt.Sprintf("https://sp1.baidu.com/8aQDcjqpAAV3otqbppnN2DJv/api.php?query=%s&resource_id=5809", ip)

	resp, err :=   http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return string(body)
}

// CheckIP 检测ip
func CheckIP(ip string)(ok bool)  {
	ipAddr := net.ParseIP(ip)
	if ipAddr != nil && strings.Contains(ip, "."){
		return true
	}
	ipAddr = net.ParseIP(ip)
	return ipAddr != nil && strings.Contains(ip, ":")
}

