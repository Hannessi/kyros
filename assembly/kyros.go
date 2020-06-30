package main

import (
	"github.com/hannessi/kyros/log"
	traderGrid "github.com/hannessi/kyros/packages/trader/grid"
	traderInterface "github.com/hannessi/kyros/packages/trader/interface"
	"os"
	"os/signal"
)

func main() {
	version := "0.0.1"
	log.Info("Starting Kyros v"+version)
	gridTrader, err := traderGrid.New(traderGrid.CreateNewRequest{Broker:nil})
	if err != nil {
		log.Error("could not create trader: "+err.Error())
	}

	_, err = gridTrader.Start(traderInterface.StartRequest{})
	if err != nil {
		log.Error("could not start trader: "+err.Error())
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			log.Info("Kyros is shutting down..")
			return
		}
	}
}
