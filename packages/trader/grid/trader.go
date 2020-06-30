package traderGrid

import (
	"errors"
	"fmt"
	"github.com/hannessi/kyros/log"
	brokerInterface "github.com/hannessi/kyros/packages/broker/interface"
	traderInterface "github.com/hannessi/kyros/packages/trader/interface"
	"time"
)

func New(request CreateNewRequest) (*Trader, error) {
	trader := Trader{
		Broker: request.Broker,
		grid:   createGrid(10000,0,500), // todo create config
		isActive: false,
	}

	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for range ticker.C {
			if trader.isActive {
				if err := trader.timedRun(); err != nil {
					log.Error("could not complete Trader.timedRun: "+ err.Error())
				}
			}
		}
	}()

	if err := trader.validate(); err != nil {
		log.Error("could not validate trader: "+err.Error())
		return nil, errors.New("could not validate trader: "+err.Error())
	}

	return &trader, nil
}

type CreateNewRequest struct {
	Broker *brokerInterface.Broker
}

type Trader struct {
	Broker *brokerInterface.Broker
	grid   *grid
	isActive bool
}

func (t *Trader) validate() error {
	//if t.Broker == nil {
	//	return errors.New("Trader.Broker is not valid")
	//}

	if t.grid == nil {
		return errors.New("Trader.grid is not valid")
	}

	return nil
}

func (t *Trader) Start(request traderInterface.StartRequest) (*traderInterface.StartResponse, error) {
	t.isActive = true
	return &traderInterface.StartResponse{}, nil
}

func (t *Trader) Stop(request traderInterface.StopRequest) (*traderInterface.StopResponse, error) {
	t.isActive = false
	return &traderInterface.StopResponse{}, nil
}

func (t *Trader) timedRun() error {
	log.Debug(t.grid.ToString())
	t.grid.clear()
	t.setOpenOrders()
	t.setOpenTrades()
	t.createNewOrder()
	return errors.New("not implemented yet")
}

func (t *Trader) setOpenOrders() error {
	// get open order using broker

	// loop through grid and populate the grid
}

func (t *Trader) setOpenTrades() error {
	// get open trader using broker

	// loop through grid and populate the grid
}

func (t *Trader) createNewOrder() error {
	// check is order/trade has not been set in enveloping levels

	// create order if needed.
}


func createGrid(max, min, step int) *grid {
	levels := make([]level, 0)

	current := max

	for current > min {
		levels = append(levels, level{
			price:    float64(current),
			hasOrder: false,
			hasTrade: false,
		})

		current -= step
	}

	return &grid{
		levels: levels,
	}
}

type grid struct {
	levels []level
}

type level struct {
	price    float64
	hasOrder bool
	hasTrade bool
}

func (g *grid) ToString() string {
	s := "\n"

	for i, level := range g.levels {
		s = s + fmt.Sprintf("Level %d: %6.4f - t:%v - o:%v\n",i, level.price, level.hasTrade, level.hasOrder )
	}
	return s
}

func (g *grid) clear() {
	for i := range g.levels {
		g.levels[i].hasOrder = false
		g.levels[i].hasTrade = false
	}
}
