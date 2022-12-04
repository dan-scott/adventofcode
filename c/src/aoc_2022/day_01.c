//
// Created by Dan Scott on 3/12/22.
//

#include "aoc_2022.h"

PART_FN(2022, 1, 1) {
    int ch;
    uint32_t max = 0;
    uint32_t current = 0;
    uint32_t current_sum = 0;
    while (*input != '\0') {
        if (*input == '\n') {
            current_sum += current;
            current = 0;
            if (*(input + 1) == '\n') {
                if (current_sum > max) {
                    max = current_sum;
                }
                current_sum = 0;
                input++;
            }
        } else {
            current = current * 10 + (*input - '0');
        }
        input++;
    }
    printf("2022 Day 1 part 1: %u\n", max);
}

PART_FN(2022, 1, 2) {
    char ch;
    uint32_t max_0 = 0;
    uint32_t max_1 = 0;
    uint32_t max_2 = 0;
    uint32_t current = 0;
    uint32_t current_sum = 0;
    while (*input != '\0') {
        if (*input == '\n') {
            current_sum += current;
            current = 0;
            if (*(input + 1) == '\n') {
                if (current_sum > max_0) {
                    max_2 = max_1;
                    max_1 = max_0;
                    max_0 = current_sum;
                } else if (current_sum > max_1) {
                    max_2 = max_1;
                    max_1 = current_sum;
                } else if (current_sum > max_2) {
                    max_2 = current_sum;
                }
                current = 0;
                current_sum = 0;
            }
        } else {
            current = current * 10 + (*input - '0');
        }
        input++;
    }
    printf("2022 Day 1 part 2: %u\n", max_0 + max_1 + max_2);
}
