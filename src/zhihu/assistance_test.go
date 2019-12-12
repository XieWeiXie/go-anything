package zhihu

import (
	"fmt"
	"testing"
)

func TestGetPageSource(t *testing.T) {
	content := hotZhiHuSource()
	fmt.Println(hotZhiHuList(content))
}
