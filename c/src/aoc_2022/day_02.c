//
// Created by Dan Scott on 3/12/22.
//

#include "aoc_2022.h"

//0 = R
//1 = P
//2 = S

static uint8_t score_sheet[9] = {
        4, 8, 3,
        1, 5, 9,
        7, 2, 6,
};

// 0 = lose
// 1 = draw
// 2 = win
static uint8_t decode_sheet[9] = {
        3, 4, 8,
        1, 5, 9,
        2, 6, 7,
};

PART_FN(2022, 2, 1) {
    uint32_t total = 0;
    while (*input != '\0') {
        int elf_hand = *input - 'A';
        input += 2;
        int my_hand = *input - 'X';
        total += score_sheet[3 * elf_hand + my_hand];
        input += 2;
    }
    printf("2022 Day 2 part 1: %u\n", total);
}

PART_FN(2022, 2, 2) {
    uint32_t total = 0;
    while (*input != '\0') {
        int elf_hand = *input - 'A';
        input += 2;
        int my_hand = *input - 'X';
        total += decode_sheet[3 * elf_hand + my_hand];
        input += 2;
    }
    printf("2022 Day 2 part 2: %u\n", total);
}
