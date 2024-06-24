const std = @import("std");

pub const year = 2019;
pub const day = 3;

const Vec2 = struct {
    x: i32,
    y: i32,

    fn manhattan(self: *Vec2, from: Vec2) i32 {
        return @intCast(@abs(from.x - self.x) + @abs(from.y - self.y));
    }

    fn add(self: *Vec2, other: Vec2) void {
        self.x += other.x;
        self.y += other.y;
    }

    fn unitDir(dir: Direction) Vec2 {
        return switch (dir) {
            .U => .{ .x = 0, .y = 1 },
            .D => .{ .x = 0, .y = -1 },
            .L => .{ .x = -1, .y = 0 },
            .R => .{ .x = 1, .y = 0 },
        };
    }
};

const ZERO_VEC: Vec2 = .{ .x = 0, .y = 0 };

const WireMap = std.hash_map.AutoHashMap(Vec2, bool);

const Direction = enum {
    const Self = @This();
    R,
    L,
    U,
    D,

    pub fn from(char: u8) Self {
        return switch (char) {
            'U' => .U,
            'D' => .D,
            'L' => .L,
            else => .R,
        };
    }
};

pub fn part1solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var map = WireMap.init(allocator);
    defer map.deinit();
    var lineSplit = std.mem.splitScalar(u8, input, '\n');
    const wire1 = lineSplit.next() orelse return "failed";
    const wire2 = lineSplit.next() orelse return "failed";

    var wire1Iter = WireIterator.from(wire1);
    var pos = Vec2{ .x = 0, .y = 0 };
    while (wire1Iter.next()) |step| {
        const unitDelta = Vec2.unitDir(step.dir);
        for (0..step.dist) |_| {
            pos.add(unitDelta);
            try map.put(pos, true);
        }
    }

    var wire2Iter = WireIterator.from(wire2);
    pos.x = 0;
    pos.y = 0;
    var minDist: i32 = std.math.maxInt(i32);
    while (wire2Iter.next()) |step| {
        const unitDelta = Vec2.unitDir(step.dir);
        for (0..step.dist) |_| {
            pos.add(unitDelta);
            if (map.contains(pos)) {
                const dist = pos.manhattan(ZERO_VEC);
                minDist = @min(minDist, dist);
            }
        }
    }
    return std.fmt.allocPrint(allocator, "{}", .{minDist});
}

const WireHopCount = struct {
    wire1Hops: i32,
    wire2Hops: i32,

    pub fn combinedHops(self: *WireHopCount) i32 {
        return self.wire1Hops + self.wire2Hops;
    }
};

const WireHopMap = std.hash_map.AutoHashMap(Vec2, WireHopCount);

pub fn part2solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var map = WireHopMap.init(allocator);
    defer map.deinit();
    var lineSplit = std.mem.splitScalar(u8, input, '\n');
    const wire1 = lineSplit.next() orelse return "failed";
    const wire2 = lineSplit.next() orelse return "failed";

    var wire1Iter = WireIterator.from(wire1);
    var pos = Vec2{ .x = 0, .y = 0 };
    var hopCt: i32 = 0;
    while (wire1Iter.next()) |step| {
        const unitDelta = Vec2.unitDir(step.dir);
        for (0..step.dist) |_| {
            hopCt += 1;
            pos.add(unitDelta);
            if (map.getPtr(pos)) |posHops| {
                posHops.wire1Hops = @min(posHops.wire1Hops, hopCt);
            } else {
                try map.put(pos, .{ .wire1Hops = hopCt, .wire2Hops = std.math.maxInt(i32) });
            }
        }
    }

    pos.x = 0;
    pos.y = 0;
    var wire2Iter = WireIterator.from(wire2);
    var minHops: i32 = std.math.maxInt(i32);
    hopCt = 0;
    while (wire2Iter.next()) |step| {
        const unitDelta = Vec2.unitDir(step.dir);
        for (0..step.dist) |_| {
            hopCt += 1;
            pos.add(unitDelta);
            if (map.getPtr(pos)) |posHops| {
                posHops.wire2Hops = @min(posHops.wire2Hops, hopCt);
                if (posHops.wire1Hops > -1) {
                    minHops = @min(minHops, posHops.combinedHops());
                }
            } else {
                try map.put(pos, .{ .wire1Hops = -1, .wire2Hops = hopCt });
            }
        }
    }

    return std.fmt.allocPrint(allocator, "{}", .{minHops});
}

const WireIterator = struct {
    const Self = @This();

    pub const WireStep = struct {
        dist: usize,
        dir: Direction,
    };

    tokenIterator: std.mem.TokenIterator(u8, .scalar),

    pub fn from(wireInput: []const u8) Self {
        return .{
            .tokenIterator = std.mem.tokenizeScalar(u8, wireInput, ','),
        };
    }

    pub fn next(self: *Self) ?Self.WireStep {
        if (self.tokenIterator.next()) |t| {
            const dir = Direction.from(t[0]);
            const dist = std.fmt.parseInt(usize, t[1..], 10) catch return null;
            return .{
                .dir = dir,
                .dist = dist,
            };
        }
        return null;
    }
};

test "part1Solver" {
    const cases = [_]struct { input: []const u8, expected: []const u8 }{
        .{
            .input =
            \\R8,U5,L5,D3
            \\U7,R6,D4,L4
            ,
            .expected = "6",
        },
        .{
            .input =
            \\R75,D30,R83,U83,L12,D49,R71,U7,L72
            \\U62,R66,U55,R34,D71,R55,D58,R83
            ,
            .expected = "159",
        },
        .{
            .input =
            \\R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
            \\U98,R91,D20,R16,D67,R40,U7,R15,U6,R7
            ,
            .expected = "135",
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
            .input =
            \\R8,U5,L5,D3
            \\U7,R6,D4,L4
            ,
            .expected = "30",
        },
        .{
            .input =
            \\R75,D30,R83,U83,L12,D49,R71,U7,L72
            \\U62,R66,U55,R34,D71,R55,D58,R83
            ,
            .expected = "610",
        },
        .{
            .input =
            \\R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
            \\U98,R91,D20,R16,D67,R40,U7,R15,U6,R7
            ,
            .expected = "410",
        },
    };
    const allocator = std.testing.allocator;
    for (cases) |case| {
        const actual = try part2solver(allocator, case.input);
        defer allocator.free(actual);
        try std.testing.expectEqualStrings(case.expected, actual);
    }
}
