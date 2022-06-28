package qup-sdk

import "testing"

func TestRunUpdate(t *testing.T) {
	err := RunUpdate(UpdateResponse{
		Version: 2022062800,
		Script: "2022062800.sh",
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
