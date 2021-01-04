// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package schema

import (
	"strconv"

	flatbuffers "github.com/google/flatbuffers/go"
)

type POS int8

const (
	POSverb     POS = 0
	POSnoun     POS = 1
	POSadj      POS = 2
	POSpronoun  POS = 3
	POSadv      POS = 4
	POSprep     POS = 5
	POSconj     POS = 6
	POSintj     POS = 7
	POSprefix   POS = 8
	POSsuffix   POS = 9
	POSparticle POS = 10
	POSarticle  POS = 11
	POSunknown  POS = 12
)

var EnumNamesPOS = map[POS]string{
	POSverb:     "verb",
	POSnoun:     "noun",
	POSadj:      "adj",
	POSpronoun:  "pronoun",
	POSadv:      "adv",
	POSprep:     "prep",
	POSconj:     "conj",
	POSintj:     "intj",
	POSprefix:   "prefix",
	POSsuffix:   "suffix",
	POSparticle: "particle",
	POSarticle:  "article",
	POSunknown:  "unknown",
}

var EnumValuesPOS = map[string]POS{
	"verb":     POSverb,
	"noun":     POSnoun,
	"adj":      POSadj,
	"pronoun":  POSpronoun,
	"adv":      POSadv,
	"prep":     POSprep,
	"conj":     POSconj,
	"intj":     POSintj,
	"prefix":   POSprefix,
	"suffix":   POSsuffix,
	"particle": POSparticle,
	"article":  POSarticle,
	"unknown":  POSunknown,
}

func (v POS) String() string {
	if s, ok := EnumNamesPOS[v]; ok {
		return s
	}
	return "POS(" + strconv.FormatInt(int64(v), 10) + ")"
}

type Etymology struct {
	_tab flatbuffers.Table
}

func GetRootAsEtymology(buf []byte, offset flatbuffers.UOffsetT) *Etymology {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Etymology{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Etymology) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Etymology) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Etymology) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Etymology) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Etymology) Usages(obj *Usage, j int) bool {
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

func (rcv *Etymology) UsagesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func EtymologyStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func EtymologyAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(id), 0)
}
func EtymologyAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(description), 0)
}
func EtymologyAddUsages(builder *flatbuffers.Builder, usages flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(usages), 0)
}
func EtymologyStartUsagesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EtymologyEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Group struct {
	_tab flatbuffers.Table
}

func GetRootAsGroup(buf []byte, offset flatbuffers.UOffsetT) *Group {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Group{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Group) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Group) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Group) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Group) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Group) Definitions(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Group) DefinitionsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func GroupStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func GroupAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(id), 0)
}
func GroupAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(description), 0)
}
func GroupAddDefinitions(builder *flatbuffers.Builder, definitions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(definitions), 0)
}
func GroupStartDefinitionsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func GroupEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Usage struct {
	_tab flatbuffers.Table
}

func GetRootAsUsage(buf []byte, offset flatbuffers.UOffsetT) *Usage {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Usage{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Usage) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Usage) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Usage) Pos() POS {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return POS(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Usage) MutatePos(n POS) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func (rcv *Usage) Definitions(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Usage) DefinitionsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Usage) Groups(obj *Group, j int) bool {
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

func (rcv *Usage) GroupsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func UsageStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func UsageAddPos(builder *flatbuffers.Builder, pos POS) {
	builder.PrependInt8Slot(0, int8(pos), 0)
}
func UsageAddDefinitions(builder *flatbuffers.Builder, definitions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(definitions), 0)
}
func UsageStartDefinitionsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func UsageAddGroups(builder *flatbuffers.Builder, groups flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(groups), 0)
}
func UsageStartGroupsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func UsageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Entry struct {
	_tab flatbuffers.Table
}

func GetRootAsEntry(buf []byte, offset flatbuffers.UOffsetT) *Entry {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Entry{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Entry) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Entry) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Entry) Term() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Entry) Etymologies(obj *Etymology, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Entry) EtymologiesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func EntryStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func EntryAddTerm(builder *flatbuffers.Builder, term flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(term), 0)
}
func EntryAddEtymologies(builder *flatbuffers.Builder, etymologies flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(etymologies), 0)
}
func EntryStartEtymologiesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EntryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
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
