cmake_minimum_required(VERSION 3.23)
project(c C)

set(CMAKE_C_STANDARD 11)

add_compile_definitions(_INPUT_DIR="${PROJECT_SOURCE_DIR}/../inputs/")

include_directories(src)

add_executable(c
        src/defines.h
        src/main.c)
