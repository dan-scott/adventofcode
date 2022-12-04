//
// Created by Dan Scott on 5/12/22.
//

#include "aoc_2022.h"

size_t chomp_int(const char **input, int t) {
    size_t num = 0;
    while (**input != t) {
        num += num * 10 + (**input - '0');
        *input += 1;
    }
    *input += 1;
    return num;
}

PART_FN(4, 1) {
    size_t count = 0;
    while (*input != '\0') {
        size_t left_start = chomp_int(&input, '-');
        size_t left_end = chomp_int(&input, ',');
        size_t right_start = chomp_int(&input, '-');
        size_t right_end = chomp_int(&input, '\n');

        if ((left_start <= right_start && right_end <= left_end) ||
            (right_start <= left_start && left_end <= right_end)) {
            count++;
        }
    }
    printf("2022 Day 4 part 1: %zu\n", count);
}

PART_FN(4, 2) {
    size_t count = 0;
    while (*input != '\0') {
        size_t left_start = chomp_int(&input, '-');
        size_t left_end = chomp_int(&input, ',');
        size_t right_start = chomp_int(&input, '-');
        size_t right_end = chomp_int(&input, '\n');

        if ((left_start <= right_start && right_start <= left_end) ||
                (right_start <= left_start && left_start <= right_end)) {
            count++;
        }
    }
    printf("2022 Day 4 part 2: %zu\n", count);
}