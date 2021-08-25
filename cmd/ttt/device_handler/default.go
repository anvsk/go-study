package device_handler

type DefaultHandler struct {
	*BaseHandler
}

func NewDefault(d DeviceData) DeviceDataHandlerInterface {
	return &DefaultHandler{
		BaseHandler: NewHandler(d),
	}
}
