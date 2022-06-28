package updatesclientsdk

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type UpdateResponse struct {
	Version   uint   `json:"version"`
	Alpha     bool   `json:"alpha"`
	Beta      bool   `json:"beta"`
	Candidate bool   `json:"candidate"`
	Release   bool   `json:"release"`
	Script    string `json:"script"`
	Files     []File `json:"files"`
}

type File struct {
	Name string `json:"name"`
	Sum  string `json:"sum"`
}

func GetNextUpdate(project, last string) (UpdateResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprint("https://updates.quantfall.com/qup/v1/public/update/", project, "/", last), nil)
	if err != nil {
		return UpdateResponse{}, err
	}

	client := &http.Client{}

	var update UpdateResponse

	resp, err := client.Do(req)
	if err != nil {
		return UpdateResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return UpdateResponse{}, fmt.Errorf("no next update")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return UpdateResponse{}, err
	}

	err = json.Unmarshal(body, &update)
	if err != nil {
		return UpdateResponse{}, err
	}

	return update, nil
}

func DownloadFiles(project string, update UpdateResponse) error {
	for _, file := range(update.Files){
		for i := 0; i < 2; i++ {
			err := downloadFile(project, file.Name, update.Version)
			if err != nil {
				return err
			}
	
			f, err := ioutil.ReadFile(filepath.Join(fmt.Sprint(update.Version), file.Name))
			
			hash := md5.Sum(f)
			
			if hex.EncodeToString(hash[:]) != file.Sum {
				return fmt.Errorf("file corrupted")
			}
		}
	}

	return nil
}

func downloadFile(project, file string, version uint) error {
	req, err := http.NewRequest("GET", fmt.Sprint("https://updates.quantfall.com/qup/v1/public/file/", project, "/", version, "/", file), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = os.MkdirAll(fmt.Sprint(version), 0755)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(fmt.Sprint(version), file), body, 0755)
	if err != nil {
		return err
	}

	return nil
}
