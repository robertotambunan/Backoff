# Jaen is a Backoff

### How to Install
go get -v github.com/robertotambunan/Jaen

### Usage
Jaen is a backoff. Backoff will provide you time.Duration to repeat your action. When we have case like connecting to database, or getting data from database, we will get error when something bad happen, such as bad internet network, server issues, etc. All we need is something that can make us retry the action. Trying is always good, don't give up with only trying once.


### Example:
```
stmtOut, err := db.Prepare("SELECT squareNumber FROM squarenum WHERE number = ?")
if err != nil {
	panic(err.Error())
}
var squareNum int

b := &jaen.Backoff{
    MaxTime: 10 * time.Second,
	MinTime: 100 * time.Millisecond,
}

for{
    err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
    if err != nil && b.getTotalTouch() < 5 { //trying until 5 times
        timeToSleep := b.Do()
        time.Sleep(timeToSleep)
    } else{
        if err != nil{
            log.Println("Still can't get it", b.getTotalTouch()) //Still error after trying several times, this case 5 times
        }
        b.Reset()
        break; //break the loop
    }
}
```