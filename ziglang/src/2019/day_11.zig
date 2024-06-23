const std = @import("std");
const util = @import("../utils.zig");
const Vec2 = @import("../vec2.zig").Vec2;
const VM = @import("./int_code_vm.zig").IntcodeVM;

pub const year = 2019;
pub const day = 11;

pub fn part1solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var vm = try VM.init(allocator, input);
    defer vm.deinit();
    vm.run();

    var hull = Hull.init(allocator);
    defer hull.deinit();
    var robot = Robot{ .loc = Vec.of(0, 0), .dir = Dir.up() };

    var painted = std.AutoHashMap(Vec, void).init(allocator);
    defer painted.deinit();

    while (vm.state != .done) {
        const current_color = hull.colorAt(robot.loc);
        vm.input(@intFromEnum(current_color));
        const next_color = vm.output();
        try hull.paint(robot.loc, @enumFromInt(next_color));
        try painted.put(robot.loc, {});

        const turn = vm.output();
        if (turn == 0) {
            robot.dir.left();
        } else {
            robot.dir.right();
        }
        const delta = robot.dir.vec();
        _ = robot.loc.mutAdd(@constCast(&delta));
    }

    return std.fmt.allocPrint(allocator, "{}", .{painted.count()});
}

pub fn part2solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var vm = try VM.init(allocator, input);
    defer vm.deinit();
    vm.run();

    var hull = Hull.init(allocator);
    defer hull.deinit();

    try hull.paint(Vec.of(0, 0), .white);

    var robot = Robot{ .loc = Vec.of(0, 0), .dir = Dir.up() };

    while (vm.state != .done) {
        const current_color = hull.colorAt(robot.loc);
        vm.input(@intFromEnum(current_color));
        const next_color = vm.output();
        try hull.paint(robot.loc, @enumFromInt(next_color));

        const turn = vm.output();
        if (turn == 0) {
            robot.dir.left();
        } else {
            robot.dir.right();
        }
        const delta = robot.dir.vec();
        _ = robot.loc.mutAdd(@constCast(&delta));
    }

    const width: usize = @intCast(hull.max.x - hull.min.x + 1);
    const height: usize = @intCast(hull.max.y - hull.min.y + 1);

    var outpt_arr = std.ArrayList(u8).init(allocator);
    defer outpt_arr.deinit();

    try outpt_arr.append('\n');

    for (0..height) |y| {
        for (0..width) |x| {
            const pt = Vec.of(@intCast(x), @intCast(y));
            if (hull.colorAt(hull.min.add(@constCast(&pt))) == .white) {
                try outpt_arr.append('#');
            } else {
                try outpt_arr.append(' ');
            }
        }
        try outpt_arr.append('\n');
    }

    return try outpt_arr.toOwnedSliceSentinel(0);
}

const Color = enum {
    black,
    white,
};

const Dir = struct {
    // north = 0,
    // east = 1,
    // south = 2,
    // west = 3,
    dir: u2,

    const Self = @This();
    const vecs = [_]Vec{
        Vec.of(0, -1), // north
        Vec.of(1, 0), // east
        Vec.of(0, 1), // south
        Vec.of(-1, 0), // west
    };

    pub fn up() Self {
        return .{ .dir = 0 };
    }

    pub fn left(self: *Self) void {
        self.dir = @addWithOverflow(self.dir, 3)[0];
    }

    pub fn right(self: *Self) void {
        self.dir = @addWithOverflow(self.dir, 1)[0];
    }

    pub fn vec(self: *const Self) Vec {
        return vecs[@intCast(self.dir)];
    }
};

const Vec = Vec2(isize);

const Hull = struct {
    min: Vec,
    max: Vec,
    plates: Map,
    allocator: std.mem.Allocator,

    const Self = @This();
    const Map = std.AutoHashMap(Vec, Color);

    pub fn init(allocator: std.mem.Allocator) Self {
        return .{
            .min = Vec.of(std.math.maxInt(isize), std.math.maxInt(isize)),
            .max = Vec.of(std.math.minInt(isize), std.math.minInt(isize)),
            .plates = Map.init(allocator),
            .allocator = allocator,
        };
    }

    pub fn deinit(self: *Self) void {
        self.plates.deinit();
    }

    pub fn colorAt(self: *const Self, loc: Vec) Color {
        return self.plates.get(loc) orelse .black;
    }

    pub fn paint(self: *Self, loc: Vec, color: Color) !void {
        try self.plates.put(loc, color);
        self.min.x = @min(self.min.x, loc.x);
        self.max.x = @max(self.max.x, loc.x);
        self.min.y = @min(self.min.y, loc.y);
        self.max.y = @max(self.max.y, loc.y);
    }
};

const Robot = struct {
    loc: Vec,
    dir: Dir,
};
