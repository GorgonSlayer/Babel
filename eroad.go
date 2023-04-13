package main

import (
	"net/http"

	"radiola.co.nz/babel/src/intake"
	"radiola.co.nz/babel/src/service"
	"radiola.co.nz/babel/src/util/logger"
)

/** Testing E-Road process **/
func ERoadLoop() {
	client := http.Client{}
	l := logger.NewLogger(false, "eroad.log")
	eri := intake.NewERoadAPI(
		"InsertYourKeysHere",
		"https://api.eroad.co.nz",
		l,
	)
	res, _ := eri.RetrieveIntrospect(&client)
	ers := service.NewERoadService(l)
	ers.ProcessIntrospection(res)
}

func main() {
	ERoadLoop()
}
