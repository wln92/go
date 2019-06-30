package service

import (
	"fmt"
	"io"
	"log"
	"protos"
	"strconv"

	"golang.org/x/net/context"
)

//对外提供的工厂函数
func NewUserService() *UserService {
	return &UserService{}
}

//**************************************************************
// 接口实现，接口定义是在proto生成的.pb.go文件中
//**************************************************************

// 接口实现对象，属性成员根据而业务自定义
type UserService struct {
}

// Get接口方法实现
func (this *UserService) Get(ctx context.Context, req *protos.UserReq) (*protos.User, error) {
	return &protos.User{Id: 1, Name: "shuai"}, nil
}

// GetList接口方法实现
func (this *UserService) GetList(req *protos.UserReq, stream protos.IUserService_GetListServer) error {
	fmt.Println(*req)
	// 流式返回多条数据
	for i := 0; i < 5; i++ {
		stream.Send(&protos.User{Id: int32(i), Name: "我是" + strconv.Itoa(i)})
	}
	return nil
}

// WaitGet接口方法实现
func (this *UserService) WaitGet(reqStream protos.IUserService_WaitGetServer) error {
	for { // 接收流式请求并返回单一对象
		userReq, err := reqStream.Recv()
		if err != io.EOF {
			fmt.Println("流请求~", *userReq)
		} else {
			return reqStream.SendAndClose(&protos.User{Id: 100, Name: "shuai"})
		}
	}
}

//双向流：请求流和响应流异步
func (this *UserService) LoopGet(reqStream protos.IUserService_LoopGetServer) error {
	for {
		userReq, err := reqStream.Recv()
		if err == io.EOF { //请求结束
			return nil
		}
		if err != nil {
			return err
		}
		if err = reqStream.Send(&protos.User{Id: userReq.Id, Name: "shuai"}); err != nil {
			return err
		}
	}
}
