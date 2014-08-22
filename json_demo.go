package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Analysis_results struct {
	Id               int64
	Sample_code      string `xorm:"varchar(50)"`
	Sample_sub_id    string `xorm:"varchar(50)"`
	Sample_date      time.Time
	Sample_name      string `xorm:"varchar(50)"`
	Sample_user_name string `xorm:"varchar(50)"`
	Sample_user_id   int64
	Mode             string `xorm:"varchar(50)"`
	Burn             string `xorm:"varchar(50)"`
	Element_name     string `xorm:"varchar(50)"`
	Element_code     string `xorm:"varchar(50)"`
	Element_value    string `xorm:"varchar(50)"`
	Element_unit     string `xorm:"varchar(20)"`
	Create_time      time.Time
	Create_user_code string `xorm:"varchar(50)"`
	Create_user_name string `xorm:"varchar(50)"`
	Is_in_use        bool
	Userd_by         string `xorm:"varchar(50)"`
	Remark           string `xorm:"varchar(50)"`
	Attribute1       string `xorm:"varchar(50)"`
	Attribute2       string `xorm:"varchar(50)"`
	Attribute3       string `xorm:"varchar(50)"`
	Attribute4       string `xorm:"varchar(50)"`
	Attribute5       string `xorm:"varchar(50)"`
}

func main() {
	str := `
{
  "attribute1": "11",
  "attribute2": "11",
  "attribute3": "22",
  "attribute4": "33",
  "attribute5": "11",
  "burn": "111",
  "create_time": "2014-08-21T20:00:00Z",
  "create_user_code": "aa",
  "create_user_name": "bb",
  "element_code": "aa",
  "element_name": "dd",
  "element_unit": "ee",
  "element_value": "ff",
  "is_in_use": true,
  "mode": "aa",
  "remark": "ff",
  "sample_code": "123",
  "sample_date": "2014-08-21T20:00:00Z",
  "sample_name": "test",
  "sample_sub_id": "test",
  "sample_user_id": 24,
  "sample_user_name": "test",
  "userd_by": "test"
}`
	var analysis Analysis_results
	err := json.Unmarshal([]byte(str), &analysis)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(analysis)
}
