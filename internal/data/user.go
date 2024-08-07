package data

import (
	"binanceexchange_user/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type User struct {
	ID                  uint64    `gorm:"primarykey;type:int"`
	Address             string    `gorm:"type:varchar(100);not null"`
	ApiStatus           uint64    `gorm:"type:int;not null"`
	ApiKey              string    `gorm:"type:varchar(200);not null"`
	ApiSecret           string    `gorm:"type:varchar(200);not null"`
	BindTraderStatus    uint64    `gorm:"type:int;not null"`
	BindTraderStatusTfi uint64    `gorm:"type:int;not null"`
	IsDai               uint64    `gorm:"type:int;not null"`
	UseNewSystem        uint64    `gorm:"type:int;not null"`
	CreatedAt           time.Time `gorm:"type:datetime;not null"`
	UpdatedAt           time.Time `gorm:"type:datetime;not null"`
}

type UserInfo struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	BId       uint64    `gorm:"type:bigint(20);not null"`
	BaseMoney float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type TraderInfo struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	TraderId  uint64    `gorm:"type:int;not null"`
	BId       uint64    `gorm:"type:bigint(20);not null"`
	BaseMoney float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserBalance struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	Balance   string    `gorm:"type:varchar(100);not null"`
	Cost      uint64    `gorm:"type:bigint(20);not null"`
	Open      uint64    `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserBalanceRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	Amount    string    `gorm:"type:varchar(100);not null"`
	Balance   string    `gorm:"type:varchar(100);not null"`
	tx        string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserAmount struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	Amount    int64     `gorm:"type:bigint(20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserAmountRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	OrderId   uint64    `gorm:"type:bigint(20);not null"`
	Amount    int64     `gorm:"type:bigint(20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type Trader struct {
	ID          uint64    `gorm:"primarykey;type:int"`
	IsOpen      uint64    `gorm:"type:int;not null"`
	Amount      uint64    `gorm:"type:bigint(20);not null"`
	BaseMoney   float64   `gorm:"type:decimal(40,8);not null"`
	Name        string    `gorm:"type:varchar(100);not null"`
	PortfolioId string    `gorm:"column:portfolioId;type:varchar(100);not null"`
	CreatedAt   time.Time `gorm:"type:datetime;not null"`
	UpdatedAt   time.Time `gorm:"type:datetime;not null"`
}

type LhCoinSymbol struct {
	ID                uint64 `gorm:"primarykey;type:int"`
	Symbol            string `gorm:"type:varchar(100);not null"`
	QuantityPrecision int64  `gorm:"type:int;not null"`
	PricePrecision    int64  `gorm:"type:int;not null"`
}

type LhTraderPosition struct {
	ID       uint64  `gorm:"primarykey;type:int"`
	TraderId uint64  `gorm:"type:int;not null"`
	Symbol   string  `gorm:"type:varchar(100);not null"`
	Side     string  `gorm:"type:varchar(100);not null"`
	Type     string  `gorm:"type:varchar(100);not null"`
	Qty      float64 `gorm:"type:decimal(65,20);not null"`
}

type ZyTraderPosition struct {
	ID             uint64  `gorm:"primarykey;type:int"`
	PositionSide   string  `gorm:"column:positionSide;type:varchar(100);not null"`
	Symbol         string  `gorm:"type:varchar(100);not null"`
	PositionAmount float64 `gorm:"column:positionAmount;type:decimal(60,8);not null"`
}

type TraderPosition struct {
	ID             uint64  `gorm:"primarykey;type:int"`
	PositionSide   string  `gorm:"type:varchar(100);not null"`
	Symbol         string  `gorm:"type:varchar(100);not null"`
	PositionAmount float64 `gorm:"type:decimal(60,8);not null"`
}

type TraderEmail struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	TraderId  uint64    `gorm:"type:int;not null"`
	Email     string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type TraderPositionNew struct {
	ID     uint64  `gorm:"primarykey;type:int"`
	Closed uint64  `gorm:"type:bigint(20);not null"`
	Opened uint64  `gorm:"type:bigint(20);not null"`
	Symbol string  `gorm:"type:varchar(100);not null"`
	Side   string  `gorm:"type:varchar(100);not null"`
	Qty    float64 `gorm:"type:decimal(65,20);not null"`
}

type UserBindTrader struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	TraderId  uint64    `gorm:"type:int;not null"`
	Amount    uint64    `gorm:"type:bigint(20);not null"`
	Status    uint64    `gorm:"type:int;not null"`
	InitOrder uint64    `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserBindAfterUnbind struct {
	ID           uint64  `gorm:"primarykey;type:int"`
	UserId       uint64  `gorm:"type:int;not null"`
	TraderId     uint64  `gorm:"type:int;not null"`
	Status       uint64  `gorm:"type:int;not null"`
	Symbol       string  `gorm:"type:varchar(100);not null"`
	PositionSide string  `gorm:"type:varchar(100);not null"`
	Quantity     float64 `gorm:"type:decimal(65,20);not null"`
	Amount       uint64  `gorm:"type:bigint(20);not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserOrder struct {
	ID            uint64  `gorm:"primarykey;type:int"`
	UserId        uint64  `gorm:"type:int;not null"`
	TraderId      uint64  `gorm:"type:int;not null"`
	ClientOrderId string  `gorm:"type:varchar(100)"`
	OrderId       string  `gorm:"type:varchar(100);not null"`
	Symbol        string  `gorm:"type:varchar(100);not null"`
	Side          string  `gorm:"type:varchar(100);not null"`
	PositionSide  string  `gorm:"type:varchar(100);not null"`
	Quantity      float64 `gorm:"type:decimal(65,20);not null"`
	Price         float64 `gorm:"type:decimal(65,20);not null"`
	TraderQty     float64 `gorm:"type:decimal(65,20);not null"`
	OrderType     string  `gorm:"type:varchar(100);not null"`
	ClosePosition string  `gorm:"type:varchar(100);not null"`
	CumQuote      float64 `gorm:"type:decimal(65,20);not null"`
	ExecutedQty   float64 `gorm:"type:decimal(65,20);not null"`
	AvgPrice      float64 `gorm:"type:decimal(65,20);not null"`
	HandleStatus  int64   `gorm:"type:int;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UserOrderErr struct {
	ID            uint64  `gorm:"primarykey;type:int"`
	UserId        uint64  `gorm:"type:int;not null"`
	TraderId      uint64  `gorm:"type:int;not null"`
	ClientOrderId string  `gorm:"type:varchar(100)"`
	OrderId       string  `gorm:"type:varchar(100);not null"`
	Symbol        string  `gorm:"type:varchar(100);not null"`
	Side          string  `gorm:"type:varchar(100);not null"`
	PositionSide  string  `gorm:"type:varchar(100);not null"`
	Quantity      float64 `gorm:"type:decimal(65,20);not null"`
	Price         float64 `gorm:"type:decimal(65,20);not null"`
	TraderQty     float64 `gorm:"type:decimal(65,20);not null"`
	OrderType     string  `gorm:"type:varchar(100);not null"`
	ClosePosition string  `gorm:"type:varchar(100);not null"`
	CumQuote      float64 `gorm:"type:decimal(65,20);not null"`
	ExecutedQty   float64 `gorm:"type:decimal(65,20);not null"`
	AvgPrice      float64 `gorm:"type:decimal(65,20);not null"`
	HandleStatus  int64   `gorm:"type:int;not null"`
	Code          int64   `gorm:"type:int"`
	Msg           string  `gorm:"type:varchar(200)"`
	InitOrder     int64   `gorm:"type:int"`
	Proportion    float64 `gorm:"type:decimal(65,20)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type TradingBoxOpen struct {
	ID             uint64    `gorm:"primarykey;type:int"`
	TokenId        uint64    `gorm:"type:int;not null"`
	Amount         uint64    `gorm:"type:varchar(200);not null"`
	AmountTotal    uint64    `gorm:"type:varchar(200);not null"`
	WithdrawStatus uint64    `gorm:"type:int;not null"`
	AmountRate     float64   `gorm:"type:decimal(65,20);not null"`
	Total          float64   `gorm:"type:decimal(65,20);not null"`
	WithdrawAmount float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt      time.Time `gorm:"type:datetime;not null"`
	UpdatedAt      time.Time `gorm:"type:datetime;not null"`
}

type BinanceTradeHistory struct {
	ID                  uint64  `gorm:"primarykey;type:int"`
	TraderNum           uint64  `gorm:"type:bigint(20);not null"`
	Time                uint64  `gorm:"type:bigint(20);not null"`
	Symbol              string  `gorm:"type:varchar(45);not null"`
	Side                string  `gorm:"type:varchar(45);not null"`
	Price               float64 `gorm:"type:decimal(65,20);not null"`
	Fee                 float64 `gorm:"type:decimal(65,20);not null"`
	FeeAsset            string  `gorm:"type:varchar(45);not null"`
	Quantity            float64 `gorm:"type:decimal(65,20);not null"`
	QuantityAsset       string  `gorm:"type:varchar(45);not null"`
	RealizedProfit      float64 `gorm:"type:decimal(65,20);not null"`
	RealizedProfitAsset string  `gorm:"type:varchar(45);not null"`
	BaseAsset           string  `gorm:"type:varchar(45);not null"`
	Qty                 float64 `gorm:"type:decimal(65,20);not null"`
	PositionSide        string  `gorm:"type:varchar(45);not null"`
	ActiveBuy           string  `gorm:"type:varchar(45);not null"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type BinancePositionHistory struct {
	ID              uint64  `gorm:"primarykey;type:int"`
	Symbol          string  `gorm:"type:varchar(45);not null"`
	Side            string  `gorm:"type:varchar(45);not null"`
	Status          string  `gorm:"type:varchar(45);not null"`
	Closed          uint64  `gorm:"type:bigint(20);not null"`
	Opened          uint64  `gorm:"type:bigint(20);not null"`
	AvgCost         float64 `gorm:"type:decimal(65,20);not null"`
	AvgClosePrice   float64 `gorm:"type:decimal(65,20);not null"`
	ClosingPnl      float64 `gorm:"type:decimal(65,20);not null"`
	MaxOpenInterest float64 `gorm:"type:decimal(65,20);not null"`
	ClosedVolume    float64 `gorm:"type:decimal(65,20);not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type UserOrderDo struct {
	ID           uint64  `gorm:"primarykey;type:int"`
	ApiKey       string  `gorm:"type:varchar(200);not null"`
	ApiSecret    string  `gorm:"type:varchar(200);not null"`
	ApiKeyTwo    string  `gorm:"type:varchar(200);not null"`
	ApiSecretTwo string  `gorm:"type:varchar(200);not null"`
	Symbol       string  `gorm:"type:varchar(200);not null"`
	Status       uint64  `gorm:"type:int;not null"`
	CloseRate    uint64  `gorm:"type:int;not null"`
	CloseBase    uint64  `gorm:"type:int;not null"`
	Qty          float64 `gorm:"type:decimal(65,20);not null"`
	Price        float64 `gorm:"type:decimal(65,20);not null"`
	ClosePrice   float64 `gorm:"type:decimal(65,20);not null"`
	Amount       float64 `gorm:"type:decimal(65,20);not null"`
	QtyTwo       float64 `gorm:"type:decimal(65,20);not null"`
	PriceTwo     float64 `gorm:"type:decimal(65,20);not null"`
	AmountTwo    float64 `gorm:"type:decimal(65,20);not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserOrderDoNew struct {
	ID           uint64  `gorm:"primarykey;type:int"`
	ApiKey       string  `gorm:"type:varchar(200);not null"`
	ApiSecret    string  `gorm:"type:varchar(200);not null"`
	ApiKeyTwo    string  `gorm:"type:varchar(200);not null"`
	ApiSecretTwo string  `gorm:"type:varchar(200);not null"`
	Symbol       string  `gorm:"type:varchar(200);not null"`
	SymbolTwo    string  `gorm:"type:varchar(200);not null"`
	Status       uint64  `gorm:"type:int;not null"`
	Qty          float64 `gorm:"type:decimal(65,20);not null"`
	QtyTwo       float64 `gorm:"type:decimal(65,20);not null"`
	Side         string  `gorm:"type:varchar(200);not null"`
	SideTwo      string  `gorm:"type:varchar(200);not null"`
	OrderId      string  `gorm:"type:varchar(200);not null"`
	OrderIdTwo   string  `gorm:"type:varchar(200);not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type FilData struct {
	ID       uint64 `gorm:"primarykey;type:int"`
	FromUser string `gorm:"type:varchar(200);not null"`
	FromType uint64 `gorm:"type:int;not null"`
	To       string `gorm:"type:varchar(200);not null"`
	Value    string `gorm:"type:varchar(200);not null"`
}

type BinanceTrade struct {
	ID        uint64 `gorm:"primarykey;type:int"`
	TraderNum uint64 `gorm:"type:bigint(20);not null"`
	Status    uint64 `gorm:"type:bigint(20);not null"`
}

type BinanceUserRepo struct {
	data *Data
	log  *log.Helper
}

func NewBinanceUserRepo(data *Data, logger log.Logger) biz.BinanceUserRepo {
	return &BinanceUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// InsertUser .
func (b *BinanceUserRepo) InsertUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	insertUser := &User{
		Address: user.Address,
	}

	res := b.data.DB(ctx).Table("new_user").Create(&insertUser)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_ERROR", "创建数据失败")
	}

	return &biz.User{
		ID:        insertUser.ID,
		Address:   insertUser.Address,
		ApiKey:    insertUser.ApiKey,
		ApiSecret: insertUser.ApiSecret,
		CreatedAt: insertUser.CreatedAt,
		UpdatedAt: insertUser.UpdatedAt,
	}, nil
}

// UpdateUserApiStatus .
func (b *BinanceUserRepo) UpdateUserApiStatus(ctx context.Context, userId uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user").Where("id=?", userId).
		Updates(map[string]interface{}{"api_status": 1, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_ERROR", "UPDATE_USER_ERROR")
	}

	return true, nil
}

// UpdateUser .
func (b *BinanceUserRepo) UpdateUser(ctx context.Context, userId uint64, apiKey string, apiSecret string) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user").Where("id=?", userId).
		Updates(map[string]interface{}{"api_key": apiKey, "api_secret": apiSecret, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_ERROR", "UPDATE_USER_ERROR")
	}

	return true, nil
}

// UpdateUserBindTraderStatus .
func (b *BinanceUserRepo) UpdateUserBindTraderStatus(ctx context.Context, userId uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user").Where("id=?", userId).
		Updates(map[string]interface{}{"bind_trader_status": 1, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_ERROR", "UPDATE_USER_ERROR")
	}

	return true, nil
}

// UpdateUserIsDai .
func (b *BinanceUserRepo) UpdateUserIsDai(ctx context.Context, address string) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user").Where("address=?", address).
		Updates(map[string]interface{}{"is_dai": 1, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_ERROR", "UPDATE_USER_ERROR")
	}

	return true, nil
}

// UpdateUserBindTraderStatusNo .
func (b *BinanceUserRepo) UpdateUserBindTraderStatusNo(ctx context.Context, userId uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user").Where("id=?", userId).
		Updates(map[string]interface{}{"bind_trader_status": 2, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_ERROR", "UPDATE_USER_ERROR")
	}

	return true, nil
}

// UpdateUserBindTraderStatusNoTfi .
func (b *BinanceUserRepo) UpdateUserBindTraderStatusNoTfi(ctx context.Context, userId uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user").Where("id=?", userId).
		Updates(map[string]interface{}{"bind_trader_status_tfi": 2, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_ERROR", "UPDATE_USER_ERROR")
	}

	return true, nil
}

// InsertUserBalance .
func (b *BinanceUserRepo) InsertUserBalance(ctx context.Context, userBalance *biz.UserBalance) (bool, error) {
	insertUserBalance := &UserBalance{
		UserId:  userBalance.UserId,
		Balance: userBalance.Balance,
		Cost:    userBalance.Cost,
		Open:    userBalance.Open,
	}

	res := b.data.DB(ctx).Table("new_user_balance").Create(&insertUserBalance)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_BALANCE_ERROR", "创建数据失败")
	}

	return true, nil
}

// InsertUserBalanceTwo .
func (b *BinanceUserRepo) InsertUserBalanceTwo(ctx context.Context, userBalance *biz.UserBalance) (bool, error) {
	insertUserBalance := &UserBalance{
		UserId:  userBalance.UserId,
		Balance: userBalance.Balance,
		Cost:    userBalance.Cost,
		Open:    userBalance.Open,
	}

	res := b.data.DB(ctx).Table("new_user_balance_two").Create(&insertUserBalance)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_BALANCE_TWO_ERROR", "创建数据失败")
	}

	return true, nil
}

// UpdatesUserBalance .
func (b *BinanceUserRepo) UpdatesUserBalance(ctx context.Context, userId uint64, balance string, cost uint64, open uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_balance").Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance": balance, "cost": cost, "open": open, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_BALANCE_ERROR", "UPDATE_USER_BALANCE_ERROR")
	}

	return true, nil
}

// UpdatesUserBalanceTwo .
func (b *BinanceUserRepo) UpdatesUserBalanceTwo(ctx context.Context, userId uint64, balance string, cost uint64, open uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_balance_two").Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance": balance, "cost": cost, "open": open, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_BALANCE_TWO_ERROR", "UPDATE_USER_BALANCE_TWO_ERROR")
	}

	return true, nil
}

// InsertUserBalanceRecord .
func (b *BinanceUserRepo) InsertUserBalanceRecord(ctx context.Context, userBalanceRecord *biz.UserBalanceRecord) (bool, error) {
	insertUserBalanceRecord := &UserBalanceRecord{
		UserId:  userBalanceRecord.UserId,
		Amount:  userBalanceRecord.Amount,
		Balance: userBalanceRecord.Balance,
		tx:      userBalanceRecord.Tx,
	}

	res := b.data.DB(ctx).Table("new_user_balance_record").Create(&insertUserBalanceRecord)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_BALANCE_RECORD_ERROR", "创建数据失败")
	}

	return true, nil
}

// InsertUserBalanceRecordTwo .
func (b *BinanceUserRepo) InsertUserBalanceRecordTwo(ctx context.Context, userBalanceRecord *biz.UserBalanceRecord) (bool, error) {
	insertUserBalanceRecord := &UserBalanceRecord{
		UserId:  userBalanceRecord.UserId,
		Amount:  userBalanceRecord.Amount,
		Balance: userBalanceRecord.Balance,
		tx:      userBalanceRecord.Tx,
	}

	res := b.data.DB(ctx).Table("new_user_balance_record_two").Create(&insertUserBalanceRecord)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_BALANCE_TWO_RECORD_ERROR", "创建数据失败")
	}

	return true, nil
}

// InsertUserAmount .
func (b *BinanceUserRepo) InsertUserAmount(ctx context.Context, userAmount *biz.UserAmount) (bool, error) {
	insertUserAmount := &UserAmount{
		UserId: userAmount.UserId,
		Amount: 0,
	}

	res := b.data.DB(ctx).Table("new_user_amount").Create(&insertUserAmount)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_AMOUNT_ERROR", "创建数据失败")
	}

	return true, nil
}

// InsertUserAmountTwo .
func (b *BinanceUserRepo) InsertUserAmountTwo(ctx context.Context, userAmount *biz.UserAmount) (bool, error) {
	insertUserAmount := &UserAmount{
		UserId: userAmount.UserId,
		Amount: 0,
	}

	res := b.data.DB(ctx).Table("new_user_amount_two").Create(&insertUserAmount)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_AMOUNT_TWO_ERROR", "创建数据失败")
	}

	return true, nil
}

// UpdatesUserAmount .
func (b *BinanceUserRepo) UpdatesUserAmount(ctx context.Context, userId uint64, amount int64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_amount").Where("user_id=?", userId).
		Updates(map[string]interface{}{"amount": gorm.Expr("amount + ?", amount), "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_AMOUNT_ERROR", "UPDATE_USER_AMOUNT_ERROR")
	}

	return true, nil
}

// UpdatesUserAmountTwo .
func (b *BinanceUserRepo) UpdatesUserAmountTwo(ctx context.Context, userId uint64, amount int64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_amount_two").Where("user_id=?", userId).
		Updates(map[string]interface{}{"amount": gorm.Expr("amount + ?", amount), "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_AMOUNT_TWO_ERROR", "UPDATE_USER_AMOUNT_ERROR")
	}

	return true, nil
}

// InsertUserAmountRecord .
func (b *BinanceUserRepo) InsertUserAmountRecord(ctx context.Context, userAmountRecord *biz.UserAmountRecord) (bool, error) {
	insertUserAmountRecord := &UserAmountRecord{
		UserId:  userAmountRecord.UserId,
		OrderId: userAmountRecord.OrderId,
		Amount:  userAmountRecord.Amount,
	}

	res := b.data.DB(ctx).Table("new_user_amount_record").Create(&insertUserAmountRecord)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_AMOUNT_RECORD_ERROR", "创建数据失败")
	}

	return true, nil
}

// InsertUserAmountRecordTwo .
func (b *BinanceUserRepo) InsertUserAmountRecordTwo(ctx context.Context, userAmountRecord *biz.UserAmountRecord) (bool, error) {
	insertUserAmountRecord := &UserAmountRecord{
		UserId:  userAmountRecord.UserId,
		OrderId: userAmountRecord.OrderId,
		Amount:  userAmountRecord.Amount,
	}

	res := b.data.DB(ctx).Table("new_user_amount_record_two").Create(&insertUserAmountRecord)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_AMOUNT_RECORD_TWO_ERROR", "创建数据失败")
	}

	return true, nil
}

// InsertUserBindTrader .
func (b *BinanceUserRepo) InsertUserBindTrader(ctx context.Context, userId uint64, traderId uint64, amount uint64) (*biz.UserBindTrader, error) {
	insertUserBindTrader := &UserBindTrader{
		UserId:   userId,
		TraderId: traderId,
		Amount:   amount,
	}

	res := b.data.DB(ctx).Table("new_user_bind_trader").Create(&insertUserBindTrader)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_BIND_TRADER_ERROR", "创建数据失败")
	}

	return &biz.UserBindTrader{
		ID:        insertUserBindTrader.ID,
		UserId:    insertUserBindTrader.UserId,
		TraderId:  insertUserBindTrader.TraderId,
		Amount:    insertUserBindTrader.Amount,
		CreatedAt: insertUserBindTrader.CreatedAt,
		UpdatedAt: insertUserBindTrader.UpdatedAt,
	}, nil
}

// InsertUserBindTraderTwo .
func (b *BinanceUserRepo) InsertUserBindTraderTwo(ctx context.Context, userId uint64, traderId uint64, amount uint64) (*biz.UserBindTrader, error) {
	insertUserBindTrader := &UserBindTrader{
		UserId:   userId,
		TraderId: traderId,
		Amount:   amount,
	}

	res := b.data.DB(ctx).Table("new_user_bind_trader_two").Create(&insertUserBindTrader)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_BIND_TRADER_ERROR", "创建数据失败")
	}

	return &biz.UserBindTrader{
		ID:        insertUserBindTrader.ID,
		UserId:    insertUserBindTrader.UserId,
		TraderId:  insertUserBindTrader.TraderId,
		Amount:    insertUserBindTrader.Amount,
		CreatedAt: insertUserBindTrader.CreatedAt,
		UpdatedAt: insertUserBindTrader.UpdatedAt,
	}, nil
}

// UpdatesUserBindTraderStatus .
func (b *BinanceUserRepo) UpdatesUserBindTraderStatus(ctx context.Context, userId uint64, status uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_bind_trader").Where("user_id=?", userId).
		Updates(map[string]interface{}{"status": status, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_BIND_TRADER_ERROR", "UPDATE_USER_BIND_TRADER_ERROR")
	}

	return true, nil
}

// UpdatesUserBindTraderStatusAndInitOrderById .
func (b *BinanceUserRepo) UpdatesUserBindTraderStatusAndInitOrderById(ctx context.Context, id uint64, status uint64, initOrder uint64, amount uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_bind_trader").Where("id=?", id).
		Updates(map[string]interface{}{"status": status, "init_order": initOrder, "amount": amount, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_BIND_TRADER_ERROR", "UPDATE_USER_BIND_TRADER_ERROR")
	}

	return true, nil
}

// UpdatesUserBindTraderInitOrderById .
func (b *BinanceUserRepo) UpdatesUserBindTraderInitOrderById(ctx context.Context, id uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_bind_trader").Where("id=?", id).
		Updates(map[string]interface{}{"init_order": 1, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_BIND_TRADER_ERROR", "UPDATE_USER_BIND_TRADER_ERROR")
	}

	return true, nil
}

// UpdatesUserBindTraderTwoInitOrderById .
func (b *BinanceUserRepo) UpdatesUserBindTraderTwoInitOrderById(ctx context.Context, id uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_bind_trader_two").Where("id=?", id).
		Updates(map[string]interface{}{"init_order": 1, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_BIND_TRADER_TWO_ERROR", "UPDATE_USER_BIND_TRADER_ERROR")
	}

	return true, nil
}

// UpdatesUserBindTraderTwoUnbindById .
func (b *BinanceUserRepo) UpdatesUserBindTraderTwoUnbindById(ctx context.Context, id uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_bind_trader_two").Where("id=?", id).
		Updates(map[string]interface{}{"status": 1, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_BIND_TRADER_TWO_ERROR", "UPDATE_USER_BIND_TRADER_ERROR")
	}

	return true, nil
}

// UpdatesUserBindTraderTwoById .
func (b *BinanceUserRepo) UpdatesUserBindTraderTwoById(ctx context.Context, id uint64, amount uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_bind_trader_two").Where("id=?", id).
		Updates(map[string]interface{}{"init_order": 0, "status": 0, "amount": amount, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_BIND_TRADER_TWO_ERROR", "UPDATE_USER_BIND_TRADER_ERROR")
	}

	return true, nil
}

// UpdatesUserBindTraderTwoAdminOverOrderById .
func (b *BinanceUserRepo) UpdatesUserBindTraderTwoAdminOverOrderById(ctx context.Context, id uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_bind_trader_two").Where("id=?", id).
		Updates(map[string]interface{}{"init_order": 2, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_BIND_TRADER_TWO_ERROR", "UPDATE_USER_BIND_TRADER_ERROR")
	}

	return true, nil
}

// UpdatesUserBindTraderTwoAfterAdminOverOrderById .
func (b *BinanceUserRepo) UpdatesUserBindTraderTwoAfterAdminOverOrderById(ctx context.Context, id uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_bind_trader_two").Where("id=?", id).
		Updates(map[string]interface{}{"init_order": 1, "status": 1, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_BIND_TRADER_TWO_ERROR", "UPDATE_USER_BIND_TRADER_ERROR")
	}

	return true, nil
}

// UpdatesUserBindTraderStatusById .
func (b *BinanceUserRepo) UpdatesUserBindTraderStatusById(ctx context.Context, id uint64, status uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_bind_trader").Where("id=?", id).
		Updates(map[string]interface{}{"status": status, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_BIND_TRADER_ERROR", "UPDATE_USER_BIND_TRADER_ERROR")
	}

	return true, nil
}

// UpdatesUserBindTraderStatusAndAmountById .
func (b *BinanceUserRepo) UpdatesUserBindTraderStatusAndAmountById(ctx context.Context, id uint64, status uint64, amount uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_bind_trader").Where("id=?", id).
		Updates(map[string]interface{}{"status": status, "amount": amount, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_BIND_TRADER_ERROR", "UPDATE_USER_BIND_TRADER_ERROR")
	}

	return true, nil
}

// DeleteUserBindTrader .
func (b *BinanceUserRepo) DeleteUserBindTrader(ctx context.Context, userId uint64) (bool, error) {
	var (
		err error
	)

	if err = b.data.DB(ctx).Table("new_user_bind_trader").Where("user_id=?", userId).Delete(&UserBindTrader{}).Error; nil != err {
		return false, errors.NotFound("DELETE_USER_BIND_TRADER_ERROR", "DELETE_USER_BIND_TRADER_ERROR")
	}

	return true, nil
}

// InsertUserOrderNew .
func (b *BinanceUserRepo) InsertUserOrderNew(ctx context.Context, order *biz.UserOrder) (*biz.UserOrder, error) {
	insertUserOrder := &UserOrder{
		UserId:        order.UserId,
		TraderId:      order.TraderId,
		ClientOrderId: order.ClientOrderId,
		OrderId:       order.OrderId,
		Symbol:        order.Symbol,
		Side:          order.Side,
		PositionSide:  order.PositionSide,
		Quantity:      order.Quantity,
		Price:         order.Price,
		TraderQty:     order.TraderQty,
		OrderType:     order.OrderType,
		ClosePosition: order.ClosePosition,
		CumQuote:      order.CumQuote,
		ExecutedQty:   order.ExecutedQty,
		AvgPrice:      order.AvgPrice,
		HandleStatus:  0,
	}

	res := b.data.DB(ctx).Table("new_user_order_two_" + strconv.FormatInt(int64(order.UserId), 10)).Create(&insertUserOrder)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_ORDER_ERROR", "创建数据失败")
	}

	return &biz.UserOrder{
		ID:            insertUserOrder.ID,
		UserId:        insertUserOrder.UserId,
		TraderId:      insertUserOrder.TraderId,
		ClientOrderId: insertUserOrder.ClientOrderId,
		OrderId:       insertUserOrder.OrderId,
		Symbol:        insertUserOrder.Symbol,
		Side:          insertUserOrder.Side,
		PositionSide:  insertUserOrder.PositionSide,
		Quantity:      insertUserOrder.Quantity,
		Price:         insertUserOrder.Price,
		TraderQty:     insertUserOrder.TraderQty,
		OrderType:     insertUserOrder.OrderType,
		ClosePosition: insertUserOrder.ClosePosition,
		CumQuote:      insertUserOrder.CumQuote,
		ExecutedQty:   insertUserOrder.ExecutedQty,
		AvgPrice:      insertUserOrder.AvgPrice,
		CreatedAt:     insertUserOrder.CreatedAt,
		UpdatedAt:     insertUserOrder.UpdatedAt,
	}, nil
}

// InsertUserOrder .
func (b *BinanceUserRepo) InsertUserOrder(ctx context.Context, order *biz.UserOrder) (*biz.UserOrder, error) {
	insertUserOrder := &UserOrder{
		ID:            order.ID,
		UserId:        order.UserId,
		TraderId:      order.TraderId,
		ClientOrderId: order.ClientOrderId,
		OrderId:       order.OrderId,
		Symbol:        order.Symbol,
		Side:          order.Side,
		PositionSide:  order.PositionSide,
		Quantity:      order.Quantity,
		Price:         order.Price,
		TraderQty:     order.TraderQty,
		OrderType:     order.OrderType,
		ClosePosition: order.ClosePosition,
		CumQuote:      order.CumQuote,
		ExecutedQty:   order.ExecutedQty,
		AvgPrice:      order.AvgPrice,
		HandleStatus:  0,
	}

	res := b.data.DB(ctx).Table("new_user_order").Create(&insertUserOrder)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_ORDER_ERROR", "创建数据失败")
	}

	return &biz.UserOrder{
		ID:            insertUserOrder.ID,
		UserId:        insertUserOrder.UserId,
		TraderId:      insertUserOrder.TraderId,
		ClientOrderId: insertUserOrder.ClientOrderId,
		OrderId:       insertUserOrder.OrderId,
		Symbol:        insertUserOrder.Symbol,
		Side:          insertUserOrder.Side,
		PositionSide:  insertUserOrder.PositionSide,
		Quantity:      insertUserOrder.Quantity,
		Price:         insertUserOrder.Price,
		TraderQty:     insertUserOrder.TraderQty,
		OrderType:     insertUserOrder.OrderType,
		ClosePosition: insertUserOrder.ClosePosition,
		CumQuote:      insertUserOrder.CumQuote,
		ExecutedQty:   insertUserOrder.ExecutedQty,
		AvgPrice:      insertUserOrder.AvgPrice,
		CreatedAt:     insertUserOrder.CreatedAt,
		UpdatedAt:     insertUserOrder.UpdatedAt,
	}, nil
}

// InsertUserOrderErr .
func (b *BinanceUserRepo) InsertUserOrderErr(ctx context.Context, order *biz.UserOrderErr) (*biz.UserOrderErr, error) {
	insertUserOrder := &UserOrderErr{
		ID:            order.ID,
		UserId:        order.UserId,
		TraderId:      order.TraderId,
		ClientOrderId: order.ClientOrderId,
		OrderId:       order.OrderId,
		Symbol:        order.Symbol,
		Side:          order.Side,
		PositionSide:  order.PositionSide,
		Quantity:      order.Quantity,
		Price:         order.Price,
		TraderQty:     order.TraderQty,
		OrderType:     order.OrderType,
		ClosePosition: order.ClosePosition,
		CumQuote:      order.CumQuote,
		ExecutedQty:   order.ExecutedQty,
		AvgPrice:      order.AvgPrice,
		Code:          order.Code,
		Msg:           order.Msg,
		InitOrder:     order.InitOrder,
		Proportion:    order.Proportion,
		HandleStatus:  0,
	}

	res := b.data.DB(ctx).Table("new_user_order_err").Create(&insertUserOrder)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_ORDER_ERR_ERROR", "创建数据失败")
	}

	return &biz.UserOrderErr{
		ID:            insertUserOrder.ID,
		UserId:        insertUserOrder.UserId,
		TraderId:      insertUserOrder.TraderId,
		ClientOrderId: insertUserOrder.ClientOrderId,
		OrderId:       insertUserOrder.OrderId,
		Symbol:        insertUserOrder.Symbol,
		Side:          insertUserOrder.Side,
		PositionSide:  insertUserOrder.PositionSide,
		Quantity:      insertUserOrder.Quantity,
		Price:         insertUserOrder.Price,
		TraderQty:     insertUserOrder.TraderQty,
		OrderType:     insertUserOrder.OrderType,
		ClosePosition: insertUserOrder.ClosePosition,
		CumQuote:      insertUserOrder.CumQuote,
		ExecutedQty:   insertUserOrder.ExecutedQty,
		AvgPrice:      insertUserOrder.AvgPrice,
		InitOrder:     insertUserOrder.InitOrder,
		Code:          insertUserOrder.Code,
		Msg:           insertUserOrder.Msg,
		Proportion:    insertUserOrder.Proportion,
		CreatedAt:     insertUserOrder.CreatedAt,
		UpdatedAt:     insertUserOrder.UpdatedAt,
	}, nil
}

// InsertUserOrderTwoNew .
func (b *BinanceUserRepo) InsertUserOrderTwoNew(ctx context.Context, order *biz.UserOrder) (*biz.UserOrder, error) {
	insertUserOrder := &UserOrder{
		ID:            order.ID,
		UserId:        order.UserId,
		TraderId:      order.TraderId,
		ClientOrderId: order.ClientOrderId,
		OrderId:       order.OrderId,
		Symbol:        order.Symbol,
		Side:          order.Side,
		PositionSide:  order.PositionSide,
		Quantity:      order.Quantity,
		Price:         order.Price,
		TraderQty:     order.TraderQty,
		OrderType:     order.OrderType,
		ClosePosition: order.ClosePosition,
		CumQuote:      order.CumQuote,
		ExecutedQty:   order.ExecutedQty,
		AvgPrice:      order.AvgPrice,
		HandleStatus:  0,
	}

	res := b.data.DB(ctx).Table("new_user_order_two_" + strconv.FormatInt(int64(order.UserId), 10)).Create(&insertUserOrder)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_ORDER_TWO_ERROR", "创建数据失败")
	}

	return &biz.UserOrder{
		ID:            insertUserOrder.ID,
		UserId:        insertUserOrder.UserId,
		TraderId:      insertUserOrder.TraderId,
		ClientOrderId: insertUserOrder.ClientOrderId,
		OrderId:       insertUserOrder.OrderId,
		Symbol:        insertUserOrder.Symbol,
		Side:          insertUserOrder.Side,
		PositionSide:  insertUserOrder.PositionSide,
		Quantity:      insertUserOrder.Quantity,
		Price:         insertUserOrder.Price,
		TraderQty:     insertUserOrder.TraderQty,
		OrderType:     insertUserOrder.OrderType,
		ClosePosition: insertUserOrder.ClosePosition,
		CumQuote:      insertUserOrder.CumQuote,
		ExecutedQty:   insertUserOrder.ExecutedQty,
		AvgPrice:      insertUserOrder.AvgPrice,
		CreatedAt:     insertUserOrder.CreatedAt,
		UpdatedAt:     insertUserOrder.UpdatedAt,
	}, nil
}

// InsertUserOrderTwo .
func (b *BinanceUserRepo) InsertUserOrderTwo(ctx context.Context, order *biz.UserOrder) (*biz.UserOrder, error) {
	insertUserOrder := &UserOrder{
		ID:            order.ID,
		UserId:        order.UserId,
		TraderId:      order.TraderId,
		ClientOrderId: order.ClientOrderId,
		OrderId:       order.OrderId,
		Symbol:        order.Symbol,
		Side:          order.Side,
		PositionSide:  order.PositionSide,
		Quantity:      order.Quantity,
		Price:         order.Price,
		TraderQty:     order.TraderQty,
		OrderType:     order.OrderType,
		ClosePosition: order.ClosePosition,
		CumQuote:      order.CumQuote,
		ExecutedQty:   order.ExecutedQty,
		AvgPrice:      order.AvgPrice,
		HandleStatus:  0,
	}

	res := b.data.DB(ctx).Table("new_user_order_two").Create(&insertUserOrder)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_ORDER_TWO_ERROR", "创建数据失败")
	}

	return &biz.UserOrder{
		ID:            insertUserOrder.ID,
		UserId:        insertUserOrder.UserId,
		TraderId:      insertUserOrder.TraderId,
		ClientOrderId: insertUserOrder.ClientOrderId,
		OrderId:       insertUserOrder.OrderId,
		Symbol:        insertUserOrder.Symbol,
		Side:          insertUserOrder.Side,
		PositionSide:  insertUserOrder.PositionSide,
		Quantity:      insertUserOrder.Quantity,
		Price:         insertUserOrder.Price,
		TraderQty:     insertUserOrder.TraderQty,
		OrderType:     insertUserOrder.OrderType,
		ClosePosition: insertUserOrder.ClosePosition,
		CumQuote:      insertUserOrder.CumQuote,
		ExecutedQty:   insertUserOrder.ExecutedQty,
		AvgPrice:      insertUserOrder.AvgPrice,
		CreatedAt:     insertUserOrder.CreatedAt,
		UpdatedAt:     insertUserOrder.UpdatedAt,
	}, nil
}

// InsertUserOrderErrTwo .
func (b *BinanceUserRepo) InsertUserOrderErrTwo(ctx context.Context, order *biz.UserOrderErr) (*biz.UserOrderErr, error) {
	insertUserOrder := &UserOrderErr{
		ID:            order.ID,
		UserId:        order.UserId,
		TraderId:      order.TraderId,
		ClientOrderId: order.ClientOrderId,
		OrderId:       order.OrderId,
		Symbol:        order.Symbol,
		Side:          order.Side,
		PositionSide:  order.PositionSide,
		Quantity:      order.Quantity,
		Price:         order.Price,
		TraderQty:     order.TraderQty,
		OrderType:     order.OrderType,
		ClosePosition: order.ClosePosition,
		CumQuote:      order.CumQuote,
		ExecutedQty:   order.ExecutedQty,
		AvgPrice:      order.AvgPrice,
		Code:          order.Code,
		Msg:           order.Msg,
		InitOrder:     order.InitOrder,
		Proportion:    order.Proportion,
		HandleStatus:  0,
	}

	res := b.data.DB(ctx).Table("new_user_order_err_two").Create(&insertUserOrder)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_ORDER_ERR_TWO_ERROR", "创建数据失败")
	}

	return &biz.UserOrderErr{
		ID:            insertUserOrder.ID,
		UserId:        insertUserOrder.UserId,
		TraderId:      insertUserOrder.TraderId,
		ClientOrderId: insertUserOrder.ClientOrderId,
		OrderId:       insertUserOrder.OrderId,
		Symbol:        insertUserOrder.Symbol,
		Side:          insertUserOrder.Side,
		PositionSide:  insertUserOrder.PositionSide,
		Quantity:      insertUserOrder.Quantity,
		Price:         insertUserOrder.Price,
		TraderQty:     insertUserOrder.TraderQty,
		OrderType:     insertUserOrder.OrderType,
		ClosePosition: insertUserOrder.ClosePosition,
		CumQuote:      insertUserOrder.CumQuote,
		ExecutedQty:   insertUserOrder.ExecutedQty,
		AvgPrice:      insertUserOrder.AvgPrice,
		InitOrder:     insertUserOrder.InitOrder,
		Code:          insertUserOrder.Code,
		Msg:           insertUserOrder.Msg,
		Proportion:    insertUserOrder.Proportion,
		CreatedAt:     insertUserOrder.CreatedAt,
		UpdatedAt:     insertUserOrder.UpdatedAt,
	}, nil
}

// UpdatesUserOrderHandleStatus .
func (b *BinanceUserRepo) UpdatesUserOrderHandleStatus(ctx context.Context, id uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_order").Where("id=? and handle_status<?", id, 1).
		Updates(map[string]interface{}{"handle_status": 1, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_ORDER_ERROR", "UPDATE_USER_ORDER_ERROR")
	}

	return true, nil
}

// UpdatesUserOrderTwoHandleStatus .
func (b *BinanceUserRepo) UpdatesUserOrderTwoHandleStatus(ctx context.Context, id uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_order_two").Where("id=? and handle_status<?", id, 1).
		Updates(map[string]interface{}{"handle_status": 1, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_ORDER_TWO_ERROR", "UPDATE_USER_ORDER_ERROR")
	}

	return true, nil
}

// InsertUserBindAfterUnbind .
func (b *BinanceUserRepo) InsertUserBindAfterUnbind(ctx context.Context, u *biz.UserBindAfterUnbind) (*biz.UserBindAfterUnbind, error) {
	insertUserBindAfterUnbind := &UserBindAfterUnbind{
		ID:           u.ID,
		UserId:       u.UserId,
		TraderId:     u.TraderId,
		Symbol:       u.Symbol,
		PositionSide: u.PositionSide,
		Quantity:     u.Quantity,
		Amount:       u.Amount,
		Status:       0,
	}

	res := b.data.DB(ctx).Table("new_user_bind_after_unbind").Create(&insertUserBindAfterUnbind)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_BIND_AFTER_UNBIND_ERROR", "创建数据失败")
	}

	return &biz.UserBindAfterUnbind{
		ID:           insertUserBindAfterUnbind.ID,
		UserId:       insertUserBindAfterUnbind.UserId,
		TraderId:     insertUserBindAfterUnbind.TraderId,
		Symbol:       insertUserBindAfterUnbind.Symbol,
		PositionSide: insertUserBindAfterUnbind.PositionSide,
		Quantity:     insertUserBindAfterUnbind.Quantity,
		Status:       insertUserBindAfterUnbind.Status,
		Amount:       insertUserBindAfterUnbind.Amount,
		CreatedAt:    insertUserBindAfterUnbind.CreatedAt,
		UpdatedAt:    insertUserBindAfterUnbind.UpdatedAt,
	}, nil
}

// UpdatesUserBindAfterUnbind .
func (b *BinanceUserRepo) UpdatesUserBindAfterUnbind(ctx context.Context, id uint64, status uint64, quantity float64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_bind_after_unbind").Where("id=? and status<?", id, 1).
		Updates(map[string]interface{}{"status": status, "quantity": quantity, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_AFTER_UNBIND_ERROR", "UPDATE_USER_AFTER_UNBIND_ERROR")
	}

	return true, nil
}

// UpdatesUserBindAfterUnbindTwo .
func (b *BinanceUserRepo) UpdatesUserBindAfterUnbindTwo(ctx context.Context, id uint64, status uint64, quantity float64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_bind_after_unbind_two").Where("id=? and status<?", id, 1).
		Updates(map[string]interface{}{"status": status, "quantity": quantity, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_AFTER_UNBIND_TWO_ERROR", "UPDATE_USER_AFTER_UNBIND_ERROR")
	}

	return true, nil
}

// GetUsers .
func (b *BinanceUserRepo) GetUsers() ([]*biz.User, error) {
	var users []*User
	if err := b.data.db.Table("new_user").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	res := make([]*biz.User, 0)
	for _, v := range users {
		res = append(res, &biz.User{
			ID:               v.ID,
			Address:          v.Address,
			ApiKey:           v.ApiKey,
			ApiStatus:        v.ApiStatus,
			BindTraderStatus: v.BindTraderStatus,
			ApiSecret:        v.ApiSecret,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUsersByUserIds .
func (b *BinanceUserRepo) GetUsersByUserIds(userIds []uint64) (map[uint64]*biz.User, error) {
	var users []*User
	if err := b.data.db.Table("new_user").Where("id in(?)", userIds).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	res := make(map[uint64]*biz.User, 0)
	for _, v := range users {
		res[v.ID] = &biz.User{
			ID:               v.ID,
			Address:          v.Address,
			ApiKey:           v.ApiKey,
			ApiStatus:        v.ApiStatus,
			BindTraderStatus: v.BindTraderStatus,
			ApiSecret:        v.ApiSecret,
			IsDai:            v.IsDai,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
			UseNewSystem:     v.UseNewSystem,
		}
	}

	return res, nil
}

// GetUsersByBindUserStatus .
func (b *BinanceUserRepo) GetUsersByBindUserStatus() ([]*biz.User, error) {
	var users []*User
	if err := b.data.db.Table("new_user").Where("bind_trader_status=?", 0).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	res := make([]*biz.User, 0)
	for _, v := range users {
		res = append(res, &biz.User{
			ID:               v.ID,
			Address:          v.Address,
			ApiStatus:        v.ApiStatus,
			ApiKey:           v.ApiKey,
			BindTraderStatus: v.BindTraderStatus,
			ApiSecret:        v.ApiSecret,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUsersByBindUserStatusReBind .
func (b *BinanceUserRepo) GetUsersByBindUserStatusReBind() ([]*biz.User, error) {
	var users []*User
	if err := b.data.db.Table("new_user").Where("bind_trader_status=?", 2).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	res := make([]*biz.User, 0)
	for _, v := range users {
		res = append(res, &biz.User{
			ID:                  v.ID,
			Address:             v.Address,
			ApiStatus:           v.ApiStatus,
			ApiKey:              v.ApiKey,
			BindTraderStatus:    v.BindTraderStatus,
			BindTraderStatusTfi: v.BindTraderStatusTfi,
			ApiSecret:           v.ApiSecret,
			CreatedAt:           v.CreatedAt,
			UpdatedAt:           v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserByAddress .
func (b *BinanceUserRepo) GetUserByAddress(ctx context.Context, address string) (*biz.User, error) {
	var user *User
	if err := b.data.DB(ctx).Table("new_user").Where("address=?", address).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	return &biz.User{
		ID:               user.ID,
		ApiStatus:        user.ApiStatus,
		Address:          user.Address,
		ApiKey:           user.ApiKey,
		BindTraderStatus: user.BindTraderStatus,
		ApiSecret:        user.ApiSecret,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
		IsDai:            user.IsDai,
	}, nil
}

// GetUserById .
func (b *BinanceUserRepo) GetUserById(ctx context.Context, userId uint64) (*biz.User, error) {
	var user *User
	if err := b.data.DB(ctx).Table("new_user").Where("id=?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	return &biz.User{
		ID:               user.ID,
		ApiStatus:        user.ApiStatus,
		Address:          user.Address,
		ApiKey:           user.ApiKey,
		BindTraderStatus: user.BindTraderStatus,
		ApiSecret:        user.ApiSecret,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
	}, nil
}

// GetUserByApiKeyAndApiSecret .
func (b *BinanceUserRepo) GetUserByApiKeyAndApiSecret(key string, secret string) (*biz.User, error) {
	var user *User
	if err := b.data.db.Table("new_user").Where("api_key=? or api_secret=?", key, secret).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	return &biz.User{
		ID:               user.ID,
		Address:          user.Address,
		ApiStatus:        user.ApiStatus,
		ApiKey:           user.ApiKey,
		BindTraderStatus: user.BindTraderStatus,
		ApiSecret:        user.ApiSecret,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
	}, nil
}

// GetUserBalance .
func (b *BinanceUserRepo) GetUserBalance(ctx context.Context, userId uint64) (*biz.UserBalance, error) {
	var userBalance *UserBalance
	if err := b.data.DB(ctx).Table("new_user_balance").Where("user_id=?", userId).First(&userBalance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BALANCE_ERROR", err.Error())
	}

	return &biz.UserBalance{
		ID:        userBalance.ID,
		UserId:    userBalance.UserId,
		Balance:   userBalance.Balance,
		Cost:      userBalance.Cost,
		Open:      userBalance.Open,
		CreatedAt: userBalance.CreatedAt,
		UpdatedAt: userBalance.UpdatedAt,
	}, nil
}

// GetUserBalanceTwo .
func (b *BinanceUserRepo) GetUserBalanceTwo(ctx context.Context, userId uint64) (*biz.UserBalance, error) {
	var userBalance *UserBalance
	if err := b.data.DB(ctx).Table("new_user_balance_two").Where("user_id=?", userId).First(&userBalance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BALANCE_TWO_ERROR", err.Error())
	}

	return &biz.UserBalance{
		ID:        userBalance.ID,
		UserId:    userBalance.UserId,
		Balance:   userBalance.Balance,
		Cost:      userBalance.Cost,
		Open:      userBalance.Open,
		CreatedAt: userBalance.CreatedAt,
		UpdatedAt: userBalance.UpdatedAt,
	}, nil
}

// GetUserAmount .
func (b *BinanceUserRepo) GetUserAmount(ctx context.Context, userId uint64) (*biz.UserAmount, error) {
	var userAmount *UserAmount
	if err := b.data.DB(ctx).Table("new_user_amount").Where("user_id=?", userId).First(&userAmount).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_AMOUNT_TWO_ERROR", err.Error())
	}

	return &biz.UserAmount{
		ID:        userAmount.ID,
		UserId:    userAmount.UserId,
		Amount:    userAmount.Amount,
		CreatedAt: userAmount.CreatedAt,
		UpdatedAt: userAmount.UpdatedAt,
	}, nil
}

// GetUserAmountTwo .
func (b *BinanceUserRepo) GetUserAmountTwo(ctx context.Context, userId uint64) (*biz.UserAmount, error) {
	var userAmount *UserAmount
	if err := b.data.DB(ctx).Table("new_user_amount_two").Where("user_id=?", userId).First(&userAmount).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_AMOUNT_TWO_ERROR", err.Error())
	}

	return &biz.UserAmount{
		ID:        userAmount.ID,
		UserId:    userAmount.UserId,
		Amount:    userAmount.Amount,
		CreatedAt: userAmount.CreatedAt,
		UpdatedAt: userAmount.UpdatedAt,
	}, nil
}

// GetUserBalanceByUserIds .
func (b *BinanceUserRepo) GetUserBalanceByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*biz.UserBalance, error) {
	var userBalance []*UserBalance
	if err := b.data.DB(ctx).Table("new_user_balance").Where("user_id in(?)", userIds).Find(&userBalance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BALANCE_ERROR", err.Error())
	}

	res := make(map[uint64]*biz.UserBalance, 0)
	for _, v := range userBalance {
		res[v.UserId] = &biz.UserBalance{
			ID:        v.ID,
			UserId:    v.UserId,
			Balance:   v.Balance,
			Cost:      v.Cost,
			Open:      v.Open,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return res, nil
}

// GetUserBalanceTwoByUserIds .
func (b *BinanceUserRepo) GetUserBalanceTwoByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*biz.UserBalance, error) {
	var userBalance []*UserBalance
	if err := b.data.DB(ctx).Table("new_user_balance_two").Where("user_id in(?)", userIds).Find(&userBalance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BALANCE_TWO_ERROR", err.Error())
	}

	res := make(map[uint64]*biz.UserBalance, 0)
	for _, v := range userBalance {
		res[v.UserId] = &biz.UserBalance{
			ID:        v.ID,
			UserId:    v.UserId,
			Balance:   v.Balance,
			Cost:      v.Cost,
			Open:      v.Open,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return res, nil
}

// GetUserAmountByUserIds .
func (b *BinanceUserRepo) GetUserAmountByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*biz.UserAmount, error) {
	var userAmount []*UserAmount
	if err := b.data.DB(ctx).Table("new_user_amount").Where("user_id in(?)", userIds).Find(&userAmount).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_AMOUNT_ERROR", err.Error())
	}

	res := make(map[uint64]*biz.UserAmount, 0)
	for _, v := range userAmount {
		res[v.UserId] = &biz.UserAmount{
			ID:        v.ID,
			UserId:    v.UserId,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return res, nil
}

// GetUserAmountTwoByUserIds .
func (b *BinanceUserRepo) GetUserAmountTwoByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*biz.UserAmount, error) {
	var userAmount []*UserAmount
	if err := b.data.DB(ctx).Table("new_user_amount_two").Where("user_id in(?)", userIds).Find(&userAmount).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_AMOUNT_TWO_ERROR", err.Error())
	}

	res := make(map[uint64]*biz.UserAmount, 0)
	for _, v := range userAmount {
		res[v.UserId] = &biz.UserAmount{
			ID:        v.ID,
			UserId:    v.UserId,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return res, nil
}

// GetTradersOrderByAmountDesc .
func (b *BinanceUserRepo) GetTradersOrderByAmountDesc() ([]*biz.Trader, error) {
	var traders []*Trader
	if err := b.data.db.Table("trader").Where("is_open=?", 1).Order("amount desc").Find(&traders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_TRADER_ERROR", err.Error())
	}

	res := make([]*biz.Trader, 0)
	for _, v := range traders {
		res = append(res, &biz.Trader{
			ID:        v.ID,
			Status:    v.IsOpen,
			Amount:    v.Amount,
			BaseMoney: v.BaseMoney,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

// GetTradersByTraderNum .
func (b *BinanceUserRepo) GetTradersByTraderNum() (map[string]*biz.Trader, error) {
	var traders []*Trader
	res := make(map[string]*biz.Trader, 0)
	if err := b.data.db.Table("trader").Where("is_open=?", 1).Find(&traders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "FIND_TRADER_ERROR", err.Error())
	}

	for _, v := range traders {
		res[v.PortfolioId] = &biz.Trader{
			ID:          v.ID,
			Status:      v.IsOpen,
			Amount:      v.Amount,
			BaseMoney:   v.BaseMoney,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			Name:        v.Name,
			PortfolioId: v.PortfolioId,
		}
	}

	return res, nil
}

// GetTraders .
func (b *BinanceUserRepo) GetTraders() (map[uint64]*biz.Trader, error) {
	var traders []*Trader
	if err := b.data.db.Table("trader").Find(&traders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_TRADER_ERROR", err.Error())
	}

	res := make(map[uint64]*biz.Trader, 0)
	for _, v := range traders {
		res[v.ID] = &biz.Trader{
			ID:          v.ID,
			Status:      v.IsOpen,
			Amount:      v.Amount,
			BaseMoney:   v.BaseMoney,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			Name:        v.Name,
			PortfolioId: v.PortfolioId,
		}
	}

	return res, nil
}

// GetUserBindTraderTwoById .
func (b *BinanceUserRepo) GetUserBindTraderTwoById(id uint64) (*biz.UserBindTrader, error) {
	var userBindTrader *UserBindTrader
	if err := b.data.db.Table("new_user_bind_trader_two").Where("id=?", id).First(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_ERROR", err.Error())
	}

	return &biz.UserBindTrader{
		ID:        userBindTrader.ID,
		UserId:    userBindTrader.UserId,
		TraderId:  userBindTrader.TraderId,
		Amount:    userBindTrader.Amount,
		Status:    userBindTrader.Status,
		InitOrder: userBindTrader.InitOrder,
		CreatedAt: userBindTrader.CreatedAt,
		UpdatedAt: userBindTrader.UpdatedAt,
	}, nil
}

// GetUserBindTraderTwoByUserId .
func (b *BinanceUserRepo) GetUserBindTraderTwoByUserId(userId uint64) ([]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	if err := b.data.db.Table("new_user_bind_trader_two").Where("user_id=?", userId).Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_ERROR", err.Error())
	}

	res := make([]*biz.UserBindTrader, 0)
	for _, v := range userBindTrader {
		res = append(res, &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			Amount:    v.Amount,
			InitOrder: v.InitOrder,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Status:    v.Status,
		})
	}

	return res, nil
}

// GetUserBindTraderByUserId .
func (b *BinanceUserRepo) GetUserBindTraderByUserId(userId uint64) ([]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	if err := b.data.db.Table("new_user_bind_trader").Where("user_id=?", userId).Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_ERROR", err.Error())
	}

	res := make([]*biz.UserBindTrader, 0)
	for _, v := range userBindTrader {
		res = append(res, &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Status:    v.Status,
		})
	}

	return res, nil
}

// GetUserBindTrader .
func (b *BinanceUserRepo) GetUserBindTrader() (map[uint64][]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	if err := b.data.db.Table("new_user_bind_trader").Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_ERROR", err.Error())
	}
	res := make(map[uint64][]*biz.UserBindTrader, 0)
	for _, v := range userBindTrader {
		if _, ok := res[v.UserId]; !ok {
			res[v.UserId] = make([]*biz.UserBindTrader, 0)
		}

		res[v.UserId] = append(res[v.UserId], &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Status:    v.Status,
		})
	}

	return res, nil
}

// GetUserBindTraderByStatus .
func (b *BinanceUserRepo) GetUserBindTraderByStatus(status uint64) (map[uint64][]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	if err := b.data.db.Table("new_user_bind_trader").Where("status=?", status).Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_ERROR", err.Error())
	}
	res := make(map[uint64][]*biz.UserBindTrader, 0)
	for _, v := range userBindTrader {
		if _, ok := res[v.TraderId]; !ok {
			res[v.TraderId] = make([]*biz.UserBindTrader, 0)
		}

		res[v.TraderId] = append(res[v.TraderId], &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			Amount:    v.Amount,
			Status:    v.Status,
			InitOrder: v.InitOrder,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserBindTraderByInitOrder .
func (b *BinanceUserRepo) GetUserBindTraderByInitOrder() (map[uint64][]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	if err := b.data.db.Table("new_user_bind_trader").Where("init_order=? and status=?", 0, 0).Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_ERROR", err.Error())
	}
	res := make(map[uint64][]*biz.UserBindTrader, 0)
	for _, v := range userBindTrader {
		if _, ok := res[v.TraderId]; !ok {
			res[v.TraderId] = make([]*biz.UserBindTrader, 0)
		}

		res[v.TraderId] = append(res[v.TraderId], &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			Amount:    v.Amount,
			Status:    v.Status,
			InitOrder: v.InitOrder,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserBindTraderByAlreadyInitOrder .
func (b *BinanceUserRepo) GetUserBindTraderByAlreadyInitOrder() (map[uint64][]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	if err := b.data.db.Table("new_user_bind_trader").Where("init_order=? and status=?", 1, 0).Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_ERROR", err.Error())
	}
	res := make(map[uint64][]*biz.UserBindTrader, 0)
	for _, v := range userBindTrader {
		if _, ok := res[v.TraderId]; !ok {
			res[v.TraderId] = make([]*biz.UserBindTrader, 0)
		}

		res[v.TraderId] = append(res[v.TraderId], &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			Amount:    v.Amount,
			Status:    v.Status,
			InitOrder: v.InitOrder,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserBindTraderByOverOrder .
func (b *BinanceUserRepo) GetUserBindTraderByOverOrder() (map[uint64][]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	if err := b.data.db.Table("new_user_bind_trader").Where("init_order=? and status=?", 2, 0).Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_ERROR", err.Error())
	}
	res := make(map[uint64][]*biz.UserBindTrader, 0)
	for _, v := range userBindTrader {
		if _, ok := res[v.TraderId]; !ok {
			res[v.TraderId] = make([]*biz.UserBindTrader, 0)
		}

		res[v.TraderId] = append(res[v.TraderId], &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			Amount:    v.Amount,
			Status:    v.Status,
			InitOrder: v.InitOrder,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserBindTraderTwoByInitOrder .
func (b *BinanceUserRepo) GetUserBindTraderTwoByInitOrder() (map[uint64][]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	if err := b.data.db.Table("new_user_bind_trader_two").Where("init_order=? and status=?", 0, 0).Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_TWO_ERROR", err.Error())
	}
	res := make(map[uint64][]*biz.UserBindTrader, 0)
	for _, v := range userBindTrader {
		if _, ok := res[v.TraderId]; !ok {
			res[v.TraderId] = make([]*biz.UserBindTrader, 0)
		}

		res[v.TraderId] = append(res[v.TraderId], &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			Amount:    v.Amount,
			Status:    v.Status,
			InitOrder: v.InitOrder,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserBindTraderTwoByAlreadyInitOrder .
func (b *BinanceUserRepo) GetUserBindTraderTwoByAlreadyInitOrder() (map[uint64][]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	if err := b.data.db.Table("new_user_bind_trader_two").Where("init_order=? and status=?", 1, 0).Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_TWO_ERROR", err.Error())
	}
	res := make(map[uint64][]*biz.UserBindTrader, 0)
	for _, v := range userBindTrader {
		if _, ok := res[v.TraderId]; !ok {
			res[v.TraderId] = make([]*biz.UserBindTrader, 0)
		}

		res[v.TraderId] = append(res[v.TraderId], &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			Amount:    v.Amount,
			Status:    v.Status,
			InitOrder: v.InitOrder,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserBindTraderTwoByOverOrder .
func (b *BinanceUserRepo) GetUserBindTraderTwoByOverOrder() (map[uint64][]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	if err := b.data.db.Table("new_user_bind_trader_two").Where("init_order=? and status=?", 2, 0).Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_TWO_ERROR", err.Error())
	}
	res := make(map[uint64][]*biz.UserBindTrader, 0)
	for _, v := range userBindTrader {
		if _, ok := res[v.TraderId]; !ok {
			res[v.TraderId] = make([]*biz.UserBindTrader, 0)
		}

		res[v.TraderId] = append(res[v.TraderId], &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			Amount:    v.Amount,
			Status:    v.Status,
			InitOrder: v.InitOrder,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserBindTraderMapByUserIds .
func (b *BinanceUserRepo) GetUserBindTraderMapByUserIds(userIds []uint64) (map[uint64][]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	if err := b.data.db.Table("new_user_bind_trader").Where("user_id in(?)", userIds).Order("amount desc").Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_ERROR", err.Error())
	}
	res := make(map[uint64][]*biz.UserBindTrader, 0)
	for _, v := range userBindTrader {
		if _, ok := res[v.UserId]; !ok {
			res[v.UserId] = make([]*biz.UserBindTrader, 0)
		}

		res[v.UserId] = append(res[v.UserId], &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Status:    v.Status,
			InitOrder: v.InitOrder,
		})
	}

	return res, nil
}

// GetUserBindTraderByTraderIds .
func (b *BinanceUserRepo) GetUserBindTraderByTraderIds(traderIds []uint64) (map[uint64][]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	if err := b.data.db.Table("new_user_bind_trader").Where("trader_id in(?) and status=?", traderIds, 0).Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_ERROR", err.Error())
	}

	res := make(map[uint64][]*biz.UserBindTrader, 0)
	for _, v := range userBindTrader {
		if _, ok := res[v.TraderId]; !ok {
			res[v.TraderId] = make([]*biz.UserBindTrader, 0)
		}

		res[v.TraderId] = append(res[v.TraderId], &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Status:    v.Status,
			InitOrder: v.InitOrder,
		})
	}

	return res, nil
}

// GetUserBindTraderTwoByTraderIds .
func (b *BinanceUserRepo) GetUserBindTraderTwoByTraderIds(traderIds []uint64) (map[uint64][]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	if err := b.data.db.Table("new_user_bind_trader_two").Where("trader_id in(?) and status=?", traderIds, 0).Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_ERROR", err.Error())
	}

	res := make(map[uint64][]*biz.UserBindTrader, 0)
	for _, v := range userBindTrader {
		if _, ok := res[v.TraderId]; !ok {
			res[v.TraderId] = make([]*biz.UserBindTrader, 0)
		}

		res[v.TraderId] = append(res[v.TraderId], &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Status:    v.Status,
			InitOrder: v.InitOrder,
		})
	}

	return res, nil
}

// GetUserBindAfterUnbindByUserIdAndTraderIdAndSymbol .
func (b *BinanceUserRepo) GetUserBindAfterUnbindByUserIdAndTraderIdAndSymbol(ctx context.Context, userId uint64, traderId uint64, symbol string, positionSide string) (*biz.UserBindAfterUnbind, error) {
	var userBindAfterUnbind *UserBindAfterUnbind
	if err := b.data.DB(ctx).Table("new_user_bind_after_unbind").
		Where("user_id=? and trader_id=? and status=? and symbol=? and position_side=?", userId, traderId, 0, symbol, positionSide).
		First(&userBindAfterUnbind).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_AFTER_UNBIND_ERROR", err.Error())
	}

	return &biz.UserBindAfterUnbind{
		ID:           userBindAfterUnbind.ID,
		UserId:       userBindAfterUnbind.UserId,
		TraderId:     userBindAfterUnbind.TraderId,
		Symbol:       userBindAfterUnbind.Symbol,
		PositionSide: userBindAfterUnbind.PositionSide,
		Quantity:     userBindAfterUnbind.Quantity,
		Status:       userBindAfterUnbind.Status,
		Amount:       userBindAfterUnbind.Amount,
		CreatedAt:    userBindAfterUnbind.CreatedAt,
		UpdatedAt:    userBindAfterUnbind.UpdatedAt,
	}, nil
}

// GetUserBindAfterUnbindByTraderIds .
func (b *BinanceUserRepo) GetUserBindAfterUnbindByTraderIds(traderIds []uint64) (map[uint64][]*biz.UserBindAfterUnbind, error) {
	var userBindAfterUnbind []*UserBindAfterUnbind
	if err := b.data.db.Table("new_user_bind_after_unbind").Where("trader_id in(?) and status=?", traderIds, 0).Find(&userBindAfterUnbind).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_AFTER_UNBIND_ERROR", err.Error())
	}

	res := make(map[uint64][]*biz.UserBindAfterUnbind, 0)
	for _, v := range userBindAfterUnbind {
		if _, ok := res[v.TraderId]; !ok {
			res[v.TraderId] = make([]*biz.UserBindAfterUnbind, 0)
		}

		res[v.TraderId] = append(res[v.TraderId], &biz.UserBindAfterUnbind{
			ID:           v.ID,
			UserId:       v.UserId,
			TraderId:     v.TraderId,
			Symbol:       v.Symbol,
			PositionSide: v.PositionSide,
			Quantity:     v.Quantity,
			Status:       v.Status,
			Amount:       v.Amount,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserBindAfterUnbindByStatus .
func (b *BinanceUserRepo) GetUserBindAfterUnbindByStatus() ([]*biz.UserBindAfterUnbind, error) {
	var userBindAfterUnbind []*UserBindAfterUnbind
	if err := b.data.db.Table("new_user_bind_after_unbind").Where("status=?", 0).Find(&userBindAfterUnbind).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_AFTER_UNBIND_ERROR", err.Error())
	}

	res := make([]*biz.UserBindAfterUnbind, 0)
	for _, v := range userBindAfterUnbind {
		res = append(res, &biz.UserBindAfterUnbind{
			ID:           v.ID,
			UserId:       v.UserId,
			TraderId:     v.TraderId,
			Symbol:       v.Symbol,
			PositionSide: v.PositionSide,
			Quantity:     v.Quantity,
			Status:       v.Status,
			Amount:       v.Amount,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserBindAfterUnbindTwoByStatus .
func (b *BinanceUserRepo) GetUserBindAfterUnbindTwoByStatus() ([]*biz.UserBindAfterUnbind, error) {
	var userBindAfterUnbind []*UserBindAfterUnbind
	if err := b.data.db.Table("new_user_bind_after_unbind_two").Where("status=?", 0).Find(&userBindAfterUnbind).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_AFTER_UNBIND_TWO_ERROR", err.Error())
	}

	res := make([]*biz.UserBindAfterUnbind, 0)
	for _, v := range userBindAfterUnbind {
		res = append(res, &biz.UserBindAfterUnbind{
			ID:           v.ID,
			UserId:       v.UserId,
			TraderId:     v.TraderId,
			Symbol:       v.Symbol,
			PositionSide: v.PositionSide,
			Quantity:     v.Quantity,
			Status:       v.Status,
			Amount:       v.Amount,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserOrderByUserTraderIdAndSymbol .
func (b *BinanceUserRepo) GetUserOrderByUserTraderIdAndSymbol(userId uint64, traderId uint64, symbol string) ([]*biz.UserOrder, error) {
	var userOrder []*UserOrder
	if err := b.data.db.Table("new_user_order").
		Where("user_id=? and trader_id=? and symbol=?", userId, traderId, symbol).
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_ERROR", err.Error())
	}

	res := make([]*biz.UserOrder, 0)
	for _, v := range userOrder {
		res = append(res, &biz.UserOrder{
			ID:            v.ID,
			UserId:        v.UserId,
			TraderId:      v.TraderId,
			ClientOrderId: v.ClientOrderId,
			OrderId:       v.OrderId,
			Symbol:        v.Symbol,
			Side:          v.Side,
			PositionSide:  v.PositionSide,
			Quantity:      v.Quantity,
			Price:         v.Price,
			TraderQty:     v.TraderQty,
			OrderType:     v.OrderType,
			ClosePosition: v.ClosePosition,
			CumQuote:      v.CumQuote,
			ExecutedQty:   v.ExecutedQty,
			AvgPrice:      v.AvgPrice,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserOrders .
func (b *BinanceUserRepo) GetUserOrders() ([]*biz.UserOrder, error) {
	var userOrder []*UserOrder
	if err := b.data.db.Table("new_user_order_two").Order("id asc").
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_TWO_ERROR", err.Error())
	}

	res := make([]*biz.UserOrder, 0)
	for _, v := range userOrder {
		res = append(res, &biz.UserOrder{
			ID:            v.ID,
			UserId:        v.UserId,
			TraderId:      v.TraderId,
			ClientOrderId: v.ClientOrderId,
			OrderId:       v.OrderId,
			Symbol:        v.Symbol,
			Side:          v.Side,
			PositionSide:  v.PositionSide,
			Quantity:      v.Quantity,
			Price:         v.Price,
			TraderQty:     v.TraderQty,
			OrderType:     v.OrderType,
			ClosePosition: v.ClosePosition,
			CumQuote:      v.CumQuote,
			ExecutedQty:   v.ExecutedQty,
			AvgPrice:      v.AvgPrice,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserOrderTwoByUserTraderIdAndSymbol .
func (b *BinanceUserRepo) GetUserOrderTwoByUserTraderIdAndSymbol(userId uint64, traderId uint64, symbol string) ([]*biz.UserOrder, error) {
	var userOrder []*UserOrder
	if err := b.data.db.Table("new_user_order_two").
		Where("user_id=? and trader_id=? and symbol=?", userId, traderId, symbol).
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_TWO_ERROR", err.Error())
	}

	res := make([]*biz.UserOrder, 0)
	for _, v := range userOrder {
		res = append(res, &biz.UserOrder{
			ID:            v.ID,
			UserId:        v.UserId,
			TraderId:      v.TraderId,
			ClientOrderId: v.ClientOrderId,
			OrderId:       v.OrderId,
			Symbol:        v.Symbol,
			Side:          v.Side,
			PositionSide:  v.PositionSide,
			Quantity:      v.Quantity,
			Price:         v.Price,
			TraderQty:     v.TraderQty,
			OrderType:     v.OrderType,
			ClosePosition: v.ClosePosition,
			CumQuote:      v.CumQuote,
			ExecutedQty:   v.ExecutedQty,
			AvgPrice:      v.AvgPrice,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserOrderTwoByUserTraderIdAndSymbolNew .
func (b *BinanceUserRepo) GetUserOrderTwoByUserTraderIdAndSymbolNew(userId uint64, traderId uint64, symbol string) ([]*biz.UserOrder, error) {
	var userOrder []*UserOrder
	if err := b.data.db.Table("new_user_order_two_"+strconv.FormatInt(int64(userId), 10)).
		Where("user_id=? and trader_id=? and symbol=?", userId, traderId, symbol).
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_TWO_ERROR", err.Error())
	}

	res := make([]*biz.UserOrder, 0)
	for _, v := range userOrder {
		res = append(res, &biz.UserOrder{
			ID:            v.ID,
			UserId:        v.UserId,
			TraderId:      v.TraderId,
			ClientOrderId: v.ClientOrderId,
			OrderId:       v.OrderId,
			Symbol:        v.Symbol,
			Side:          v.Side,
			PositionSide:  v.PositionSide,
			Quantity:      v.Quantity,
			Price:         v.Price,
			TraderQty:     v.TraderQty,
			OrderType:     v.OrderType,
			ClosePosition: v.ClosePosition,
			CumQuote:      v.CumQuote,
			ExecutedQty:   v.ExecutedQty,
			AvgPrice:      v.AvgPrice,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserOrderByUserTraderId .
func (b *BinanceUserRepo) GetUserOrderByUserTraderId(userId uint64, traderId uint64) (map[string][]*biz.UserOrder, error) {
	var userOrder []*UserOrder
	if err := b.data.db.Table("new_user_order").
		Where("user_id=? and trader_id=?", userId, traderId).
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_ERROR", err.Error())
	}

	res := make(map[string][]*biz.UserOrder, 0)
	for _, v := range userOrder {
		if _, ok := res[v.Symbol]; !ok {
			res[v.Symbol] = make([]*biz.UserOrder, 0)

		}

		res[v.Symbol] = append(res[v.Symbol], &biz.UserOrder{
			ID:            v.ID,
			UserId:        v.UserId,
			TraderId:      v.TraderId,
			ClientOrderId: v.ClientOrderId,
			OrderId:       v.OrderId,
			Symbol:        v.Symbol,
			Side:          v.Side,
			PositionSide:  v.PositionSide,
			Quantity:      v.Quantity,
			Price:         v.Price,
			TraderQty:     v.TraderQty,
			OrderType:     v.OrderType,
			ClosePosition: v.ClosePosition,
			CumQuote:      v.CumQuote,
			ExecutedQty:   v.ExecutedQty,
			AvgPrice:      v.AvgPrice,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserOrderByUserIdAndSymbolAndPositionSide .
func (b *BinanceUserRepo) GetUserOrderByUserIdAndSymbolAndPositionSide(userId uint64, symbol string, positionSide string) ([]*biz.UserOrder, error) {
	var userOrder []*UserOrder
	if err := b.data.db.Table("new_user_order").
		Where("user_id=? and symbol=? and position_side=?", userId, symbol, positionSide).
		Order("id asc").
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_ERROR", err.Error())
	}

	res := make([]*biz.UserOrder, 0)
	for _, v := range userOrder {
		res = append(res, &biz.UserOrder{
			ID:            v.ID,
			UserId:        v.UserId,
			TraderId:      v.TraderId,
			ClientOrderId: v.ClientOrderId,
			OrderId:       v.OrderId,
			Symbol:        v.Symbol,
			Side:          v.Side,
			PositionSide:  v.PositionSide,
			Quantity:      v.Quantity,
			Price:         v.Price,
			TraderQty:     v.TraderQty,
			OrderType:     v.OrderType,
			ClosePosition: v.ClosePosition,
			CumQuote:      v.CumQuote,
			ExecutedQty:   v.ExecutedQty,
			AvgPrice:      v.AvgPrice,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserOrderById .
func (b *BinanceUserRepo) GetUserOrderById(orderId int64) (*biz.UserOrder, error) {
	var userOrder *UserOrder
	if err := b.data.db.Table("new_user_order").
		Where("id=?", orderId).
		First(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_ERROR", err.Error())
	}

	return &biz.UserOrder{
		ID:            userOrder.ID,
		UserId:        userOrder.UserId,
		TraderId:      userOrder.TraderId,
		ClientOrderId: userOrder.ClientOrderId,
		OrderId:       userOrder.OrderId,
		Symbol:        userOrder.Symbol,
		Side:          userOrder.Side,
		PositionSide:  userOrder.PositionSide,
		Quantity:      userOrder.Quantity,
		Price:         userOrder.Price,
		TraderQty:     userOrder.TraderQty,
		OrderType:     userOrder.OrderType,
		ClosePosition: userOrder.ClosePosition,
		CumQuote:      userOrder.CumQuote,
		ExecutedQty:   userOrder.ExecutedQty,
		AvgPrice:      userOrder.AvgPrice,
		HandleStatus:  userOrder.HandleStatus,
		CreatedAt:     userOrder.CreatedAt,
		UpdatedAt:     userOrder.UpdatedAt,
	}, nil
}

// GetUserOrderByHandleStatus .
func (b *BinanceUserRepo) GetUserOrderByHandleStatus() ([]*biz.UserOrder, error) {
	var userOrder []*UserOrder
	if err := b.data.db.Table("new_user_order").
		Where("((side=? and position_side=?) or (side=? and position_side=?)) and handle_status=?", "SELL", "LONG", "BUY", "SHORT", 0).
		Order("id asc").
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_ERROR", err.Error())
	}

	res := make([]*biz.UserOrder, 0)
	for _, v := range userOrder {
		res = append(res, &biz.UserOrder{
			ID:            v.ID,
			UserId:        v.UserId,
			TraderId:      v.TraderId,
			ClientOrderId: v.ClientOrderId,
			OrderId:       v.OrderId,
			Symbol:        v.Symbol,
			Side:          v.Side,
			PositionSide:  v.PositionSide,
			Quantity:      v.Quantity,
			Price:         v.Price,
			TraderQty:     v.TraderQty,
			OrderType:     v.OrderType,
			ClosePosition: v.ClosePosition,
			CumQuote:      v.CumQuote,
			ExecutedQty:   v.ExecutedQty,
			AvgPrice:      v.AvgPrice,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserOrderByUserIdGroupBySymbol .
func (b *BinanceUserRepo) GetUserOrderByUserIdGroupBySymbol(userId uint64) ([]*biz.UserOrder, error) {
	var userOrder []*UserOrder
	if err := b.data.db.Table("new_user_order").
		Where("user_id=?", userId).
		Group("symbol").
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_ERROR", err.Error())
	}

	res := make([]*biz.UserOrder, 0)
	for _, v := range userOrder {
		res = append(res, &biz.UserOrder{
			ID:            v.ID,
			UserId:        v.UserId,
			TraderId:      v.TraderId,
			ClientOrderId: v.ClientOrderId,
			OrderId:       v.OrderId,
			Symbol:        v.Symbol,
			Side:          v.Side,
			PositionSide:  v.PositionSide,
			Quantity:      v.Quantity,
			Price:         v.Price,
			TraderQty:     v.TraderQty,
			OrderType:     v.OrderType,
			ClosePosition: v.ClosePosition,
			CumQuote:      v.CumQuote,
			ExecutedQty:   v.ExecutedQty,
			AvgPrice:      v.AvgPrice,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserOrderByIds .
func (b *BinanceUserRepo) GetUserOrderByIds(ids []int64) ([]*biz.UserOrder, error) {
	var userOrder []*UserOrder
	if err := b.data.db.Table("new_user_order").
		Where("id in(?)", ids).
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_ERROR", err.Error())
	}

	res := make([]*biz.UserOrder, 0)
	for _, v := range userOrder {
		res = append(res, &biz.UserOrder{
			ID:            v.ID,
			UserId:        v.UserId,
			TraderId:      v.TraderId,
			ClientOrderId: v.ClientOrderId,
			OrderId:       v.OrderId,
			Symbol:        v.Symbol,
			Side:          v.Side,
			PositionSide:  v.PositionSide,
			Quantity:      v.Quantity,
			Price:         v.Price,
			TraderQty:     v.TraderQty,
			OrderType:     v.OrderType,
			ClosePosition: v.ClosePosition,
			CumQuote:      v.CumQuote,
			ExecutedQty:   v.ExecutedQty,
			AvgPrice:      v.AvgPrice,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
			HandleStatus:  v.HandleStatus,
		})
	}

	return res, nil
}

// GetUserOrderTwoByIds .
func (b *BinanceUserRepo) GetUserOrderTwoByIds(ids []int64) ([]*biz.UserOrder, error) {
	var userOrder []*UserOrder
	if err := b.data.db.Table("new_user_order_two").
		Where("id in(?)", ids).
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_TWO_ERROR", err.Error())
	}

	res := make([]*biz.UserOrder, 0)
	for _, v := range userOrder {
		res = append(res, &biz.UserOrder{
			ID:            v.ID,
			UserId:        v.UserId,
			TraderId:      v.TraderId,
			ClientOrderId: v.ClientOrderId,
			OrderId:       v.OrderId,
			Symbol:        v.Symbol,
			Side:          v.Side,
			PositionSide:  v.PositionSide,
			Quantity:      v.Quantity,
			Price:         v.Price,
			TraderQty:     v.TraderQty,
			OrderType:     v.OrderType,
			ClosePosition: v.ClosePosition,
			CumQuote:      v.CumQuote,
			ExecutedQty:   v.ExecutedQty,
			AvgPrice:      v.AvgPrice,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
			HandleStatus:  v.HandleStatus,
		})
	}

	return res, nil
}

// GetUserOrderByUserIdMapId .
func (b *BinanceUserRepo) GetUserOrderByUserIdMapId(userId uint64) (map[string]*biz.UserOrder, error) {
	var userOrder []*UserOrder
	if err := b.data.db.Table("new_user_order").
		Where("user_id=?", userId).
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_ERROR", err.Error())
	}

	res := make(map[string]*biz.UserOrder, 0)
	for _, v := range userOrder {
		res[v.OrderId] = &biz.UserOrder{
			ID:            v.ID,
			UserId:        v.UserId,
			TraderId:      v.TraderId,
			ClientOrderId: v.ClientOrderId,
			OrderId:       v.OrderId,
			Symbol:        v.Symbol,
			Side:          v.Side,
			PositionSide:  v.PositionSide,
			Quantity:      v.Quantity,
			Price:         v.Price,
			TraderQty:     v.TraderQty,
			OrderType:     v.OrderType,
			ClosePosition: v.ClosePosition,
			CumQuote:      v.CumQuote,
			ExecutedQty:   v.ExecutedQty,
			AvgPrice:      v.AvgPrice,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
			HandleStatus:  v.HandleStatus,
		}
	}

	return res, nil
}

// GetSymbol .
func (b *BinanceUserRepo) GetSymbol() (map[string]*biz.Symbol, error) {
	var lhCoinSymbol []*LhCoinSymbol
	if err := b.data.db.Table("lh_coin_symbol").Find(&lhCoinSymbol).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_LH_COIN_SYMBOL_ERROR", err.Error())
	}

	res := make(map[string]*biz.Symbol, 0)
	for _, v := range lhCoinSymbol {
		res[v.Symbol+"USDT"] = &biz.Symbol{
			ID:                v.ID,
			Symbol:            v.Symbol,
			QuantityPrecision: v.QuantityPrecision,
			PricePrecision:    v.PricePrecision,
		}
	}

	return res, nil
}

// GetTraderPosition .
func (b *BinanceUserRepo) GetTraderPosition(traderId uint64) ([]*biz.TraderPosition, error) {
	var lhTraderPosition []*LhTraderPosition
	if err := b.data.db.Table("lh_trader_position").Where("trader_id=?", traderId).Find(&lhTraderPosition).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_LH_TRADER_POSITION_ERROR", err.Error())
	}

	res := make([]*biz.TraderPosition, 0)
	for _, v := range lhTraderPosition {
		res = append(res, &biz.TraderPosition{
			ID:           v.ID,
			TraderId:     v.TraderId,
			Symbol:       v.Symbol,
			Qty:          v.Qty,
			Side:         v.Side,
			PositionSide: v.Type,
		})
	}

	return res, nil
}

// GetOpeningTraderPositionNewNew .
func (b *BinanceUserRepo) GetOpeningTraderPositionNewNew(traderNum string) ([]*biz.ZyTraderPosition, error) {
	var lhTraderPosition []*TraderPosition

	res := make([]*biz.ZyTraderPosition, 0)
	tableName := "trader_position_" + traderNum
	if err := b.data.db.Table(tableName).Where("position_amount>?", 0).Find(&lhTraderPosition).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "FIND_LH_TRADER_POSITION_ERROR", err.Error())
	}

	for _, v := range lhTraderPosition {
		res = append(res, &biz.ZyTraderPosition{
			ID:             v.ID,
			Symbol:         v.Symbol,
			PositionAmount: v.PositionAmount,
			PositionSide:   v.PositionSide,
		})
	}

	return res, nil
}

// GetOpeningTraderPositionNewNewMap .
func (b *BinanceUserRepo) GetOpeningTraderPositionNewNewMap(traderNum string) (map[uint64]*biz.ZyTraderPosition, error) {
	var lhTraderPosition []*TraderPosition

	res := make(map[uint64]*biz.ZyTraderPosition, 0)
	tableName := "trader_position_" + traderNum
	if err := b.data.db.Table(tableName).Where("position_amount>?", 0).Find(&lhTraderPosition).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "FIND_LH_TRADER_POSITION_ERROR", err.Error())
	}

	for _, v := range lhTraderPosition {
		res[v.ID] = &biz.ZyTraderPosition{
			ID:             v.ID,
			Symbol:         v.Symbol,
			PositionAmount: v.PositionAmount,
			PositionSide:   v.PositionSide,
		}
	}

	return res, nil
}

// GetTraderEmail .
func (b *BinanceUserRepo) GetTraderEmail(traderId uint64) (map[uint64][]*biz.TraderEmail, error) {
	var traderEmail []*TraderEmail

	res := make(map[uint64][]*biz.TraderEmail, 0)
	tableName := "trader_email"
	if err := b.data.db.Table(tableName).Where("trader_id=?", traderId).Find(&traderEmail).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "FIND_LH_TRADER_POSITION_ERROR", err.Error())
	}

	for _, v := range traderEmail {
		if _, ok := res[v.TraderId]; !ok {
			res[v.TraderId] = make([]*biz.TraderEmail, 0)
		}

		res[v.TraderId] = append(res[v.TraderId], &biz.TraderEmail{
			ID:        v.ID,
			Email:     v.Email,
			CreatedAt: v.CreatedAt,
			TraderId:  v.TraderId,
		})
	}

	return res, nil
}

// GetOpeningTraderPosition .
func (b *BinanceUserRepo) GetOpeningTraderPosition(traderNum string) ([]*biz.ZyTraderPosition, error) {
	var lhTraderPosition []*ZyTraderPosition

	res := make([]*biz.ZyTraderPosition, 0)
	tableName := "zy_trader_position" + traderNum
	if err := b.data.db.Table(tableName).Where("positionAmount>?", 0).Find(&lhTraderPosition).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "FIND_LH_TRADER_POSITION_ERROR", err.Error())
	}

	for _, v := range lhTraderPosition {
		res = append(res, &biz.ZyTraderPosition{
			ID:             v.ID,
			Symbol:         v.Symbol,
			PositionAmount: v.PositionAmount,
			PositionSide:   v.PositionSide,
		})
	}

	return res, nil
}

// GetOpeningTraderPositionNew .
func (b *BinanceUserRepo) GetOpeningTraderPositionNew(traderNum string) ([]*biz.TraderPositionNew, error) {
	var lhTraderPosition []*TraderPositionNew

	res := make([]*biz.TraderPositionNew, 0)
	tableName := "new_binance_position_" + traderNum + "_history"
	if err := b.data.db.Table(tableName).Where("closed=?", 0).Find(&lhTraderPosition).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "FIND_LH_TRADER_POSITION_ERROR", err.Error())
	}

	for _, v := range lhTraderPosition {
		res = append(res, &biz.TraderPositionNew{
			ID:     v.ID,
			Symbol: v.Symbol,
			Qty:    v.Qty,
			Side:   v.Side,
			Closed: v.Closed,
			Opened: v.Opened,
		})
	}

	return res, nil
}

// InsertTradingBoxOpen .
func (b *BinanceUserRepo) InsertTradingBoxOpen(ctx context.Context, box *biz.TradingBoxOpen) (*biz.TradingBoxOpen, error) {
	insertBox := &TradingBoxOpen{
		TokenId:        box.TokenId,
		Amount:         box.Amount,
		AmountTotal:    box.AmountTotal,
		WithdrawStatus: box.WithdrawStatus,
		AmountRate:     box.AmountRate,
		Total:          box.Total,
		WithdrawAmount: box.WithdrawAmount,
	}

	res := b.data.DB(ctx).Table("new_trading_box_open").Create(&insertBox)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_TRADING_BOX_OPEN_ERROR", "创建数据失败")
	}

	return &biz.TradingBoxOpen{
		ID:             box.ID,
		TokenId:        box.TokenId,
		Amount:         box.Amount,
		AmountTotal:    box.AmountTotal,
		WithdrawStatus: box.WithdrawStatus,
		AmountRate:     box.AmountRate,
		Total:          box.Total,
		WithdrawAmount: box.WithdrawAmount,
		CreatedAt:      box.CreatedAt,
		UpdatedAt:      box.UpdatedAt,
	}, nil
}

// UpdateTradingBoxOpenStatus .
func (b *BinanceUserRepo) UpdateTradingBoxOpenStatus(ctx context.Context, tokenId uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_trading_box_open").Where("token_id=? and withdraw_status=?", tokenId, 0).
		Updates(map[string]interface{}{"withdraw_status": 1, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_TRADING_BOX_OPEN_ERROR", "UPDATE_TRADING_BOX_ERROR")
	}

	return true, nil
}

// GetTradingBoxOpen .
func (b *BinanceUserRepo) GetTradingBoxOpen() ([]*biz.TradingBoxOpen, error) {
	var tradingBoxOpen []*TradingBoxOpen
	if err := b.data.db.Table("new_trading_box_open").Find(&tradingBoxOpen).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_TRADING_BOX_OPEN_ERROR", err.Error())
	}

	res := make([]*biz.TradingBoxOpen, 0)
	for _, box := range tradingBoxOpen {
		res = append(res, &biz.TradingBoxOpen{
			ID:             box.ID,
			TokenId:        box.TokenId,
			Amount:         box.Amount,
			AmountTotal:    box.AmountTotal,
			WithdrawStatus: box.WithdrawStatus,
			AmountRate:     box.AmountRate,
			Total:          box.Total,
			WithdrawAmount: box.WithdrawAmount,
			CreatedAt:      box.CreatedAt,
			UpdatedAt:      box.UpdatedAt,
		})
	}

	return res, nil
}

// GetTradingBoxOpenMap .
func (b *BinanceUserRepo) GetTradingBoxOpenMap() (map[uint64]*biz.TradingBoxOpen, error) {
	var tradingBoxOpen []*TradingBoxOpen
	if err := b.data.db.Table("new_trading_box_open").Find(&tradingBoxOpen).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_TRADING_BOX_OPEN_ERROR", err.Error())
	}

	res := make(map[uint64]*biz.TradingBoxOpen, 0)
	for _, box := range tradingBoxOpen {
		res[box.TokenId] = &biz.TradingBoxOpen{
			ID:             box.ID,
			TokenId:        box.TokenId,
			Amount:         box.Amount,
			AmountTotal:    box.AmountTotal,
			WithdrawStatus: box.WithdrawStatus,
			AmountRate:     box.AmountRate,
			Total:          box.Total,
			WithdrawAmount: box.WithdrawAmount,
			CreatedAt:      box.CreatedAt,
			UpdatedAt:      box.UpdatedAt,
		}
	}

	return res, nil
}

// GetTradingBoxOpenMapByStatus .
func (b *BinanceUserRepo) GetTradingBoxOpenMapByStatus(status uint64) (map[uint64]*biz.TradingBoxOpen, error) {
	var tradingBoxOpen []*TradingBoxOpen
	if err := b.data.db.Table("new_trading_box_open").Where("withdraw_status=?", status).Find(&tradingBoxOpen).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_TRADING_BOX_OPEN_ERROR", err.Error())
	}

	res := make(map[uint64]*biz.TradingBoxOpen, 0)
	for _, box := range tradingBoxOpen {
		res[box.TokenId] = &biz.TradingBoxOpen{
			ID:             box.ID,
			TokenId:        box.TokenId,
			Amount:         box.Amount,
			AmountTotal:    box.AmountTotal,
			WithdrawStatus: box.WithdrawStatus,
			AmountRate:     box.AmountRate,
			Total:          box.Total,
			WithdrawAmount: box.WithdrawAmount,
			CreatedAt:      box.CreatedAt,
			UpdatedAt:      box.UpdatedAt,
		}
	}

	return res, nil
}

// GetBinanceTrade .
func (b *BinanceUserRepo) GetBinanceTrade() ([]*biz.BinanceTrade, error) {
	var binanceTrade []*BinanceTrade
	if err := b.data.db.Table("new_binance_trader").Where("status=?", 0).Find(&binanceTrade).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_BINANCE_TRADER_ERROR", err.Error())
	}

	res := make([]*biz.BinanceTrade, 0)
	for _, v := range binanceTrade {
		res = append(res, &biz.BinanceTrade{
			ID:        v.ID,
			TraderNum: v.TraderNum,
			Status:    v.Status,
		})
	}

	return res, nil
}

// GetBinanceTradeHistory .
func (b *BinanceUserRepo) GetBinanceTradeHistory(traderNum uint64) ([]*biz.BinanceTradeHistory, error) {
	var binanceTradeHistory []*BinanceTradeHistory
	if err := b.data.db.Table("new_binance_trade_history").Where("trader_num=?", traderNum).Find(&binanceTradeHistory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_BINANCE_TRADER_HISTORY_ERROR", err.Error())
	}

	res := make([]*biz.BinanceTradeHistory, 0)
	for _, v := range binanceTradeHistory {
		res = append(res, &biz.BinanceTradeHistory{
			ID:                  v.ID,
			Time:                v.Time,
			Symbol:              v.Symbol,
			Side:                v.Side,
			Price:               v.Price,
			Fee:                 v.Fee,
			FeeAsset:            v.FeeAsset,
			Quantity:            v.Quantity,
			QuantityAsset:       v.QuantityAsset,
			RealizedProfit:      v.RealizedProfit,
			RealizedProfitAsset: v.RealizedProfitAsset,
			BaseAsset:           v.BaseAsset,
			Qty:                 v.Qty,
			PositionSide:        v.PositionSide,
			ActiveBuy:           v.ActiveBuy,
			CreatedAt:           v.CreatedAt,
			UpdatedAt:           v.UpdatedAt,
		})
	}

	return res, nil
}

// GetBinanceTradeHistoryByTraderNumNewest .
func (b *BinanceUserRepo) GetBinanceTradeHistoryByTraderNumNewest(traderNum uint64, limit int) ([]*biz.BinanceTradeHistory, error) {
	var binanceTradeHistory []*BinanceTradeHistory

	res := make([]*biz.BinanceTradeHistory, 0)

	tableName := "new_binance_trade_" + strconv.FormatInt(int64(traderNum), 10) + "_history"
	if err := b.data.db.Table(tableName).Order("id desc").Limit(limit).Find(&binanceTradeHistory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return res, errors.New(500, "FIND_BINANCE_TRADER_HISTORY_ERROR", err.Error())
	}

	for _, v := range binanceTradeHistory {
		res = append(res, &biz.BinanceTradeHistory{
			ID:                  v.ID,
			Time:                v.Time,
			Symbol:              v.Symbol,
			Side:                v.Side,
			Price:               v.Price,
			Fee:                 v.Fee,
			FeeAsset:            v.FeeAsset,
			Quantity:            v.Quantity,
			QuantityAsset:       v.QuantityAsset,
			RealizedProfit:      v.RealizedProfit,
			RealizedProfitAsset: v.RealizedProfitAsset,
			BaseAsset:           v.BaseAsset,
			Qty:                 v.Qty,
			PositionSide:        v.PositionSide,
			ActiveBuy:           v.ActiveBuy,
			CreatedAt:           v.CreatedAt,
			UpdatedAt:           v.UpdatedAt,
		})
	}
	return res, nil
}

// InsertBinanceTradeHistory .
func (b *BinanceUserRepo) InsertBinanceTradeHistory(ctx context.Context, history *biz.BinanceTradeHistory, traderNum uint64) (*biz.BinanceTradeHistory, error) {
	insert := &BinanceTradeHistory{
		Time:                history.Time,
		Symbol:              history.Symbol,
		Side:                history.Side,
		Price:               history.Price,
		Fee:                 history.Fee,
		FeeAsset:            history.FeeAsset,
		Quantity:            history.Quantity,
		QuantityAsset:       history.QuantityAsset,
		RealizedProfit:      history.RealizedProfit,
		RealizedProfitAsset: history.RealizedProfitAsset,
		BaseAsset:           history.BaseAsset,
		Qty:                 history.Qty,
		PositionSide:        history.PositionSide,
		ActiveBuy:           history.ActiveBuy,
	}

	tableName := "new_binance_trade_" + strconv.FormatInt(int64(traderNum), 10) + "_history"
	res := b.data.DB(ctx).Table(tableName).Create(&insert)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_BINANCE_TRADER_HISTORY_ERROR", "创建数据失败")
	}

	return &biz.BinanceTradeHistory{
		ID:                  insert.ID,
		Time:                insert.Time,
		Symbol:              insert.Symbol,
		Side:                insert.Side,
		Price:               insert.Price,
		Fee:                 insert.Fee,
		FeeAsset:            insert.FeeAsset,
		Quantity:            insert.Quantity,
		QuantityAsset:       insert.QuantityAsset,
		RealizedProfit:      insert.RealizedProfit,
		RealizedProfitAsset: insert.RealizedProfitAsset,
		BaseAsset:           insert.BaseAsset,
		Qty:                 insert.Qty,
		PositionSide:        insert.PositionSide,
		ActiveBuy:           insert.ActiveBuy,
		CreatedAt:           insert.CreatedAt,
		UpdatedAt:           insert.UpdatedAt,
	}, nil
}

// GetBinancePositionHistory .
func (b *BinanceUserRepo) GetBinancePositionHistory(traderNum uint64) ([]*biz.BinancePositionHistory, error) {
	var binancePositionHistory []*BinancePositionHistory
	if err := b.data.db.Table("new_binance_position_history").Where("trader_num=?", traderNum).Find(&binancePositionHistory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_BINANCE_POSITION_HISTORY_ERROR", err.Error())
	}

	res := make([]*biz.BinancePositionHistory, 0)
	for _, v := range binancePositionHistory {
		res = append(res, &biz.BinancePositionHistory{
			ID:              v.ID,
			Symbol:          v.Symbol,
			Side:            v.Side,
			Closed:          v.Closed,
			Opened:          v.Opened,
			AvgCost:         v.AvgCost,
			AvgClosePrice:   v.AvgClosePrice,
			ClosingPnl:      v.ClosingPnl,
			MaxOpenInterest: v.MaxOpenInterest,
			ClosedVolume:    v.ClosedVolume,
			Status:          v.Status,
			CreatedAt:       v.CreatedAt,
			UpdatedAt:       v.UpdatedAt,
		})
	}

	return res, nil
}

// GetBinancePositionHistoryByTraderNumNewest .
func (b *BinanceUserRepo) GetBinancePositionHistoryByTraderNumNewest(traderNum uint64) (*biz.BinancePositionHistory, error) {
	var binancePositionHistory *BinancePositionHistory
	if err := b.data.db.Table("new_binance_position_history").Where("trader_num=?", traderNum).Order("id desc").First(&binancePositionHistory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_BINANCE_POSITION_HISTORY_ERROR", err.Error())
	}

	return &biz.BinancePositionHistory{
		ID:              binancePositionHistory.ID,
		Symbol:          binancePositionHistory.Symbol,
		Side:            binancePositionHistory.Side,
		Closed:          binancePositionHistory.Closed,
		Opened:          binancePositionHistory.Opened,
		AvgCost:         binancePositionHistory.AvgCost,
		AvgClosePrice:   binancePositionHistory.AvgClosePrice,
		ClosingPnl:      binancePositionHistory.ClosingPnl,
		MaxOpenInterest: binancePositionHistory.MaxOpenInterest,
		Status:          binancePositionHistory.Status,
		ClosedVolume:    binancePositionHistory.ClosedVolume,
		CreatedAt:       binancePositionHistory.CreatedAt,
		UpdatedAt:       binancePositionHistory.UpdatedAt,
	}, nil
}

// InsertBinancePositionHistory .
func (b *BinanceUserRepo) InsertBinancePositionHistory(ctx context.Context, binancePositionHistory *biz.BinancePositionHistory) (*biz.BinancePositionHistory, error) {
	insert := &BinancePositionHistory{
		Symbol:          binancePositionHistory.Symbol,
		Side:            binancePositionHistory.Side,
		Closed:          binancePositionHistory.Closed,
		Opened:          binancePositionHistory.Opened,
		AvgCost:         binancePositionHistory.AvgCost,
		AvgClosePrice:   binancePositionHistory.AvgClosePrice,
		ClosingPnl:      binancePositionHistory.ClosingPnl,
		MaxOpenInterest: binancePositionHistory.MaxOpenInterest,
		ClosedVolume:    binancePositionHistory.ClosedVolume,
		Status:          binancePositionHistory.Status,
	}

	res := b.data.DB(ctx).Table("new_binance_position_history").Create(&insert)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_BINANCE_POSITION_HISTORY_ERROR", "创建数据失败")
	}

	return &biz.BinancePositionHistory{
		ID:              insert.ID,
		Symbol:          insert.Symbol,
		Side:            insert.Side,
		Closed:          insert.Closed,
		Opened:          insert.Opened,
		AvgCost:         insert.AvgCost,
		AvgClosePrice:   insert.AvgClosePrice,
		ClosingPnl:      insert.ClosingPnl,
		MaxOpenInterest: insert.MaxOpenInterest,
		ClosedVolume:    insert.ClosedVolume,
		CreatedAt:       insert.CreatedAt,
		UpdatedAt:       insert.UpdatedAt,
	}, nil
}

// InsertUserOrderDo .
func (b *BinanceUserRepo) InsertUserOrderDo(ctx context.Context, userOrderDo *biz.UserOrderDo) error {
	insert := &UserOrderDo{
		ApiKey:       userOrderDo.ApiKey,
		ApiSecret:    userOrderDo.ApiSecret,
		ApiKeyTwo:    userOrderDo.ApiKeyTwo,
		ApiSecretTwo: userOrderDo.ApiSecretTwo,
		Symbol:       userOrderDo.Symbol,
		Status:       userOrderDo.Status,
		CloseRate:    userOrderDo.CloseRate,
		CloseBase:    userOrderDo.CloseBase,
		Qty:          userOrderDo.Qty,
		Price:        userOrderDo.Price,
		ClosePrice:   userOrderDo.ClosePrice,
		Amount:       userOrderDo.Amount,
		QtyTwo:       userOrderDo.QtyTwo,
		PriceTwo:     userOrderDo.PriceTwo,
		AmountTwo:    userOrderDo.AmountTwo,
	}

	res := b.data.DB(ctx).Table("new_user_order_do").Create(&insert)
	if res.Error != nil {
		return errors.New(500, "CREATE_USER_ORDER_DO_ERROR", "创建数据失败")
	}

	return nil
}

// InsertUserOrderDoNew .
func (b *BinanceUserRepo) InsertUserOrderDoNew(ctx context.Context, userOrderDo *biz.UserOrderDoNew) error {
	insert := &UserOrderDoNew{
		ApiKey:       userOrderDo.ApiKey,
		ApiSecret:    userOrderDo.ApiSecret,
		ApiKeyTwo:    userOrderDo.ApiKeyTwo,
		ApiSecretTwo: userOrderDo.ApiSecretTwo,
		Symbol:       userOrderDo.Symbol,
		SymbolTwo:    userOrderDo.SymbolTwo,
		Status:       userOrderDo.Status,
		Qty:          userOrderDo.Qty,
		QtyTwo:       userOrderDo.QtyTwo,
		Side:         userOrderDo.Side,
		SideTwo:      userOrderDo.SideTwo,
		OrderIdTwo:   userOrderDo.OrderIdTwo,
		OrderId:      userOrderDo.OrderId,
	}

	res := b.data.DB(ctx).Table("new_user_order_do_new").Create(&insert)
	if res.Error != nil {
		return errors.New(500, "CREATE_USER_ORDER_DO_NEW_ERROR", "创建数据失败")
	}

	return nil
}

// UpdateUserOrderDo .
func (b *BinanceUserRepo) UpdateUserOrderDo(ctx context.Context, id uint64, closeOrderId string) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_order_do_new").Where("id=?", id).
		Updates(map[string]interface{}{"status": 2, "close_order_id": closeOrderId, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_ORDER_DO_ERROR", "UPDATE_USER_ORDER_DO_ERROR")
	}

	return true, nil
}

// UpdateUserOrderDoThree .
func (b *BinanceUserRepo) UpdateUserOrderDoThree(ctx context.Context, id uint64, qty float64, qtyTwo float64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if 0 < qty {
		if err = b.data.DB(ctx).Table("new_user_order_do_new").Where("id=?", id).
			Updates(map[string]interface{}{"qty": qty, "updated_at": now}).Error; nil != err {
			return false, errors.NotFound("UPDATE_USER_ORDER_DO_ERROR", "UPDATE_USER_ORDER_DO_ERROR")
		}
	} else if 0 < qtyTwo {
		if err = b.data.DB(ctx).Table("new_user_order_do_new").Where("id=?", id).
			Updates(map[string]interface{}{"qty_two": qtyTwo, "updated_at": now}).Error; nil != err {
			return false, errors.NotFound("UPDATE_USER_ORDER_DO_ERROR", "UPDATE_USER_ORDER_DO_ERROR")
		}
	}

	return true, nil
}

// UpdateUserOrderDoTwo .
func (b *BinanceUserRepo) UpdateUserOrderDoTwo(ctx context.Context, id uint64, orderId string, orderIdTwo string) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if 0 < len(orderId) {
		if err = b.data.DB(ctx).Table("new_user_order_do_new").Where("id=?", id).
			Updates(map[string]interface{}{"order_id": orderId, "updated_at": now}).Error; nil != err {
			return false, errors.NotFound("UPDATE_USER_ORDER_DO_ERROR", "UPDATE_USER_ORDER_DO_ERROR")
		}
	} else if 0 < len(orderIdTwo) {
		if err = b.data.DB(ctx).Table("new_user_order_do_new").Where("id=?", id).
			Updates(map[string]interface{}{"order_id_two": orderIdTwo, "updated_at": now}).Error; nil != err {
			return false, errors.NotFound("UPDATE_USER_ORDER_DO_ERROR", "UPDATE_USER_ORDER_DO_ERROR")
		}
	}

	return true, nil
}

// GetUserOrderDoNew .
func (b *BinanceUserRepo) GetUserOrderDoNew() ([]*biz.UserOrderDoNew, error) {
	var userOrderDo []*UserOrderDoNew
	if err := b.data.db.Table("new_user_order_do_new").Where("status=?", 1).Find(&userOrderDo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_DO_New_ERROR", err.Error())
	}

	res := make([]*biz.UserOrderDoNew, 0)
	for _, v := range userOrderDo {
		res = append(res, &biz.UserOrderDoNew{
			ID:           v.ID,
			ApiKey:       v.ApiKey,
			ApiSecret:    v.ApiSecret,
			ApiKeyTwo:    v.ApiKeyTwo,
			ApiSecretTwo: v.ApiSecretTwo,
			Symbol:       v.Symbol,
			SymbolTwo:    v.SymbolTwo,
			Status:       v.Status,
			Qty:          v.Qty,
			QtyTwo:       v.QtyTwo,
			Side:         v.Side,
			SideTwo:      v.SideTwo,
			OrderId:      v.OrderId,
			OrderIdTwo:   v.OrderIdTwo,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserOrderDoLast .
func (b *BinanceUserRepo) GetUserOrderDoLast() ([]*biz.UserOrderDo, error) {
	var userOrderDo []*UserOrderDo
	if err := b.data.db.Table("new_user_order_do_new").Where("status=?", 1).Order("id desc").Find(&userOrderDo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_DO_ERROR", err.Error())
	}

	res := make([]*biz.UserOrderDo, 0)
	for _, v := range userOrderDo {
		res = append(res, &biz.UserOrderDo{
			ID:           v.ID,
			ApiKey:       v.ApiKey,
			ApiSecret:    v.ApiSecret,
			ApiKeyTwo:    v.ApiKeyTwo,
			ApiSecretTwo: v.ApiSecretTwo,
			Symbol:       v.Symbol,
			Status:       v.Status,
			CloseRate:    v.CloseRate,
			CloseBase:    v.CloseBase,
			Qty:          v.Qty,
			Price:        v.Price,
			ClosePrice:   v.ClosePrice,
			Amount:       v.Amount,
			QtyTwo:       v.QtyTwo,
			PriceTwo:     v.PriceTwo,
			AmountTwo:    v.AmountTwo,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserOrderDo .
func (b *BinanceUserRepo) GetUserOrderDo() ([]*biz.UserOrderDo, error) {
	var userOrderDo []*UserOrderDo
	if err := b.data.db.Table("new_user_order_do").Where("status=?", 1).Find(&userOrderDo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_DO_ERROR", err.Error())
	}

	res := make([]*biz.UserOrderDo, 0)
	for _, v := range userOrderDo {
		res = append(res, &biz.UserOrderDo{
			ID:           v.ID,
			ApiKey:       v.ApiKey,
			ApiSecret:    v.ApiSecret,
			ApiKeyTwo:    v.ApiKeyTwo,
			ApiSecretTwo: v.ApiSecretTwo,
			Symbol:       v.Symbol,
			Status:       v.Status,
			CloseRate:    v.CloseRate,
			CloseBase:    v.CloseBase,
			Qty:          v.Qty,
			Price:        v.Price,
			ClosePrice:   v.ClosePrice,
			Amount:       v.Amount,
			QtyTwo:       v.QtyTwo,
			PriceTwo:     v.PriceTwo,
			AmountTwo:    v.AmountTwo,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		})
	}

	return res, nil
}

// InsertFilData .
func (b *BinanceUserRepo) InsertFilData(ctx context.Context, filData *biz.FilData2) error {
	insert := &FilData{
		ID:       filData.ID,
		FromUser: filData.From,
		FromType: filData.FromType,
		To:       filData.To,
		Value:    filData.Value,
	}

	res := b.data.DB(ctx).Table("fil_data").Create(&insert)
	if res.Error != nil {
		return errors.New(500, "CREATE_FIL_DATA_ERROR", "创建数据失败")
	}

	return nil
}

// GetFilData .
func (b *BinanceUserRepo) GetFilData(address ...[]string) (map[string][]*biz.FilData2, error) {
	var filData []*FilData
	if err := b.data.db.Table("fil_data").Where("from_user in(?)", address).Find(&filData).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_FIL_DATA_ERROR", err.Error())
	}

	res := make(map[string][]*biz.FilData2, 0)
	for _, v := range filData {
		if _, ok := res[v.FromUser]; !ok {
			res[v.FromUser] = make([]*biz.FilData2, 0)
		}

		res[v.FromUser] = append(res[v.FromUser], &biz.FilData2{
			ID:       v.ID,
			From:     v.FromUser,
			FromType: v.FromType,
			To:       v.To,
			Value:    v.Value,
		})
	}

	return res, nil
}

// SAddOrderSetSellLongOrBuyShort 平仓订单信息放入缓存.
func (b *BinanceUserRepo) SAddOrderSetSellLongOrBuyShort(ctx context.Context, OrderId int64) error {
	var err error
	err = b.data.rdb.SAdd(ctx, "new_user_order_sell_long_buy_short", OrderId).Err()
	return err
}

// SMembersOrderSetSellLongOrBuyShort 平仓订单信息.
func (b *BinanceUserRepo) SMembersOrderSetSellLongOrBuyShort(ctx context.Context) ([]string, error) {
	return b.data.rdb.SMembers(ctx, "new_user_order_sell_long_buy_short").Result()
}

// SRemOrderSetSellLongOrBuyShort 平仓订单信息放入缓存.
func (b *BinanceUserRepo) SRemOrderSetSellLongOrBuyShort(ctx context.Context, OrderId int64) error {
	var err error
	err = b.data.rdb.SRem(ctx, "new_user_order_sell_long_buy_short", OrderId).Err()
	return err
}

// SAddOrderSetSellLongOrBuyShortTwo 平仓订单信息放入缓存.
func (b *BinanceUserRepo) SAddOrderSetSellLongOrBuyShortTwo(ctx context.Context, OrderId int64) error {
	var err error
	err = b.data.rdb.SAdd(ctx, "new_user_order_sell_long_buy_short_two", OrderId).Err()
	return err
}

// SMembersOrderSetSellLongOrBuyShortTwo 平仓订单信息.
func (b *BinanceUserRepo) SMembersOrderSetSellLongOrBuyShortTwo(ctx context.Context) ([]string, error) {
	return b.data.rdb.SMembers(ctx, "new_user_order_sell_long_buy_short_two").Result()
}

// SRemOrderSetSellLongOrBuyShortTwo 平仓订单信息放入缓存.
func (b *BinanceUserRepo) SRemOrderSetSellLongOrBuyShortTwo(ctx context.Context, OrderId int64) error {
	var err error
	err = b.data.rdb.SRem(ctx, "new_user_order_sell_long_buy_short_two", OrderId).Err()
	return err
}

// GetUsersNew .
func (b *BinanceUserRepo) GetUsersNew() ([]*biz.User, error) {
	var users []*User
	res := make([]*biz.User, 0)

	if err := b.data.db.Table("new_user").
		Where("use_new_system=? and is_dai=? and bind_trader_status_tfi=?", 2, 1, 1).
		Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	for _, v := range users {
		res = append(res, &biz.User{
			ID:               v.ID,
			Address:          v.Address,
			ApiKey:           v.ApiKey,
			ApiStatus:        v.ApiStatus,
			BindTraderStatus: v.BindTraderStatus,
			ApiSecret:        v.ApiSecret,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		})
	}

	return res, nil
}

// GetNewUserBindTraderTwoByUserIdMap .
func (b *BinanceUserRepo) GetNewUserBindTraderTwoByUserIdMap() (map[uint64][]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	res := make(map[uint64][]*biz.UserBindTrader, 0)

	if err := b.data.db.Table("new_user_bind_trader_two").Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_TWO_ERROR", err.Error())
	}

	for _, v := range userBindTrader {
		if _, ok := res[v.UserId]; !ok {
			res[v.UserId] = make([]*biz.UserBindTrader, 0)
		}

		res[v.UserId] = append(res[v.UserId], &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			Amount:    v.Amount,
			Status:    v.Status,
			InitOrder: v.InitOrder,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserOrderTwoByUserTraderIdNew .
func (b *BinanceUserRepo) GetUserOrderTwoByUserTraderIdNew(userId uint64, traderId uint64) ([]*biz.UserOrder, error) {
	var userOrder []*UserOrder

	res := make([]*biz.UserOrder, 0)
	if err := b.data.db.Table("new_user_order_two_"+strconv.FormatInt(int64(userId), 10)).
		Where("user_id=? and trader_id=?", userId, traderId).
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_TWO_ERROR", err.Error())
	}

	for _, v := range userOrder {
		res = append(res, &biz.UserOrder{
			ID:            v.ID,
			UserId:        v.UserId,
			TraderId:      v.TraderId,
			ClientOrderId: v.ClientOrderId,
			OrderId:       v.OrderId,
			Symbol:        v.Symbol,
			Side:          v.Side,
			PositionSide:  v.PositionSide,
			Quantity:      v.Quantity,
			Price:         v.Price,
			TraderQty:     v.TraderQty,
			OrderType:     v.OrderType,
			ClosePosition: v.ClosePosition,
			CumQuote:      v.CumQuote,
			ExecutedQty:   v.ExecutedQty,
			AvgPrice:      v.AvgPrice,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return res, nil
}

// UpdateTraderInfo .
func (b *BinanceUserRepo) UpdateTraderInfo(ctx context.Context, traderId uint64, baseMoney float64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_trader_info").Where("trader_id=?", traderId).
		Updates(map[string]interface{}{"base_money": baseMoney, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_TRADER_INFO_ERROR", "UPDATE_TRADER_INFO_ERROR")
	}

	return true, nil
}

// UpdateUserInfo .
func (b *BinanceUserRepo) UpdateUserInfo(ctx context.Context, userId uint64, baseMoney float64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("new_user_info").Where("user_id=?", userId).
		Updates(map[string]interface{}{"base_money": baseMoney, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_INFO_ERROR", "UPDATE_USER_INFO_ERROR")
	}

	return true, nil
}

// GetUserInfo .
func (b *BinanceUserRepo) GetUserInfo() (map[uint64]*biz.UserInfo, error) {
	var userOrder []*UserInfo

	res := make(map[uint64]*biz.UserInfo, 0)
	if err := b.data.db.Table("new_user_info").
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "FIND_USER_INFO_ERROR", err.Error())
	}

	for _, v := range userOrder {
		res[v.UserId] = &biz.UserInfo{
			ID:        v.ID,
			UserId:    v.UserId,
			BId:       v.BId,
			BaseMoney: v.BaseMoney,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return res, nil
}

// GetTraderInfo .
func (b *BinanceUserRepo) GetTraderInfo() (map[uint64]*biz.TraderInfo, error) {
	var userOrder []*TraderInfo

	res := make(map[uint64]*biz.TraderInfo, 0)
	if err := b.data.db.Table("new_trader_info").
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "FIND_TRADER_INFO_ERROR", err.Error())
	}

	for _, v := range userOrder {
		res[v.TraderId] = &biz.TraderInfo{
			ID:        v.ID,
			TraderId:  v.TraderId,
			BId:       v.BId,
			BaseMoney: v.BaseMoney,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return res, nil
}
