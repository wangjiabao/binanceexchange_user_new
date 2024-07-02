package biz

import (
	v1 "binanceexchange_user/api/binanceexchange_user/v1"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

type User struct {
	ID                  uint64
	Address             string
	PlayType            uint64
	ApiStatus           uint64
	ApiKey              string
	ApiSecret           string
	BindTraderStatus    uint64
	BindTraderStatusTfi uint64
	IsDai               uint64
	UseNewSystem        uint64
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type UserBalance struct {
	ID        uint64
	UserId    uint64
	Balance   string
	Cost      uint64
	Open      uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserBalanceRecord struct {
	ID        uint64
	UserId    uint64
	Amount    string
	Balance   string
	Tx        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserAmount struct {
	ID        uint64
	UserId    uint64
	Amount    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserAmountRecord struct {
	ID        uint64
	UserId    uint64
	OrderId   uint64
	Amount    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Trader struct {
	ID          uint64
	Status      uint64
	Amount      uint64
	BaseMoney   float64
	PortfolioId string
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TraderPosition struct {
	ID           uint64
	TraderId     uint64
	Symbol       string
	Qty          float64
	Side         string
	PositionSide string
}

type TraderPositionNew struct {
	ID     uint64
	Symbol string
	Qty    float64
	Side   string
	Closed uint64
	Opened uint64
}

type UserBindTrader struct {
	ID        uint64
	UserId    uint64
	TraderId  uint64
	Amount    uint64
	Status    uint64
	InitOrder uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserBindAfterUnbind struct {
	ID           uint64
	UserId       uint64
	TraderId     uint64
	Symbol       string
	PositionSide string
	Quantity     float64
	Status       uint64
	Amount       uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserOrder struct {
	ID            uint64
	UserId        uint64
	TraderId      uint64
	ClientOrderId string
	OrderId       string
	Symbol        string
	Side          string
	PositionSide  string
	Quantity      float64
	Price         float64
	TraderQty     float64
	OrderType     string
	ClosePosition string
	CumQuote      float64
	ExecutedQty   float64
	AvgPrice      float64
	HandleStatus  int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UserOrderErr struct {
	ID            uint64
	UserId        uint64
	TraderId      uint64
	ClientOrderId string
	OrderId       string
	Symbol        string
	Side          string
	PositionSide  string
	Quantity      float64
	Price         float64
	TraderQty     float64
	OrderType     string
	ClosePosition string
	CumQuote      float64
	ExecutedQty   float64
	AvgPrice      float64
	HandleStatus  int64
	InitOrder     int64
	Code          int64
	Msg           string
	Proportion    float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type TradingBoxOpen struct {
	ID             uint64
	TokenId        uint64
	Amount         uint64
	AmountTotal    uint64
	WithdrawStatus uint64
	AmountRate     float64
	Total          float64
	WithdrawAmount float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type BinanceTradeHistory struct {
	ID                  uint64
	Time                uint64
	Symbol              string
	Side                string
	Price               float64
	Fee                 float64
	FeeAsset            string
	Quantity            float64
	QuantityAsset       string
	RealizedProfit      float64
	RealizedProfitAsset string
	BaseAsset           string
	Qty                 float64
	PositionSide        string
	ActiveBuy           string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type BinancePositionHistory struct {
	ID              uint64
	Symbol          string
	Side            string
	Closed          uint64
	Opened          uint64
	AvgCost         float64
	AvgClosePrice   float64
	ClosingPnl      float64
	MaxOpenInterest float64
	ClosedVolume    float64
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type UserOrderDo struct {
	ID           uint64
	ApiKey       string
	ApiSecret    string
	ApiKeyTwo    string
	ApiSecretTwo string
	Symbol       string
	Status       uint64
	CloseRate    uint64
	CloseBase    uint64
	Qty          float64
	Price        float64
	ClosePrice   float64
	Amount       float64
	QtyTwo       float64
	PriceTwo     float64
	AmountTwo    float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserOrderDoNew struct {
	ID           uint64
	ApiKey       string
	ApiSecret    string
	ApiKeyTwo    string
	ApiSecretTwo string
	Symbol       string
	SymbolTwo    string
	Status       uint64
	Qty          float64
	QtyTwo       float64
	Side         string
	SideTwo      string
	OrderId      string
	OrderIdTwo   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Symbol struct {
	ID                uint64
	Symbol            string
	QuantityPrecision int64
	PricePrecision    int64
}

type BinanceLeverAge struct {
	LeverAge int64
	symbol   string
}

type BinanceMarginType struct {
	Code int64
	Msg  string
}

type BinancePositionSide struct {
	Code int64
	Msg  string
}

type OrderInfo struct {
	Code int64
	Msg  string
}

type BinanceSymbolPrice struct {
	Symbol string
	Price  string
}

type OrderHistory struct {
	OrderId      int64
	Qty          string
	Symbol       string
	Price        string
	Side         string
	RealizedPnl  string
	QuoteQty     string
	PositionSide string
	Time         int64
}

type OrderData struct {
	Coin     string
	Type     string
	Price    string
	Side     string
	Qty      string
	Position string
}

type BinanceUserBalance struct {
	Balance            string
	Asset              string
	CrossUnPnl         string
	AvailableBalance   string
	CrossWalletBalance string
	MaxWithdrawAmount  string
}

type BinanceOrder struct {
	OrderId       int64
	ExecutedQty   string
	ClientOrderId string
	Symbol        string
	AvgPrice      string
	CumQuote      string
	Side          string
	PositionSide  string
	ClosePosition bool
	Type          string
	Status        string
}

type BinanceTradeHistoryResp struct {
	Data *BinanceTradeHistoryData
}

type BinanceTradeHistoryData struct {
	Total uint64
	List  []*BinanceTradeHistoryDataList
}

type BinanceTrade struct {
	ID        uint64
	TraderNum uint64
	Status    uint64
}

type BinanceTradeHistoryDataList struct {
	Time                uint64
	Symbol              string
	Side                string
	Price               float64
	Fee                 float64
	FeeAsset            string
	Quantity            float64
	QuantityAsset       string
	RealizedProfit      float64
	RealizedProfitAsset string
	BaseAsset           string
	Qty                 float64
	PositionSide        string
	ActiveBuy           bool
}

type BinancePositionHistoryResp struct {
	Data *BinancePositionHistoryData
}

type BinancePositionHistoryData struct {
	Total uint64
	List  []*BinancePositionHistoryDataList
}

type BinancePositionHistoryDataList struct {
	Symbol          string
	Side            string
	Closed          uint64
	Opened          uint64
	AvgCost         float64
	AvgClosePrice   float64
	ClosingPnl      float64
	MaxOpenInterest float64
	ClosedVolume    float64
	Status          string
}

type FilData struct {
	TotalCount uint64
	Transfers  []*FilDataList
}

type FilDataList struct {
	Timestamp uint64
	From      string
	To        string
	Value     string
	Type      string
}

type FilData2 struct {
	ID       uint64
	From     string
	FromType uint64
	To       string
	Value    string
}

type BinanceUserRepo interface {
	InsertUser(ctx context.Context, User *User) (*User, error)
	UpdateUserApiStatus(ctx context.Context, userId uint64) (bool, error)
	UpdateUser(ctx context.Context, userId uint64, apiKey string, apiSecret string) (bool, error)
	UpdateUserBindTraderStatus(ctx context.Context, userId uint64) (bool, error)
	UpdateUserIsDai(ctx context.Context, address string) (bool, error)
	UpdateUserBindTraderStatusNo(ctx context.Context, userId uint64) (bool, error)
	UpdateUserBindTraderStatusNoTfi(ctx context.Context, userId uint64) (bool, error)
	InsertUserBalance(ctx context.Context, userBalance *UserBalance) (bool, error)
	InsertUserBalanceTwo(ctx context.Context, userBalance *UserBalance) (bool, error)
	UpdatesUserBalance(ctx context.Context, userId uint64, balance string, cost uint64, open uint64) (bool, error)
	UpdatesUserBalanceTwo(ctx context.Context, userId uint64, balance string, cost uint64, open uint64) (bool, error)
	InsertUserBalanceRecord(ctx context.Context, userBalance *UserBalanceRecord) (bool, error)
	InsertUserBalanceRecordTwo(ctx context.Context, userBalanceRecord *UserBalanceRecord) (bool, error)
	InsertUserAmount(ctx context.Context, userAmount *UserAmount) (bool, error)
	InsertUserAmountTwo(ctx context.Context, userAmount *UserAmount) (bool, error)
	UpdatesUserAmount(ctx context.Context, userId uint64, amount int64) (bool, error)
	UpdatesUserAmountTwo(ctx context.Context, userId uint64, amount int64) (bool, error)
	InsertUserAmountRecord(ctx context.Context, userAmount *UserAmountRecord) (bool, error)
	InsertUserAmountRecordTwo(ctx context.Context, userAmount *UserAmountRecord) (bool, error)
	InsertUserBindTrader(ctx context.Context, userId uint64, traderId uint64, amount uint64) (*UserBindTrader, error)
	InsertUserBindTraderTwo(ctx context.Context, userId uint64, traderId uint64, amount uint64) (*UserBindTrader, error)
	UpdatesUserBindTraderStatus(ctx context.Context, userId uint64, status uint64) (bool, error)
	UpdatesUserBindTraderStatusById(ctx context.Context, id uint64, status uint64) (bool, error)
	UpdatesUserBindTraderStatusAndInitOrderById(ctx context.Context, id uint64, status uint64, initOrder uint64, amount uint64) (bool, error)
	UpdatesUserBindTraderInitOrderById(ctx context.Context, id uint64) (bool, error)
	UpdatesUserBindTraderTwoInitOrderById(ctx context.Context, id uint64) (bool, error)
	UpdatesUserBindTraderTwoById(ctx context.Context, id uint64, amount uint64) (bool, error)
	UpdatesUserBindTraderTwoAdminOverOrderById(ctx context.Context, id uint64) (bool, error)
	UpdatesUserBindTraderTwoAfterAdminOverOrderById(ctx context.Context, id uint64) (bool, error)
	UpdatesUserBindTraderStatusAndAmountById(ctx context.Context, id uint64, status uint64, amount uint64) (bool, error)
	DeleteUserBindTrader(ctx context.Context, userId uint64) (bool, error)
	InsertUserOrder(ctx context.Context, order *UserOrder) (*UserOrder, error)
	InsertUserOrderErr(ctx context.Context, order *UserOrderErr) (*UserOrderErr, error)
	InsertUserOrderErrTwo(ctx context.Context, order *UserOrderErr) (*UserOrderErr, error)
	InsertUserOrderTwo(ctx context.Context, order *UserOrder) (*UserOrder, error)
	UpdatesUserOrderHandleStatus(ctx context.Context, id uint64) (bool, error)
	UpdatesUserOrderTwoHandleStatus(ctx context.Context, id uint64) (bool, error)
	InsertUserBindAfterUnbind(ctx context.Context, u *UserBindAfterUnbind) (*UserBindAfterUnbind, error)
	UpdatesUserBindAfterUnbind(ctx context.Context, id uint64, status uint64, quantity float64) (bool, error)
	UpdatesUserBindAfterUnbindTwo(ctx context.Context, id uint64, status uint64, quantity float64) (bool, error)
	GetUsers() ([]*User, error)
	GetUsersByUserIds(userIds []uint64) (map[uint64]*User, error)
	GetUsersByBindUserStatus() ([]*User, error)
	GetUsersByBindUserStatusReBind() ([]*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	GetUserById(ctx context.Context, userId uint64) (*User, error)
	GetUserByApiKeyAndApiSecret(key string, secret string) (*User, error)
	GetUserBalance(ctx context.Context, userId uint64) (*UserBalance, error)
	GetUserBalanceTwo(ctx context.Context, userId uint64) (*UserBalance, error)
	GetUserAmount(ctx context.Context, userId uint64) (*UserAmount, error)
	GetUserAmountTwo(ctx context.Context, userId uint64) (*UserAmount, error)
	GetUserBalanceByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*UserBalance, error)
	GetUserBalanceTwoByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*UserBalance, error)
	GetUserAmountByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*UserAmount, error)
	GetUserAmountTwoByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*UserAmount, error)
	GetTradersOrderByAmountDesc() ([]*Trader, error)
	GetTraders() (map[uint64]*Trader, error)
	GetTradersByTraderNum() (map[string]*Trader, error)
	GetUserBindTraderByUserId(userId uint64) ([]*UserBindTrader, error)
	GetUserBindTraderTwoByUserId(userId uint64) ([]*UserBindTrader, error)
	GetUserBindTraderTwoById(id uint64) (*UserBindTrader, error)
	GetUserBindTrader() (map[uint64][]*UserBindTrader, error)
	GetUserBindTraderByStatus(status uint64) (map[uint64][]*UserBindTrader, error)
	GetUserBindTraderMapByUserIds(userIds []uint64) (map[uint64][]*UserBindTrader, error)
	GetUserBindTraderByTraderIds(traderIds []uint64) (map[uint64][]*UserBindTrader, error)
	GetUserBindTraderTwoByTraderIds(traderIds []uint64) (map[uint64][]*UserBindTrader, error)
	GetUserBindAfterUnbindByTraderIds(traderIds []uint64) (map[uint64][]*UserBindAfterUnbind, error)
	GetUserBindAfterUnbindByStatus() ([]*UserBindAfterUnbind, error)
	GetUserBindAfterUnbindTwoByStatus() ([]*UserBindAfterUnbind, error)
	GetUserBindTraderByInitOrder() (map[uint64][]*UserBindTrader, error)
	GetUserBindTraderTwoByInitOrder() (map[uint64][]*UserBindTrader, error)
	GetUserBindTraderByAlreadyInitOrder() (map[uint64][]*UserBindTrader, error)
	GetUserBindTraderTwoByAlreadyInitOrder() (map[uint64][]*UserBindTrader, error)
	GetUserBindTraderByOverOrder() (map[uint64][]*UserBindTrader, error)
	GetUserBindTraderTwoByOverOrder() (map[uint64][]*UserBindTrader, error)
	GetUserBindAfterUnbindByUserIdAndTraderIdAndSymbol(ctx context.Context, userId uint64, traderId uint64, symbol string, positionSide string) (*UserBindAfterUnbind, error)
	GetSymbol() (map[string]*Symbol, error)
	GetUserOrderByUserTraderIdAndSymbol(userId uint64, traderId uint64, symbol string) ([]*UserOrder, error)
	GetUserOrderTwoByUserTraderIdAndSymbol(userId uint64, traderId uint64, symbol string) ([]*UserOrder, error)
	GetUserOrderByUserTraderId(userId uint64, traderId uint64) (map[string][]*UserOrder, error)
	GetUserOrderByUserIdMapId(userId uint64) (map[string]*UserOrder, error)
	GetUserOrderByUserIdAndSymbolAndPositionSide(userId uint64, symbol string, positionSide string) ([]*UserOrder, error)
	GetUserOrderByHandleStatus() ([]*UserOrder, error)
	GetUserOrderByUserIdGroupBySymbol(userId uint64) ([]*UserOrder, error)
	GetUserOrderByIds(ids []int64) ([]*UserOrder, error)
	GetUserOrderTwoByIds(ids []int64) ([]*UserOrder, error)
	GetUserOrderById(orderId int64) (*UserOrder, error)
	GetTraderPosition(traderId uint64) ([]*TraderPosition, error)
	GetOpeningTraderPositionNew(traderNum string) ([]*TraderPositionNew, error)
	InsertTradingBoxOpen(ctx context.Context, box *TradingBoxOpen) (*TradingBoxOpen, error)
	UpdateTradingBoxOpenStatus(ctx context.Context, tokenId uint64) (bool, error)
	GetTradingBoxOpen() ([]*TradingBoxOpen, error)
	GetTradingBoxOpenMap() (map[uint64]*TradingBoxOpen, error)
	GetTradingBoxOpenMapByStatus(status uint64) (map[uint64]*TradingBoxOpen, error)
	GetBinanceTradeHistory(traderNum uint64) ([]*BinanceTradeHistory, error)
	GetBinanceTradeHistoryByTraderNumNewest(traderNum uint64, limit int) ([]*BinanceTradeHistory, error)
	GetBinanceTrade() ([]*BinanceTrade, error)
	InsertBinanceTradeHistory(ctx context.Context, history *BinanceTradeHistory, traderNum uint64) (*BinanceTradeHistory, error)
	GetBinancePositionHistory(traderNum uint64) ([]*BinancePositionHistory, error)
	GetBinancePositionHistoryByTraderNumNewest(traderNum uint64) (*BinancePositionHistory, error)
	InsertBinancePositionHistory(ctx context.Context, binancePositionHistory *BinancePositionHistory) (*BinancePositionHistory, error)
	InsertUserOrderDo(ctx context.Context, userOrderDo *UserOrderDo) error
	InsertUserOrderDoNew(ctx context.Context, userOrderDo *UserOrderDoNew) error
	UpdateUserOrderDo(ctx context.Context, id uint64, closePrice string) (bool, error)
	UpdateUserOrderDoTwo(ctx context.Context, id uint64, orderId string, orderIdTwo string) (bool, error)
	UpdateUserOrderDoThree(ctx context.Context, id uint64, qty float64, qtyTwo float64) (bool, error)
	GetUserOrderDoNew() ([]*UserOrderDoNew, error)
	GetUserOrderDo() ([]*UserOrderDo, error)
	GetUserOrderDoLast() ([]*UserOrderDo, error)
	InsertFilData(ctx context.Context, filData *FilData2) error
	GetFilData(address ...[]string) (map[string][]*FilData2, error)

	SAddOrderSetSellLongOrBuyShort(ctx context.Context, OrderId int64) error
	SMembersOrderSetSellLongOrBuyShort(ctx context.Context) ([]string, error)
	SRemOrderSetSellLongOrBuyShort(ctx context.Context, OrderId int64) error

	SAddOrderSetSellLongOrBuyShortTwo(ctx context.Context, OrderId int64) error
	SMembersOrderSetSellLongOrBuyShortTwo(ctx context.Context) ([]string, error)
	SRemOrderSetSellLongOrBuyShortTwo(ctx context.Context, OrderId int64) error
}

// BinanceUserUsecase is a BinanceData usecase.
type BinanceUserUsecase struct {
	binanceUserRepo BinanceUserRepo
	tx              Transaction
	log             *log.Helper
}

// NewBinanceDataUsecase new a BinanceData usecase.
func NewBinanceDataUsecase(binanceUserRepo BinanceUserRepo, tx Transaction, logger log.Logger) *BinanceUserUsecase {
	return &BinanceUserUsecase{binanceUserRepo: binanceUserRepo, tx: tx, log: log.NewHelper(logger)}
}

func (b *BinanceUserUsecase) SetUserBalanceAndUser(ctx context.Context, address string, balance string, cost uint64, open bool, playType uint64) error {
	var (
		user            *User
		userBalance     *UserBalance
		balance1        string
		balance2        string
		cost1           uint64
		cost2           uint64
		open1           uint64
		open2           uint64
		err             error
		lastUserBalance string // 上次用户余额
	)

	if 1 == playType {
		balance1 = balance
		cost1 = cost
		// todo
		//if open {
		open1 = 1
		//}

	} else if 2 == playType {
		balance2 = balance
		cost2 = cost
		//if open {
		open2 = 1
		//}

		var balanceTmp int64
		lengthToKeep := len(balance2) - 18 // todo 当len=19时表示1个代币，截取到10

		if lengthToKeep > 0 {
			balanceTmpStr := balance2[:lengthToKeep]
			balanceTmp, err = strconv.ParseInt(balanceTmpStr, 10, 64)
			if nil != err || 0 >= balanceTmp {
				return nil
			}

			if 100 > balanceTmp { // 最少100个tfi
				return nil
			}

			if balanceTmp <= int64(cost2) {
				cost2 = uint64(balanceTmp)
			}

		} else {
			// 不足1个tfi
			cost2 = 0
		}

	}
	user, err = b.binanceUserRepo.GetUserByAddress(ctx, address)
	if nil != err {
		return err
	}

	// 初始化 用户和余额
	if nil == user {
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			user, err = b.binanceUserRepo.InsertUser(ctx, &User{
				Address: address, // 初始化玩法，模式
			})
			if nil != err {
				return err
			}

			if nil == user {
				return errors.New(500, "CREATE_USER_ERROR", "未发现创建的用户")
			}

			_, err = b.binanceUserRepo.InsertUserBalance(ctx, &UserBalance{
				UserId:  user.ID,
				Balance: balance1,
				Cost:    cost1,
				Open:    open1,
			})
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserBalanceTwo(ctx, &UserBalance{
				UserId:  user.ID,
				Balance: balance2,
				Cost:    cost2,
				Open:    open2,
			})
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserAmount(ctx, &UserAmount{
				UserId: user.ID,
			})
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserAmountTwo(ctx, &UserAmount{
				UserId: user.ID,
			})
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserBalanceRecord(ctx, &UserBalanceRecord{
				UserId:  user.ID,
				Amount:  balance1,
				Balance: lastUserBalance,
			})
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserBalanceRecordTwo(ctx, &UserBalanceRecord{
				UserId:  user.ID,
				Amount:  balance2,
				Balance: lastUserBalance,
			})
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			return err
		}

	} else {
		// 模式对应才给充值，初始化用户时已经决定了，条件很关键
		if 1 == playType {
			userBalance, err = b.binanceUserRepo.GetUserBalance(ctx, user.ID)
			if nil != err {
				return err
			}

			// 无变化
			if userBalance.Balance == balance1 && userBalance.Cost == cost1 && userBalance.Open == open1 {
				return nil
			}

			// 有变化
			tmpAmount := new(big.Int)
			if userBalance.Balance != balance1 {
				// 上次余额
				tmpLastUserBalance := new(big.Int)
				tmpLastUserBalance.SetString(userBalance.Balance, 10)
				// 本次余额
				tmpCurrentUserBalance := new(big.Int)
				tmpCurrentUserBalance.SetString(balance1, 10)
				// 增长
				tmpAmount.Sub(tmpCurrentUserBalance, tmpLastUserBalance)
			}

			tmpCost := userBalance.Cost
			if tmpCost != cost1 {
				tmpCost = cost1
			}

			// 更新余额
			if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = b.binanceUserRepo.UpdatesUserBalance(ctx, user.ID, balance1, tmpCost, open1)
				if nil != err {
					return err
				}

				if userBalance.Balance != balance1 {
					_, err = b.binanceUserRepo.InsertUserBalanceRecord(ctx, &UserBalanceRecord{
						UserId:  user.ID,
						Amount:  tmpAmount.String(),
						Balance: userBalance.Balance,
					})
					if nil != err {
						return err
					}
				}

				// 开单额度变更时，更改绑定状态
				if userBalance.Cost != tmpCost {
					_, err = b.binanceUserRepo.UpdateUserBindTraderStatusNo(ctx, user.ID)
					if nil != err {
						return err
					}
				}

				return nil
			}); err != nil {
				return err
			}
		} else if 2 == playType {
			userBalance, err = b.binanceUserRepo.GetUserBalanceTwo(ctx, user.ID)
			if nil != err {
				return err
			}

			// 无变化
			if userBalance.Balance == balance2 && userBalance.Cost == cost2 && userBalance.Open == open2 {
				return nil
			}

			// 有变化
			tmpAmount := new(big.Int)
			if userBalance.Balance != balance2 {
				// 上次余额
				tmpLastUserBalance := new(big.Int)
				tmpLastUserBalance.SetString(userBalance.Balance, 10)
				// 本次余额
				tmpCurrentUserBalance := new(big.Int)
				tmpCurrentUserBalance.SetString(balance2, 10)
				// 增长
				tmpAmount.Sub(tmpCurrentUserBalance, tmpLastUserBalance)
			}

			tmpCost := userBalance.Cost
			if tmpCost != cost2 {
				tmpCost = cost2
			}

			// 更新余额
			if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = b.binanceUserRepo.UpdatesUserBalanceTwo(ctx, user.ID, balance2, tmpCost, open2)
				if nil != err {
					return err
				}

				if userBalance.Balance != balance2 {
					_, err = b.binanceUserRepo.InsertUserBalanceRecordTwo(ctx, &UserBalanceRecord{
						UserId:  user.ID,
						Amount:  tmpAmount.String(),
						Balance: userBalance.Balance,
					})
					if nil != err {
						return err
					}
				}

				// 开单额度变更时，更改绑定状态
				if userBalance.Cost != tmpCost {
					_, err = b.binanceUserRepo.UpdateUserBindTraderStatusNoTfi(ctx, user.ID)
					if nil != err {
						return err
					}
				}

				return nil
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

func isZero(x float64) bool {
	const epsilon = 1e-9 // 定义一个足够小的阈值
	return math.Abs(x) < epsilon
}

func (b *BinanceUserUsecase) UpdateUser(ctx context.Context, user *User, apiKey string, apiSecret string) error {
	var (
		user2  *User
		symbol map[string]*Symbol
		err    error
	)

	user2, err = b.binanceUserRepo.GetUserByApiKeyAndApiSecret(apiKey, apiSecret)
	if nil != err {
		return err
	}

	if nil == user2 || (user2.ID == user.ID) { // api不存在 或 地址和api指向相同ID(同一条记录)
		if apiKey != user.ApiKey || apiSecret != user.ApiSecret { // api_key或api_secret发生了变化
			if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = b.binanceUserRepo.UpdateUser(ctx, user.ID, apiKey, apiSecret)
				if nil != err {
					return err
				}

				return nil
			}); err != nil {
				return err
			}

			var (
				res *BinancePositionSide
			)
			res, err = requestBinancePositionSide(apiKey, apiSecret)
			if nil != err {
				fmt.Println("更改持仓模式异常", err, user)
				return err
			}

			if -2014 == res.Code {
				fmt.Println(user)
				return errors.New(int(res.Code), res.Msg, "api信息不正确")
			}

			if -2015 == res.Code {
				fmt.Println(user)
				return errors.New(int(res.Code), res.Msg, "api信息不正确")
			}

			if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = b.binanceUserRepo.UpdateUserApiStatus(ctx, user.ID)
				if nil != err {
					return err
				}

				return nil
			}); err != nil {
				return err
			}

			symbol, err = b.binanceUserRepo.GetSymbol()
			if nil != err {
				return err
			}

			for k, _ := range symbol {
				_, err = requestBinanceLeverAge(k, int64(5), apiKey, apiSecret)
				if nil != err {
					continue
				}

				_, err = requestBinanceMarginType(k, apiKey, apiSecret)
				if nil != err {
					continue
				}
			}
		}
	}

	return nil
}

func (b *BinanceUserUsecase) GetUsers() ([]*User, error) {
	return b.binanceUserRepo.GetUsers()
}

func (b *BinanceUserUsecase) GetTradingBoxOpen() ([]*TradingBoxOpen, error) {
	return b.binanceUserRepo.GetTradingBoxOpen()
}

func (b *BinanceUserUsecase) GetTradingBoxOpenMap() (map[uint64]*TradingBoxOpen, error) {
	return b.binanceUserRepo.GetTradingBoxOpenMap()
}

func (b *BinanceUserUsecase) GetTradingBoxOpenMapByStatus(status uint64) (map[uint64]*TradingBoxOpen, error) {
	return b.binanceUserRepo.GetTradingBoxOpenMapByStatus(status)
}

func (b *BinanceUserUsecase) InsertTradingBoxOpen(context context.Context, tokenId uint64, amount uint64, amountTotal uint64, amountRate float64, total float64, withdrawAmount float64) error {
	_, err := b.binanceUserRepo.InsertTradingBoxOpen(context, &TradingBoxOpen{
		TokenId:        tokenId,
		Amount:         amount,
		AmountTotal:    amountTotal,
		WithdrawStatus: 0,
		AmountRate:     amountRate,
		Total:          total,
		WithdrawAmount: withdrawAmount,
	})

	return err
}

func (b *BinanceUserUsecase) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	var (
		user           *User
		userBalance    *UserBalance
		userBalanceTwo *UserBalance
		userAmount     *UserAmount
		userAmountTwo  *UserAmount
		err            error
	)

	if 0 >= len(req.Address) || 300 < len(req.Address) {
		return &v1.GetUserReply{}, nil
	}

	user, err = b.binanceUserRepo.GetUserByAddress(ctx, req.Address)
	if nil != err {
		return &v1.GetUserReply{}, nil
	}

	if nil == user {
		return &v1.GetUserReply{}, nil
	}

	userBalance, err = b.binanceUserRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return &v1.GetUserReply{}, nil
	}

	userBalanceTwo, err = b.binanceUserRepo.GetUserBalanceTwo(ctx, user.ID)
	if nil != err {
		return &v1.GetUserReply{}, nil
	}

	userAmount, err = b.binanceUserRepo.GetUserAmount(ctx, user.ID)
	if nil != err {
		return &v1.GetUserReply{}, nil
	}

	userAmountTwo, err = b.binanceUserRepo.GetUserAmountTwo(ctx, user.ID)
	if nil != err {
		return &v1.GetUserReply{}, nil
	}
	return &v1.GetUserReply{Status: int64(user.ApiStatus), Play: int64(user.PlayType), Balance: userBalance.Balance, Amount: userAmount.Amount, BalanceTfi: userBalanceTwo.Balance, AmountTfi: userAmountTwo.Amount}, err
}

// BindTrader
// 用户的绑定交易员，和重新换绑的方法
// 重新换绑会将所有交易员绑定历史更新为不可用
// 重新绑定，且历史开单未全部平仓的交易员的代币数据会加入新数据方便下单逻辑查询，平仓使用
func (b *BinanceUserUsecase) BindTrader(ctx context.Context) (*v1.BindTraderReply, error) {
	var (
		err error
	)

	// 处理绑定
	err = b.UserBindTrader(ctx, nil)
	if nil != err {
		return nil, err
	}

	// 因为额度换绑
	err = b.ReBindTrader(ctx)
	if nil != err {
		return nil, err
	}

	// 处理更换单个trader
	err = b.ChangeBindTrader(ctx)
	if nil != err {
		return nil, err
	}

	return nil, nil
}

func (b *BinanceUserUsecase) UserBindTrader(ctx context.Context, users []*User) error {

	var (
		userBalance map[uint64]*UserBalance
		traders     []*Trader
		err         error
	)

	// 获取未绑定交易员的用户
	if nil == users {
		users, err = b.binanceUserRepo.GetUsersByBindUserStatus()
		if nil != err {
			return err
		}
	}

	if 0 == len(users) {
		return nil
	}

	// 获取用户cost
	userIds := make([]uint64, 0)
	for _, vUsers := range users {
		userIds = append(userIds, vUsers.ID)
	}
	userBalance, err = b.binanceUserRepo.GetUserBalanceByUserIds(ctx, userIds)
	if nil != err {
		return err
	}

	// 按amount排序好的交易员
	traders, err = b.binanceUserRepo.GetTradersOrderByAmountDesc()
	if nil != err {
		return err
	}

	if 0 == len(traders) {
		return nil
	}

	// 1. 遍历，轮询用户
	// 2. 轮询交易员（查询时已经按amount desc排序，用户cost的百分之30>=当前交易员的amount字段，建立绑定关系。
	// 3. 第二轮保留最后的交易员额度限制，将此和用户剩余可用额度对比，按顺序分配给此额度及以下的交易员。
	// 4. 轮询结束时，用户额度剩余不予理会。
	for _, vUsers := range users {
		if _, ok := userBalance[vUsers.ID]; !ok {
			continue
		}

		// 初始化
		tmpCost := userBalance[vUsers.ID].Cost
		bindTrader := make(map[uint64]*Trader, 0)

		// 第一轮
		for _, vTraders := range traders {
			if 0 >= tmpCost {
				break
			}

			if 0 >= vTraders.Amount {
				continue
			}

			if 100 >= vTraders.Amount {
				continue
			}

			if tmpCost*30/100 >= vTraders.Amount {
				// 绑定
				if _, ok := bindTrader[vTraders.ID]; ok {
					continue
				}

				bindTrader[vTraders.ID] = vTraders
				tmpCost -= vTraders.Amount
			}
		}

		// 第二轮，跳过分配限制的额度，剩下的按顺序分配
		for _, vTraders := range traders {
			if 100 > tmpCost {
				break
			}

			if 0 >= vTraders.Amount || 100 < vTraders.Amount {
				continue
			}

			// 绑定
			if _, ok := bindTrader[vTraders.ID]; ok {
				continue
			}

			bindTrader[vTraders.ID] = vTraders
			tmpCost -= vTraders.Amount
		}

		// 上述待绑定交易员结果集
		if 0 >= len(bindTrader) {
			continue
		}

		// 新增 和 更新
		insertUserBindTrader := make([]*UserBindTrader, 0)
		for k, v := range bindTrader {

			insertUserBindTrader = append(insertUserBindTrader, &UserBindTrader{
				UserId:   vUsers.ID,
				TraderId: k,
				Amount:   v.Amount,
			})
		}

		// 写入
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.UpdateUserBindTraderStatus(ctx, vUsers.ID)
			if nil != err {
				return err
			}

			for _, vInsertUserBindTrader := range insertUserBindTrader {
				_, err = b.binanceUserRepo.InsertUserBindTrader(ctx, vInsertUserBindTrader.UserId, vInsertUserBindTrader.TraderId, vInsertUserBindTrader.Amount)
				if nil != err {
					return err
				}
			}

			return nil
		}); err != nil {
			fmt.Println("换绑交易员写入异常", vUsers, err)
			continue
		}
	}

	return nil
}

func (b *BinanceUserUsecase) ReBindTrader(ctx context.Context) error {

	var (
		userBalance       map[uint64]*UserBalance
		userBindTraderMap map[uint64][]*UserBindTrader
		users             []*User
		traders           []*Trader
		symbols           map[string]*Symbol
		err               error
	)

	// 获取需要更新绑定交易员的用户
	users, err = b.binanceUserRepo.GetUsersByBindUserStatusReBind()
	if nil != err {
		return err
	}

	if 0 == len(users) {
		return nil
	}

	// 获取用户cost
	userIds := make([]uint64, 0)
	for _, vUsers := range users {
		userIds = append(userIds, vUsers.ID)
	}
	userBalance, err = b.binanceUserRepo.GetUserBalanceByUserIds(ctx, userIds)
	if nil != err {
		return err
	}

	// 获取用户绑定记录
	userBindTraderMap, err = b.binanceUserRepo.GetUserBindTraderMapByUserIds(userIds)
	if nil != err {
		return err
	}

	// 按amount排序好的交易员
	traders, err = b.binanceUserRepo.GetTradersOrderByAmountDesc()
	if nil != err {
		return err
	}

	if 0 == len(traders) {
		return nil
	}

	// 代币
	symbols, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return err
	}

	// 1. 遍历，轮询用户
	// 2. 轮询交易员（查询时已经按amount desc排序，用户cost的百分之30>=当前交易员的amount字段，建立绑定关系。
	// 3. 第二轮保留最后的交易员额度限制，将此和用户剩余可用额度对比，按顺序分配给此额度及以下的交易员。
	// 4. 轮询结束时，用户额度剩余不予理会。
	for _, vUsers := range users {
		if _, exists := userBalance[vUsers.ID]; !exists {
			continue
		}

		tmpCost := userBalance[vUsers.ID].Cost

		// 不存在绑定记录, 新增
		if _, exists := userBindTraderMap[vUsers.ID]; !exists {
			// 新增
			bindTrader := make(map[uint64]*Trader, 0)
			// 第一轮
			for _, vTraders := range traders {
				if 0 >= tmpCost {
					break
				}

				if 0 >= vTraders.Amount {
					continue
				}

				if 100 >= vTraders.Amount {
					continue
				}

				if tmpCost*30/100 >= vTraders.Amount {
					// 绑定
					if _, ok := bindTrader[vTraders.ID]; ok {
						continue
					}

					bindTrader[vTraders.ID] = vTraders
					tmpCost -= vTraders.Amount
				}
			}

			// 第二轮，跳过分配限制的额度，剩下的按顺序分配
			for _, vTraders := range traders {
				if 100 > tmpCost {
					break
				}

				if 0 >= vTraders.Amount || 100 < vTraders.Amount {
					continue
				}

				// 绑定
				if _, ok := bindTrader[vTraders.ID]; ok {
					continue
				}

				bindTrader[vTraders.ID] = vTraders
				tmpCost -= vTraders.Amount
			}

			// 上述待绑定交易员结果集
			if 0 >= len(bindTrader) {
				continue
			}

			insertUserBindTrader := make([]*UserBindTrader, 0)
			for k, v := range bindTrader {

				insertUserBindTrader = append(insertUserBindTrader, &UserBindTrader{
					UserId:   vUsers.ID,
					TraderId: k,
					Amount:   v.Amount,
				})
			}

			// 写入
			if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = b.binanceUserRepo.UpdateUserBindTraderStatus(ctx, vUsers.ID)
				if nil != err {
					return err
				}

				for _, vInsertUserBindTrader := range insertUserBindTrader {
					_, err = b.binanceUserRepo.InsertUserBindTrader(ctx, vInsertUserBindTrader.UserId, vInsertUserBindTrader.TraderId, vInsertUserBindTrader.Amount)
					if nil != err {
						return err
					}
				}

				return nil
			}); err != nil {
				fmt.Println("换绑交易员写入异常", vUsers, err)
			}

		} else {

			// 存在正在绑定记录，检查
			var (
				bindCost               uint64
				alreadyBindTraderCount uint64
			)

			alreadyBindTrader := make(map[uint64]uint64, 0)
			selfUpdateBindTrader := make(map[uint64]*UserBindTrader, 0)
			for _, vUserBindTraderMap := range userBindTraderMap[vUsers.ID] {
				if 0 == vUserBindTraderMap.Status {
					bindCost += vUserBindTraderMap.Amount
					alreadyBindTraderCount++
				}

				// 因为额度调整换绑的不算
				if 3 == vUserBindTraderMap.Status {
					selfUpdateBindTrader[vUserBindTraderMap.TraderId] = vUserBindTraderMap
					continue
				}

				alreadyBindTrader[vUserBindTraderMap.TraderId] = vUserBindTraderMap.TraderId
			}

			// 需要新增
			if bindCost < tmpCost {
				tmpCost -= bindCost

				bindTrader := make(map[uint64]*Trader, 0)

				// 第一轮
				for _, vTraders := range traders {
					// 已经绑定过的交易员
					if _, ok := alreadyBindTrader[vTraders.ID]; ok {
						continue
					}

					if 0 >= tmpCost {
						break
					}

					if 0 >= vTraders.Amount {
						continue
					}

					if 100 >= vTraders.Amount {
						continue
					}

					if tmpCost*30/100 >= vTraders.Amount {
						// 绑定
						if _, ok := bindTrader[vTraders.ID]; ok {
							continue
						}

						bindTrader[vTraders.ID] = vTraders
						tmpCost -= vTraders.Amount
					}
				}

				// 第二轮，跳过分配限制的额度，剩下的按顺序分配
				for _, vTraders := range traders {
					// 已经绑定过的交易员
					if _, ok := alreadyBindTrader[vTraders.ID]; ok {
						continue
					}

					if 100 > tmpCost {
						break
					}

					if 0 >= vTraders.Amount || 100 < vTraders.Amount {
						continue
					}

					// 绑定
					if _, ok := bindTrader[vTraders.ID]; ok {
						continue
					}

					bindTrader[vTraders.ID] = vTraders
					tmpCost -= vTraders.Amount
				}

				// 上述待绑定交易员结果集
				if 0 >= len(bindTrader) {
					continue
				}

				insertUserBindTrader := make([]*UserBindTrader, 0)
				for k, v := range bindTrader {
					insertUserBindTrader = append(insertUserBindTrader, &UserBindTrader{
						UserId:   vUsers.ID,
						TraderId: k,
						Amount:   v.Amount,
					})
				}

				// 写入
				if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
					_, err = b.binanceUserRepo.UpdateUserBindTraderStatus(ctx, vUsers.ID)
					if nil != err {
						return err
					}

					for _, vInsertUserBindTrader := range insertUserBindTrader {
						// 存在的写入，不存在更新
						if _, ok := selfUpdateBindTrader[vInsertUserBindTrader.TraderId]; ok {
							_, err = b.binanceUserRepo.UpdatesUserBindTraderStatusAndInitOrderById(ctx, selfUpdateBindTrader[vInsertUserBindTrader.TraderId].ID, 0, 0, selfUpdateBindTrader[vInsertUserBindTrader.TraderId].Amount)
						} else {
							_, err = b.binanceUserRepo.InsertUserBindTrader(ctx, vInsertUserBindTrader.UserId, vInsertUserBindTrader.TraderId, vInsertUserBindTrader.Amount)
						}

						if nil != err {
							return err
						}
					}

					return nil
				}); err != nil {
					fmt.Println("换绑交易员写入异常", vUsers, err)
				}

				continue
			} else if bindCost > tmpCost {
				// 需要减掉
				tmpMoreCost := bindCost - tmpCost
				updateBindTrader := make([]*UserBindTrader, 0)
				// 查询的结果已经排序好按amount
				for _, vUserBindTraderMap := range userBindTraderMap[vUsers.ID] {
					if 0 == vUserBindTraderMap.Status && tmpMoreCost >= vUserBindTraderMap.Amount {
						if 0 >= tmpMoreCost {
							break
						} else if tmpMoreCost >= vUserBindTraderMap.Amount {
							tmpMoreCost -= vUserBindTraderMap.Amount
							updateBindTrader = append(updateBindTrader, vUserBindTraderMap)
						}
					}
				}

				// 更新未status=1不再绑定的状态
				for _, vUpdateBindTrader := range updateBindTrader {
					// 解绑后的还会平单的数据
					userBindAfterUnbind := make([]*UserBindAfterUnbind, 0)
					if 1 == vUpdateBindTrader.InitOrder {
						// 添加未关单数据到后续表
						var (
							currentOrders map[string][]*UserOrder
						)
						currentOrders, err = b.binanceUserRepo.GetUserOrderByUserTraderId(vUpdateBindTrader.UserId, vUpdateBindTrader.TraderId)
						if nil != err {
							continue
						}
						if 0 < len(currentOrders) {
							// 用户下遍历对应的交易员
							for symbol, vCurrentOrders := range currentOrders {
								// 精度
								if _, ok := symbols[symbol]; !ok {
									continue
								}

								var (
									historyQuantityFloatMore  float64
									historyQuantityFloatEmpty float64
								)

								for _, vVCurrentOrders := range vCurrentOrders {

									// 历史开多和平多
									if "BUY" == vVCurrentOrders.Side && "LONG" == vVCurrentOrders.PositionSide {
										historyQuantityFloatMore += vVCurrentOrders.ExecutedQty
									} else if "SELL" == vVCurrentOrders.Side && "LONG" == vVCurrentOrders.PositionSide {
										historyQuantityFloatMore -= vVCurrentOrders.ExecutedQty
									}

									// 历史开空和平空
									if "SELL" == vVCurrentOrders.Side && "SHORT" == vVCurrentOrders.PositionSide {
										historyQuantityFloatEmpty += vVCurrentOrders.ExecutedQty
									} else if "BUY" == vVCurrentOrders.Side && "SHORT" == vVCurrentOrders.PositionSide {
										historyQuantityFloatEmpty -= vVCurrentOrders.ExecutedQty
									}
								}

								// 当前代币的多单还剩多少, 有的记录
								if !isZero(historyQuantityFloatMore) && 0 < historyQuantityFloatMore {
									userBindAfterUnbind = append(userBindAfterUnbind, &UserBindAfterUnbind{
										UserId:       vUpdateBindTrader.UserId,
										TraderId:     vUpdateBindTrader.TraderId,
										Amount:       vUpdateBindTrader.Amount,
										Symbol:       symbol,
										PositionSide: "LONG",
										Quantity:     historyQuantityFloatMore,
									})
								}

								// 当前代币的空单还剩多少, 有的记录
								if !isZero(historyQuantityFloatEmpty) && 0 < historyQuantityFloatEmpty {
									userBindAfterUnbind = append(userBindAfterUnbind, &UserBindAfterUnbind{
										UserId:       vUpdateBindTrader.UserId,
										TraderId:     vUpdateBindTrader.TraderId,
										Amount:       vUpdateBindTrader.Amount,
										Symbol:       symbol,
										PositionSide: "SHORT",
										Quantity:     historyQuantityFloatEmpty,
									})
								}
							}
						}
					}

					// 写入
					if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
						_, err = b.binanceUserRepo.UpdatesUserBindTraderStatusById(ctx, vUpdateBindTrader.ID, 3)
						if nil != err {
							return err
						}

						// 放置平仓表
						for _, vUserBindAfterUnbind := range userBindAfterUnbind {
							_, err = b.binanceUserRepo.InsertUserBindAfterUnbind(ctx, vUserBindAfterUnbind)
							if nil != err {
								return err
							}
						}

						return nil
					}); err != nil {
						fmt.Println("用户额度变更换绑失败", err, vUpdateBindTrader)
					}
				}

				// 写入
				if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
					_, err = b.binanceUserRepo.UpdateUserBindTraderStatus(ctx, vUsers.ID)
					if nil != err {
						return err
					}

					return nil
				}); err != nil {
					fmt.Println("用户额度变更换绑失败", err)
				}

				continue
			} else {
				continue
			}
		}

	}

	return nil
}

func (b *BinanceUserUsecase) ChangeBindTrader(ctx context.Context) error {
	var (
		userBindTrader map[uint64][]*UserBindTrader
		traders        map[uint64]*Trader
		symbols        map[string]*Symbol
		err            error
	)

	// 获取待更换的绑定关系
	userBindTrader, err = b.binanceUserRepo.GetUserBindTraderByStatus(2)
	if nil != err {
		return err
	}

	if 0 == len(userBindTrader) {
		return nil
	}

	// 代币
	symbols, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return err
	}

	// 全部交易员
	traders, err = b.binanceUserRepo.GetTraders()
	if nil != err {
		return err
	}

	// 每个换绑的用户，换一个未绑定的同等amount的可用的交易员
	for traderId, vUserBindTrader := range userBindTrader {
		// 获取amount相同的traders
		if _, ok := traders[traderId]; !ok {
			fmt.Println("无法换绑，不存在交易员", traderId)
			continue
		}

		// 可替换的，去除当前的
		okTraders := make([]*Trader, 0)
		for _, vTraders := range traders {
			if traderId != vTraders.ID && vTraders.Amount == traders[traderId].Amount && 1 == vTraders.Status {
				okTraders = append(okTraders, vTraders)
			}
		}

		// 遍历需要换绑的用户
		for _, vVUserBindTrader := range vUserBindTrader {
			var (
				userBindTradersAll []*UserBindTrader
				insertTrader       *Trader
			)

			// 已经绑定的
			userBindTradersAll, err = b.binanceUserRepo.GetUserBindTraderByUserId(vVUserBindTrader.UserId)
			if nil != err {
				continue
			}
			userBindTradersAllMap := make(map[uint64]*UserBindTrader, 0)
			selfUpdateBindTrader := make(map[uint64]uint64, 0)
			for _, vUserBindTradersAll := range userBindTradersAll {
				// 因为额度调整，自己换绑的不算
				if 3 == vUserBindTradersAll.Status {
					selfUpdateBindTrader[vUserBindTradersAll.TraderId] = vUserBindTradersAll.ID
					continue
				}

				userBindTradersAllMap[vUserBindTradersAll.TraderId] = vUserBindTradersAll
			}

			// 差集，未绑定过的
			for _, vOkTraders := range okTraders {
				if _, ok := userBindTradersAllMap[vOkTraders.ID]; !ok {
					insertTrader = vOkTraders // 新增
					break
				}
			}

			if nil == insertTrader {
				continue
			}

			// 解绑后的还会平单的数据
			userBindAfterUnbind := make([]*UserBindAfterUnbind, 0)
			// 添加未关单数据到后续表
			if 1 == vVUserBindTrader.InitOrder {
				var (
					currentOrders map[string][]*UserOrder
				)
				currentOrders, err = b.binanceUserRepo.GetUserOrderByUserTraderId(vVUserBindTrader.UserId, vVUserBindTrader.TraderId)
				if nil != err {
					continue
				}
				if 0 < len(currentOrders) {
					// 用户下遍历对应的交易员
					for symbol, vCurrentOrders := range currentOrders {
						// 精度
						if _, ok := symbols[symbol]; !ok {
							continue
						}

						var (
							historyQuantityFloatMore  float64
							historyQuantityFloatEmpty float64
						)

						for _, vVCurrentOrders := range vCurrentOrders {

							// 历史开多和平多
							if "BUY" == vVCurrentOrders.Side && "LONG" == vVCurrentOrders.PositionSide {
								historyQuantityFloatMore += vVCurrentOrders.ExecutedQty
							} else if "SELL" == vVCurrentOrders.Side && "LONG" == vVCurrentOrders.PositionSide {
								historyQuantityFloatMore -= vVCurrentOrders.ExecutedQty
							}

							// 历史开空和平空
							if "SELL" == vVCurrentOrders.Side && "SHORT" == vVCurrentOrders.PositionSide {
								historyQuantityFloatEmpty += vVCurrentOrders.ExecutedQty
							} else if "BUY" == vVCurrentOrders.Side && "SHORT" == vVCurrentOrders.PositionSide {
								historyQuantityFloatEmpty -= vVCurrentOrders.ExecutedQty
							}
						}

						// 当前代币的多单还剩多少, 有的记录
						if !isZero(historyQuantityFloatMore) && 0 < historyQuantityFloatMore {
							userBindAfterUnbind = append(userBindAfterUnbind, &UserBindAfterUnbind{
								UserId:       vVUserBindTrader.UserId,
								TraderId:     vVUserBindTrader.TraderId,
								Amount:       vVUserBindTrader.Amount,
								Symbol:       symbol,
								PositionSide: "LONG",
								Quantity:     historyQuantityFloatMore,
							})
						}

						// 当前代币的空单还剩多少, 有的记录
						if !isZero(historyQuantityFloatEmpty) && 0 < historyQuantityFloatEmpty {
							userBindAfterUnbind = append(userBindAfterUnbind, &UserBindAfterUnbind{
								UserId:       vVUserBindTrader.UserId,
								TraderId:     vVUserBindTrader.TraderId,
								Amount:       vVUserBindTrader.Amount,
								Symbol:       symbol,
								PositionSide: "SHORT",
								Quantity:     historyQuantityFloatEmpty,
							})
						}

					}
				}
			}

			// 写入
			if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = b.binanceUserRepo.UpdatesUserBindTraderStatusById(ctx, vVUserBindTrader.ID, 1)
				if nil != err {
					return err
				}

				// 存在则更新
				if _, ok := selfUpdateBindTrader[insertTrader.ID]; ok {
					_, err = b.binanceUserRepo.UpdatesUserBindTraderStatusAndInitOrderById(ctx, selfUpdateBindTrader[insertTrader.ID], 0, 0, insertTrader.Amount)
				} else {
					_, err = b.binanceUserRepo.InsertUserBindTrader(ctx, vVUserBindTrader.UserId, insertTrader.ID, vVUserBindTrader.Amount)
				}

				if nil != err {
					return err
				}

				// 放置平仓表
				for _, vUserBindAfterUnbind := range userBindAfterUnbind {
					_, err = b.binanceUserRepo.InsertUserBindAfterUnbind(ctx, vUserBindAfterUnbind)
					if nil != err {
						return err
					}
				}

				return nil
			}); err != nil {
				return err
			}

		}

	}

	return nil
}

// Analyze 测试方法
func (b *BinanceUserUsecase) Analyze(ctx context.Context, req *v1.AnalyzeRequest) (*v1.AnalyzeReply, error) {
	//var (
	//	orders     []*UserOrder
	//	ordersUser map[string]*UserOrder
	//	order1     map[string]*UserOrder
	//	total      int64
	//	u          float64
	//	totalU     int64
	//	u2         float64
	//	totalU2    int64
	//	err        error
	//)
	//
	//orders, err = b.binanceUserRepo.GetUserOrderByUserIdGroupBySymbol(4)
	//if nil != err {
	//	return nil, nil
	//}
	//
	//ordersUser, err = b.binanceUserRepo.GetUserOrderByUserIdMapId(4)
	//if nil != err {
	//	return nil, nil
	//}
	//
	//order1 = make(map[string]*UserOrder, 0)
	//
	//for _, v := range orders {
	//	var (
	//		binanceOrder []*OrderHistory
	//		start        = int64(1707346800000)
	//	)
	//
	//	for start <= 1708337023116 {
	//
	//		var (
	//			tmpBinanceOrder []*OrderHistory
	//			tmpStart        string
	//			end             string
	//		)
	//
	//		tmpStart = strconv.FormatInt(start, 10)
	//		if start+432000000 >= 1708337023116 {
	//			end = strconv.FormatInt(1708337023116, 10)
	//		} else {
	//			end = strconv.FormatInt(start+432000000, 10)
	//		}
	//
	//		//fmt.Println(v.Symbol, tmpStart, end)
	//		tmpBinanceOrder, err = requestBinanceOrderHistory(
	//			"Bc8Id20hm7sBIrKewsRAvA4RX2MPmy546X6q1wOGlsBQsSoRdn6Uze6lHAim8YY1",
	//			"DAzqoU9KbrQfxrWNMKdbNYgtWSxZpFVPSRPxeN1k08ed7cNGTACc7dy9oK8eZwrG",
	//			v.Symbol,
	//			0,
	//			tmpStart, end)
	//		if nil != err {
	//			fmt.Println("错误查询", err, v.Symbol)
	//		}
	//
	//		for _, vTmpBinanceOrder := range tmpBinanceOrder {
	//			binanceOrder = append(binanceOrder, vTmpBinanceOrder)
	//		}
	//
	//		start += 432000000
	//
	//	}
	//
	//	for _, vBinanceOrder := range binanceOrder {
	//		tmpOrderId := strconv.FormatInt(vBinanceOrder.OrderId, 10)
	//		//if _, ok := ordersUser[tmpOrderId]; !ok {
	//		//	fmt.Println("不存在系统下单", tmpOrderId)
	//		//}
	//
	//		total++
	//		// 平多
	//		if "SELL" == vBinanceOrder.Side && "LONG" == vBinanceOrder.PositionSide {
	//			var tmp float64
	//			tmp, err = strconv.ParseFloat(vBinanceOrder.RealizedPnl, 64)
	//			if nil != err {
	//				return nil, err
	//			}
	//
	//			fmt.Println("收益此时：", u, vBinanceOrder.RealizedPnl, tmp, vBinanceOrder)
	//
	//			u += tmp
	//			totalU++
	//
	//			if _, ok := ordersUser[tmpOrderId]; ok {
	//				u2 += tmp
	//
	//				if _, ok2 := order1[tmpOrderId]; !ok2 {
	//					totalU2++
	//					order1[tmpOrderId] = ordersUser[tmpOrderId]
	//				}
	//			}
	//		}
	//
	//		// 平空
	//		if "BUY" == vBinanceOrder.Side && "SHORT" == vBinanceOrder.PositionSide {
	//			var tmp float64
	//			tmp, err = strconv.ParseFloat(vBinanceOrder.RealizedPnl, 64)
	//			if nil != err {
	//				return nil, err
	//			}
	//
	//			fmt.Println("收益此时：", u, vBinanceOrder.RealizedPnl, tmp, vBinanceOrder)
	//
	//			u += tmp
	//			totalU++
	//
	//			if _, ok := ordersUser[tmpOrderId]; ok {
	//				fmt.Println("系统收益此时：", u2)
	//
	//				u2 += tmp
	//
	//				if _, ok2 := order1[tmpOrderId]; !ok2 {
	//
	//					totalU2++
	//
	//					order1[tmpOrderId] = ordersUser[tmpOrderId]
	//				}
	//			}
	//		}
	//
	//		// 开多
	//		if "BUY" == vBinanceOrder.Side && "LONG" == vBinanceOrder.PositionSide {
	//
	//		}
	//
	//		// 开空
	//		if "SELL" == vBinanceOrder.Side && "SHORT" == vBinanceOrder.PositionSide {
	//
	//		}
	//	}
	//
	//}
	//
	//fmt.Println("共单数：", total, "收益：", u, "平仓单共：", totalU, "系统订单：", len(ordersUser), "系统收益：", u2, "系统收益共", totalU2, "a", len(order1))
	var (
		userOrders []*UserOrder
		err        error
	)
	userOrders, err = b.binanceUserRepo.GetUserOrderByHandleStatus()
	if nil != err {
		return nil, err
	}

	for _, v := range userOrders {
		err = b.binanceUserRepo.SAddOrderSetSellLongOrBuyShort(ctx, int64(v.ID))
		if nil != err {
			fmt.Println("错误的数据写入, redis", v, err)
		}
	}

	return nil, nil
}

// ListenTraders 用户下单
func (b *BinanceUserUsecase) ListenTraders(ctx context.Context, req *v1.ListenTraderAndUserOrderRequest) (*v1.ListenTraderAndUserOrderReply, error) {
	var (
		wg sync.WaitGroup
	)

	fmt.Println("老系统", req.SendBody)

	wg.Add(1) // 启动一个goroutine就登记+1
	go b.ListenTradersHandle(ctx, req, &wg)

	wg.Add(1) // 启动一个goroutine就登记+1
	go b.ListenTradersHandleTwo(ctx, req, &wg)

	wg.Wait() // 等待所有登记的goroutine都结束

	return &v1.ListenTraderAndUserOrderReply{
		Status: "ok",
	}, nil
}

// ListenTradersNew 用户下单
func (b *BinanceUserUsecase) ListenTradersNew(ctx context.Context, req *v1.ListenTraderAndUserOrderRequest) (*v1.ListenTraderAndUserOrderReply, error) {
	var (
		wg sync.WaitGroup
	)
	fmt.Println("新系统", req.SendBody)

	wg.Add(1) // 启动一个goroutine就登记+1
	go b.ListenTradersHandleTwo(ctx, req, &wg)
	wg.Wait() // 等待所有登记的goroutine都结束

	return &v1.ListenTraderAndUserOrderReply{
		Status: "ok",
	}, nil
}

// ListenTradersHandle 用户下单
func (b *BinanceUserUsecase) ListenTradersHandle(ctx context.Context, req *v1.ListenTraderAndUserOrderRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	var (
		traderIds      []uint64
		userBindTrader map[uint64][]*UserBindTrader
		userIds        []uint64
		userIdsMap     map[uint64]uint64
		users          map[uint64]*User
		userBalance    map[uint64]*UserBalance
		userAmount     map[uint64]*UserAmount
		symbol         map[string]*Symbol
		err            error
	)

	traderIds = make([]uint64, 0)
	for _, vOrders := range req.SendBody.Orders {
		if 0 < vOrders.Uid {
			traderIds = append(traderIds, vOrders.Uid)
		}
	}

	if 0 < len(traderIds) {
		userBindTrader, err = b.binanceUserRepo.GetUserBindTraderByTraderIds(traderIds)
		if nil != err {
			return
		}
	} else {
		return
	}

	userIds = make([]uint64, 0)
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserBindTrader := range userBindTrader {
		for _, vVUserBindTrader := range vUserBindTrader {
			// 初始化仓位
			if 1 != vVUserBindTrader.InitOrder {
				continue
			}

			if _, ok := userIdsMap[vVUserBindTrader.UserId]; ok {
				continue
			}

			userIdsMap[vVUserBindTrader.UserId] = vVUserBindTrader.UserId
			userIds = append(userIds, vVUserBindTrader.UserId)
		}
	}

	if 0 >= len(userIds) {
		return
	}

	// 获取用户信息，余额信息，收益信息
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		return
	}

	userBalance, err = b.binanceUserRepo.GetUserBalanceByUserIds(ctx, userIds)
	if nil != err {
		return
	}

	userAmount, err = b.binanceUserRepo.GetUserAmountByUserIds(ctx, userIds)
	if nil != err {
		return
	}

	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return
	}

	for _, vOrders := range req.SendBody.Orders {
		for _, vOrdersData := range vOrders.Data {
			if _, exists := userBindTrader[vOrders.Uid]; !exists {
				continue
			}

			for _, vUserBindTrader := range userBindTrader[vOrders.Uid] {
				if 0 == vUserBindTrader.Status { // 绑定
					// 初始化仓位
					if 1 != vUserBindTrader.InitOrder {
						continue
					}

					if _, ok := users[vUserBindTrader.UserId]; !ok {
						continue
					}

					// todo 暂时使用检测没用api信息
					if 0 >= users[vUserBindTrader.UserId].ApiStatus {
						continue
					}

					if _, ok := userBalance[vUserBindTrader.UserId]; !ok {
						continue
					}

					if _, ok := userAmount[vUserBindTrader.UserId]; !ok {
						continue
					}

					// 判断是开单还是关单，sell long 关多 buy short 关空
					if ("SELL" == vOrdersData.Side && "LONG" == vOrdersData.Type) || ("BUY" == vOrdersData.Side && "SHORT" == vOrdersData.Type) {

					} else if ("SELL" == vOrdersData.Side && "SHORT" == vOrdersData.Type) || ("BUY" == vOrdersData.Side && "LONG" == vOrdersData.Type) {
						// todo 暂时只关不开
						if 0 == users[vUserBindTrader.UserId].IsDai {
							if 5 != vOrders.Uid && 83 != vOrders.Uid {
								continue
							}
						}

						// 精度按代币18位，截取小数点后到5位计算
						var balanceTmp int64
						lengthToKeep := len(userBalance[vUserBindTrader.UserId].Balance) - 13

						if lengthToKeep > 0 {
							balanceTmpStr := userBalance[vUserBindTrader.UserId].Balance[:lengthToKeep]
							balanceTmp, err = strconv.ParseInt(balanceTmpStr, 10, 64)
							if nil != err || 0 >= balanceTmp {
								continue
							}
						} else {
							continue
						}

						// 余额不足，收益大于余额的10倍
						if userAmount[vUserBindTrader.UserId].Amount > balanceTmp*10 {
							continue
						}

					} else {
						continue
					}

					// 精度
					if _, ok := symbol[vOrdersData.Symbol]; !ok {
						continue
					}

					// 发送订单
					wg.Add(1) // 启动一个goroutine就登记+1
					go b.userOrderGoroutine(ctx, wg, &OrderData{
						Coin:     vOrdersData.Symbol,
						Type:     vOrdersData.Type,
						Price:    vOrdersData.Price,
						Side:     vOrdersData.Side,
						Qty:      vOrdersData.Qty,
						Position: vOrdersData.Position,
					}, vOrders.BaseMoney, users[vUserBindTrader.UserId], vUserBindTrader, symbol[vOrdersData.Symbol].QuantityPrecision, 0, 0, 0)
				}
			}
		}
	}
}

func (b *BinanceUserUsecase) userOrderGoroutine(ctx context.Context, wg *sync.WaitGroup, order *OrderData, amount string, user *User, userBindTrader *UserBindTrader, quantityPrecision int64, initOrderReq uint64, proportion float64, overOrderReq uint64) {
	defer wg.Done() // goroutine结束就登记-1

	var (
		binanceOrder  *BinanceOrder
		orderType     = "MARKET"
		positionSide  string
		quantity      string
		qty           float64
		price         float64
		traderAmount  float64
		currentOrders []*UserOrder
		currentOrder  *UserOrder
		insertOrder   *UserOrder
		err           error
	)

	// 新订单数据
	currentOrder = &UserOrder{
		UserId:        userBindTrader.UserId,
		TraderId:      userBindTrader.TraderId,
		Symbol:        order.Coin,
		Side:          "",
		PositionSide:  "",
		Quantity:      0,
		Price:         0,
		TraderQty:     0,
		OrderType:     orderType,
		ClosePosition: "",
		CumQuote:      0,
		ExecutedQty:   0,
		AvgPrice:      0,
	}

	if 0 == overOrderReq { // 非全平仓下单
		qty, err = strconv.ParseFloat(order.Qty, 64)
		if nil != err {
			fmt.Println(err, order, userBindTrader)
			return
		}
		currentOrder.TraderQty = qty

		price, err = strconv.ParseFloat(order.Price, 64)
		if nil != err {
			fmt.Println(err, order, userBindTrader)
			return
		}
		currentOrder.Price = price

		traderAmount, err = strconv.ParseFloat(amount, 64)
		if nil != err {
			fmt.Println(err, order, userBindTrader)
			return
		}
	}

	if "LONG" == order.Type {
		positionSide = "LONG"
	} else if "SHORT" == order.Type {
		positionSide = "SHORT" // 空
	} else {
		fmt.Println("err position side")
		return
	}
	currentOrder.PositionSide = positionSide

	var (
		quantityFloat float64
	)
	if 1 == initOrderReq { // 初始化下单
		if 0 > currentOrder.Price {
			fmt.Println("err price", currentOrder.Price, order, userBindTrader)
			return
		}

		quantityFloat = float64(userBindTrader.Amount) * proportion / currentOrder.Price
	} else if 1 == overOrderReq { // 全平仓

	} else {
		quantityFloat = float64(userBindTrader.Amount) * qty / traderAmount // 本次开单数量
	}

	var historyQuantityFloat float64
	// 本次平仓
	if ("SELL" == order.Side && "LONG" == order.Type) || ("BUY" == order.Side && "SHORT" == order.Type) {
		// 订单统计
		currentOrders, err = b.binanceUserRepo.GetUserOrderByUserTraderIdAndSymbol(userBindTrader.UserId, userBindTrader.TraderId, order.Coin)
		if nil != err {
			fmt.Println(err, order)
			return
		}
		// 查出用户的BUY单的币的数量，在对应的trader下，超过了BUY不能SELL todo 使用数据库量太大以后
		if 0 >= len(currentOrders) {
			return
		}

		// 多的部分不管，按剩余的数量关 todo 交易员全部平仓，少的部分另一个程序解决
		for _, vCurrentOrders := range currentOrders {
			// 本次平多
			if "SELL" == order.Side && "LONG" == order.Type {
				// 历史开多和平多
				if "BUY" == vCurrentOrders.Side && "LONG" == vCurrentOrders.PositionSide {
					historyQuantityFloat += vCurrentOrders.ExecutedQty
				} else if "SELL" == vCurrentOrders.Side && "LONG" == vCurrentOrders.PositionSide {
					historyQuantityFloat -= vCurrentOrders.ExecutedQty
				}
			} else if "BUY" == order.Side && "SHORT" == order.Type {
				// 历史开空和平空
				if "SELL" == vCurrentOrders.Side && "SHORT" == vCurrentOrders.PositionSide {
					historyQuantityFloat += vCurrentOrders.ExecutedQty
				} else if "BUY" == vCurrentOrders.Side && "SHORT" == vCurrentOrders.PositionSide {
					historyQuantityFloat -= vCurrentOrders.ExecutedQty
				}
			}
		}

		// 平单历史数量不足了
		if isZero(historyQuantityFloat) || 0 > historyQuantityFloat {
			fmt.Println("trader的开单数量小于关单数量了，可能是精度问题", order.Coin, userBindTrader.UserId, userBindTrader.TraderId, historyQuantityFloat)
			if 1 == overOrderReq {
				fmt.Println("全平")
			}

			return
		}

		if 1 == overOrderReq { // 全平仓
			quantityFloat = historyQuantityFloat
		} else {
			// 按仓位百分比平仓
			if 0 < len(order.Position) {
				var (
					tmpPositionTotal float64
				)

				tmpPositionTotal, err = strconv.ParseFloat(order.Position, 64)
				if nil != err {
					fmt.Println(err, order, userBindTrader)
					return
				}

				if 0 < tmpPositionTotal && 0 < qty {
					quantityFloat = historyQuantityFloat * qty / tmpPositionTotal // 本次平仓数量
				}
			}

			// 超过了净开单数量
			// todo 并发操作时数据不一致的可能性，会导致数量对不上，例如下单程序和更换绑定程序同时执行时，是否程序中统计的总数字会漏计算这里的新增的订单的数字
			if quantityFloat > historyQuantityFloat {
				quantityFloat = historyQuantityFloat
			}
		}

	} else if ("SELL" == order.Side && "SHORT" == order.Type) || ("BUY" == order.Side && "LONG" == order.Type) {
		// 开仓

	} else {
		fmt.Println("err order side")
		return
	}

	currentOrder.Side = order.Side

	// 精度调整
	if 0 >= quantityPrecision {
		quantity = fmt.Sprintf("%d", int64(quantityFloat))
	} else {
		quantity = strconv.FormatFloat(quantityFloat, 'f', int(quantityPrecision), 64)
	}

	currentOrder.Quantity, err = strconv.ParseFloat(quantity, 64)
	if nil != err {
		fmt.Println(err)
		return
	}

	if 0 >= currentOrder.Quantity {
		fmt.Println("下单数值太小", quantity, currentOrder, quantityPrecision)
		return
	}

	// 下单信息
	var orderInfo *OrderInfo
	// 请求下单
	binanceOrder, orderInfo, err = requestBinanceOrder(order.Coin, order.Side, orderType, positionSide, quantity, user.ApiKey, user.ApiSecret)
	if nil != err {
		fmt.Println(err)
		return
	}

	// 下单异常
	if 0 >= binanceOrder.OrderId {
		// 写入
		orderErr := &UserOrderErr{
			UserId:        currentOrder.UserId,
			TraderId:      currentOrder.TraderId,
			ClientOrderId: currentOrder.ClientOrderId,
			OrderId:       currentOrder.OrderId,
			Symbol:        currentOrder.Symbol,
			Side:          currentOrder.Side,
			PositionSide:  currentOrder.PositionSide,
			Quantity:      quantityFloat,
			Price:         currentOrder.Price,
			TraderQty:     currentOrder.TraderQty,
			OrderType:     currentOrder.OrderType,
			ClosePosition: currentOrder.ClosePosition,
			CumQuote:      currentOrder.CumQuote,
			ExecutedQty:   currentOrder.ExecutedQty,
			AvgPrice:      currentOrder.AvgPrice,
			HandleStatus:  currentOrder.HandleStatus,
			InitOrder:     int64(initOrderReq),
			Code:          orderInfo.Code,
			Msg:           orderInfo.Msg,
			Proportion:    proportion,
		}

		if 1 == overOrderReq {
			orderErr.InitOrder = 2
		}

		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.InsertUserOrderErr(ctx, orderErr)
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			fmt.Println(err, orderErr)
			return
		}

		return // 返回
	}

	currentOrder.OrderId = strconv.FormatInt(binanceOrder.OrderId, 10)

	currentOrder.CumQuote, err = strconv.ParseFloat(binanceOrder.CumQuote, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	currentOrder.ExecutedQty, err = strconv.ParseFloat(binanceOrder.ExecutedQty, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	currentOrder.AvgPrice, err = strconv.ParseFloat(binanceOrder.AvgPrice, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	// 写入
	if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
		insertOrder, err = b.binanceUserRepo.InsertUserOrder(ctx, currentOrder)
		if nil != err {
			return err
		}

		return nil
	}); err != nil {
		fmt.Println(err, currentOrder)
		return
	}

	// 计算收益 卖单去binance查询本单收益
	if ("SELL" == order.Side && "LONG" == order.Type) || ("BUY" == order.Side && "SHORT" == order.Type) {
		if 0 >= insertOrder.ID {
			fmt.Println("错误的数据写入", insertOrder)
			return
		}

		err = b.binanceUserRepo.SAddOrderSetSellLongOrBuyShort(ctx, int64(insertOrder.ID))
		if nil != err {
			fmt.Println("错误的数据写入, redis", insertOrder, err)
			return
		}
	}

	return
}

// ListenTradersHandleTwo 用户下单 tfi
func (b *BinanceUserUsecase) ListenTradersHandleTwo(ctx context.Context, req *v1.ListenTraderAndUserOrderRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	var (
		newSystemReq   bool
		traderIds      []uint64
		traderNums     []uint64
		userBindTrader map[uint64][]*UserBindTrader
		userIds        []uint64
		userIdsMap     map[uint64]uint64
		users          map[uint64]*User
		symbol         map[string]*Symbol
		err            error
	)

	traderIds = make([]uint64, 0)
	traderNums = make([]uint64, 0)
	for _, vOrders := range req.SendBody.Orders {
		if 0 < vOrders.Uid {
			traderIds = append(traderIds, vOrders.Uid)
		}

		if 0 < vOrders.TraderNum {
			traderNums = append(traderNums, vOrders.TraderNum)
		}

	}

	// 新系统需要的交易员信息
	traders := make(map[string]*Trader, 0)
	tradersByIdMap := make(map[uint64]*Trader, 0)
	if 0 < len(traderIds) { // 先判断旧系统参数
		userBindTrader, err = b.binanceUserRepo.GetUserBindTraderTwoByTraderIds(traderIds)
		if nil != err {
			return
		}
	} else if 0 < len(traderNums) {
		traders, err = b.binanceUserRepo.GetTradersByTraderNum()
		if nil != err {
			return
		}

		for _, vTraderNum := range traderNums {
			tmpStr := strconv.FormatUint(vTraderNum, 10)
			if _, ok := traders[tmpStr]; ok {
				traderIds = append(traderIds, traders[tmpStr].ID) // 转化未旧系统映射的交易员id做后续处理
				tradersByIdMap[traders[tmpStr].ID] = traders[tmpStr]
			}
		}

		userBindTrader, err = b.binanceUserRepo.GetUserBindTraderTwoByTraderIds(traderIds)
		if nil != err {
			return
		}

		// 新系统下单请求
		newSystemReq = true
	} else {
		return
	}

	userIds = make([]uint64, 0)
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserBindTrader := range userBindTrader {
		for _, vVUserBindTrader := range vUserBindTrader {
			// 初始化仓位
			if 1 != vVUserBindTrader.InitOrder {
				continue
			}

			if _, ok := userIdsMap[vVUserBindTrader.UserId]; ok {
				continue
			}

			userIdsMap[vVUserBindTrader.UserId] = vVUserBindTrader.UserId
			userIds = append(userIds, vVUserBindTrader.UserId)
		}
	}

	if 0 >= len(userIds) {
		return
	}

	// 获取用户信息，余额信息，收益信息
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		return
	}

	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return
	}

	for _, vOrders := range req.SendBody.Orders {
		for _, vOrdersData := range vOrders.Data {
			var (
				orderUid          = vOrders.Uid
				tmpUserBindTrader []*UserBindTrader
			)

			// 新系统
			if newSystemReq {
				tmpStr := strconv.FormatUint(vOrders.TraderNum, 10)
				if _, ok := traders[tmpStr]; !ok {
					fmt.Println("新系统，未匹配到交易员", vOrders)
				}
				orderUid = traders[tmpStr].ID
			}

			if _, exists := userBindTrader[orderUid]; !exists {
				continue
			}

			tmpUserBindTrader = userBindTrader[orderUid]

			for _, vUserBindTrader := range tmpUserBindTrader {
				if 0 == vUserBindTrader.Status { // 绑定
					// 初始化仓位
					if 1 != vUserBindTrader.InitOrder {
						continue
					}

					if _, ok := users[vUserBindTrader.UserId]; !ok {
						continue
					}

					// todo 暂时使用检测没用api信息
					if 0 >= users[vUserBindTrader.UserId].ApiStatus {
						continue
					}

					// 精度
					if _, ok := symbol[vOrdersData.Symbol]; !ok {
						continue
					}

					// todo 暂时只关不开
					if ("SELL" == vOrdersData.Side && "SHORT" == vOrdersData.Type) || ("BUY" == vOrdersData.Side && "LONG" == vOrdersData.Type) {
						if 0 == users[vUserBindTrader.UserId].IsDai {
							if 5 != orderUid && 83 != orderUid {
								continue
							}
						}
					}

					var (
						tmpBaseMoney string
					)
					if newSystemReq {
						// 新系统请求
						if 1 == users[vUserBindTrader.UserId].UseNewSystem { // 新系统用户
							if _, ok := tradersByIdMap[vUserBindTrader.TraderId]; !ok {
								fmt.Println("新系统下单错误，未找到交易员", vUserBindTrader.TraderId)
								continue
							}

							tmpBaseMoney = strconv.FormatFloat(tradersByIdMap[vUserBindTrader.TraderId].BaseMoney, 'f', -1, 64)
						} else {
							continue // 不是新系统用户，跳过
						}
					} else {
						// 老系统请求
						if 1 == users[vUserBindTrader.UserId].UseNewSystem { // 新系统用户
							continue
						} else {
							tmpBaseMoney = vOrders.BaseMoney
						}
					}

					//
					//fmt.Println(tmpBaseMoney, users[vUserBindTrader.UserId], vOrdersData)
					// 发送订单
					wg.Add(1) // 启动一个goroutine就登记+1
					go b.userOrderGoroutineTwo(ctx, wg, &OrderData{
						Coin:     vOrdersData.Symbol,
						Type:     vOrdersData.Type,
						Price:    vOrdersData.Price,
						Side:     vOrdersData.Side,
						Qty:      vOrdersData.Qty,
						Position: vOrdersData.Position,
					}, tmpBaseMoney, users[vUserBindTrader.UserId], vUserBindTrader, symbol[vOrdersData.Symbol].QuantityPrecision, 0, 0, 0)
				}
			}
		}
	}
}

func (b *BinanceUserUsecase) userOrderGoroutineTwo(ctx context.Context, wg *sync.WaitGroup, order *OrderData, amount string, user *User, userBindTrader *UserBindTrader, quantityPrecision int64, initOrderReq uint64, proportion float64, overOrderReq uint64) {
	defer wg.Done() // goroutine结束就登记-1

	var (
		binanceOrder  *BinanceOrder
		orderType     = "MARKET"
		positionSide  string
		quantity      string
		qty           float64
		price         float64
		traderAmount  float64
		currentOrders []*UserOrder
		currentOrder  *UserOrder
		insertOrder   *UserOrder
		err           error
	)

	// 新订单数据
	currentOrder = &UserOrder{
		UserId:        userBindTrader.UserId,
		TraderId:      userBindTrader.TraderId,
		Symbol:        order.Coin,
		Side:          "",
		PositionSide:  "",
		Quantity:      0,
		Price:         0,
		TraderQty:     0,
		OrderType:     orderType,
		ClosePosition: "",
		CumQuote:      0,
		ExecutedQty:   0,
		AvgPrice:      0,
	}

	if 0 == overOrderReq { // 非全平仓下单
		qty, err = strconv.ParseFloat(order.Qty, 64)
		if nil != err {
			fmt.Println(err, order, userBindTrader)
			return
		}
		currentOrder.TraderQty = qty

		traderAmount, err = strconv.ParseFloat(amount, 64)
		if nil != err {
			fmt.Println(err, order, userBindTrader)
			return
		}
	}

	if "LONG" == order.Type {
		positionSide = "LONG"
	} else if "SHORT" == order.Type {
		positionSide = "SHORT" // 空
	} else {
		fmt.Println("err position side")
		return
	}
	currentOrder.PositionSide = positionSide

	var (
		quantityFloat float64
	)
	if 1 == initOrderReq { // 初始化下单
		price, err = strconv.ParseFloat(order.Price, 64)
		if nil != err {
			fmt.Println(err, order, userBindTrader)
			return
		}
		currentOrder.Price = price

		if 0 > currentOrder.Price {
			fmt.Println("err price", currentOrder.Price, order, userBindTrader)
			return
		}

		quantityFloat = float64(userBindTrader.Amount) * proportion / currentOrder.Price
	} else if 1 == overOrderReq { // 全平仓

	} else {
		quantityFloat = float64(userBindTrader.Amount) * qty / traderAmount // 本次开单数量
	}

	var historyQuantityFloat float64
	// 本次平仓
	if ("SELL" == order.Side && "LONG" == order.Type) || ("BUY" == order.Side && "SHORT" == order.Type) {
		// 订单统计
		currentOrders, err = b.binanceUserRepo.GetUserOrderTwoByUserTraderIdAndSymbol(userBindTrader.UserId, userBindTrader.TraderId, order.Coin)
		if nil != err {
			fmt.Println(err, order)
			return
		}
		// 查出用户的BUY单的币的数量，在对应的trader下，超过了BUY不能SELL todo 使用数据库量太大以后
		if 0 >= len(currentOrders) {
			return
		}

		// 多的部分不管，按剩余的数量关 todo 交易员全部平仓，少的部分另一个程序解决
		for _, vCurrentOrders := range currentOrders {
			// 本次平多
			if "SELL" == order.Side && "LONG" == order.Type {
				// 历史开多和平多
				if "BUY" == vCurrentOrders.Side && "LONG" == vCurrentOrders.PositionSide {
					historyQuantityFloat += vCurrentOrders.ExecutedQty
				} else if "SELL" == vCurrentOrders.Side && "LONG" == vCurrentOrders.PositionSide {
					historyQuantityFloat -= vCurrentOrders.ExecutedQty
				}
			} else if "BUY" == order.Side && "SHORT" == order.Type {
				// 历史开空和平空
				if "SELL" == vCurrentOrders.Side && "SHORT" == vCurrentOrders.PositionSide {
					historyQuantityFloat += vCurrentOrders.ExecutedQty
				} else if "BUY" == vCurrentOrders.Side && "SHORT" == vCurrentOrders.PositionSide {
					historyQuantityFloat -= vCurrentOrders.ExecutedQty
				}
			}
		}

		// 开单历史数量不足了
		if isZero(historyQuantityFloat) || 0 > historyQuantityFloat {
			fmt.Println("trader的开单数量小于关单数量了，可能是精度问题", order.Coin, userBindTrader.UserId, userBindTrader.TraderId, historyQuantityFloat)
			if 1 == overOrderReq {
				fmt.Println("全平功能")
			}
			return
		}

		if 1 == overOrderReq { // 全平仓
			quantityFloat = historyQuantityFloat
		} else {
			// 按仓位百分比平仓
			if 0 < len(order.Position) {
				var (
					tmpPositionTotal float64
				)

				tmpPositionTotal, err = strconv.ParseFloat(order.Position, 64)
				if nil != err {
					fmt.Println(err, order, userBindTrader)
					return
				}

				if 0 < tmpPositionTotal && 0 < qty {
					quantityFloat = historyQuantityFloat * qty / tmpPositionTotal // 本次平仓数量
				}
			}

			// 超过了净开单数量
			// todo 并发操作时数据不一致的可能性，会导致数量对不上，例如下单程序和更换绑定程序同时执行时，是否程序中统计的总数字会漏计算这里的新增的订单的数字
			if quantityFloat > historyQuantityFloat {
				quantityFloat = historyQuantityFloat
			}
		}

	} else if ("SELL" == order.Side && "SHORT" == order.Type) || ("BUY" == order.Side && "LONG" == order.Type) {
		// 开仓

	} else {
		fmt.Println("err order side")
		return
	}

	currentOrder.Side = order.Side

	// 精度调整
	if 0 >= quantityPrecision {
		quantity = fmt.Sprintf("%d", int64(quantityFloat))
	} else {
		quantity = strconv.FormatFloat(quantityFloat, 'f', int(quantityPrecision), 64)
	}

	currentOrder.Quantity, err = strconv.ParseFloat(quantity, 64)
	if nil != err {
		fmt.Println(err)
		return
	}

	if 0 >= currentOrder.Quantity {
		fmt.Println("下单数值太小", quantity, currentOrder, quantityPrecision)
		return
	}

	var orderInfo *OrderInfo
	// 请求下单
	binanceOrder, orderInfo, err = requestBinanceOrder(order.Coin, order.Side, orderType, positionSide, quantity, user.ApiKey, user.ApiSecret)
	if nil != err {
		fmt.Println(err)
		return
	}

	// 下单异常
	if 0 >= binanceOrder.OrderId {
		// 写入
		orderErr := &UserOrderErr{
			UserId:        currentOrder.UserId,
			TraderId:      currentOrder.TraderId,
			ClientOrderId: currentOrder.ClientOrderId,
			OrderId:       currentOrder.OrderId,
			Symbol:        currentOrder.Symbol,
			Side:          currentOrder.Side,
			PositionSide:  currentOrder.PositionSide,
			Quantity:      quantityFloat,
			Price:         currentOrder.Price,
			TraderQty:     currentOrder.TraderQty,
			OrderType:     currentOrder.OrderType,
			ClosePosition: currentOrder.ClosePosition,
			CumQuote:      currentOrder.CumQuote,
			ExecutedQty:   currentOrder.ExecutedQty,
			AvgPrice:      currentOrder.AvgPrice,
			HandleStatus:  currentOrder.HandleStatus,
			InitOrder:     int64(initOrderReq),
			Code:          orderInfo.Code,
			Msg:           orderInfo.Msg,
			Proportion:    proportion,
		}

		if 1 == overOrderReq {
			orderErr.InitOrder = 2
		}

		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.InsertUserOrderErrTwo(ctx, orderErr)
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			fmt.Println(err, orderErr)
			return
		}

		return // 返回
	}

	currentOrder.OrderId = strconv.FormatInt(binanceOrder.OrderId, 10)

	currentOrder.CumQuote, err = strconv.ParseFloat(binanceOrder.CumQuote, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	currentOrder.ExecutedQty, err = strconv.ParseFloat(binanceOrder.ExecutedQty, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	currentOrder.AvgPrice, err = strconv.ParseFloat(binanceOrder.AvgPrice, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	// 写入
	if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
		insertOrder, err = b.binanceUserRepo.InsertUserOrderTwo(ctx, currentOrder)
		if nil != err {
			return err
		}

		return nil
	}); err != nil {
		fmt.Println(err, insertOrder)
		return
	}

	// 计算收益 卖单去binance查询本单收益
	if ("SELL" == order.Side && "LONG" == order.Type) || ("BUY" == order.Side && "SHORT" == order.Type) {
		if 0 >= insertOrder.ID {
			fmt.Println("错误的数据写入", insertOrder)
			return
		}

		err = b.binanceUserRepo.SAddOrderSetSellLongOrBuyShortTwo(ctx, int64(insertOrder.ID))
		if nil != err {
			fmt.Println("错误的数据写入, redis", insertOrder, err)
			return
		}
	}

	return
}

// OrderHandle 收益处理
func (b *BinanceUserUsecase) OrderHandle(ctx context.Context, req *v1.OrderHandleRequest) (*v1.OrderHandleReply, error) {
	var (
		wg          sync.WaitGroup
		orderIdsStr []string
		orderIds    []int64
		userOrders  []*UserOrder
		userIdsMap  map[uint64]uint64
		userIds     []uint64
		users       map[uint64]*User
		err         error
	)

	orderIdsStr, err = b.binanceUserRepo.SMembersOrderSetSellLongOrBuyShort(ctx)
	if nil != err {
		fmt.Println(err)
	}

	if 0 >= len(orderIdsStr) {
		return &v1.OrderHandleReply{}, nil
	}

	// 订单信息
	orderIds = make([]int64, 0)
	for _, vOrderIdsStr := range orderIdsStr {
		var tmp int64
		tmp, err = strconv.ParseInt(vOrderIdsStr, 10, 64)
		if nil != err {
			fmt.Println(err)
		}

		orderIds = append(orderIds, tmp)
	}
	userOrders, err = b.binanceUserRepo.GetUserOrderByIds(orderIds)
	if nil != err {
		fmt.Println("更新平单后收益数据失败, 空查询, 订单不存在", err, orderIds)
		return &v1.OrderHandleReply{}, nil
	}

	// 用户信息
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserOrders := range userOrders {
		userIdsMap[vUserOrders.UserId] = vUserOrders.UserId
	}
	userIds = make([]uint64, 0)
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		fmt.Println("更新平单后收益数据失败, 空查询，用户查询不存在", err, userIds)
		return &v1.OrderHandleReply{}, nil
	}

	for k, vUserOrders := range userOrders {
		if 100 <= k {
			break
		}

		if _, ok := users[vUserOrders.UserId]; !ok {
			fmt.Println("更新平单收益，不存在用户", vUserOrders)
			continue
		}

		wg.Add(1) // 启动一个goroutine就登记+1
		go b.userOrderHandleGoroutine(ctx, &wg, vUserOrders, users[vUserOrders.UserId])
	}

	wg.Wait() // 等待所有登记的goroutine都结束

	return &v1.OrderHandleReply{}, nil
}

func (b *BinanceUserUsecase) userOrderHandleGoroutine(ctx context.Context, wg *sync.WaitGroup, userOrder *UserOrder, user *User) {
	defer wg.Done() // goroutine结束就登记-1

	var (
		binanceOrder []*OrderHistory
		income       float64
		orderIdInt   int64
		err          error
	)

	// 平
	if !(("SELL" == userOrder.Side && "LONG" == userOrder.PositionSide) || ("BUY" == userOrder.Side && "SHORT" == userOrder.PositionSide)) {
		fmt.Println("非平仓", userOrder.ID)
		return
	}

	if 0 < userOrder.HandleStatus {
		//继续执行删除
	} else {

		// 查询
		orderIdInt, err = strconv.ParseInt(userOrder.OrderId, 10, 64)

		binanceOrder, err = requestBinanceOrderHistory(user.ApiKey, user.ApiSecret, userOrder.Symbol, orderIdInt, "", "")
		if nil != err {
			fmt.Println("更新平单后收益数据失败, 错误查询", err, userOrder.ID)
			return
		}

		if 0 >= len(binanceOrder) {
			fmt.Println("更新平单后收益数据失败, 空数据", binanceOrder, userOrder.ID)
			return
		}

		// 收益
		for _, vBinanceOrder := range binanceOrder {
			var (
				tmpIncome float64
			)
			tmpIncome, err = strconv.ParseFloat(vBinanceOrder.RealizedPnl, 64)
			if nil != err {
				fmt.Println("更新平单后收益数据失败", err, vBinanceOrder)
				return
			}

			income += tmpIncome
		}

		// 写入
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.UpdatesUserOrderHandleStatus(ctx, userOrder.ID)
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.UpdatesUserAmount(ctx, userOrder.UserId, int64(income*100000))
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserAmountRecord(ctx, &UserAmountRecord{
				UserId:    userOrder.UserId,
				OrderId:   userOrder.ID,
				Amount:    int64(income * 100000),
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			})
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			fmt.Println("更新平单后收益数据失败", err, userOrder.ID)
			return
		}
	}

	err = b.binanceUserRepo.SRemOrderSetSellLongOrBuyShort(ctx, int64(userOrder.ID))
	if nil != err {
		fmt.Println("更新平单后收益数据失败, 错误的数据删除, redis，order_id:", int64(userOrder.ID), userOrder.ID, err)
		return
	}

	return
}

// OrderHandleTwo 收益处理
func (b *BinanceUserUsecase) OrderHandleTwo(ctx context.Context, req *v1.OrderHandleRequest) (*v1.OrderHandleReply, error) {
	var (
		wg          sync.WaitGroup
		orderIdsStr []string
		orderIds    []int64
		userOrders  []*UserOrder
		userIdsMap  map[uint64]uint64
		userIds     []uint64
		users       map[uint64]*User
		err         error
	)

	orderIdsStr, err = b.binanceUserRepo.SMembersOrderSetSellLongOrBuyShortTwo(ctx)
	if nil != err {
		fmt.Println(err)
	}

	if 0 >= len(orderIdsStr) {
		return &v1.OrderHandleReply{}, nil
	}

	// 订单信息
	orderIds = make([]int64, 0)
	for _, vOrderIdsStr := range orderIdsStr {
		var tmp int64
		tmp, err = strconv.ParseInt(vOrderIdsStr, 10, 64)
		if nil != err {
			fmt.Println(err)
		}

		orderIds = append(orderIds, tmp)
	}
	userOrders, err = b.binanceUserRepo.GetUserOrderTwoByIds(orderIds)
	if nil != err {
		fmt.Println("更新平单后收益数据失败, 空查询, 订单不存在", err, orderIds)
		return &v1.OrderHandleReply{}, nil
	}

	// 用户信息
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserOrders := range userOrders {
		userIdsMap[vUserOrders.UserId] = vUserOrders.UserId
	}
	userIds = make([]uint64, 0)
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		fmt.Println("更新平单后收益数据失败, 空查询，用户查询不存在", err, userIds)
		return &v1.OrderHandleReply{}, nil
	}

	for k, vUserOrders := range userOrders {
		if 100 <= k {
			break
		}

		if _, ok := users[vUserOrders.UserId]; !ok {
			fmt.Println("更新平单收益，不存在用户", vUserOrders)
			continue
		}

		wg.Add(1) // 启动一个goroutine就登记+1
		go b.userOrderHandleGoroutineTwo(ctx, &wg, vUserOrders, users[vUserOrders.UserId])
	}

	wg.Wait() // 等待所有登记的goroutine都结束

	return &v1.OrderHandleReply{}, nil
}

func (b *BinanceUserUsecase) userOrderHandleGoroutineTwo(ctx context.Context, wg *sync.WaitGroup, userOrder *UserOrder, user *User) {
	defer wg.Done() // goroutine结束就登记-1

	var (
		binanceOrder []*OrderHistory
		income       float64
		orderIdInt   int64
		err          error
	)

	// 平
	if !(("SELL" == userOrder.Side && "LONG" == userOrder.PositionSide) || ("BUY" == userOrder.Side && "SHORT" == userOrder.PositionSide)) {
		fmt.Println("非平仓", userOrder.ID)
		return
	}

	if 0 < userOrder.HandleStatus {
		//继续执行删除
	} else {

		// 查询
		orderIdInt, err = strconv.ParseInt(userOrder.OrderId, 10, 64)

		binanceOrder, err = requestBinanceOrderHistory(user.ApiKey, user.ApiSecret, userOrder.Symbol, orderIdInt, "", "")
		if nil != err {
			fmt.Println("更新平单后收益数据失败, 错误查询", err, userOrder.ID)
			return
		}

		if 0 >= len(binanceOrder) {
			fmt.Println("更新平单后收益数据失败, 空数据", binanceOrder, userOrder.ID)
			return
		}

		// 收益
		for _, vBinanceOrder := range binanceOrder {
			var (
				tmpIncome float64
			)
			tmpIncome, err = strconv.ParseFloat(vBinanceOrder.RealizedPnl, 64)
			if nil != err {
				fmt.Println("更新平单后收益数据失败", err, vBinanceOrder)
				return
			}

			income += tmpIncome
		}

		// 写入
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.UpdatesUserOrderTwoHandleStatus(ctx, userOrder.ID)
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.UpdatesUserAmountTwo(ctx, userOrder.UserId, int64(income*100000))
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserAmountRecordTwo(ctx, &UserAmountRecord{
				UserId:    userOrder.UserId,
				OrderId:   userOrder.ID,
				Amount:    int64(income * 100000),
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			})
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			fmt.Println("更新平单后收益数据失败", err, userOrder.ID)
			return
		}
	}

	err = b.binanceUserRepo.SRemOrderSetSellLongOrBuyShortTwo(ctx, int64(userOrder.ID))
	if nil != err {
		fmt.Println("更新平单后收益数据失败, 错误的数据删除, redis，order_id:", int64(userOrder.ID), userOrder.ID, err)
		return
	}

	return
}

// CloseOrderAfterBind 强制关单
func (b *BinanceUserUsecase) CloseOrderAfterBind(ctx context.Context, req *v1.CloseOrderAfterBindRequest) (*v1.CloseOrderAfterBindReply, error) {
	var (
		wg                  sync.WaitGroup
		userBindAfterUnbind []*UserBindAfterUnbind
		userIdsMap          map[uint64]uint64
		userIds             []uint64
		users               map[uint64]*User
		symbol              map[string]*Symbol
		err                 error
	)

	userBindAfterUnbind, err = b.binanceUserRepo.GetUserBindAfterUnbindByStatus()
	if nil != err {
		return nil, err
	}

	// 用户信息
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserBindAfterUnbind := range userBindAfterUnbind {
		userIdsMap[vUserBindAfterUnbind.UserId] = vUserBindAfterUnbind.UserId
	}
	userIds = make([]uint64, 0)
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		fmt.Println("绑定后的关单，空查询，用户查询不存在", err, userIds)
		return nil, nil
	}
	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		fmt.Println("绑定后的关单，空查询，代币查询不存在", err, userIds)
		return nil, nil
	}

	for _, vUserBIndAfterUnbind := range userBindAfterUnbind {
		// 精度
		if _, ok := symbol[vUserBIndAfterUnbind.Symbol]; !ok {
			continue
		}

		if _, ok := users[vUserBIndAfterUnbind.UserId]; !ok {
			continue
		}

		// todo 暂时使用检测没用api信息
		if 0 >= users[vUserBIndAfterUnbind.UserId].ApiStatus {
			continue
		}

		wg.Add(1) // 启动一个goroutine就登记+1
		go b.CloseOrderAfterBindGoroutine(ctx, &wg, vUserBIndAfterUnbind, users[vUserBIndAfterUnbind.UserId], symbol[vUserBIndAfterUnbind.Symbol].QuantityPrecision)
	}

	wg.Wait() // 等待所有登记的goroutine都结束

	return nil, nil
}

func (b *BinanceUserUsecase) CloseOrderAfterBindGoroutine(ctx context.Context, wg *sync.WaitGroup, userBindAfterUnbind *UserBindAfterUnbind, user *User, quantityPrecision int64) {
	defer wg.Done() // goroutine结束就登记-1

	var (
		currentOrder *UserOrder
		insertOrder  *UserOrder
		binanceOrder *BinanceOrder
		orderType    = "MARKET"
		err          error
	)

	// 新订单数据
	currentOrder = &UserOrder{
		UserId:        userBindAfterUnbind.UserId,
		TraderId:      userBindAfterUnbind.TraderId,
		Symbol:        userBindAfterUnbind.Symbol,
		Side:          "",
		PositionSide:  "",
		Quantity:      0,
		Price:         0,
		TraderQty:     0,
		OrderType:     orderType,
		ClosePosition: "",
		CumQuote:      0,
		ExecutedQty:   0,
		AvgPrice:      0,
	}

	currentOrder.TraderQty = userBindAfterUnbind.Quantity
	if "LONG" == userBindAfterUnbind.PositionSide {
		currentOrder.Side = "SELL" //平多
	} else if "SHORT" == userBindAfterUnbind.PositionSide {
		currentOrder.Side = "BUY"
	} else {
		fmt.Println("err position side")
		return
	}
	currentOrder.PositionSide = userBindAfterUnbind.PositionSide

	var (
		quantity string
	)

	quantityFloat := userBindAfterUnbind.Quantity
	// 精度调整
	if 0 >= quantityPrecision {
		quantity = fmt.Sprintf("%d", int64(quantityFloat))
	} else {
		quantity = strconv.FormatFloat(quantityFloat, 'f', int(quantityPrecision), 64)
	}
	currentOrder.Quantity, err = strconv.ParseFloat(quantity, 64)
	if nil != err {
		fmt.Println(err)
		return
	}

	if 0 >= currentOrder.Quantity {
		fmt.Println("下单数值太小", quantity, currentOrder, quantityPrecision)
		return
	}

	var orderInfo *OrderInfo
	// 请求下单
	binanceOrder, orderInfo, err = requestBinanceOrder(userBindAfterUnbind.Symbol, currentOrder.Side, orderType, currentOrder.PositionSide, quantity, user.ApiKey, user.ApiSecret)
	if nil != err {
		fmt.Println(err)
		return
	}

	// 下单异常
	if 0 >= binanceOrder.OrderId {
		// 写入
		orderErr := &UserOrderErr{
			UserId:        currentOrder.UserId,
			TraderId:      currentOrder.TraderId,
			ClientOrderId: currentOrder.ClientOrderId,
			OrderId:       currentOrder.OrderId,
			Symbol:        currentOrder.Symbol,
			Side:          currentOrder.Side,
			PositionSide:  currentOrder.PositionSide,
			Quantity:      quantityFloat,
			Price:         currentOrder.Price,
			TraderQty:     currentOrder.TraderQty,
			OrderType:     currentOrder.OrderType,
			ClosePosition: currentOrder.ClosePosition,
			CumQuote:      currentOrder.CumQuote,
			ExecutedQty:   currentOrder.ExecutedQty,
			AvgPrice:      currentOrder.AvgPrice,
			HandleStatus:  currentOrder.HandleStatus,
			Code:          orderInfo.Code,
			Msg:           orderInfo.Msg,
		}

		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.InsertUserOrderErr(ctx, orderErr)
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			fmt.Println(err, orderErr)
			return
		}

		return // 返回
	}

	currentOrder.OrderId = strconv.FormatInt(binanceOrder.OrderId, 10)

	currentOrder.CumQuote, err = strconv.ParseFloat(binanceOrder.CumQuote, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	currentOrder.ExecutedQty, err = strconv.ParseFloat(binanceOrder.ExecutedQty, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	currentOrder.AvgPrice, err = strconv.ParseFloat(binanceOrder.AvgPrice, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	// 写入
	if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
		insertOrder, err = b.binanceUserRepo.InsertUserOrder(ctx, currentOrder)
		if nil != err {
			return err
		}

		return nil
	}); err != nil {
		fmt.Println(err, currentOrder)
		return
	}

	// 计算收益 卖单去binance查询本单收益
	if ("SELL" == insertOrder.Side && "LONG" == insertOrder.PositionSide) || ("BUY" == insertOrder.Side && "SHORT" == insertOrder.PositionSide) {
		if 0 >= insertOrder.ID {
			fmt.Println("错误的数据写入", insertOrder)
			return
		}

		err = b.binanceUserRepo.SAddOrderSetSellLongOrBuyShort(ctx, int64(insertOrder.ID))
		if nil != err {
			fmt.Println("错误的数据写入, redis", insertOrder, err)
			return
		}
	}

	// 新事务 写入防止锁等待可能影响订单的insert
	if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
		if 0 < insertOrder.ID {
			_, err = b.binanceUserRepo.UpdatesUserBindAfterUnbind(ctx, userBindAfterUnbind.ID, 1, userBindAfterUnbind.Quantity-currentOrder.ExecutedQty)
			if nil != err {
				return err
			}
		}

		return nil
	}); err != nil {
		fmt.Println(err, currentOrder)
		return
	}

	return
}

// CloseOrderAfterBindTwo 强制关单tfi
func (b *BinanceUserUsecase) CloseOrderAfterBindTwo(ctx context.Context, req *v1.CloseOrderAfterBindRequest) (*v1.CloseOrderAfterBindReply, error) {
	var (
		wg                  sync.WaitGroup
		userBindAfterUnbind []*UserBindAfterUnbind
		userIdsMap          map[uint64]uint64
		userIds             []uint64
		users               map[uint64]*User
		symbol              map[string]*Symbol
		err                 error
	)

	userBindAfterUnbind, err = b.binanceUserRepo.GetUserBindAfterUnbindTwoByStatus()
	if nil != err {
		return nil, err
	}

	// 用户信息
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserBindAfterUnbind := range userBindAfterUnbind {
		userIdsMap[vUserBindAfterUnbind.UserId] = vUserBindAfterUnbind.UserId
	}
	userIds = make([]uint64, 0)
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		fmt.Println("绑定后的关单，空查询，用户查询不存在", err, userIds)
		return nil, nil
	}
	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		fmt.Println("绑定后的关单，空查询，代币查询不存在", err, userIds)
		return nil, nil
	}

	for _, vUserBIndAfterUnbind := range userBindAfterUnbind {
		// 精度
		if _, ok := symbol[vUserBIndAfterUnbind.Symbol]; !ok {
			continue
		}

		if _, ok := users[vUserBIndAfterUnbind.UserId]; !ok {
			continue
		}

		// todo 暂时使用检测没用api信息
		if 0 >= users[vUserBIndAfterUnbind.UserId].ApiStatus {
			continue
		}

		wg.Add(1) // 启动一个goroutine就登记+1
		go b.CloseOrderAfterBindGoroutineTwo(ctx, &wg, vUserBIndAfterUnbind, users[vUserBIndAfterUnbind.UserId], symbol[vUserBIndAfterUnbind.Symbol].QuantityPrecision)
	}

	wg.Wait() // 等待所有登记的goroutine都结束

	return nil, nil
}

func (b *BinanceUserUsecase) CloseOrderAfterBindGoroutineTwo(ctx context.Context, wg *sync.WaitGroup, userBindAfterUnbind *UserBindAfterUnbind, user *User, quantityPrecision int64) {
	defer wg.Done() // goroutine结束就登记-1

	var (
		currentOrder *UserOrder
		insertOrder  *UserOrder
		binanceOrder *BinanceOrder
		orderType    = "MARKET"
		err          error
	)

	// 新订单数据
	currentOrder = &UserOrder{
		UserId:        userBindAfterUnbind.UserId,
		TraderId:      userBindAfterUnbind.TraderId,
		Symbol:        userBindAfterUnbind.Symbol,
		Side:          "",
		PositionSide:  "",
		Quantity:      0,
		Price:         0,
		TraderQty:     0,
		OrderType:     orderType,
		ClosePosition: "",
		CumQuote:      0,
		ExecutedQty:   0,
		AvgPrice:      0,
	}

	currentOrder.TraderQty = userBindAfterUnbind.Quantity
	if "LONG" == userBindAfterUnbind.PositionSide {
		currentOrder.Side = "SELL" //平多
	} else if "SHORT" == userBindAfterUnbind.PositionSide {
		currentOrder.Side = "BUY"
	} else {
		fmt.Println("err position side")
		return
	}
	currentOrder.PositionSide = userBindAfterUnbind.PositionSide

	var (
		quantity string
	)

	quantityFloat := userBindAfterUnbind.Quantity
	// 精度调整
	if 0 >= quantityPrecision {
		quantity = fmt.Sprintf("%d", int64(quantityFloat))
	} else {
		quantity = strconv.FormatFloat(quantityFloat, 'f', int(quantityPrecision), 64)
	}
	currentOrder.Quantity, err = strconv.ParseFloat(quantity, 64)
	if nil != err {
		fmt.Println(err)
		return
	}

	if 0 >= currentOrder.Quantity {
		fmt.Println("下单数值太小", quantity, currentOrder, quantityPrecision)
		return
	}

	var orderInfo *OrderInfo
	// 请求下单
	binanceOrder, orderInfo, err = requestBinanceOrder(userBindAfterUnbind.Symbol, currentOrder.Side, orderType, currentOrder.PositionSide, quantity, user.ApiKey, user.ApiSecret)
	if nil != err {
		fmt.Println(err)
		return
	}

	// 下单异常
	if 0 >= binanceOrder.OrderId {
		// 写入
		orderErr := &UserOrderErr{
			UserId:        currentOrder.UserId,
			TraderId:      currentOrder.TraderId,
			ClientOrderId: currentOrder.ClientOrderId,
			OrderId:       currentOrder.OrderId,
			Symbol:        currentOrder.Symbol,
			Side:          currentOrder.Side,
			PositionSide:  currentOrder.PositionSide,
			Quantity:      quantityFloat,
			Price:         currentOrder.Price,
			TraderQty:     currentOrder.TraderQty,
			OrderType:     currentOrder.OrderType,
			ClosePosition: currentOrder.ClosePosition,
			CumQuote:      currentOrder.CumQuote,
			ExecutedQty:   currentOrder.ExecutedQty,
			AvgPrice:      currentOrder.AvgPrice,
			HandleStatus:  currentOrder.HandleStatus,
			Code:          orderInfo.Code,
			Msg:           orderInfo.Msg,
		}

		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.InsertUserOrderErrTwo(ctx, orderErr)
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			fmt.Println(err, orderErr)
			return
		}

		return // 返回
	}

	currentOrder.OrderId = strconv.FormatInt(binanceOrder.OrderId, 10)

	currentOrder.CumQuote, err = strconv.ParseFloat(binanceOrder.CumQuote, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	currentOrder.ExecutedQty, err = strconv.ParseFloat(binanceOrder.ExecutedQty, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	currentOrder.AvgPrice, err = strconv.ParseFloat(binanceOrder.AvgPrice, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	// 写入
	if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
		insertOrder, err = b.binanceUserRepo.InsertUserOrderTwo(ctx, currentOrder)
		if nil != err {
			return err
		}

		return nil
	}); err != nil {
		fmt.Println("解绑后平仓，", err, currentOrder)
		return
	}

	// 计算收益 卖单去binance查询本单收益
	if ("SELL" == insertOrder.Side && "LONG" == insertOrder.PositionSide) || ("BUY" == insertOrder.Side && "SHORT" == insertOrder.PositionSide) {
		if 0 >= insertOrder.ID {
			fmt.Println("解绑后平仓，错误的数据写入", insertOrder)
			return
		}

		err = b.binanceUserRepo.SAddOrderSetSellLongOrBuyShortTwo(ctx, int64(insertOrder.ID))
		if nil != err {
			fmt.Println("解绑后平仓，错误的数据写入, redis", insertOrder, err)
			return
		}
	}

	// 新事务 写入防止锁等待可能影响订单的insert
	if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
		if 0 < insertOrder.ID {
			_, err = b.binanceUserRepo.UpdatesUserBindAfterUnbindTwo(ctx, userBindAfterUnbind.ID, 1, userBindAfterUnbind.Quantity-currentOrder.ExecutedQty)
			if nil != err {
				return err
			}
		}

		return nil
	}); err != nil {
		fmt.Println("解绑后平仓，", err, currentOrder)
		return
	}

	return
}

// InitOrderAfterBind 初始化仓位
func (b *BinanceUserUsecase) InitOrderAfterBind(ctx context.Context, req *v1.InitOrderAfterBindRequest) (*v1.InitOrderAfterBindReply, error) {
	var (
		wg              sync.WaitGroup
		userBindTraders map[uint64][]*UserBindTrader
		traders         map[uint64]*Trader
		userIds         []uint64
		userIdsMap      map[uint64]uint64
		users           map[uint64]*User
		userBalance     map[uint64]*UserBalance
		userAmount      map[uint64]*UserAmount
		symbol          map[string]*Symbol
		err             error
	)

	userBindTraders, err = b.binanceUserRepo.GetUserBindTraderByInitOrder()
	if nil != err {
		return nil, err
	}

	userIds = make([]uint64, 0)
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserBindTrader := range userBindTraders {
		for _, vVUserBindTrader := range vUserBindTrader {
			if _, ok := userIdsMap[vVUserBindTrader.UserId]; ok {
				continue
			}

			userIdsMap[vVUserBindTrader.UserId] = vVUserBindTrader.UserId
			userIds = append(userIds, vVUserBindTrader.UserId)
		}
	}

	if 0 >= len(userIds) {
		return nil, nil
	}

	traders, err = b.binanceUserRepo.GetTraders()
	if nil != err {
		return nil, nil
	}

	// 获取用户信息，余额信息，收益信息
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		return nil, nil
	}

	userBalance, err = b.binanceUserRepo.GetUserBalanceByUserIds(ctx, userIds)
	if nil != err {
		return nil, nil
	}

	userAmount, err = b.binanceUserRepo.GetUserAmountByUserIds(ctx, userIds)
	if nil != err {
		return nil, nil
	}

	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return nil, nil
	}

	for traderId, vUserBindTraders := range userBindTraders {
		var (
			traderPositions []*TraderPosition
		)

		// 先更新状态
		for _, vVUserBindTraders := range vUserBindTraders {
			if _, ok := users[vVUserBindTraders.UserId]; !ok {
				continue
			}

			// todo 暂时使用检测没用api信息
			if 0 >= users[vVUserBindTraders.UserId].ApiStatus {
				continue
			}

			if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = b.binanceUserRepo.UpdatesUserBindTraderInitOrderById(ctx, vVUserBindTraders.ID)
				if nil != err {
					return err
				}

				return nil
			}); err != nil {
				fmt.Println(err, "修改初始化失败", traderId, vVUserBindTraders)
				continue
			}
		}

		if _, ok := traders[traderId]; !ok {
			fmt.Println("初始化仓位，不存在交易员", traderId)
			continue
		}

		traderPositions, err = b.binanceUserRepo.GetTraderPosition(traderId)
		if nil != err {
			fmt.Println("初始化仓位，空仓位", traderId, err)
			continue
		}

		// 按仓位
		for _, vTraderPositions := range traderPositions {
			var (
				price      *BinanceSymbolPrice
				priceFloat float64
			)
			// 查询币价
			price, err = requestBinanceSymbolPrice(vTraderPositions.Symbol)
			if nil != err {
				fmt.Println("初始化仓位，价格查询", traderId, price, err)
				continue
			}

			priceFloat, err = strconv.ParseFloat(price.Price, 64)
			if nil != err {
				fmt.Println("初始化仓位，价格查询，转化", traderId, price, err)
				continue
			}

			// 交易员仓位u占保证金比例
			var proportion float64
			if 0 >= vTraderPositions.Qty || 0 >= traders[traderId].BaseMoney {
				continue
			}

			proportion = vTraderPositions.Qty * priceFloat / traders[traderId].BaseMoney
			if 0 >= proportion || 0 >= priceFloat {
				fmt.Println("初始化仓位，比例计算", vTraderPositions, proportion, vTraderPositions.Qty, priceFloat, traders[traderId].BaseMoney)
				continue
			}

			for _, vVUserBindTraders := range vUserBindTraders {
				if 0 == vVUserBindTraders.Status { // 绑定
					if _, ok := users[vVUserBindTraders.UserId]; !ok {
						continue
					}

					// todo 暂时使用检测没用api信息
					if 0 >= users[vVUserBindTraders.UserId].ApiStatus {
						continue
					}

					if _, ok := userBalance[vVUserBindTraders.UserId]; !ok {
						continue
					}

					if _, ok := userAmount[vVUserBindTraders.UserId]; !ok {
						continue
					}

					// 判断是开单还是关单，sell long 关多 buy short 关空
					if ("SELL" == vTraderPositions.Side && "SHORT" == vTraderPositions.PositionSide) || ("BUY" == vTraderPositions.Side && "LONG" == vTraderPositions.PositionSide) {
						// todo 暂时只关不开
						if 0 == users[vVUserBindTraders.UserId].IsDai {
							if 5 != vVUserBindTraders.TraderId && 83 != vVUserBindTraders.TraderId {
								continue
							}
						}

						// 精度按代币18位，截取小数点后到5位计算
						var balanceTmp int64
						lengthToKeep := len(userBalance[vVUserBindTraders.UserId].Balance) - 13

						if lengthToKeep > 0 {
							balanceTmpStr := userBalance[vVUserBindTraders.UserId].Balance[:lengthToKeep]
							balanceTmp, err = strconv.ParseInt(balanceTmpStr, 10, 64)
							if nil != err || 0 >= balanceTmp {
								continue
							}
						} else {
							continue
						}

						// 余额不足，收益大于余额的10倍
						if userAmount[vVUserBindTraders.UserId].Amount > balanceTmp*10 {
							continue
						}
					} else {
						continue
					}

					// 精度
					if _, ok := symbol[vTraderPositions.Symbol]; !ok {
						continue
					}

					// 发送订单
					wg.Add(1) // 启动一个goroutine就登记+1
					go b.userOrderGoroutine(ctx, &wg, &OrderData{
						Coin:  vTraderPositions.Symbol,
						Type:  vTraderPositions.PositionSide,
						Price: price.Price,
						Side:  vTraderPositions.Side,
						Qty:   "0",
					}, "0", users[vVUserBindTraders.UserId], vVUserBindTraders, symbol[vTraderPositions.Symbol].QuantityPrecision, 1, proportion, 0)
				}

			}
		}

	}

	wg.Wait()

	return nil, nil
}

// InitOrderAfterBindTwo 初始化仓位tfi
func (b *BinanceUserUsecase) InitOrderAfterBindTwo(ctx context.Context, req *v1.InitOrderAfterBindRequest) (*v1.InitOrderAfterBindReply, error) {
	var (
		wg              sync.WaitGroup
		userBindTraders map[uint64][]*UserBindTrader
		traders         map[uint64]*Trader
		userIds         []uint64
		userIdsMap      map[uint64]uint64
		users           map[uint64]*User
		symbol          map[string]*Symbol
		err             error
	)

	userBindTraders, err = b.binanceUserRepo.GetUserBindTraderTwoByInitOrder()
	if nil != err {
		return nil, err
	}

	userIds = make([]uint64, 0)
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserBindTrader := range userBindTraders {
		for _, vVUserBindTrader := range vUserBindTrader {
			if _, ok := userIdsMap[vVUserBindTrader.UserId]; ok {
				continue
			}

			userIdsMap[vVUserBindTrader.UserId] = vVUserBindTrader.UserId
			userIds = append(userIds, vVUserBindTrader.UserId)
		}
	}

	if 0 >= len(userIds) {
		return nil, nil
	}

	traders, err = b.binanceUserRepo.GetTraders()
	if nil != err {
		return nil, nil
	}

	// 获取用户信息，余额信息，收益信息
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		return nil, nil
	}

	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return nil, nil
	}

	for traderId, vUserBindTraders := range userBindTraders {
		var (
			traderPositions           []*TraderPosition
			traderOpeningPositionsNew []*TraderPositionNew
		)

		// 先更新状态
		for _, vVUserBindTraders := range vUserBindTraders {
			if _, ok := users[vVUserBindTraders.UserId]; !ok {
				continue
			}

			// todo 暂时使用检测没用api信息
			if 0 >= users[vVUserBindTraders.UserId].ApiStatus {
				continue
			}

			if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = b.binanceUserRepo.UpdatesUserBindTraderTwoInitOrderById(ctx, vVUserBindTraders.ID)
				if nil != err {
					return err
				}

				return nil
			}); err != nil {
				fmt.Println(err, "修改初始化失败", traderId, vVUserBindTraders)
				continue
			}
		}

		if _, ok := traders[traderId]; !ok {
			fmt.Println("初始化仓位，不存在交易员", traderId)
			continue
		}

		// 老系统的
		traderPositions, err = b.binanceUserRepo.GetTraderPosition(traderId)
		if nil != err {
			fmt.Println("初始化仓位，空仓位", traderId, err)
			continue
		}
		// 按仓位
		for _, vTraderPositions := range traderPositions {
			var (
				price      *BinanceSymbolPrice
				priceFloat float64
			)
			// 查询币价
			price, err = requestBinanceSymbolPrice(vTraderPositions.Symbol)
			if nil != err {
				fmt.Println("初始化仓位，价格查询", traderId, price, err)
				continue
			}

			priceFloat, err = strconv.ParseFloat(price.Price, 64)
			if nil != err {
				fmt.Println("初始化仓位，价格查询，转化", traderId, price, err)
				continue
			}

			// 交易员仓位u占保证金比例
			var proportion float64
			if 0 >= vTraderPositions.Qty || 0 >= traders[traderId].BaseMoney {
				continue
			}

			proportion = vTraderPositions.Qty * priceFloat / traders[traderId].BaseMoney
			if 0 >= proportion || 0 >= priceFloat {
				fmt.Println("初始化仓位，比例计算", vTraderPositions, proportion, vTraderPositions.Qty, priceFloat, traders[traderId].BaseMoney)
				continue
			}

			for _, vVUserBindTraders := range vUserBindTraders {
				if 0 == vVUserBindTraders.Status { // 绑定
					if _, ok := users[vVUserBindTraders.UserId]; !ok {
						continue
					}

					// todo 暂时使用检测没用api信息
					if 0 >= users[vVUserBindTraders.UserId].ApiStatus {
						continue
					}

					// 使用的新系统
					if 1 == users[vVUserBindTraders.UserId].UseNewSystem {
						continue
					}

					// 精度
					if _, ok := symbol[vTraderPositions.Symbol]; !ok {
						continue
					}

					if ("SELL" == vTraderPositions.Side && "SHORT" == vTraderPositions.PositionSide) || ("BUY" == vTraderPositions.Side && "LONG" == vTraderPositions.PositionSide) {
						// todo 暂时只关不开
						if 0 == users[vVUserBindTraders.UserId].IsDai {
							if 5 != vVUserBindTraders.TraderId && 83 != vVUserBindTraders.TraderId {
								continue
							}
						}
					}

					// 发送订单
					wg.Add(1) // 启动一个goroutine就登记+1
					go b.userOrderGoroutineTwo(ctx, &wg, &OrderData{
						Coin:  vTraderPositions.Symbol,
						Type:  vTraderPositions.PositionSide,
						Price: price.Price,
						Side:  vTraderPositions.Side,
						Qty:   "0",
					}, "0", users[vVUserBindTraders.UserId], vVUserBindTraders, symbol[vTraderPositions.Symbol].QuantityPrecision, 1, proportion, 0)
				}
			}
		}

		// 新系统的
		traderOpeningPositionsNew, err = b.binanceUserRepo.GetOpeningTraderPositionNew(traders[traderId].PortfolioId)
		if nil != err {
			fmt.Println("初始化仓位，空仓位", traderId, err)
			continue
		}
		// 按仓位
		for _, vTraderPositions := range traderOpeningPositionsNew {
			var (
				price      *BinanceSymbolPrice
				priceFloat float64
			)
			// 查询币价
			price, err = requestBinanceSymbolPrice(vTraderPositions.Symbol)
			if nil != err {
				fmt.Println("初始化仓位，价格查询", traderId, price, err)
				continue
			}

			priceFloat, err = strconv.ParseFloat(price.Price, 64)
			if nil != err {
				fmt.Println("初始化仓位，价格查询，转化", traderId, vTraderPositions, price, err)
				continue
			}

			// 交易员仓位u占保证金比例
			var proportion float64
			if 0 >= vTraderPositions.Qty || 0 >= traders[traderId].BaseMoney {
				continue
			}

			proportion = vTraderPositions.Qty * priceFloat / traders[traderId].BaseMoney
			if 0 >= proportion || 0 >= priceFloat {
				fmt.Println("初始化仓位，比例计算", vTraderPositions, proportion, vTraderPositions.Qty, priceFloat, traders[traderId].BaseMoney)
				continue
			}

			for _, vVUserBindTraders := range vUserBindTraders {
				if 0 == vVUserBindTraders.Status { // 绑定
					if _, ok := users[vVUserBindTraders.UserId]; !ok {
						continue
					}

					// todo 暂时使用检测没用api信息
					if 0 >= users[vVUserBindTraders.UserId].ApiStatus {
						continue
					}

					// 使用的新系统
					if 1 != users[vVUserBindTraders.UserId].UseNewSystem {
						continue
					}

					// 精度
					if _, ok := symbol[vTraderPositions.Symbol]; !ok {
						continue
					}

					side := "SELL"
					if "SHORT" == vTraderPositions.Side || "LONG" == vTraderPositions.Side {
						if 0 == users[vVUserBindTraders.UserId].IsDai {
							continue
						}
					} else {
						continue
					}

					if "LONG" == vTraderPositions.Side {
						side = "BUY"
					}

					fmt.Println("新系统初始化：", vTraderPositions, users[vVUserBindTraders.UserId].Address)
					// 发送订单
					wg.Add(1) // 启动一个goroutine就登记+1
					go b.userOrderGoroutineTwo(ctx, &wg, &OrderData{
						Coin:  vTraderPositions.Symbol,
						Type:  vTraderPositions.Side,
						Price: price.Price,
						Side:  side,
						Qty:   "0",
					}, "0", users[vVUserBindTraders.UserId], vVUserBindTraders, symbol[vTraderPositions.Symbol].QuantityPrecision, 1, proportion, 0)
				}
			}
		}
	}

	wg.Wait()

	return nil, nil
}

func (b *BinanceUserUsecase) ExchangeUserLeverAge(ctx context.Context, req *v1.ExchangeUserLeverAgeRequest) (*v1.ExchangeUserLeverAgeReply, error) {
	var (
		symbol map[string]*Symbol
		user   *User
		err    error
	)
	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return nil, err
	}

	user, err = b.binanceUserRepo.GetUserById(ctx, req.UserId)
	if nil != err {
		return nil, err
	}

	if nil == user {
		return nil, nil
	}

	if "all" == req.Symbol {
		for k, _ := range symbol {
			fmt.Println("仓位杠杆修改：", k, int64(5), user.ApiKey, user.ApiSecret)
			_, err = requestBinanceLeverAge(k, int64(5), user.ApiKey, user.ApiSecret)
			if nil != err {
				fmt.Println(err)
				return nil, err
			}
		}
	} else {
		for k, _ := range symbol {
			if k != req.Symbol {
				continue
			}

			fmt.Println("仓位杠杆修改：", k, int64(5), user.ApiKey, user.ApiSecret)
			_, err = requestBinanceLeverAge(k, int64(5), user.ApiKey, user.ApiSecret)
			if nil != err {
				fmt.Println(err)
				return nil, err
			}
		}
	}

	return nil, nil
}

// OrderAdminTwo 管理员下单手动
func (b *BinanceUserUsecase) OrderAdminTwo(ctx context.Context, req *v1.OrderAdminTwoRequest) (*v1.OrderAdminTwoReply, error) {
	var (
		wg              sync.WaitGroup
		userBindTraders map[uint64][]*UserBindTrader
		traders         map[uint64]*Trader
		userIds         []uint64
		userIdsMap      map[uint64]uint64
		users           map[uint64]*User
		symbol          map[string]*Symbol
		err             error
	)

	// 指定用户
	if 0 >= req.UserId || 0 >= len(req.Symbol) || 0 >= req.TraderId {
		return nil, nil
	}

	userBindTraders, err = b.binanceUserRepo.GetUserBindTraderTwoByAlreadyInitOrder()
	if nil != err {
		return nil, err
	}

	userIds = make([]uint64, 0)
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserBindTrader := range userBindTraders {
		for _, vVUserBindTrader := range vUserBindTrader {
			// 需要的
			if vVUserBindTrader.UserId != req.UserId {
				continue
			}

			if _, ok := userIdsMap[vVUserBindTrader.UserId]; ok {
				continue
			}

			userIdsMap[vVUserBindTrader.UserId] = vVUserBindTrader.UserId
			userIds = append(userIds, vVUserBindTrader.UserId)
		}
	}

	if 0 >= len(userIds) {
		return nil, nil
	}

	traders, err = b.binanceUserRepo.GetTraders()
	if nil != err {
		return nil, nil
	}

	// 获取用户信息，余额信息，收益信息
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		return nil, nil
	}

	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return nil, err
	}

	for traderId, vUserBindTraders := range userBindTraders {
		if req.TraderId != traderId { // 参数交易员
			continue
		}

		var (
			traderOpeningPositionsNew []*TraderPositionNew
		)

		if _, ok := traders[traderId]; !ok {
			fmt.Println("初始化仓位，不存在交易员", traderId)
			continue
		}

		// 新系统的
		traderOpeningPositionsNew, err = b.binanceUserRepo.GetOpeningTraderPositionNew(traders[traderId].PortfolioId)
		if nil != err {
			fmt.Println("初始化仓位，空仓位", traderId, err)
			continue
		}
		// 按仓位
		for _, vTraderPositions := range traderOpeningPositionsNew {
			var (
				price      *BinanceSymbolPrice
				priceFloat float64
			)

			if req.Symbol != vTraderPositions.Symbol { // 参数仓位
				continue
			}

			// 查询币价
			price, err = requestBinanceSymbolPrice(vTraderPositions.Symbol)
			if nil != err {
				fmt.Println("初始化仓位，价格查询", traderId, price, err)
				continue
			}

			priceFloat, err = strconv.ParseFloat(price.Price, 64)
			if nil != err {
				fmt.Println("初始化仓位，价格查询，转化", traderId, vTraderPositions, price, err)
				continue
			}

			// 交易员仓位u占保证金比例
			var proportion float64
			if 0 >= vTraderPositions.Qty || 0 >= traders[traderId].BaseMoney {
				continue
			}

			proportion = vTraderPositions.Qty * priceFloat / traders[traderId].BaseMoney
			if 0 >= proportion || 0 >= priceFloat {
				fmt.Println("初始化仓位，比例计算", vTraderPositions, proportion, vTraderPositions.Qty, priceFloat, traders[traderId].BaseMoney)
				continue
			}

			for _, vVUserBindTraders := range vUserBindTraders {
				if req.UserId != vVUserBindTraders.UserId { // 参数用户
					continue
				}

				if 0 == vVUserBindTraders.Status { // 绑定
					if _, ok := users[vVUserBindTraders.UserId]; !ok {
						continue
					}

					// todo 暂时使用检测没用api信息
					if 0 >= users[vVUserBindTraders.UserId].ApiStatus {
						continue
					}

					// 使用的新系统
					if 1 != users[vVUserBindTraders.UserId].UseNewSystem {
						continue
					}

					// 精度
					if _, ok := symbol[vTraderPositions.Symbol]; !ok {
						continue
					}

					side := "SELL"
					if "SHORT" == vTraderPositions.Side || "LONG" == vTraderPositions.Side {
						if 0 == users[vVUserBindTraders.UserId].IsDai {
							continue
						}
					} else {
						continue
					}

					if "LONG" == vTraderPositions.Side {
						side = "BUY"
					}

					for k, _ := range symbol {
						if k != vTraderPositions.Symbol {
							continue
						}
						fmt.Println("新系统仓位自定义补充，仓位杠杆修改：", k, int64(5), users[vVUserBindTraders.UserId].ApiKey, users[vVUserBindTraders.UserId].ApiSecret)
						_, err = requestBinanceLeverAge(k, int64(5), users[vVUserBindTraders.UserId].ApiKey, users[vVUserBindTraders.UserId].ApiSecret)
						if nil != err {
							fmt.Println(err)
							return nil, err
						}
					}

					fmt.Println("新系统仓位自定义补充：", vTraderPositions, users[vVUserBindTraders.UserId].Address, side)
					// 发送订单
					wg.Add(1) // 启动一个goroutine就登记+1
					go b.userOrderGoroutineTwo(ctx, &wg, &OrderData{
						Coin:  vTraderPositions.Symbol,
						Type:  vTraderPositions.Side,
						Price: price.Price,
						Side:  side,
						Qty:   "0",
					}, "0", users[vVUserBindTraders.UserId], vVUserBindTraders, symbol[vTraderPositions.Symbol].QuantityPrecision, 1, proportion, 0)
				}
			}
		}
	}

	wg.Wait()

	return nil, nil
}

// OverOrder 平仓
func (b *BinanceUserUsecase) OverOrder(ctx context.Context, req *v1.OverOrderAfterBindRequest) (*v1.OverOrderAfterBindReply, error) {
	var (
		wg              sync.WaitGroup
		userBindTraders map[uint64][]*UserBindTrader
		userIds         []uint64
		userIdsMap      map[uint64]uint64
		users           map[uint64]*User
		symbol          map[string]*Symbol
		err             error
	)

	userBindTraders, err = b.binanceUserRepo.GetUserBindTraderByAlreadyInitOrder()
	if nil != err {
		return nil, err
	}

	userIds = make([]uint64, 0)
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserBindTrader := range userBindTraders {
		for _, vVUserBindTrader := range vUserBindTrader {
			if _, ok := userIdsMap[vVUserBindTrader.UserId]; ok {
				continue
			}

			userIdsMap[vVUserBindTrader.UserId] = vVUserBindTrader.UserId
			userIds = append(userIds, vVUserBindTrader.UserId)
		}
	}

	if 0 >= len(userIds) {
		return nil, nil
	}

	// 获取用户信息，余额信息，收益信息
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		return nil, nil
	}

	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return nil, nil
	}

	for traderId, vUserBindTraders := range userBindTraders {
		var (
			traderPositions []*TraderPosition
		)

		traderPositions, err = b.binanceUserRepo.GetTraderPosition(traderId)
		if nil != err {
			fmt.Println("初始化仓位，空仓位", traderId, err)
			continue
		}

		// 按仓位
		for _, vTraderPositions := range traderPositions {
			if 0 < vTraderPositions.Qty {
				continue
			}

			for _, vVUserBindTraders := range vUserBindTraders {
				if 0 == vVUserBindTraders.Status { // 绑定
					if _, ok := users[vVUserBindTraders.UserId]; !ok {
						continue
					}

					// todo 暂时使用检测没用api信息
					if 0 >= users[vVUserBindTraders.UserId].ApiStatus {
						continue
					}

					// 精度
					if _, ok := symbol[vTraderPositions.Symbol]; !ok {
						continue
					}

					// 平仓
					var side string
					if "LONG" == vTraderPositions.PositionSide {
						side = "SELL"
					} else if "SHORT" == vTraderPositions.PositionSide {
						side = "BUY"
					} else {
						continue
					}

					// 订单统计
					var currentOrders []*UserOrder
					currentOrders, err = b.binanceUserRepo.GetUserOrderByUserTraderIdAndSymbol(vVUserBindTraders.UserId, vVUserBindTraders.TraderId, vTraderPositions.Symbol)
					if nil != err {
						fmt.Println(err)
						continue
					}
					// 查出用户的BUY单的币的数量，在对应的trader下，超过了BUY不能SELL todo 使用数据库量太大以后
					if 0 >= len(currentOrders) {
						continue
					}
					var historyQuantityFloat float64
					// 多的部分不管，按剩余的数量关 todo 交易员全部平仓，少的部分另一个程序解决
					for _, vCurrentOrders := range currentOrders {
						// 本次平多
						if "SELL" == side && "LONG" == vTraderPositions.PositionSide {
							// 历史开多和平多
							if "BUY" == vCurrentOrders.Side && "LONG" == vCurrentOrders.PositionSide {
								historyQuantityFloat += vCurrentOrders.ExecutedQty
							} else if "SELL" == vCurrentOrders.Side && "LONG" == vCurrentOrders.PositionSide {
								historyQuantityFloat -= vCurrentOrders.ExecutedQty
							}
						} else if "BUY" == side && "SHORT" == vTraderPositions.PositionSide {
							// 历史开空和平空
							if "SELL" == vCurrentOrders.Side && "SHORT" == vCurrentOrders.PositionSide {
								historyQuantityFloat += vCurrentOrders.ExecutedQty
							} else if "BUY" == vCurrentOrders.Side && "SHORT" == vCurrentOrders.PositionSide {
								historyQuantityFloat -= vCurrentOrders.ExecutedQty
							}
						}
					}

					// 开单历史数量不足了
					if isZero(historyQuantityFloat) || 0 > historyQuantityFloat {
						continue
					}

					// 发送订单
					wg.Add(1) // 启动一个goroutine就登记+1
					go b.userOrderGoroutine(ctx, &wg, &OrderData{
						Coin:  vTraderPositions.Symbol,
						Type:  vTraderPositions.PositionSide,
						Price: "0",
						Side:  side,
						Qty:   "0",
					}, "0", users[vVUserBindTraders.UserId], vVUserBindTraders, symbol[vTraderPositions.Symbol].QuantityPrecision, 0, 0, 1)
				}

			}
		}

	}

	wg.Wait()

	return nil, nil
}

// OverOrderTwo 平仓
func (b *BinanceUserUsecase) OverOrderTwo(ctx context.Context, req *v1.OverOrderAfterBindRequest) (*v1.OverOrderAfterBindReply, error) {
	var (
		wg              sync.WaitGroup
		userBindTraders map[uint64][]*UserBindTrader
		userIds         []uint64
		userIdsMap      map[uint64]uint64
		users           map[uint64]*User
		symbol          map[string]*Symbol
		err             error
	)

	userBindTraders, err = b.binanceUserRepo.GetUserBindTraderTwoByAlreadyInitOrder()
	if nil != err {
		return nil, err
	}

	userIds = make([]uint64, 0)
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserBindTrader := range userBindTraders {
		for _, vVUserBindTrader := range vUserBindTrader {
			if _, ok := userIdsMap[vVUserBindTrader.UserId]; ok {
				continue
			}

			userIdsMap[vVUserBindTrader.UserId] = vVUserBindTrader.UserId
			userIds = append(userIds, vVUserBindTrader.UserId)
		}
	}

	if 0 >= len(userIds) {
		return nil, nil
	}

	// 获取用户信息，余额信息，收益信息
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		return nil, nil
	}

	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return nil, nil
	}

	for traderId, vUserBindTraders := range userBindTraders {
		var (
			traderPositions []*TraderPosition
		)

		traderPositions, err = b.binanceUserRepo.GetTraderPosition(traderId)
		if nil != err {
			fmt.Println("初始化仓位，空仓位", traderId, err)
			continue
		}

		// 按仓位
		for _, vTraderPositions := range traderPositions {
			if 0 < vTraderPositions.Qty {
				continue
			}

			for _, vVUserBindTraders := range vUserBindTraders {
				if 0 == vVUserBindTraders.Status { // 绑定
					if _, ok := users[vVUserBindTraders.UserId]; !ok {
						continue
					}

					// todo 暂时使用检测没用api信息
					if 0 >= users[vVUserBindTraders.UserId].ApiStatus {
						continue
					}

					// 新系统用户不使用
					if 1 == users[vVUserBindTraders.UserId].UseNewSystem {
						continue
					}

					// 精度
					if _, ok := symbol[vTraderPositions.Symbol]; !ok {
						continue
					}

					// 平仓
					var side string
					if "LONG" == vTraderPositions.PositionSide {
						side = "SELL"
					} else if "SHORT" == vTraderPositions.PositionSide {
						side = "BUY"
					} else {
						continue
					}

					// 订单统计
					var currentOrders []*UserOrder
					currentOrders, err = b.binanceUserRepo.GetUserOrderTwoByUserTraderIdAndSymbol(vVUserBindTraders.UserId, vVUserBindTraders.TraderId, vTraderPositions.Symbol)
					if nil != err {
						fmt.Println(err)
						continue
					}
					// 查出用户的BUY单的币的数量，在对应的trader下，超过了BUY不能SELL todo 使用数据库量太大以后
					if 0 >= len(currentOrders) {
						continue
					}
					var historyQuantityFloat float64
					// 多的部分不管，按剩余的数量关 todo 交易员全部平仓，少的部分另一个程序解决
					for _, vCurrentOrders := range currentOrders {
						// 本次平多
						if "SELL" == side && "LONG" == vTraderPositions.PositionSide {
							// 历史开多和平多
							if "BUY" == vCurrentOrders.Side && "LONG" == vCurrentOrders.PositionSide {
								historyQuantityFloat += vCurrentOrders.ExecutedQty
							} else if "SELL" == vCurrentOrders.Side && "LONG" == vCurrentOrders.PositionSide {
								historyQuantityFloat -= vCurrentOrders.ExecutedQty
							}
						} else if "BUY" == side && "SHORT" == vTraderPositions.PositionSide {
							// 历史开空和平空
							if "SELL" == vCurrentOrders.Side && "SHORT" == vCurrentOrders.PositionSide {
								historyQuantityFloat += vCurrentOrders.ExecutedQty
							} else if "BUY" == vCurrentOrders.Side && "SHORT" == vCurrentOrders.PositionSide {
								historyQuantityFloat -= vCurrentOrders.ExecutedQty
							}
						}
					}

					// 开单历史数量不足了
					if isZero(historyQuantityFloat) || 0 > historyQuantityFloat {
						continue
					}

					// 发送订单
					wg.Add(1) // 启动一个goroutine就登记+1
					go b.userOrderGoroutineTwo(ctx, &wg, &OrderData{
						Coin:  vTraderPositions.Symbol,
						Type:  vTraderPositions.PositionSide,
						Price: "0",
						Side:  side,
						Qty:   "0",
					}, "0", users[vVUserBindTraders.UserId], vVUserBindTraders, symbol[vTraderPositions.Symbol].QuantityPrecision, 0, 0, 1)
				}

			}
		}

	}

	wg.Wait()

	return nil, nil
}

// AdminOverOrder 平仓
func (b *BinanceUserUsecase) AdminOverOrder(ctx context.Context, req *v1.OverOrderAfterBindRequest) (*v1.OverOrderAfterBindReply, error) {
	var (
		wg              sync.WaitGroup
		userBindTraders map[uint64][]*UserBindTrader
		userIds         []uint64
		userIdsMap      map[uint64]uint64
		users           map[uint64]*User
		symbol          map[string]*Symbol
		err             error
	)

	userBindTraders, err = b.binanceUserRepo.GetUserBindTraderByOverOrder()
	if nil != err {
		return nil, err
	}

	userIds = make([]uint64, 0)
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserBindTrader := range userBindTraders {
		for _, vVUserBindTrader := range vUserBindTrader {
			if _, ok := userIdsMap[vVUserBindTrader.UserId]; ok {
				continue
			}

			userIdsMap[vVUserBindTrader.UserId] = vVUserBindTrader.UserId
			userIds = append(userIds, vVUserBindTrader.UserId)
		}
	}

	if 0 >= len(userIds) {
		return nil, nil
	}

	// 获取用户信息，余额信息，收益信息
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		return nil, nil
	}

	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return nil, nil
	}

	for traderId, vUserBindTraders := range userBindTraders {
		var (
			traderPositions []*TraderPosition
		)

		traderPositions, err = b.binanceUserRepo.GetTraderPosition(traderId)
		if nil != err {
			fmt.Println("初始化仓位，空仓位", traderId, err)
			continue
		}

		// 按仓位
		for _, vTraderPositions := range traderPositions {
			for _, vVUserBindTraders := range vUserBindTraders {
				if 0 == vVUserBindTraders.Status { // 绑定
					if _, ok := users[vVUserBindTraders.UserId]; !ok {
						continue
					}

					// todo 暂时使用检测没用api信息
					if 0 >= users[vVUserBindTraders.UserId].ApiStatus {
						continue
					}

					// 精度
					if _, ok := symbol[vTraderPositions.Symbol]; !ok {
						continue
					}

					// 平仓
					var side string
					if "LONG" == vTraderPositions.PositionSide {
						side = "SELL"
					} else if "SHORT" == vTraderPositions.PositionSide {
						side = "BUY"
					} else {
						continue
					}

					// 发送订单
					wg.Add(1) // 启动一个goroutine就登记+1
					go b.userOrderGoroutine(ctx, &wg, &OrderData{
						Coin:  vTraderPositions.Symbol,
						Type:  vTraderPositions.PositionSide,
						Price: "0",
						Side:  side,
						Qty:   "0",
					}, "0", users[vVUserBindTraders.UserId], vVUserBindTraders, symbol[vTraderPositions.Symbol].QuantityPrecision, 0, 0, 1)
				}

			}
		}

	}

	wg.Wait()

	return nil, nil
}

// AdminOverOrderTwo 平仓
func (b *BinanceUserUsecase) AdminOverOrderTwo(ctx context.Context, req *v1.OverOrderAfterBindRequest) (*v1.OverOrderAfterBindReply, error) {
	return nil, b.handleAdminOverOrderTwo(ctx)
}

func (b *BinanceUserUsecase) AdminOverOrderTwoByInfo(ctx context.Context, req *v1.AdminOverOrderTwoByInfoRequest) (*v1.AdminOverOrderTwoByInfoReply, error) {
	var (
		err      error
		side     string
		quantity string
	)
	var orderInfo *OrderInfo
	if "SHORT" == req.PositionSide {
		side = "BUY"
	} else if "LONG" == req.PositionSide {
		side = "SELL"
	} else {
		fmt.Println("错误的自定义关仓方向", req)
		return nil, nil
	}
	if 0 > req.Num {
		fmt.Println("错误的自定义关仓数量", req)
		return nil, nil
	}

	quantity = strconv.FormatFloat(req.Num, 'f', -1, 64)

	// 请求下单
	_, orderInfo, err = requestBinanceOrder(req.Symbol, side, "MARKET", req.PositionSide, quantity, req.ApiKey, req.ApiSecret)
	if nil != err {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(orderInfo)
	return nil, nil
}

func (b *BinanceUserUsecase) handleAdminOverOrderTwo(ctx context.Context) error {
	var (
		wg              sync.WaitGroup
		userBindTraders map[uint64][]*UserBindTrader
		userIds         []uint64
		userIdsMap      map[uint64]uint64
		users           map[uint64]*User
		symbol          map[string]*Symbol
		err             error
	)

	userBindTraders, err = b.binanceUserRepo.GetUserBindTraderTwoByOverOrder()
	if nil != err {
		return err
	}

	userIds = make([]uint64, 0)
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserBindTrader := range userBindTraders {
		for _, vVUserBindTrader := range vUserBindTrader {
			if _, ok := userIdsMap[vVUserBindTrader.UserId]; ok {
				continue
			}

			userIdsMap[vVUserBindTrader.UserId] = vVUserBindTrader.UserId
			userIds = append(userIds, vVUserBindTrader.UserId)
		}
	}

	if 0 >= len(userIds) {
		return nil
	}

	// 获取用户信息，余额信息，收益信息
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		return nil
	}

	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return nil
	}

	for traderId, vUserBindTraders := range userBindTraders {
		var (
			traderPositions []*TraderPosition
		)

		traderPositions, err = b.binanceUserRepo.GetTraderPosition(traderId)
		if nil != err {
			fmt.Println("初始化仓位，空仓位", traderId, err)
			continue
		}

		// 按仓位
		for _, vTraderPositions := range traderPositions {
			for _, vVUserBindTraders := range vUserBindTraders {
				if 0 == vVUserBindTraders.Status { // 绑定
					if _, ok := users[vVUserBindTraders.UserId]; !ok {
						continue
					}

					// todo 暂时使用检测没用api信息
					if 0 >= users[vVUserBindTraders.UserId].ApiStatus {
						continue
					}

					// 精度
					if _, ok := symbol[vTraderPositions.Symbol]; !ok {
						continue
					}

					// 平仓
					var side string
					if "LONG" == vTraderPositions.PositionSide {
						side = "SELL"
					} else if "SHORT" == vTraderPositions.PositionSide {
						side = "BUY"
					} else {
						continue
					}

					// 发送订单
					wg.Add(1) // 启动一个goroutine就登记+1
					go b.userOrderGoroutineTwo(ctx, &wg, &OrderData{
						Coin:  vTraderPositions.Symbol,
						Type:  vTraderPositions.PositionSide,
						Price: "0",
						Side:  side,
						Qty:   "0",
					}, "0", users[vVUserBindTraders.UserId], vVUserBindTraders, symbol[vTraderPositions.Symbol].QuantityPrecision, 0, 0, 1)
				}

			}
		}

	}

	wg.Wait()

	return nil
}

// GetTermBinanceCurrentBalance 平仓
func (b *BinanceUserUsecase) GetTermBinanceCurrentBalance(ctx context.Context, userId uint64) (float64, error) {
	var (
		user           *User
		binanceBalance *BinanceUserBalance
		balanceFloat   float64
		crossUnPnl     float64
		err            error
	)

	user, err = b.binanceUserRepo.GetUserById(ctx, userId)
	if nil != err {
		return 0, err
	}

	binanceBalance, err = requestBinanceUserBalance(user.ApiKey, user.ApiSecret)
	if nil != err {
		return 0, err
	}

	balanceFloat, err = strconv.ParseFloat(binanceBalance.Balance, 64)
	if nil != err {
		return 0, err
	}

	crossUnPnl, err = strconv.ParseFloat(binanceBalance.CrossUnPnl, 64)
	if nil != err {
		return 0, err
	}

	return balanceFloat + crossUnPnl, nil

}

// 以下是请求binance的api

func requestBinanceOrder(symbol string, side string, orderType string, positionSide string, quantity string, apiKey string, secretKey string) (*BinanceOrder, *OrderInfo, error) {
	var (
		client       *http.Client
		req          *http.Request
		resp         *http.Response
		res          *BinanceOrder
		resOrderInfo *OrderInfo
		data         string
		b            []byte
		err          error
		apiUrl       = "https://fapi.binance.com/fapi/v1/order"
	)

	//fmt.Println(symbol, side, orderType, positionSide, quantity, apiKey, secretKey)
	// 时间
	now := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	// 拼请求数据
	data = "symbol=" + symbol + "&side=" + side + "&type=" + orderType + "&positionSide=" + positionSide + "&newOrderRespType=" + "RESULT" + "&quantity=" + quantity + "&timestamp=" + now

	// 加密
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	// 构造请求

	req, err = http.NewRequest("POST", apiUrl, strings.NewReader(data+"&signature="+signature))
	if err != nil {
		return nil, nil, err
	}
	// 添加头信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-MBX-APIKEY", apiKey)

	// 请求执行
	client = &http.Client{Timeout: 3 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(b), err)
		return nil, nil, err
	}

	var o BinanceOrder
	err = json.Unmarshal(b, &o)
	if err != nil {
		fmt.Println(string(b), err)
		return nil, nil, err
	}

	res = &BinanceOrder{
		OrderId:       o.OrderId,
		ExecutedQty:   o.ExecutedQty,
		ClientOrderId: o.ClientOrderId,
		Symbol:        o.Symbol,
		AvgPrice:      o.AvgPrice,
		CumQuote:      o.CumQuote,
		Side:          o.Side,
		PositionSide:  o.PositionSide,
		ClosePosition: o.ClosePosition,
		Type:          o.Type,
	}

	if 0 >= res.OrderId {
		//fmt.Println(string(b))
		err = json.Unmarshal(b, &resOrderInfo)
		if err != nil {
			fmt.Println(string(b), err)
			return nil, nil, err
		}
	}

	return res, resOrderInfo, nil
}

func requestBinanceOrderStop(symbol string, side string, positionSide string, quantity string, stopPrice string, price string, apiKey string, secretKey string) (*BinanceOrder, *OrderInfo, error) {
	//fmt.Println(symbol, side, positionSide, quantity, stopPrice, price, apiKey, secretKey)
	var (
		client       *http.Client
		req          *http.Request
		resp         *http.Response
		res          *BinanceOrder
		resOrderInfo *OrderInfo
		data         string
		b            []byte
		err          error
		apiUrl       = "https://fapi.binance.com/fapi/v1/order"
	)

	// 时间
	now := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	// 拼请求数据
	data = "symbol=" + symbol + "&side=" + side + "&type=STOP&stopPrice=" + stopPrice + "&price=" + price + "&positionSide=" + positionSide + "&newOrderRespType=" + "RESULT" + "&quantity=" + quantity + "&timestamp=" + now

	// 加密
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	// 构造请求

	req, err = http.NewRequest("POST", apiUrl, strings.NewReader(data+"&signature="+signature))
	if err != nil {
		return nil, nil, err
	}
	// 添加头信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-MBX-APIKEY", apiKey)

	// 请求执行
	client = &http.Client{Timeout: 3 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(b), err)
		return nil, nil, err
	}

	var o BinanceOrder
	err = json.Unmarshal(b, &o)
	if err != nil {
		fmt.Println(string(b), err)
		return nil, nil, err
	}
	//fmt.Println(string(b), err)
	res = &BinanceOrder{
		OrderId:       o.OrderId,
		ExecutedQty:   o.ExecutedQty,
		ClientOrderId: o.ClientOrderId,
		Symbol:        o.Symbol,
		AvgPrice:      o.AvgPrice,
		CumQuote:      o.CumQuote,
		Side:          o.Side,
		PositionSide:  o.PositionSide,
		ClosePosition: o.ClosePosition,
		Type:          o.Type,
	}

	if 0 >= res.OrderId {
		err = json.Unmarshal(b, &resOrderInfo)
		if err != nil {
			fmt.Println(string(b), err)
			return nil, nil, err
		}
	}

	return res, resOrderInfo, nil
}

// 更改杠杆倍率
func requestBinanceLeverAge(symbol string, leverAge int64, apiKey string, secretKey string) (*BinanceLeverAge, error) {
	var (
		client *http.Client
		req    *http.Request
		resp   *http.Response
		res    *BinanceLeverAge
		data   string
		b      []byte
		err    error
		apiUrl = "https://fapi.binance.com/fapi/v1/leverage"
	)

	// 时间
	now := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	// 拼请求数据
	data = "symbol=" + symbol + "&leverage=" + strconv.FormatInt(leverAge, 10) + "&timestamp=" + now
	// 加密
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	// 构造请求

	req, err = http.NewRequest("POST", apiUrl, strings.NewReader(data+"&signature="+signature))
	if err != nil {
		return nil, err
	}
	// 添加头信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-MBX-APIKEY", apiKey)

	// 请求执行
	client = &http.Client{Timeout: 3 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var l BinanceLeverAge
	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res = &BinanceLeverAge{
		LeverAge: l.LeverAge,
		symbol:   l.symbol,
	}

	return res, nil
}

// 更改逐全仓模式
func requestBinanceMarginType(symbol string, apiKey string, secretKey string) (*BinanceMarginType, error) {
	var (
		client *http.Client
		req    *http.Request
		resp   *http.Response
		res    *BinanceMarginType
		data   string
		b      []byte
		err    error
		apiUrl = "https://fapi.binance.com/fapi/v1/marginType"
	)

	// 时间
	now := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	// 拼请求数据
	data = "symbol=" + symbol + "&marginType=CROSSED" + "&timestamp=" + now
	// 加密
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	// 构造请求

	req, err = http.NewRequest("POST", apiUrl, strings.NewReader(data+"&signature="+signature))
	if err != nil {
		return nil, err
	}
	// 添加头信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-MBX-APIKEY", apiKey)

	// 请求执行
	client = &http.Client{Timeout: 3 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var l BinanceMarginType
	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res = &BinanceMarginType{
		Code: l.Code,
		Msg:  l.Msg,
	}

	return res, nil
}

// 更改持仓模式
func requestBinancePositionSide(apiKey string, secretKey string) (*BinancePositionSide, error) {
	var (
		client *http.Client
		req    *http.Request
		resp   *http.Response
		res    *BinancePositionSide
		data   string
		b      []byte
		err    error
		apiUrl = "https://fapi.binance.com/fapi/v1/positionSide/dual"
	)

	// 时间
	now := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	// 拼请求数据
	data = "dualSidePosition=true&timestamp=" + now
	// 加密
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	// 构造请求

	req, err = http.NewRequest("POST", apiUrl, strings.NewReader(data+"&signature="+signature))
	if err != nil {
		return nil, err
	}
	// 添加头信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-MBX-APIKEY", apiKey)

	// 请求执行
	client = &http.Client{Timeout: 3 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var l BinancePositionSide
	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res = &BinancePositionSide{
		Code: l.Code,
		Msg:  l.Msg,
	}

	return res, nil
}

// 交易对币价
func requestBinanceSymbolPrice(symbol string) (*BinanceSymbolPrice, error) {
	var (
		client *http.Client
		req    *http.Request
		resp   *http.Response
		res    *BinanceSymbolPrice
		data   string
		b      []byte
		err    error
		apiUrl = "https://fapi.binance.com/fapi/v2/ticker/price?symbol=" + symbol
	)

	// 拼请求数据
	req, err = http.NewRequest("GET", apiUrl, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	// 添加头信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 请求执行
	client = &http.Client{Timeout: 3 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var l BinanceSymbolPrice
	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res = &BinanceSymbolPrice{
		Symbol: l.Symbol,
		Price:  l.Price,
	}

	return res, nil
}

func requestBinanceOrderInfo(symbol string, orderId int64, apiKey string, secretKey string) (*BinanceOrder, error) {
	var (
		client *http.Client
		req    *http.Request
		resp   *http.Response
		res    *BinanceOrder
		data   string
		b      []byte
		err    error
		apiUrl = "https://fapi.binance.com/fapi/v1/order"
	)

	// 时间
	now := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	// 拼请求数据
	data = "symbol=" + symbol + "&orderId=" + strconv.FormatInt(orderId, 10) + "&timestamp=" + now
	// 加密
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	// 构造请求

	req, err = http.NewRequest("GET", apiUrl, strings.NewReader(data+"&signature="+signature))
	if err != nil {
		return nil, err
	}
	// 添加头信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-MBX-APIKEY", apiKey)

	// 请求执行
	client = &http.Client{Timeout: 3 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var o BinanceOrder
	err = json.Unmarshal(b, &o)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res = &BinanceOrder{
		OrderId:       o.OrderId,
		ExecutedQty:   o.ExecutedQty,
		ClientOrderId: o.ClientOrderId,
		Symbol:        o.Symbol,
		AvgPrice:      o.AvgPrice,
		CumQuote:      o.CumQuote,
		Side:          o.Side,
		PositionSide:  o.PositionSide,
		ClosePosition: o.ClosePosition,
		Type:          o.Type,
		Status:        o.Status,
	}

	return res, nil
}

func requestBinanceOrderHistory(apiKey string, secretKey string, symbol string, orderId int64, startTime string, endTime string) ([]*OrderHistory, error) {
	var (
		client *http.Client
		req    *http.Request
		resp   *http.Response
		res    []*OrderHistory
		data   string
		b      []byte
		err    error
		apiUrl = "https://fapi.binance.com/fapi/v1/userTrades"
	)

	// 时间
	now := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	// 拼请求数据
	if 0 < len(startTime) && 0 < len(endTime) {
		data = "symbol=" + symbol + "&startTime=" + startTime + "&endTime=" + endTime + "&limit=1000&timestamp=" + now
	} else if 0 < orderId {
		data = "symbol=" + symbol + "&orderId=" + strconv.FormatInt(orderId, 10) + "&timestamp=" + now
	} else {
		return nil, nil
	}

	// 加密
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	// 构造请求

	req, err = http.NewRequest("GET", apiUrl+"?"+data+"&signature="+signature, nil)
	if err != nil {
		return nil, err
	}

	// 添加头信息
	req.Header.Set("X-MBX-APIKEY", apiKey)

	// 请求执行
	client = &http.Client{Timeout: 3 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(b))
		fmt.Println(err)
		return nil, err
	}

	var i []*OrderHistory
	err = json.Unmarshal(b, &i)
	if err != nil {
		fmt.Println(string(b))
		return nil, err
	}

	res = make([]*OrderHistory, 0)
	for _, v := range i {
		res = append(res, v)
	}

	return res, nil
}

// 用户余额
func requestBinanceUserBalance(apiKey string, secretKey string) (*BinanceUserBalance, error) {
	var (
		client *http.Client
		req    *http.Request
		resp   *http.Response
		res    *BinanceUserBalance
		data   string
		b      []byte
		err    error
		apiUrl = "https://fapi.binance.com/fapi/v2/balance"
	)

	// 时间
	now := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	// 拼请求数据
	data = "timestamp=" + now
	// 加密
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	// 构造请求

	req, err = http.NewRequest("GET", apiUrl, strings.NewReader(data+"&signature="+signature))
	if err != nil {
		return nil, err
	}
	// 添加头信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-MBX-APIKEY", apiKey)

	// 请求执行
	client = &http.Client{Timeout: 3 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var o []*BinanceUserBalance
	err = json.Unmarshal(b, &o)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, v := range o {
		if "USDT" == v.Asset {
			res = &BinanceUserBalance{
				Balance:            v.Balance,
				CrossUnPnl:         v.CrossUnPnl,
				AvailableBalance:   v.AvailableBalance,
				CrossWalletBalance: v.CrossWalletBalance,
				MaxWithdrawAmount:  v.MaxWithdrawAmount,
			}
		}
	}

	return res, nil
}

func (b *BinanceUserUsecase) PullBinanceTradeHistory(ctx context.Context) error {
	var (
		binanceTrade []*BinanceTrade
		wg           sync.WaitGroup
		proxy        []*ProxyData
		err          error
	)

	binanceTrade, err = b.binanceUserRepo.GetBinanceTrade()
	if nil != err {
		return err
	}

	proxy, err = requestProxy()
	if nil != err {
		return err
	}

	for k, v := range binanceTrade {
		wg.Add(1) // 启动一个goroutine就登记+1
		go b.pullBinanceTradeGoroutine(ctx, &wg, k, v, proxy)
	}

	wg.Wait()
	return err
}

func (b *BinanceUserUsecase) pullBinanceTradeGoroutine(ctx context.Context, wg *sync.WaitGroup, kProxy int, v *BinanceTrade, proxy []*ProxyData) {
	defer wg.Done() // goroutine结束就登记-1

	// 数据库对比数据
	var (
		compareMax                     = 50 // 预设最大对比条数，小于最大限制50条
		currentCompareMax              int  // 实际获得对比条数
		err                            error
		binanceTradeHistoryNewestGroup []*BinanceTradeHistory
	)

	binanceTradeHistoryNewestGroup, err = b.binanceUserRepo.GetBinanceTradeHistoryByTraderNumNewest(v.TraderNum, compareMax)
	if nil != err {
		return
	}
	currentCompareMax = len(binanceTradeHistoryNewestGroup)
	fmt.Println(currentCompareMax)

	var (
		wg1            sync.WaitGroup
		requestMaxPage = int64(1) // 一次查10页，一页用俩ip，最多分20个ip一个号
		proxyKMin      int64
		proxyKMax      int64
	)
	// 每人人分20个代理
	if kProxy*20+20 <= len(proxy) {
		proxyKMin = int64(kProxy) * 20
		proxyKMax = int64(kProxy)*20 + 19
	} else {
		fmt.Println("代理数量不足", kProxy, v, len(proxy))
		return
	}

	var binanceTradeHistoryByPage sync.Map
	//binanceTradeHistoryByPage = make(map[int64][]*BinanceTradeHistoryDataList, 0)
	for i := int64(0); i < requestMaxPage; i++ {
		if proxyKMin+i+requestMaxPage > proxyKMax {
			fmt.Println("错误的分配ip", proxyKMax, proxyKMin+i+requestMaxPage)
		}

		// 分配ip
		tmpProxy := make([]*ProxyData, 0)
		tmpProxy = append(tmpProxy, proxy[proxyKMin+i])
		tmpProxy = append(tmpProxy, proxy[proxyKMin+i+requestMaxPage])

		wg1.Add(1) // 启动一个goroutine就登记+1
		go handleGoroutine(ctx, &wg1, i+1, v, tmpProxy, &binanceTradeHistoryByPage)
	}

	wg1.Wait()

	for i := int64(0); i < requestMaxPage; i++ {
		tmpBinanceTradeHistoryByPage, ok := binanceTradeHistoryByPage.Load(i)
		fmt.Println("页数:", i)
		fmt.Println(tmpBinanceTradeHistoryByPage, ok)
		//for _, vVBinanceTradeHistoryByPage := range tmpBinanceTradeHistoryByPage {
		//	fmt.Println(vVBinanceTradeHistoryByPage)
		//}
	}

	//for page, vBinanceTradeHistoryByPage := range binanceTradeHistoryByPage {
	//	fmt.Println("页数:", page)
	//	for _, vVBinanceTradeHistoryByPage := range vBinanceTradeHistoryByPage {
	//		fmt.Println(vVBinanceTradeHistoryByPage)
	//	}
	//}

	//insertBinanceTrade := make([]*BinanceTradeHistory, 0)

	//// 不断爬数据，一直请求，直到和数据表中数据最后一条一致，预估一次有个几百条，每页50条，每两秒一次，期间可能有新数据，在本次可能被忽略
	//last := true
	//tmpPageNumber := int64(1)
	//for last {
	//	var (
	//		binanceTradeHistory []*BinanceTradeHistoryDataList
	//	)
	//
	//	for i := proxyKMin; i <= proxyKMax; i++ {
	//		tmpProxy := "http://" + proxy[i].Ip + ":" + strconv.FormatInt(proxy[i].Port, 10) + "/"
	//		binanceTradeHistory, err = requestProxyBinanceTradeHistory(tmpProxy, tmpPageNumber, 50, int64(v.TraderNum))
	//		if nil != err {
	//			fmt.Println(err)
	//			continue
	//		}
	//
	//		break
	//	}
	//
	//	//binanceTradeHistory, err = requestBinanceTradeHistory(tmpPageNumber, 50, int64(v.TraderNum))
	//	//if nil != err {
	//	//	time.Sleep(2 * time.Second)
	//	//	fmt.Println(err)
	//	//	continue
	//	//}
	//
	//	if nil == binanceTradeHistory {
	//		return // 直接返回
	//	}
	//	if 0 >= len(binanceTradeHistory) {
	//		return // 直接返回
	//	}
	//
	//	//currentLen := len(binanceTradeHistory)
	//	//for kPage, vBinanceTradeHistory := range binanceTradeHistory {
	//	//	// 第1-n条一致认为是已有数据
	//	//
	//	//	if 0 == currentCompareMax { // 数据库无数据
	//	//		// 不做限制
	//	//	} else { // 有数据且不足对比条数
	//	//
	//	//		// 页面上的数据一定大于等于数据库中对比条数
	//	//		for kDatabase, binanceTradeHistoryNewest := range binanceTradeHistoryNewestGroup {
	//	//			if vBinanceTradeHistory.Time == binanceTradeHistoryNewest.Time &&
	//	//				vBinanceTradeHistory.Symbol == binanceTradeHistoryNewest.Symbol &&
	//	//				vBinanceTradeHistory.Side == binanceTradeHistoryNewest.Side &&
	//	//				vBinanceTradeHistory.PositionSide == binanceTradeHistoryNewest.PositionSide &&
	//	//				IsEqual(vBinanceTradeHistory.Qty, binanceTradeHistoryNewest.Qty) && // 数量
	//	//				IsEqual(vBinanceTradeHistory.Price, binanceTradeHistoryNewest.Price) && //价格
	//	//				IsEqual(vBinanceTradeHistory.RealizedProfit, binanceTradeHistoryNewest.RealizedProfit) &&
	//	//				IsEqual(vBinanceTradeHistory.Quantity, binanceTradeHistoryNewest.Quantity) &&
	//	//				IsEqual(vBinanceTradeHistory.Fee, binanceTradeHistoryNewest.Fee) {
	//	//
	//	//			}
	//	//		}
	//	//
	//	//		time.Sleep(2 * time.Second)
	//	//		last = false // 终止
	//	//		break
	//	//	}
	//	//
	//	//	// 类型处理
	//	//	tmpActiveBuy := "false"
	//	//	if vBinanceTradeHistory.ActiveBuy {
	//	//		tmpActiveBuy = "true"
	//	//	}
	//	//
	//	//	// 追加
	//	//	insertBinanceTrade = append(insertBinanceTrade, &BinanceTradeHistory{
	//	//		Time:                vBinanceTradeHistory.Time,
	//	//		Symbol:              vBinanceTradeHistory.Symbol,
	//	//		Side:                vBinanceTradeHistory.Side,
	//	//		Price:               vBinanceTradeHistory.Price,
	//	//		Fee:                 vBinanceTradeHistory.Fee,
	//	//		FeeAsset:            vBinanceTradeHistory.FeeAsset,
	//	//		Quantity:            vBinanceTradeHistory.Quantity,
	//	//		QuantityAsset:       vBinanceTradeHistory.QuantityAsset,
	//	//		RealizedProfit:      vBinanceTradeHistory.RealizedProfit,
	//	//		RealizedProfitAsset: vBinanceTradeHistory.RealizedProfitAsset,
	//	//		BaseAsset:           vBinanceTradeHistory.BaseAsset,
	//	//		Qty:                 vBinanceTradeHistory.Qty,
	//	//		PositionSide:        vBinanceTradeHistory.PositionSide,
	//	//		ActiveBuy:           tmpActiveBuy,
	//	//	})
	//	//}
	//	//
	//	//// 不满50条，查到了最后，一般出现在大于50条的初始化
	//	//if 50 > len(binanceTradeHistory) {
	//	//	break
	//	//}
	//	//
	//	//tmpPageNumber++
	//	//time.Sleep(2 * time.Second)
	//}
	//
	//// insert
	//insertBinanceTradeLen := len(insertBinanceTrade)
	//if 0 < insertBinanceTradeLen {
	//	for i := insertBinanceTradeLen - 1; i >= 0; i-- {
	//		_, err = b.binanceUserRepo.InsertBinanceTradeHistory(ctx, insertBinanceTrade[i], v.TraderNum)
	//		if nil != err {
	//			fmt.Println(err)
	//		}
	//	}
	//}
}

func handleGoroutine(ctx context.Context, wg1 *sync.WaitGroup, tmpPageNumber int64, v *BinanceTrade, proxy []*ProxyData, data *sync.Map) {
	defer wg1.Done() // goroutine结束就登记-1

	//if _, ok := data[tmpPageNumber]; !ok {
	//	data[tmpPageNumber] = make([]*BinanceTradeHistoryDataList, 0)
	//}

	var (
		err                 error
		binanceTradeHistory []*BinanceTradeHistoryDataList
	)

	for _, vProxy := range proxy {
		//tmpProxy := "http://" + vProxy.Ip + ":" + strconv.FormatInt(vProxy.Port, 10) + "/"
		//fmt.Println(tmpProxy, "页数:", tmpPageNumber)
		//binanceTradeHistory, err = requestProxyBinanceTradeHistory(tmpProxy, tmpPageNumber, 50, int64(v.TraderNum))
		//if nil != err {
		//	fmt.Println(err)
		//	continue
		//}

		fmt.Println(vProxy)
		binanceTradeHistory, err = requestBinanceTradeHistory(tmpPageNumber, 50, int64(v.TraderNum))
		if nil != err {
			time.Sleep(2 * time.Second)
			fmt.Println(err)
			continue
		}

		break
	}

	//binanceTradeHistory, err = requestBinanceTradeHistory(tmpPageNumber, 50, int64(v.TraderNum))
	//if nil != err {
	//	time.Sleep(2 * time.Second)
	//	fmt.Println(err)
	//	continue
	//}

	if nil == binanceTradeHistory {
		return // 直接返回
	}
	if 0 >= len(binanceTradeHistory) {
		return // 直接返回
	}

	data.Store(tmpPageNumber, binanceTradeHistory)
	return

	//currentLen := len(binanceTradeHistory)
	//for kPage, vBinanceTradeHistory := range binanceTradeHistory {
	//	// 第1-n条一致认为是已有数据
	//
	//	if 0 == currentCompareMax { // 数据库无数据
	//		// 不做限制
	//	} else { // 有数据且不足对比条数
	//
	//		// 页面上的数据一定大于等于数据库中对比条数
	//		for kDatabase, binanceTradeHistoryNewest := range binanceTradeHistoryNewestGroup {
	//			if vBinanceTradeHistory.Time == binanceTradeHistoryNewest.Time &&
	//				vBinanceTradeHistory.Symbol == binanceTradeHistoryNewest.Symbol &&
	//				vBinanceTradeHistory.Side == binanceTradeHistoryNewest.Side &&
	//				vBinanceTradeHistory.PositionSide == binanceTradeHistoryNewest.PositionSide &&
	//				IsEqual(vBinanceTradeHistory.Qty, binanceTradeHistoryNewest.Qty) && // 数量
	//				IsEqual(vBinanceTradeHistory.Price, binanceTradeHistoryNewest.Price) && //价格
	//				IsEqual(vBinanceTradeHistory.RealizedProfit, binanceTradeHistoryNewest.RealizedProfit) &&
	//				IsEqual(vBinanceTradeHistory.Quantity, binanceTradeHistoryNewest.Quantity) &&
	//				IsEqual(vBinanceTradeHistory.Fee, binanceTradeHistoryNewest.Fee) {
	//
	//			}
	//		}
	//
	//		time.Sleep(2 * time.Second)
	//		last = false // 终止
	//		break
	//	}
	//
	//	// 类型处理
	//	tmpActiveBuy := "false"
	//	if vBinanceTradeHistory.ActiveBuy {
	//		tmpActiveBuy = "true"
	//	}
	//
	//	// 追加
	//	insertBinanceTrade = append(insertBinanceTrade, &BinanceTradeHistory{
	//		Time:                vBinanceTradeHistory.Time,
	//		Symbol:              vBinanceTradeHistory.Symbol,
	//		Side:                vBinanceTradeHistory.Side,
	//		Price:               vBinanceTradeHistory.Price,
	//		Fee:                 vBinanceTradeHistory.Fee,
	//		FeeAsset:            vBinanceTradeHistory.FeeAsset,
	//		Quantity:            vBinanceTradeHistory.Quantity,
	//		QuantityAsset:       vBinanceTradeHistory.QuantityAsset,
	//		RealizedProfit:      vBinanceTradeHistory.RealizedProfit,
	//		RealizedProfitAsset: vBinanceTradeHistory.RealizedProfitAsset,
	//		BaseAsset:           vBinanceTradeHistory.BaseAsset,
	//		Qty:                 vBinanceTradeHistory.Qty,
	//		PositionSide:        vBinanceTradeHistory.PositionSide,
	//		ActiveBuy:           tmpActiveBuy,
	//	})
	//}
	//
	//// 不满50条，查到了最后，一般出现在大于50条的初始化
	//if 50 > len(binanceTradeHistory) {
	//	break
	//}
	//
	//tmpPageNumber++
	//time.Sleep(2 * time.Second)
}

// PullBinancePositionHistory 暂时不用
func (b *BinanceUserUsecase) PullBinancePositionHistory(ctx context.Context) error {
	var (
		binanceTrade []*BinanceTrade
		proxy        []*Proxy
		err          error
	)

	binanceTrade, err = b.binanceUserRepo.GetBinanceTrade()
	if nil != err {
		return err
	}

	proxy, err = getProxy()
	if nil != err {
		return err
	}

	for _, v := range binanceTrade {
		var (
			binancePositionHistoryNewest *BinancePositionHistory
		)
		binancePositionHistoryNewest, err = b.binanceUserRepo.GetBinancePositionHistoryByTraderNumNewest(v.TraderNum)
		if nil != err {
			return err
		}

		insertBinancePosition := make([]*BinancePositionHistory, 0)

		last := true
		tmpPageNumber := int64(1)
		for last {
			var (
				binancePositionHistory []*BinancePositionHistoryDataList
			)

			for i := len(proxy) - 1; i >= 0; i-- {
				if i < len(proxy)-3 { // 最多3次
					break
				}

				tmpProxy := "http://" + proxy[i].Ip + ":" + strconv.FormatInt(proxy[i].Port, 10) + "/"
				binancePositionHistory, err = requestProxyBinancePositionHistory(tmpProxy, tmpPageNumber, 50, int64(v.TraderNum))
				if nil != err {
					fmt.Println(err)
					continue
				}

				break
			}

			//binanceTradeHistory, err = requestBinanceTradeHistory(tmpPageNumber, 50, int64(v.TraderNum))
			//if nil != err {
			//	fmt.Println(err)
			//	continue
			//}

			if nil == binancePositionHistory {
				break
			}

			for _, vBinancePositionHistory := range binancePositionHistory {

				if nil != binancePositionHistoryNewest {
					if vBinancePositionHistory.Symbol == binancePositionHistoryNewest.Symbol &&
						vBinancePositionHistory.Side == binancePositionHistoryNewest.Side &&
						vBinancePositionHistory.Opened == binancePositionHistoryNewest.Opened &&
						vBinancePositionHistory.Closed == binancePositionHistoryNewest.Closed &&
						vBinancePositionHistory.Status == binancePositionHistoryNewest.Status &&
						IsEqual(vBinancePositionHistory.AvgCost, binancePositionHistoryNewest.AvgCost) &&
						IsEqual(vBinancePositionHistory.ClosingPnl, binancePositionHistoryNewest.ClosingPnl) {
						last = false // 终止
						break
					}

				}

				insertBinancePosition = append(insertBinancePosition, &BinancePositionHistory{
					Symbol:          vBinancePositionHistory.Symbol,
					Side:            vBinancePositionHistory.Side,
					Closed:          vBinancePositionHistory.Closed,
					Opened:          vBinancePositionHistory.Opened,
					AvgCost:         vBinancePositionHistory.AvgCost,
					AvgClosePrice:   vBinancePositionHistory.AvgClosePrice,
					ClosingPnl:      vBinancePositionHistory.ClosingPnl,
					MaxOpenInterest: vBinancePositionHistory.MaxOpenInterest,
					ClosedVolume:    vBinancePositionHistory.ClosedVolume,
					Status:          vBinancePositionHistory.Status,
				})
			}

			// 不满50条
			if 50 > len(binancePositionHistory) {
				break
			}

			tmpPageNumber++
			time.Sleep(2 * time.Second)
		}

		// insert
		insertBinanceTradeLen := len(insertBinancePosition)
		if 0 < insertBinanceTradeLen {
			for i := insertBinanceTradeLen - 1; i >= 0; i-- {
				_, err = b.binanceUserRepo.InsertBinancePositionHistory(ctx, insertBinancePosition[i])
				if nil != err {
					fmt.Println(err)
				}
			}
		}

		time.Sleep(200 * time.Millisecond) // 200毫秒一个代单员，减小被封代理IP概率
	}

	return err
}

// GetBinancePositionHistory 暂时不用
func (b *BinanceUserUsecase) GetBinancePositionHistory(ctx context.Context, req *v1.GetBinanceTraderPositionHistoryRequest) (*v1.GetBinanceTraderPositionHistoryReply, error) {
	var (
		proxy []*Proxy
		err   error
	)

	proxy, err = getProxy()
	if nil != err {
		return &v1.GetBinanceTraderPositionHistoryReply{}, err
	}

	var (
		binancePositionHistoryAll []*BinancePositionHistoryDataList
	)

	last := true
	tmpPageNumber := int64(1)
	for last {
		var (
			binancePositionHistory []*BinancePositionHistoryDataList
		)

		for i := len(proxy) - 1; i >= 0; i-- {
			if i < len(proxy)-3 { // 最多3次直接返回
				return &v1.GetBinanceTraderPositionHistoryReply{}, err
			}

			tmpProxy := "http://" + proxy[i].Ip + ":" + strconv.FormatInt(proxy[i].Port, 10) + "/"
			binancePositionHistory, err = requestProxyBinancePositionHistory(tmpProxy, tmpPageNumber, 50, int64(req.Id))
			if nil != err {
				fmt.Println(err)
				continue
			}

			break
		}

		//binancePositionHistory, err = requestBinancePositionHistory(tmpPageNumber, 50, int64(req.Id))
		//if nil != err {
		//	fmt.Println(err)
		//	last = false
		//}

		if nil == binancePositionHistory {
			last = false
		}

		for _, vBinancePositionHistory := range binancePositionHistory {
			binancePositionHistoryAll = append(binancePositionHistoryAll, vBinancePositionHistory)
		}

		// 不满50条
		if 50 > len(binancePositionHistory) {
			last = false
		}

		tmpPageNumber++
		time.Sleep(2 * time.Second)
	}

	var (
		winTotal       uint64
		lostTotal      uint64
		winTotalValue  float64
		lostTotalValue float64
	)

	res := &v1.GetBinanceTraderPositionHistoryReply{}
	res.Positions = make([]*v1.GetBinanceTraderPositionHistoryReply_Position, 0)

	for _, vBinancePositionHistoryAll := range binancePositionHistoryAll {
		if "All Closed" != vBinancePositionHistoryAll.Status {
			continue
		}

		if 0.000000001 > vBinancePositionHistoryAll.ClosingPnl {
			lostTotal++
			lostTotalValue += math.Abs(vBinancePositionHistoryAll.ClosingPnl)
		} else {
			winTotal++
			winTotalValue += vBinancePositionHistoryAll.ClosingPnl
		}

		res.Positions = append(res.Positions, &v1.GetBinanceTraderPositionHistoryReply_Position{
			Symbol:     vBinancePositionHistoryAll.Symbol,
			Side:       vBinancePositionHistoryAll.Side,
			ClosingPnl: vBinancePositionHistoryAll.ClosingPnl,
		})
	}

	// 胜率
	if 0 == lostTotal {
		res.WinRate = 1
	} else {
		res.WinRate = float64(winTotal) / float64(lostTotal)
	}

	// 盈亏比
	var (
		lostValueAvg float64
		winValueAvg  float64
	)

	if 0 < winTotalValue && 0 < winTotal {
		winValueAvg = winTotalValue / float64(winTotal)
	}

	if 0 < lostTotalValue && 0 < lostTotal {
		lostValueAvg = lostTotalValue / float64(lostTotal)
	}

	if 0 < lostValueAvg && 0 < winValueAvg {
		res.WinLostRate = winValueAvg / lostValueAvg
	}

	// 收益率
	res.RevenueRate = res.WinRate * res.WinLostRate

	return res, err
}

func IsEqual(f1, f2 float64) bool {
	if f1 > f2 {
		return f1-f2 < 0.000000001
	} else {
		return f2-f1 < 0.000000001
	}
}

// 拉取交易员交易历史
func requestProxyBinanceTradeHistory(proxyAddr string, pageNumber int64, pageSize int64, portfolioId int64) ([]*BinanceTradeHistoryDataList, error) {
	var (
		resp   *http.Response
		res    []*BinanceTradeHistoryDataList
		b      []byte
		err    error
		apiUrl = "https://www.binance.com/bapi/futures/v1/friendly/future/copy-trade/lead-portfolio/trade-history"
	)

	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		log.Fatal(err)
	}
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}
	httpClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	// 构造请求
	contentType := "application/json"
	data := `{"pageNumber":` + strconv.FormatInt(pageNumber, 10) + `,"pageSize":` + strconv.FormatInt(pageSize, 10) + `,portfolioId:` + strconv.FormatInt(portfolioId, 10) + `}`
	resp, err = httpClient.Post(apiUrl, contentType, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var l *BinanceTradeHistoryResp
	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if nil == l.Data {
		return res, nil
	}

	if nil == l.Data.List {
		return res, nil
	}

	res = make([]*BinanceTradeHistoryDataList, 0)
	for _, v := range l.Data.List {
		res = append(res, v)
	}

	return res, nil
}

// 拉取交易员交易历史
func requestBinanceTradeHistory(pageNumber int64, pageSize int64, portfolioId int64) ([]*BinanceTradeHistoryDataList, error) {
	var (
		resp   *http.Response
		res    []*BinanceTradeHistoryDataList
		b      []byte
		err    error
		apiUrl = "https://www.binance.com/bapi/futures/v1/friendly/future/copy-trade/lead-portfolio/trade-history"
	)

	// 构造请求
	contentType := "application/json"
	data := `{"pageNumber":` + strconv.FormatInt(pageNumber, 10) + `,"pageSize":` + strconv.FormatInt(pageSize, 10) + `,portfolioId:` + strconv.FormatInt(portfolioId, 10) + `}`
	resp, err = http.Post(apiUrl, contentType, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var l *BinanceTradeHistoryResp
	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if nil == l.Data {
		return res, nil
	}

	if nil == l.Data.List {
		return res, nil
	}

	res = make([]*BinanceTradeHistoryDataList, 0)
	for _, v := range l.Data.List {
		res = append(res, v)
	}

	return res, nil
}

type ProxyData struct {
	Ip   string
	Port int64
}

type ProxyRep struct {
	Data []*ProxyData
}

// 拉取代理列表
func requestProxy() ([]*ProxyData, error) {
	var (
		resp   *http.Response
		b      []byte
		res    []*ProxyData
		err    error
		apiUrl = "http://api.ipipgo.com/ip?cty=00&c=500&pt=1&ft=json&pat=\\n&rep=1&key=a37f63e3&ts=3"
	)

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	// 构造请求
	resp, err = httpClient.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var l *ProxyRep
	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res = make([]*ProxyData, 0)
	if nil == l.Data || 0 >= len(l.Data) {
		return res, nil
	}

	for _, v := range l.Data {
		res = append(res, v)
	}

	return res, nil
}

// 拉取交易员仓位历史
func requestProxyBinancePositionHistory(proxyAddr string, pageNumber int64, pageSize int64, portfolioId int64) ([]*BinancePositionHistoryDataList, error) {
	var (
		resp   *http.Response
		res    []*BinancePositionHistoryDataList
		b      []byte
		err    error
		apiUrl = "https://www.binance.com/bapi/futures/v1/friendly/future/copy-trade/lead-portfolio/position-history"
	)

	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		log.Fatal(err)
	}
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}
	httpClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	// 构造请求
	contentType := "application/json"
	data := `{"pageNumber":` + strconv.FormatInt(pageNumber, 10) + `,"pageSize":` + strconv.FormatInt(pageSize, 10) + `,portfolioId:` + strconv.FormatInt(portfolioId, 10) + `}`
	resp, err = httpClient.Post(apiUrl, contentType, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var l *BinancePositionHistoryResp
	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if nil == l.Data {
		return res, nil
	}

	if nil == l.Data.List {
		return res, nil
	}

	res = make([]*BinancePositionHistoryDataList, 0)
	for _, v := range l.Data.List {
		res = append(res, v)
	}

	return res, nil
}

// 拉取交易员交易历史
func requestBinancePositionHistory(pageNumber int64, pageSize int64, portfolioId int64) ([]*BinancePositionHistoryDataList, error) {
	var (
		resp   *http.Response
		res    []*BinancePositionHistoryDataList
		b      []byte
		err    error
		apiUrl = "https://www.binance.com/bapi/futures/v1/friendly/future/copy-trade/lead-portfolio/position-history"
	)

	// 构造请求
	contentType := "application/json"
	data := `{"pageNumber":` + strconv.FormatInt(pageNumber, 10) + `,"pageSize":` + strconv.FormatInt(pageSize, 10) + `,portfolioId:` + strconv.FormatInt(portfolioId, 10) + `}`
	resp, err = http.Post(apiUrl, contentType, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var l *BinancePositionHistoryResp
	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if nil == l.Data {
		return res, nil
	}

	if nil == l.Data.List {
		return res, nil
	}

	res = make([]*BinancePositionHistoryDataList, 0)
	for _, v := range l.Data.List {
		res = append(res, v)
	}

	return res, nil
}

type Proxy struct {
	Ip   string
	Port int64
}

func getProxy() ([]*Proxy, error) {
	var (
		p   []*Proxy
		err error
	)

	content, err := ioutil.ReadFile("/www/wwwroot/copyuserdata.trading-fi.com/data/proxy_ip.txt")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = json.Unmarshal(content, &p)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return p, nil
}

// UserOrderDo 新增刷单
func (b *BinanceUserUsecase) UserOrderDo(ctx context.Context, req *v1.UserOrderDoRequest) (*v1.UserOrderDoReply, error) {
	var (
		err error
		//usdtBtc       = req.Amount
		//priceBtc      *BinanceSymbolPrice
		priceFloatEth    float64
		priceFloatEthTwo float64
		//priceFloatBtc float64
		qtyEth      float64
		qtyEthTwo   float64
		priceEth    *BinanceSymbolPrice
		priceEthTwo *BinanceSymbolPrice
		//qtyBtc        float64
		symbol map[string]*Symbol
	)
	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return nil, err
	}

	// 精度
	if _, ok := symbol[req.Symbol]; !ok {
		return nil, errors.New(500, "SYMBOL ERR", "不存在查询代币精度")
	}
	if _, ok := symbol[req.SymbolTwo]; !ok {
		return nil, errors.New(500, "SYMBOL TWO ERR", "不存在查询代币精度，第二个")
	}

	// 下单
	// 精度调整
	var (
		quantityEth             string
		quantityEthTwo          string
		binanceOrderEth         *BinanceOrder
		binanceOrderEthTwo      *BinanceOrder
		binanceOrderEthClose    *BinanceOrder
		binanceOrderEthTwoClose *BinanceOrder
	)

	if 1 == req.Redo {
		priceFloatEth = req.PriceOpen
		priceFloatEthTwo = req.PriceOpenTwo
		var userOrderDo []*UserOrderDo
		userOrderDo, err = b.binanceUserRepo.GetUserOrderDoLast()
		if nil != err {
			return nil, err
		}

		if 0 >= len(userOrderDo) {
			return nil, errors.New(500, "Order Err", "没有订单")
		}

		// 查询币价 第一个
		qtyEth = req.Amount / priceFloatEth
		if 0 >= symbol[req.Symbol].QuantityPrecision {
			quantityEth = fmt.Sprintf("%d", int64(qtyEth))
		} else {
			quantityEth = strconv.FormatFloat(qtyEth, 'f', int(symbol[req.Symbol].QuantityPrecision), 64)
		}

		// 查询币价 第二个
		qtyEthTwo = req.AmountTwo / priceFloatEthTwo
		if 0 >= symbol[req.SymbolTwo].QuantityPrecision {
			quantityEthTwo = fmt.Sprintf("%d", int64(qtyEthTwo))
		} else {
			quantityEthTwo = strconv.FormatFloat(qtyEthTwo, 'f', int(symbol[req.SymbolTwo].QuantityPrecision), 64)
		}

		var (
			positionSide         string // todo 变量名字反了 side对应的值是sell,long
			positionSideClose    string // todo 变量名字反了 side对应的值是sell,long
			limitPrice           string
			positionSideTwo      string
			positionSideTwoClose string
			limitPriceTwo        string
		)

		if "SHORT" == req.Side { // 开空
			positionSide = "SELL"
			positionSideClose = "BUY"
		} else if "LONG" == req.Side { // 开多
			positionSide = "BUY"
			positionSideClose = "SELL"
		} else {
			return nil, errors.New(500, "Side Err", "下单方向错误")
		}

		if "SHORT" == req.SideTwo { // 开空
			positionSideTwo = "SELL"
			positionSideTwoClose = "BUY"
		} else if "LONG" == req.SideTwo { // 开多
			positionSideTwo = "BUY"
			positionSideTwoClose = "SELL"
		} else {
			return nil, errors.New(500, "Side Err", "下单方向错误")
		}

		if 1 == req.RedoNum {
			binanceOrderEth, _, err = requestBinanceOrder(req.Symbol, positionSide, "MARKET", req.Side, quantityEth, req.ApiKey, req.ApiSecret)
			if nil != err {
				return nil, err
			}
			if 0 >= binanceOrderEth.OrderId {
				return nil, errors.New(500, "Order Err", "开仓错误")
			}

			var (
				qty float64
			)
			qty, err = strconv.ParseFloat(binanceOrderEth.ExecutedQty, 64)
			if nil != err {
				fmt.Println(err, "qty", binanceOrderEth)
			}

			_, err = b.binanceUserRepo.UpdateUserOrderDoThree(ctx, userOrderDo[0].ID, qty, 0)
			if nil != err {
				return nil, err
			}
		} else if 2 == req.RedoNum {
			binanceOrderEthTwo, _, err = requestBinanceOrder(req.SymbolTwo, positionSideTwo, "MARKET", req.SideTwo, quantityEthTwo, req.ApiKeyTwo, req.ApiSecretTwo)
			if nil != err {
				return nil, err
			}

			if 0 >= binanceOrderEthTwo.OrderId {
				return nil, errors.New(500, "Order Err", "开仓错误")
			}

			var (
				qtyTwo float64
			)
			qtyTwo, err = strconv.ParseFloat(binanceOrderEthTwo.ExecutedQty, 64)
			if nil != err {
				fmt.Println(err, "two qty", binanceOrderEthTwo)
			}

			_, err = b.binanceUserRepo.UpdateUserOrderDoThree(ctx, userOrderDo[0].ID, 0, qtyTwo)
			if nil != err {
				return nil, err
			}
		} else if 3 == req.RedoNum {
			limitPrice = strconv.FormatFloat(req.Price, 'f', int(symbol[req.Symbol].PricePrecision), 64)
			binanceOrderEthClose, _, err = requestBinanceOrderStop(req.Symbol, positionSideClose, req.Side, quantityEth, limitPrice, limitPrice, req.ApiKey, req.ApiSecret)
			if nil != err {
				return nil, err
			}

			if 0 >= binanceOrderEthClose.OrderId {
				return nil, errors.New(500, "Order Err", "补单错误")
			}

			_, err = b.binanceUserRepo.UpdateUserOrderDoTwo(ctx, userOrderDo[0].ID, strconv.FormatInt(binanceOrderEthClose.OrderId, 10), "")
			if nil != err {
				return nil, err
			}
		} else if 4 == req.RedoNum {
			limitPriceTwo = strconv.FormatFloat(req.PriceTwo, 'f', int(symbol[req.SymbolTwo].PricePrecision), 64)
			binanceOrderEthTwoClose, _, err = requestBinanceOrderStop(req.SymbolTwo, positionSideTwoClose, req.SideTwo, quantityEthTwo, limitPriceTwo, limitPriceTwo, req.ApiKeyTwo, req.ApiSecretTwo)
			if nil != err {
				return nil, err
			}

			if 0 >= binanceOrderEthTwoClose.OrderId {
				return nil, errors.New(500, "Order Err", "补单错误")
			}

			_, err = b.binanceUserRepo.UpdateUserOrderDoTwo(ctx, userOrderDo[0].ID, "", strconv.FormatInt(binanceOrderEthTwoClose.OrderId, 10))
			if nil != err {
				return nil, err
			}
		} else if 5 == req.RedoNum {
			limitPrice = strconv.FormatFloat(req.Price, 'f', int(symbol[req.Symbol].PricePrecision), 64)
			binanceOrderEthClose, _, err = requestBinanceOrderStop(req.Symbol, positionSideClose, req.Side, quantityEth, limitPrice, limitPrice, req.ApiKey, req.ApiSecret)
			if nil != err {
				return nil, err
			}

			if 0 >= binanceOrderEthClose.OrderId {
				return nil, errors.New(500, "Order Err", "补单错误")
			}

			_, err = b.binanceUserRepo.UpdateUserOrderDoTwo(ctx, userOrderDo[0].ID, strconv.FormatInt(binanceOrderEthClose.OrderId, 10), "")
			if nil != err {
				return nil, err
			}

			limitPriceTwo = strconv.FormatFloat(req.PriceTwo, 'f', int(symbol[req.SymbolTwo].PricePrecision), 64)
			binanceOrderEthTwoClose, _, err = requestBinanceOrderStop(req.SymbolTwo, positionSideTwoClose, req.SideTwo, quantityEthTwo, limitPriceTwo, limitPriceTwo, req.ApiKeyTwo, req.ApiSecretTwo)
			if nil != err {
				return nil, err
			}

			if 0 >= binanceOrderEthTwoClose.OrderId {
				return nil, errors.New(500, "Order Err", "补单错误")
			}

			_, err = b.binanceUserRepo.UpdateUserOrderDoTwo(ctx, userOrderDo[0].ID, "", strconv.FormatInt(binanceOrderEthTwoClose.OrderId, 10))
			if nil != err {
				return nil, err
			}
		}

	} else {
		_, err = requestBinanceLeverAge(req.Symbol, req.Num, req.ApiKey, req.ApiSecret)
		if nil != err {
			return nil, err
		}

		_, err = requestBinanceLeverAge(req.SymbolTwo, req.NumTwo, req.ApiKeyTwo, req.ApiSecretTwo)
		if nil != err {
			return nil, err
		}

		// 查询币价 第一个
		priceEth, err = requestBinanceSymbolPrice(req.Symbol)
		if nil != err {
			return nil, err
		}
		priceFloatEth, err = strconv.ParseFloat(priceEth.Price, 64)
		if nil != err {
			return nil, err
		}

		qtyEth = req.Amount / priceFloatEth
		if 0 >= symbol[req.Symbol].QuantityPrecision {
			quantityEth = fmt.Sprintf("%d", int64(qtyEth))
		} else {
			quantityEth = strconv.FormatFloat(qtyEth, 'f', int(symbol[req.Symbol].QuantityPrecision), 64)
		}

		// 查询币价 第二个
		priceEthTwo, err = requestBinanceSymbolPrice(req.SymbolTwo)
		if nil != err {
			return nil, err
		}

		priceFloatEthTwo, err = strconv.ParseFloat(priceEthTwo.Price, 64)
		if nil != err {
			return nil, err
		}
		qtyEthTwo = req.AmountTwo / priceFloatEthTwo
		if 0 >= symbol[req.SymbolTwo].QuantityPrecision {
			quantityEthTwo = fmt.Sprintf("%d", int64(qtyEthTwo))
		} else {
			quantityEthTwo = strconv.FormatFloat(qtyEthTwo, 'f', int(symbol[req.SymbolTwo].QuantityPrecision), 64)
		}

		var (
			positionSide string // todo 变量名字反了 side对应的值是sell,long
			//positionSideClose    string // todo 变量名字反了 side对应的值是sell,long
			//limitPrice           string
			positionSideTwo string
			//positionSideTwoClose string
			//limitPriceTwo        string
		)

		if "SHORT" == req.Side { // 开空
			positionSide = "SELL"
			//positionSideClose = "BUY"
		} else if "LONG" == req.Side { // 开多
			positionSide = "BUY"
			//positionSideClose = "SELL"
		} else {
			return nil, errors.New(500, "Side Err", "下单方向错误")
		}

		if "SHORT" == req.SideTwo { // 开空
			positionSideTwo = "SELL"
			//positionSideTwoClose = "BUY"
		} else if "LONG" == req.SideTwo { // 开多
			positionSideTwo = "BUY"
			//positionSideTwoClose = "SELL"
		} else {
			return nil, errors.New(500, "Side Err", "下单方向错误")
		}

		binanceOrderEth, _, err = requestBinanceOrder(req.Symbol, positionSide, "MARKET", req.Side, quantityEth, req.ApiKey, req.ApiSecret)
		if nil != err {
			return nil, err
		}

		binanceOrderEthTwo, _, err = requestBinanceOrder(req.SymbolTwo, positionSideTwo, "MARKET", req.SideTwo, quantityEthTwo, req.ApiKeyTwo, req.ApiSecretTwo)
		if nil != err {
			return nil, err
		}

		if 0 >= binanceOrderEth.OrderId && 0 >= binanceOrderEthTwo.OrderId {
			return nil, errors.New(500, "Order Err", "两单开仓错误")
		}

		fmt.Println(binanceOrderEth, binanceOrderEthTwo)

		//priceFloatEth, err = strconv.ParseFloat(binanceOrderEth.AvgPrice, 64)
		//if nil != err {
		//	fmt.Println(err, "qty", binanceOrderEth)
		//	return nil, errors.New(500, "Order Err", "下单错误，第2没下成功")
		//}
		//
		//priceFloatEthTwo, err = strconv.ParseFloat(binanceOrderEthTwo.AvgPrice, 64)
		//if nil != err {
		//	fmt.Println(err, "two qty", binanceOrderEthTwo)
		//	return nil, errors.New(500, "Order Err", "下单错误，第2没下成功")
		//}

		//// 多空的止损价
		//qtyEthLimitLong := priceFloatEth - (priceFloatEth*(req.Amount/float64(req.Num)-req.CloseAmount))/req.Amount
		//qtyEthLimitShort := priceFloatEth + (priceFloatEth*(req.Amount/float64(req.Num)-req.CloseAmount))/req.Amount
		//
		//// 多空的止损价
		//qtyEthLimitLongTwo := priceFloatEthTwo - (priceFloatEthTwo*(req.AmountTwo/float64(req.NumTwo)-req.CloseAmountTwo))/req.AmountTwo
		//qtyEthLimitShortTwo := priceFloatEthTwo + (priceFloatEthTwo*(req.AmountTwo/float64(req.NumTwo)-req.CloseAmountTwo))/req.AmountTwo
		//
		//fmt.Println(priceFloatEthTwo, qtyEthLimitShortTwo, req.AmountTwo/float64(req.NumTwo)-req.CloseAmountTwo, req.AmountTwo)
		//if "SHORT" == req.Side { // 开空
		//	limitPrice = strconv.FormatFloat(qtyEthLimitShort, 'f', int(symbol[req.Symbol].PricePrecision), 64)
		//} else if "LONG" == req.Side { // 开多
		//	limitPrice = strconv.FormatFloat(qtyEthLimitLong, 'f', int(symbol[req.Symbol].PricePrecision), 64)
		//} else {
		//	return nil, errors.New(500, "Side Err", "下单方向错误")
		//}
		//
		//if "SHORT" == req.SideTwo { // 开空
		//	limitPriceTwo = strconv.FormatFloat(qtyEthLimitShortTwo, 'f', int(symbol[req.SymbolTwo].PricePrecision), 64)
		//} else if "LONG" == req.SideTwo { // 开多
		//	limitPriceTwo = strconv.FormatFloat(qtyEthLimitLongTwo, 'f', int(symbol[req.SymbolTwo].PricePrecision), 64)
		//} else {
		//	return nil, errors.New(500, "Side Err", "下单方向错误，第二单")
		//}

		//limitPrice = strconv.FormatFloat(req.Price, 'f', int(symbol[req.Symbol].PricePrecision), 64)
		//binanceOrderEthClose, _, err = requestBinanceOrderStop(req.Symbol, positionSideClose, req.Side, quantityEth, limitPrice, limitPrice, req.ApiKey, req.ApiSecret)
		//if nil != err {
		//	return nil, err
		//}
		//
		//limitPriceTwo = strconv.FormatFloat(req.PriceTwo, 'f', int(symbol[req.SymbolTwo].PricePrecision), 64)
		//binanceOrderEthTwoClose, _, err = requestBinanceOrderStop(req.SymbolTwo, positionSideTwoClose, req.SideTwo, quantityEthTwo, limitPriceTwo, limitPriceTwo, req.ApiKeyTwo, req.ApiSecretTwo)
		//if nil != err {
		//	return nil, err
		//}

		var (
			userOrderDoEth *UserOrderDoNew
			qty            float64
			qtyTwo         float64
		)
		qty, err = strconv.ParseFloat(binanceOrderEth.ExecutedQty, 64)
		if nil != err {
			fmt.Println(err, "qty", binanceOrderEth)
		}

		qtyTwo, err = strconv.ParseFloat(binanceOrderEthTwo.ExecutedQty, 64)
		if nil != err {
			fmt.Println(err, "two qty", binanceOrderEthTwo)
		}

		userOrderDoEth = &UserOrderDoNew{
			ApiKey:       req.ApiKey,
			ApiSecret:    req.ApiSecret,
			ApiKeyTwo:    req.ApiKeyTwo,
			ApiSecretTwo: req.ApiSecretTwo,
			Symbol:       req.Symbol,
			SymbolTwo:    req.SymbolTwo,
			Status:       1,
			Qty:          qty,
			QtyTwo:       qtyTwo,
			SideTwo:      req.SideTwo,
			Side:         req.Side,
			//OrderId:      strconv.FormatInt(binanceOrderEthClose.OrderId, 10),
			//OrderIdTwo:   strconv.FormatInt(binanceOrderEthTwoClose.OrderId, 10),
		}

		err = b.binanceUserRepo.InsertUserOrderDoNew(ctx, userOrderDoEth)
		if nil != err {
			return nil, err
		}

		if 0 >= binanceOrderEth.OrderId {
			return nil, errors.New(500, "Order Err", "下单错误，第1单没下成功")
		}

		if 0 >= binanceOrderEthTwo.OrderId {
			return nil, errors.New(500, "Order Err", "下单错误，第2单没下成功")
		}

		//if 0 >= binanceOrderEthClose.OrderId {
		//	return nil, errors.New(500, "Order Err", "下单错误，第1止损没下成功")
		//}
		//
		//if 0 >= binanceOrderEthTwoClose.OrderId {
		//	return nil, errors.New(500, "Order Err", "下单错误，第2止损没下成功")
		//}
	}

	return nil, nil
}

// UserOrderDoHandlePrice 刷单处理
func (b *BinanceUserUsecase) UserOrderDoHandlePrice(ctx context.Context, req *v1.UserOrderDoHandlePriceRequest) (*v1.UserOrderDoHandlePriceReply, error) {
	var (
		err         error
		userOrderDo []*UserOrderDoNew
		symbol      map[string]*Symbol
	)
	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return nil, err
	}

	//// 精度
	//if _, ok := symbol["EOSUSDT"]; !ok {
	//	return nil, errors.New(500, "EOSUSDT SYMBOL ERR", "不存在查询代币精度")
	//}
	//if _, ok := symbol["BTCUSDT"]; !ok {
	//	return nil, errors.New(500, "BTCUSDT SYMBOL ERR", "不存在查询代币精度")
	//}

	// 查询数据
	userOrderDo, err = b.binanceUserRepo.GetUserOrderDoNew()
	if nil != err {
		return nil, err
	}

	if nil == userOrderDo || 0 >= len(userOrderDo) {
		return nil, nil
	}

	for _, vUserOrderDo := range userOrderDo {
		var (
			orderId           int64
			orderIdTwo        int64
			closeSymbol       string
			qty               float64
			qtyStr            string
			positionSide      string
			side              string
			apiKey            string
			apiSecret         string
			orderInfo         *BinanceOrder
			orderInfoTwo      *BinanceOrder
			binanceOrderClose *BinanceOrder
			closeOne          bool
		)

		// 检测哪个仓位平了
		if 0 < len(vUserOrderDo.OrderIdTwo) {
			orderId, err = strconv.ParseInt(vUserOrderDo.OrderId, 10, 64)
			if 0 < orderId {
				orderInfo, err = requestBinanceOrderInfo(vUserOrderDo.Symbol, orderId, vUserOrderDo.ApiKey, vUserOrderDo.ApiSecret)
				if nil != err || nil == orderInfo {
					fmt.Println(vUserOrderDo, "刷单查询，查询，错误信息", orderInfo, err)
					continue
				}

				if 0 < len(orderInfo.Status) && "NEW" != orderInfo.Status {
					closeSymbol = vUserOrderDo.SymbolTwo
					positionSide = vUserOrderDo.SideTwo
					if "LONG" == positionSide {
						side = "SELL"
					} else {
						side = "BUY"
					}
					qty = vUserOrderDo.QtyTwo
					apiKey = vUserOrderDo.ApiKeyTwo
					apiSecret = vUserOrderDo.ApiSecretTwo
					closeOne = true
				}
				//fmt.Println(orderInfo)
			}
		}

		if 0 < len(vUserOrderDo.OrderIdTwo) {
			orderIdTwo, err = strconv.ParseInt(vUserOrderDo.OrderIdTwo, 10, 64)
			if 0 < orderIdTwo {
				orderInfoTwo, err = requestBinanceOrderInfo(vUserOrderDo.SymbolTwo, orderIdTwo, vUserOrderDo.ApiKeyTwo, vUserOrderDo.ApiSecretTwo)
				if nil != err || nil == orderInfoTwo {
					fmt.Println(vUserOrderDo, "刷单查询，查询，错误信息", orderInfo, err)
					continue
				}

				if 0 < len(orderInfoTwo.Status) && "NEW" != orderInfoTwo.Status {
					closeSymbol = vUserOrderDo.Symbol
					positionSide = vUserOrderDo.Side
					if "LONG" == positionSide {
						side = "SELL"
					} else {
						side = "BUY"
					}
					qty = vUserOrderDo.Qty
					apiKey = vUserOrderDo.ApiKey
					apiSecret = vUserOrderDo.ApiSecret
					closeOne = true
				}

				//fmt.Println(orderInfoTwo)
			}
		}

		//fmt.Println(orderInfo, orderInfoTwo)
		if closeOne {
			fmt.Println(orderInfo, orderInfoTwo)
			// 关仓
			qtyStr = strconv.FormatFloat(qty, 'f', int(symbol[closeSymbol].QuantityPrecision), 64)
			//fmt.Println(closeSymbol, side, "MARKET", positionSide, qtyStr, apiKey, apiSecret, binanceOrderClose)

			if 0 >= qty {
				continue
			}

			binanceOrderClose, _, err = requestBinanceOrder(closeSymbol, side, "MARKET", positionSide, qtyStr, apiKey, apiSecret)
			if nil != err {
				return nil, err
			}

			if 0 >= binanceOrderClose.OrderId {
				fmt.Println(binanceOrderClose)
				return nil, errors.New(500, "Order Err", "关仓下单错误")
			}

			_, err = b.binanceUserRepo.UpdateUserOrderDo(ctx, vUserOrderDo.ID, strconv.FormatInt(binanceOrderClose.OrderId, 10))
			if nil != err {
				return nil, err
			}
		}
	}

	return nil, nil
}

// PullFilData .
func (b *BinanceUserUsecase) PullFilData(ctx context.Context, req *v1.PullFilDataRequest) (*v1.PullFilDataReply, error) {
	type AddressTypeAndData struct {
		AddressType uint64
		FilDataList []*FilDataList
	}

	var (
		err error
		// 全局人send记录
		globalData map[string]*AddressTypeAndData
	)
	globalData = make(map[string]*AddressTypeAndData, 0)

	var (
		getSendFilDataList func(address string, globalData map[string]*AddressTypeAndData, num uint64) // 定义递归方法, 获取send数据
	)
	getSendFilDataList = func(address string, globalData map[string]*AddressTypeAndData, num uint64) {
		fmt.Println("开始查询：", address)
		var (
			forCondition = true
			filDataList  []*FilDataList
		)

		if 1 >= len(address) {
			fmt.Println(address, "err address")
			return
		}

		// 存在的地址信息，不再查询
		if _, ok := globalData[address]; ok {
			return
		}

		// 初始化
		globalData[address] = &AddressTypeAndData{
			AddressType: 1, // 1普通账户，2交易所账户，3f0地址
		}
		globalData[address].FilDataList = make([]*FilDataList, 0)
		if "0" == string(address[1]) {
			globalData[address].AddressType = 3
			if 0 < num { // 非第一次节点查询，
				return
			}
		}

		filDataList = make([]*FilDataList, 0)
		last := time.Now().Add(-259200 * time.Second).Unix() // 三天前时间戳
		for i := int64(0); forCondition && i < 1000; i++ {
			var (
				total          uint64
				tmpFilDataList []*FilDataList
			)

			tmpFilDataList, total, err = requestFilTransfer(address, 100, i)
			if nil != err {
				fmt.Println(err)
				break
			}

			if 0 < num { // 非首次入参查询
				// 交易所账户，时间线，3天以内仍在转账
				if 0 < len(tmpFilDataList) && uint64(last) < tmpFilDataList[0].Timestamp {
					globalData[address].AddressType = 2
					break
				}

				// 交易所账户，总转账在1000以上
				if 1000 < total {
					globalData[address].AddressType = 2
					break
				}
			}

			// 遍历每一条
			for _, v := range tmpFilDataList {
				// 时间线以后的，直接返回
				if v.Timestamp < 1646582400 {
					// 结束循环，最外层查询
					forCondition = false
					break
				}

				// 只要发送订单
				if "send" != v.Type {
					continue
				}

				if v.From == v.To {
					continue
				}

				// 最少1个fil
				if 19 > len(v.Value) {
					continue
				}

				// 追加
				filDataList = append(filDataList, v)
			}

			// 不足100条，结束了
			if 100 > len(tmpFilDataList) {
				break
			}
		}

		// 跟新全局信息
		globalData[address].FilDataList = filDataList

		// 递归
		if 0 < len(filDataList) {
			for _, vFilDataList := range filDataList {
				getSendFilDataList(vFilDataList.To, globalData, 1)
			}
		}
	}

	getSendFilDataList(req.Address, globalData, 0)

	//var (
	//	filDataList []*FilDataList
	//)

	//for _, v := range getSendFilDataList(req.Address, globalData) {
	//	fmt.Println(v)
	//}

	for _, vG := range globalData {
		for _, vVG := range vG.FilDataList {
			err = b.binanceUserRepo.InsertFilData(ctx, &FilData2{
				From:     vVG.From,
				FromType: vG.AddressType,
				To:       vVG.To,
				Value:    vVG.Value,
			})
			if nil != err {
				fmt.Println(err)
				continue
			}
		}
	}

	return nil, nil
}

// 获取fil交易历史
func requestFilTransfer(address string, pageSize int64, page int64) ([]*FilDataList, uint64, error) {
	var (
		resp *http.Response
		res  []*FilDataList
		b    []byte
		err  error
	)

	apiUrl := "https://filfox.info/api/v1/address/" + address + "/transfers?pageSize=" + strconv.FormatInt(pageSize, 10) + "&page=" + strconv.FormatInt(page, 10) + "&type=transfer"

	// 构造请求
	resp, err = http.Get(apiUrl)
	if err != nil {
		return nil, 0, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}

	var l *FilData
	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}

	if 0 >= l.TotalCount {
		return res, 0, nil
	}

	if nil == l.Transfers {
		return res, 0, nil
	}

	res = make([]*FilDataList, 0)
	for _, v := range l.Transfers {
		res = append(res, v)
	}

	return res, l.TotalCount, nil
}

// GetFilData .
func (b *BinanceUserUsecase) GetFilData(ctx context.Context, req *v1.GetFilDataRequest) (*v1.GetFilDataReply, error) {
	var (
		err error
		//filDataF0 map[string][]*FilData2
		globalFilData map[string][]*v1.GetFilDataReply_DataList
	)

	//f0 := make([]string, 0)
	////f0 = append(f0, "f01050978", "f0693127", "f0490501", "f097432", "f01307626", "f02723", "f0515264", "f083920", "f02609", "f0165111", "f01236627", "f01368089")
	//f0 = append(f0, req.Address)
	//filDataF0, err = b.binanceUserRepo.GetFilData(f0)
	//if nil != err {
	//	return nil, err
	//}

	var (
		getFilData func(address string, globalFilData map[string][]*v1.GetFilDataReply_DataList) []*v1.GetFilDataReply_DataList
	)
	getFilData = func(address string, globalFilData map[string][]*v1.GetFilDataReply_DataList) []*v1.GetFilDataReply_DataList {
		var (
			filData map[string][]*FilData2
		)
		fmt.Println(address)

		// 	存在返回
		if _, ok := globalFilData[address]; ok {
			return globalFilData[address] // 已经查过的人already字段为true
		}

		tmp := make([]string, 0)
		tmp = append(tmp, address)
		filData, err = b.binanceUserRepo.GetFilData(tmp)
		if nil != err {
			fmt.Println(err)
			return nil
		}

		if 0 >= len(filData) {
			return nil
		}

		res := make([]*v1.GetFilDataReply_DataList, 0)
		globalRes := make([]*v1.GetFilDataReply_DataList, 0)
		for addressK, v := range filData {
			for _, vFilData := range v {
				globalRes = append(globalRes, &v1.GetFilDataReply_DataList{
					To:        vFilData.To,
					Value:     vFilData.Value,
					ListChild: nil,  // 递归
					Already:   true, // 标记链路上
				})
			}
			// 挂上去，查过的
			if _, ok := globalFilData[addressK]; !ok {
				globalFilData[addressK] = globalRes
			}

			for _, vFilData := range v {
				res = append(res, &v1.GetFilDataReply_DataList{
					To:        vFilData.To,
					Value:     vFilData.Value,
					ListChild: getFilData(vFilData.To, globalFilData), // 递归
				})
			}
		}

		return res
	}

	//res := make([]*v1.GetFilDataReply_DataList, 0)
	//for _, vFilDataF0 := range filDataF0 {
	//	for _, vVFilDataF0 := range vFilDataF0 {
	//		tmp := make([]string, 0)
	//		tmp = append(tmp, vVFilDataF0.To)
	//
	//		res = append(res, &v1.GetFilDataReply_DataList{
	//			To:        vVFilDataF0.To,
	//			Value:     vVFilDataF0.Value,
	//			ListChild: getFilData(tmp), // 递归
	//		})
	//	}
	//}

	globalFilData = make(map[string][]*v1.GetFilDataReply_DataList, 0)
	return &v1.GetFilDataReply{
		List: getFilData(req.Address, globalFilData),
	}, nil
}

// GetBinanceTradersTrade .
func (b *BinanceUserUsecase) GetBinanceTradersTrade(ctx context.Context, req *v1.GetBinanceTradersTradeRequest) (*v1.GetBinanceTradersTradeReply, error) {
	var (
		NewBinanceTradeHistory []*BinanceTradeHistory
		err                    error
	)
	NewBinanceTradeHistory, err = b.binanceUserRepo.GetBinanceTradeHistoryByTraderNumNewest(req.TraderNum, 6000)
	if nil != err {
		return nil, err
	}

	res := &v1.GetBinanceTradersTradeReply{List: make([]*v1.GetBinanceTradersTradeReply_DataList, 0)}
	for _, v := range NewBinanceTradeHistory {
		res.List = append(res.List, &v1.GetBinanceTradersTradeReply_DataList{
			Time:         v.Time,
			CreatedAt:    v.CreatedAt.String(),
			Symbol:       v.Symbol,
			Side:         v.Side,
			PositionSide: v.PositionSide,
			Qty:          v.Qty,
		})
	}

	return res, err
}

// GetUserBindData .
func (b *BinanceUserUsecase) GetUserBindData(ctx context.Context, req *v1.GetUserBindDataRequest) (*v1.GetUserBindDataReply, error) {
	var (
		err                  error
		NewUserBindTraderTwo map[uint64][]*UserBindTrader
		userIds              []uint64
		userIdsMap           map[uint64]uint64
	)

	NewUserBindTraderTwo, err = b.binanceUserRepo.GetUserBindTraderTwoByAlreadyInitOrder()
	if nil != err {
		return nil, err
	}

	userIds = make([]uint64, 0)
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserBindTrader := range NewUserBindTraderTwo {
		for _, vVUserBindTrader := range vUserBindTrader {
			if _, ok := userIdsMap[vVUserBindTrader.UserId]; ok {
				continue
			}

			userIdsMap[vVUserBindTrader.UserId] = vVUserBindTrader.UserId
			userIds = append(userIds, vVUserBindTrader.UserId)
		}
	}

	if 0 >= len(userIds) {
		return nil, nil
	}

	var (
		traders map[uint64]*Trader
		users   map[uint64]*User
	)
	traders, err = b.binanceUserRepo.GetTraders()
	if nil != err {
		return nil, nil
	}

	// 获取用户信息，余额信息，收益信息
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		return nil, nil
	}

	res := &v1.GetUserBindDataReply{
		List: make([]*v1.GetUserBindDataReply_DataList, 0),
	}

	for traderId, vUserBindTraders := range NewUserBindTraderTwo {
		var (
			traderName string
			traderNum  string
		)
		if _, ok := traders[traderId]; ok {
			traderNum = traders[traderId].PortfolioId
			traderName = traders[traderId].Name
		}

		// 先更新状态
		for _, vVUserBindTraders := range vUserBindTraders {
			if _, ok := users[vVUserBindTraders.UserId]; !ok {
				continue
			}

			// 带单
			if 1 > users[vVUserBindTraders.UserId].IsDai {
				continue
			}

			// 正在开启
			if 0 == vVUserBindTraders.Status && 1 == vVUserBindTraders.InitOrder {
				res.List = append(res.List, &v1.GetUserBindDataReply_DataList{
					Id:         vVUserBindTraders.ID,
					UserId:     vVUserBindTraders.UserId,
					Address:    users[vVUserBindTraders.UserId].Address,
					Amount:     vVUserBindTraders.Amount,
					TraderNum:  traderNum,
					TraderName: traderName,
				})
			}
		}
	}

	return res, nil
}

// InsertUserBindData .
func (b *BinanceUserUsecase) InsertUserBindData(ctx context.Context, req *v1.InsertUserBindDataRequest) (*v1.InsertUserBindDataReply, error) {
	var (
		err             error
		traders         map[uint64]*Trader
		traderId        uint64
		userId          = req.SendBody.UserId
		userBindTraders []*UserBindTrader
	)
	traders, err = b.binanceUserRepo.GetTraders()
	if nil != err {
		return &v1.InsertUserBindDataReply{
			Msg: "未查询到交易员信息",
			Res: false,
		}, nil
	}

	for _, trader := range traders {
		// 可用的
		if 1 != trader.Status {
			continue
		}

		if req.SendBody.TraderNum == trader.PortfolioId {
			traderId = trader.ID
		}
	}

	if 0 >= traderId {
		return &v1.InsertUserBindDataReply{
			Msg: "未查询到可用交易员信息",
			Res: false,
		}, nil
	}

	// 改为带单用户
	if 0 < len(req.SendBody.Address) {
		var (
			user *User
		)
		user, err = b.binanceUserRepo.GetUserByAddress(ctx, req.SendBody.Address)
		if nil != err {
			return &v1.InsertUserBindDataReply{
				Msg: "未查询到用户信息",
				Res: false,
			}, nil

		}

		if 1 > user.IsDai {
			_, err = b.binanceUserRepo.UpdateUserIsDai(ctx, req.SendBody.Address)
			if nil != err {
				return &v1.InsertUserBindDataReply{
					Msg: "修改未带单员状态失败",
					Res: false,
				}, nil

			}
		}

		userId = user.ID
	}

	userBindTraders, err = b.binanceUserRepo.GetUserBindTraderTwoByUserId(userId)
	if nil != err {
		return &v1.InsertUserBindDataReply{
			Msg: "未查询到用户信息",
			Res: false,
		}, nil
	}

	if nil == userBindTraders {
		_, err = b.binanceUserRepo.InsertUserBindTraderTwo(ctx, userId, traderId, req.SendBody.Amount)
		if nil != err {
			return &v1.InsertUserBindDataReply{
				Msg: "错误的添加",
				Res: false,
			}, nil
		}
	} else {

		var tmpAlreadyBind *UserBindTrader
		for _, vUserBindTraders := range userBindTraders {
			if vUserBindTraders.TraderId == traderId {
				tmpAlreadyBind = vUserBindTraders
			}
		}

		if nil == tmpAlreadyBind {
			_, err = b.binanceUserRepo.InsertUserBindTraderTwo(ctx, userId, traderId, req.SendBody.Amount)
			if nil != err {
				return &v1.InsertUserBindDataReply{
					Msg: "错误的添加",
					Res: false,
				}, nil
			}
		} else {
			// 正在运行的会被清理掉
			if 1 == tmpAlreadyBind.InitOrder && 0 == tmpAlreadyBind.Status {
				// 更新状态
				_, err = b.binanceUserRepo.UpdatesUserBindTraderTwoAdminOverOrderById(ctx, tmpAlreadyBind.ID)
				if nil != err {
					return &v1.InsertUserBindDataReply{
						Msg: "错误的更新",
						Res: false,
					}, nil
				}

				err = b.handleAdminOverOrderTwo(ctx)
				if nil != err {
					return &v1.InsertUserBindDataReply{
						Msg: "错误的平仓",
						Res: false,
					}, nil
				}
			}

			_, err = b.binanceUserRepo.UpdatesUserBindTraderTwoById(ctx, tmpAlreadyBind.ID, req.SendBody.Amount)
			if nil != err {
				return &v1.InsertUserBindDataReply{
					Msg: "错误的更新",
					Res: false,
				}, nil
			}
		}
	}

	return &v1.InsertUserBindDataReply{
		Msg: "ok",
		Res: true,
	}, nil
}

// DeleteUserBindData .
func (b *BinanceUserUsecase) DeleteUserBindData(ctx context.Context, req *v1.DeleteUserBindDataRequest) (*v1.DeleteUserBindDataReply, error) {
	var (
		err             error
		userBindTraders *UserBindTrader
	)
	userBindTraders, err = b.binanceUserRepo.GetUserBindTraderTwoById(req.SendBody.Id)
	if nil != err || nil == userBindTraders {
		return &v1.DeleteUserBindDataReply{
			Msg: "未查询到信息",
			Res: false,
		}, nil
	}

	// 更新状态
	_, err = b.binanceUserRepo.UpdatesUserBindTraderTwoAdminOverOrderById(ctx, userBindTraders.ID)
	if nil != err {
		return &v1.DeleteUserBindDataReply{
			Msg: "错误的更新",
			Res: false,
		}, nil
	}

	err = b.handleAdminOverOrderTwo(ctx)
	if nil != err {
		return &v1.DeleteUserBindDataReply{
			Msg: "错误的平仓",
			Res: false,
		}, nil
	}

	_, err = b.binanceUserRepo.UpdatesUserBindTraderTwoAfterAdminOverOrderById(ctx, userBindTraders.ID)
	if nil != err {
		return &v1.DeleteUserBindDataReply{
			Msg: "错误的更新2",
			Res: false,
		}, nil
	}

	return &v1.DeleteUserBindDataReply{
		Msg: "ok",
		Res: true,
	}, nil
}
