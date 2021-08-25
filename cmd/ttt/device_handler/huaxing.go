package device_handler

type HuaxingHandler struct {
	*BaseHandler
}

func NewHuaxing(d DeviceData) DeviceDataHandlerInterface {
	return &HuaxingHandler{
		BaseHandler: NewHandler(d),
	}
}

func (h *HuaxingHandler) SetAlarmdData() {
	println("HuaxingHandler SetAlarmdData ")
}
