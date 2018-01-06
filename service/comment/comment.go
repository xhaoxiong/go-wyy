package comment

import (
	"fmt"
	"go-wyy/service/encrypt"
	"net/http"
	"net/url"
	"strings"
	"io/ioutil"
	"strconv"
	"go-wyy/models"
	"encoding/json"
)

/**
	id:歌曲id
	limit:每次请求条数
	offset:请求起始点
 */
func GetComments(id string, offset int, limit int) (comment *models.Comment, err error) {
	if id != "" {
		id = "404465600"
	}
	rid := ""
	strOffset := strconv.Itoa(offset)
	strLimit := strconv.Itoa(limit)
	total := "true"
	initStr1 := `{"rid": "` + rid + `","offset": "` + strOffset + `","total": "` + total + `","limit": "` + strLimit + `","csrf_token": ""}`
	params1, key1, err := encrypt.EncParams(initStr1)
	if err != nil {
		panic(err)
	}
	// 发送POST请求得到最后包含url的结果
	comment,err = Comments(params1, key1, id)




	if err != nil {
		fmt.Println(err)
		return comment, err
	}
	return comment, err
}

func GetAllComment(id string) (data interface{}, err error) {
	var comments []*models.Comments
	offset := 0
	for {
		data, err := GetComments(id, offset, offset+20)
		if err != nil {
			return data, err
		}
		//此处开启协程将数据存入数据库
		fmt.Println(data.Comments[0].User.NickName)
		fmt.Println(data.Total)
		if offset > int(data.Total) {
			break
		}
		offset+=20
	}
	return comments,err
}

func Comments(params string, encSecKey string, id string) (comment *models.Comment,err error) {
	client := &http.Client{}
	form := url.Values{}
	form.Set("params", params)
	form.Set("encSecKey", encSecKey)
	body := strings.NewReader(form.Encode())
	request, _ := http.NewRequest("POST", "http://music.163.com/weapi/v1/resource/comments/R_SO_4_"+id+"?csrf_token=", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Referer", "http://music.163.com")
	request.Header.Set("Content-Length", (string)(body.Len()))
	request.Header.Set("Cookie", "_ntes_nnid=f2c441d1440900d6daa9611bab3dc027,1515122355101; _ntes_nuid=f2c441d1440900d6daa9611bab3dc027; JSESSIONID-WYYY=Romq%5CHodpBq4TPCfWRRUHDzrkWgXCGH1ets%2FVAos2KQ2Yf76eNPz0g%2BKn2NnO9i%2F01IPS%2FWdgcu%2FfmswuAZRIrBo90IXKMTfVa%2F%2BjCt1e4jBagq9omzJ2fb7V72YXO2IR%2BKAEBgt90FvJCDe2I%2FfnGCEV3cxJ6aUf86E%5CQWZ3xzSoKid%3A1515145674403; _iuqxldmzr_=32; __utma=94650624.753127024.1515122355.1515135769.1515140831.4; __utmb=94650624.25.10.1515140831; __utmc=94650624; __utmz=94650624.1515122355.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none)")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.101 Safari/537.36")
	// 发起请求
	response, reqErr := client.Do(request)
	// 错误处理
	if reqErr != nil {
		fmt.Println("Fatal error ", reqErr.Error())
		return comment,reqErr
	}
	defer response.Body.Close()
	resBody, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(resBody, &comment)
	if err != nil {
		fmt.Println(err)
		return comment,err
	}
	return comment,nil

}
