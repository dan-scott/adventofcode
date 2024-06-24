const std = @import("std");

pub const year = 2019;
pub const day = 4;

fn getRange(input: []const u8) !struct { min: usize, max: usize } {
    var parts = std.mem.splitScalar(u8, input, '-');
    const a = parts.next() orelse unreachable;
    const b = parts.next() orelse unreachable;
    const min = try std.fmt.parseInt(usize, a, 10);
    const max = try std.fmt.parseInt(usize, b, 10);
    return .{ .min = min, .max = max };
}

fn getDigits(val: usize) [6]usize {
    return [6]usize{ val / 100_000, (val % 100_000) / 10_000, (val % 10_000) / 1_000, (val % 1_000) / 100, (val % 100) / 10, val % 10 };
}

pub fn part1solver(alloc: std.mem.Allocator, input: []const u8) ![]const u8 {
    const range = try getRange(input);
    var ct: u32 = 0;
    for (range.min..(range.max + 1)) |val| {
        const digits = getDigits(val);
        var has_double = false;
        var increasing = true;
        var prev_dgt = digits[0];
        check: for (1..6) |idx| {
            if (digits[idx] == prev_dgt) {
                has_double = true;
            } else if (digits[idx] < prev_dgt) {
                increasing = false;
                break :check;
            }
            prev_dgt = digits[idx];
        }
        if (has_double and increasing) {
            ct += 1;
        }
    }
    return std.fmt.allocPrint(alloc, "{}", .{ct});
}

pub fn part2solver(alloc: std.mem.Allocator, input: []const u8) ![]const u8 {
    const range = try getRange(input);
    var ct: u32 = 0;
    for (range.min..(range.max + 1)) |val| {
        const digits = getDigits(val);
        var has_double = false;
        var increasing = true;
        var current_group_size: u8 = 0;
        var prev_dgt = digits[0];
        check: for (1..6) |idx| {
            if (digits[idx] == prev_dgt) {
                current_group_size += 1;
            } else if (digits[idx] < prev_dgt) {
                increasing = false;
                break :check;
            } else if (current_group_size > 0) {
                if (current_group_size == 1) {
                    has_double = true;
                }
                current_group_size = 0;
            }
            prev_dgt = digits[idx];
        }
        if (current_group_size == 1) {
            has_double = true;
        }
        if (has_double and increasing) {
            ct += 1;
        }
    }
    return std.fmt.allocPrint(alloc, "{}", .{ct});
}

test "part1Solver" {
    const cases = [_]struct { input: []const u8, expected: []const u8 }{
        .{
            .input = "111111-111111",
            .expected = "1",
        },
        .{
            .input = "223450-223450",
            .expected = "0",
        },
        .{
            .input = "123789-123789",
            .expected = "0",
        },
    };
    const allocator = std.testing.allocator;
    for (cases) |case| {
        const actual = try part1solver(allocator, case.input);
        defer allocator.free(actual);
        try std.testing.expectEqualStrings(case.expected, actual);
    }
}

test "part2Solver" {
    const cases = [_]struct { input: []const u8, expected: []const u8 }{
        .{
            .input = "112233-112233",
            .expected = "1",
        },
        .{
            .input = "123444-123444",
            .expected = "0",
        },
        .{
            .input = "111122-111122",
            .expected = "1",
        },
    };
    const allocator = std.testing.allocator;
    for (cases) |case| {
        const actual = try part2solver(allocator, case.input);
        defer allocator.free(actual);
        try std.testing.expectEqualStrings(case.expected, actual);
    }
}
