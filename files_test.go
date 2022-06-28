package qupsdk

import (
	"fmt"
	"testing"
	"time"
)

func TestGetLastUpdate(t *testing.T){
	v, err := GetLastUpdate("update.txt")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestSetLastUpdate(t *testing.T){
	err := SetLastUpdate("update.txt", fmt.Sprint(time.Now().UTC().Unix()))
	if err != nil {
		t.Error(err)
	}
}
