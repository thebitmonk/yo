#ifndef _DB_INTERFACE_
#define _DB_INTERFACE_

#include "service_common/StoreDB_types.h"
#include <iostream>
#include <vector>

using namespace std;
using namespace storedb;
class DBInterface
{
    public:
        virtual ~DBInterface() {}
        virtual bool openDB(const string& path) = 0;
        virtual bool put(const string& key, const string& value) = 0;
        virtual bool search(const string& prefix, vector<string>& values) = 0;
};

#endif
