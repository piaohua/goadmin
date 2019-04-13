package models

import (
	"testing"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func LoadOrm(aliasName string) {
	beego.LoadAppConfig("ini", "../conf/app.conf")
	driverName := "mysql"
	logfile := beego.AppConfig.String(driverName + ".logfile")
	beego.AppConfig.Set((driverName + ".logfile"), ("../" + logfile))
	initOrm(aliasName)
	err := orm.RunSyncdb(aliasName, false, true)
	if err != nil {
		panic(err)
	}
}

func TestLogger(t *testing.T) {
	LoadOrm("default")
	//beego.LoadAppConfig("ini", "../conf/app.conf")
	//driverName := "mysql"
	//logfile := beego.AppConfig.String(driverName + ".logfile")
	//beego.AppConfig.Set((driverName + ".logfile"), ("../" + logfile))
	//aliasName := "default"
	//initOrm(aliasName)
	//err := orm.RunSyncdb(aliasName, false, true)
	//if err != nil {
	//	t.Errorf("sync db err %v", err)
	//}

	module := &StatSendClick{
		Timestamp:   int(time.Now().Unix()),
		Cover:       1,
		Scene:       1,
		Version:     1,
		SendNumber:  1,
		SendTimes:   1,
		ClickNumber: 1,
		ClickTimes:  1,
		Regist:      1,
		Login:       1,
	}
	t.Logf("module %#v", module)
	time.Sleep(3 * time.Second)
	o := orm.NewOrm()
	//num, err := o.Update(module)
	//t.Logf("update num %d", num)
	//if err != nil {
	//	t.Fatalf("update err : %v", err)
	//}
	id, err := o.Insert(module)
	t.Logf("insert id %d", id)
	if err != nil {
		t.Fatalf("insert err : %v", err)
	}
	time.Sleep(3 * time.Second)
}

// [INSERT INTO `log_send_clieck_stat` (`timestamp`, `cover`, `scene`, `version`, `send_number`, `send_times`, `click_number`, `click_times`, `regist`, `login`, `created_at`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)] - `1539240112`, `1`, `1`, `1`, `1`, `1`, `1`, `1`, `1`, `1`, `2018-10-11 14:41:55.3868347 +0800 CST`

// TestReportActive ...
func TestReportActive(t *testing.T) {
	LoadOrm("default")
	orm.ResetModelCache()
	LoadOrm("report")
	module := &StatActive{
		ProjectId: 11,
		Dau:       1,
		RegMan:    1,
		LogMan:    1,
		BackMan:   1,
		Left2:     1,
		Log15:     1,
	}
	t.Logf("module %#v", module)
	time.Sleep(3 * time.Second)
	o := orm.NewOrm()
	//num, err := o.Update(module)
	//t.Logf("update num %d", num)
	//if err != nil {
	//	t.Fatalf("update err : %v", err)
	//}
	id, err := o.Insert(module)
	t.Logf("insert id %d", id)
	if err != nil {
		t.Fatalf("insert err : %v", err)
	}
	time.Sleep(3 * time.Second)
}

// TestChat ...
func TestChat(t *testing.T) {
	LoadOrm("default")
	chat := SystemChat{
		UserId:   1,
		UserName: "admin",
		Message:  "hello world!",
		//CreatedAt: time.Now(),
	}
	t.Logf("chat %#v", chat)
	time.Sleep(3 * time.Second)
	ok, err := chat.Save()
	t.Logf("ok %v, err %v", ok, err)
	time.Sleep(3 * time.Second)
}
