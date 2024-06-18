const std = @import("std");
const utils = @import("../utils.zig");

pub fn run(allocator: std.mem.Allocator) !void {
    try utils.dayRunner(allocator, @import("./day_01.zig"));
    try utils.dayRunner(allocator, @import("./day_02.zig"));
    try utils.dayRunner(allocator, @import("./day_03.zig"));
    try utils.dayRunner(allocator, @import("./day_04.zig"));
    try utils.dayRunner(allocator, @import("./day_05.zig"));
    try utils.dayRunner(allocator, @import("./day_06.zig"));
    try utils.dayRunner(allocator, @import("./day_07.zig"));
    try utils.dayRunner(allocator, @import("./day_08.zig"));
    try utils.dayRunner(allocator, @import("./day_09.zig"));
    try utils.dayRunner(allocator, @import("./day_10.zig"));
}

test "aoc_2019" {
    _ = @import("./int_code_vm.zig");
    _ = @import("./day_01.zig");
    _ = @import("./day_02.zig");
    _ = @import("./day_03.zig");
    _ = @import("./day_04.zig");
    _ = @import("./day_05.zig");
    _ = @import("./day_06.zig");
    _ = @import("./day_07.zig");
    _ = @import("./day_08.zig");
    _ = @import("./day_09.zig");
    _ = @import("./day_10.zig");
}
