// Code generated by MockGen. DO NOT EDIT.
// Source: types/module/module.go

// Package mocks is a generated GoMock package.
package mocks

import (
	json "encoding/json"
	gomock "github.com/golang/mock/gomock"
	mux "github.com/gorilla/mux"
	runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	client "github.com/sisu-network/cosmos-sdk/client"
	codec "github.com/sisu-network/cosmos-sdk/codec"
	types "github.com/sisu-network/cosmos-sdk/codec/types"
	types0 "github.com/sisu-network/cosmos-sdk/types"
	module "github.com/sisu-network/cosmos-sdk/types/module"
	types1 "github.com/sisu-network/tendermint/abci/types"
	cobra "github.com/spf13/cobra"
	reflect "reflect"
)

// MockAppModuleBasic is a mock of AppModuleBasic interface
type MockAppModuleBasic struct {
	ctrl     *gomock.Controller
	recorder *MockAppModuleBasicMockRecorder
}

// MockAppModuleBasicMockRecorder is the mock recorder for MockAppModuleBasic
type MockAppModuleBasicMockRecorder struct {
	mock *MockAppModuleBasic
}

// NewMockAppModuleBasic creates a new mock instance
func NewMockAppModuleBasic(ctrl *gomock.Controller) *MockAppModuleBasic {
	mock := &MockAppModuleBasic{ctrl: ctrl}
	mock.recorder = &MockAppModuleBasicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppModuleBasic) EXPECT() *MockAppModuleBasicMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockAppModuleBasic) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockAppModuleBasicMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAppModuleBasic)(nil).Name))
}

// RegisterLegacyAminoCodec mocks base method
func (m *MockAppModuleBasic) RegisterLegacyAminoCodec(arg0 *codec.LegacyAmino) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterLegacyAminoCodec", arg0)
}

// RegisterLegacyAminoCodec indicates an expected call of RegisterLegacyAminoCodec
func (mr *MockAppModuleBasicMockRecorder) RegisterLegacyAminoCodec(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterLegacyAminoCodec", reflect.TypeOf((*MockAppModuleBasic)(nil).RegisterLegacyAminoCodec), arg0)
}

// RegisterInterfaces mocks base method
func (m *MockAppModuleBasic) RegisterInterfaces(arg0 types.InterfaceRegistry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInterfaces", arg0)
}

// RegisterInterfaces indicates an expected call of RegisterInterfaces
func (mr *MockAppModuleBasicMockRecorder) RegisterInterfaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInterfaces", reflect.TypeOf((*MockAppModuleBasic)(nil).RegisterInterfaces), arg0)
}

// DefaultGenesis mocks base method
func (m *MockAppModuleBasic) DefaultGenesis(arg0 codec.JSONMarshaler) json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultGenesis", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// DefaultGenesis indicates an expected call of DefaultGenesis
func (mr *MockAppModuleBasicMockRecorder) DefaultGenesis(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultGenesis", reflect.TypeOf((*MockAppModuleBasic)(nil).DefaultGenesis), arg0)
}

// ValidateGenesis mocks base method
func (m *MockAppModuleBasic) ValidateGenesis(arg0 codec.JSONMarshaler, arg1 client.TxEncodingConfig, arg2 json.RawMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateGenesis", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateGenesis indicates an expected call of ValidateGenesis
func (mr *MockAppModuleBasicMockRecorder) ValidateGenesis(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateGenesis", reflect.TypeOf((*MockAppModuleBasic)(nil).ValidateGenesis), arg0, arg1, arg2)
}

// RegisterRESTRoutes mocks base method
func (m *MockAppModuleBasic) RegisterRESTRoutes(arg0 client.Context, arg1 *mux.Router) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterRESTRoutes", arg0, arg1)
}

// RegisterRESTRoutes indicates an expected call of RegisterRESTRoutes
func (mr *MockAppModuleBasicMockRecorder) RegisterRESTRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRESTRoutes", reflect.TypeOf((*MockAppModuleBasic)(nil).RegisterRESTRoutes), arg0, arg1)
}

// RegisterGRPCGatewayRoutes mocks base method
func (m *MockAppModuleBasic) RegisterGRPCGatewayRoutes(arg0 client.Context, arg1 *runtime.ServeMux) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterGRPCGatewayRoutes", arg0, arg1)
}

// RegisterGRPCGatewayRoutes indicates an expected call of RegisterGRPCGatewayRoutes
func (mr *MockAppModuleBasicMockRecorder) RegisterGRPCGatewayRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCGatewayRoutes", reflect.TypeOf((*MockAppModuleBasic)(nil).RegisterGRPCGatewayRoutes), arg0, arg1)
}

// GetTxCmd mocks base method
func (m *MockAppModuleBasic) GetTxCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetTxCmd indicates an expected call of GetTxCmd
func (mr *MockAppModuleBasicMockRecorder) GetTxCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxCmd", reflect.TypeOf((*MockAppModuleBasic)(nil).GetTxCmd))
}

// GetQueryCmd mocks base method
func (m *MockAppModuleBasic) GetQueryCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetQueryCmd indicates an expected call of GetQueryCmd
func (mr *MockAppModuleBasicMockRecorder) GetQueryCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryCmd", reflect.TypeOf((*MockAppModuleBasic)(nil).GetQueryCmd))
}

// MockAppModuleGenesis is a mock of AppModuleGenesis interface
type MockAppModuleGenesis struct {
	ctrl     *gomock.Controller
	recorder *MockAppModuleGenesisMockRecorder
}

// MockAppModuleGenesisMockRecorder is the mock recorder for MockAppModuleGenesis
type MockAppModuleGenesisMockRecorder struct {
	mock *MockAppModuleGenesis
}

// NewMockAppModuleGenesis creates a new mock instance
func NewMockAppModuleGenesis(ctrl *gomock.Controller) *MockAppModuleGenesis {
	mock := &MockAppModuleGenesis{ctrl: ctrl}
	mock.recorder = &MockAppModuleGenesisMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppModuleGenesis) EXPECT() *MockAppModuleGenesisMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockAppModuleGenesis) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockAppModuleGenesisMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAppModuleGenesis)(nil).Name))
}

// RegisterLegacyAminoCodec mocks base method
func (m *MockAppModuleGenesis) RegisterLegacyAminoCodec(arg0 *codec.LegacyAmino) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterLegacyAminoCodec", arg0)
}

// RegisterLegacyAminoCodec indicates an expected call of RegisterLegacyAminoCodec
func (mr *MockAppModuleGenesisMockRecorder) RegisterLegacyAminoCodec(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterLegacyAminoCodec", reflect.TypeOf((*MockAppModuleGenesis)(nil).RegisterLegacyAminoCodec), arg0)
}

// RegisterInterfaces mocks base method
func (m *MockAppModuleGenesis) RegisterInterfaces(arg0 types.InterfaceRegistry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInterfaces", arg0)
}

// RegisterInterfaces indicates an expected call of RegisterInterfaces
func (mr *MockAppModuleGenesisMockRecorder) RegisterInterfaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInterfaces", reflect.TypeOf((*MockAppModuleGenesis)(nil).RegisterInterfaces), arg0)
}

// DefaultGenesis mocks base method
func (m *MockAppModuleGenesis) DefaultGenesis(arg0 codec.JSONMarshaler) json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultGenesis", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// DefaultGenesis indicates an expected call of DefaultGenesis
func (mr *MockAppModuleGenesisMockRecorder) DefaultGenesis(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultGenesis", reflect.TypeOf((*MockAppModuleGenesis)(nil).DefaultGenesis), arg0)
}

// ValidateGenesis mocks base method
func (m *MockAppModuleGenesis) ValidateGenesis(arg0 codec.JSONMarshaler, arg1 client.TxEncodingConfig, arg2 json.RawMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateGenesis", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateGenesis indicates an expected call of ValidateGenesis
func (mr *MockAppModuleGenesisMockRecorder) ValidateGenesis(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateGenesis", reflect.TypeOf((*MockAppModuleGenesis)(nil).ValidateGenesis), arg0, arg1, arg2)
}

// RegisterRESTRoutes mocks base method
func (m *MockAppModuleGenesis) RegisterRESTRoutes(arg0 client.Context, arg1 *mux.Router) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterRESTRoutes", arg0, arg1)
}

// RegisterRESTRoutes indicates an expected call of RegisterRESTRoutes
func (mr *MockAppModuleGenesisMockRecorder) RegisterRESTRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRESTRoutes", reflect.TypeOf((*MockAppModuleGenesis)(nil).RegisterRESTRoutes), arg0, arg1)
}

// RegisterGRPCGatewayRoutes mocks base method
func (m *MockAppModuleGenesis) RegisterGRPCGatewayRoutes(arg0 client.Context, arg1 *runtime.ServeMux) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterGRPCGatewayRoutes", arg0, arg1)
}

// RegisterGRPCGatewayRoutes indicates an expected call of RegisterGRPCGatewayRoutes
func (mr *MockAppModuleGenesisMockRecorder) RegisterGRPCGatewayRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCGatewayRoutes", reflect.TypeOf((*MockAppModuleGenesis)(nil).RegisterGRPCGatewayRoutes), arg0, arg1)
}

// GetTxCmd mocks base method
func (m *MockAppModuleGenesis) GetTxCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetTxCmd indicates an expected call of GetTxCmd
func (mr *MockAppModuleGenesisMockRecorder) GetTxCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxCmd", reflect.TypeOf((*MockAppModuleGenesis)(nil).GetTxCmd))
}

// GetQueryCmd mocks base method
func (m *MockAppModuleGenesis) GetQueryCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetQueryCmd indicates an expected call of GetQueryCmd
func (mr *MockAppModuleGenesisMockRecorder) GetQueryCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryCmd", reflect.TypeOf((*MockAppModuleGenesis)(nil).GetQueryCmd))
}

// InitGenesis mocks base method
func (m *MockAppModuleGenesis) InitGenesis(arg0 types0.Context, arg1 codec.JSONMarshaler, arg2 json.RawMessage) []types1.ValidatorUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitGenesis", arg0, arg1, arg2)
	ret0, _ := ret[0].([]types1.ValidatorUpdate)
	return ret0
}

// InitGenesis indicates an expected call of InitGenesis
func (mr *MockAppModuleGenesisMockRecorder) InitGenesis(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitGenesis", reflect.TypeOf((*MockAppModuleGenesis)(nil).InitGenesis), arg0, arg1, arg2)
}

// ExportGenesis mocks base method
func (m *MockAppModuleGenesis) ExportGenesis(arg0 types0.Context, arg1 codec.JSONMarshaler) json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportGenesis", arg0, arg1)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// ExportGenesis indicates an expected call of ExportGenesis
func (mr *MockAppModuleGenesisMockRecorder) ExportGenesis(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportGenesis", reflect.TypeOf((*MockAppModuleGenesis)(nil).ExportGenesis), arg0, arg1)
}

// MockAppModule is a mock of AppModule interface
type MockAppModule struct {
	ctrl     *gomock.Controller
	recorder *MockAppModuleMockRecorder
}

// MockAppModuleMockRecorder is the mock recorder for MockAppModule
type MockAppModuleMockRecorder struct {
	mock *MockAppModule
}

// NewMockAppModule creates a new mock instance
func NewMockAppModule(ctrl *gomock.Controller) *MockAppModule {
	mock := &MockAppModule{ctrl: ctrl}
	mock.recorder = &MockAppModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppModule) EXPECT() *MockAppModuleMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockAppModule) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockAppModuleMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAppModule)(nil).Name))
}

// RegisterLegacyAminoCodec mocks base method
func (m *MockAppModule) RegisterLegacyAminoCodec(arg0 *codec.LegacyAmino) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterLegacyAminoCodec", arg0)
}

// RegisterLegacyAminoCodec indicates an expected call of RegisterLegacyAminoCodec
func (mr *MockAppModuleMockRecorder) RegisterLegacyAminoCodec(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterLegacyAminoCodec", reflect.TypeOf((*MockAppModule)(nil).RegisterLegacyAminoCodec), arg0)
}

// RegisterInterfaces mocks base method
func (m *MockAppModule) RegisterInterfaces(arg0 types.InterfaceRegistry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInterfaces", arg0)
}

// RegisterInterfaces indicates an expected call of RegisterInterfaces
func (mr *MockAppModuleMockRecorder) RegisterInterfaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInterfaces", reflect.TypeOf((*MockAppModule)(nil).RegisterInterfaces), arg0)
}

// DefaultGenesis mocks base method
func (m *MockAppModule) DefaultGenesis(arg0 codec.JSONMarshaler) json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultGenesis", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// DefaultGenesis indicates an expected call of DefaultGenesis
func (mr *MockAppModuleMockRecorder) DefaultGenesis(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultGenesis", reflect.TypeOf((*MockAppModule)(nil).DefaultGenesis), arg0)
}

// ValidateGenesis mocks base method
func (m *MockAppModule) ValidateGenesis(arg0 codec.JSONMarshaler, arg1 client.TxEncodingConfig, arg2 json.RawMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateGenesis", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateGenesis indicates an expected call of ValidateGenesis
func (mr *MockAppModuleMockRecorder) ValidateGenesis(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateGenesis", reflect.TypeOf((*MockAppModule)(nil).ValidateGenesis), arg0, arg1, arg2)
}

// RegisterRESTRoutes mocks base method
func (m *MockAppModule) RegisterRESTRoutes(arg0 client.Context, arg1 *mux.Router) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterRESTRoutes", arg0, arg1)
}

// RegisterRESTRoutes indicates an expected call of RegisterRESTRoutes
func (mr *MockAppModuleMockRecorder) RegisterRESTRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRESTRoutes", reflect.TypeOf((*MockAppModule)(nil).RegisterRESTRoutes), arg0, arg1)
}

// RegisterGRPCGatewayRoutes mocks base method
func (m *MockAppModule) RegisterGRPCGatewayRoutes(arg0 client.Context, arg1 *runtime.ServeMux) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterGRPCGatewayRoutes", arg0, arg1)
}

// RegisterGRPCGatewayRoutes indicates an expected call of RegisterGRPCGatewayRoutes
func (mr *MockAppModuleMockRecorder) RegisterGRPCGatewayRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCGatewayRoutes", reflect.TypeOf((*MockAppModule)(nil).RegisterGRPCGatewayRoutes), arg0, arg1)
}

// GetTxCmd mocks base method
func (m *MockAppModule) GetTxCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetTxCmd indicates an expected call of GetTxCmd
func (mr *MockAppModuleMockRecorder) GetTxCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxCmd", reflect.TypeOf((*MockAppModule)(nil).GetTxCmd))
}

// GetQueryCmd mocks base method
func (m *MockAppModule) GetQueryCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetQueryCmd indicates an expected call of GetQueryCmd
func (mr *MockAppModuleMockRecorder) GetQueryCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryCmd", reflect.TypeOf((*MockAppModule)(nil).GetQueryCmd))
}

// InitGenesis mocks base method
func (m *MockAppModule) InitGenesis(arg0 types0.Context, arg1 codec.JSONMarshaler, arg2 json.RawMessage) []types1.ValidatorUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitGenesis", arg0, arg1, arg2)
	ret0, _ := ret[0].([]types1.ValidatorUpdate)
	return ret0
}

// InitGenesis indicates an expected call of InitGenesis
func (mr *MockAppModuleMockRecorder) InitGenesis(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitGenesis", reflect.TypeOf((*MockAppModule)(nil).InitGenesis), arg0, arg1, arg2)
}

// ExportGenesis mocks base method
func (m *MockAppModule) ExportGenesis(arg0 types0.Context, arg1 codec.JSONMarshaler) json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportGenesis", arg0, arg1)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// ExportGenesis indicates an expected call of ExportGenesis
func (mr *MockAppModuleMockRecorder) ExportGenesis(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportGenesis", reflect.TypeOf((*MockAppModule)(nil).ExportGenesis), arg0, arg1)
}

// RegisterInvariants mocks base method
func (m *MockAppModule) RegisterInvariants(arg0 types0.InvariantRegistry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInvariants", arg0)
}

// RegisterInvariants indicates an expected call of RegisterInvariants
func (mr *MockAppModuleMockRecorder) RegisterInvariants(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInvariants", reflect.TypeOf((*MockAppModule)(nil).RegisterInvariants), arg0)
}

// Route mocks base method
func (m *MockAppModule) Route() types0.Route {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Route")
	ret0, _ := ret[0].(types0.Route)
	return ret0
}

// Route indicates an expected call of Route
func (mr *MockAppModuleMockRecorder) Route() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Route", reflect.TypeOf((*MockAppModule)(nil).Route))
}

// QuerierRoute mocks base method
func (m *MockAppModule) QuerierRoute() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuerierRoute")
	ret0, _ := ret[0].(string)
	return ret0
}

// QuerierRoute indicates an expected call of QuerierRoute
func (mr *MockAppModuleMockRecorder) QuerierRoute() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerierRoute", reflect.TypeOf((*MockAppModule)(nil).QuerierRoute))
}

// LegacyQuerierHandler mocks base method
func (m *MockAppModule) LegacyQuerierHandler(arg0 *codec.LegacyAmino) types0.Querier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LegacyQuerierHandler", arg0)
	ret0, _ := ret[0].(types0.Querier)
	return ret0
}

// LegacyQuerierHandler indicates an expected call of LegacyQuerierHandler
func (mr *MockAppModuleMockRecorder) LegacyQuerierHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LegacyQuerierHandler", reflect.TypeOf((*MockAppModule)(nil).LegacyQuerierHandler), arg0)
}

// RegisterServices mocks base method
func (m *MockAppModule) RegisterServices(arg0 module.Configurator) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterServices", arg0)
}

// RegisterServices indicates an expected call of RegisterServices
func (mr *MockAppModuleMockRecorder) RegisterServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterServices", reflect.TypeOf((*MockAppModule)(nil).RegisterServices), arg0)
}

// BeginBlock mocks base method
func (m *MockAppModule) BeginBlock(arg0 types0.Context, arg1 types1.RequestBeginBlock) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BeginBlock", arg0, arg1)
}

// BeginBlock indicates an expected call of BeginBlock
func (mr *MockAppModuleMockRecorder) BeginBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginBlock", reflect.TypeOf((*MockAppModule)(nil).BeginBlock), arg0, arg1)
}

// EndBlock mocks base method
func (m *MockAppModule) EndBlock(arg0 types0.Context, arg1 types1.RequestEndBlock) []types1.ValidatorUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndBlock", arg0, arg1)
	ret0, _ := ret[0].([]types1.ValidatorUpdate)
	return ret0
}

// EndBlock indicates an expected call of EndBlock
func (mr *MockAppModuleMockRecorder) EndBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndBlock", reflect.TypeOf((*MockAppModule)(nil).EndBlock), arg0, arg1)
}
