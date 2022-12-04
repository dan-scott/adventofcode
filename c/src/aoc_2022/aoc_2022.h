//
// Created by Dan Scott on 3/12/22.
//

#pragma once

#include <stdio.h>
#include <stdlib.h>

#include "defines.h"

#define PART_FN(day, part) void PART(2022, day, part)(const char* input)

#define DAY(day) PART_FN(day, 1); PART_FN(day, 2);

DAY(1)

DAY(2)

DAY(3)

DAY(4)

#undef DAY
