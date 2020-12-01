// Code generated by mockery v1.0.0. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
)

// RolesGetter is an autogenerated mock type for the RolesGetter type
type RolesGetter struct {
	mock.Mock
}

// Roles provides a mock function with given fields: namespace
func (_m *RolesGetter) Roles(namespace string) v1alpha1.RoleInterface {
	ret := _m.Called(namespace)

	var r0 v1alpha1.RoleInterface
	if rf, ok := ret.Get(0).(func(string) v1alpha1.RoleInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.RoleInterface)
		}
	}

	return r0
}