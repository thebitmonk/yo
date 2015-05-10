cmake_minimum_required(VERSION 2.8.4)

set(PROJECT_NAME thrift)
project(${PROJECT_NAME} C CXX)

set(SERVICE_COMMON_DIR ${PROJECT_SOURCE_DIR}/service_common)

add_subdirectory(${SERVICE_COMMON_DIR})
