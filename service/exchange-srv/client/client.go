package client

import (
	"wq-fotune-backend/libs/env"
	"wq-fotune-backend/libs/micro_client"
	pb "wq-fotune-backend/service/exchange-srv/proto"
	quoteCli "wq-fotune-backend/service/quote-srv/client"
	pbQuote "wq-fotune-backend/service/quote-srv/proto"
	userCli "wq-fotune-backend/service/user-srv/client"
	pbUser "wq-fotune-backend/service/user-srv/proto"
	walletCli "wq-fotune-backend/service/wallet-srv/client"
	pbWallet "wq-fotune-backend/service/wallet-srv/proto"
)

func NewExOrderClient(etcdAddr string) pb.ExOrderService {
	service := micro_client.InitBase(
		etcdAddr,
		//micro.Name("exchange-srv.client"),
	)
	exOrderService := pb.NewExOrderService(env.EXCHANGE_ORDER_SRV_NAME, service.Client())
	return exOrderService
}

func NewForwardOfferClient(etcdAddr string) pb.ForwardOfferService {
	service := micro_client.InitBase(
		etcdAddr,
	)
	exOrderService := pb.NewForwardOfferService(env.EXCHANGE_ORDER_SRV_NAME, service.Client())
	return exOrderService
}

var (
	QuoteService  pbQuote.QuoteService
	UserService   pbUser.UserService
	WalletService pbWallet.WalletService
)

func InitService() {
	QuoteService = quoteCli.NewQuoteClient(env.EtcdAddr)
	UserService = userCli.NewUserClient(env.EtcdAddr)
	WalletService = walletCli.NewWalletClient(env.EtcdAddr)
}
