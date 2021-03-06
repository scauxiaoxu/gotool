package time

/**
  补充时间换算规则： 秒 > 毫秒 > 微妙 > 纳秒 > ..
  1s=10^3ms(毫秒)=10^6μs(微秒)=10^9ns(纳秒)=10^12ps(皮秒)=10^15fs(飞秒)=10^18as(阿秒)=10^21zm(仄秒)=10^24ym(幺秒)
*/

import (
	"time"
)

const Layout = "2006-01-02 15:04:05"

// time.  类似php time()函数
func Time() int64 {
	return time.Now().Unix()
}

// 返回当前时间微妙  有点类似php microtime()
func MicroTime() float64 {
	micro := time.Now().UnixNano() / 1e3
	timeF := float64(micro)
	return timeF / 1000000
}

// 类似PHP strtotime()
// 暂时只支持TIME_LAYOUT格式转化 通过err判断转化是否成功
func Strtotime(layout string) (int64, error) {
	var unixTime int64
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation(Layout, layout, loc)
	if err != nil {
		return unixTime, err
	}
	unixTime = theTime.Unix()
	return unixTime, err
}

// 类似PHP date()
func Date(format string, unixTime int64) string {
	//TODO 指定时区 time.LoadLocation windows系统需要配置go环境 否则可能会有问题
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	tm := time.Unix(unixTime, 0).In(cstSh)
	return tm.Format(format)
}
