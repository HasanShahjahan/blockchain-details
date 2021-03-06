package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlock_InValidNetwork(t *testing.T) {
	url := "http://localhost:8000/api/block/HASAN/000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf"
	actual := ApiHelper(t, url)
	fmt.Println(actual)
	assert.Equal(t, false, actual, "Should be same")
}
func TestBlock_InValidBlockHash(t *testing.T) {
	url := "http://localhost:8000/api/block/BTC/00034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf"
	actual := ApiHelper(t, url)
	assert.Equal(t, false, actual, "Should be same")
}
func TestBlock_InValidNetworkAndBlockHash(t *testing.T) {
	url := "http://localhost:8000/api/block/HASAN/000UU034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf"
	actual := ApiHelper(t, url)
	assert.Equal(t, false, actual, "Should be same")
}
func TestBlock_ValidResponse(t *testing.T) {
	url := "http://localhost:8000/api/block/BTC/000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf"
	actual := ApiHelper(t, url)
	assert.Equal(t, true, actual, "Should be same")
}