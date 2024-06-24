const std = @import("std");

pub const year = 2019;
pub const day = 1;

pub fn part1solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var total: i32 = 0;
    var lines = std.mem.splitSequence(u8, input, "\n");
    while (lines.next()) |numStr| {
        const moduleWeight = try std.fmt.parseInt(i32, numStr, 10);
        total += @divFloor(moduleWeight, 3) - 2;
    }
    return std.fmt.allocPrint(allocator, "{}", .{total});
}

pub fn part2solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var total: i32 = 0;
    var lines = std.mem.splitSequence(u8, input, "\n");
    while (lines.next()) |numStr| {
        const moduleWeight = try std.fmt.parseInt(i32, numStr, 10);
        total += calcFuel(moduleWeight);
    }
    return std.fmt.allocPrint(allocator, "{}", .{total});
}

fn calcFuel(weight: i32) i32 {
    var nextWeight = weight;
    var total: i32 = 0;
    while (true) {
        nextWeight = @divFloor(nextWeight, 3) - 2;
        total += nextWeight;
        if (nextWeight < 9) {
            return total;
        }
    }
}

test "part1Solver" {
    const alloc = std.testing.allocator;
    const input =
        \\12
        \\14
        \\1969
        \\100756
    ;
    const expected = "34241";
    const solution = try part1solver(alloc, input);
    defer alloc.free(solution);

    try std.testing.expectEqualStrings(expected, solution);
}

test "part2Solver" {
    const alloc = std.testing.allocator;
    const cases = [_]struct { input: []const u8, expected: []const u8 }{
        .{ .input = "14", .expected = "2" },
        .{ .input = "1969", .expected = "966" },
        .{ .input = "100756", .expected = "50346" },
    };
    for (cases) |case| {
        const sln = try part2solver(alloc, case.input);
        defer alloc.free(sln);
        try std.testing.expectEqualStrings(case.expected, sln);
    }
}
