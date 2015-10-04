package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
	"strings"
)

type YahooFinanceStruct struct {
	Query struct {
		Count       int    `json:"count"`
		Created     string `json:"created"`
		Diagnostics struct {
			Build_version string `json:"build-version"`
			Cache         struct {
				Content              string `json:"content"`
				Execution_start_time string `json:"execution-start-time"`
				Execution_stop_time  string `json:"execution-stop-time"`
				Execution_time       string `json:"execution-time"`
				Method               string `json:"method"`
				Type                 string `json:"type"`
			} `json:"cache"`
			Javascript struct {
				Execution_start_time string `json:"execution-start-time"`
				Execution_stop_time  string `json:"execution-stop-time"`
				Execution_time       string `json:"execution-time"`
				Instructions_used    string `json:"instructions-used"`
				Table_name           string `json:"table-name"`
			} `json:"javascript"`
			PubliclyCallable string `json:"publiclyCallable"`
			Query            struct {
				Content              string `json:"content"`
				Execution_start_time string `json:"execution-start-time"`
				Execution_stop_time  string `json:"execution-stop-time"`
				Execution_time       string `json:"execution-time"`
				Params               string `json:"params"`
			} `json:"query"`
			Service_time string `json:"service-time"`
			URL          []struct {
				Content              string `json:"content"`
				Execution_start_time string `json:"execution-start-time"`
				Execution_stop_time  string `json:"execution-stop-time"`
				Execution_time       string `json:"execution-time"`
			} `json:"url"`
			User_time string `json:"user-time"`
		} `json:"diagnostics"`
		Lang    string `json:"lang"`
		Results struct {
			Quote []struct {
				AfterHoursChangeRealtime                       interface{} `json:"AfterHoursChangeRealtime"`
				AnnualizedGain                                 interface{} `json:"AnnualizedGain"`
				Ask                                            string      `json:"Ask"`
				AskRealtime                                    interface{} `json:"AskRealtime"`
				AverageDailyVolume                             string      `json:"AverageDailyVolume"`
				Bid                                            string      `json:"Bid"`
				BidRealtime                                    interface{} `json:"BidRealtime"`
				BookValue                                      string      `json:"BookValue"`
				Change                                         string      `json:"Change"`
				ChangeFromFiftydayMovingAverage                string      `json:"ChangeFromFiftydayMovingAverage"`
				ChangeFromTwoHundreddayMovingAverage           string      `json:"ChangeFromTwoHundreddayMovingAverage"`
				ChangeFromYearHigh                             string      `json:"ChangeFromYearHigh"`
				ChangeFromYearLow                              string      `json:"ChangeFromYearLow"`
				ChangePercentRealtime                          interface{} `json:"ChangePercentRealtime"`
				ChangeRealtime                                 interface{} `json:"ChangeRealtime"`
				ChangePercentChange                            string      `json:"Change_PercentChange"`
				ChangeinPercent                                string      `json:"ChangeinPercent"`
				Commission                                     interface{} `json:"Commission"`
				Currency                                       string      `json:"Currency"`
				DaysHigh                                       string      `json:"DaysHigh"`
				DaysLow                                        string      `json:"DaysLow"`
				DaysRange                                      string      `json:"DaysRange"`
				DaysRangeRealtime                              interface{} `json:"DaysRangeRealtime"`
				DaysValueChange                                interface{} `json:"DaysValueChange"`
				DaysValueChangeRealtime                        interface{} `json:"DaysValueChangeRealtime"`
				DividendPayDate                                string      `json:"DividendPayDate"`
				DividendShare                                  string      `json:"DividendShare"`
				DividendYield                                  string      `json:"DividendYield"`
				EBITDA                                         string      `json:"EBITDA"`
				EPSEstimateCurrentYear                         string      `json:"EPSEstimateCurrentYear"`
				EPSEstimateNextQuarter                         string      `json:"EPSEstimateNextQuarter"`
				EPSEstimateNextYear                            string      `json:"EPSEstimateNextYear"`
				EarningsShare                                  string      `json:"EarningsShare"`
				ErrorIndicationreturnedforsymbolchangedinvalid interface{} `json:"ErrorIndicationreturnedforsymbolchangedinvalid"`
				ExDividendDate                                 string      `json:"ExDividendDate"`
				FiftydayMovingAverage                          string      `json:"FiftydayMovingAverage"`
				HighLimit                                      interface{} `json:"HighLimit"`
				HoldingsGain                                   interface{} `json:"HoldingsGain"`
				HoldingsGainPercent                            interface{} `json:"HoldingsGainPercent"`
				HoldingsGainPercentRealtime                    interface{} `json:"HoldingsGainPercentRealtime"`
				HoldingsGainRealtime                           interface{} `json:"HoldingsGainRealtime"`
				HoldingsValue                                  interface{} `json:"HoldingsValue"`
				HoldingsValueRealtime                          interface{} `json:"HoldingsValueRealtime"`
				LastTradeDate                                  string      `json:"LastTradeDate"`
				LastTradePriceOnly                             string      `json:"LastTradePriceOnly"`
				LastTradeRealtimeWithTime                      interface{} `json:"LastTradeRealtimeWithTime"`
				LastTradeTime                                  string      `json:"LastTradeTime"`
				LastTradeWithTime                              string      `json:"LastTradeWithTime"`
				LowLimit                                       interface{} `json:"LowLimit"`
				MarketCapRealtime                              interface{} `json:"MarketCapRealtime"`
				MarketCapitalization                           string      `json:"MarketCapitalization"`
				MoreInfo                                       interface{} `json:"MoreInfo"`
				Name                                           string      `json:"Name"`
				Notes                                          interface{} `json:"Notes"`
				OneyrTargetPrice                               string      `json:"OneyrTargetPrice"`
				Open                                           string      `json:"Open"`
				OrderBookRealtime                              interface{} `json:"OrderBookRealtime"`
				PEGRatio                                       string      `json:"PEGRatio"`
				PERatio                                        string      `json:"PERatio"`
				PERatioRealtime                                interface{} `json:"PERatioRealtime"`
				PercebtChangeFromYearHigh                      string      `json:"PercebtChangeFromYearHigh"`
				PercentChange                                  string      `json:"PercentChange"`
				PercentChangeFromFiftydayMovingAverage         string      `json:"PercentChangeFromFiftydayMovingAverage"`
				PercentChangeFromTwoHundreddayMovingAverage    string      `json:"PercentChangeFromTwoHundreddayMovingAverage"`
				PercentChangeFromYearLow                       string      `json:"PercentChangeFromYearLow"`
				PreviousClose                                  string      `json:"PreviousClose"`
				PriceBook                                      string      `json:"PriceBook"`
				PriceEPSEstimateCurrentYear                    string      `json:"PriceEPSEstimateCurrentYear"`
				PriceEPSEstimateNextYear                       string      `json:"PriceEPSEstimateNextYear"`
				PricePaid                                      interface{} `json:"PricePaid"`
				PriceSales                                     string      `json:"PriceSales"`
				SharesOwned                                    interface{} `json:"SharesOwned"`
				ShortRatio                                     string      `json:"ShortRatio"`
				StockExchange                                  string      `json:"StockExchange"`
				Symbol                                         string      `json:"Symbol"`
				TickerTrend                                    interface{} `json:"TickerTrend"`
				TradeDate                                      interface{} `json:"TradeDate"`
				TwoHundreddayMovingAverage                     string      `json:"TwoHundreddayMovingAverage"`
				Volume                                         string      `json:"Volume"`
				YearHigh                                       string      `json:"YearHigh"`
				YearLow                                        string      `json:"YearLow"`
				YearRange                                      string      `json:"YearRange"`
				symbol                                         string      `json:"symbol"`
			} `json:"quote"`
		} `json:"results"`
	} `json:"query"`
}

type ClientBuyRequest struct {
	Budget        float32       //allocated in client
	Stocks        []ClientStock //allocated in client
	UninvestedBal float32       //allocated in client
}

type ClientStock struct {
	Symbol             string  //allocated in client
	StockName          string  //allocated in server
	PercentAllocated   int     //allocated in client
	BuyPrice           float32 //allocated in server
	NumberOfStocks     int     //allocated in server
	CurrentMarketPrice float32 //will be allocated in server during view
	LeftoverAmount     float32 //allocated in server
}

type ServerBuyResponse struct {
	TradeID int
	Entry   TransactionMapEntry
}

type ClientViewRequest struct {
}

type ServerViewResponse struct {
}

type TransactionMapEntry struct {
	StockList        []ClientStock
	UninvestedAmount float32
	InitialBudget    float32
}

var tradeIDCounter int
var transactionMap = make(map[int]TransactionMapEntry)

const yahooPrefixURL = "https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20yahoo.finance.quotes%20where%20symbol%20in%20(%22"
const yahooPostfixURL = "%22)&format=json&diagnostics=true&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys&callback="

type YahooFinanceServer struct{}

func (this *YahooFinanceServer) View(tradeID int, reply *ServerBuyResponse) error {
	isSingleStock := false

	fmt.Println(" Trade ID entered is : ", tradeID)
	if tradeID > len(transactionMap) {
		error_msg := "Invalid Trade ID. Maximum permitted value of trade ID is : -------" + strconv.Itoa(len(transactionMap))
		return fmt.Errorf(error_msg, tradeID)
	}

	fmt.Println("Received a request to View stocks, Generating URL for invocation")
	mapEntry := transactionMap[tradeID]
	if len(mapEntry.StockList) == 1 {
		isSingleStock = true
		var tempStock ClientStock
		tempStock.Symbol = "XYZ"
		mapEntry.StockList = append(mapEntry.StockList, tempStock)
	}

	url := generateURL(mapEntry.StockList)

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var yahooFinanceData YahooFinanceStruct // TopTracks
	err = json.Unmarshal(body, &yahooFinanceData)
	if err != nil {
		panic(err.Error())
	}
	if isSingleStock {
		mapEntry.StockList = mapEntry.StockList[:1]
		yahooFinanceData.Query.Results.Quote = yahooFinanceData.Query.Results.Quote[:1]
	}

	fmt.Println("----------------------------------------------------\nData is generated from Yahoo! Finance")

	outputMapEntry := populateCurrentMarketValues(yahooFinanceData, mapEntry)
	reply.TradeID = tradeID
	reply.Entry = outputMapEntry
	fmt.Println("----------------------------------------------------\nCompleted stock evaluation for trade ID", tradeIDCounter)

	return nil
}

func (this *YahooFinanceServer) Buy(request ClientBuyRequest, reply *ServerBuyResponse) error {
	var yahooFinanceData YahooFinanceStruct
	isSingleStock := false

	fmt.Println("Received a request to buy stocks, Generating URL for invocation")

	if len(request.Stocks) == 1 {
		isSingleStock = true
		var tempStock ClientStock
		tempStock.Symbol = "XYZ"
		request.Stocks = append(request.Stocks, tempStock)
	}

	url := generateURL(request.Stocks)

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &yahooFinanceData)
	if err != nil {
		panic(err.Error())
	}

	if isSingleStock {
		request.Stocks = request.Stocks[:1]
		yahooFinanceData.Query.Results.Quote = yahooFinanceData.Query.Results.Quote[:1]
	}

	fmt.Println("----------------------------------------------------\nData is generated from Yahoo! Finance")
	var currentMapEntry TransactionMapEntry
	evaluateStockData(&request, &yahooFinanceData, &currentMapEntry)
	tradeIDCounter++
	reply.TradeID = tradeIDCounter
	reply.Entry = currentMapEntry
	fmt.Println("----------------------------------------------------\nCompleted stock evaluation for trade ID", tradeIDCounter)
	transactionMap[tradeIDCounter] = currentMapEntry
	return nil
}

func populateCurrentMarketValues(financeData YahooFinanceStruct, mapEntry TransactionMapEntry) TransactionMapEntry {

	outputMapEntry := new(TransactionMapEntry)
	outputMapEntry.InitialBudget = mapEntry.InitialBudget
	outputMapEntry.UninvestedAmount = mapEntry.UninvestedAmount
	outputMapEntry.StockList = make([]ClientStock, len(mapEntry.StockList))
	copy(outputMapEntry.StockList, mapEntry.StockList)
	for i := 0; i < len(outputMapEntry.StockList); i++ {
		tempFloatVal, err := strconv.ParseFloat(financeData.Query.Results.Quote[i].Ask, 64)
		if err != nil {
			fmt.Println(err)
			//panic(err.Error())
			tempFloatVal = 0
		}
		outputMapEntry.StockList[i].CurrentMarketPrice = float32(tempFloatVal)
	}
	return *outputMapEntry
}

func evaluateStockData(request *ClientBuyRequest, financeData *YahooFinanceStruct, mapEntry *TransactionMapEntry) {

	fmt.Println("----------------------------------------------------\nStarting evaluation of stock data ")
	var totalUninvestedBal float32
	newStockList := make([]ClientStock, len(request.Stocks))

	copy(newStockList, request.Stocks)

	for i := 0; i < len(newStockList); i++ {

		for x := 0; x < len(financeData.Query.Results.Quote); x++ {

			if strings.Compare(newStockList[i].Symbol, financeData.Query.Results.Quote[x].Symbol) == 0 {
				//Found the matching stock
				fmt.Println("----------------------------------------------------\nCurrent stock being evaluated is :", financeData.Query.Results.Quote[x].Symbol)
				fmt.Println("Market price is : ", financeData.Query.Results.Quote[x].Ask)
				tempFloatVal, err := strconv.ParseFloat(financeData.Query.Results.Quote[x].Ask, 64)
				if err != nil {
					//panic(err.Error())
					tempFloatVal = 0
				}
				newStockList[i].BuyPrice = float32(tempFloatVal)
				fmt.Println("Float value is ", newStockList[i].BuyPrice)

				newStockList[i].StockName = financeData.Query.Results.Quote[x].Name
				budgetForStock := request.Budget * float32(newStockList[i].PercentAllocated) / float32(100.0)
				fmt.Println("Budget calculated is : ", budgetForStock)
				if newStockList[i].BuyPrice == 0 {
					//Stock data not found
					newStockList[i].NumberOfStocks = 0
				} else {
					newStockList[i].NumberOfStocks = int(budgetForStock / newStockList[i].BuyPrice)
				}

				fmt.Println("No. of stocks calculated is : ", newStockList[i].NumberOfStocks)
				newStockList[i].LeftoverAmount = budgetForStock - (float32(newStockList[i].NumberOfStocks) * newStockList[i].BuyPrice)
				fmt.Println("Leftover balance calculated is : ", newStockList[i].LeftoverAmount)
				totalUninvestedBal += newStockList[i].LeftoverAmount
				fmt.Println("Total uninvested balance is : ", totalUninvestedBal)
				break
			}
		}
		if len(newStockList[i].StockName) == 0 {
			//Data was not found for this stock
			fmt.Println("----------------------------------------------------\nStock data was not found for this stock code : ", newStockList[i].Symbol)
			newStockList[i].StockName = "NOT FOUND"
			newStockList[i].LeftoverAmount = request.Budget * float32(newStockList[i].PercentAllocated) / float32(100.0)
			newStockList[i].NumberOfStocks = 0
		}
	}
	mapEntry.StockList = newStockList
	mapEntry.UninvestedAmount = totalUninvestedBal
	mapEntry.InitialBudget = request.Budget
}

func generateURL(clientStocks []ClientStock) string {

	var buffer bytes.Buffer
	buffer.WriteString(yahooPrefixURL)
	for i := 0; i < len(clientStocks); i++ {
		buffer.WriteString(clientStocks[i].Symbol)
		buffer.WriteString("%22%2C%22")
	}
	tempString := strings.TrimSuffix(buffer.String(), "%22%2C%22")
	buffer.Reset()
	buffer.WriteString(tempString)
	buffer.WriteString(yahooPostfixURL)
	url := buffer.String()
	fmt.Println("----------------------------------------------------\nURL generated is :\n", url)
	return url
}

func server() {
	rpc.Register(new(YahooFinanceServer))
	ln, err := net.Listen("tcp", ":9999")
	fmt.Println("Starting up server")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server started")
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}
func main() {
	go server()
	var input string
	for 1 == 1 {
		fmt.Scanln(&input)
		if strings.EqualFold("exit", input) {
			break
		}
	}
}
