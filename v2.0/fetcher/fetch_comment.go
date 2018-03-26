package fetcher

import (
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"go-wyy/service/encrypt"
)


func GetComments(id string, offset int, limit int)([]byte, error) {

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
	return FetchComments(params1, key1, id)

}




func FetchComments(params string, encSecKey string, id string) ([]byte, error) {
	client := &http.Client{}
	form := url.Values{}
	form.Set("params", params)
	form.Set("encSecKey", encSecKey)
	body := strings.NewReader(form.Encode())
	request, _ := http.NewRequest("POST", "http://music.163.com/weapi/v1/resource/comments/R_SO_4_"+id+"?csrf_token=", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Referer", "http://music.163.com")
	request.Header.Set("Content-Length", (string)(body.Len()))
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.101 Safari/537.36")
	request.Header.Set("Cookie","_ntes_nnid=f2c441d1440900d6daa9611bab3dc027,1515122355101; _ntes_nuid=f2c441d1440900d6daa9611bab3dc027; __utmz=94650624.1515122355.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); _iuqxldmzr_=32; __remember_me=true; JSESSIONID-WYYY=YXZtk7tOBJ4b3gOVrX2hl5%2BBriZyYVR5kNX3D3G5oWFRcY3J1cvGnMJRZx6JXgVSRNhFKO3O%5CmRiRACwWjrhBnkmK3dgGyTawDSAAmF%2Fct5T%2BhYVRy1BnxCgx%5CYrAUrjnQ8jEJQ1VHJTdNhqS4p9jVxHdRcc7iv5cQn649a%5CsBTc46WR%3A1515402120148; __utma=94650624.753127024.1515122355.1515218466.1515400320.9; __utmc=94650624; MUSIC_U=0120a3f48157438f759f2034b3925668ae731f8ae462a842927650798e0d663c97d1b459676c0cc693e926b3c390b8ba205ba14613b02d6c02d1ccf53040f6087d9739a0cccfd7eebf122d59fa1ed6a2; __csrf=5aa926378397ed694496ebf6486c5dfc; __utmb=94650624.5.10.1515400320")
	// 发起请求
	response, reqErr := client.Do(request)
	// 错误处理
	if reqErr != nil {
		fmt.Println("Fatal error ", reqErr.Error())
		return nil, reqErr
	}
	defer response.Body.Close()

	if response.StatusCode!=http.StatusOK{
		fmt.Println("Error:status code", response.StatusCode)
		return nil, fmt.Errorf("wrong status code:%d", response.StatusCode)
	}


	return ioutil.ReadAll(response.Body)
}