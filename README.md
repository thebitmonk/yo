# yo
like man but less verbose and with real practical examples of various tools / commands

This project is WIP and has been developed during SaveTheHacker hackathon.
The storage layer has been kept very general since I plan to develop
personal dropbox and task manager sort of solution on top of that.

In terms of design ther are 5 major parts:
* cli (golang) :  takes care of arg parsing and messaging display to the user
* client (golang) : http client talks to api server. Understands error code and 
                    and sends back appropriate message to cli
* api-server (golang) : processes requests. The authentication and validation 
                   shall be plugged at this layer.
* thrift server (C++) : api server shall talk via pool of thrift clients to 
                  thrift RPC server
* storage layer : all the data is stored in the form of binary key, value. 
                RocksDB from the family of level DB has been used for storing as of now.
                 The entire system is pluggable since it is built against an interface.

TODO:
----
* thread safe thrift client pool
* multiplexing requests via channel to the thrift client pool
* email flow for a new user. Not necessary for search but must while adding
  a new command and usage.
 
