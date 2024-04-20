package HandlerInfo

type POSTHandlerInfo struct {
	Params      map[string]interface{}
	RequestBody map[string]interface{}
}

type POSTHandleFunc func(POSTHandlerInfo) (string, error)
