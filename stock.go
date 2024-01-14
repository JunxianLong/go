//基本的GET请求
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Data       *Data  `json:"data"`
}

type Data struct {
	Infos []*Info `json:"info"`
}

func newData() *Data {
	return &Data{
		Infos: make([]*Info, 0),
	}
}

type Info struct {
	Code       string `json:"code"`        //:"603268"
	Name       string `json:"name"`        //:"松发股份"
	ReasonType string `json:"reason_type"` //"储能+陶瓷产品"
}

func NewResponse() *Response {
	return &Response{
		Data: newData(),
	}
}

func main() {

	path := "stock.txt"
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)

	year := "2023"
	month := "12"
	dates := make([]string, 0)
	//days := []string{"01", "02", "03", "04", "05", "06", "07", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}
	days := []string{"08"}
	breaks := map[string]struct{}{
		//"01": struct{}{},
		//"06": struct{}{},
		//"07": struct{}{},
		//"10": struct{}{},
		//"10": struct{}{},
		//"10": struct{}{},
		//"10": struct{}{},
		//"10": struct{}{},
		//"10": struct{}{},
		//"10": struct{}{},
		//"10": struct{}{},
		//"10": struct{}{},
	}
	for _, day := range days {
		if _, ok := breaks[day]; ok {
			continue
		}
		dates = append(dates, fmt.Sprintf("%s%s%s", year, month, day))
	}
	for _, date := range dates {
		write.WriteString(fmt.Sprintf("日期:%s\n", date))
		url := fmt.Sprintf("https://data.10jqka.com.cn/dataapi/limit_up/limit_up_pool?page=1&limit=90&field=199112,10,9001,330323,330324,330325,9002,330329,133971,133970,1968584,3475914,9003,9004&filter=HS,GEM2STAR&date=%s&order_field=330324&order_type=0&_=1705160494666", date)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		resp.Body.Close()
		res := NewResponse()
		if err := json.Unmarshal(body, res); err != nil {
			panic(err)
		}
		for i, v := range res.Data.Infos {
			content := fmt.Sprintf("id:%d,股票编号:%s,股票名称:%s,涨停原因:%s\n", i+1, v.Code, v.Name, v.ReasonType)
			write.WriteString(content)
		}
		write.Flush()
		fmt.Printf("日期:%s,涨停个数:%d\n", date, len(res.Data.Infos))
	}
}
