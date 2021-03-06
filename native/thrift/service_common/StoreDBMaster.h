/**
 * Autogenerated by Thrift Compiler (0.9.2)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#ifndef StoreDBMaster_H
#define StoreDBMaster_H

#include <thrift/TDispatchProcessor.h>
#include "StoreDB_types.h"

namespace storedb {

class StoreDBMasterIf {
 public:
  virtual ~StoreDBMasterIf() {}
  virtual void isOk(Response& _return) = 0;
  virtual void add(Response& _return, const std::string& key, const std::string& value) = 0;
  virtual void search(MultiResponse& _return, const std::string& prefix) = 0;
};

class StoreDBMasterIfFactory {
 public:
  typedef StoreDBMasterIf Handler;

  virtual ~StoreDBMasterIfFactory() {}

  virtual StoreDBMasterIf* getHandler(const ::apache::thrift::TConnectionInfo& connInfo) = 0;
  virtual void releaseHandler(StoreDBMasterIf* /* handler */) = 0;
};

class StoreDBMasterIfSingletonFactory : virtual public StoreDBMasterIfFactory {
 public:
  StoreDBMasterIfSingletonFactory(const boost::shared_ptr<StoreDBMasterIf>& iface) : iface_(iface) {}
  virtual ~StoreDBMasterIfSingletonFactory() {}

  virtual StoreDBMasterIf* getHandler(const ::apache::thrift::TConnectionInfo&) {
    return iface_.get();
  }
  virtual void releaseHandler(StoreDBMasterIf* /* handler */) {}

 protected:
  boost::shared_ptr<StoreDBMasterIf> iface_;
};

class StoreDBMasterNull : virtual public StoreDBMasterIf {
 public:
  virtual ~StoreDBMasterNull() {}
  void isOk(Response& /* _return */) {
    return;
  }
  void add(Response& /* _return */, const std::string& /* key */, const std::string& /* value */) {
    return;
  }
  void search(MultiResponse& /* _return */, const std::string& /* prefix */) {
    return;
  }
};


class StoreDBMaster_isOk_args {
 public:

  static const char* ascii_fingerprint; // = "99914B932BD37A50B983C5E7C90AE93B";
  static const uint8_t binary_fingerprint[16]; // = {0x99,0x91,0x4B,0x93,0x2B,0xD3,0x7A,0x50,0xB9,0x83,0xC5,0xE7,0xC9,0x0A,0xE9,0x3B};

  StoreDBMaster_isOk_args(const StoreDBMaster_isOk_args&);
  StoreDBMaster_isOk_args& operator=(const StoreDBMaster_isOk_args&);
  StoreDBMaster_isOk_args() {
  }

  virtual ~StoreDBMaster_isOk_args() throw();

  bool operator == (const StoreDBMaster_isOk_args & /* rhs */) const
  {
    return true;
  }
  bool operator != (const StoreDBMaster_isOk_args &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const StoreDBMaster_isOk_args & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  friend std::ostream& operator<<(std::ostream& out, const StoreDBMaster_isOk_args& obj);
};


class StoreDBMaster_isOk_pargs {
 public:

  static const char* ascii_fingerprint; // = "99914B932BD37A50B983C5E7C90AE93B";
  static const uint8_t binary_fingerprint[16]; // = {0x99,0x91,0x4B,0x93,0x2B,0xD3,0x7A,0x50,0xB9,0x83,0xC5,0xE7,0xC9,0x0A,0xE9,0x3B};


  virtual ~StoreDBMaster_isOk_pargs() throw();

  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  friend std::ostream& operator<<(std::ostream& out, const StoreDBMaster_isOk_pargs& obj);
};

typedef struct _StoreDBMaster_isOk_result__isset {
  _StoreDBMaster_isOk_result__isset() : success(false) {}
  bool success :1;
} _StoreDBMaster_isOk_result__isset;

class StoreDBMaster_isOk_result {
 public:

  static const char* ascii_fingerprint; // = "1E3FEAF846610A0C7D48EF3D6B340A7F";
  static const uint8_t binary_fingerprint[16]; // = {0x1E,0x3F,0xEA,0xF8,0x46,0x61,0x0A,0x0C,0x7D,0x48,0xEF,0x3D,0x6B,0x34,0x0A,0x7F};

  StoreDBMaster_isOk_result(const StoreDBMaster_isOk_result&);
  StoreDBMaster_isOk_result& operator=(const StoreDBMaster_isOk_result&);
  StoreDBMaster_isOk_result() {
  }

  virtual ~StoreDBMaster_isOk_result() throw();
  Response success;

  _StoreDBMaster_isOk_result__isset __isset;

  void __set_success(const Response& val);

  bool operator == (const StoreDBMaster_isOk_result & rhs) const
  {
    if (!(success == rhs.success))
      return false;
    return true;
  }
  bool operator != (const StoreDBMaster_isOk_result &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const StoreDBMaster_isOk_result & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  friend std::ostream& operator<<(std::ostream& out, const StoreDBMaster_isOk_result& obj);
};

typedef struct _StoreDBMaster_isOk_presult__isset {
  _StoreDBMaster_isOk_presult__isset() : success(false) {}
  bool success :1;
} _StoreDBMaster_isOk_presult__isset;

class StoreDBMaster_isOk_presult {
 public:

  static const char* ascii_fingerprint; // = "1E3FEAF846610A0C7D48EF3D6B340A7F";
  static const uint8_t binary_fingerprint[16]; // = {0x1E,0x3F,0xEA,0xF8,0x46,0x61,0x0A,0x0C,0x7D,0x48,0xEF,0x3D,0x6B,0x34,0x0A,0x7F};


  virtual ~StoreDBMaster_isOk_presult() throw();
  Response* success;

  _StoreDBMaster_isOk_presult__isset __isset;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);

  friend std::ostream& operator<<(std::ostream& out, const StoreDBMaster_isOk_presult& obj);
};

typedef struct _StoreDBMaster_add_args__isset {
  _StoreDBMaster_add_args__isset() : key(false), value(false) {}
  bool key :1;
  bool value :1;
} _StoreDBMaster_add_args__isset;

class StoreDBMaster_add_args {
 public:

  static const char* ascii_fingerprint; // = "07A9615F837F7D0A952B595DD3020972";
  static const uint8_t binary_fingerprint[16]; // = {0x07,0xA9,0x61,0x5F,0x83,0x7F,0x7D,0x0A,0x95,0x2B,0x59,0x5D,0xD3,0x02,0x09,0x72};

  StoreDBMaster_add_args(const StoreDBMaster_add_args&);
  StoreDBMaster_add_args& operator=(const StoreDBMaster_add_args&);
  StoreDBMaster_add_args() : key(), value() {
  }

  virtual ~StoreDBMaster_add_args() throw();
  std::string key;
  std::string value;

  _StoreDBMaster_add_args__isset __isset;

  void __set_key(const std::string& val);

  void __set_value(const std::string& val);

  bool operator == (const StoreDBMaster_add_args & rhs) const
  {
    if (!(key == rhs.key))
      return false;
    if (!(value == rhs.value))
      return false;
    return true;
  }
  bool operator != (const StoreDBMaster_add_args &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const StoreDBMaster_add_args & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  friend std::ostream& operator<<(std::ostream& out, const StoreDBMaster_add_args& obj);
};


class StoreDBMaster_add_pargs {
 public:

  static const char* ascii_fingerprint; // = "07A9615F837F7D0A952B595DD3020972";
  static const uint8_t binary_fingerprint[16]; // = {0x07,0xA9,0x61,0x5F,0x83,0x7F,0x7D,0x0A,0x95,0x2B,0x59,0x5D,0xD3,0x02,0x09,0x72};


  virtual ~StoreDBMaster_add_pargs() throw();
  const std::string* key;
  const std::string* value;

  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  friend std::ostream& operator<<(std::ostream& out, const StoreDBMaster_add_pargs& obj);
};

typedef struct _StoreDBMaster_add_result__isset {
  _StoreDBMaster_add_result__isset() : success(false) {}
  bool success :1;
} _StoreDBMaster_add_result__isset;

class StoreDBMaster_add_result {
 public:

  static const char* ascii_fingerprint; // = "1E3FEAF846610A0C7D48EF3D6B340A7F";
  static const uint8_t binary_fingerprint[16]; // = {0x1E,0x3F,0xEA,0xF8,0x46,0x61,0x0A,0x0C,0x7D,0x48,0xEF,0x3D,0x6B,0x34,0x0A,0x7F};

  StoreDBMaster_add_result(const StoreDBMaster_add_result&);
  StoreDBMaster_add_result& operator=(const StoreDBMaster_add_result&);
  StoreDBMaster_add_result() {
  }

  virtual ~StoreDBMaster_add_result() throw();
  Response success;

  _StoreDBMaster_add_result__isset __isset;

  void __set_success(const Response& val);

  bool operator == (const StoreDBMaster_add_result & rhs) const
  {
    if (!(success == rhs.success))
      return false;
    return true;
  }
  bool operator != (const StoreDBMaster_add_result &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const StoreDBMaster_add_result & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  friend std::ostream& operator<<(std::ostream& out, const StoreDBMaster_add_result& obj);
};

typedef struct _StoreDBMaster_add_presult__isset {
  _StoreDBMaster_add_presult__isset() : success(false) {}
  bool success :1;
} _StoreDBMaster_add_presult__isset;

class StoreDBMaster_add_presult {
 public:

  static const char* ascii_fingerprint; // = "1E3FEAF846610A0C7D48EF3D6B340A7F";
  static const uint8_t binary_fingerprint[16]; // = {0x1E,0x3F,0xEA,0xF8,0x46,0x61,0x0A,0x0C,0x7D,0x48,0xEF,0x3D,0x6B,0x34,0x0A,0x7F};


  virtual ~StoreDBMaster_add_presult() throw();
  Response* success;

  _StoreDBMaster_add_presult__isset __isset;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);

  friend std::ostream& operator<<(std::ostream& out, const StoreDBMaster_add_presult& obj);
};

typedef struct _StoreDBMaster_search_args__isset {
  _StoreDBMaster_search_args__isset() : prefix(false) {}
  bool prefix :1;
} _StoreDBMaster_search_args__isset;

class StoreDBMaster_search_args {
 public:

  static const char* ascii_fingerprint; // = "EFB929595D312AC8F305D5A794CFEDA1";
  static const uint8_t binary_fingerprint[16]; // = {0xEF,0xB9,0x29,0x59,0x5D,0x31,0x2A,0xC8,0xF3,0x05,0xD5,0xA7,0x94,0xCF,0xED,0xA1};

  StoreDBMaster_search_args(const StoreDBMaster_search_args&);
  StoreDBMaster_search_args& operator=(const StoreDBMaster_search_args&);
  StoreDBMaster_search_args() : prefix() {
  }

  virtual ~StoreDBMaster_search_args() throw();
  std::string prefix;

  _StoreDBMaster_search_args__isset __isset;

  void __set_prefix(const std::string& val);

  bool operator == (const StoreDBMaster_search_args & rhs) const
  {
    if (!(prefix == rhs.prefix))
      return false;
    return true;
  }
  bool operator != (const StoreDBMaster_search_args &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const StoreDBMaster_search_args & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  friend std::ostream& operator<<(std::ostream& out, const StoreDBMaster_search_args& obj);
};


class StoreDBMaster_search_pargs {
 public:

  static const char* ascii_fingerprint; // = "EFB929595D312AC8F305D5A794CFEDA1";
  static const uint8_t binary_fingerprint[16]; // = {0xEF,0xB9,0x29,0x59,0x5D,0x31,0x2A,0xC8,0xF3,0x05,0xD5,0xA7,0x94,0xCF,0xED,0xA1};


  virtual ~StoreDBMaster_search_pargs() throw();
  const std::string* prefix;

  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  friend std::ostream& operator<<(std::ostream& out, const StoreDBMaster_search_pargs& obj);
};

typedef struct _StoreDBMaster_search_result__isset {
  _StoreDBMaster_search_result__isset() : success(false) {}
  bool success :1;
} _StoreDBMaster_search_result__isset;

class StoreDBMaster_search_result {
 public:

  static const char* ascii_fingerprint; // = "537259DD471FCB08C8C2FDE2D8E1FA09";
  static const uint8_t binary_fingerprint[16]; // = {0x53,0x72,0x59,0xDD,0x47,0x1F,0xCB,0x08,0xC8,0xC2,0xFD,0xE2,0xD8,0xE1,0xFA,0x09};

  StoreDBMaster_search_result(const StoreDBMaster_search_result&);
  StoreDBMaster_search_result& operator=(const StoreDBMaster_search_result&);
  StoreDBMaster_search_result() {
  }

  virtual ~StoreDBMaster_search_result() throw();
  MultiResponse success;

  _StoreDBMaster_search_result__isset __isset;

  void __set_success(const MultiResponse& val);

  bool operator == (const StoreDBMaster_search_result & rhs) const
  {
    if (!(success == rhs.success))
      return false;
    return true;
  }
  bool operator != (const StoreDBMaster_search_result &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const StoreDBMaster_search_result & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  friend std::ostream& operator<<(std::ostream& out, const StoreDBMaster_search_result& obj);
};

typedef struct _StoreDBMaster_search_presult__isset {
  _StoreDBMaster_search_presult__isset() : success(false) {}
  bool success :1;
} _StoreDBMaster_search_presult__isset;

class StoreDBMaster_search_presult {
 public:

  static const char* ascii_fingerprint; // = "537259DD471FCB08C8C2FDE2D8E1FA09";
  static const uint8_t binary_fingerprint[16]; // = {0x53,0x72,0x59,0xDD,0x47,0x1F,0xCB,0x08,0xC8,0xC2,0xFD,0xE2,0xD8,0xE1,0xFA,0x09};


  virtual ~StoreDBMaster_search_presult() throw();
  MultiResponse* success;

  _StoreDBMaster_search_presult__isset __isset;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);

  friend std::ostream& operator<<(std::ostream& out, const StoreDBMaster_search_presult& obj);
};

class StoreDBMasterClient : virtual public StoreDBMasterIf {
 public:
  StoreDBMasterClient(boost::shared_ptr< ::apache::thrift::protocol::TProtocol> prot) {
    setProtocol(prot);
  }
  StoreDBMasterClient(boost::shared_ptr< ::apache::thrift::protocol::TProtocol> iprot, boost::shared_ptr< ::apache::thrift::protocol::TProtocol> oprot) {
    setProtocol(iprot,oprot);
  }
 private:
  void setProtocol(boost::shared_ptr< ::apache::thrift::protocol::TProtocol> prot) {
  setProtocol(prot,prot);
  }
  void setProtocol(boost::shared_ptr< ::apache::thrift::protocol::TProtocol> iprot, boost::shared_ptr< ::apache::thrift::protocol::TProtocol> oprot) {
    piprot_=iprot;
    poprot_=oprot;
    iprot_ = iprot.get();
    oprot_ = oprot.get();
  }
 public:
  boost::shared_ptr< ::apache::thrift::protocol::TProtocol> getInputProtocol() {
    return piprot_;
  }
  boost::shared_ptr< ::apache::thrift::protocol::TProtocol> getOutputProtocol() {
    return poprot_;
  }
  void isOk(Response& _return);
  void send_isOk();
  void recv_isOk(Response& _return);
  void add(Response& _return, const std::string& key, const std::string& value);
  void send_add(const std::string& key, const std::string& value);
  void recv_add(Response& _return);
  void search(MultiResponse& _return, const std::string& prefix);
  void send_search(const std::string& prefix);
  void recv_search(MultiResponse& _return);
 protected:
  boost::shared_ptr< ::apache::thrift::protocol::TProtocol> piprot_;
  boost::shared_ptr< ::apache::thrift::protocol::TProtocol> poprot_;
  ::apache::thrift::protocol::TProtocol* iprot_;
  ::apache::thrift::protocol::TProtocol* oprot_;
};

class StoreDBMasterProcessor : public ::apache::thrift::TDispatchProcessor {
 protected:
  boost::shared_ptr<StoreDBMasterIf> iface_;
  virtual bool dispatchCall(::apache::thrift::protocol::TProtocol* iprot, ::apache::thrift::protocol::TProtocol* oprot, const std::string& fname, int32_t seqid, void* callContext);
 private:
  typedef  void (StoreDBMasterProcessor::*ProcessFunction)(int32_t, ::apache::thrift::protocol::TProtocol*, ::apache::thrift::protocol::TProtocol*, void*);
  typedef std::map<std::string, ProcessFunction> ProcessMap;
  ProcessMap processMap_;
  void process_isOk(int32_t seqid, ::apache::thrift::protocol::TProtocol* iprot, ::apache::thrift::protocol::TProtocol* oprot, void* callContext);
  void process_add(int32_t seqid, ::apache::thrift::protocol::TProtocol* iprot, ::apache::thrift::protocol::TProtocol* oprot, void* callContext);
  void process_search(int32_t seqid, ::apache::thrift::protocol::TProtocol* iprot, ::apache::thrift::protocol::TProtocol* oprot, void* callContext);
 public:
  StoreDBMasterProcessor(boost::shared_ptr<StoreDBMasterIf> iface) :
    iface_(iface) {
    processMap_["isOk"] = &StoreDBMasterProcessor::process_isOk;
    processMap_["add"] = &StoreDBMasterProcessor::process_add;
    processMap_["search"] = &StoreDBMasterProcessor::process_search;
  }

  virtual ~StoreDBMasterProcessor() {}
};

class StoreDBMasterProcessorFactory : public ::apache::thrift::TProcessorFactory {
 public:
  StoreDBMasterProcessorFactory(const ::boost::shared_ptr< StoreDBMasterIfFactory >& handlerFactory) :
      handlerFactory_(handlerFactory) {}

  ::boost::shared_ptr< ::apache::thrift::TProcessor > getProcessor(const ::apache::thrift::TConnectionInfo& connInfo);

 protected:
  ::boost::shared_ptr< StoreDBMasterIfFactory > handlerFactory_;
};

class StoreDBMasterMultiface : virtual public StoreDBMasterIf {
 public:
  StoreDBMasterMultiface(std::vector<boost::shared_ptr<StoreDBMasterIf> >& ifaces) : ifaces_(ifaces) {
  }
  virtual ~StoreDBMasterMultiface() {}
 protected:
  std::vector<boost::shared_ptr<StoreDBMasterIf> > ifaces_;
  StoreDBMasterMultiface() {}
  void add(boost::shared_ptr<StoreDBMasterIf> iface) {
    ifaces_.push_back(iface);
  }
 public:
  void isOk(Response& _return) {
    size_t sz = ifaces_.size();
    size_t i = 0;
    for (; i < (sz - 1); ++i) {
      ifaces_[i]->isOk(_return);
    }
    ifaces_[i]->isOk(_return);
    return;
  }

  void add(Response& _return, const std::string& key, const std::string& value) {
    size_t sz = ifaces_.size();
    size_t i = 0;
    for (; i < (sz - 1); ++i) {
      ifaces_[i]->add(_return, key, value);
    }
    ifaces_[i]->add(_return, key, value);
    return;
  }

  void search(MultiResponse& _return, const std::string& prefix) {
    size_t sz = ifaces_.size();
    size_t i = 0;
    for (; i < (sz - 1); ++i) {
      ifaces_[i]->search(_return, prefix);
    }
    ifaces_[i]->search(_return, prefix);
    return;
  }

};

} // namespace

#endif
