package main

import (
	"errors"
	"fmt"
	"math/rand"
)

//服务端接口
type ConsistentHashServer interface {
	//添加主机节点
	addServer(server string) error

	//找到存储数据节点
	fidServer(key string) (server string)
}

//服务器列表
type ServerList struct {

	//服务器列表
	serverList map[string]string

	//节点对应位置
	virtualPosList map[int]string
}

func (serverList *ServerList) fidServer(key string) string {

	//根据key获取位置，这里就随便算了
	pos := int(rand.Intn(100000))

	for p, s := range serverList.virtualPosList {
		if p < pos {
			return s
		}
	}

	return serverList.virtualPosList[0]
}

//添加主机节点
func (serverList *ServerList) addServer(server string) error {

	//去除无效节点
	if server == "" {
		return errors.New("无效的服务。")
	}

	//初始化
	if serverList.serverList == nil {
		serverList.serverList = map[string]string{}
	}

	if serverList.virtualPosList == nil {
		serverList.virtualPosList = map[int]string{}
	}

	//已挂载则直接返回
	if _, ok := serverList.serverList[server]; ok {
		return nil
	}

	//设置虚拟节点为5个，挂载节点
	for i := 0; i < 5; i++ {

		//哈希算法计算位置-这里随便算了一个随机数代替
		pos := int(rand.Intn(100000))

		//挂载
		serverList.virtualPosList[pos] = server
	}

	serverList.serverList[server] = server

	return nil
}

//一致哈希
func main() {

	//继承接口
	var s ConsistentHashServer
	s = &ServerList{}

	//添加机器
	s.addServer("127.0.0.1")
	s.addServer("192.168.0.27")

	fmt.Println(s.fidServer("1"))
}
