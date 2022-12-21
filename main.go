package main

import (
	"radiola.co.nz/babel/src/overseer"
	"radiola.co.nz/babel/src/util/logger"
)

/** This function should initialise the application and ensure the overseer loop kicks off. **/
func initialise(logName string) {
	l := logger.NewLogger(false, logName)
	overseer := overseer.NewOverseer(l)
}

/** Main loop for this application**/
func main() {

}
