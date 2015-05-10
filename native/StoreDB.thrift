namespace cpp storedb

struct Key {
  1: required i64 id
  2: required binary command
  3: optional binary email
}

struct Value {
  1: required binary command 
  2: required binary options 
  3: required binary description
  4: i64 upVotes
  5: i64 downVotes  
}

struct Response {
  1: required bool status
  2: optional double time
  3: optional binary data
}

struct MultiResponse {
  1: required bool status
  2: optional double time
  3: optional list<binary> data
}

service StoreDBMaster {
  Response isOk(),
  Response add(1:binary key, 2:binary value)
  MultiResponse search(1:binary prefix)
}
