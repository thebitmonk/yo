cmake_minimum_required(VERSION 2.8.4)

ADD_DEFINITIONS(-Wall -std=c++11)

# find thrift headers and libs
find_package(Thrift REQUIRED)
if(Thrift_FOUND)
    include_directories(${Thrift_INCLUDE_DIR})
    include_directories(${Thrift_INCLUDE_DIR}/thrift)
else()
    message(WARNING "Thrift NOT found")
endif()

find_package(Boost REQUIRED COMPONENTS system filesystem)
if(Boost_FOUND)
    include_directories(${Boost_INCLUDE_DIRS})
    message(STATUS "Boost found, include dir: ${Boost_INCLUDE_DIRS}")
else()
    message(WARNING "Boost NOT found")
endif()

#----------------------------------------------------------
# Project
#----------------------------------------------------------

set(PROJECT_NAME_STR storedb_master)
project(${PROJECT_NAME_STR} C CXX)

#----------------------------------------------------------
# Build Master Executable 
#----------------------------------------------------------

file(GLOB MASTER_SOURCE_FILES StoreDBMaster_server.cpp)

set(PROJECT_EXE_NAME storedb_master_server.out)

add_executable(${PROJECT_EXE_NAME} ${MASTER_SOURCE_FILES})

target_include_directories(${PROJECT_EXE_NAME} PUBLIC
    ${THRIFT_DIR}
    ${DB_DIR}/include
)

target_link_libraries(${PROJECT_EXE_NAME} db service_common ${Thrift_STATIC_LIB} pthread)
