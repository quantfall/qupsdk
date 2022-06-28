package updatesclientsdk

import "testing"

func TestGetNextUpdate(t *testing.T) {
	u, err := GetNextUpdate("utmstack", "0")
	if err != nil{
		t.Error(err)
	}
	t.Log(u)
}