package gobay

import "fmt"
import "os"
import "strings"
import "io"
import "io/ioutil"
import "bufio"
import "bytes"
import "text/template"
import "crypto/rand"
import "gopkg.in/yaml.v2"
import "github.com/satori/go.uuid"

func version() string {
	return "v1.0"
}

func fileExists(p string) bool {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return false
	}
	return true
}

func banner() {
	fmt.Println("*******************************************************")
	fmt.Printf("*                       Gobay %s                     *\n", version())
	fmt.Println("*******************************************************")
}
func fileType(p string) string {
	parts := strings.Split(p, ".")
	ftype := parts[len(parts)-1]
	return ftype
}
func copyFile(d string, p string) error {
	in, err := os.Open(p)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(d)
	if err != nil {
		return err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	err = out.Sync()
	return err
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
func fileGetContents(p string) ([]byte, error) {
	return ioutil.ReadFile(p)
}
func compileGoString(name string, text string, obj interface{}, fmap template.FuncMap) (string, error) {
	tmpl, err := template.New(name).Funcs(fmap).Parse(text)
	if err != nil {
		return "", err
	}
	var can bytes.Buffer
	err = tmpl.Execute(&can, obj)
	if err != nil {
		return "", err
	}
	return can.String(), nil
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

func WrapCall(name string, pre string, text string, post string) string {
	s := `<?xml version="1.0" encoding="utf-8"?>
<%sRequest xmlns="urn:ebay:apis:eBLBaseComponents">
	<RequesterCredentials> 
		<eBayAuthToken>{{ .EbayAuthToken }}</eBayAuthToken> 
	</RequesterCredentials> 
	<MessageID>{{ .MessageID }}</MessageID>
	<Version>{{ .CompatLevel }}</Version>
	<WarningLevel>{{ .WarningLevel }}</WarningLevel>
%s
%s
%s
</%sRequest>
`
	return fmt.Sprintf(s, name, pre, text, post, name)
}

func LoadConfiguration(y []byte, e *map[interface{}]interface{}) error {
	return yaml.Unmarshal(y, e)
}
func InStringSlice(s []string, r string) bool {

	for _, c := range s {
		if r == c {
			return true
		}
	}

	return false
}

var runSandboxTests bool

func shouldRunSandbox() bool {
	runSandboxTests = false // turn this to true if you have
	// setup your sandbox account
	if fileExists("../secret.yml") == false {
		fmt.Printf("Not Running sandbox tests because no secret\n")
		return false
	}
	if runSandboxTests == false {
		fmt.Printf("Not Running sandbox tests because of settings\n")
	}
	return runSandboxTests
}
