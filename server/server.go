package main

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"git-wip-us.apache.org/repos/asf/thrift.git/lib/go/thrift"
	"github.com/gorilla/mux"
	"github.com/thebitmonk/storedb"
	"github.com/thebitmonk/utils"
	"log"
	"net/http"
)

const (
	_email       string = "email"
	_password    string = "password"
	_endpoint    string = "endpoint"
	_command     string = "command"
	_options     string = "options"
	_description string = "description"
	_prefix      string = "prefix"
)

var tclient *storedb.StoreDBMasterClient

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	log.Print("Home handler triggered")
}

func generateId() int64 {
	var n int64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return n
}

func AddHandler(rw http.ResponseWriter, r *http.Request) {
	log.Print("New Add handler triggered")

	_ = r.ParseForm()
	id := generateId()
	email := r.Form[_email]
	command := r.Form[_command]
	options := r.Form[_options]
	description := r.Form[_description]

	//key := storedb.Key{Id: id, Command: []byte(command[0]), Email: []byte(email[0])}
	value := storedb.Value{Command: []byte(command[0]), Options: []byte(options[0]), Description: []byte(description[0]), UpVotes: 0, DownVotes: 0}

	//serkey, _ := thrift_utils.SerializeKey(&key)
	serval, _ := thrift_utils.SerializeValue(&value)

	serkey := fmt.Sprintf("%s::id::%d::email::%s", command[0], id, email[0])
	fmt.Printf("Serialized key %s\n\n", serkey)

	resp := &storedb.Response{}
	resp, _ = tclient.Add([]byte(serkey), []byte(serval))
	fmt.Println(string((*resp).Data))
	rw.Write((*resp).Data)
}

func SearchHandler(rw http.ResponseWriter, r *http.Request) {
	log.Print("Search handler triggered")

	_ = r.ParseForm()
	prefix := r.Form[_prefix]
	resp := &storedb.MultiResponse{}
	resp, _ = tclient.Search([]byte(prefix[0]))
	fmt.Println((*resp).Status)

	var subcommands = make([]storedb.Value, len((*resp).Data))
	for i, serval := range (*resp).Data {
		_value := &storedb.Value{}
		thrift_utils.DeserializeValue(string(serval), _value)
		subcommands[i] = *_value
		fmt.Println(string(_value.Options))
		fmt.Println(string(_value.Description))
	}
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(subcommands); err != nil {
		panic(err)
	}
}

func UpvoteHandler(rw http.ResponseWriter, r *http.Request) {
	log.Print("Upvote handler triggered")
}

func DownvoteHandler(rw http.ResponseWriter, r *http.Request) {
	log.Print("Downvote handler triggered")
}

// TODO: thrift is not thread safe. Make a client pool and http interface must
// call via this thrift client pool
func getClient() (c *storedb.StoreDBMasterClient) {

	var addr = "127.0.0.1:9090"
	socket, _ := thrift.NewTSocket(addr)

	transportFactory := thrift.NewTBufferedTransportFactory(1024)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	var transport thrift.TTransport

	transport = transportFactory.GetTransport(socket)
	transport.Open()
	tclient := storedb.NewStoreDBMasterClientFactory(transport, protocolFactory)
	return tclient

}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	tclient = getClient()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/add", AddHandler).Methods("POST")
	r.HandleFunc("/search", SearchHandler).Methods("POST")
	r.HandleFunc("/upvote", UpvoteHandler).Methods("POST")
	r.HandleFunc("/downvote", DownvoteHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}
