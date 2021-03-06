// Code generated by counterfeiter. DO NOT EDIT.
package v2actionfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/util/configv3"
)

type FakeConfig struct {
	AccessTokenStub        func() string
	accessTokenMutex       sync.RWMutex
	accessTokenArgsForCall []struct {
	}
	accessTokenReturns struct {
		result1 string
	}
	accessTokenReturnsOnCall map[int]struct {
		result1 string
	}
	BinaryNameStub        func() string
	binaryNameMutex       sync.RWMutex
	binaryNameArgsForCall []struct {
	}
	binaryNameReturns struct {
		result1 string
	}
	binaryNameReturnsOnCall map[int]struct {
		result1 string
	}
	DialTimeoutStub        func() time.Duration
	dialTimeoutMutex       sync.RWMutex
	dialTimeoutArgsForCall []struct {
	}
	dialTimeoutReturns struct {
		result1 time.Duration
	}
	dialTimeoutReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	PollingIntervalStub        func() time.Duration
	pollingIntervalMutex       sync.RWMutex
	pollingIntervalArgsForCall []struct {
	}
	pollingIntervalReturns struct {
		result1 time.Duration
	}
	pollingIntervalReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	RefreshTokenStub        func() string
	refreshTokenMutex       sync.RWMutex
	refreshTokenArgsForCall []struct {
	}
	refreshTokenReturns struct {
		result1 string
	}
	refreshTokenReturnsOnCall map[int]struct {
		result1 string
	}
	SSHOAuthClientStub        func() string
	sSHOAuthClientMutex       sync.RWMutex
	sSHOAuthClientArgsForCall []struct {
	}
	sSHOAuthClientReturns struct {
		result1 string
	}
	sSHOAuthClientReturnsOnCall map[int]struct {
		result1 string
	}
	SetAccessTokenStub        func(string)
	setAccessTokenMutex       sync.RWMutex
	setAccessTokenArgsForCall []struct {
		arg1 string
	}
	SetRefreshTokenStub        func(string)
	setRefreshTokenMutex       sync.RWMutex
	setRefreshTokenArgsForCall []struct {
		arg1 string
	}
	SetTargetInformationStub        func(configv3.TargetInformationArgs)
	setTargetInformationMutex       sync.RWMutex
	setTargetInformationArgsForCall []struct {
		arg1 configv3.TargetInformationArgs
	}
	SetTokenInformationStub        func(string, string, string)
	setTokenInformationMutex       sync.RWMutex
	setTokenInformationArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	SetUAAClientCredentialsStub        func(string, string)
	setUAAClientCredentialsMutex       sync.RWMutex
	setUAAClientCredentialsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	SetUAAGrantTypeStub        func(string)
	setUAAGrantTypeMutex       sync.RWMutex
	setUAAGrantTypeArgsForCall []struct {
		arg1 string
	}
	SkipSSLValidationStub        func() bool
	skipSSLValidationMutex       sync.RWMutex
	skipSSLValidationArgsForCall []struct {
	}
	skipSSLValidationReturns struct {
		result1 bool
	}
	skipSSLValidationReturnsOnCall map[int]struct {
		result1 bool
	}
	StagingTimeoutStub        func() time.Duration
	stagingTimeoutMutex       sync.RWMutex
	stagingTimeoutArgsForCall []struct {
	}
	stagingTimeoutReturns struct {
		result1 time.Duration
	}
	stagingTimeoutReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	StartupTimeoutStub        func() time.Duration
	startupTimeoutMutex       sync.RWMutex
	startupTimeoutArgsForCall []struct {
	}
	startupTimeoutReturns struct {
		result1 time.Duration
	}
	startupTimeoutReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	TargetStub        func() string
	targetMutex       sync.RWMutex
	targetArgsForCall []struct {
	}
	targetReturns struct {
		result1 string
	}
	targetReturnsOnCall map[int]struct {
		result1 string
	}
	UAAGrantTypeStub        func() string
	uAAGrantTypeMutex       sync.RWMutex
	uAAGrantTypeArgsForCall []struct {
	}
	uAAGrantTypeReturns struct {
		result1 string
	}
	uAAGrantTypeReturnsOnCall map[int]struct {
		result1 string
	}
	UnsetOrganizationAndSpaceInformationStub        func()
	unsetOrganizationAndSpaceInformationMutex       sync.RWMutex
	unsetOrganizationAndSpaceInformationArgsForCall []struct {
	}
	UnsetSpaceInformationStub        func()
	unsetSpaceInformationMutex       sync.RWMutex
	unsetSpaceInformationArgsForCall []struct {
	}
	VerboseStub        func() (bool, []string)
	verboseMutex       sync.RWMutex
	verboseArgsForCall []struct {
	}
	verboseReturns struct {
		result1 bool
		result2 []string
	}
	verboseReturnsOnCall map[int]struct {
		result1 bool
		result2 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfig) AccessToken() string {
	fake.accessTokenMutex.Lock()
	ret, specificReturn := fake.accessTokenReturnsOnCall[len(fake.accessTokenArgsForCall)]
	fake.accessTokenArgsForCall = append(fake.accessTokenArgsForCall, struct {
	}{})
	fake.recordInvocation("AccessToken", []interface{}{})
	fake.accessTokenMutex.Unlock()
	if fake.AccessTokenStub != nil {
		return fake.AccessTokenStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.accessTokenReturns
	return fakeReturns.result1
}

func (fake *FakeConfig) AccessTokenCallCount() int {
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	return len(fake.accessTokenArgsForCall)
}

func (fake *FakeConfig) AccessTokenCalls(stub func() string) {
	fake.accessTokenMutex.Lock()
	defer fake.accessTokenMutex.Unlock()
	fake.AccessTokenStub = stub
}

func (fake *FakeConfig) AccessTokenReturns(result1 string) {
	fake.accessTokenMutex.Lock()
	defer fake.accessTokenMutex.Unlock()
	fake.AccessTokenStub = nil
	fake.accessTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) AccessTokenReturnsOnCall(i int, result1 string) {
	fake.accessTokenMutex.Lock()
	defer fake.accessTokenMutex.Unlock()
	fake.AccessTokenStub = nil
	if fake.accessTokenReturnsOnCall == nil {
		fake.accessTokenReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.accessTokenReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) BinaryName() string {
	fake.binaryNameMutex.Lock()
	ret, specificReturn := fake.binaryNameReturnsOnCall[len(fake.binaryNameArgsForCall)]
	fake.binaryNameArgsForCall = append(fake.binaryNameArgsForCall, struct {
	}{})
	fake.recordInvocation("BinaryName", []interface{}{})
	fake.binaryNameMutex.Unlock()
	if fake.BinaryNameStub != nil {
		return fake.BinaryNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.binaryNameReturns
	return fakeReturns.result1
}

func (fake *FakeConfig) BinaryNameCallCount() int {
	fake.binaryNameMutex.RLock()
	defer fake.binaryNameMutex.RUnlock()
	return len(fake.binaryNameArgsForCall)
}

func (fake *FakeConfig) BinaryNameCalls(stub func() string) {
	fake.binaryNameMutex.Lock()
	defer fake.binaryNameMutex.Unlock()
	fake.BinaryNameStub = stub
}

func (fake *FakeConfig) BinaryNameReturns(result1 string) {
	fake.binaryNameMutex.Lock()
	defer fake.binaryNameMutex.Unlock()
	fake.BinaryNameStub = nil
	fake.binaryNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) BinaryNameReturnsOnCall(i int, result1 string) {
	fake.binaryNameMutex.Lock()
	defer fake.binaryNameMutex.Unlock()
	fake.BinaryNameStub = nil
	if fake.binaryNameReturnsOnCall == nil {
		fake.binaryNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.binaryNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) DialTimeout() time.Duration {
	fake.dialTimeoutMutex.Lock()
	ret, specificReturn := fake.dialTimeoutReturnsOnCall[len(fake.dialTimeoutArgsForCall)]
	fake.dialTimeoutArgsForCall = append(fake.dialTimeoutArgsForCall, struct {
	}{})
	fake.recordInvocation("DialTimeout", []interface{}{})
	fake.dialTimeoutMutex.Unlock()
	if fake.DialTimeoutStub != nil {
		return fake.DialTimeoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.dialTimeoutReturns
	return fakeReturns.result1
}

func (fake *FakeConfig) DialTimeoutCallCount() int {
	fake.dialTimeoutMutex.RLock()
	defer fake.dialTimeoutMutex.RUnlock()
	return len(fake.dialTimeoutArgsForCall)
}

func (fake *FakeConfig) DialTimeoutCalls(stub func() time.Duration) {
	fake.dialTimeoutMutex.Lock()
	defer fake.dialTimeoutMutex.Unlock()
	fake.DialTimeoutStub = stub
}

func (fake *FakeConfig) DialTimeoutReturns(result1 time.Duration) {
	fake.dialTimeoutMutex.Lock()
	defer fake.dialTimeoutMutex.Unlock()
	fake.DialTimeoutStub = nil
	fake.dialTimeoutReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) DialTimeoutReturnsOnCall(i int, result1 time.Duration) {
	fake.dialTimeoutMutex.Lock()
	defer fake.dialTimeoutMutex.Unlock()
	fake.DialTimeoutStub = nil
	if fake.dialTimeoutReturnsOnCall == nil {
		fake.dialTimeoutReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.dialTimeoutReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) PollingInterval() time.Duration {
	fake.pollingIntervalMutex.Lock()
	ret, specificReturn := fake.pollingIntervalReturnsOnCall[len(fake.pollingIntervalArgsForCall)]
	fake.pollingIntervalArgsForCall = append(fake.pollingIntervalArgsForCall, struct {
	}{})
	fake.recordInvocation("PollingInterval", []interface{}{})
	fake.pollingIntervalMutex.Unlock()
	if fake.PollingIntervalStub != nil {
		return fake.PollingIntervalStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pollingIntervalReturns
	return fakeReturns.result1
}

func (fake *FakeConfig) PollingIntervalCallCount() int {
	fake.pollingIntervalMutex.RLock()
	defer fake.pollingIntervalMutex.RUnlock()
	return len(fake.pollingIntervalArgsForCall)
}

func (fake *FakeConfig) PollingIntervalCalls(stub func() time.Duration) {
	fake.pollingIntervalMutex.Lock()
	defer fake.pollingIntervalMutex.Unlock()
	fake.PollingIntervalStub = stub
}

func (fake *FakeConfig) PollingIntervalReturns(result1 time.Duration) {
	fake.pollingIntervalMutex.Lock()
	defer fake.pollingIntervalMutex.Unlock()
	fake.PollingIntervalStub = nil
	fake.pollingIntervalReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) PollingIntervalReturnsOnCall(i int, result1 time.Duration) {
	fake.pollingIntervalMutex.Lock()
	defer fake.pollingIntervalMutex.Unlock()
	fake.PollingIntervalStub = nil
	if fake.pollingIntervalReturnsOnCall == nil {
		fake.pollingIntervalReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.pollingIntervalReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) RefreshToken() string {
	fake.refreshTokenMutex.Lock()
	ret, specificReturn := fake.refreshTokenReturnsOnCall[len(fake.refreshTokenArgsForCall)]
	fake.refreshTokenArgsForCall = append(fake.refreshTokenArgsForCall, struct {
	}{})
	fake.recordInvocation("RefreshToken", []interface{}{})
	fake.refreshTokenMutex.Unlock()
	if fake.RefreshTokenStub != nil {
		return fake.RefreshTokenStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.refreshTokenReturns
	return fakeReturns.result1
}

func (fake *FakeConfig) RefreshTokenCallCount() int {
	fake.refreshTokenMutex.RLock()
	defer fake.refreshTokenMutex.RUnlock()
	return len(fake.refreshTokenArgsForCall)
}

func (fake *FakeConfig) RefreshTokenCalls(stub func() string) {
	fake.refreshTokenMutex.Lock()
	defer fake.refreshTokenMutex.Unlock()
	fake.RefreshTokenStub = stub
}

func (fake *FakeConfig) RefreshTokenReturns(result1 string) {
	fake.refreshTokenMutex.Lock()
	defer fake.refreshTokenMutex.Unlock()
	fake.RefreshTokenStub = nil
	fake.refreshTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) RefreshTokenReturnsOnCall(i int, result1 string) {
	fake.refreshTokenMutex.Lock()
	defer fake.refreshTokenMutex.Unlock()
	fake.RefreshTokenStub = nil
	if fake.refreshTokenReturnsOnCall == nil {
		fake.refreshTokenReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.refreshTokenReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) SSHOAuthClient() string {
	fake.sSHOAuthClientMutex.Lock()
	ret, specificReturn := fake.sSHOAuthClientReturnsOnCall[len(fake.sSHOAuthClientArgsForCall)]
	fake.sSHOAuthClientArgsForCall = append(fake.sSHOAuthClientArgsForCall, struct {
	}{})
	fake.recordInvocation("SSHOAuthClient", []interface{}{})
	fake.sSHOAuthClientMutex.Unlock()
	if fake.SSHOAuthClientStub != nil {
		return fake.SSHOAuthClientStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sSHOAuthClientReturns
	return fakeReturns.result1
}

func (fake *FakeConfig) SSHOAuthClientCallCount() int {
	fake.sSHOAuthClientMutex.RLock()
	defer fake.sSHOAuthClientMutex.RUnlock()
	return len(fake.sSHOAuthClientArgsForCall)
}

func (fake *FakeConfig) SSHOAuthClientCalls(stub func() string) {
	fake.sSHOAuthClientMutex.Lock()
	defer fake.sSHOAuthClientMutex.Unlock()
	fake.SSHOAuthClientStub = stub
}

func (fake *FakeConfig) SSHOAuthClientReturns(result1 string) {
	fake.sSHOAuthClientMutex.Lock()
	defer fake.sSHOAuthClientMutex.Unlock()
	fake.SSHOAuthClientStub = nil
	fake.sSHOAuthClientReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) SSHOAuthClientReturnsOnCall(i int, result1 string) {
	fake.sSHOAuthClientMutex.Lock()
	defer fake.sSHOAuthClientMutex.Unlock()
	fake.SSHOAuthClientStub = nil
	if fake.sSHOAuthClientReturnsOnCall == nil {
		fake.sSHOAuthClientReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.sSHOAuthClientReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) SetAccessToken(arg1 string) {
	fake.setAccessTokenMutex.Lock()
	fake.setAccessTokenArgsForCall = append(fake.setAccessTokenArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetAccessToken", []interface{}{arg1})
	fake.setAccessTokenMutex.Unlock()
	if fake.SetAccessTokenStub != nil {
		fake.SetAccessTokenStub(arg1)
	}
}

func (fake *FakeConfig) SetAccessTokenCallCount() int {
	fake.setAccessTokenMutex.RLock()
	defer fake.setAccessTokenMutex.RUnlock()
	return len(fake.setAccessTokenArgsForCall)
}

func (fake *FakeConfig) SetAccessTokenCalls(stub func(string)) {
	fake.setAccessTokenMutex.Lock()
	defer fake.setAccessTokenMutex.Unlock()
	fake.SetAccessTokenStub = stub
}

func (fake *FakeConfig) SetAccessTokenArgsForCall(i int) string {
	fake.setAccessTokenMutex.RLock()
	defer fake.setAccessTokenMutex.RUnlock()
	argsForCall := fake.setAccessTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfig) SetRefreshToken(arg1 string) {
	fake.setRefreshTokenMutex.Lock()
	fake.setRefreshTokenArgsForCall = append(fake.setRefreshTokenArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetRefreshToken", []interface{}{arg1})
	fake.setRefreshTokenMutex.Unlock()
	if fake.SetRefreshTokenStub != nil {
		fake.SetRefreshTokenStub(arg1)
	}
}

func (fake *FakeConfig) SetRefreshTokenCallCount() int {
	fake.setRefreshTokenMutex.RLock()
	defer fake.setRefreshTokenMutex.RUnlock()
	return len(fake.setRefreshTokenArgsForCall)
}

func (fake *FakeConfig) SetRefreshTokenCalls(stub func(string)) {
	fake.setRefreshTokenMutex.Lock()
	defer fake.setRefreshTokenMutex.Unlock()
	fake.SetRefreshTokenStub = stub
}

func (fake *FakeConfig) SetRefreshTokenArgsForCall(i int) string {
	fake.setRefreshTokenMutex.RLock()
	defer fake.setRefreshTokenMutex.RUnlock()
	argsForCall := fake.setRefreshTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfig) SetTargetInformation(arg1 configv3.TargetInformationArgs) {
	fake.setTargetInformationMutex.Lock()
	fake.setTargetInformationArgsForCall = append(fake.setTargetInformationArgsForCall, struct {
		arg1 configv3.TargetInformationArgs
	}{arg1})
	fake.recordInvocation("SetTargetInformation", []interface{}{arg1})
	fake.setTargetInformationMutex.Unlock()
	if fake.SetTargetInformationStub != nil {
		fake.SetTargetInformationStub(arg1)
	}
}

func (fake *FakeConfig) SetTargetInformationCallCount() int {
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	return len(fake.setTargetInformationArgsForCall)
}

func (fake *FakeConfig) SetTargetInformationCalls(stub func(configv3.TargetInformationArgs)) {
	fake.setTargetInformationMutex.Lock()
	defer fake.setTargetInformationMutex.Unlock()
	fake.SetTargetInformationStub = stub
}

func (fake *FakeConfig) SetTargetInformationArgsForCall(i int) configv3.TargetInformationArgs {
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	argsForCall := fake.setTargetInformationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfig) SetTokenInformation(arg1 string, arg2 string, arg3 string) {
	fake.setTokenInformationMutex.Lock()
	fake.setTokenInformationArgsForCall = append(fake.setTokenInformationArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("SetTokenInformation", []interface{}{arg1, arg2, arg3})
	fake.setTokenInformationMutex.Unlock()
	if fake.SetTokenInformationStub != nil {
		fake.SetTokenInformationStub(arg1, arg2, arg3)
	}
}

func (fake *FakeConfig) SetTokenInformationCallCount() int {
	fake.setTokenInformationMutex.RLock()
	defer fake.setTokenInformationMutex.RUnlock()
	return len(fake.setTokenInformationArgsForCall)
}

func (fake *FakeConfig) SetTokenInformationCalls(stub func(string, string, string)) {
	fake.setTokenInformationMutex.Lock()
	defer fake.setTokenInformationMutex.Unlock()
	fake.SetTokenInformationStub = stub
}

func (fake *FakeConfig) SetTokenInformationArgsForCall(i int) (string, string, string) {
	fake.setTokenInformationMutex.RLock()
	defer fake.setTokenInformationMutex.RUnlock()
	argsForCall := fake.setTokenInformationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeConfig) SetUAAClientCredentials(arg1 string, arg2 string) {
	fake.setUAAClientCredentialsMutex.Lock()
	fake.setUAAClientCredentialsArgsForCall = append(fake.setUAAClientCredentialsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("SetUAAClientCredentials", []interface{}{arg1, arg2})
	fake.setUAAClientCredentialsMutex.Unlock()
	if fake.SetUAAClientCredentialsStub != nil {
		fake.SetUAAClientCredentialsStub(arg1, arg2)
	}
}

func (fake *FakeConfig) SetUAAClientCredentialsCallCount() int {
	fake.setUAAClientCredentialsMutex.RLock()
	defer fake.setUAAClientCredentialsMutex.RUnlock()
	return len(fake.setUAAClientCredentialsArgsForCall)
}

func (fake *FakeConfig) SetUAAClientCredentialsCalls(stub func(string, string)) {
	fake.setUAAClientCredentialsMutex.Lock()
	defer fake.setUAAClientCredentialsMutex.Unlock()
	fake.SetUAAClientCredentialsStub = stub
}

func (fake *FakeConfig) SetUAAClientCredentialsArgsForCall(i int) (string, string) {
	fake.setUAAClientCredentialsMutex.RLock()
	defer fake.setUAAClientCredentialsMutex.RUnlock()
	argsForCall := fake.setUAAClientCredentialsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeConfig) SetUAAGrantType(arg1 string) {
	fake.setUAAGrantTypeMutex.Lock()
	fake.setUAAGrantTypeArgsForCall = append(fake.setUAAGrantTypeArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetUAAGrantType", []interface{}{arg1})
	fake.setUAAGrantTypeMutex.Unlock()
	if fake.SetUAAGrantTypeStub != nil {
		fake.SetUAAGrantTypeStub(arg1)
	}
}

func (fake *FakeConfig) SetUAAGrantTypeCallCount() int {
	fake.setUAAGrantTypeMutex.RLock()
	defer fake.setUAAGrantTypeMutex.RUnlock()
	return len(fake.setUAAGrantTypeArgsForCall)
}

func (fake *FakeConfig) SetUAAGrantTypeCalls(stub func(string)) {
	fake.setUAAGrantTypeMutex.Lock()
	defer fake.setUAAGrantTypeMutex.Unlock()
	fake.SetUAAGrantTypeStub = stub
}

func (fake *FakeConfig) SetUAAGrantTypeArgsForCall(i int) string {
	fake.setUAAGrantTypeMutex.RLock()
	defer fake.setUAAGrantTypeMutex.RUnlock()
	argsForCall := fake.setUAAGrantTypeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfig) SkipSSLValidation() bool {
	fake.skipSSLValidationMutex.Lock()
	ret, specificReturn := fake.skipSSLValidationReturnsOnCall[len(fake.skipSSLValidationArgsForCall)]
	fake.skipSSLValidationArgsForCall = append(fake.skipSSLValidationArgsForCall, struct {
	}{})
	fake.recordInvocation("SkipSSLValidation", []interface{}{})
	fake.skipSSLValidationMutex.Unlock()
	if fake.SkipSSLValidationStub != nil {
		return fake.SkipSSLValidationStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.skipSSLValidationReturns
	return fakeReturns.result1
}

func (fake *FakeConfig) SkipSSLValidationCallCount() int {
	fake.skipSSLValidationMutex.RLock()
	defer fake.skipSSLValidationMutex.RUnlock()
	return len(fake.skipSSLValidationArgsForCall)
}

func (fake *FakeConfig) SkipSSLValidationCalls(stub func() bool) {
	fake.skipSSLValidationMutex.Lock()
	defer fake.skipSSLValidationMutex.Unlock()
	fake.SkipSSLValidationStub = stub
}

func (fake *FakeConfig) SkipSSLValidationReturns(result1 bool) {
	fake.skipSSLValidationMutex.Lock()
	defer fake.skipSSLValidationMutex.Unlock()
	fake.SkipSSLValidationStub = nil
	fake.skipSSLValidationReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeConfig) SkipSSLValidationReturnsOnCall(i int, result1 bool) {
	fake.skipSSLValidationMutex.Lock()
	defer fake.skipSSLValidationMutex.Unlock()
	fake.SkipSSLValidationStub = nil
	if fake.skipSSLValidationReturnsOnCall == nil {
		fake.skipSSLValidationReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.skipSSLValidationReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeConfig) StagingTimeout() time.Duration {
	fake.stagingTimeoutMutex.Lock()
	ret, specificReturn := fake.stagingTimeoutReturnsOnCall[len(fake.stagingTimeoutArgsForCall)]
	fake.stagingTimeoutArgsForCall = append(fake.stagingTimeoutArgsForCall, struct {
	}{})
	fake.recordInvocation("StagingTimeout", []interface{}{})
	fake.stagingTimeoutMutex.Unlock()
	if fake.StagingTimeoutStub != nil {
		return fake.StagingTimeoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stagingTimeoutReturns
	return fakeReturns.result1
}

func (fake *FakeConfig) StagingTimeoutCallCount() int {
	fake.stagingTimeoutMutex.RLock()
	defer fake.stagingTimeoutMutex.RUnlock()
	return len(fake.stagingTimeoutArgsForCall)
}

func (fake *FakeConfig) StagingTimeoutCalls(stub func() time.Duration) {
	fake.stagingTimeoutMutex.Lock()
	defer fake.stagingTimeoutMutex.Unlock()
	fake.StagingTimeoutStub = stub
}

func (fake *FakeConfig) StagingTimeoutReturns(result1 time.Duration) {
	fake.stagingTimeoutMutex.Lock()
	defer fake.stagingTimeoutMutex.Unlock()
	fake.StagingTimeoutStub = nil
	fake.stagingTimeoutReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) StagingTimeoutReturnsOnCall(i int, result1 time.Duration) {
	fake.stagingTimeoutMutex.Lock()
	defer fake.stagingTimeoutMutex.Unlock()
	fake.StagingTimeoutStub = nil
	if fake.stagingTimeoutReturnsOnCall == nil {
		fake.stagingTimeoutReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.stagingTimeoutReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) StartupTimeout() time.Duration {
	fake.startupTimeoutMutex.Lock()
	ret, specificReturn := fake.startupTimeoutReturnsOnCall[len(fake.startupTimeoutArgsForCall)]
	fake.startupTimeoutArgsForCall = append(fake.startupTimeoutArgsForCall, struct {
	}{})
	fake.recordInvocation("StartupTimeout", []interface{}{})
	fake.startupTimeoutMutex.Unlock()
	if fake.StartupTimeoutStub != nil {
		return fake.StartupTimeoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.startupTimeoutReturns
	return fakeReturns.result1
}

func (fake *FakeConfig) StartupTimeoutCallCount() int {
	fake.startupTimeoutMutex.RLock()
	defer fake.startupTimeoutMutex.RUnlock()
	return len(fake.startupTimeoutArgsForCall)
}

func (fake *FakeConfig) StartupTimeoutCalls(stub func() time.Duration) {
	fake.startupTimeoutMutex.Lock()
	defer fake.startupTimeoutMutex.Unlock()
	fake.StartupTimeoutStub = stub
}

func (fake *FakeConfig) StartupTimeoutReturns(result1 time.Duration) {
	fake.startupTimeoutMutex.Lock()
	defer fake.startupTimeoutMutex.Unlock()
	fake.StartupTimeoutStub = nil
	fake.startupTimeoutReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) StartupTimeoutReturnsOnCall(i int, result1 time.Duration) {
	fake.startupTimeoutMutex.Lock()
	defer fake.startupTimeoutMutex.Unlock()
	fake.StartupTimeoutStub = nil
	if fake.startupTimeoutReturnsOnCall == nil {
		fake.startupTimeoutReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.startupTimeoutReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) Target() string {
	fake.targetMutex.Lock()
	ret, specificReturn := fake.targetReturnsOnCall[len(fake.targetArgsForCall)]
	fake.targetArgsForCall = append(fake.targetArgsForCall, struct {
	}{})
	fake.recordInvocation("Target", []interface{}{})
	fake.targetMutex.Unlock()
	if fake.TargetStub != nil {
		return fake.TargetStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.targetReturns
	return fakeReturns.result1
}

func (fake *FakeConfig) TargetCallCount() int {
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	return len(fake.targetArgsForCall)
}

func (fake *FakeConfig) TargetCalls(stub func() string) {
	fake.targetMutex.Lock()
	defer fake.targetMutex.Unlock()
	fake.TargetStub = stub
}

func (fake *FakeConfig) TargetReturns(result1 string) {
	fake.targetMutex.Lock()
	defer fake.targetMutex.Unlock()
	fake.TargetStub = nil
	fake.targetReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) TargetReturnsOnCall(i int, result1 string) {
	fake.targetMutex.Lock()
	defer fake.targetMutex.Unlock()
	fake.TargetStub = nil
	if fake.targetReturnsOnCall == nil {
		fake.targetReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.targetReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) UAAGrantType() string {
	fake.uAAGrantTypeMutex.Lock()
	ret, specificReturn := fake.uAAGrantTypeReturnsOnCall[len(fake.uAAGrantTypeArgsForCall)]
	fake.uAAGrantTypeArgsForCall = append(fake.uAAGrantTypeArgsForCall, struct {
	}{})
	fake.recordInvocation("UAAGrantType", []interface{}{})
	fake.uAAGrantTypeMutex.Unlock()
	if fake.UAAGrantTypeStub != nil {
		return fake.UAAGrantTypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.uAAGrantTypeReturns
	return fakeReturns.result1
}

func (fake *FakeConfig) UAAGrantTypeCallCount() int {
	fake.uAAGrantTypeMutex.RLock()
	defer fake.uAAGrantTypeMutex.RUnlock()
	return len(fake.uAAGrantTypeArgsForCall)
}

func (fake *FakeConfig) UAAGrantTypeCalls(stub func() string) {
	fake.uAAGrantTypeMutex.Lock()
	defer fake.uAAGrantTypeMutex.Unlock()
	fake.UAAGrantTypeStub = stub
}

func (fake *FakeConfig) UAAGrantTypeReturns(result1 string) {
	fake.uAAGrantTypeMutex.Lock()
	defer fake.uAAGrantTypeMutex.Unlock()
	fake.UAAGrantTypeStub = nil
	fake.uAAGrantTypeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) UAAGrantTypeReturnsOnCall(i int, result1 string) {
	fake.uAAGrantTypeMutex.Lock()
	defer fake.uAAGrantTypeMutex.Unlock()
	fake.UAAGrantTypeStub = nil
	if fake.uAAGrantTypeReturnsOnCall == nil {
		fake.uAAGrantTypeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.uAAGrantTypeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) UnsetOrganizationAndSpaceInformation() {
	fake.unsetOrganizationAndSpaceInformationMutex.Lock()
	fake.unsetOrganizationAndSpaceInformationArgsForCall = append(fake.unsetOrganizationAndSpaceInformationArgsForCall, struct {
	}{})
	fake.recordInvocation("UnsetOrganizationAndSpaceInformation", []interface{}{})
	fake.unsetOrganizationAndSpaceInformationMutex.Unlock()
	if fake.UnsetOrganizationAndSpaceInformationStub != nil {
		fake.UnsetOrganizationAndSpaceInformationStub()
	}
}

func (fake *FakeConfig) UnsetOrganizationAndSpaceInformationCallCount() int {
	fake.unsetOrganizationAndSpaceInformationMutex.RLock()
	defer fake.unsetOrganizationAndSpaceInformationMutex.RUnlock()
	return len(fake.unsetOrganizationAndSpaceInformationArgsForCall)
}

func (fake *FakeConfig) UnsetOrganizationAndSpaceInformationCalls(stub func()) {
	fake.unsetOrganizationAndSpaceInformationMutex.Lock()
	defer fake.unsetOrganizationAndSpaceInformationMutex.Unlock()
	fake.UnsetOrganizationAndSpaceInformationStub = stub
}

func (fake *FakeConfig) UnsetSpaceInformation() {
	fake.unsetSpaceInformationMutex.Lock()
	fake.unsetSpaceInformationArgsForCall = append(fake.unsetSpaceInformationArgsForCall, struct {
	}{})
	fake.recordInvocation("UnsetSpaceInformation", []interface{}{})
	fake.unsetSpaceInformationMutex.Unlock()
	if fake.UnsetSpaceInformationStub != nil {
		fake.UnsetSpaceInformationStub()
	}
}

func (fake *FakeConfig) UnsetSpaceInformationCallCount() int {
	fake.unsetSpaceInformationMutex.RLock()
	defer fake.unsetSpaceInformationMutex.RUnlock()
	return len(fake.unsetSpaceInformationArgsForCall)
}

func (fake *FakeConfig) UnsetSpaceInformationCalls(stub func()) {
	fake.unsetSpaceInformationMutex.Lock()
	defer fake.unsetSpaceInformationMutex.Unlock()
	fake.UnsetSpaceInformationStub = stub
}

func (fake *FakeConfig) Verbose() (bool, []string) {
	fake.verboseMutex.Lock()
	ret, specificReturn := fake.verboseReturnsOnCall[len(fake.verboseArgsForCall)]
	fake.verboseArgsForCall = append(fake.verboseArgsForCall, struct {
	}{})
	fake.recordInvocation("Verbose", []interface{}{})
	fake.verboseMutex.Unlock()
	if fake.VerboseStub != nil {
		return fake.VerboseStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.verboseReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConfig) VerboseCallCount() int {
	fake.verboseMutex.RLock()
	defer fake.verboseMutex.RUnlock()
	return len(fake.verboseArgsForCall)
}

func (fake *FakeConfig) VerboseCalls(stub func() (bool, []string)) {
	fake.verboseMutex.Lock()
	defer fake.verboseMutex.Unlock()
	fake.VerboseStub = stub
}

func (fake *FakeConfig) VerboseReturns(result1 bool, result2 []string) {
	fake.verboseMutex.Lock()
	defer fake.verboseMutex.Unlock()
	fake.VerboseStub = nil
	fake.verboseReturns = struct {
		result1 bool
		result2 []string
	}{result1, result2}
}

func (fake *FakeConfig) VerboseReturnsOnCall(i int, result1 bool, result2 []string) {
	fake.verboseMutex.Lock()
	defer fake.verboseMutex.Unlock()
	fake.VerboseStub = nil
	if fake.verboseReturnsOnCall == nil {
		fake.verboseReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 []string
		})
	}
	fake.verboseReturnsOnCall[i] = struct {
		result1 bool
		result2 []string
	}{result1, result2}
}

func (fake *FakeConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	fake.binaryNameMutex.RLock()
	defer fake.binaryNameMutex.RUnlock()
	fake.dialTimeoutMutex.RLock()
	defer fake.dialTimeoutMutex.RUnlock()
	fake.pollingIntervalMutex.RLock()
	defer fake.pollingIntervalMutex.RUnlock()
	fake.refreshTokenMutex.RLock()
	defer fake.refreshTokenMutex.RUnlock()
	fake.sSHOAuthClientMutex.RLock()
	defer fake.sSHOAuthClientMutex.RUnlock()
	fake.setAccessTokenMutex.RLock()
	defer fake.setAccessTokenMutex.RUnlock()
	fake.setRefreshTokenMutex.RLock()
	defer fake.setRefreshTokenMutex.RUnlock()
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	fake.setTokenInformationMutex.RLock()
	defer fake.setTokenInformationMutex.RUnlock()
	fake.setUAAClientCredentialsMutex.RLock()
	defer fake.setUAAClientCredentialsMutex.RUnlock()
	fake.setUAAGrantTypeMutex.RLock()
	defer fake.setUAAGrantTypeMutex.RUnlock()
	fake.skipSSLValidationMutex.RLock()
	defer fake.skipSSLValidationMutex.RUnlock()
	fake.stagingTimeoutMutex.RLock()
	defer fake.stagingTimeoutMutex.RUnlock()
	fake.startupTimeoutMutex.RLock()
	defer fake.startupTimeoutMutex.RUnlock()
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	fake.uAAGrantTypeMutex.RLock()
	defer fake.uAAGrantTypeMutex.RUnlock()
	fake.unsetOrganizationAndSpaceInformationMutex.RLock()
	defer fake.unsetOrganizationAndSpaceInformationMutex.RUnlock()
	fake.unsetSpaceInformationMutex.RLock()
	defer fake.unsetSpaceInformationMutex.RUnlock()
	fake.verboseMutex.RLock()
	defer fake.verboseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConfig) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v2action.Config = new(FakeConfig)
