// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "github.com/compose-spec/compose-go/types"
	api "github.com/docker/compose/v2/pkg/api"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockService) Build(ctx context.Context, project *types.Project, options api.BuildOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", ctx, project, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockServiceMockRecorder) Build(ctx, project, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockService)(nil).Build), ctx, project, options)
}

// Convert mocks base method.
func (m *MockService) Convert(ctx context.Context, project *types.Project, options api.ConvertOptions) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Convert", ctx, project, options)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Convert indicates an expected call of Convert.
func (mr *MockServiceMockRecorder) Convert(ctx, project, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Convert", reflect.TypeOf((*MockService)(nil).Convert), ctx, project, options)
}

// Copy mocks base method.
func (m *MockService) Copy(ctx context.Context, projectName string, options api.CopyOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy", ctx, projectName, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Copy indicates an expected call of Copy.
func (mr *MockServiceMockRecorder) Copy(ctx, projectName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockService)(nil).Copy), ctx, projectName, options)
}

// Create mocks base method.
func (m *MockService) Create(ctx context.Context, project *types.Project, options api.CreateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, project, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockServiceMockRecorder) Create(ctx, project, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), ctx, project, options)
}

// Down mocks base method.
func (m *MockService) Down(ctx context.Context, projectName string, options api.DownOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Down", ctx, projectName, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Down indicates an expected call of Down.
func (mr *MockServiceMockRecorder) Down(ctx, projectName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Down", reflect.TypeOf((*MockService)(nil).Down), ctx, projectName, options)
}

// Events mocks base method.
func (m *MockService) Events(ctx context.Context, projectName string, options api.EventsOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Events", ctx, projectName, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Events indicates an expected call of Events.
func (mr *MockServiceMockRecorder) Events(ctx, projectName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Events", reflect.TypeOf((*MockService)(nil).Events), ctx, projectName, options)
}

// Exec mocks base method.
func (m *MockService) Exec(ctx context.Context, projectName string, options api.RunOptions) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", ctx, projectName, options)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockServiceMockRecorder) Exec(ctx, projectName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockService)(nil).Exec), ctx, projectName, options)
}

// Images mocks base method.
func (m *MockService) Images(ctx context.Context, projectName string, options api.ImagesOptions) ([]api.ImageSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Images", ctx, projectName, options)
	ret0, _ := ret[0].([]api.ImageSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Images indicates an expected call of Images.
func (mr *MockServiceMockRecorder) Images(ctx, projectName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Images", reflect.TypeOf((*MockService)(nil).Images), ctx, projectName, options)
}

// Kill mocks base method.
func (m *MockService) Kill(ctx context.Context, projectName string, options api.KillOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kill", ctx, projectName, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Kill indicates an expected call of Kill.
func (mr *MockServiceMockRecorder) Kill(ctx, projectName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockService)(nil).Kill), ctx, projectName, options)
}

// List mocks base method.
func (m *MockService) List(ctx context.Context, options api.ListOptions) ([]api.Stack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, options)
	ret0, _ := ret[0].([]api.Stack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockServiceMockRecorder) List(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockService)(nil).List), ctx, options)
}

// Logs mocks base method.
func (m *MockService) Logs(ctx context.Context, projectName string, consumer api.LogConsumer, options api.LogOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logs", ctx, projectName, consumer, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logs indicates an expected call of Logs.
func (mr *MockServiceMockRecorder) Logs(ctx, projectName, consumer, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logs", reflect.TypeOf((*MockService)(nil).Logs), ctx, projectName, consumer, options)
}

// Pause mocks base method.
func (m *MockService) Pause(ctx context.Context, projectName string, options api.PauseOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pause", ctx, projectName, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pause indicates an expected call of Pause.
func (mr *MockServiceMockRecorder) Pause(ctx, projectName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pause", reflect.TypeOf((*MockService)(nil).Pause), ctx, projectName, options)
}

// Port mocks base method.
func (m *MockService) Port(ctx context.Context, projectName, service string, port int, options api.PortOptions) (string, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Port", ctx, projectName, service, port, options)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Port indicates an expected call of Port.
func (mr *MockServiceMockRecorder) Port(ctx, projectName, service, port, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Port", reflect.TypeOf((*MockService)(nil).Port), ctx, projectName, service, port, options)
}

// Ps mocks base method.
func (m *MockService) Ps(ctx context.Context, projectName string, options api.PsOptions) ([]api.ContainerSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ps", ctx, projectName, options)
	ret0, _ := ret[0].([]api.ContainerSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ps indicates an expected call of Ps.
func (mr *MockServiceMockRecorder) Ps(ctx, projectName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ps", reflect.TypeOf((*MockService)(nil).Ps), ctx, projectName, options)
}

// Pull mocks base method.
func (m *MockService) Pull(ctx context.Context, project *types.Project, options api.PullOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull", ctx, project, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pull indicates an expected call of Pull.
func (mr *MockServiceMockRecorder) Pull(ctx, project, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockService)(nil).Pull), ctx, project, options)
}

// Push mocks base method.
func (m *MockService) Push(ctx context.Context, project *types.Project, options api.PushOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", ctx, project, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockServiceMockRecorder) Push(ctx, project, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockService)(nil).Push), ctx, project, options)
}

// Remove mocks base method.
func (m *MockService) Remove(ctx context.Context, projectName string, options api.RemoveOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, projectName, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockServiceMockRecorder) Remove(ctx, projectName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockService)(nil).Remove), ctx, projectName, options)
}

// Restart mocks base method.
func (m *MockService) Restart(ctx context.Context, projectName string, options api.RestartOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restart", ctx, projectName, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restart indicates an expected call of Restart.
func (mr *MockServiceMockRecorder) Restart(ctx, projectName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restart", reflect.TypeOf((*MockService)(nil).Restart), ctx, projectName, options)
}

// RunOneOffContainer mocks base method.
func (m *MockService) RunOneOffContainer(ctx context.Context, project *types.Project, opts api.RunOptions) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunOneOffContainer", ctx, project, opts)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunOneOffContainer indicates an expected call of RunOneOffContainer.
func (mr *MockServiceMockRecorder) RunOneOffContainer(ctx, project, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunOneOffContainer", reflect.TypeOf((*MockService)(nil).RunOneOffContainer), ctx, project, opts)
}

// Start mocks base method.
func (m *MockService) Start(ctx context.Context, projectName string, options api.StartOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", ctx, projectName, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockServiceMockRecorder) Start(ctx, projectName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockService)(nil).Start), ctx, projectName, options)
}

// Stop mocks base method.
func (m *MockService) Stop(ctx context.Context, projectName string, options api.StopOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", ctx, projectName, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockServiceMockRecorder) Stop(ctx, projectName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockService)(nil).Stop), ctx, projectName, options)
}

// Top mocks base method.
func (m *MockService) Top(ctx context.Context, projectName string, services []string) ([]api.ContainerProcSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Top", ctx, projectName, services)
	ret0, _ := ret[0].([]api.ContainerProcSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Top indicates an expected call of Top.
func (mr *MockServiceMockRecorder) Top(ctx, projectName, services interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Top", reflect.TypeOf((*MockService)(nil).Top), ctx, projectName, services)
}

// UnPause mocks base method.
func (m *MockService) UnPause(ctx context.Context, projectName string, options api.PauseOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnPause", ctx, projectName, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnPause indicates an expected call of UnPause.
func (mr *MockServiceMockRecorder) UnPause(ctx, projectName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnPause", reflect.TypeOf((*MockService)(nil).UnPause), ctx, projectName, options)
}

// Up mocks base method.
func (m *MockService) Up(ctx context.Context, project *types.Project, options api.UpOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Up", ctx, project, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Up indicates an expected call of Up.
func (mr *MockServiceMockRecorder) Up(ctx, project, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Up", reflect.TypeOf((*MockService)(nil).Up), ctx, project, options)
}

// MockLogConsumer is a mock of LogConsumer interface.
type MockLogConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockLogConsumerMockRecorder
}

// MockLogConsumerMockRecorder is the mock recorder for MockLogConsumer.
type MockLogConsumerMockRecorder struct {
	mock *MockLogConsumer
}

// NewMockLogConsumer creates a new mock instance.
func NewMockLogConsumer(ctrl *gomock.Controller) *MockLogConsumer {
	mock := &MockLogConsumer{ctrl: ctrl}
	mock.recorder = &MockLogConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogConsumer) EXPECT() *MockLogConsumerMockRecorder {
	return m.recorder
}

// Log mocks base method.
func (m *MockLogConsumer) Log(containerName, service, message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Log", containerName, service, message)
}

// Log indicates an expected call of Log.
func (mr *MockLogConsumerMockRecorder) Log(containerName, service, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockLogConsumer)(nil).Log), containerName, service, message)
}

// Register mocks base method.
func (m *MockLogConsumer) Register(container string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", container)
}

// Register indicates an expected call of Register.
func (mr *MockLogConsumerMockRecorder) Register(container interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockLogConsumer)(nil).Register), container)
}

// Status mocks base method.
func (m *MockLogConsumer) Status(container, msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Status", container, msg)
}

// Status indicates an expected call of Status.
func (mr *MockLogConsumerMockRecorder) Status(container, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockLogConsumer)(nil).Status), container, msg)
}
