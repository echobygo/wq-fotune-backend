package service

import (
	"context"
	"encoding/json"
	"github.com/shopspring/decimal"
	"strings"
	"time"
	"wq-fotune-backend/libs/logger"
	"wq-fotune-backend/pkg/response"
	"wq-fotune-backend/service/exchange-srv/client"
	"wq-fotune-backend/service/exchange-srv/model"
	pb "wq-fotune-backend/service/exchange-srv/proto"
	globalF "wq-fotune-backend/service/forward-offer-srv/global"
	quoteCron "wq-fotune-backend/service/quote-srv/cron"
	pbQuote "wq-fotune-backend/service/quote-srv/proto"
	fotune_srv_user "wq-fotune-backend/service/user-srv/proto"
)

func (e *ExOrderService) GetTradeSymbols(exchange, symbol string) ([]*pb.Symbol, error) {
	symbolList := make([]*pb.Symbol, 0)
	//获取可交易的品种币对 顺便获取行情
	symbols := e.dao.GetAllSymbolWithState(1, exchange, symbol) //state 1open  0 close
	if len(symbols) == 0 {
		return nil, response.NewDataNotFound(ErrID, "没有品种数据")
	}
	ticksResp, errTick := client.QuoteService.GetTicksWithExchangeSymbol(context.Background(), &pbQuote.GetTicksSymbolReq{
		Symbol:   symbol,
		Exchange: exchange,
	})
	var ticks []quoteCron.Ticker
	if errTick == nil {
		if err := json.Unmarshal(ticksResp.Ticks, &ticks); err != nil {
			logger.Warnf("解析tick数据失败 GetTradeSymbols %v", err)
		}
	}
	for _, s := range symbols {
		newSymbol := strings.ToUpper(strings.ReplaceAll(s.Symbol, "/", "-"))
		price := "0"
		change := "+0.00%"
		data := &pb.Symbol{
			Symbol:   s.Symbol,
			Exchange: s.Exchange,
			State:    s.State,
			Unit:     s.Unit,
			Price:    "0",
			Change:   "",
		}
		if errTick != nil || len(ticks) == 0 {
			logger.Warnf("GetTradeSymbols func GetOkexTicks出错了 %v", errTick)
			data.Price = price
			data.Change = change
			symbolList = append(symbolList, data)
			continue
		}
		for _, tick := range ticks {
			if tick.Symbol == newSymbol {
				data.Price = globalF.Float64ToString(tick.Last)
				data.Change = tick.Change
				symbolList = append(symbolList, data)
			}
		}

	}
	return symbolList, nil
}

func (e *ExOrderService) GetTradeCount(userId, strategyId string) (int32, error) {
	return e.dao.TradeCount(userId, strategyId)
}

func (e *ExOrderService) GetTradeList(userId, strategyId string, pageNum, pageSize int32) []*model.WqTrade {
	return e.dao.GetTradeList(userId, strategyId, pageNum, pageSize)
}

func (e *ExOrderService) GetProfitRealTime(userId, strategyId string) []*model.WqProfit {
	return e.dao.GetProfitRealTime(userId, strategyId)
}

func (e *ExOrderService) GetProfitByStrID(userId, strategyId string) (*model.WqProfit, error) {
	return e.dao.GetProfitByID(userId, strategyId)
}

func (e *ExOrderService) GetProfitDailyByStrID(userId, strategyId string, limit int) []*model.WqProfitDaily {
	return e.dao.GetProfitDailyList(userId, strategyId, limit)
}

func (e *ExOrderService) GetProfitSortByRate() []*model.WqProfit {
	sql := ""
	orderBy := "rate_return desc"
	return e.dao.GetProfitListBySql(sql, orderBy)
}

func (e *ExOrderService) GetProfitSortByRateYear() []*model.WqProfit {
	sql := ""
	orderBy := "rate_return_year desc"
	return e.dao.GetProfitListBySql(sql, orderBy)
}

func (e *ExOrderService) SaveRateRankToRedis(key string, data interface{}) {
	if err := e.cacheService.CacheData(key, data, time.Minute*5); err != nil {
		logger.Warnf("保存排名数据失败 %v", err)
	}
}

func (e *ExOrderService) CacheRateReturn() {
	logger.Infof("开始更新收益排名")
	var saveData []*model.RateRank
	profitList := e.GetProfitSortByRate()
	index := 0
	userIdList := make(map[string]bool, 0)
	for _, profit := range profitList {
		user, err := client.UserService.GetUserInfo(context.Background(), &fotune_srv_user.UserInfoReq{UserID: profit.UserID})
		if err != nil {
			logger.Warnf("CacheRateReturn 查找用户信息失败 %v 用户id %s", err, profit.UserID)
			continue
		}
		if _, ok := userIdList[user.UserID]; ok {
			continue
		}
		userIdList[user.UserID] = true
		index += 1
		if index >= 21 {
			continue
		}
		saveData = append(saveData, &model.RateRank{
			ID:             index,
			UserId:         user.UserID,
			Avatar:         user.Avatar,
			Name:           user.Name,
			RateReturn:     decimal.NewFromFloat(profit.RateReturn).String() + "%",
			RateReturnYear: decimal.NewFromFloat(profit.RateReturnYear).String() + "%",
		})
	}
	marshal, err := json.Marshal(saveData)
	if err != nil {
		logger.Warnf("json Marshal 排名数据失败 %v", err)
		return
	}
	e.SaveRateRankToRedis("rateReturnSort", marshal)
	logger.Infof("结束更新收益排名")
}

func (e *ExOrderService) CacheRateReturnYear() {
	logger.Infof("开始更新年化收益排名")
	var saveData []*model.RateRank
	profitList := e.GetProfitSortByRateYear()
	index := 0
	userIdList := make(map[string]bool, 0)
	for _, profit := range profitList {
		user, err := client.UserService.GetUserInfo(context.Background(), &fotune_srv_user.UserInfoReq{UserID: profit.UserID})
		if err != nil {
			logger.Warnf("CacheRateReturnYear 查找用户信息失败 %v 用户id %s", err, profit.UserID)
			continue
		}
		if _, ok := userIdList[user.UserID]; ok {
			continue
		}
		userIdList[user.UserID] = true
		index += 1
		if index >= 21 {
			continue
		}
		saveData = append(saveData, &model.RateRank{
			ID:             index,
			UserId:         user.UserID,
			Avatar:         user.Avatar,
			Name:           user.Name,
			RateReturn:     decimal.NewFromFloat(profit.RateReturn).String() + "%",
			RateReturnYear: decimal.NewFromFloat(profit.RateReturnYear).String() + "%",
		})
	}
	marshal, err := json.Marshal(saveData)
	if err != nil {
		logger.Warnf("json Marshal 排名数据失败 %v", err)
		return
	}
	e.SaveRateRankToRedis("rateReturnYearSort", marshal)
	logger.Infof("结束更新收益排名")
}

func (e *ExOrderService) GetSymbolRankWithRateYear() []*model.WqSymbolRecommend {
	return e.dao.GetSymbolRecommend(1)
}

func (e *ExOrderService) GetBtcTickPrice() float64 {
	ticksResp, err := client.QuoteService.GetTicksWithExchangeSymbol(context.Background(), &pbQuote.GetTicksSymbolReq{
		Symbol:   "usdt",
		Exchange: "binance",
	})
	price := 0.0

	if err != nil {
		logger.Errorf("获取行情失败 %v", err)
		return price
	}
	var tickList []quoteCron.Ticker
	if err := json.Unmarshal(ticksResp.Ticks, &tickList); err != nil {
		logger.Warnf("解析tick数据失败 GetTradeSymbols %v", err)
		return price
	}
	for _, ticker := range tickList {
		if ticker.Symbol == "BTC-USDT" {
			price = ticker.Last
			return price
		}
	}
	return price
}
