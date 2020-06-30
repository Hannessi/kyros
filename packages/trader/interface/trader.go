package traderInterface

type Trader interface {
	Start(StartRequest) (*StartResponse, error)
	Stop(StopRequest) (*StopResponse, error)
}

type StartRequest struct {
}

type StartResponse struct {
}

type StopRequest struct {
}

type StopResponse struct {
}
