package device_handler

// 对外统一接口方法
type DeviceDataHandlerInterface interface {
	Handle()
	SetAlarmdData()
}

// 父类
type BaseHandler struct {
	DeviceData
}

// 设备信息参数
type DeviceData struct {
	MainPumpSpeed     int     `json:"mainPumpSpeed"`     // 实时主贲
	RealMainWater     int     `json:"realMainWater"`     // 主缸水位
	RealTem           float32 `json:"realTem"`           // 主缸温度
	AuxiRealMainWater int     `json:"auxiRealMainWater"` // 副缸水位
	AuxiRealTem       float32 `json:"auxiRealTem"`       // 副缸温度
}

// 构造方法
func NewHandler(d DeviceData) *BaseHandler {
	return &BaseHandler{
		DeviceData: d,
	}
}

func (t *BaseHandler) Handle() {
	println("default Handle\n")
}

func (base *BaseHandler) SetAlarmdData() {
	println("default SetAlarmdData ")
}

func TestFunc() {
	list := map[string]DeviceDataHandlerInterface{
		"hg":      NewHuaxing(DeviceData{}),
		"default": NewDefault(DeviceData{}),
	}
	println("====Handle====")
	list["hg"].Handle()
	list["default"].Handle()
	println("====SetAlarmdData====")
	list["hg"].SetAlarmdData()
	list["default"].SetAlarmdData()
	println("====END====")

}
