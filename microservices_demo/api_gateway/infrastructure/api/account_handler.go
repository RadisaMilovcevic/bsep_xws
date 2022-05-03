package api

import (
	"context"
	"encoding/json"
	"github.com/RadisaMilovcevic/bsep_xws/microservices_demo/api_gateway/domain"
	"github.com/RadisaMilovcevic/bsep_xws/microservices_demo/api_gateway/infrastructure/services"
	account "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common/proto/account_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type AccountHandler struct {
	accountClientAddress string
}

func NewAccountHandler(accountClientAddress string) Handler {
	return &AccountHandler{
		accountClientAddress: accountClientAddress,
	}
}

func (handler *AccountHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/account/account/{id}", handler.GetDetails)
	if err != nil {
		panic(err)
	}
}

func (handler *AccountHandler) GetDetails(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	accountDetails := &domain.Account{Id: id}
	accountClient := services.NewShippingClient(handler.accountClientAddress)
	accountInfo, err := accountClient.Get(context.TODO(), &account.GetRequest{Id: accountDetails.Id})
	response, err := json.Marshal(accountDetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
