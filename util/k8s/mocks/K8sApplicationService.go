// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	bean "github.com/devtron-labs/devtron/api/bean"
	application "github.com/devtron-labs/devtron/client/k8s/application"

	client "github.com/devtron-labs/devtron/api/helm-app"

	cluster "github.com/devtron-labs/devtron/pkg/cluster"

	context "context"

	io "io"

	k8s "github.com/devtron-labs/devtron/util/k8s"

	mock "github.com/stretchr/testify/mock"

	rest "k8s.io/client-go/rest"
)

// K8sApplicationService is an autogenerated mock type for the K8sApplicationService type
type K8sApplicationService struct {
	mock.Mock
}

// CreateResource provides a mock function with given fields: request
func (_m *K8sApplicationService) CreateResource(request *k8s.ResourceRequestBean) (*application.ManifestResponse, error) {
	ret := _m.Called(request)

	var r0 *application.ManifestResponse
	if rf, ok := ret.Get(0).(func(*k8s.ResourceRequestBean) *application.ManifestResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*application.ManifestResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*k8s.ResourceRequestBean) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteResource provides a mock function with given fields: request
func (_m *K8sApplicationService) DeleteResource(request *k8s.ResourceRequestBean) (*application.ManifestResponse, error) {
	ret := _m.Called(request)

	var r0 *application.ManifestResponse
	if rf, ok := ret.Get(0).(func(*k8s.ResourceRequestBean) *application.ManifestResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*application.ManifestResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*k8s.ResourceRequestBean) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterServiceAndIngress provides a mock function with given fields: resourceTreeInf, validRequests, appDetail, appId
func (_m *K8sApplicationService) FilterServiceAndIngress(resourceTreeInf map[string]interface{}, validRequests []k8s.ResourceRequestBean, appDetail bean.AppDetailContainer, appId string) []k8s.ResourceRequestBean {
	ret := _m.Called(resourceTreeInf, validRequests, appDetail, appId)

	var r0 []k8s.ResourceRequestBean
	if rf, ok := ret.Get(0).(func(map[string]interface{}, []k8s.ResourceRequestBean, bean.AppDetailContainer, string) []k8s.ResourceRequestBean); ok {
		r0 = rf(resourceTreeInf, validRequests, appDetail, appId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]k8s.ResourceRequestBean)
		}
	}

	return r0
}

// GetManifestsByBatch provides a mock function with given fields: ctx, request
func (_m *K8sApplicationService) GetManifestsByBatch(ctx context.Context, request []k8s.ResourceRequestBean) ([]k8s.BatchResourceResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 []k8s.BatchResourceResponse
	if rf, ok := ret.Get(0).(func(context.Context, []k8s.ResourceRequestBean) []k8s.BatchResourceResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]k8s.BatchResourceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []k8s.ResourceRequestBean) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPodLogs provides a mock function with given fields: request
func (_m *K8sApplicationService) GetPodLogs(request *k8s.ResourceRequestBean) (io.ReadCloser, error) {
	ret := _m.Called(request)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(*k8s.ResourceRequestBean) io.ReadCloser); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*k8s.ResourceRequestBean) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResource provides a mock function with given fields: request
func (_m *K8sApplicationService) GetResource(request *k8s.ResourceRequestBean) (*application.ManifestResponse, error) {
	ret := _m.Called(request)

	var r0 *application.ManifestResponse
	if rf, ok := ret.Get(0).(func(*k8s.ResourceRequestBean) *application.ManifestResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*application.ManifestResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*k8s.ResourceRequestBean) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourceInfo provides a mock function with given fields:
func (_m *K8sApplicationService) GetResourceInfo() (*k8s.ResourceInfo, error) {
	ret := _m.Called()

	var r0 *k8s.ResourceInfo
	if rf, ok := ret.Get(0).(func() *k8s.ResourceInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*k8s.ResourceInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRestConfigByCluster provides a mock function with given fields: _a0
func (_m *K8sApplicationService) GetRestConfigByCluster(_a0 *cluster.ClusterBean) (*rest.Config, error) {
	ret := _m.Called(_a0)

	var r0 *rest.Config
	if rf, ok := ret.Get(0).(func(*cluster.ClusterBean) *rest.Config); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Config)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cluster.ClusterBean) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRestConfigByClusterId provides a mock function with given fields: clusterId
func (_m *K8sApplicationService) GetRestConfigByClusterId(clusterId int) (*rest.Config, error) {
	ret := _m.Called(clusterId)

	var r0 *rest.Config
	if rf, ok := ret.Get(0).(func(int) *rest.Config); ok {
		r0 = rf(clusterId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Config)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(clusterId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUrlsByBatch provides a mock function with given fields: resp
func (_m *K8sApplicationService) GetUrlsByBatch(resp []k8s.BatchResourceResponse) []interface{} {
	ret := _m.Called(resp)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func([]k8s.BatchResourceResponse) []interface{}); ok {
		r0 = rf(resp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	return r0
}

// ListEvents provides a mock function with given fields: request
func (_m *K8sApplicationService) ListEvents(request *k8s.ResourceRequestBean) (*application.EventsResponse, error) {
	ret := _m.Called(request)

	var r0 *application.EventsResponse
	if rf, ok := ret.Get(0).(func(*k8s.ResourceRequestBean) *application.EventsResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*application.EventsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*k8s.ResourceRequestBean) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateResource provides a mock function with given fields: request
func (_m *K8sApplicationService) UpdateResource(request *k8s.ResourceRequestBean) (*application.ManifestResponse, error) {
	ret := _m.Called(request)

	var r0 *application.ManifestResponse
	if rf, ok := ret.Get(0).(func(*k8s.ResourceRequestBean) *application.ManifestResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*application.ManifestResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*k8s.ResourceRequestBean) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateResourceRequest provides a mock function with given fields: appIdentifier, request
func (_m *K8sApplicationService) ValidateResourceRequest(appIdentifier *client.AppIdentifier, request *application.K8sRequestBean) (bool, error) {
	ret := _m.Called(appIdentifier, request)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*client.AppIdentifier, *application.K8sRequestBean) bool); ok {
		r0 = rf(appIdentifier, request)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.AppIdentifier, *application.K8sRequestBean) error); ok {
		r1 = rf(appIdentifier, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewK8sApplicationService interface {
	mock.TestingT
	Cleanup(func())
}

// NewK8sApplicationService creates a new instance of K8sApplicationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewK8sApplicationService(t mockConstructorTestingTNewK8sApplicationService) *K8sApplicationService {
	mock := &K8sApplicationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
