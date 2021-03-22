cmake_minimum_required(VERSION 3.17)
project(dll)

set(CMAKE_C_STANDARD 99)

include_directories(keiro)
include_directories(keiro/minhook)
include_directories(imgui)

add_library(dll sigscan.cpp sigscan.h ntdlr.h _patterns.h mem.cpp mem.h main.cpp)