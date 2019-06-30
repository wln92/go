package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"./protos" // 此为自定义的protos包，存放的是.proto文件和对应的.pb.go文件

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

var (
	// 命令行参数-host，默认地址本机9000端口
	addr = flag.String("host", "127.0.0.1:9000", "")
	// UserService服务存根，可直接调用服务方法
	userCli protos.IUserServiceClient
)

func main() {
	// 创建grpc的服务连接
	conn, err := grpc.Dial(*addr)
	if err != nil {
		log.Fatal("failed to connect : ", err)
	}
	// 存根
	userCli = protos.NewIUserServiceClient(conn)

	// 测试前三种数据传递方式，第四种省略(二三结合)
	TestGet()
	TestGetList()
	TestWaitGet()
}

func TestGet() {
	// 测试调用方法Get，返回user对象
	user, err := userCli.Get(context.Background(), &protos.UserReq{Id: 1})
	if err != nil {
		fmt.Printf("Get connect failed :%v", err)
		return
	}
	log.Println("Get响应数据：", *user)
}

func TestGetList() {
	// 测试调用方法GetList，返回一个Stream流，循环获取多个user对象
	recvStream, err := userCli.GetList(context.Background(), &protos.UserReq{Id: 1})
	if err != nil {
		fmt.Printf("GetList connect failed :%v", err)
		return
	}
	for {
		user, err := recvStream.Recv()
		if err == io.EOF {
			break
		}
		log.Println("GetList获取的一条响应数据：", *user)
	}
}

func TestWaitGet() {
	// 测试调用方法WaitGet，传入多条请求数据，返回一个user对象
	sendStream, err := userCli.WaitGet(context.Background())
	if err != nil {
		fmt.Printf("WaitGet connect failed :%v", err)
		return
	}
	for i := 0; i < 5; i++ { // 一次传入5条请求数据
		if err = sendStream.Send(&protos.UserReq{Id: int32(i)}); err != nil {
			fmt.Printf("WaitGet send failed :%v", err)
			return
		}
	}
	// 服务端接受全部请求数据后，返回一个user对象
	user, err := sendStream.CloseAndRecv()
	if err != nil {
		fmt.Printf("WaitGet recv failed :%v", err)
		return
	}
	log.Println("WaitGet响应数据：", *user)
}
