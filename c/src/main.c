//
// Created by Dan Scott on 3/12/22.
//


#include <stdlib.h>
#include <stdio.h>
#include <strings.h>

#include "defines.h"
#include "aoc_2022/aoc_2022.h"

#define RUN_DAY(year, day)  get_input_file(&input, year, day);  \
                            PART(year, day, 1)(input);          \
                            rewind(input);                      \
                            PART(year, day, 2)(input);          \
                            fclose(input);

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


int main() {
    FILE *input;

    RUN_DAY(2022,1)

    return EXIT_SUCCESS;
}