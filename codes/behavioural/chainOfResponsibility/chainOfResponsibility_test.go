package chainofresponsibility

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	cashier := &cashier{}
	// 为医疗部门设定下一个目标
	medical := &medical{}
	medical.setNext(cashier)
	// 为医生部门设定下一个目标
	doctor := &doctor{}
	doctor.setNext(medical)
	// 为接待部门设定下一个目标
	reception := &reception{}
	reception.setNext(doctor)
	patient := &patient{name: "abec"}
	// 病人探访
	reception.execute(patient)
}
