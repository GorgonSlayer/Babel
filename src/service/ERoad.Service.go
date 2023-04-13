package service

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"
	"radiola.co.nz/babel/src/model/intakeResponse"
	"radiola.co.nz/babel/src/util/logger"
)

/** Interface for external access **/
type IERoadService interface {
	ProcessIntrospection(res *http.Response) (intakeResponse.ERoadIntrospectAsset, error)
}

/** State stored for ERoad service **/
type ERoadService struct {
	logger logger.Logger
}

/** Constructor **/
func NewERoadService(l logger.Logger) IERoadService {
	return ERoadService{
		logger: l,
	}
}

/** This service process an HTTP response, filters the Machines out and returns structs of ERoadIntrospectAssets. **/
func (ers ERoadService) ProcessIntrospection(res *http.Response) (intakeResponse.ERoadIntrospectAsset, error) {
	var eria intakeResponse.ERoadIntrospectAsset
	if res.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(res.Body)
		err := decoder.Decode(&eria)
		if err != nil {
			ers.logger.Zap.Error("Error during ProcessIntrospection json unmarshalling", zap.Any("Error", err))
		}
		ers.logger.Zap.Debug("Content of Introspect Call", zap.Any("Introspect", eria))
		return eria, err
	}
	return eria, errors.New("an error occurred in process introspection (eroad), the response status code was not 200")
}

/** Reads the struct interface of strings to return machine (vehicle) entities **/
func (ers ERoadService) FilterMachines() []intakeResponse.ERoadIntrospectAsset {
	return nil
}
