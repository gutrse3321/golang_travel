package main
import (
    "fmt"
    "time"
)

var week time.Duration

func main() {
    // 获取当天时间
    // go1.0.3 e.g. Wed Dec 21 09:52:14 +0100 RST 2011
    // shell: 2018-09-26 15:16:46.784968228 +0000 UTC
    t := time.Now()
    fmt.Println(t)

    // 获取当天时间
    // shell: 2018-09-26 15:16:46.784968228 +0000 UTC
    // 貌似1.11版本直接使用UTC格式了？？没去查明
    t = time.Now().UTC()
    fmt.Println(t)

    // 计算次数
    // shell: 2018-09-26 15:16:46.784968228 +0000 UTC
    week = 60 * 60 * 24 * 7 * 1e9
    week_from_now := t.Add(week)
    fmt.Println(week_from_now)

    // 格式化时间
    // shell: 26 Sep 18 15:24 UTC
    fmt.Println(t.Format(time.RFC822))

    // shell: Wed Sep 26 15:25:32 2018
    fmt.Println(t.Format(time.ANSIC))

    // 其实仅仅是使用这个格式，下面那个列子也是
    // shell: 26 Sep 2018 15:26
    fmt.Println(t.Format("02 Jan 2006 15:04"))
    // shell: 2018-09-26 15:28:02.791779809 +0000 UTC => 20180926
    s := t.Format("20060102")
    fmt.Println(t, "=>", s)
}