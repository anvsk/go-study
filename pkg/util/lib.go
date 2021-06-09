package util

import (
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "errors"
    "math/rand"
    "net"
    "reflect"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
)

func InArray(need interface{}, haystack interface{}) (exists bool, index int) {
    exists = false
    index = -1
    switch reflect.TypeOf(haystack).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(haystack)
        for i := 0; i < s.Len(); i++ {
            if reflect.DeepEqual(need, s.Index(i).Interface()) == true {
                index = i
                exists = true
                return
            }
        }
    }
    return
}

// 类型互转，通过json编解码
func Type2type(from interface{}, to interface{}) {
    bytes, err := json.Marshal(from)
    if err != nil {
        return
    }
    json.Unmarshal(bytes, to)
}

// BindArgsWithGin 绑定请求参数
func BindArgsWithGin(c *gin.Context, req interface{}) error {
    return c.ShouldBindWith(req, binding.Default(c.Request.Method, c.ContentType()))
}

// MakeMD5 MD5加密
func MakeMD5(data string) string {
    h := md5.New()
    h.Write([]byte(data)) // 需要加密的字符串为 123456
    cipherStr := h.Sum(nil)
    return hex.EncodeToString(cipherStr) // 输出加密结果
}

// Random 生成随机数
func Random(min, max int) int {
    if min == max {
        return max
    }
    max = max + 1
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return min + r.Intn(max-min)
}

// RandomStr 随机字符串
func RandomStr(l int) string {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    seed := "1234567890QWERTYUIOPASDFGHJKLZXCVBNM"
    str := ""
    length := len(seed)
    for i := 0; i < l; i++ {
        point := r.Intn(length)
        str = str + seed[point:point+1]
    }
    return str
}

// BuildPassword 构建用户密码
func BuildPassword(password, salt string) string {
    return MakeMD5(password + salt)
}

// TernaryOperation 三元操作符
func TernaryOperation(exist bool, res, el interface{}) interface{} {
    if exist {
        return res
    }
    return el
}

// 获取机器ip
func GetLocalIP() (string, error) {
    var (
        addrs   []net.Addr
        addr    net.Addr
        err     error
        ipNet   *net.IPNet
        isIpNet bool
    )

    if addrs, err = net.InterfaceAddrs(); err != nil {
        return "", err
    }

    // 获取第一个非IO的网卡
    for _, addr = range addrs {
        // ipv4  ipv6
        // 如果能反解成ip地址 则为我们需要的地址
        if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
            // 是ip地址 不是 unix socket地址
            // 继续判断 是ipv4 还是 ipv6
            // 跳过ipv6
            if ipNet.IP.To4() != nil {
                return ipNet.IP.String(), nil
            }
        }
    }
    return "", errors.New("ip not found")
}

// StrArrExist 检测string数组中是否包含某个字符串
func StrArrExist(arr []string, check string) bool {
    for _, v := range arr {
        if v == check {
            return true
        }
    }
    return false
}

// RetryFunc 带重试的func
func RetryFunc(times int, f func() error) error {
    var (
        reTimes int
        err     error
    )
RETRY:
    if err = f(); err != nil {
        if reTimes == times {
            return err
        }
        time.Sleep(time.Duration(1) * time.Second)
        reTimes++
        goto RETRY
    }
    return nil
}
