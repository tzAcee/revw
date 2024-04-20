package HandlerInfo

type PUTHandlerInfo struct {
	Params map[string]any
}

type PUTHandleFunc func(PUTHandlerInfo) (string, error)
