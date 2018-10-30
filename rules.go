package main

import (
	"fmt"
	"net"

	"k8s.io/apimachinery/pkg/util/intstr"
)

// RuleName is ..
type RuleName string

// Rules
const (
	MaskBitsSize RuleName = "MaskBitsSize"
	ListSize     RuleName = "ListSize"
	LabelValues  RuleName = "LabelValues"
	LabelCount   RuleName = "LabelCount"
)

// Rule is ...
type Rule struct {
	Name     RuleName           `json:"name"`
	Operator Operator           `json:"operator"`
	Key      string             `json:"key"`
	Value    intstr.IntOrString `json:"value"`
}

func (v *Rule) isValidMaskBitsSize(l []string) (bool, error) {

	for _, i := range l {

		_, network, _ := net.ParseCIDR(i)
		m, _ := network.Mask.Size()

		ok, err := operatorExec(m, v.Value.IntValue(), v.Operator)
		if err != nil {
			return false, err
		}

		if !ok {
			return false, fmt.Errorf(
				"error InvalidNaskBitsSize: mask size must be %s %v",
				v.Operator, v.Value.IntValue())
		}

	}

	return true, nil

}

func (v *Rule) isValidListSize(l []string) (bool, error) {

	ok, err := operatorExec(len(l), v.Value.IntValue(), v.Operator)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, fmt.Errorf(
			"error InvalidListSize: list size must be %s %v ",
			v.Operator, v.Value.IntValue())
	}

	return true, nil

}

func (v *Rule) isValidLabelValues(labels map[string]string) (bool, error) {

	ok, err := operatorExec(len(labels), v.Value.IntValue(), v.Operator)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, fmt.Errorf(
			"error InvalidLabelCount: the numbers of labels must be %s %v",
			v.Operator, v.Value.IntValue())
	}

	return true, nil

}

func (v *Rule) isValidLabelCount(labels map[string]string) (bool, error) {

	ok, err := operatorExec(len(labels), v.Value.IntValue(), v.Operator)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, fmt.Errorf(
			"error InvalidLabelCount: the numbers of labels must be %s %v",
			v.Operator, v.Value.IntValue())
	}

	return true, nil

}
