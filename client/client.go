package main

import (
	"fmt"
	"net/rpc"
	"strings"
)

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

func printRequestDetails(currentRequest ClientBuyRequest, totalBalance int) {

	fmt.Println("----------------------------------------------------\nSummary of request is as follows :")
	fmt.Println("Total Budget : ", currentRequest.Budget)
	fmt.Println("----------------------------------------------------")
	for i := 0; i < len(currentRequest.Stocks); i++ {
		fmt.Println("Stock symbol : ", currentRequest.Stocks[i].Symbol)
		fmt.Println("Pecentage allocated to stock : ", currentRequest.Stocks[i].PercentAllocated)
	}

	if totalBalance > 0 {
		fmt.Println("Remaining balance percentage is : ", totalBalance)
		fmt.Println("Uninvested balance amount is : ", currentRequest.UninvestedBal)

	}
}

func printResponseDetails(currentResponse ServerBuyResponse) {

	viewResponse := false
	var totalMarketBalance float32
	fmt.Println("----------------------------------------------------\nSummary of response is as follows :")
	fmt.Println("Trade ID : ", currentResponse.TradeID, "\n----------------------------------------------------")
	fmt.Println("Initial Budget : ", currentResponse.Entry.InitialBudget, "\n----------------------------------------------------")
	for i := 0; i < len(currentResponse.Entry.StockList); i++ {
		fmt.Println("Stock symbol : ", currentResponse.Entry.StockList[i].Symbol)
		fmt.Println("Stock Name : ", currentResponse.Entry.StockList[i].StockName)
		fmt.Println("Pecentage allocated to stock : ", currentResponse.Entry.StockList[i].PercentAllocated, "%")
		fmt.Println("No. of stocks bought : ", currentResponse.Entry.StockList[i].NumberOfStocks)
		fmt.Println("Buy Price : ", currentResponse.Entry.StockList[i].BuyPrice)
		if currentResponse.Entry.StockList[i].CurrentMarketPrice != 0 {
			fmt.Println("----------------------------------------------------\nCurrent Market Price : ", currentResponse.Entry.StockList[i].CurrentMarketPrice)
			viewResponse = true
			fmt.Println("----------------------------------------------------")
			if currentResponse.Entry.StockList[i].CurrentMarketPrice >= currentResponse.Entry.StockList[i].BuyPrice {
				fmt.Println("Total Profit on this stock : +", (currentResponse.Entry.StockList[i].CurrentMarketPrice - currentResponse.Entry.StockList[i].BuyPrice))
			} else {
				fmt.Println("Total Loss on this stock : ", (currentResponse.Entry.StockList[i].CurrentMarketPrice - currentResponse.Entry.StockList[i].BuyPrice))
			}
			totalMarketBalance += (currentResponse.Entry.StockList[i].CurrentMarketPrice * float32(currentResponse.Entry.StockList[i].NumberOfStocks))
		}
		fmt.Println("----------------------------------------------------")
	}
	fmt.Println("Uninvested Amount : ", currentResponse.Entry.UninvestedAmount)
	if viewResponse {
		fmt.Println("----------------------------------------------------\n Total current market worth of stock : ", (totalMarketBalance + currentResponse.Entry.UninvestedAmount), "\n-------------------------------------------")

	}
}

func clientBuy(request ClientBuyRequest) ServerBuyResponse {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Initiating call to server : ")

	var result ServerBuyResponse
	err = c.Call("YahooFinanceServer.Buy", request,
		&result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Function invocation successful!")
		return result
	}
	fmt.Println("Returning nil from client function")
	return result
}

func clientView(tradeID int) (ServerBuyResponse, error) {
	fmt.Println("Trade ID entered in clientView is : ", tradeID)
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("----------------------------------------------------\nInitiating call to server")
	var result ServerBuyResponse
	err = c.Call("YahooFinanceServer.View", tradeID,
		&result)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	fmt.Println("----------------------------------------------------\nFunction invocation successful!")
	return result, nil

}

func mergeStocks(stockArray []ClientStock, currentStock ClientStock) []ClientStock {
	var i int
	var stockFound bool
	outputArray := stockArray[:]

	for i = 0; i < len(stockArray); i++ {
		if strings.Compare(currentStock.Symbol, stockArray[i].Symbol) == 0 {
			stockFound = true
			fmt.Println("You have already entered this stock option earlier. Adding your allocated percentage to the percentage allocated previously.")
			stockArray[i].PercentAllocated += currentStock.PercentAllocated
		}
	}
	if !stockFound {
		outputArray = append(outputArray, currentStock)
	}
	return outputArray
}

func main() {
	var userChoice int
	var userBudget float32
	var userStockCode string
	var userYesNoChoice string
	var percentageAllocated int
	var transactionCommitted bool
	var totalBalance int
	var currentRequest ClientBuyRequest
	userContinue := true
	stockSymbols := make(map[int]string)
	stockSymbols[1] = "AAPL"
	stockSymbols[2] = "MSFT"
	stockSymbols[3] = "GOOGL"
	stockSymbols[4] = "YHOO"
	stockSymbols[5] = "CSCO"
	stockSymbols[6] = "ADBE"
	stockSymbols[7] = "LNKD"
	stockSymbols[8] = "VMW"
	stockSymbols[9] = "FB"

	for userContinue {
		fmt.Println("----------------------------------------------------\nChoose from the below 2 options\n1. Buy some stocks.\n2. View a stock portfolio.")
		fmt.Scanf("%d", &userChoice)
		switch userChoice {
		case 1: //User chose to buy some stocks

			transactionCommitted = false
			totalBalance = 100

			fmt.Println("----------------------------------------------------\nEnter your desired budget(in USD $) : ")
			fmt.Scanf("%f", &userBudget)
			currentRequest.Budget = userBudget
			currentRequest.Stocks = make([]ClientStock, 0)

			for !transactionCommitted {

				currentStock := new(ClientStock)

				fmt.Println("----------------------------------------------------\nYou can choose from the below stocks or enter your own stock code by pressing 0 :",
					"\n1.Apple Inc. (AAPL)",
					"\n2.Microsoft Corporation (MSFT)",
					"\n3.Google Inc. (GOOGL)",
					"\n4.Yahoo! Inc. (YHOO)",
					"\n5.Cisco Systems, Inc. (CSCO)",
					"\n6.Adobe Systems Incorporated (ADBE)",
					"\n7.LinkedIn Corporation (LNKD)",
					"\n8.VMware, Inc. (VMW)",
					"\n9.Facebook, Inc. (FB)",
					"\n0.Enter your own stock code")

				fmt.Scanf("%d", &userChoice)
				if userChoice < 0 && userChoice > 9 {
					//Invalid user choice
					panic("----------------------------------------------------\nYou have entered an invalid choice")
				} else if userChoice == 0 {
					//User wants to enter his own stock code
					fmt.Println("----------------------------------------------------\nEnter your custom stock code : ")
					fmt.Scanln(&userStockCode)
					currentStock.Symbol = userStockCode
				} else {
					//User has selected a stock code from the List
					currentStock.Symbol = stockSymbols[userChoice]
				}

				fmt.Println("----------------------------------------------------\nEnter percentage allocated to this stock (1-", totalBalance, ")")
				fmt.Scanf("%d", &percentageAllocated)

				if percentageAllocated < 0 || percentageAllocated > totalBalance {
					panic("----------------------------------------------------\nYou have entered an invalid percentage : ")
				}

				currentStock.PercentAllocated = percentageAllocated
				//Add code here to check for repeated stocks. If found, add their percentages
				//currentRequest.Stocks = append(currentRequest.Stocks, *currentStock)
				currentRequest.Stocks = mergeStocks(currentRequest.Stocks, *currentStock)
				totalBalance -= percentageAllocated

				if totalBalance == 0 {
					fmt.Println("----------------------------------------------------\nYou do not have any more money to purchase more stocks ")
					transactionCommitted = true
				} else {
					fmt.Println("----------------------------------------------------\n You have allocated only ", (100 - totalBalance), "% of your budget. Do you want to add any more stocks?(Y/N)")
					fmt.Scanln(&userYesNoChoice)
					if strings.Compare("N", strings.ToUpper(userYesNoChoice)) == 0 {
						fmt.Println("----------------------------------------------------\nCommitting the transaction and sending request to server")
						transactionCommitted = true
					}
				}
			}
			currentRequest.UninvestedBal = float32(totalBalance) * userBudget / float32(100.0)

			//Transaction is committed. Summary of request is displayed
			printRequestDetails(currentRequest, totalBalance)

			//Request is sent to server
			currentResponse := clientBuy(currentRequest)

			//Print the response
			printResponseDetails(currentResponse)

			fmt.Println("----------------------------------------------------\nDo you want to go again ?(Y/N)")
			fmt.Scanln(&userYesNoChoice)
			if strings.Compare("N", strings.ToUpper(userYesNoChoice)) == 0 {
				userContinue = false
			}

			break
		case 2: //User wants to view a portfolio
			fmt.Println("----------------------------------------------------\nEnter the Trade ID for which you want to view the portfolio : ")
			fmt.Scanf("%d", &userChoice)

			response, err := clientView(userChoice)
			if err == nil {
				printResponseDetails(response)

			} else {
				fmt.Println(err)
			}

			fmt.Println("----------------------------------------------------\nDo you want to go again ?(Y/N)")
			fmt.Scanln(&userYesNoChoice)
			if strings.Compare("N", strings.ToUpper(userYesNoChoice)) == 0 {
				userContinue = false
			}

			break
		default:
			fmt.Println("----------------------------------------------------\nInvalid choice")
		}
	}
	fmt.Println("------------------------THANK YOU!----------------------------")
}
