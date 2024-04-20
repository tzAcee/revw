package HandlerInfo

type GETHandlerInfo struct {
	Params map[string]interface{}
}

type GETHandleFunc func(GETHandlerInfo) (string, error)
