package main

import (
	"fmt"
	"sync"
)

type Empty interface { //命名空接口
}

type Registry interface {
	createSubnet(network, sn, data string) (string, error)
	updateSubnet(network, sn, data string) (string, error)
	deleteSubnet(network, sn string) (string, error)
	watchSubnets(network string, since uint64) (string, error)
	getSubnets(a ...interface{}) (string, error) //匿名空接口
	SubnetRegistry                               //接口嵌套
}

type SubnetRegistry interface {
	createConfig(network, sn, data string) (string, error)
	updateConfig(network, sn, data string) (string, error)
	deleteConfig(network, sn string) (string, error)
	watchConfig(network string, since uint64) (string, error)
}

/*
 (1)只要某个类型(例如:etcdSubnetRegistry)拥有该接口的所有方法签名,即算实现该接口
    无需显式声明实现了那个接口(例如:Registry),该种风格称为Structural Typing
 (2)接口只有方法签名,没有数据字段
 (3)接口可以嵌入到其他接口,或结构中
 (4)空接口可以作为任何数据类型的容器,
    也可以说任何类型都实现了空接口,任何类型都是空接口的子类型
    空接口类似Java中的java.lang.Object,C#中System.Object
*/
type etcdSubnetRegistry struct {
	mux     sync.Mutex
	cli     string
	etcdCfg string
}

func (esr *etcdSubnetRegistry) createSubnet(network, sn, data string) (string, error) {
	//TODO:
	return "", nil
}
func (esr *etcdSubnetRegistry) updateSubnet(network, sn, data string) (string, error) {
	//TODO:
	return "", nil
}
func (esr *etcdSubnetRegistry) deleteSubnet(network, sn string) (string, error) {
	//TODO:
	return "", nil
}
func (esr *etcdSubnetRegistry) watchSubnets(network string, since uint64) (string, error) {
	//TODO:
	return "", nil
}

func (esr *etcdSubnetRegistry) getSubnets(a ...interface{}) (string, error) {
	//TODO:
	return "", nil
}

func (esr *etcdSubnetRegistry) createConfig(network, sn, data string) (string, error) {
	//TODO:
	return "", nil
}
func (esr *etcdSubnetRegistry) updateConfig(network, sn, data string) (string, error) {
	//TODO:
	return "", nil
}
func (esr *etcdSubnetRegistry) deleteConfig(network, sn string) (string, error) {
	//TODO:
	return "", nil
}
func (esr *etcdSubnetRegistry) watchConfig(network string, since uint64) (string, error) {
	//TODO:
	return "", nil
}
