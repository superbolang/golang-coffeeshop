# golang-coffeeshoop

This is a project to practice goroutine in golang. Here is the scenario :   

A coffee shop operates by serving pre-order from customers, it has 5 baristas ready to take order, serving and finish. The program simulate 4 different scenarios for process :  
- Sequential process  
  Each barista handle only one order at a time, no other barista is allowed to work while there is coffee in the making  
- Concurrent process with goroutine  
  The orders are processed concurrent by available barista in each process (taking order, serving, finish). It makes possible for one barita to do multiple work, since every process is open to each barista  
- Goroutine with waitgroup  
  Waitgroup prevents goroutine from going rampage, it keeps the output to wait othres goroutine to finish their work and then release the output. While the output still random, it prevents possible loss from previous scenario  
- Goroutine with waitgroup and channel  
  Goroutine is a close operation, meaning it operates inside its own function. What if we want the result to be used for another process ? Here comes channel, a mechanism that allows goroutines (concurrently executing functions) to communicate and synchronize with each other  
  
## How to use  

In the main function there are 4 functions, comment or uncomment each to run. When ask for order input, try with small number and increase it to observe the result  

## Expected result  

🚗 orderWithouthGoroutine()  
- You'll see some sequential output, each bartender working on one order from taking order to complete  

🚕 orderWithGoroutine()  
- You'll see some random bartender working on random process (take order, serving, finish) 
- To see the result you must input order at least 1.000
- Here comes the problem, some times you'll experience main function finishes earlier than your schemed goroutine and that's normal. It this case there will be unfinished goroutine  
- Otherwise you'll experience your goroutine finishes eralier than main function, in this case goroutine will be executed again  
- For your reference, main function is also a goroutine by default 😅  

🚌 orderWait()  
- You'll see all goroutine finish their work, but still in random order  

🚙 orderWithChannel()  
- You'll see a much nicer look, each order is done by the same goroutine