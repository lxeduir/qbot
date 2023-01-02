package public

import (
	"regexp"
	"strings"
)

type CqCode struct {
	CQ    string
	Name  []string
	Value []string
}
type CqCodes struct {
	CqCode []CqCode
}

func (c CqCode) Get() string {
	str := "[CQ:" + c.CQ
	for i, v := range c.Name {
		str += "," + v + "=" + c.Value[i]
	}
	str += "]"
	return str
}
func (C CqCodes) Get() string {
	str := ""
	for _, c := range C.CqCode {
		str += "[CQ:" + c.CQ
		for i, v := range c.Name {
			if v == "type" && c.Value[i] == "flash" {
				continue
			}
			str += "," + v + "=" + c.Value[i]
		}
		str += "]"
	}
	return str
}
func CQGets(C []CqCode) string {
	str := ""
	for _, c := range C {
		str += "[CQ:" + c.CQ
		for i, v := range c.Name {
			if v == "type" && c.Value[i] == "flash" {
				continue
			}
			str += "," + v + "=" + c.Value[i]
		}
		str += "]"
	}
	return str
}
func Parse(CQ string) CqCodes {
	var Cq CqCodes
	cq := CqExtract(CQ)
	var matchsReg = regexp.MustCompile(`\[CQ:([^,]+),([^=]+=[^,]+)(?:,([^=]+=[^,]+))*]`)
	var matchReg = regexp.MustCompile(`\[CQ:([^,]+),([^=]+=[^,]+)(?:,([^=]+=[^,]+))*](?:,\[CQ:([^,]+),([^=]+=[^,]+)(?:,([^=]+=[^,]+))*])*`)
	for _, c := range cq {
		if c == "" {
			continue
		}
		var C CqCode
		var matchs = matchsReg.FindAllStringSubmatch(c, -1)[0]
		C.CQ = matchs[1]
		nv := matchReg.FindAllStringSubmatch(c, -1)[0][2:]
		for _, v := range nv {
			if v == "" {
				continue
			}
			C.Name = append(C.Name, strings.Split(v, "=")[0])
			C.Value = append(C.Value, strings.Split(v, "=")[1])
		}
		Cq.CqCode = append(Cq.CqCode, C)
	}

	return Cq
}

func CqExtract(CQ string) (s []string) {
	CQ = strings.Replace(CQ, "]", "]---@#$%^", -1)
	CQ = strings.Replace(CQ, "=,", "=?,", -1)
	s1 := strings.Split(CQ, "---@#$%^")
	for _, v := range s1 {
		flag := false
		str := ""
		for _, x := range v {
			if x == '[' {
				flag = true
			} else if x == ']' {
				str += "]"
				s = append(s, str)
				str = ""
				break
			}
			if flag {
				str += string(x)
			}
		}
	}
	return s
}
