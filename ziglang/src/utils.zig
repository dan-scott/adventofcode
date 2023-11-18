const std = @import("std");
pub const log_level: std.log.Level = .info;

pub const DayInputError = error{
    RootDirNotSet,
    InputFileNotFound,
};

pub fn getDay(allocator: std.mem.Allocator, year: u16, day: u8) ![]const u8 {
    const root = std.os.getenv("ADVENT_OF_CODE_ROOT") orelse {
        std.log.err("ADVENT_OF_CODE_ROOT env var not set", .{});
        return DayInputError.RootDirNotSet;
    };

    const subPathStr = try std.fmt.allocPrint(allocator, "inputs/{}/{}.txt", .{ year, day });
    defer allocator.free(subPathStr);

    const path = try std.fs.path.join(allocator, &[_][]const u8{ root, subPathStr });
    defer allocator.free(path);

    const dayFile = std.fs.openFileAbsolute(path, .{}) catch {
        std.log.err("Unable to open day file: {s}", .{path});
        return DayInputError.InputFileNotFound;
    };
    const fileSize = (try dayFile.stat()).size;

    return dayFile.readToEndAlloc(allocator, fileSize);
}

pub fn dayRunner(allocator: std.mem.Allocator, day: anytype) !void {
    var arena = std.heap.ArenaAllocator.init(allocator);
    defer arena.deinit();
    const aAlloc = arena.allocator();

    const input = try getDay(aAlloc, day.year, day.day);
    const trimmed = std.mem.trimRight(u8, input, "\n");

    std.debug.print("{d} Day {d}\n", .{ day.year, day.day });

    const part1result = try day.part1solver(aAlloc, trimmed);
    const part2result = try day.part2solver(aAlloc, trimmed);

    std.debug.print("     Part 1: {s}\n     Part 2: {s}\n", .{ part1result, part2result });
}

test "read day file" {
    const alloc = std.testing.allocator;
    const dayContent = try getDay(alloc, 2016, 1);
    defer alloc.free(dayContent);

    try std.testing.expectEqualSlices(u8, "L5,", dayContent[0..3]);
}
