package viper_test

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/baetyl/baetyl-go/log"
	"github.com/fsnotify/fsnotify"
	logg "github.com/pieterclaerhout/go-log"
	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	cfg := Config{}
	Viper(&cfg, "collector_local.yaml")
	// logg.InfoDump(cfg, "cfg")
	// fmt.Println(v.AllKeys())
	// fmt.Println(v.Get("tx"))
	// time.Sleep(5 * time.Second)
	// logg.InfoDump(cfg, "cfg")

}

// 优先级: 命令行 > 环境变量 > 默认值
// defaultPath 默认路径
// assignObject 赋值对象
func Viper(assignObject interface{}, defaultPath ...string) *viper.Viper {
	var config string
	flag.StringVar(&config, "f", "", "choose config file.")
	flag.Parse()
	if config == "" {
		// fmt.Printf("您正在使用config的默认值,config的路径为%v\n", "./collector_local.yaml")
		if configEnv := os.Getenv("COLLECTOR_CONFIG"); configEnv == "" {
			config = defaultPath[0]
			fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config)
		} else {
			config = configEnv
			fmt.Printf("您正在使用COLLECTOR_CONFIG环境变量,config的路径为%v\n", config)
		}
	} else {
		fmt.Printf("您正在使用命令行的-f参数传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(assignObject); err != nil {
			fmt.Println(err)
		}
	})
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		fmt.Println(err)
	}
	logg.InfoDump(cfg, "cfg")
	return v
}

type Config struct {
	Zaplog   log.Config `yaml:"zaplog" json:"zaplog"`
	TxConfig `yaml:"tx" mapstructure:"tx"`
}

type TxConfig struct {
	Address               string `yaml:"address" json:"address"`
	IntervalSendDuration  string `yaml:"interval-send-duration" json:"interval-send-duration"`
	IntervalSendDuration2 string `yaml:"interval_send_duration"`
}
