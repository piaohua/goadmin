package libs

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

// UcFirst ...
func UcFirst(str string) string {
	if str == "" {
		return ""
	}
	bytes := []byte(str)
	bytes[0] = []byte(strings.ToUpper(string(bytes[0])))[0]
	return string(bytes)
}

// MD5 ...
func MD5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

// TodayTime today time
func TodayTime() time.Time {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	return today
}

// ExecCmd 执行shell命令
func ExecCmd(c string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", c)
	out, err := cmd.Output()
	return out, err
}

//Load load info by json
func Load(filePath string, v interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// Format 跟 PHP 中 date 类似的使用方式，如果 ts 没传递，则使用当前时间
func Format(format string, ts ...time.Time) string {
	patterns := []string{
		// 年
		"Y", "2006", // 4 位数字完整表示的年份
		"y", "06", // 2 位数字表示的年份

		// 月
		"m", "01", // 数字表示的月份，有前导零
		"n", "1", // 数字表示的月份，没有前导零
		"M", "Jan", // 三个字母缩写表示的月份
		"F", "January", // 月份，完整的文本格式，例如 January 或者 March

		// 日
		"d", "02", // 月份中的第几天，有前导零的 2 位数字
		"j", "2", // 月份中的第几天，没有前导零

		"D", "Mon", // 星期几，文本表示，3 个字母
		"l", "Monday", // 星期几，完整的文本格式;L的小写字母

		// 时间
		"g", "3", // 小时，12 小时格式，没有前导零
		"G", "15", // 小时，24 小时格式，没有前导零
		"h", "03", // 小时，12 小时格式，有前导零
		"H", "15", // 小时，24 小时格式，有前导零

		"a", "pm", // 小写的上午和下午值
		"A", "PM", // 小写的上午和下午值

		"i", "04", // 有前导零的分钟数
		"s", "05", // 秒数，有前导零
	}
	replacer := strings.NewReplacer(patterns...)
	format = replacer.Replace(format)

	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}
	return t.Format(format)
}

// StrToLocalTime  字符串转换成本地Time时间表示
func StrToLocalTime(value string) time.Time {
	if value == "" {
		return time.Time{}
	}
	zoneName, offset := time.Now().Zone()

	zoneValue := offset / 3600 * 100
	if zoneValue > 0 {
		value += fmt.Sprintf(" +%04d", zoneValue)
	} else {
		value += fmt.Sprintf(" -%04d", zoneValue)
	}

	if zoneName != "" {
		value += " " + zoneName
	}
	return StrToTime(value)
}

// StrToTime 字符串转换成Time时间表示
func StrToTime(value string) time.Time {
	if value == "" {
		return time.Time{}
	}
	layouts := []string{
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05 -0700 MST",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02 -0700 MST",
		"2006-01-02 -0700",
		"2006-01-02",
		"2006/01/02 -0700 MST",
		"2006/01/02 -0700",
		"2006/01/02",
		"2006-01-02 15:04:05 -0700 -0700",
		"2006/01/02 15:04:05 -0700 -0700",
		"2006-01-02 -0700 -0700",
		"2006/01/02 -0700 -0700",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			return t
		}
	}
	panic(err)
}

// Time2Str 返回字符串格式化时间
func Time2Str(t time.Time) string {
	return Format("2006-01-02 15:04:05", t)
}

// Time2LocalStr 返回字符串格式化本地时间
func Time2LocalStr(t time.Time) string {
	return Format("2006-01-02 15:04:05", t.Local())
}

// TodayStr 返回字符串格式化本地时间
func TodayStr(day int) string {
	t := time.Now().AddDate(0, 0, day)
	return Format("01/02/2006", t.Local())
}

// CurrTimeStr 返回当前时间字符串
func CurrTimeStr() string {
	return time.Now().Local().Format("20060102150405")
}

//Today2Str today
func Today2Str() (start, end string) {
	t := time.Now().Local()
	todayStart := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	todayEnd := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Local)
	start = todayStart.Format("2006-01-02 15:04:05")
	end = todayEnd.Format("2006-01-02 15:04:05")
	return
}

// Str2Local str格式化时间转本地时间戳
func Str2Local(t string) (int64, error) {
	the_time, err := time.ParseInLocation("01/02/2006", t, time.Local)
	if err == nil {
		return the_time.Unix(), err
	}
	return 0, err
}

// Datetime2Str 指定日期开始结束时间
func Datetime2Str(t string) (start, end string) {
	the_time, err := time.ParseInLocation("01/02/2006", t, time.Local)
	if err != nil {
		return
	}
	start = Format("01/02/2006", the_time.Local())
	end = Format("01/02/2006", the_time.AddDate(0, 0, 1).Local())
	return
}

// Datetime2Format 指定日期开始结束时间
func Datetime2Format(t string) (start, end string) {
	the_time, err := time.ParseInLocation("01/02/2006", t, time.Local)
	if err != nil {
		return
	}
	start = the_time.Format("2006-01-02 15:04:05")
	end = the_time.AddDate(0, 0, 1).Format("2006-01-02 15:04:05")
	return
}

// Datetime2Day 指定日期数值格式
func Datetime2Day(t string) (today int) {
	the_time, err := time.ParseInLocation("01/02/2006", t, time.Local)
	if err != nil {
		return
	}
	today = Time2DayDate(the_time)
	return
}

// Str2LocalTime str格式化时间转本地时间
func Str2LocalTime(s string) (t time.Time, err error) {
	t, err = time.ParseInLocation("01/02/2006", s, time.Local)
	return
}

// Stamp2Time 时间截转换为Time时间
func Stamp2Time(t int64) time.Time {
	return time.Unix(t, 0)
}

//Time2DayDate 时间yyyymmdd
func Time2DayDate(t time.Time) int {
	return t.Year()*10000 + int(t.Month())*100 + t.Day()
}

//PathExists 文件或目录路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//WriteFile json文件写入文件
func WriteFile(filePath string, i interface{}) error {
	data, err := json.Marshal(i)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
