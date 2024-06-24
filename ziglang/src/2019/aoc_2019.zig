const std = @import("std");
const dayRunner = @import("../util.zig").runner.dayRunner;

pub fn run(allocator: std.mem.Allocator) !void {
    try dayRunner(allocator, @import("./day_01.zig"));
    try dayRunner(allocator, @import("./day_02.zig"));
    try dayRunner(allocator, @import("./day_03.zig"));
    try dayRunner(allocator, @import("./day_04.zig"));
    try dayRunner(allocator, @import("./day_05.zig"));
    try dayRunner(allocator, @import("./day_06.zig"));
    try dayRunner(allocator, @import("./day_07.zig"));
    try dayRunner(allocator, @import("./day_08.zig"));
    try dayRunner(allocator, @import("./day_09.zig"));
    try dayRunner(allocator, @import("./day_10.zig"));
    try dayRunner(allocator, @import("./day_11.zig"));
    try dayRunner(allocator, @import("./day_12.zig"));
}

test "aoc_2019" {
    std.testing.refAllDecls(@This());
}
