package updatesclientsdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
		return UpdateResponse{}, fmt.Errorf("there is no next update")
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
