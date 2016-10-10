package main

import (
    "io/ioutil"
    "net/http"
    "os"
    "bufio"
    "log"
    "flag"
    "github.com/satori/go.uuid"
    "crypto/rand"
    "fmt"
)

var (
    Log *log.Logger
    logpath = flag.String("log","/tmp/gobay_notifications.log","Log path")
    bind = flag.String("bind",":8000","where to bind")
)
func createLogFile(logpath string) error {
    if fileExists(logpath) {
        return nil
    }
    file, err := os.Create(logpath)
    if err != nil {
        return err
    }
    Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
    return nil
}
func fileExists(p string) bool {
    if _, err := os.Stat(p); os.IsNotExist(err) {
        return false
    }
    return true
}
func pseudoUUID() (string, error) {

    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {

        return "", err
    }

    id := uuid.NewV4()

    return id.String(), nil
}
func filePutContents(p string, txt string) error {
    f, err := os.Create(p)
    if err != nil {
        return err
    }
    w := bufio.NewWriter(f)
    _, err = w.WriteString(txt)
    w.Flush()
    return nil
}
func save_request(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Server", "Gobay Notification Server")   
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        // do something with err
        
        w.WriteHeader(500)
        Log.Printf("could not read body %s",err)
        return
    }
    // write out data

    fname, err := pseudoUUID()
    if err != nil {
        w.WriteHeader(500)
        Log.Printf("could not generate pseudoUUID %s",err)
        return
    }
    err = filePutContents(fmt.Sprintf("%s.xml",fname),string(b))
    if err != nil {
        Log.Printf("could not save request as a file %s",err)
        w.WriteHeader(500)
        return
    }
    w.WriteHeader(200)
    w.Write([]byte("OK"))

}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
    flag.Parse()
    err := createLogFile(*logpath)
    if err != nil {
        fmt.Printf("could not create logfile %s",err)
        return
    }
    server := http.Server{
        Addr:    *bind,
        Handler: &theHndlr{},
    }

    mux = make(map[string]func(http.ResponseWriter, *http.Request))
    mux["/"] = save_request
    fmt.Printf("Listening: %s\n",*bind)
    server.ListenAndServe()
}

type theHndlr struct{}

func (*theHndlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
    // fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
    if h, ok := mux[r.URL.String()]; ok {
        h(w, r)
        return
    }
}
