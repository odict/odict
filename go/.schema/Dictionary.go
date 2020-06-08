// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package schema

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Dictionary struct {
	_tab flatbuffers.Table
}

func GetRootAsDictionary(buf []byte, offset flatbuffers.UOffsetT) *Dictionary {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Dictionary{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Dictionary) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Dictionary) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Dictionary) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Dictionary) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Dictionary) Entries(obj *Entry, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Dictionary) EntriesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DictionaryStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func DictionaryAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(id), 0)
}
func DictionaryAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0)
}
func DictionaryAddEntries(builder *flatbuffers.Builder, entries flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(entries), 0)
}
func DictionaryStartEntriesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DictionaryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
