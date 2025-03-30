// Code generated by mockery v2.49.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	types "github.com/agile-edge/go-mod-secrets/v4/pkg/types"
)

// SecretStoreClient is an autogenerated mock type for the SecretStoreClient type
type SecretStoreClient struct {
	mock.Mock
}

// BindUserToIdentity provides a mock function with given fields: token, identityId, authHandle, username
func (_m *SecretStoreClient) BindUserToIdentity(token string, identityId string, authHandle string, username string) error {
	ret := _m.Called(token, identityId, authHandle, username)

	if len(ret) == 0 {
		panic("no return value specified for BindUserToIdentity")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(token, identityId, authHandle, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckAuthMethodEnabled provides a mock function with given fields: token, mountPoint, authType
func (_m *SecretStoreClient) CheckAuthMethodEnabled(token string, mountPoint string, authType string) (bool, error) {
	ret := _m.Called(token, mountPoint, authType)

	if len(ret) == 0 {
		panic("no return value specified for CheckAuthMethodEnabled")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (bool, error)); ok {
		return rf(token, mountPoint, authType)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) bool); ok {
		r0 = rf(token, mountPoint, authType)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(token, mountPoint, authType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckIdentityKeyExists provides a mock function with given fields: token, keyName
func (_m *SecretStoreClient) CheckIdentityKeyExists(token string, keyName string) (bool, error) {
	ret := _m.Called(token, keyName)

	if len(ret) == 0 {
		panic("no return value specified for CheckIdentityKeyExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(token, keyName)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(token, keyName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(token, keyName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckSecretEngineInstalled provides a mock function with given fields: token, mountPoint, engine
func (_m *SecretStoreClient) CheckSecretEngineInstalled(token string, mountPoint string, engine string) (bool, error) {
	ret := _m.Called(token, mountPoint, engine)

	if len(ret) == 0 {
		panic("no return value specified for CheckSecretEngineInstalled")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (bool, error)); ok {
		return rf(token, mountPoint, engine)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) bool); ok {
		r0 = rf(token, mountPoint, engine)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(token, mountPoint, engine)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNamedIdentityKey provides a mock function with given fields: token, keyName, algorithm
func (_m *SecretStoreClient) CreateNamedIdentityKey(token string, keyName string, algorithm string) error {
	ret := _m.Called(token, keyName, algorithm)

	if len(ret) == 0 {
		panic("no return value specified for CreateNamedIdentityKey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(token, keyName, algorithm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateOrUpdateIdentity provides a mock function with given fields: token, name, metadata, policies
func (_m *SecretStoreClient) CreateOrUpdateIdentity(token string, name string, metadata map[string]string, policies []string) (string, error) {
	ret := _m.Called(token, name, metadata, policies)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrUpdateIdentity")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, map[string]string, []string) (string, error)); ok {
		return rf(token, name, metadata, policies)
	}
	if rf, ok := ret.Get(0).(func(string, string, map[string]string, []string) string); ok {
		r0 = rf(token, name, metadata, policies)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, map[string]string, []string) error); ok {
		r1 = rf(token, name, metadata, policies)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrUpdateIdentityRole provides a mock function with given fields: token, roleName, keyName, template, audience, jwtTTL
func (_m *SecretStoreClient) CreateOrUpdateIdentityRole(token string, roleName string, keyName string, template string, audience string, jwtTTL string) error {
	ret := _m.Called(token, roleName, keyName, template, audience, jwtTTL)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrUpdateIdentityRole")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string, string) error); ok {
		r0 = rf(token, roleName, keyName, template, audience, jwtTTL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateOrUpdateTokenRole provides a mock function with given fields: token, roleName, parameters
func (_m *SecretStoreClient) CreateOrUpdateTokenRole(token string, roleName string, parameters map[string]interface{}) error {
	ret := _m.Called(token, roleName, parameters)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrUpdateTokenRole")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, map[string]interface{}) error); ok {
		r0 = rf(token, roleName, parameters)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateOrUpdateUser provides a mock function with given fields: token, mountPoint, username, password, tokenTTL, tokenPolicies
func (_m *SecretStoreClient) CreateOrUpdateUser(token string, mountPoint string, username string, password string, tokenTTL string, tokenPolicies []string) error {
	ret := _m.Called(token, mountPoint, username, password, tokenTTL, tokenPolicies)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrUpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string, []string) error); ok {
		r0 = rf(token, mountPoint, username, password, tokenTTL, tokenPolicies)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateToken provides a mock function with given fields: token, parameters
func (_m *SecretStoreClient) CreateToken(token string, parameters map[string]interface{}) (map[string]interface{}, error) {
	ret := _m.Called(token, parameters)

	if len(ret) == 0 {
		panic("no return value specified for CreateToken")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) (map[string]interface{}, error)); ok {
		return rf(token, parameters)
	}
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(token, parameters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, map[string]interface{}) error); ok {
		r1 = rf(token, parameters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTokenByRole provides a mock function with given fields: token, role, parameters
func (_m *SecretStoreClient) CreateTokenByRole(token string, role string, parameters map[string]interface{}) (map[string]interface{}, error) {
	ret := _m.Called(token, role, parameters)

	if len(ret) == 0 {
		panic("no return value specified for CreateTokenByRole")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, map[string]interface{}) (map[string]interface{}, error)); ok {
		return rf(token, role, parameters)
	}
	if rf, ok := ret.Get(0).(func(string, string, map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(token, role, parameters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, map[string]interface{}) error); ok {
		r1 = rf(token, role, parameters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteIdentity provides a mock function with given fields: token, name
func (_m *SecretStoreClient) DeleteIdentity(token string, name string) error {
	ret := _m.Called(token, name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteIdentity")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(token, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: token, mountPoint, username
func (_m *SecretStoreClient) DeleteUser(token string, mountPoint string, username string) error {
	ret := _m.Called(token, mountPoint, username)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(token, mountPoint, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnableKVSecretEngine provides a mock function with given fields: token, mountPoint, kvVersion
func (_m *SecretStoreClient) EnableKVSecretEngine(token string, mountPoint string, kvVersion string) error {
	ret := _m.Called(token, mountPoint, kvVersion)

	if len(ret) == 0 {
		panic("no return value specified for EnableKVSecretEngine")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(token, mountPoint, kvVersion)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnablePasswordAuth provides a mock function with given fields: token, mountPoint
func (_m *SecretStoreClient) EnablePasswordAuth(token string, mountPoint string) error {
	ret := _m.Called(token, mountPoint)

	if len(ret) == 0 {
		panic("no return value specified for EnablePasswordAuth")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(token, mountPoint)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetIdentityByEntityId provides a mock function with given fields: token, entityId
func (_m *SecretStoreClient) GetIdentityByEntityId(token string, entityId string) (types.EntityMetadata, error) {
	ret := _m.Called(token, entityId)

	if len(ret) == 0 {
		panic("no return value specified for GetIdentityByEntityId")
	}

	var r0 types.EntityMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (types.EntityMetadata, error)); ok {
		return rf(token, entityId)
	}
	if rf, ok := ret.Get(0).(func(string, string) types.EntityMetadata); ok {
		r0 = rf(token, entityId)
	} else {
		r0 = ret.Get(0).(types.EntityMetadata)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(token, entityId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HealthCheck provides a mock function with given fields:
func (_m *SecretStoreClient) HealthCheck() (int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HealthCheck")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields: secretThreshold, secretShares
func (_m *SecretStoreClient) Init(secretThreshold int, secretShares int) (types.InitResponse, error) {
	ret := _m.Called(secretThreshold, secretShares)

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 types.InitResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (types.InitResponse, error)); ok {
		return rf(secretThreshold, secretShares)
	}
	if rf, ok := ret.Get(0).(func(int, int) types.InitResponse); ok {
		r0 = rf(secretThreshold, secretShares)
	} else {
		r0 = ret.Get(0).(types.InitResponse)
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(secretThreshold, secretShares)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InstallPolicy provides a mock function with given fields: token, policyName, policyDocument
func (_m *SecretStoreClient) InstallPolicy(token string, policyName string, policyDocument string) error {
	ret := _m.Called(token, policyName, policyDocument)

	if len(ret) == 0 {
		panic("no return value specified for InstallPolicy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(token, policyName, policyDocument)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InternalServiceLogin provides a mock function with given fields: token, authEngine, username, password
func (_m *SecretStoreClient) InternalServiceLogin(token string, authEngine string, username string, password string) (map[string]interface{}, error) {
	ret := _m.Called(token, authEngine, username, password)

	if len(ret) == 0 {
		panic("no return value specified for InternalServiceLogin")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) (map[string]interface{}, error)); ok {
		return rf(token, authEngine, username, password)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, string) map[string]interface{}); ok {
		r0 = rf(token, authEngine, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(token, authEngine, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTokenAccessors provides a mock function with given fields: token
func (_m *SecretStoreClient) ListTokenAccessors(token string) ([]string, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for ListTokenAccessors")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]string, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupAuthHandle provides a mock function with given fields: token, mountPoint
func (_m *SecretStoreClient) LookupAuthHandle(token string, mountPoint string) (string, error) {
	ret := _m.Called(token, mountPoint)

	if len(ret) == 0 {
		panic("no return value specified for LookupAuthHandle")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(token, mountPoint)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(token, mountPoint)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(token, mountPoint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupIdentity provides a mock function with given fields: token, name
func (_m *SecretStoreClient) LookupIdentity(token string, name string) (string, error) {
	ret := _m.Called(token, name)

	if len(ret) == 0 {
		panic("no return value specified for LookupIdentity")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(token, name)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(token, name)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(token, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupToken provides a mock function with given fields: token
func (_m *SecretStoreClient) LookupToken(token string) (types.TokenMetadata, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for LookupToken")
	}

	var r0 types.TokenMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (types.TokenMetadata, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) types.TokenMetadata); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(types.TokenMetadata)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupTokenAccessor provides a mock function with given fields: token, accessor
func (_m *SecretStoreClient) LookupTokenAccessor(token string, accessor string) (types.TokenMetadata, error) {
	ret := _m.Called(token, accessor)

	if len(ret) == 0 {
		panic("no return value specified for LookupTokenAccessor")
	}

	var r0 types.TokenMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (types.TokenMetadata, error)); ok {
		return rf(token, accessor)
	}
	if rf, ok := ret.Get(0).(func(string, string) types.TokenMetadata); ok {
		r0 = rf(token, accessor)
	} else {
		r0 = ret.Get(0).(types.TokenMetadata)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(token, accessor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegenRootToken provides a mock function with given fields: keys
func (_m *SecretStoreClient) RegenRootToken(keys []string) (string, error) {
	ret := _m.Called(keys)

	if len(ret) == 0 {
		panic("no return value specified for RegenRootToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) (string, error)); ok {
		return rf(keys)
	}
	if rf, ok := ret.Get(0).(func([]string) string); ok {
		r0 = rf(keys)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevokeToken provides a mock function with given fields: token
func (_m *SecretStoreClient) RevokeToken(token string) error {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for RevokeToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RevokeTokenAccessor provides a mock function with given fields: token, accessor
func (_m *SecretStoreClient) RevokeTokenAccessor(token string, accessor string) error {
	ret := _m.Called(token, accessor)

	if len(ret) == 0 {
		panic("no return value specified for RevokeTokenAccessor")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(token, accessor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unseal provides a mock function with given fields: keysBase64
func (_m *SecretStoreClient) Unseal(keysBase64 []string) error {
	ret := _m.Called(keysBase64)

	if len(ret) == 0 {
		panic("no return value specified for Unseal")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(keysBase64)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSecretStoreClient creates a new instance of SecretStoreClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSecretStoreClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *SecretStoreClient {
	mock := &SecretStoreClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
