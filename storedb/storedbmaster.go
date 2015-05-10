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

type StoreDBMaster interface {
	IsOk() (r *Response, err error)
	// Parameters:
	//  - Key
	//  - Value
	Add(key []byte, value []byte) (r *Response, err error)
	// Parameters:
	//  - Prefix
	Search(prefix []byte) (r *MultiResponse, err error)
}

type StoreDBMasterClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewStoreDBMasterClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *StoreDBMasterClient {
	return &StoreDBMasterClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewStoreDBMasterClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *StoreDBMasterClient {
	return &StoreDBMasterClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

func (p *StoreDBMasterClient) IsOk() (r *Response, err error) {
	if err = p.sendIsOk(); err != nil {
		return
	}
	return p.recvIsOk()
}

func (p *StoreDBMasterClient) sendIsOk() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("isOk", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := IsOkArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *StoreDBMasterClient) recvIsOk() (value *Response, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error1 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error2 error
		error2, err = error1.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error2
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "isOk failed: out of sequence response")
		return
	}
	result := IsOkResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Key
//  - Value
func (p *StoreDBMasterClient) Add(key []byte, value []byte) (r *Response, err error) {
	if err = p.sendAdd(key, value); err != nil {
		return
	}
	return p.recvAdd()
}

func (p *StoreDBMasterClient) sendAdd(key []byte, value []byte) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("add", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := AddArgs{
		Key:   key,
		Value: value,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *StoreDBMasterClient) recvAdd() (value *Response, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error3 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error4 error
		error4, err = error3.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error4
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "add failed: out of sequence response")
		return
	}
	result := AddResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Prefix
func (p *StoreDBMasterClient) Search(prefix []byte) (r *MultiResponse, err error) {
	if err = p.sendSearch(prefix); err != nil {
		return
	}
	return p.recvSearch()
}

func (p *StoreDBMasterClient) sendSearch(prefix []byte) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("search", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := SearchArgs{
		Prefix: prefix,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *StoreDBMasterClient) recvSearch() (value *MultiResponse, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error5 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error6 error
		error6, err = error5.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error6
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "search failed: out of sequence response")
		return
	}
	result := SearchResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type StoreDBMasterProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      StoreDBMaster
}

func (p *StoreDBMasterProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *StoreDBMasterProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *StoreDBMasterProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewStoreDBMasterProcessor(handler StoreDBMaster) *StoreDBMasterProcessor {

	self7 := &StoreDBMasterProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self7.processorMap["isOk"] = &storeDBMasterProcessorIsOk{handler: handler}
	self7.processorMap["add"] = &storeDBMasterProcessorAdd{handler: handler}
	self7.processorMap["search"] = &storeDBMasterProcessorSearch{handler: handler}
	return self7
}

func (p *StoreDBMasterProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x8 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x8.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x8

}

type storeDBMasterProcessorIsOk struct {
	handler StoreDBMaster
}

func (p *storeDBMasterProcessorIsOk) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := IsOkArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("isOk", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := IsOkResult{}
	var retval *Response
	var err2 error
	if retval, err2 = p.handler.IsOk(); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing isOk: "+err2.Error())
		oprot.WriteMessageBegin("isOk", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("isOk", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type storeDBMasterProcessorAdd struct {
	handler StoreDBMaster
}

func (p *storeDBMasterProcessorAdd) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AddArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("add", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AddResult{}
	var retval *Response
	var err2 error
	if retval, err2 = p.handler.Add(args.Key, args.Value); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing add: "+err2.Error())
		oprot.WriteMessageBegin("add", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("add", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type storeDBMasterProcessorSearch struct {
	handler StoreDBMaster
}

func (p *storeDBMasterProcessorSearch) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SearchArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("search", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := SearchResult{}
	var retval *MultiResponse
	var err2 error
	if retval, err2 = p.handler.Search(args.Prefix); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing search: "+err2.Error())
		oprot.WriteMessageBegin("search", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("search", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type IsOkArgs struct {
}

func NewIsOkArgs() *IsOkArgs {
	return &IsOkArgs{}
}

func (p *IsOkArgs) Read(iprot thrift.TProtocol) error {
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
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
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

func (p *IsOkArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("isOk_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *IsOkArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IsOkArgs(%+v)", *p)
}

type IsOkResult struct {
	Success *Response `thrift:"success,0" json:"success"`
}

func NewIsOkResult() *IsOkResult {
	return &IsOkResult{}
}

var IsOkResult_Success_DEFAULT *Response

func (p *IsOkResult) GetSuccess() *Response {
	if !p.IsSetSuccess() {
		return IsOkResult_Success_DEFAULT
	}
	return p.Success
}
func (p *IsOkResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *IsOkResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.ReadField0(iprot); err != nil {
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

func (p *IsOkResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &Response{}
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success, err)
	}
	return nil
}

func (p *IsOkResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("isOk_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *IsOkResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *IsOkResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IsOkResult(%+v)", *p)
}

type AddArgs struct {
	Key   []byte `thrift:"key,1" json:"key"`
	Value []byte `thrift:"value,2" json:"value"`
}

func NewAddArgs() *AddArgs {
	return &AddArgs{}
}

func (p *AddArgs) GetKey() []byte {
	return p.Key
}

func (p *AddArgs) GetValue() []byte {
	return p.Value
}
func (p *AddArgs) Read(iprot thrift.TProtocol) error {
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

func (p *AddArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Key = v
	}
	return nil
}

func (p *AddArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *AddArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("add_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
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

func (p *AddArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:key: %s", p, err)
	}
	if err := oprot.WriteBinary(p.Key); err != nil {
		return fmt.Errorf("%T.key (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:key: %s", p, err)
	}
	return err
}

func (p *AddArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:value: %s", p, err)
	}
	if err := oprot.WriteBinary(p.Value); err != nil {
		return fmt.Errorf("%T.value (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:value: %s", p, err)
	}
	return err
}

func (p *AddArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddArgs(%+v)", *p)
}

type AddResult struct {
	Success *Response `thrift:"success,0" json:"success"`
}

func NewAddResult() *AddResult {
	return &AddResult{}
}

var AddResult_Success_DEFAULT *Response

func (p *AddResult) GetSuccess() *Response {
	if !p.IsSetSuccess() {
		return AddResult_Success_DEFAULT
	}
	return p.Success
}
func (p *AddResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.ReadField0(iprot); err != nil {
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

func (p *AddResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &Response{}
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success, err)
	}
	return nil
}

func (p *AddResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("add_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *AddResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *AddResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddResult(%+v)", *p)
}

type SearchArgs struct {
	Prefix []byte `thrift:"prefix,1" json:"prefix"`
}

func NewSearchArgs() *SearchArgs {
	return &SearchArgs{}
}

func (p *SearchArgs) GetPrefix() []byte {
	return p.Prefix
}
func (p *SearchArgs) Read(iprot thrift.TProtocol) error {
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

func (p *SearchArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Prefix = v
	}
	return nil
}

func (p *SearchArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("search_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
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

func (p *SearchArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("prefix", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:prefix: %s", p, err)
	}
	if err := oprot.WriteBinary(p.Prefix); err != nil {
		return fmt.Errorf("%T.prefix (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:prefix: %s", p, err)
	}
	return err
}

func (p *SearchArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SearchArgs(%+v)", *p)
}

type SearchResult struct {
	Success *MultiResponse `thrift:"success,0" json:"success"`
}

func NewSearchResult() *SearchResult {
	return &SearchResult{}
}

var SearchResult_Success_DEFAULT *MultiResponse

func (p *SearchResult) GetSuccess() *MultiResponse {
	if !p.IsSetSuccess() {
		return SearchResult_Success_DEFAULT
	}
	return p.Success
}
func (p *SearchResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.ReadField0(iprot); err != nil {
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

func (p *SearchResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &MultiResponse{}
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success, err)
	}
	return nil
}

func (p *SearchResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("search_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *SearchResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *SearchResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SearchResult(%+v)", *p)
}
