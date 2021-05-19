package util

import (
    "flag"
    "io/ioutil"
    "os"

    "gopkg.in/yaml.v2"
)

var Config Cf
var Args Customization

type Cf struct {
    Env          string `yaml:"Env"`
    Debug        bool   `yaml:"Debug"`
    Store        `yaml:"Store"`
    TicketConfig `yaml:"TicketConfig"`
    Notice       `yaml:"Notice"`
}

type Store struct {
    Cache StoreConfig `json:"Cache" yaml:"Cache"`
    DB    StoreConfig `json:"DB" yaml:"DB"`
}

type StoreConfig struct {
    DefaultDriver string             `json:"DefaultDriver" yaml:"DefaultDriver"`
    Driver        []string           `json:"Driver" yaml:"Driver"`
    Connects      map[string]Connect `json:"Connects" yaml:"Connects"`
}

type Connect struct {
    Host     string `json:"Host" yaml:"Host"`
    Port     int    `json:"Port" yaml:"Port"`
    User     string `json:"User" yaml:"User"`
    Password string `json:"Password" yaml:"Password"`
    DbIndex  int    `json:"DbIndex" yaml:"DbIndex"`
    DbName   string `json:"DbName" yaml:"DbName"`
}

type TicketConfig struct {
    Debug         bool `yaml:"Debug"`
    User          `yaml:"User"`
    Customization `yaml:"Customization"`
    PortNo        map[string]int `yaml:"PortNo"`
}

type Notice struct {
    Flag      bool   `yaml:"Flag"`
    CompanyWx string `yaml:"CompanyWx"`
}

type Customization struct {
    From           string `yaml:"From"`
    To             string `yaml:"To"`
    Date           string `yaml:"Date"`
    MinShipTime    string `yaml:"MinShipTime"`
    LatestBusTime  string `yaml:"LatestBusTime"`
    LatestShipTime string `yaml:"LatestShipTime"`
    LineNum        string `yaml:"LineNum"`
    Class          string `yaml:"Class"`
}

type User struct {
    Mobile         string
    Password       string
    Authentication string
    Passengers     []string
}

func initConfig() {
    var configFilePath string
    flag.StringVar(&configFilePath, "config", "config.yml", "config file location")
    flag.Parse()

    file, err := os.Open(configFilePath)

    if err != nil {
        panic(err)
    }

    bytes, err := ioutil.ReadAll(file)

    if err != nil {
        panic(err)
    }

    cfg := Cf{}

    err = yaml.Unmarshal(bytes, &cfg)

    if err != nil {
        panic(err)
    }

    Config = cfg
}
