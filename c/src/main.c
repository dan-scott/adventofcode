//
// Created by Dan Scott on 3/12/22.
//


#include <stdlib.h>
#include <stdio.h>
#include <strings.h>

#include "defines.h"
#include "aoc_2022/aoc_2022.h"

#define RUN_DAY(year, day)  input = get_input_file(year, day);  \
                            PART(year, day, 1)(input);          \
                            PART(year, day, 2)(input);          \
                            free(input);                        \
                            printf("--------------------------------\n");

 char* get_input_file(uint16_t year, uint8_t day) {
    size_t path_length = strlen(INPUT_DIR) + 12;
    char *path = malloc(path_length);
    sprintf(path, "%s%u/%u.txt", INPUT_DIR, year, day);
    FILE *file;
    file = fopen(path, "rb");
    if (file == NULL) {
        fprintf(stderr, "Failed to open input file %s", path);
        exit(EXIT_FAILURE);
    }

    fseek(file, 0L, SEEK_END);
    size_t file_size = ftell(file);
    rewind(file);
    char* buffer = (char*)malloc(file_size + 1);
    if (buffer == NULL) {
        fprintf(stderr, "Not enough memory to read %s", path);
        exit(EXIT_FAILURE);
    }

    size_t bytes_read = fread(buffer, sizeof (char), file_size, file);
    if (bytes_read < file_size) {
        fprintf(stderr, "Could not read file %s", path);
        exit(EXIT_FAILURE);
    }

    buffer[bytes_read] = '\0';

    fclose(file);
    return buffer;
}


int main() {
    char* input;

    RUN_DAY(2022, 1)
    RUN_DAY(2022, 2)
    RUN_DAY(2022, 3)
    RUN_DAY(2022, 4)

    return EXIT_SUCCESS;
}