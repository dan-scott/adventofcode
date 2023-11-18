const std = @import("std");
const day01 = @import("./day_01.zig");
const utils = @import("../utils.zig");

pub fn run(allocator: std.mem.Allocator) !void {
    try utils.dayRunner(allocator, day01);
}

test "aoc_2019" {
    _ = @import("./day_01.zig");
}
