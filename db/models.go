package db

import (
	"fmt"
	"log"
	"time"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	_ "github.com/lunny/godbc"
)

//明细表
type Tb_weigh_datalineinfo_detail struct {
	Id          int64
	BatchNumber int64     `xorm:"BatchNumber"`
	CarNumber   string    `xorm:"varchar(30) CarNumber"`
	Weight      float64   `xorm:'Weight'`
	WeighTime   time.Time `xorm:"DateTime WeighTime"`
	OperateBit  int64     `xorm:"OperateBit"`
	Remark      string    `xorm: "Remark"`
	Attribute1  string    `xorm:"varchar(100) Attribute1"`
	Attribute2  string    `xorm:"varchar(100) Attribute2"`
	Attribute3  string    `xorm:"varchar(100) Attribute3"`
	Attribute4  string    `xorm:"varchar(100) Attribute4"`
	Attribute5  string    `xorm:"varchar(100) Attribute5"`
}

var mssql *xorm.Engine
var pqorm *xorm.Engine

func NewDetail(batchNumber int64, weight float64) error {
	_, err := mssql.Insert(&Tb_weigh_datalineinfo_detail{BatchNumber: batchNumber, Weight: weight, WeighTime: time.Now(), OperateBit: 1})
	return err
}

func init() {
	fmt.Println("main init")
	//创建mssql orm引擎
	var err error
	mssql, err = xorm.NewEngine("odbc", "driver={sql server};server=127.0.0.1;port=1433;uid=sa;pwd=tlys.oaxmz.5860247;database=tgzljl_czgl")
	if err != nil {
		log.Fatalf("err happened", err)
	}

	//创建pq orm引擎
	pqorm, err = xorm.NewEngine("postgres", "host=127.0.0.1 port=5432 user=zljl password=123 dbname=tgzljl_czgl sslmode=disable")
	if err != nil {
		log.Fatalln("pq orm created error:", err)
	}
	err = pqorm.Sync(new(Tb_weigh_datalineinfo_detail))
	if err != nil {
		log.Fatalln("pq orm sync error:", err)
	}
}
