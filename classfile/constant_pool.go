package classfile

type ConstantPool []ConstantInfo

type ExtensionConstantPool struct {
	ConstantPool
}

func (this ConstantPool) getUtf8(idx uint16) string {
	return this[idx-1].(*ConstantUtf8Info).str
}

func (this ConstantPool) GetClassInfo(idx uint16) *ConstantClassInfo {
	return this[idx-1].(*ConstantClassInfo)
}

func (this ConstantPool) getNameAndTypeInfo(idx uint16) *ConstantNameType {
	return this[idx-1].(*ConstantNameType)
}

func (this ConstantPool) GetFieldRef(idx uint16) *ConstantFieldRefInfo {
	return this[idx-1].(*ConstantFieldRefInfo)
}

func (this ConstantPool) GetMethodRef(idx uint16) *ConstantMethodRefInfo {
	return this[idx-1].(*ConstantMethodRefInfo)
}

func readConstantPool(reader *ClassReader) ConstantPool {
	poolSize := reader.readUInt16()

	var pool ConstantPool = make([]ConstantInfo, poolSize)

	for i := 1; i < int(poolSize); i++ {
		info := ReadConstantInfo(reader, pool)
		// pool = append(pool, info)
		pool[i-1] = info
		switch info.(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return pool
}
