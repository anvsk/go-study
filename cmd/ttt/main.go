package main

type DeviceDataHandler interface {
	Handle()
}

func (t *Base_devicedataHandler) Handle() {
	// fmt.Print("prepare downloading\n")
	t.IBase_devicedataHandler.Set_AlarmdData()
	// fmt.Print("finish downloading\n")
}

type Base_devicedataHandler struct {
	Data                    DeviceData
	IBase_devicedataHandler //数据处理方法
}
type DeviceData struct {
	MainPumpSpeed     int     `json:"mainPumpSpeed"`     // 实时主贲
	RealMainWater     int     `json:"realMainWater"`     // 主缸水位
	RealTem           float32 `json:"realTem"`           // 主缸温度
	AuxiRealMainWater int     `json:"auxiRealMainWater"` // 副缸水位
	AuxiRealTem       float32 `json:"auxiRealTem"`       // 副缸温度
}

type IBase_devicedataHandler interface {
	Set_AlarmdData()
}

func (base *Base_devicedataHandler) Set_AlarmdData() {
	println("default Set_AlarmdDta ")
}

func main() {
	// var list map[string]Base_devicedataHandler
	// aa := NewHx()
	// fmt.Println(aa)
	list := map[string]DeviceDataHandler{
		"hg":      NewHx(),
		"default": NewDefault(),
	}
	// list["hg"].IBase_devicedataHandler.Set_AlarmdData()
	// list["default"].IBase_devicedataHandler.Set_AlarmdData()
	list["hg"].Handle()
	list["default"].Handle()
}

func newHandler(impl IBase_devicedataHandler) *Base_devicedataHandler {
	return &Base_devicedataHandler{
		IBase_devicedataHandler: impl,
	}
}

type Hx_devicedata_handler struct {
	*Base_devicedataHandler
}

type default_devicedata_handler struct {
	*Base_devicedataHandler
}

func NewDefault() DeviceDataHandler {
	downloader := &default_devicedata_handler{}
	template := newHandler(downloader)
	downloader.Base_devicedataHandler = template
	return downloader
}

func NewHx() DeviceDataHandler {
	downloader := &Hx_devicedata_handler{}
	template := newHandler(downloader)
	downloader.Base_devicedataHandler = template
	return downloader
}

func (base *Hx_devicedata_handler) Set_DeivceData() {

}

func (base *Hx_devicedata_handler) Set_AlarmdData() {
	println("Hx_devicedata_handler Set_AlarmdDta ")
}
