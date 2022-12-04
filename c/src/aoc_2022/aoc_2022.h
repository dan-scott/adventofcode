//
// Created by Dan Scott on 3/12/22.
//

#pragma once

#include <stdio.h>
#include <stdlib.h>

#include "defines.h"

#define PART_FN(year, day, part) void PART(year, day, part)(const char* input)

#define DAY(year, day) PART_FN(year, day, 1); PART_FN(year, day, 2);

DAY(2022, 1)
DAY(2022, 2)
DAY(2022, 3)

#undef DAY
