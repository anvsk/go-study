package main

import "fmt"

type DeviceDataHandler interface {
	Handle()
	Set_AlarmdData()
}

func (t *Base_devicedataHandler) Handle() {
	fmt.Print("default Handle\n")
}

func (base *Base_devicedataHandler) Set_AlarmdData() {
	println("default Set_AlarmdDta ")
}

type Base_devicedataHandler struct {
	Data DeviceData
}
type DeviceData struct {
	MainPumpSpeed     int     `json:"mainPumpSpeed"`     // 实时主贲
	RealMainWater     int     `json:"realMainWater"`     // 主缸水位
	RealTem           float32 `json:"realTem"`           // 主缸温度
	AuxiRealMainWater int     `json:"auxiRealMainWater"` // 副缸水位
	AuxiRealTem       float32 `json:"auxiRealTem"`       // 副缸温度
}

func main() {
	list := map[string]DeviceDataHandler{
		"hg":      NewHx(),
		"default": NewDefault(),
	}
	list["hg"].Handle()
	list["default"].Handle()
	println("========")
	list["hg"].Set_AlarmdData()
	list["default"].Set_AlarmdData()
}

func newHandler() *Base_devicedataHandler {
	return &Base_devicedataHandler{}
}

type Hx_devicedata_handler struct {
	*Base_devicedataHandler
}

type default_devicedata_handler struct {
	*Base_devicedataHandler
}

func NewDefault() DeviceDataHandler {
	return &default_devicedata_handler{
		Base_devicedataHandler: newHandler(),
	}
}

func NewHx() DeviceDataHandler {
	return &Hx_devicedata_handler{
		Base_devicedataHandler: newHandler(),
	}
}

func (base *Hx_devicedata_handler) Set_DeivceData() {

}

func (base *Hx_devicedata_handler) Set_AlarmdData() {
	println("Hx_devicedata_handler Set_AlarmdDta ")
}

func (base *Hx_devicedata_handler) Handle() {
	println("Hx_devicedata_handler Handle ")
}
