const std = @import("std");
const day01 = @import("./day_01.zig");
const day02 = @import("./day_02.zig");
const utils = @import("../utils.zig");

pub fn run(allocator: std.mem.Allocator) !void {
    try utils.dayRunner(allocator, day01);
    try utils.dayRunner(allocator, day02);
}

test "aoc_2019" {
    _ = @import("./int_code_vm.zig");
    _ = @import("./day_01.zig");
    _ = @import("./day_02.zig");
}
