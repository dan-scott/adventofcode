//
// Created by Dan Scott on 3/12/22.
//


#include <stdlib.h>
#include <stdio.h>
#include <strings.h>

#include "defines.h"

void get_input_file(FILE **file, uint16_t year, uint8_t day) {
    int32_t len = strlen(INPUT_DIR) + 12;
    char *path = malloc(len);
    sprintf(path, "%s%u/%u.txt", INPUT_DIR, year, day);
    *file = fopen(path, "rb");
    if (file == NULL) {
        fprintf(stderr, "Failed to open input file %s", path);
        exit(EXIT_FAILURE);
    }
}

void part_1() {
    FILE *input;
    get_input_file(&input, 2022, 1);
    int ch;
    uint32_t max = 0;
    uint32_t current = 0;
    uint32_t current_sum = 0;
    bool prev_nl = false;
    while ((ch = fgetc(input)) != EOF) {
        if (ch == '\n') {
            if (prev_nl) {
                if (current_sum > max) {
                    max = current_sum;
                }
                current = 0;
                current_sum = 0;
            } else {
                prev_nl = true;
                current_sum += current;
                current = 0;
            }
        } else {
            prev_nl = false;
            current = current * 10 + (ch - '0');
        }
    }
    fclose(input);
    printf("2022 Day 1 part 1: %u\n", max);
}

void part_2() {
    FILE *input;
    get_input_file(&input, 2022, 1);
    char ch;
    uint32_t max_0 = 0;
    uint32_t max_1 = 0;
    uint32_t max_2 = 0;
    uint32_t current = 0;
    uint32_t current_sum = 0;
    bool prev_nl = false;
    while ((ch = fgetc(input)) != EOF) {
        if (ch == '\n') {
            if (prev_nl) {
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
            } else {
                prev_nl = true;
                current_sum += current;
                current = 0;
            }
        } else {
            prev_nl = false;
            current = current * 10 + (ch - '0');
        }
    }
    fclose(input);
    printf("2022 Day 1 part 2: %u\n", max_0 + max_1 + max_2);
}

int main() {
    part_1();
    part_2();
    return EXIT_SUCCESS;
}