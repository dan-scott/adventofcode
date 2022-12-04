cmake_minimum_required(VERSION 3.23)
project(c C)

set(CMAKE_C_STANDARD 11)

add_compile_definitions(_INPUT_DIR="${PROJECT_SOURCE_DIR}/../inputs/")

include_directories(src)

add_executable(c
        src/defines.h
        src/main.c
        src/aoc_2022/aoc_2022.h
        src/aoc_2022/day_01.c
        src/aoc_2022/day_02.c
        src/aoc_2022/day_03.c
        src/aoc_2022/day_04.c)
