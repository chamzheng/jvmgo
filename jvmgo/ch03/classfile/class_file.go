package classfile

import "fmt"

type ClassFile struct {
	maigc        uint32          //魔数字CAFE BABE
	minorVersion uint16          //子版本号
	majorVersion uint16          //主版本号
	constantPool ConstantPool    //常量池
	accessFlags  uint16          //访问控制
	thisClass    uint16          //本类
	superClass   uint16          //父类
	interfaces   []uint16        //接口
	fields       []*MemberInfo   //字段
	methods      []*MemberInfo   //方法
	attributes   []AttributeInfo //
}
