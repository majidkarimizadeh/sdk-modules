/*
Leaseweb API for aggregation packs

Testing AggregationpackAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package aggregationpack

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/majidkarimizadeh/sdk-modules/aggregationpack"
)

func Test_aggregationpack_AggregationpackAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AggregationpackAPIService GetAggregationPack", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var aggregationPackId string

		resp, httpRes, err := apiClient.AggregationpackAPI.GetAggregationPack(context.Background(), aggregationPackId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AggregationpackAPIService GetAggregationPackList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AggregationpackAPI.GetAggregationPackList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}