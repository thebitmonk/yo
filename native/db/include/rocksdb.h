#ifndef _ROCKSDB_INTERFACE_
#define _ROCKSDB_INTERFACE_

#include "db_interface.h"
#include "rocksdb/db.h"
#include "service_common/StoreDB_types.h"
#include <iostream>

using namespace rocksdb;
using namespace std;

class RocksDB: public DBInterface{

private:
    DB *db;
    void printDB();

public:
    ~RocksDB();
    bool openDB(const string& path);
    bool put(const string& key, const string& value);
    bool search(const string& prefix, vector<string>& values);
};
#endif
