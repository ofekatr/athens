// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	minio "github.com/minio/minio-go/v6"
	mock "github.com/stretchr/testify/mock"
)

// MinioCore is an autogenerated mock type for the minioCore type
type MinioCore struct {
	mock.Mock
}

// ListObjectsV2 provides a mock function with given fields: bucketName, objectPrefix, continuationToken, fetchOwner, delimiter, maxkeys, startAfter
func (_m *MinioCore) ListObjectsV2(bucketName string, objectPrefix string, continuationToken string, fetchOwner bool, delimiter string, maxkeys int, startAfter string) (minio.ListBucketV2Result, error) {
	ret := _m.Called(bucketName, objectPrefix, continuationToken, fetchOwner, delimiter, maxkeys, startAfter) // nolint: typecheck // linter does not recognize struct embedding

	var r0 minio.ListBucketV2Result
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, bool, string, int, string) (minio.ListBucketV2Result, error)); ok {
		return rf(bucketName, objectPrefix, continuationToken, fetchOwner, delimiter, maxkeys, startAfter)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, bool, string, int, string) minio.ListBucketV2Result); ok {
		r0 = rf(bucketName, objectPrefix, continuationToken, fetchOwner, delimiter, maxkeys, startAfter)
	} else {
		r0 = ret.Get(0).(minio.ListBucketV2Result)
	}

	if rf, ok := ret.Get(1).(func(string, string, string, bool, string, int, string) error); ok {
		r1 = rf(bucketName, objectPrefix, continuationToken, fetchOwner, delimiter, maxkeys, startAfter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMinioCore interface {
	mock.TestingT
	Cleanup(func())
}

// NewMinioCore creates a new instance of MinioCore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMinioCore(t mockConstructorTestingTNewMinioCore) *MinioCore {
	mock := &MinioCore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) }) // nolint: typecheck // linter does not recognize struct embedding

	return mock
}