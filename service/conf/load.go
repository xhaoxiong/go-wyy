package conf

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type DbConf struct {
	DbHost string `json:"db_host"`
	DbPort string `json:"db_port"`
	DbUser string `json:"db_user"`
	DbPass string `json:"db_pass"`
	DbName string `json:"db_name"`
}

func (this *DbConf) Load(filename string,) (dbconf *DbConf,err error) {
	var dbconff *DbConf
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("read file error")
		return dbconf,err
	}
	datajson := []byte(data)

	err = json.Unmarshal(datajson, &dbconff)

	if err != nil {

		return dbconff, err

	}
	return dbconff,nil
}



