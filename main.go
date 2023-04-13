package main

import (
	"time"

	"radiola.co.nz/babel/src/config"
	"radiola.co.nz/babel/src/overseer"
	"radiola.co.nz/babel/src/util"
	"radiola.co.nz/babel/src/util/logger"
	"radiola.co.nz/babel/src/worker"
)

/** This function should initialise the application and ensure the overseer loop kicks off. **/
func initialise(logName string) {
	l := logger.NewLogger(false, logName)
	fpc := config.GetConfigFromFile(l)
	o := overseer.NewOverseer(l)
	fpw := worker.NewFleetPinWorker(fpc.Jwt, fpc.FleetPinUrl, fpc.Host, fpc.Port, fpc.Key, fpc.Agency, l)
	item := util.Item{
		Priority: time.Now().Unix(),
		Index:    0,
		Worker:   fpw,
	}
	o.Push(&item)
	o.OverseerLoop()
}

/** Main loop for this application**/
/**
func main() {
	//l := logger.NewLogger(false, "output.log")
	initialise("output.log")
	//newMain(l)
}
**/
