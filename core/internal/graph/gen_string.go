// Code generated by "stringer -linecomment -type=MType,ParserType,FieldType -output=./gen_string.go"; DO NOT EDIT.

package graph

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[itemError-0]
	_ = x[itemEOF-1]
	_ = x[itemName-2]
	_ = x[itemOn-3]
	_ = x[itemPunctuator-4]
	_ = x[itemArgsOpen-5]
	_ = x[itemArgsClose-6]
	_ = x[itemListOpen-7]
	_ = x[itemListClose-8]
	_ = x[itemObjOpen-9]
	_ = x[itemObjClose-10]
	_ = x[itemColon-11]
	_ = x[itemEquals-12]
	_ = x[itemRequired-13]
	_ = x[itemDirective-14]
	_ = x[itemVariable-15]
	_ = x[itemSpread-16]
	_ = x[itemNumberVal-17]
	_ = x[itemStringVal-18]
	_ = x[itemBoolVal-19]
}

const _MType_name = "errorend of filelabel\"on\"punctuation !()[]{}:=()[]{}:=!@(directive)$variable...numberstringboolean"

var _MType_index = [...]uint8{0, 5, 16, 21, 25, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 67, 76, 79, 85, 91, 98}

func (i MType) String() string {
	if i < 0 || i >= MType(len(_MType_index)-1) {
		return "MType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MType_name[_MType_index[i]:_MType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[parserError-0]
	_ = x[parserEOF-1]
	_ = x[OpQuery-2]
	_ = x[OpMutate-3]
	_ = x[OpSub-4]
	_ = x[NodeStr-5]
	_ = x[NodeNum-6]
	_ = x[NodeBool-7]
	_ = x[NodeObj-8]
	_ = x[NodeList-9]
	_ = x[NodeVar-10]
	_ = x[NodeLabel-11]
}

const _ParserType_name = "parserErrorparserEOFOpQueryOpMutateOpSubNodeStrNodeNumNodeBoolNodeObjNodeListNodeVarNodeLabel"

var _ParserType_index = [...]uint8{0, 11, 20, 27, 35, 40, 47, 54, 62, 69, 77, 84, 93}

func (i ParserType) String() string {
	if i < 0 || i >= ParserType(len(_ParserType_index)-1) {
		return "ParserType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ParserType_name[_ParserType_index[i]:_ParserType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FieldUnion-1]
	_ = x[FieldMember-2]
	_ = x[FieldKeyword-3]
}

const _FieldType_name = "FieldUnionFieldMemberFieldKeyword"

var _FieldType_index = [...]uint8{0, 10, 21, 33}

func (i FieldType) String() string {
	i -= 1
	if i < 0 || i >= FieldType(len(_FieldType_index)-1) {
		return "FieldType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _FieldType_name[_FieldType_index[i]:_FieldType_index[i+1]]
}
