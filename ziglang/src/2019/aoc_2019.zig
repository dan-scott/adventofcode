const std = @import("std");
const day01 = @import("./day_01.zig");
const day02 = @import("./day_02.zig");
const day03 = @import("./day_03.zig");
const day04 = @import("./day_04.zig");
const day05 = @import("./day_05.zig");
const utils = @import("../utils.zig");

pub fn run(allocator: std.mem.Allocator) !void {
    try utils.dayRunner(allocator, day01);
    try utils.dayRunner(allocator, day02);
    try utils.dayRunner(allocator, day03);
    try utils.dayRunner(allocator, day04);
    try utils.dayRunner(allocator, day05);
}

test "aoc_2019" {
    _ = @import("./int_code_vm.zig");
    _ = @import("./day_01.zig");
    _ = @import("./day_02.zig");
    _ = @import("./day_03.zig");
    _ = @import("./day_04.zig");
    _ = @import("./day_05.zig");
}
