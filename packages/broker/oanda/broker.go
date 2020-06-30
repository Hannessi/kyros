package brokerOanda

import (
	"github.com/hannessi/gOanda"
	brokerInterface "github.com/hannessi/kyros/packages/broker/interface"
)

type Broker struct {
	OandaClient gOanda.Client
}

func (b *Broker) RetrieveOpenOrders(request brokerInterface.RetrieveOpenOrdersRequest) (*brokerInterface.RetrieveOpenOrdersResponse, error) {

}
