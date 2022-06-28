package updatesclientsdk

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func RunUpdate(update UpdateResponse) error {
	return exec.Command(filepath.Join(fmt.Sprint(update.Version), update.Script)).Run()
}
