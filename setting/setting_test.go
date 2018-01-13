package setting

import (
	"fmt"
	"testing"
)

func TestSetting(t *testing.T) {
	if err := Setting(); err != nil {
		t.Errorf("Setting Error: %v \n", err)
	}
	fmt.Println(ServerConfig)

}
