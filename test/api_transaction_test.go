/*
L3vels Api

Testing TransactionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package l3vels-sdk

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/chkhikvadze/l3vels-sdk"
)

func Test_l3vels-sdk_TransactionApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TransactionApiService TransactionById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var projectId string

		resp, httpRes, err := apiClient.TransactionApi.TransactionById(context.Background(), id, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionApiService Transactions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TransactionApi.Transactions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionApiService Webhook", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.TransactionApi.Webhook(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
