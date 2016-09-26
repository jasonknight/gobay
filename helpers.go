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
func pseudoUUID() (string,error) {

    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {

        return "",err
    }

    uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

    return uuid,nil
}