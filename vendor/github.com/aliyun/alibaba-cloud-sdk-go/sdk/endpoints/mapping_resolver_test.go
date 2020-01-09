package endpoints

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMappingResolver_TryResolve(t *testing.T) {

	regionId := "cn-hangzhou"
	productId := "ecs"
	endpoint := GetEndpointFromMap(regionId, productId)
	assert.Equal(t, "", endpoint)

	AddEndpointMapping("cn-hangzhou", "Ecs", "unreachable.aliyuncs.com")

	endpoint = GetEndpointFromMap(regionId, productId)
	assert.Equal(t, "unreachable.aliyuncs.com", endpoint)
}

func Test_MappingResolveConcurrent(t *testing.T) {
	current := len(endpointMapping.endpoint)
	cnt := 50
	var wg sync.WaitGroup
	for i := 0; i < cnt; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			endpoint := fmt.Sprintf("ecs#cn-hangzhou%d", k)
			for j := 0; j < 50; j++ {
				err := AddEndpointMapping(fmt.Sprintf("cn-hangzhou%d", k), "ecs", endpoint)
				assert.Nil(t, err)
				assert.Equal(t, endpoint, GetEndpointFromMap(fmt.Sprintf("cn-hangzhou%d", k), "ecs"))
			}
		}(i)
	}
	wg.Wait()
	assert.Equal(t, (current + cnt), len(endpointMapping.endpoint))
	// hit cache and concurrent get
	for i := 0; i < cnt; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			endpoint := fmt.Sprintf("ecs#cn-hangzhou%d", k)
			for j := 0; j < cnt; j++ {
				assert.Equal(t, endpoint, GetEndpointFromMap(fmt.Sprintf("cn-hangzhou%d", k), "ecs"))
				err := AddEndpointMapping(fmt.Sprintf("cn-hangzhou%d", k), "ecs", endpoint)
				assert.Nil(t, err)
			}
		}(i)
	}
	wg.Wait()
	assert.Equal(t, (current + cnt), len(endpointMapping.endpoint))
}
