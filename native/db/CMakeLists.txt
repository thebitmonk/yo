cmake_minimum_required(VERSION 2.8.4)

ADD_DEFINITIONS(-Wall -std=c++11 -O2  -DROCKSDB_PLATFORM_POSIX  -DOS_LINUX -fno-builtin-memcmp -DROCKSDB_ATOMIC_PRESENT -DROCKSDB_FALLOCATE_PRESENT -DSNAPPY -DGFLAGS=google -DZLIB -DBZIP2 -DLZ4)

# find thrift headers and libs
find_package(Thrift REQUIRED)
if(Thrift_FOUND)
    include_directories(${Thrift_INCLUDE_DIR})
    include_directories(${Thrift_INCLUDE_DIR}/thrift)
else()
    message(WARNING "Thrift NOT found")
endif()

find_package(RocksDB REQUIRED)
if(ROCKSDB_FOUND)
    include_directories(${ROCKSDB_INCLUDE})
    message(STATUS "RocksDB found, include dir: ${ROCKSDB_INCLUDE}")
else()
    message(WARNING "RocksDB NOT found")
endif()

#----------------------------------------------------------
# Project
#----------------------------------------------------------


set(PROJECT_NAME db)
project(${PROJECT_NAME} C CXX)


#----------------------------------------------------------
# Build DB - library
#----------------------------------------------------------

file(GLOB DB_SOURCE_FILES src/rocksdb.cpp)

set(PROJECT_LIB_NAME db)

add_library(${PROJECT_LIB_NAME} ${DB_SOURCE_FILES})

target_include_directories( ${PROJECT_LIB_NAME} PUBLIC
    ${PROJECT_SOURCE_DIR}/include
    ${THRIFT_DIR}
)

target_link_libraries(${PROJECT_LIB_NAME} ${ROCKSDB_STATIC_LIBRARIES} service_common -lpthread -lsnappy -lgflags -lz -lbz2 -llz4 ${THRIFT_STATIC_LIB})

