package wayland

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

type WaylandMessage struct {
	Address uint32
	Opcode  uint16
	Length  uint16
	Data    []byte
}

type Field interface {
	AppendToBuf([]byte) []byte
}

type BytesField []byte

func (uf BytesField) AppendToBuf(buf []byte) []byte {
	buf = append(buf, uf...)
	return buf
}

type Uint16Field uint16

func NewUint16Field() *UintField {
	uf := UintField(0)
	return &uf
}
func (uf Uint16Field) AppendToBuf(buf []byte) []byte {
	buf = binary.LittleEndian.AppendUint16(buf, uint16(uf))
	return buf
}

type Uint64Field uint64

func NewUint64Field() *UintField {
	uf := UintField(0)
	return &uf
}
func (uf Uint64Field) AppendToBuf(buf []byte) []byte {
	buf = binary.LittleEndian.AppendUint64(buf, uint64(uf))
	return buf
}

type UintField uint32

func NewUintField() *UintField {
	uf := UintField(0)
	return &uf
}
func (uf UintField) AppendToBuf(buf []byte) []byte {
	buf = binary.LittleEndian.AppendUint32(buf, uint32(uf))
	return buf
}

type IntField int32

func NewIntField() *IntField {
	uf := IntField(0)
	return &uf
}
func (uf IntField) AppendToBuf(buf []byte) []byte {
	buf = binary.LittleEndian.AppendUint32(buf, uint32(uf))
	return buf
}

type FixedField float32

func (ff FixedField) AppendToBuf(buf []byte) []byte {

	u_d := float64(ff) + (3 << (51 - 8))
	u_i := int64(math.Float64bits(u_d))
	iarr := int32(u_i)
	buf = binary.LittleEndian.AppendUint32(buf, uint32(iarr))
	return buf
}

type StringField string

func NewStringField() *StringField {
	uf := StringField("")
	return &uf
}
func (sf StringField) AppendToBuf(buf []byte) []byte {

	// length opf the string + null terminator
	buf = binary.LittleEndian.AppendUint32(buf, uint32(len(sf)+1))
	buf = append(buf, string(sf)...)
	buf = append(buf, []byte{0}...)

	mod := (len(sf) + 1) % 4
	padding := 0
	if mod != 0 {
		off := 4 - mod
		padding += off
	}
	for i := 0; i < padding; i++ {
		buf = append(buf, []byte{0}...)
	}
	return buf
}

func ParsePacketStructure(buf []byte, fields ...Field) error {
	for _, field := range fields {
		switch f := field.(type) {
		case *UintField:
			if len(buf) >= 4 {
				uintfield := UintField(binary.LittleEndian.Uint32(buf[0:4]))
				*f = uintfield
				buf = buf[4:]
			} else {
				return fmt.Errorf("could not parse packet structure")
			}
		case *IntField:
			if len(buf) >= 4 {
				uintfield := IntField(binary.LittleEndian.Uint32(buf[0:4]))
				*f = uintfield
				buf = buf[4:]
			} else {
				return fmt.Errorf("could not parse packet structure")
			}
		case *StringField:
			if len(buf) >= 4 {
				strlen := binary.LittleEndian.Uint32(buf[0:4])
				idp := strlen
				mod := idp % 4
				ide := idp
				if mod != 0 {
					off := 4 - mod
					ide += off
				}
				if len(buf) >= int(4+ide) {
					nameBytes := bytes.Trim(buf[4:4+strlen], "\x00")
					strfield := StringField(string(nameBytes))
					*f = strfield
				}
				ptr := 4 + ide
				if len(buf) >= int(ptr) {
					buf = buf[ptr:]
				} else {
					return fmt.Errorf("could not parse packet structure")
				}
			} else {
				return fmt.Errorf("could not parse packet structure")
			}
		default:
			return fmt.Errorf("unimplemented field type: %v", field)
		}
	}
	return nil
}

// PackerBuilder allows piecemeal construction of a Wayland protocol packet
type PacketBulder struct {
	opcode uint16
	id     uint32
	fields []Field
}

func NewPacketBuilder(id uint32, opcode uint16) *PacketBulder {
	return &PacketBulder{opcode: opcode, id: id}
}

func (pb *PacketBulder) WithString(str string) *PacketBulder {
	pb.fields = append(pb.fields, StringField(str))
	return pb
}

func (pb *PacketBulder) WithUint(u uint32) *PacketBulder {
	pb.fields = append(pb.fields, UintField(u))
	return pb
}

func (pb *PacketBulder) WithFixed(u float32) *PacketBulder {
	pb.fields = append(pb.fields, FixedField(u))
	return pb
}

func (pb *PacketBulder) WithBytes(b []byte) *PacketBulder {
	pb.fields = append(pb.fields, BytesField(b))
	return pb
}

func (pb *PacketBulder) WithU16Array(arr []uint16) *PacketBulder {
	pb.fields = append(pb.fields, UintField(uint32(len(arr)*2)))
	for _, v := range arr {
		pb.fields = append(pb.fields, Uint16Field(v))
	}
	return pb
}

func (pb *PacketBulder) WithU32Array(arr []uint32) *PacketBulder {
	pb.fields = append(pb.fields, UintField(uint32(len(arr)*4)))
	for _, v := range arr {
		pb.fields = append(pb.fields, UintField(v))
	}
	return pb
}

func (pb *PacketBulder) WithU64Array(arr []uint64) *PacketBulder {
	pb.fields = append(pb.fields, UintField(uint64(len(arr)*8)))
	for _, v := range arr {
		pb.fields = append(pb.fields, Uint64Field(v))
	}
	return pb
}

func (pb *PacketBulder) Build() []byte {
	body := []byte{}
	for _, f := range pb.fields {
		body = f.AppendToBuf(body)
	}
	plen := 8 + len(body)
	buf := []byte{}
	buf = binary.LittleEndian.AppendUint32(buf, pb.id)
	buf = binary.LittleEndian.AppendUint16(buf, uint16(pb.opcode))
	buf = binary.LittleEndian.AppendUint16(buf, uint16(plen))
	buf = append(buf, body...)
	return buf
}
