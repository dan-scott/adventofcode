//
// Created by Dan Scott on 4/12/22.
//

#include "aoc_2022.h"
#include <strings.h>

size_t code_to_points(int code) {
    if (code > 'a') {
        return code - 'a' + 1;
    }
    return code - 'A' + 27;
}


PART_FN(2022, 3, 1) {
    size_t points[53];
    memset(points, 0, sizeof (points));
    char buf[100];
    while (*input != '\0') {
        const char* start = input;
        while(*input != '\n') {
            input++;
        }
        size_t len = input - start;
        size_t half = len / 2;
        const char* right = start + half;
        memcpy(buf, right, half);
        buf[half] = '\0';
        while (start < right) {
            if(strchr(buf, *start)) {
                points[code_to_points(*start)]++;
                break;
            }
            start++;
        }
        input++;
    }
    size_t total = 0;
    for(size_t i = 1; i < 53; i++ ) {
        if (points[i] > 0) {
            total += i * points[i];
        }
    }
    printf("2022 Day 3 part 1: %zu\n", total);
}

PART_FN(2022, 3, 2) {
    size_t points[53];
    memset(points, 0, sizeof (points));
    char r1[100];
    char r2[100];
    char r3[100];
    while (*input != '\0') {
        char* end = memccpy(r1, input, '\n', 100);
        *(end - 1) = '\0';
        size_t r1_len = end - r1;
        input += r1_len;
        end = memccpy(r2, input, '\n', 100);
        *(end - 1) = '\0';
        input += end - r2;
        end = memccpy(r3, input, '\n', 100);
        *(end-1) = '\0';
        input += end - r3;
        for(size_t i = 0; i <  r1_len; i++) {
           if (strchr(r2, r1[i])) {
               if (strchr(r3, r1[i])) {
                   points[code_to_points(r1[i])]++;
                   break;
               }
           }
        }
    }
    size_t total = 0;
    for(size_t i = 1; i < 53; i++ ) {
        if (points[i] > 0) {
            total += i * points[i];
        }
    }
    printf("2022 Day 3 part 2: %zu\n", total);
}