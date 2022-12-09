package main

import (
	"fmt"
	"net/http"
	"radiola.co.nz/babel/src/intake"
	"radiola.co.nz/babel/src/outtake"
	"radiola.co.nz/babel/src/service"
)

/** Main loop for this application**/
func main() {
	client := &http.Client{}                                                                                                                                                                                                                                                                                                                                                                             //Client Pointer
	fleetPinKey := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiI2MzRlMTc5YWQ5MTIzMjAwMDkyZTM3NGYiLCJvcmciOiJUSe-_vVx1MDAwNlx1MDAxNO-_ve-_ve-_vSrvv73vv71TIiwidiI6MCwiZmlyc3ROYW1lIjoiQVBJICIsImxhc3ROYW1lIjoiVXNlcjIiLCJlbWFpbCI6ImFwaXVzZXIyQGVtYWlsLmNvbSIsImlhdCI6MTY2NjA2MjMwMSwiYXVkIjoiZmxlZXRwaW4uY28ubnoiLCJpc3MiOiJhcHAuZmxlZXRwaW4uY28ubnoifQ.YtTeSvCPOzyRrzXZgsXZOd78hbRm1X4Iofx_f8_m-mU" //Our FleetPin Key
	//Intake
	fp := intake.NewFleetPinAPIWorker(fleetPinKey)                             //Generate our JWT
	res, err := fp.GetAssetByIdHttpRequest(client, "5a1dc7d15b4cf60a00b1bbd7") //Calling our assets
	if err != nil {
		fmt.Println("Broke at the HTTP request phase")
		fmt.Println(err.Error())
	}
	// transformative stage
	fps := service.NewFleetPinService()
	fpa, err := fps.FleetPinAssetConstructor(res)
	if err != nil {
		fmt.Println("Broke at the FleetPinAsset phase")
		fmt.Println(err.Error())
	}
	tce := fps.FleetPinAssetToTransitClockFormatConverter(fpa)
	//Outtake
	out := outtake.NewTransitClockOuttake("bop-api-dev.dynamis.live", 1234, "444b7db1", "bayofplenty")
	req, err := out.GenerateTransitClockRequest()
	out.GenerateURLParams(req, tce)
	if err != nil {
		fmt.Println("Broke at the Marshalling phase of Transit Clock.")
		fmt.Println(err.Error())
	}
	res, err = out.FlushDataToTransitClock(client, req)
	if err != nil {
		fmt.Println("Broke at the Flush phase of Transit Clock.")
		fmt.Println(err.Error())
	}
	fmt.Printf("\nResponse from TransitClock \n")
	fmt.Println(res)
}
