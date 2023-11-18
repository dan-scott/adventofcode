const std = @import("std");
const aoc2019 = @import("./2019/aoc_2019.zig");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    var alloc = gpa.allocator();
    try aoc2019.run(alloc);
}

test "aoc" {
    _ = @import("./utils.zig");
    _ = @import("./2019/aoc_2019.zig");
}
