package methodarea

import "strings"

type Descriptor struct {
	IsInt       bool
	IsByte      bool
	IsChar      bool
	IsBool      bool
	IsShort     bool
	IsFloat     bool
	IsDouble    bool
	IsLong      bool
	IsReference bool
	IsArray     bool
}

func (this *Descriptor) DefaultValue() interface{} {
	if this.IsInt {
		return 0
	}
	if this.IsByte {
		return int8(0)
	}
	if this.IsChar {
		return uint16(0)
	}
	if this.IsBool {
		return false
	}
	if this.IsShort {
		return int16(0)
	}
	if this.IsFloat {
		return float32(0)
	}
	if this.IsDouble {
		return float64(0)
	}
	if this.IsLong {
		return int64(0)
	}
	if this.IsReference {
		return nil
	}
	if this.IsArray {
		return nil
	}
	panic("unknown descriptor")
}

func newDescriptor(descriptor string) *Descriptor {
	var isInt bool
	var isByte bool
	var isChar bool
	var isBool bool
	var isShort bool
	var isFloat bool
	var isDouble bool
	var isLong bool
	var isReference bool
	var isArray bool
	switch {
	case descriptor == "B":
		isByte = true
	case descriptor == "C":
		isChar = true
	case descriptor == "D":
		isDouble = true
	case descriptor == "F":
		isFloat = true
	case descriptor == "I":
		isInt = true
	case descriptor == "J":
		isLong = true
	case strings.HasPrefix(descriptor, "L"):
		isReference = true
	case descriptor == "S":
		isShort = true
	case descriptor == "Z":
		isBool = true
	case descriptor == "[":
		isArray = true
	}

	return &Descriptor{
		IsInt:       isInt,
		IsByte:      isByte,
		IsChar:      isChar,
		IsBool:      isBool,
		IsShort:     isShort,
		IsFloat:     isFloat,
		IsDouble:    isDouble,
		IsLong:      isLong,
		IsReference: isReference,
		IsArray:     isArray,
	}
}
