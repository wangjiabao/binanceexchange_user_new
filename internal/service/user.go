package service

import (
	abi "binanceexchange_user/abi"
	v1 "binanceexchange_user/api/binanceexchange_user/v1"
	"binanceexchange_user/internal/biz"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strconv"
)

// BinanceUserService is a BinanceData service .
type BinanceUserService struct {
	v1.UnimplementedBinanceUserServer

	buc *biz.BinanceUserUsecase
}

// NewBinanceDataService new a BinanceData service.
func NewBinanceDataService(buc *biz.BinanceUserUsecase) *BinanceUserService {
	return &BinanceUserService{buc: buc}
}

// GetUser 客户端查看用户信息接口，主要是api能否下单使用的反馈
func (b *BinanceUserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	return b.buc.GetUser(ctx, req)
}

// PullUserDeposit 模式1 用户充值总入口, 单一协程处理，遍历所有用户
func (b *BinanceUserService) PullUserDeposit(ctx context.Context, req *v1.PullUserDepositRequest) (*v1.PullUserDepositReply, error) {
	var (
		users []string
		err   error
	)

	users, err = pullUsersByDeposit()
	for _, v := range users {
		var (
			balance string
			cost    uint64
			open    bool
		)
		balance, cost, open, err = pullUserDepositForInfo(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = b.buc.SetUserBalanceAndUser(ctx, v, balance, cost, open, 1)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil, nil
}

func pullUsersByDeposit() ([]string, error) {
	var (
		users []string
	)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		contractAddress := common.HexToAddress("0x966a5a54adf9985043A4651431a3755081e0f603")
		instance, err := abi.NewDeposit(contractAddress, client)
		if err != nil {
			fmt.Println(err)
			return users, err
		}
		var (
			addresses []common.Address
		)

		addresses, err = instance.GetUsers(&bind.CallOpts{})
		if err != nil {
			return users, err
		}

		for _, v := range addresses {
			users = append(users, v.String())
		}

		break
	}

	return users, nil
}

func pullUserDepositForInfo(address string) (string, uint64, bool, error) {
	var (
		balance string
		cost    uint64
		open    bool
	)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		contractAddress := common.HexToAddress("0x966a5a54adf9985043A4651431a3755081e0f603")
		instance, err := abi.NewDeposit(contractAddress, client)
		if err != nil {
			fmt.Println(err)
			return balance, cost, open, err
		}

		var (
			tmp  *big.Int
			tmp2 *big.Int
			tmp3 bool
		)
		tmp, err = instance.UserdepositForUsdtAmount(&bind.CallOpts{}, common.HexToAddress(address))
		if err != nil {
			return balance, cost, open, err
		}

		tmp2, err = instance.UserCost(&bind.CallOpts{}, common.HexToAddress(address))
		if err != nil {
			return balance, cost, open, err
		}

		//contractAddress2 := common.HexToAddress("0xCD329BEb4574Ddf10411F71383FA0d313A65A737")
		//instance2, err := abi.NewDepositOpen(contractAddress2, client)
		//if err != nil {
		//	fmt.Println(err)
		//	return balance, cost, open, err
		//}

		//tmp3, err = instance2.UserOpen(&bind.CallOpts{}, common.HexToAddress(address))
		//if err != nil {
		//	return balance, cost, open, err
		//}

		if "0" != tmp.String() {
			balance = tmp.String()
		}

		if 0 < tmp2.Uint64() {
			cost = tmp2.Uint64()
		}

		open = tmp3

		break
	}

	return balance, cost, open, nil
}

// PullUserDeposit2 .
func (b *BinanceUserService) PullUserDeposit2(ctx context.Context, req *v1.PullUserDepositRequest) (*v1.PullUserDepositReply, error) {
	var (
		users []string
		err   error
	)

	users, err = pullUsersByStakeTfi()
	for _, v := range users {
		var (
			balance string
			cost    uint64
			open    bool
		)
		balance, cost, open, err = pullUserStakeTfiForInfo(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = b.buc.SetUserBalanceAndUser(ctx, v, balance, cost, open, 2)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil, nil
}

func pullUsersByStakeTfi() ([]string, error) {
	var (
		users []string
	)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		contractAddress := common.HexToAddress("0x21081a5278429c434A3D772D4A5ac4764E9dD815")
		instance, err := abi.NewStakeTfi(contractAddress, client)
		if err != nil {
			fmt.Println(err)
			return users, err
		}
		var (
			addresses []common.Address
		)

		addresses, err = instance.GetUsers(&bind.CallOpts{})
		if err != nil {
			return users, err
		}

		for _, v := range addresses {
			users = append(users, v.String())
		}

		break
	}

	return users, nil
}

func pullUserStakeTfiForInfo(address string) (string, uint64, bool, error) {
	var (
		balance string
		cost    uint64
		open    bool
		err     error
	)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		contractAddress := common.HexToAddress("0x21081a5278429c434A3D772D4A5ac4764E9dD815")
		instance, err := abi.NewStakeTfi(contractAddress, client)
		if err != nil {
			fmt.Println(err)
			return balance, cost, open, err
		}

		var (
			tmp  *big.Int
			tmp2 *big.Int
			tmp3 bool
		)
		//tmp3, err = instance.UserOpen(&bind.CallOpts{}, common.HexToAddress(address))
		//if err != nil {
		//	return balance, cost, open, err
		//}

		tmp, err = instance.UserTfiAmount(&bind.CallOpts{}, common.HexToAddress(address))
		if err != nil {
			return balance, cost, open, err
		}

		tmp2, err = instance.UserCost(&bind.CallOpts{}, common.HexToAddress(address))
		if err != nil {
			return balance, cost, open, err
		}

		if "0" != tmp.String() {
			balance = tmp.String()
		}

		if 0 < tmp2.Uint64() {
			cost = tmp2.Uint64()
		}

		open = tmp3
		break
	}

	return balance, cost, open, err
}

// PullUserCredentialsBsc 拉取用户api_key api_secret
func (b *BinanceUserService) PullUserCredentialsBsc(ctx context.Context, req *v1.PullUserCredentialsBscRequest) (*v1.PullUserCredentialsBscReply, error) {
	var (
		users []*biz.User
		err   error
	)

	users, err = b.buc.GetUsers()
	if nil != err {
		fmt.Println(err)
		return nil, err
	}

	for _, v := range users {
		var (
			apiKey    string
			apiSecret string
		)

		apiKey, apiSecret, err = pullUserCredentialsBscBySystemRole(v.Address)
		if nil != err {
			fmt.Println(err)
			return nil, err
		}

		err = b.buc.UpdateUser(ctx, v, apiKey, apiSecret)
		if nil != err {
			fmt.Println(err)
			continue
		}
	}

	return nil, nil
}

func pullUserCredentialsBscBySystemRole(address string) (string, string, error) {
	var (
		apiKey    string
		apiSecret string
	)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		tokenAddress := common.HexToAddress("0xF6811b1fff0bA973Cc09A12BE874E592e2Ff9a2d")
		instance, err := abi.NewUserCredentialsBsc(tokenAddress, client)
		if err != nil {
			fmt.Println(err)
			return apiKey, apiSecret, err
		}

		apiKey, apiSecret, err = instance.GetUserCredentialsBscBySystemRole(&bind.CallOpts{
			From: common.HexToAddress("0x84B9566F03f0F8A7F6b5abA2f684Df8082ed8093"),
		}, common.HexToAddress(address))
		if err != nil {
			return apiKey, apiSecret, err
		}

		break
	}

	return apiKey, apiSecret, nil
}

func (b *BinanceUserService) BindTrader(ctx context.Context, req *v1.BindTraderRequest) (*v1.BindTraderReply, error) {
	return b.buc.BindTrader(ctx)
}

func (b *BinanceUserService) ListenTraderAndUserOrder(ctx context.Context, req *v1.ListenTraderAndUserOrderRequest) (*v1.ListenTraderAndUserOrderReply, error) {
	return b.buc.ListenTraders(ctx, req)
}

func (b *BinanceUserService) OrderHandle(ctx context.Context, req *v1.OrderHandleRequest) (*v1.OrderHandleReply, error) {
	return b.buc.OrderHandle(ctx, req)
}

func (b *BinanceUserService) OrderHandleTwo(ctx context.Context, req *v1.OrderHandleRequest) (*v1.OrderHandleReply, error) {
	return b.buc.OrderHandleTwo(ctx, req)
}

func (b *BinanceUserService) Analyze(ctx context.Context, req *v1.AnalyzeRequest) (*v1.AnalyzeReply, error) {
	return b.buc.Analyze(ctx, req)
}

func (b *BinanceUserService) CloseOrderAfterBind(ctx context.Context, req *v1.CloseOrderAfterBindRequest) (*v1.CloseOrderAfterBindReply, error) {
	return b.buc.CloseOrderAfterBind(ctx, req)
}

func (b *BinanceUserService) CloseOrderAfterBindTwo(ctx context.Context, req *v1.CloseOrderAfterBindRequest) (*v1.CloseOrderAfterBindReply, error) {
	return b.buc.CloseOrderAfterBindTwo(ctx, req)
}

func (b *BinanceUserService) InitOrderAfterBind(ctx context.Context, req *v1.InitOrderAfterBindRequest) (*v1.InitOrderAfterBindReply, error) {
	return b.buc.InitOrderAfterBind(ctx, req)
}

func (b *BinanceUserService) InitOrderAfterBindTwo(ctx context.Context, req *v1.InitOrderAfterBindRequest) (*v1.InitOrderAfterBindReply, error) {
	return b.buc.InitOrderAfterBindTwo(ctx, req)
}

func (b *BinanceUserService) OverOrderAfterBind(ctx context.Context, req *v1.OverOrderAfterBindRequest) (*v1.OverOrderAfterBindReply, error) {
	return b.buc.OverOrder(ctx, req)
}

func (b *BinanceUserService) OverOrderAfterBindTwo(ctx context.Context, req *v1.OverOrderAfterBindRequest) (*v1.OverOrderAfterBindReply, error) {
	return b.buc.OverOrderTwo(ctx, req)
}

func (b *BinanceUserService) AdminOverOrderAfterBind(ctx context.Context, req *v1.OverOrderAfterBindRequest) (*v1.OverOrderAfterBindReply, error) {
	return b.buc.AdminOverOrder(ctx, req)
}

func (b *BinanceUserService) AdminOverOrderAfterBindTwo(ctx context.Context, req *v1.OverOrderAfterBindRequest) (*v1.OverOrderAfterBindReply, error) {
	return b.buc.AdminOverOrderTwo(ctx, req)
}

// PullTradingBoxOpen .
func (b *BinanceUserService) PullTradingBoxOpen(ctx context.Context, req *v1.PullTradingBoxOpenRequest) (*v1.PullTradingBoxOpenReply, error) {
	var (
		terms          []uint64
		boxTotalAmount map[uint64]uint64
		balance        float64
		tokenIds       []uint64
		boxAmount      map[uint64]uint64
		tradingBoxOpen map[uint64]*biz.TradingBoxOpen
		err            error
	)

	terms, boxTotalAmount, err = pullBoxTerm()
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	if 0 >= len(terms) {
		return nil, nil
	}

	balance, err = b.buc.GetTermBinanceCurrentBalance(ctx, 17)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	tokenIds, boxAmount, err = pullTradingBoxOpen()
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	tradingBoxOpen, err = b.buc.GetTradingBoxOpenMap()
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	for _, vTokenIds := range tokenIds {
		if _, ok := boxAmount[vTokenIds]; !ok {
			continue
		}

		if _, ok := tradingBoxOpen[vTokenIds]; ok { // 存在数据
			continue
		}

		tmpLastTerm := uint64(0)
		for _, vTerms := range terms {
			if _, ok := boxTotalAmount[vTerms]; !ok {
				continue
			}

			if tmpLastTerm < vTokenIds && vTokenIds <= vTerms { // 本期区间

				tmpRate := float64(boxAmount[vTokenIds]) / float64(boxTotalAmount[vTerms])
				err = b.buc.InsertTradingBoxOpen(ctx, vTokenIds, boxAmount[vTokenIds], boxTotalAmount[vTerms], tmpRate, balance, balance*tmpRate)
				if nil != err {
					fmt.Println(err)
					continue
				}

			}

			tmpLastTerm = vTerms
		}

	}

	return nil, nil
}

func pullTradingBoxOpen() ([]uint64, map[uint64]uint64, error) {
	var (
		openBox   []uint64
		boxAmount map[uint64]uint64
		err       error
	)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		contractAddress := common.HexToAddress("0xC2550aedbF6B1c858eFE8CEa4FFed1D383723887")
		instance, err := abi.NewTradingBox(contractAddress, client)
		if err != nil {
			fmt.Println(err)
			return openBox, boxAmount, err
		}

		var (
			tmp *big.Int
		)

		tmp, err = instance.TotalSupply(&bind.CallOpts{})
		if err != nil {
			return openBox, boxAmount, err
		}

		if 0 >= tmp.Uint64() {
			return openBox, boxAmount, err
		}

		openBox = make([]uint64, 0)
		boxAmount = make(map[uint64]uint64, 0)
		for tokenId := int64(1); tokenId <= tmp.Int64(); tokenId++ {
			var (
				tmp2 *big.Int
				tmp3 *big.Int
			)

			tmp3, err = instance.BoxTokenAmount(&bind.CallOpts{}, big.NewInt(tokenId))
			if err != nil {
				continue
			}

			lengthToKeep := len(tmp3.String()) - 15 // todo 每个盒子业务默认100u

			if lengthToKeep > 0 {
				amountTmpStr := tmp3.String()[:lengthToKeep]
				var (
					tmp4 int64
				)
				tmp4, err = strconv.ParseInt(amountTmpStr, 10, 64)
				if nil != err || 0 >= tmp4 {
					continue
				}

				boxAmount[uint64(tokenId)] = uint64(tmp4)
			} else {
				continue
			}

			tmp2, err = instance.OpenTimeStamp(&bind.CallOpts{}, big.NewInt(tokenId))
			if err != nil {
				continue
			}

			if 0 < tmp2.Uint64() {
				openBox = append(openBox, uint64(tokenId))
			}
		}

		break
	}

	return openBox, boxAmount, err
}

func pullBoxTerm() ([]uint64, map[uint64]uint64, error) {
	var (
		terms      []uint64
		termAmount map[uint64]uint64
		err        error
	)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		contractAddress := common.HexToAddress("0x40e9803682E7739F574d8fbEB581a91d42Bc8292")
		instance, err := abi.NewTradingAdmin(contractAddress, client)
		if err != nil {
			fmt.Println(err)
			return terms, termAmount, err
		}

		var (
			tmp []*big.Int
		)

		tmp, err = instance.GetTradingTerm(&bind.CallOpts{})
		if err != nil {
			return terms, termAmount, err
		}

		if 0 >= len(tmp) {
			return terms, termAmount, err
		}

		terms = make([]uint64, 0)
		termAmount = make(map[uint64]uint64, 0)
		for _, vTmp := range tmp {
			var (
				tmp3 *big.Int
			)

			tmp3, err = instance.GetTermAmount(&bind.CallOpts{}, vTmp)
			if err != nil {
				continue
			}

			lengthToKeep := len(tmp3.String()) - 15 // todo 每个盒子业务默认100u

			if lengthToKeep > 0 {
				amountTmpStr := tmp3.String()[:lengthToKeep]
				var (
					tmp4 int64
				)
				tmp4, err = strconv.ParseInt(amountTmpStr, 10, 64)
				if nil != err || 0 >= tmp4 {
					continue
				}

				termAmount[vTmp.Uint64()] = uint64(tmp4)
			} else {
				continue
			}

			terms = append(terms, vTmp.Uint64())
		}

		break
	}

	return terms, termAmount, err
}

// SettleTradingBoxOpen .
func (b *BinanceUserService) SettleTradingBoxOpen(ctx context.Context, req *v1.SettleTradingBoxOpenRequest) (*v1.SettleTradingBoxOpenReply, error) {
	//var (
	//	terms          []uint64
	//	tradingBoxOpen map[uint64]*biz.TradingBoxOpen
	//	err            error
	//)
	//
	//terms, _, err = pullBoxTerm()
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, nil
	//}
	//
	//tradingBoxOpen, err = b.buc.GetTradingBoxOpenMapByStatus(1)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, nil
	//}
	//
	//for tokenId, vTradingBoxOpen := range tradingBoxOpen {
	//	// tokenId从1开始，和 totalSupply是相等的，因为不能燃烧，燃烧需要update_role权限也不会放开
	//	tmpLastTerm := uint64(0)
	//	for _, vTerms := range terms {
	//		if tmpLastTerm >= tokenId || tokenId > vTerms { // 非本期区间
	//			continue
	//		}
	//
	//		tmpLastTerm = vTerms
	//	}
	//}

	return nil, nil
}
