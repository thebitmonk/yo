package thrift_utils

import (
	"git-wip-us.apache.org/repos/asf/thrift.git/lib/go/thrift"
	"github.com/thebitmonk/storedb"
)

// TODO: If possible do via interface {}
// If not then try reflect and type check approach

// Serialize & Deserialize
// ------------------------
//  serkey, _ := thrift_utils.SerializeKey(&key)
// _key := &storedb.Key{}
// _ = thrift_utils.DeserializeKey(data, _key)
// log.Print(_key.Id)
// log.Print(string(_key.Email))

//  serval, _ := thrift_utils.SerializeValue(&value)
// _value := &storedb.Value{}
// _ = thrift_utils.DeserializeValue(data, _value)
// log.Print(string(_value.Options))
// log.Print(string(_value.Description))

// Binary deserializer
func DeserializeKey(data string, T *storedb.Key) (err error) {
	t := thrift.NewTDeserializer()
	protocol := thrift.NewTBinaryProtocolFactoryDefault()
	t.Protocol = protocol.GetProtocol(t.Transport)
	err = t.ReadString(T, data)
	return err
}

// Binary serializer
func SerializeKey(T *storedb.Key) (data string, err error) {
	t := thrift.NewTSerializer()
	protocol := thrift.NewTBinaryProtocolFactoryDefault()
	t.Protocol = protocol.GetProtocol(t.Transport)
	return t.WriteString(T)
}

// Binary deserializer
func DeserializeValue(data string, T *storedb.Value) (err error) {
	t := thrift.NewTDeserializer()
	protocol := thrift.NewTBinaryProtocolFactoryDefault()
	t.Protocol = protocol.GetProtocol(t.Transport)
	err = t.ReadString(T, data)
	return err
}

// Binary serializer
func SerializeValue(T *storedb.Value) (data string, err error) {
	t := thrift.NewTSerializer()
	protocol := thrift.NewTBinaryProtocolFactoryDefault()
	t.Protocol = protocol.GetProtocol(t.Transport)
	return t.WriteString(T)
}
