# Gobay - eBay Trading SDK in Go 
## Incomplete but functioning - pre-alpha stage contribute if you can.

[![Build Status](https://travis-ci.org/jasonknight/gobay.svg?branch=master)](https://travis-ci.org/jasonknight/gobay)

A simple and direct eBay Trading API SDK in Golang. 

The planned API in your go code will be:

```go
import "gobay"

var results []Result
cnf, err := fileGetContents("blahblah.yml") //you need to define this function, 
// or load a file your preferred way...
if err != nil {
    t.Errorf("Failed to load test.yml %v\n", err)
}
ebay,err := gobay.NewEbayCallEx(cnf)
if err != nil {
    panic(err)
}
ebay.SetCallname("AddItems")
p := ebay.NewProduct()

p.Title = "My Fancy Product"
//... more product setting here stuff here
// or you can load with
// p.FromYAML([]byte(yaml_string_you_load))


ebay.AddProduct(p)

err := ebay.Execute(&results)

for _,res := range results {
    if res.Success() {
        fmt.Printf("Sent %+v\n",p)
    } else {
        fmt.Printf("Oops, failed\n%+v\n",res)
    }
}

```
[![Become A Patron](https://github.com/jasonknight/gobay/raw/master/assets/patreon.png)](https://www.patreon.com/user?u=4141497)