const std = @import("std");
const util = @import("../utils.zig");
const VM = @import("./int_code_vm.zig").IntcodeVM;

pub const year = 2019;
pub const day = 9;

pub fn part1solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var vm = try VM.init(allocator, input);
    defer vm.deinit();
    vm.run();
    vm.input(1);

    return std.fmt.allocPrint(allocator, "{any}", .{vm.output()});
}

pub fn part2solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var vm = try VM.init(allocator, input);
    defer vm.deinit();
    vm.run();
    vm.input(2);

    const output = vm.output();

    return std.fmt.allocPrint(allocator, "{}", .{output});
}

test "part1Solver" {
    const cases = [_]struct { input: []const u8, expected: []const u8 }{};
    const allocator = std.testing.allocator;
    for (cases) |case| {
        const actual = try part1solver(allocator, case.input);
        defer allocator.free(actual);
        try std.testing.expectEqualStrings(case.expected, actual);
    }
}

test "part2Solver" {
    const cases = [_]struct { input: []const u8, expected: []const u8 }{};
    const allocator = std.testing.allocator;
    for (cases) |case| {
        const actual = try part2solver(allocator, case.input);
        defer allocator.free(actual);
        try std.testing.expectEqualStrings(case.expected, actual);
    }
}
