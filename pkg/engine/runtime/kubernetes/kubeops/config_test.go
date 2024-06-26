package kubeops

import (
	"os"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"

	apiv1 "kusionstack.io/kusion/pkg/apis/api.kusion.io/v1"
)

func mockGetenv(result string) {
	mockey.Mock(os.Getenv).To(func(key string) string {
		return result
	}).Build()
}

func TestGetKubeConfig(t *testing.T) {
	resource := &apiv1.Resource{
		Extensions: map[string]any{
			"kubeConfig": "/home/test/kubeconfig",
		},
	}

	// Mock
	mockey.PatchConvey("test null env config", t, func() {
		mockGetenv("")
		assert.Equal(t, RecommendedKubeConfigFile, GetKubeConfig(nil))
	})
	mockey.PatchConvey("test env config", t, func() {
		mockGetenv("test")
		assert.Equal(t, "test", GetKubeConfig(resource))
	})
	mockey.PatchConvey("test resource config", t, func() {
		mockGetenv("")
		assert.Equal(t, "/home/test/kubeconfig", GetKubeConfig(resource))
	})
}
