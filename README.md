# gobay

A simple and direct eBay Trading API SDK in Golang. 

The planned API in your go code will be:

```go
import "gobay"

cnf, err := fileGetContents("test_data/test.yml")
if err != nil {
    t.Errorf("Failed to load test.yml %v\n", err)
}
call := NewEbayCallEx(cnf)
p := call.NewProduct()

p.Title = "My Fancy Product"

call.AddProcut(p)

res,err := call.Execute()

if res.Success() {
    fmt.Printf("Sent %s\n",p.Title)
}

```
