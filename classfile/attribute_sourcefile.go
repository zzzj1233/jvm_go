package classfile

type AttributeSourceFile struct {
	sourceFile string
	cp         ConstantPool
}

func (this AttributeSourceFile) readInfo(reader *ClassReader) {
	this.sourceFile = this.cp.getUtf8(reader.readUInt16())
}

type LineNumberTable struct {
	startPc    uint16
	lineNumber uint16
}

type AttributeLineNumberTable struct {
	lineNumberTable []*LineNumberTable
}

func (this AttributeLineNumberTable) readInfo(reader *ClassReader) {
	tableLength := reader.readUInt16()

	this.lineNumberTable = make([]*LineNumberTable, 0, tableLength)

	for i := 0; i < int(tableLength); i++ {
		this.lineNumberTable = append(this.lineNumberTable, &LineNumberTable{
			reader.readUInt16(),
			reader.readUInt16(),
		})
	}

}

type AttributeLocalVarTable struct {
	localVarTable []*LocalVarTable
}

func (this AttributeLocalVarTable) readInfo(reader *ClassReader) {
	tableLength := reader.readUInt16()
	this.localVarTable = make([]*LocalVarTable, 0, tableLength)

	for i := 0; i < int(tableLength); i++ {
		this.localVarTable = append(this.localVarTable, &LocalVarTable{
			LineNumberTable: LineNumberTable{
				startPc:    reader.readUInt16(),
				lineNumber: reader.readUInt16(),
			},
			nameIndex: reader.readUInt16(),
			descIndex: reader.readUInt16(),
			index:     reader.readUInt16(),
		})
	}

}

type LocalVarTable struct {
	LineNumberTable
	nameIndex uint16
	descIndex uint16
	index     uint16
}
