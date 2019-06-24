package imageconfig

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecr/ecriface"
	"github.com/containers/image/docker/reference"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

type mockEcr struct {
	mock.Mock
	ecriface.ECRAPI
}

func (m *mockEcr) GetAuthorizationToken(input *ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error) {
	f := m.Called(input)
	return f.Get(0).(*ecr.GetAuthorizationTokenOutput), f.Error(1)
}

func TestAwsEcrAuthenticator_Authenticate(t *testing.T) {
	api := &mockEcr{}
	api.
		On("GetAuthorizationToken", &ecr.GetAuthorizationTokenInput{RegistryIds: []*string{aws.String("0123")}}).
		Return(&ecr.GetAuthorizationTokenOutput{AuthorizationData: []*ecr.AuthorizationData{{AuthorizationToken: aws.String("token")}}}, nil)

	auth := &AwsEcrAuthenticator{Api: func(region string) ecriface.ECRAPI {
		assert.Equal(t, "ap-southeast-2", region)
		return api
	}}

	req, err := http.NewRequest("GET", "https://0123.dkr.ecr.ap-southeast-2.amazonaws.com/abc", nil)
	assert.NoError(t, err)

	ref, err := reference.ParseNamed("0123.dkr.ecr.ap-southeast-2.amazonaws.com/myrepo:tag")
	assert.NoError(t, err)

	err = auth.Authenticate(req, ref)
	assert.NoError(t, err)
	assert.Equal(t, "Basic token", req.Header.Get("Authorization"))
}
