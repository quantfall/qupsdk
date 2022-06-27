package updatesclientsdk

import (
	"io/ioutil"
	"os"
)

func GetLastUpdate(file string) (string, error) {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return "", err
		}
		defer f.Close()
		_, err = f.WriteString("0")
		if err != nil {
			return "", err
		}
		return "0", nil
	}

	f, err := ioutil.ReadFile(file)
	if err != nil {
		return "0", err
	}
	
	return string(f), nil
}

func SetLastUpdate(file string, last string) error {
	ioutil.WriteFile(file, []byte(last), 0755)
	return nil
}
