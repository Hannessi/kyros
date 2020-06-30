package brokerInterface

type Broker interface {
	RetrieveOpenOrders(RetrieveOpenOrdersRequest) (*RetrieveOpenOrdersResponse, error)
	//RetrieveOpenTrades(RetrieveOpenTradesRequest) (*RetrieveOpenTradesResponse, error)
	//CreateNewOrder(CreateNewOrderRequest) (*CreateNewOrderResponse, error)
}

type RetrieveOpenOrdersRequest struct {
}

type RetrieveOpenOrdersResponse struct {
}

type RetrieveOpenTradesRequest struct {
}

type RetrieveOpenTradesResponse struct {
}

type CreateNewOrderRequest struct {
}

type CreateNewOrderResponse struct {
}
