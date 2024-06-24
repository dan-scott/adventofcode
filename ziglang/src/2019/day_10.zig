const std = @import("std");
const util = @import("../util.zig");

pub const year = 2019;
pub const day = 10;

const Vec = util.vec.Vec2(isize);

const VecSet = std.AutoHashMap(Vec, void);

const Field = struct {
    width: isize,
    height: isize,
    asteroids: VecSet,
    allocator: std.mem.Allocator,

    pub fn init(allocator: std.mem.Allocator) Field {
        const set = VecSet.init(allocator);
        return .{
            .width = 0,
            .height = 0,
            .asteroids = set,
            .allocator = allocator,
        };
    }

    pub fn deinit(self: *Field) void {
        self.asteroids.deinit();
    }
};

pub fn part1solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var field = try parseInput(allocator, input);
    defer field.deinit();

    const res = try findIdeal(allocator, &field);

    return std.fmt.allocPrint(allocator, "{}", .{res.count});
}

fn findIdeal(allocator: std.mem.Allocator, field: *Field) !struct { count: usize, loc: Vec } {
    var outer_iter = field.asteroids.keyIterator();

    var max: usize = 0;
    var max_loc: Vec = Vec.of(0, 0);
    while (outer_iter.next()) |asteroid| {
        var hit_map = VecSet.init(allocator);
        defer hit_map.deinit();

        var inner_iter = field.asteroids.keyIterator();
        var count: usize = 0;
        while (inner_iter.next()) |tgt| {
            if (std.meta.eql(asteroid, tgt)) {
                continue;
            }
            const grad = asteroid.gradient(@constCast(tgt));
            const put = try hit_map.getOrPut(grad);
            if (!put.found_existing) {
                count += 1;
            }
        }
        max = @max(max, count);
        if (max == count) {
            max_loc = asteroid.*;
        }
    }

    return .{ .count = max, .loc = max_loc };
}

const Target = struct {
    loc: Vec,
    dist: f64,
};

const TgtDist = struct {
    const Self = @This();
    const TgtList = std.ArrayList(Target);
    const Map = std.AutoHashMap(Vec, TgtList);

    map: Map,
    src: Vec,
    allocator: std.mem.Allocator,

    pub fn init(allocator: std.mem.Allocator, src: Vec) Self {
        const map = Map.init(allocator);
        return .{
            .map = map,
            .src = src,
            .allocator = allocator,
        };
    }

    pub fn deinit(self: *Self) void {
        var map_iter = self.map.valueIterator();

        while (map_iter.next()) |tgt_list| {
            tgt_list.deinit();
        }

        self.map.deinit();
    }

    pub fn add(self: *Self, tgt: *const Vec) !void {
        const grad = self.src.gradient(tgt);
        const dst = self.src.dist(tgt);

        var put = try self.map.getOrPut(grad);
        if (!put.found_existing) {
            put.value_ptr.* = TgtList.init(self.allocator);
        }

        try put.value_ptr.append(.{ .loc = tgt.*, .dist = dst });
    }

    pub fn cmpByDst(context: void, a: Target, b: Target) bool {
        return std.sort.asc(f64)(context, a.dist, b.dist);
    }

    pub fn sort(self: *Self) !void {
        var lists = self.map.valueIterator();
        while (lists.next()) |list| {
            std.sort.insertion(Target, list.items, {}, Self.cmpByDst);
        }
    }
};

fn calcAngles(allocator: std.mem.Allocator, asteroid: Vec, field: *Field) !TgtDist {
    var tgt_dst = TgtDist.init(allocator, asteroid);
    var tgts = field.asteroids.keyIterator();
    while (tgts.next()) |tgt| {
        try tgt_dst.add(tgt);
    }

    try tgt_dst.sort();

    return tgt_dst;
}

pub fn part2solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var field = try parseInput(allocator, input);
    defer field.deinit();

    const ideal = try findIdeal(allocator, &field);

    var tgt_dst = try calcAngles(allocator, ideal.loc, &field);
    defer tgt_dst.deinit();

    var ordered = TgtDist.TgtList.init(allocator);
    defer ordered.deinit();

    var map_iter = tgt_dst.map.iterator();
    while (map_iter.next()) |tgt_list| {
        const angle = tgt_list.key_ptr.angle();
        for (tgt_list.value_ptr.items, 0..) |tgt, i| {
            const mul: f64 = @floatFromInt(i);
            const new_loc: Target = .{ .loc = tgt.loc, .dist = angle + (std.math.pi * 2) * mul };
            try ordered.append(new_loc);
        }
    }

    std.sort.insertion(Target, ordered.items, {}, TgtDist.cmpByDst);

    const bet = ordered.items[199];

    return std.fmt.allocPrint(allocator, "{}", .{bet.loc.x * 100 + bet.loc.y});
}

fn parseInput(allocator: std.mem.Allocator, input: []const u8) !Field {
    var field = Field.init(allocator);

    var line_iter = std.mem.tokenize(u8, input, "\n");

    var x: isize = 0;
    var y: isize = 0;
    while (line_iter.next()) |line| {
        x = 0;
        for (line) |char| {
            if (char == '#') {
                try field.asteroids.put(Vec.of(x, y), {});
            }
            x += 1;
        }
        y += 1;
    }

    field.width = x;
    field.height = y;

    return field;
}

test "part1Solver" {
    const cases = [_]struct { input: []const u8, expected: []const u8 }{
        .{
            .input =
            \\.#..#
            \\.....
            \\#####
            \\....#
            \\...##
            ,
            .expected = "8",
        },
        .{
            .input =
            \\......#.#.
            \\#..#.#....
            \\..#######.
            \\.#.#.###..
            \\.#..#.....
            \\..#....#.#
            \\#..#....#.
            \\.##.#..###
            \\##...#..#.
            \\.#....####
            ,
            .expected = "33",
        },
        .{
            .input =
            \\#.#...#.#.
            \\.###....#.
            \\.#....#...
            \\##.#.#.#.#
            \\....#.#.#.
            \\.##..###.#
            \\..#...##..
            \\..##....##
            \\......#...
            \\.####.###.
            ,
            .expected = "35",
        },
        .{
            .input =
            \\.#..#..###
            \\####.###.#
            \\....###.#.
            \\..###.##.#
            \\##.##.#.#.
            \\....###..#
            \\..#.#..#.#
            \\#..#.#.###
            \\.##...##.#
            \\.....#.#..
            ,
            .expected = "41",
        },
        .{
            .input =
            \\.#..##.###...#######
            \\##.############..##.
            \\.#.######.########.#
            \\.###.#######.####.#.
            \\#####.##.#.##.###.##
            \\..#####..#.#########
            \\####################
            \\#.####....###.#.#.##
            \\##.#################
            \\#####.##.###..####..
            \\..######..##.#######
            \\####.##.####...##..#
            \\.#####..#.######.###
            \\##...#.##########...
            \\#.##########.#######
            \\.####.#.###.###.#.##
            \\....##.##.###..#####
            \\.#.#.###########.###
            \\#.#.#.#####.####.###
            \\###.##.####.##.#..##
            ,
            .expected = "210",
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
    const cases = [_]struct { input: []const u8, expected: []const u8 }{.{
        .input =
        \\.#..##.###...#######
        \\##.############..##.
        \\.#.######.########.#
        \\.###.#######.####.#.
        \\#####.##.#.##.###.##
        \\..#####..#.#########
        \\####################
        \\#.####....###.#.#.##
        \\##.#################
        \\#####.##.###..####..
        \\..######..##.#######
        \\####.##.####...##..#
        \\.#####..#.######.###
        \\##...#.##########...
        \\#.##########.#######
        \\.####.#.###.###.#.##
        \\....##.##.###..#####
        \\.#.#.###########.###
        \\#.#.#.#####.####.###
        \\###.##.####.##.#..##
        ,
        .expected = "802",
    }};
    const allocator = std.testing.allocator;
    for (cases) |case| {
        const actual = try part2solver(allocator, case.input);
        defer allocator.free(actual);
        try std.testing.expectEqualStrings(case.expected, actual);
    }
}
