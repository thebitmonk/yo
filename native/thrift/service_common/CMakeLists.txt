cmake_minimum_required(VERSION 2.8.4)

set(PROJECT_NAME service_common)
project(${PROJECT_NAME} C CXX)

FILE(GLOB GEN_SOURCE_FILES StoreDB_constants.cpp StoreDB_types.cpp StoreDBMaster.cpp)

add_library(${PROJECT_NAME} ${GEN_SOURCE_FILES})
ADD_DEFINITIONS(-Wall -std=c++11 -O2)
