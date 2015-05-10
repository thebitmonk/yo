#include "service_common/StoreDBMaster.h"
#include <thrift/protocol/TBinaryProtocol.h>
#include <thrift/server/TSimpleServer.h>
#include <thrift/transport/TServerSocket.h>
#include <thrift/transport/TBufferTransports.h>
#include "db_interface.h"
#include "rocksdb.h"
#include <iostream>
using namespace ::apache::thrift;
using namespace ::apache::thrift::protocol;
using namespace ::apache::thrift::transport;
using namespace ::apache::thrift::server;

using namespace  ::storedb;

class StoreDBMasterHandler : virtual public StoreDBMasterIf {

 private:
   shared_ptr<DBInterface>db;

 public:
  StoreDBMasterHandler() {
    db = (shared_ptr<DBInterface>)(new RocksDB());
    bool status = db->openDB("/tmp/rocksdb");
    std::cout << "Calling master constructor" << std::endl;
  }

  void isOk(Response& _return) {
    std::cout << "Calling isOk" << std::endl;
    _return.status = "false";
    _return.time = 123; 
    _return.__set_data("hello world");
    std::cout << "Okay ends" << std::endl;
  }

  void add(Response& _return, const std::string& key, const std::string& value) {
    std::cout << "Calling Add" << std::endl;
    bool status = db->put(key, value); 
    if(status){
      std::cout << "Key value successfully added to DB " << std::endl;
      _return.__set_status(true);
      _return.__set_data("Successfully added");
    }else{
      std::cout << "Key value addition failed ... " << std::endl;
      _return.__set_status(false);
      _return.__set_data("Failure occurred");
    }
  }

  void search(MultiResponse& _return, const std::string& prefix) {
    std::cout << "Calling Search" << std::endl;
    std::vector<std::string>values;
    bool status = db->search(prefix, values);

    if(status){
      std::cout << "Successful prefix search in DB " << std::endl;
      _return.__set_status(true);
      _return.__set_data(values);
    }else{
      std::cout << "Prefix search failed failed ... " << std::endl;
      _return.__set_status(false);
    }
  }
};

int main(int argc, char **argv) {
  int port = 9090;
  ::boost::shared_ptr<StoreDBMasterHandler> handler(new StoreDBMasterHandler());
  ::boost::shared_ptr<TProcessor> processor(new StoreDBMasterProcessor(handler));
  ::boost::shared_ptr<TServerTransport> serverTransport(new TServerSocket(port));
  ::boost::shared_ptr<TTransportFactory> transportFactory(new TBufferedTransportFactory());
  ::boost::shared_ptr<TProtocolFactory> protocolFactory(new TBinaryProtocolFactory());

  TSimpleServer server(processor, serverTransport, transportFactory, protocolFactory);
  server.serve();
  return 0;
}

