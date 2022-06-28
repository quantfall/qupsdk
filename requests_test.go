package qupsdk

import "testing"

func TestGetNextUpdate(t *testing.T) {
	u, err := GetNextUpdate("utmstack", "0")
	if err != nil {
		t.Error(err)
	}
	t.Log(u)
}

func TestDownloadFile(t *testing.T) {
	err := DownloadFiles("utmstack", UpdateResponse{
		Version: 2022062800,
		Files: []File{
			{
				Name: "2022062800.sh",
				Sum:  "e7356fe01e251c9a1953a74b4eedf26f",
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
}
