# Copied and modified from https://github.com/cloudera/Impala/blob/master/cmake_modules/FindThrift.cmake

# - Find Thrift (a cross platform RPC lib/tool)
# This module defines
#  Thrift_VERSION, version string of ant if found
#  Thrift_INCLUDE_DIR, where to find THRIFT headers
#  Thrift_CONTRIB_DIR, where contrib thrift files (e.g. fb303.thrift) are installed
#  Thrift_FOUND, If false, do not try to use ant
#  Thrift_STATIC_LIB, path to static library

# Use following block to unset any cache entries
#unset(Thrift_INCLUDE_DIR CACHE)
#unset(Thrift_CONTRIB_DIR CACHE)
#unset(Thrift_LIBS CACHE)
#unset(Thrift_STATIC_LIB CACHE)
#unset(Thrift_COMPILER CACHE)

# prefer the thrift version supplied in THRIFT_HOME
message(STATUS "THRIFT_HOME: $ENV{THRIFT_HOME}")

# find include dir path
find_path(Thrift_INCLUDE_DIR
  NAMES thrift/Thrift.h
  HINTS $ENV{THRIFT_HOME}/include/ /usr/local/include/ /usr/include/
)

# find contrib dir path
find_path(Thrift_CONTRIB_DIR
  NAMES share/fb303/if/fb303.thrift
  HINTS $ENV{THRIFT_HOME} /usr/local/
)

# find static lib path
find_path(Thrift_STATIC_LIB_PATH
  NAMES libthrift.a
  HINTS $ENV{THRIFT_HOME}/lib/ /usr/local/lib/ /usr/lib/
)

# find compiler path
find_program(Thrift_COMPILER
  NAMES thrift
  HINTS $ENV{THRIFT_HOME}/bin/ /usr/local/bin/ /usr/bin/
  NO_DEFAULT_PATH
)

# set remaining variables
if (Thrift_STATIC_LIB_PATH)
  set(Thrift_FOUND TRUE)
  set(Thrift_STATIC_LIB ${Thrift_STATIC_LIB_PATH}/libthrift.a)
  set(ThriftNB_STATIC_LIB ${Thrift_STATIC_LIB_PATH}/libthriftnb.a)
  exec_program(${Thrift_COMPILER}
    ARGS -version OUTPUT_VARIABLE Thrift_VERSION RETURN_VALUE Thrift_RETURN)
else ()
  set(Thrift_FOUND FALSE)
endif ()

if (Thrift_FOUND)
  message(STATUS "Thrift version: ${Thrift_VERSION}")
  message(STATUS "Thrift include dir: ${Thrift_INCLUDE_DIR}")
  message(STATUS "Thrift contrib dir: ${Thrift_CONTRIB_DIR}")
  message(STATUS "Thrift static library: ${Thrift_STATIC_LIB}")
  message(STATUS "Thrift compiler: ${Thrift_COMPILER}")
else ()
  message(STATUS "Thrift compiler/libraries NOT found")
endif ()