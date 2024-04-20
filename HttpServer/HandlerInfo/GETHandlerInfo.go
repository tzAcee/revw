package HandlerInfo

type GETHandlerInfo struct {
	Params map[string]any
}

type GETHandleFunc func(GETHandlerInfo) (string, error)
