#include "rocksdb.h"
#include <iostream>
#include <thread>
#include "rocksdb/options.h"
#include "service_common/StoreDB_types.h"

using namespace std;

// Private Methods


void RocksDB::printDB(){
    cout << "Printing DB " << endl;
    rocksdb::Iterator* it = db->NewIterator(rocksdb::ReadOptions());
    int count = 0;
    for (it->SeekToFirst(); it->Valid() && count < 10; it->Next()) {
        count++;
        cout << it->key().ToString() << ":" << endl;
    }
    assert(it->status().ok());
    delete it;
}

bool RocksDB::openDB(const string& path){
    cout << "Open called " << endl;

    Options options;
    options.IncreaseParallelism();
    options.OptimizeLevelStyleCompaction();
    options.PrepareForBulkLoad();
    std::size_t N = 512;
    options.write_buffer_size = N;
    options.create_if_missing = true;

    Status s = DB::Open(options, path, &db);
    cout << "DB created " << s.ok() << endl;
    return s.ok();
};

bool RocksDB::put(const string& key, const string& value){
    cout << "Put called " << endl;
    Status s = db->Put(WriteOptions(), key, value);
    printDB();
    return s.ok();
};

bool RocksDB::search(const string& prefix, vector<string>& values){
    cout << "Prefix search called" << endl;
    auto iter = db->NewIterator(ReadOptions());
    for (iter->Seek(prefix); iter->Valid() && iter->key().starts_with(prefix); iter->Next()) {
      values.push_back(iter->value().ToString());
    }
    return true;
};

RocksDB::~RocksDB(){
    cout << "Delete db " << endl;
    delete db;
}
