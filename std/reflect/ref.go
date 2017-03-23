package main

import (
	"reflect"
	"fmt"
)

type ExecutionServiceServer interface {
	Create(string, string) (string, error)
	Start(string, string) (string, error)
	Update(string, string) (string, error)
	Pause(string, string) (string, error)
	Resume(string, string) (string, error)
	Delete(string, string) (string, error)
	Get(string, string) (string, error)
	List(string, string) (string, error)
	StartProcess(string, string) (string, error)
	GetProcess(string, string) (string, error)
	SignalProcess(string, string) (string, error)
	DeleteProcess(string, string) (string, error)
	ListProcesses(string, string) (string, error)
}

type Service struct {
	executor   string
	supervisor string
}

func New(executor, supervisor string) (*Service) {
	return &Service{
		executor: executor,
		supervisor: supervisor,
	}
}

func (s *Service) Create(string, string) (string, error) {
	fmt.Println("Create")
	return "Create", nil
}
func (s *Service) Start(string, string) (string, error) {
	fmt.Println("Start")
	return "Start", nil
}
func (s *Service) Update(string, string) (string, error) {
	fmt.Println("Update")
	return "Update", nil
}
func (s *Service) Pause(string, string) (string, error) {
	fmt.Println("Pause")
	return "Pause", nil
}
func (s *Service) Resume(string, string) (string, error) {
	fmt.Println("Resume")
	return "Resume", nil
}
func (s *Service) Delete(string, string) (string, error) {
	fmt.Println("Delete")
	return "Delete", nil
}
func (s *Service) Get(string, string) (string, error) {
	fmt.Println("Get")
	return "Get", nil
}
func (s *Service) List(string, string) (string, error) {
	fmt.Println("List")
	return "List", nil
}
func (s *Service) StartProcess(string, string) (string, error) {
	fmt.Println("StartProcess")
	return "StartProcess", nil
}
func (s *Service) GetProcess(string, string) (string, error) {
	fmt.Println("GetProcess")
	return "GetProcess", nil
}
func (s *Service) SignalProcess(string, string) (string, error) {
	fmt.Println("SignalProcess")
	return "SignalProcess", nil
}
func (s *Service) DeleteProcess(string, string) (string, error) {
	fmt.Println("DeleteProcess")
	return "DeleteProcess", nil
}
func (s *Service) ListProcesses(string, string) (string, error) {
	fmt.Println("ListProcesses")
	return "ListProcesses", nil
}
/*Extra Method from Interface*/
func (s *Service) MonitorProcesses(string, string) (string, error) {
	fmt.Println("monitorProcesses")
	return "monitorProcesses", nil
}

type ServiceDesc struct {
	ServiceName string
				// The pointer to the service interface. Used to check whether the user
				// provided implementation satisfies the interface requirements.
	HandlerType interface{} //https://golang.org/doc/faq#nil_error
	Methods     []MethodDesc
	Metadata    interface{}
}

// MethodDesc represents an RPC service's method specification.
type MethodDesc struct {
	MethodName string
	Handler    string
}

var _ExecutionService_serviceDesc = ServiceDesc{
	ServiceName: "containerd.v1.ExecutionService",
	HandlerType: (*ExecutionServiceServer)(nil), //https://golang.org/doc/faq#nil_error
	Methods: []MethodDesc{
		{
			MethodName: "Create",
			Handler:    "_ExecutionService_Create_Handler",
		},
		{
			MethodName: "Start",
			Handler:    "_ExecutionService_Start_Handler",
		},
		{
			MethodName: "Update",
			Handler:    "_ExecutionService_Update_Handler",
		},
		{
			MethodName: "Pause",
			Handler:    "_ExecutionService_Pause_Handler",
		},
		{
			MethodName: "Resume",
			Handler:    "_ExecutionService_Resume_Handler",
		},
		{
			MethodName: "Delete",
			Handler:    "_ExecutionService_Delete_Handler",
		},
		{
			MethodName: "Get",
			Handler:    "_ExecutionService_Get_Handler",
		},
		{
			MethodName: "List",
			Handler:    "_ExecutionService_List_Handler",
		},
		{
			MethodName: "StartProcess",
			Handler:    "_ExecutionService_StartProcess_Handler",
		},
		{
			MethodName: "GetProcess",
			Handler:    "_ExecutionService_GetProcess_Handler",
		},
		{
			MethodName: "SignalProcess",
			Handler:    "_ExecutionService_SignalProcess_Handler",
		},
		{
			MethodName: "DeleteProcess",
			Handler:    "_ExecutionService_DeleteProcess_Handler",
		},
		{
			MethodName: "ListProcesses",
			Handler:    "_ExecutionService_ListProcesses_Handler",
		},
	},
	Metadata: "execution.proto", //接口类型可以接收任何类型的值
}
/**
该示例说明如果实现接口则必须实现接口中所有的方法
 */
func main() {
	var val interface{} = int64(58)
	fmt.Println(reflect.TypeOf(val))
	val = 50
	fmt.Println(reflect.TypeOf(val))

	srv := New("excutor", "supervisor");
	RegisterService2(&_ExecutionService_serviceDesc, srv)
}

func RegisterService(sd *ServiceDesc, service interface{}) {
	ht := reflect.TypeOf(sd.HandlerType).Elem()
	serviceType := reflect.TypeOf(service)
	if !serviceType.Implements(ht) {
		fmt.Printf("grpc: Server.RegisterService found the handler of type %v that does not satisfy %v", serviceType, ht)
	} else {
		fmt.Printf("grpc: Server.RegisterService found the handler of type %v that  satisfy %v", serviceType, ht)
	}
}

func RegisterService2(sd *ServiceDesc, service interface{}) {

	handlerType := reflect.TypeOf(sd.HandlerType).Name()
	fmt.Printf("handlerType %v \n", handlerType)

	ht := reflect.TypeOf(sd.HandlerType).Elem()
	fmt.Printf("ht %v \n", ht)

	serviceType := reflect.TypeOf(service)
	if !serviceType.Implements(ht) {
		fmt.Printf("grpc: Server.RegisterService found the handler of type %v that does not satisfy %v \n", serviceType, ht)
	} else {
		fmt.Printf("grpc: Server.RegisterService found the handler of type %v that satisfy %v \n", serviceType, ht)
	}
}