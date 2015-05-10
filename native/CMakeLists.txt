cmake_minimum_required(VERSION 2.8.4)

# where to find cmake modules
set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${CMAKE_SOURCE_DIR}/cmake/Modules")

set(PROJECT_NAME_STR CommandStoreDB)
project(${PROJECT_NAME_STR} C CXX)

# Ensures that we do an out of source build

MACRO(MACRO_ENSURE_OUT_OF_SOURCE_BUILD MSG)
     STRING(COMPARE EQUAL "${CMAKE_SOURCE_DIR}"
     "${CMAKE_BINARY_DIR}" insource)
     GET_FILENAME_COMPONENT(PARENTDIR ${CMAKE_SOURCE_DIR} PATH)
     STRING(COMPARE EQUAL "${CMAKE_SOURCE_DIR}"
     "${PARENTDIR}" insourcesubdir)
    IF(insource OR insourcesubdir)
        MESSAGE(FATAL_ERROR "${MSG}")
    ENDIF(insource OR insourcesubdir)
ENDMACRO(MACRO_ENSURE_OUT_OF_SOURCE_BUILD)

MACRO_ENSURE_OUT_OF_SOURCE_BUILD("${CMAKE_PROJECT_NAME} requires an out of source build. Use build.sh")


set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Wextra -std=c++11 -g -O2")

set(THRIFT_DIR ${PROJECT_SOURCE_DIR}/thrift)
set(DB_DIR ${PROJECT_SOURCE_DIR}/db)
set(MASTER_SERVER_DIR ${PROJECT_SOURCE_DIR}/master)

add_subdirectory(${THRIFT_DIR})
add_subdirectory(${DB_DIR})
add_subdirectory(${MASTER_SERVER_DIR})
