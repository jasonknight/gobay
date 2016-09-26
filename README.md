# gobay

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
p := call.NewProduct()

p.Title = "My Fancy Product"

call.AddProduct(p)

res,err := call.Execute()

if res.Success() {
    fmt.Printf("Sent %s\n",p.Title)
}

```
