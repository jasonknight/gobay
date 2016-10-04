# gobay
[![Build Status](https://travis-ci.org/jasonknight/gobay.svg?branch=master)](https://travis-ci.org/jasonknight/gobay)

[![Become A Patron](https://github.com/jasonknight/gobay/raw/master/assets/patreon.png)](https://www.patreon.com/user?u=4141497)

A simple and direct eBay Trading API SDK in Golang. 

The planned API in your go code will be:

```go
import "gobay"

cnf, err := fileGetContents("blahblah.yml") //you need to define this function
if err != nil {
    t.Errorf("Failed to load test.yml %v\n", err)
}
call,err := gobay.NewEbayCallEx(cnf)
if err != nil {
    panic(err)
}
call.SetCallname("AddItems")
p := call.NewProduct()

p.Title = "My Fancy Product"

call.AddProduct(p)

results,err := call.Execute()

for _,res := range results {
    if res.Success() {
        fmt.Printf("Sent %+v\n",p)
    } else {
        fmt.Printf("Oops, failed\n%+v\n",res)
    }
}

```
