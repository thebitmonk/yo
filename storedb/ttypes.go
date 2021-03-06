// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package storedb

import (
	"bytes"
	"fmt"
	"git-wip-us.apache.org/repos/asf/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type Key struct {
	Id      int64  `thrift:"id,1,required" json:"id"`
	Command []byte `thrift:"command,2,required" json:"command"`
	Email   []byte `thrift:"email,3" json:"email"`
}

func NewKey() *Key {
	return &Key{}
}

func (p *Key) GetId() int64 {
	return p.Id
}

func (p *Key) GetCommand() []byte {
	return p.Command
}

var Key_Email_DEFAULT []byte

func (p *Key) GetEmail() []byte {
	return p.Email
}
func (p *Key) IsSetEmail() bool {
	return p.Email != nil
}

func (p *Key) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Key) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Id = v
	}
	return nil
}

func (p *Key) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Command = v
	}
	return nil
}

func (p *Key) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Email = v
	}
	return nil
}

func (p *Key) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Key"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Key) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:id: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Id)); err != nil {
		return fmt.Errorf("%T.id (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:id: %s", p, err)
	}
	return err
}

func (p *Key) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("command", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:command: %s", p, err)
	}
	if err := oprot.WriteBinary(p.Command); err != nil {
		return fmt.Errorf("%T.command (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:command: %s", p, err)
	}
	return err
}

func (p *Key) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetEmail() {
		if err := oprot.WriteFieldBegin("email", thrift.STRING, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:email: %s", p, err)
		}
		if err := oprot.WriteBinary(p.Email); err != nil {
			return fmt.Errorf("%T.email (3) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:email: %s", p, err)
		}
	}
	return err
}

func (p *Key) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Key(%+v)", *p)
}

type Value struct {
	Command     []byte `thrift:"command,1,required" json:"command"`
	Options     []byte `thrift:"options,2,required" json:"options"`
	Description []byte `thrift:"description,3,required" json:"description"`
	UpVotes     int64  `thrift:"upVotes,4" json:"upVotes"`
	DownVotes   int64  `thrift:"downVotes,5" json:"downVotes"`
}

func NewValue() *Value {
	return &Value{}
}

func (p *Value) GetCommand() []byte {
	return p.Command
}

func (p *Value) GetOptions() []byte {
	return p.Options
}

func (p *Value) GetDescription() []byte {
	return p.Description
}

func (p *Value) GetUpVotes() int64 {
	return p.UpVotes
}

func (p *Value) GetDownVotes() int64 {
	return p.DownVotes
}
func (p *Value) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Value) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Command = v
	}
	return nil
}

func (p *Value) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Options = v
	}
	return nil
}

func (p *Value) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Description = v
	}
	return nil
}

func (p *Value) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 4: %s", err)
	} else {
		p.UpVotes = v
	}
	return nil
}

func (p *Value) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 5: %s", err)
	} else {
		p.DownVotes = v
	}
	return nil
}

func (p *Value) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Value"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Value) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("command", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:command: %s", p, err)
	}
	if err := oprot.WriteBinary(p.Command); err != nil {
		return fmt.Errorf("%T.command (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:command: %s", p, err)
	}
	return err
}

func (p *Value) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("options", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:options: %s", p, err)
	}
	if err := oprot.WriteBinary(p.Options); err != nil {
		return fmt.Errorf("%T.options (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:options: %s", p, err)
	}
	return err
}

func (p *Value) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("description", thrift.STRING, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:description: %s", p, err)
	}
	if err := oprot.WriteBinary(p.Description); err != nil {
		return fmt.Errorf("%T.description (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:description: %s", p, err)
	}
	return err
}

func (p *Value) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("upVotes", thrift.I64, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:upVotes: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.UpVotes)); err != nil {
		return fmt.Errorf("%T.upVotes (4) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:upVotes: %s", p, err)
	}
	return err
}

func (p *Value) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("downVotes", thrift.I64, 5); err != nil {
		return fmt.Errorf("%T write field begin error 5:downVotes: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.DownVotes)); err != nil {
		return fmt.Errorf("%T.downVotes (5) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 5:downVotes: %s", p, err)
	}
	return err
}

func (p *Value) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Value(%+v)", *p)
}

type Response struct {
	Status bool     `thrift:"status,1,required" json:"status"`
	Time   *float64 `thrift:"time,2" json:"time"`
	Data   []byte   `thrift:"data,3" json:"data"`
}

func NewResponse() *Response {
	return &Response{}
}

func (p *Response) GetStatus() bool {
	return p.Status
}

var Response_Time_DEFAULT float64

func (p *Response) GetTime() float64 {
	if !p.IsSetTime() {
		return Response_Time_DEFAULT
	}
	return *p.Time
}

var Response_Data_DEFAULT []byte

func (p *Response) GetData() []byte {
	return p.Data
}
func (p *Response) IsSetTime() bool {
	return p.Time != nil
}

func (p *Response) IsSetData() bool {
	return p.Data != nil
}

func (p *Response) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Response) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *Response) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Time = &v
	}
	return nil
}

func (p *Response) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Data = v
	}
	return nil
}

func (p *Response) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Response"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Response) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("status", thrift.BOOL, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:status: %s", p, err)
	}
	if err := oprot.WriteBool(bool(p.Status)); err != nil {
		return fmt.Errorf("%T.status (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:status: %s", p, err)
	}
	return err
}

func (p *Response) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetTime() {
		if err := oprot.WriteFieldBegin("time", thrift.DOUBLE, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:time: %s", p, err)
		}
		if err := oprot.WriteDouble(float64(*p.Time)); err != nil {
			return fmt.Errorf("%T.time (2) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:time: %s", p, err)
		}
	}
	return err
}

func (p *Response) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetData() {
		if err := oprot.WriteFieldBegin("data", thrift.STRING, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:data: %s", p, err)
		}
		if err := oprot.WriteBinary(p.Data); err != nil {
			return fmt.Errorf("%T.data (3) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:data: %s", p, err)
		}
	}
	return err
}

func (p *Response) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Response(%+v)", *p)
}

type MultiResponse struct {
	Status bool     `thrift:"status,1,required" json:"status"`
	Time   *float64 `thrift:"time,2" json:"time"`
	Data   [][]byte `thrift:"data,3" json:"data"`
}

func NewMultiResponse() *MultiResponse {
	return &MultiResponse{}
}

func (p *MultiResponse) GetStatus() bool {
	return p.Status
}

var MultiResponse_Time_DEFAULT float64

func (p *MultiResponse) GetTime() float64 {
	if !p.IsSetTime() {
		return MultiResponse_Time_DEFAULT
	}
	return *p.Time
}

var MultiResponse_Data_DEFAULT [][]byte

func (p *MultiResponse) GetData() [][]byte {
	return p.Data
}
func (p *MultiResponse) IsSetTime() bool {
	return p.Time != nil
}

func (p *MultiResponse) IsSetData() bool {
	return p.Data != nil
}

func (p *MultiResponse) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *MultiResponse) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Status = v
	}
	return nil
}

func (p *MultiResponse) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Time = &v
	}
	return nil
}

func (p *MultiResponse) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([][]byte, 0, size)
	p.Data = tSlice
	for i := 0; i < size; i++ {
		var _elem0 []byte
		if v, err := iprot.ReadBinary(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_elem0 = v
		}
		p.Data = append(p.Data, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *MultiResponse) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("MultiResponse"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *MultiResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("status", thrift.BOOL, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:status: %s", p, err)
	}
	if err := oprot.WriteBool(bool(p.Status)); err != nil {
		return fmt.Errorf("%T.status (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:status: %s", p, err)
	}
	return err
}

func (p *MultiResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetTime() {
		if err := oprot.WriteFieldBegin("time", thrift.DOUBLE, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:time: %s", p, err)
		}
		if err := oprot.WriteDouble(float64(*p.Time)); err != nil {
			return fmt.Errorf("%T.time (2) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:time: %s", p, err)
		}
	}
	return err
}

func (p *MultiResponse) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetData() {
		if err := oprot.WriteFieldBegin("data", thrift.LIST, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:data: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.Data)); err != nil {
			return fmt.Errorf("error writing list begin: %s", err)
		}
		for _, v := range p.Data {
			if err := oprot.WriteBinary(v); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p, err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:data: %s", p, err)
		}
	}
	return err
}

func (p *MultiResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MultiResponse(%+v)", *p)
}
