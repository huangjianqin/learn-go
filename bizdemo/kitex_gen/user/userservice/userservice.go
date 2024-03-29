// Code generated by Kitex v0.4.2. DO NOT EDIT.

package userservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	user "learn-go/bizdemo/kitex_gen/user"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateUser": kitex.NewMethodInfo(createUserHandler, newCreateUserArgs, newCreateUserResult, false),
		"MGetUser":   kitex.NewMethodInfo(mGetUserHandler, newMGetUserArgs, newMGetUserResult, false),
		"CheckUser":  kitex.NewMethodInfo(checkUserHandler, newCheckUserArgs, newCheckUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.2",
		Extra:           extra,
	}
	return svcInfo
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.CreateUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).CreateUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateUserArgs:
		success, err := handler.(user.UserService).CreateUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateUserResult)
		realResult.Success = success
	}
	return nil
}
func newCreateUserArgs() interface{} {
	return &CreateUserArgs{}
}

func newCreateUserResult() interface{} {
	return &CreateUserResult{}
}

type CreateUserArgs struct {
	Req *user.CreateUserRequest
}

func (p *CreateUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.CreateUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateUserArgs) Unmarshal(in []byte) error {
	msg := new(user.CreateUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateUserArgs_Req_DEFAULT *user.CreateUserRequest

func (p *CreateUserArgs) GetReq() *user.CreateUserRequest {
	if !p.IsSetReq() {
		return CreateUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateUserResult struct {
	Success *user.CreateUserResponse
}

var CreateUserResult_Success_DEFAULT *user.CreateUserResponse

func (p *CreateUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.CreateUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateUserResult) Unmarshal(in []byte) error {
	msg := new(user.CreateUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateUserResult) GetSuccess() *user.CreateUserResponse {
	if !p.IsSetSuccess() {
		return CreateUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.CreateUserResponse)
}

func (p *CreateUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func mGetUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.MGetUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).MGetUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *MGetUserArgs:
		success, err := handler.(user.UserService).MGetUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MGetUserResult)
		realResult.Success = success
	}
	return nil
}
func newMGetUserArgs() interface{} {
	return &MGetUserArgs{}
}

func newMGetUserResult() interface{} {
	return &MGetUserResult{}
}

type MGetUserArgs struct {
	Req *user.MGetUserRequest
}

func (p *MGetUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.MGetUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MGetUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MGetUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MGetUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MGetUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *MGetUserArgs) Unmarshal(in []byte) error {
	msg := new(user.MGetUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MGetUserArgs_Req_DEFAULT *user.MGetUserRequest

func (p *MGetUserArgs) GetReq() *user.MGetUserRequest {
	if !p.IsSetReq() {
		return MGetUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MGetUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type MGetUserResult struct {
	Success *user.MGetUserResponse
}

var MGetUserResult_Success_DEFAULT *user.MGetUserResponse

func (p *MGetUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.MGetUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MGetUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MGetUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MGetUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MGetUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *MGetUserResult) Unmarshal(in []byte) error {
	msg := new(user.MGetUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MGetUserResult) GetSuccess() *user.MGetUserResponse {
	if !p.IsSetSuccess() {
		return MGetUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MGetUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.MGetUserResponse)
}

func (p *MGetUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func checkUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.CheckUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).CheckUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CheckUserArgs:
		success, err := handler.(user.UserService).CheckUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CheckUserResult)
		realResult.Success = success
	}
	return nil
}
func newCheckUserArgs() interface{} {
	return &CheckUserArgs{}
}

func newCheckUserResult() interface{} {
	return &CheckUserResult{}
}

type CheckUserArgs struct {
	Req *user.CheckUserRequest
}

func (p *CheckUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.CheckUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CheckUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CheckUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CheckUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CheckUserArgs) Unmarshal(in []byte) error {
	msg := new(user.CheckUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckUserArgs_Req_DEFAULT *user.CheckUserRequest

func (p *CheckUserArgs) GetReq() *user.CheckUserRequest {
	if !p.IsSetReq() {
		return CheckUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckUserResult struct {
	Success *user.CheckUserResponse
}

var CheckUserResult_Success_DEFAULT *user.CheckUserResponse

func (p *CheckUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.CheckUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CheckUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CheckUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CheckUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CheckUserResult) Unmarshal(in []byte) error {
	msg := new(user.CheckUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckUserResult) GetSuccess() *user.CheckUserResponse {
	if !p.IsSetSuccess() {
		return CheckUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.CheckUserResponse)
}

func (p *CheckUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateUser(ctx context.Context, Req *user.CreateUserRequest) (r *user.CreateUserResponse, err error) {
	var _args CreateUserArgs
	_args.Req = Req
	var _result CreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetUser(ctx context.Context, Req *user.MGetUserRequest) (r *user.MGetUserResponse, err error) {
	var _args MGetUserArgs
	_args.Req = Req
	var _result MGetUserResult
	if err = p.c.Call(ctx, "MGetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckUser(ctx context.Context, Req *user.CheckUserRequest) (r *user.CheckUserResponse, err error) {
	var _args CheckUserArgs
	_args.Req = Req
	var _result CheckUserResult
	if err = p.c.Call(ctx, "CheckUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
