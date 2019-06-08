// Code generated by counterfeiter. DO NOT EDIT.
package cffakes

import (
	"sync"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/oahcran/audit-services-cli-plugin/cf"
)

type FakeCfClient struct {
	GetServiceInstancesByUrlStub        func(plugin.CliConnection, string) ([]cf.ServiceInstance, error)
	getServiceInstancesByUrlMutex       sync.RWMutex
	getServiceInstancesByUrlArgsForCall []struct {
		arg1 plugin.CliConnection
		arg2 string
	}
	getServiceInstancesByUrlReturns struct {
		result1 []cf.ServiceInstance
		result2 error
	}
	getServiceInstancesByUrlReturnsOnCall map[int]struct {
		result1 []cf.ServiceInstance
		result2 error
	}
	GetServiceSummaryBySpaceStub        func(plugin.CliConnection, string) ([]cf.ServiceSummaryBySpace, error)
	getServiceSummaryBySpaceMutex       sync.RWMutex
	getServiceSummaryBySpaceArgsForCall []struct {
		arg1 plugin.CliConnection
		arg2 string
	}
	getServiceSummaryBySpaceReturns struct {
		result1 []cf.ServiceSummaryBySpace
		result2 error
	}
	getServiceSummaryBySpaceReturnsOnCall map[int]struct {
		result1 []cf.ServiceSummaryBySpace
		result2 error
	}
	ListOrgsStub        func(plugin.CliConnection) ([]cf.Org, error)
	listOrgsMutex       sync.RWMutex
	listOrgsArgsForCall []struct {
		arg1 plugin.CliConnection
	}
	listOrgsReturns struct {
		result1 []cf.Org
		result2 error
	}
	listOrgsReturnsOnCall map[int]struct {
		result1 []cf.Org
		result2 error
	}
	ListServicePlansStub        func(plugin.CliConnection) ([]cf.ServicePlan, error)
	listServicePlansMutex       sync.RWMutex
	listServicePlansArgsForCall []struct {
		arg1 plugin.CliConnection
	}
	listServicePlansReturns struct {
		result1 []cf.ServicePlan
		result2 error
	}
	listServicePlansReturnsOnCall map[int]struct {
		result1 []cf.ServicePlan
		result2 error
	}
	ListServicesStub        func(plugin.CliConnection) ([]cf.Service, error)
	listServicesMutex       sync.RWMutex
	listServicesArgsForCall []struct {
		arg1 plugin.CliConnection
	}
	listServicesReturns struct {
		result1 []cf.Service
		result2 error
	}
	listServicesReturnsOnCall map[int]struct {
		result1 []cf.Service
		result2 error
	}
	ListSpacesByOrgStub        func(plugin.CliConnection, string) ([]cf.Space, error)
	listSpacesByOrgMutex       sync.RWMutex
	listSpacesByOrgArgsForCall []struct {
		arg1 plugin.CliConnection
		arg2 string
	}
	listSpacesByOrgReturns struct {
		result1 []cf.Space
		result2 error
	}
	listSpacesByOrgReturnsOnCall map[int]struct {
		result1 []cf.Space
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCfClient) GetServiceInstancesByUrl(arg1 plugin.CliConnection, arg2 string) ([]cf.ServiceInstance, error) {
	fake.getServiceInstancesByUrlMutex.Lock()
	ret, specificReturn := fake.getServiceInstancesByUrlReturnsOnCall[len(fake.getServiceInstancesByUrlArgsForCall)]
	fake.getServiceInstancesByUrlArgsForCall = append(fake.getServiceInstancesByUrlArgsForCall, struct {
		arg1 plugin.CliConnection
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetServiceInstancesByUrl", []interface{}{arg1, arg2})
	fake.getServiceInstancesByUrlMutex.Unlock()
	if fake.GetServiceInstancesByUrlStub != nil {
		return fake.GetServiceInstancesByUrlStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getServiceInstancesByUrlReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCfClient) GetServiceInstancesByUrlCallCount() int {
	fake.getServiceInstancesByUrlMutex.RLock()
	defer fake.getServiceInstancesByUrlMutex.RUnlock()
	return len(fake.getServiceInstancesByUrlArgsForCall)
}

func (fake *FakeCfClient) GetServiceInstancesByUrlCalls(stub func(plugin.CliConnection, string) ([]cf.ServiceInstance, error)) {
	fake.getServiceInstancesByUrlMutex.Lock()
	defer fake.getServiceInstancesByUrlMutex.Unlock()
	fake.GetServiceInstancesByUrlStub = stub
}

func (fake *FakeCfClient) GetServiceInstancesByUrlArgsForCall(i int) (plugin.CliConnection, string) {
	fake.getServiceInstancesByUrlMutex.RLock()
	defer fake.getServiceInstancesByUrlMutex.RUnlock()
	argsForCall := fake.getServiceInstancesByUrlArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCfClient) GetServiceInstancesByUrlReturns(result1 []cf.ServiceInstance, result2 error) {
	fake.getServiceInstancesByUrlMutex.Lock()
	defer fake.getServiceInstancesByUrlMutex.Unlock()
	fake.GetServiceInstancesByUrlStub = nil
	fake.getServiceInstancesByUrlReturns = struct {
		result1 []cf.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeCfClient) GetServiceInstancesByUrlReturnsOnCall(i int, result1 []cf.ServiceInstance, result2 error) {
	fake.getServiceInstancesByUrlMutex.Lock()
	defer fake.getServiceInstancesByUrlMutex.Unlock()
	fake.GetServiceInstancesByUrlStub = nil
	if fake.getServiceInstancesByUrlReturnsOnCall == nil {
		fake.getServiceInstancesByUrlReturnsOnCall = make(map[int]struct {
			result1 []cf.ServiceInstance
			result2 error
		})
	}
	fake.getServiceInstancesByUrlReturnsOnCall[i] = struct {
		result1 []cf.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeCfClient) GetServiceSummaryBySpace(arg1 plugin.CliConnection, arg2 string) ([]cf.ServiceSummaryBySpace, error) {
	fake.getServiceSummaryBySpaceMutex.Lock()
	ret, specificReturn := fake.getServiceSummaryBySpaceReturnsOnCall[len(fake.getServiceSummaryBySpaceArgsForCall)]
	fake.getServiceSummaryBySpaceArgsForCall = append(fake.getServiceSummaryBySpaceArgsForCall, struct {
		arg1 plugin.CliConnection
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetServiceSummaryBySpace", []interface{}{arg1, arg2})
	fake.getServiceSummaryBySpaceMutex.Unlock()
	if fake.GetServiceSummaryBySpaceStub != nil {
		return fake.GetServiceSummaryBySpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getServiceSummaryBySpaceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCfClient) GetServiceSummaryBySpaceCallCount() int {
	fake.getServiceSummaryBySpaceMutex.RLock()
	defer fake.getServiceSummaryBySpaceMutex.RUnlock()
	return len(fake.getServiceSummaryBySpaceArgsForCall)
}

func (fake *FakeCfClient) GetServiceSummaryBySpaceCalls(stub func(plugin.CliConnection, string) ([]cf.ServiceSummaryBySpace, error)) {
	fake.getServiceSummaryBySpaceMutex.Lock()
	defer fake.getServiceSummaryBySpaceMutex.Unlock()
	fake.GetServiceSummaryBySpaceStub = stub
}

func (fake *FakeCfClient) GetServiceSummaryBySpaceArgsForCall(i int) (plugin.CliConnection, string) {
	fake.getServiceSummaryBySpaceMutex.RLock()
	defer fake.getServiceSummaryBySpaceMutex.RUnlock()
	argsForCall := fake.getServiceSummaryBySpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCfClient) GetServiceSummaryBySpaceReturns(result1 []cf.ServiceSummaryBySpace, result2 error) {
	fake.getServiceSummaryBySpaceMutex.Lock()
	defer fake.getServiceSummaryBySpaceMutex.Unlock()
	fake.GetServiceSummaryBySpaceStub = nil
	fake.getServiceSummaryBySpaceReturns = struct {
		result1 []cf.ServiceSummaryBySpace
		result2 error
	}{result1, result2}
}

func (fake *FakeCfClient) GetServiceSummaryBySpaceReturnsOnCall(i int, result1 []cf.ServiceSummaryBySpace, result2 error) {
	fake.getServiceSummaryBySpaceMutex.Lock()
	defer fake.getServiceSummaryBySpaceMutex.Unlock()
	fake.GetServiceSummaryBySpaceStub = nil
	if fake.getServiceSummaryBySpaceReturnsOnCall == nil {
		fake.getServiceSummaryBySpaceReturnsOnCall = make(map[int]struct {
			result1 []cf.ServiceSummaryBySpace
			result2 error
		})
	}
	fake.getServiceSummaryBySpaceReturnsOnCall[i] = struct {
		result1 []cf.ServiceSummaryBySpace
		result2 error
	}{result1, result2}
}

func (fake *FakeCfClient) ListOrgs(arg1 plugin.CliConnection) ([]cf.Org, error) {
	fake.listOrgsMutex.Lock()
	ret, specificReturn := fake.listOrgsReturnsOnCall[len(fake.listOrgsArgsForCall)]
	fake.listOrgsArgsForCall = append(fake.listOrgsArgsForCall, struct {
		arg1 plugin.CliConnection
	}{arg1})
	fake.recordInvocation("ListOrgs", []interface{}{arg1})
	fake.listOrgsMutex.Unlock()
	if fake.ListOrgsStub != nil {
		return fake.ListOrgsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listOrgsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCfClient) ListOrgsCallCount() int {
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	return len(fake.listOrgsArgsForCall)
}

func (fake *FakeCfClient) ListOrgsCalls(stub func(plugin.CliConnection) ([]cf.Org, error)) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = stub
}

func (fake *FakeCfClient) ListOrgsArgsForCall(i int) plugin.CliConnection {
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	argsForCall := fake.listOrgsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCfClient) ListOrgsReturns(result1 []cf.Org, result2 error) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = nil
	fake.listOrgsReturns = struct {
		result1 []cf.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeCfClient) ListOrgsReturnsOnCall(i int, result1 []cf.Org, result2 error) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = nil
	if fake.listOrgsReturnsOnCall == nil {
		fake.listOrgsReturnsOnCall = make(map[int]struct {
			result1 []cf.Org
			result2 error
		})
	}
	fake.listOrgsReturnsOnCall[i] = struct {
		result1 []cf.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeCfClient) ListServicePlans(arg1 plugin.CliConnection) ([]cf.ServicePlan, error) {
	fake.listServicePlansMutex.Lock()
	ret, specificReturn := fake.listServicePlansReturnsOnCall[len(fake.listServicePlansArgsForCall)]
	fake.listServicePlansArgsForCall = append(fake.listServicePlansArgsForCall, struct {
		arg1 plugin.CliConnection
	}{arg1})
	fake.recordInvocation("ListServicePlans", []interface{}{arg1})
	fake.listServicePlansMutex.Unlock()
	if fake.ListServicePlansStub != nil {
		return fake.ListServicePlansStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listServicePlansReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCfClient) ListServicePlansCallCount() int {
	fake.listServicePlansMutex.RLock()
	defer fake.listServicePlansMutex.RUnlock()
	return len(fake.listServicePlansArgsForCall)
}

func (fake *FakeCfClient) ListServicePlansCalls(stub func(plugin.CliConnection) ([]cf.ServicePlan, error)) {
	fake.listServicePlansMutex.Lock()
	defer fake.listServicePlansMutex.Unlock()
	fake.ListServicePlansStub = stub
}

func (fake *FakeCfClient) ListServicePlansArgsForCall(i int) plugin.CliConnection {
	fake.listServicePlansMutex.RLock()
	defer fake.listServicePlansMutex.RUnlock()
	argsForCall := fake.listServicePlansArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCfClient) ListServicePlansReturns(result1 []cf.ServicePlan, result2 error) {
	fake.listServicePlansMutex.Lock()
	defer fake.listServicePlansMutex.Unlock()
	fake.ListServicePlansStub = nil
	fake.listServicePlansReturns = struct {
		result1 []cf.ServicePlan
		result2 error
	}{result1, result2}
}

func (fake *FakeCfClient) ListServicePlansReturnsOnCall(i int, result1 []cf.ServicePlan, result2 error) {
	fake.listServicePlansMutex.Lock()
	defer fake.listServicePlansMutex.Unlock()
	fake.ListServicePlansStub = nil
	if fake.listServicePlansReturnsOnCall == nil {
		fake.listServicePlansReturnsOnCall = make(map[int]struct {
			result1 []cf.ServicePlan
			result2 error
		})
	}
	fake.listServicePlansReturnsOnCall[i] = struct {
		result1 []cf.ServicePlan
		result2 error
	}{result1, result2}
}

func (fake *FakeCfClient) ListServices(arg1 plugin.CliConnection) ([]cf.Service, error) {
	fake.listServicesMutex.Lock()
	ret, specificReturn := fake.listServicesReturnsOnCall[len(fake.listServicesArgsForCall)]
	fake.listServicesArgsForCall = append(fake.listServicesArgsForCall, struct {
		arg1 plugin.CliConnection
	}{arg1})
	fake.recordInvocation("ListServices", []interface{}{arg1})
	fake.listServicesMutex.Unlock()
	if fake.ListServicesStub != nil {
		return fake.ListServicesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listServicesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCfClient) ListServicesCallCount() int {
	fake.listServicesMutex.RLock()
	defer fake.listServicesMutex.RUnlock()
	return len(fake.listServicesArgsForCall)
}

func (fake *FakeCfClient) ListServicesCalls(stub func(plugin.CliConnection) ([]cf.Service, error)) {
	fake.listServicesMutex.Lock()
	defer fake.listServicesMutex.Unlock()
	fake.ListServicesStub = stub
}

func (fake *FakeCfClient) ListServicesArgsForCall(i int) plugin.CliConnection {
	fake.listServicesMutex.RLock()
	defer fake.listServicesMutex.RUnlock()
	argsForCall := fake.listServicesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCfClient) ListServicesReturns(result1 []cf.Service, result2 error) {
	fake.listServicesMutex.Lock()
	defer fake.listServicesMutex.Unlock()
	fake.ListServicesStub = nil
	fake.listServicesReturns = struct {
		result1 []cf.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeCfClient) ListServicesReturnsOnCall(i int, result1 []cf.Service, result2 error) {
	fake.listServicesMutex.Lock()
	defer fake.listServicesMutex.Unlock()
	fake.ListServicesStub = nil
	if fake.listServicesReturnsOnCall == nil {
		fake.listServicesReturnsOnCall = make(map[int]struct {
			result1 []cf.Service
			result2 error
		})
	}
	fake.listServicesReturnsOnCall[i] = struct {
		result1 []cf.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeCfClient) ListSpacesByOrg(arg1 plugin.CliConnection, arg2 string) ([]cf.Space, error) {
	fake.listSpacesByOrgMutex.Lock()
	ret, specificReturn := fake.listSpacesByOrgReturnsOnCall[len(fake.listSpacesByOrgArgsForCall)]
	fake.listSpacesByOrgArgsForCall = append(fake.listSpacesByOrgArgsForCall, struct {
		arg1 plugin.CliConnection
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ListSpacesByOrg", []interface{}{arg1, arg2})
	fake.listSpacesByOrgMutex.Unlock()
	if fake.ListSpacesByOrgStub != nil {
		return fake.ListSpacesByOrgStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listSpacesByOrgReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCfClient) ListSpacesByOrgCallCount() int {
	fake.listSpacesByOrgMutex.RLock()
	defer fake.listSpacesByOrgMutex.RUnlock()
	return len(fake.listSpacesByOrgArgsForCall)
}

func (fake *FakeCfClient) ListSpacesByOrgCalls(stub func(plugin.CliConnection, string) ([]cf.Space, error)) {
	fake.listSpacesByOrgMutex.Lock()
	defer fake.listSpacesByOrgMutex.Unlock()
	fake.ListSpacesByOrgStub = stub
}

func (fake *FakeCfClient) ListSpacesByOrgArgsForCall(i int) (plugin.CliConnection, string) {
	fake.listSpacesByOrgMutex.RLock()
	defer fake.listSpacesByOrgMutex.RUnlock()
	argsForCall := fake.listSpacesByOrgArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCfClient) ListSpacesByOrgReturns(result1 []cf.Space, result2 error) {
	fake.listSpacesByOrgMutex.Lock()
	defer fake.listSpacesByOrgMutex.Unlock()
	fake.ListSpacesByOrgStub = nil
	fake.listSpacesByOrgReturns = struct {
		result1 []cf.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCfClient) ListSpacesByOrgReturnsOnCall(i int, result1 []cf.Space, result2 error) {
	fake.listSpacesByOrgMutex.Lock()
	defer fake.listSpacesByOrgMutex.Unlock()
	fake.ListSpacesByOrgStub = nil
	if fake.listSpacesByOrgReturnsOnCall == nil {
		fake.listSpacesByOrgReturnsOnCall = make(map[int]struct {
			result1 []cf.Space
			result2 error
		})
	}
	fake.listSpacesByOrgReturnsOnCall[i] = struct {
		result1 []cf.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCfClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getServiceInstancesByUrlMutex.RLock()
	defer fake.getServiceInstancesByUrlMutex.RUnlock()
	fake.getServiceSummaryBySpaceMutex.RLock()
	defer fake.getServiceSummaryBySpaceMutex.RUnlock()
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	fake.listServicePlansMutex.RLock()
	defer fake.listServicePlansMutex.RUnlock()
	fake.listServicesMutex.RLock()
	defer fake.listServicesMutex.RUnlock()
	fake.listSpacesByOrgMutex.RLock()
	defer fake.listSpacesByOrgMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCfClient) recordInvocation(key string, args []interface{}) {
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

var _ cf.CfClient = new(FakeCfClient)
