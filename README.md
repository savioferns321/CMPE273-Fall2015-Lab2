# CMPE273-Fall2015-Lab2
------------------------------------
Lab 2 : Yahoo Stock simulator
------------------------------------

Run Notes : 
1) Run the server.go to start the server.
2) Run the client.go to start the client.

Client
--------
Choose whether to buy stocks or view portfolio.

1. Buy stocks : 

a. Enter the budget for the transaction.
b. Choose from a stock code or enter your own stock code.
c. Enter the percentage allocated to the stock.
d. Repeat steps b to c until the percentage remaining is 0 or you do not want to enter any more stocks.

   Response from server is displayed.


2. View Portfolio : 

a. Enter the trade ID to view a portfolio.

If the trade ID is present on the server side, a response detailing the current price of the stock and
other total profit/loss about the transaction is displayed.
Otherwise, an error is displayed to the client.



