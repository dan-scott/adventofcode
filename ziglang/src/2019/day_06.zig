const std = @import("std");
const util = @import("../utils.zig");

pub const year = 2019;
pub const day = 6;

pub fn part1solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var orbits = try parseInput(allocator, input);
    defer orbits.deinit();

    const total = try countOrbits("COM", &orbits, 0);

    return std.fmt.allocPrint(allocator, "{}", .{total});
}

fn countOrbits(obj: []const u8, orbit_map: *const OrbitMap, depth: usize) !usize {
    var total: usize = 0;
    const maybe_children = orbit_map.children.get(obj);
    if (maybe_children) |children| {
        for (children.items) |child| {
            total += try countOrbits(child, orbit_map, depth + 1);
        }
    }

    return depth + total;
}

pub fn part2solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var orbits = try parseInput(allocator, input);
    defer orbits.deinit();

    var you_idx: usize = 0;
    var santa_idx: usize = 0;

    var you_iter = orbits.parentIter("YOU");
    you: while (you_iter.next()) |you_parent| {
        var santa_iter = orbits.parentIter("SAN");
        santa_idx = 0;
        while (santa_iter.next()) |santa_parent| {
            if (std.mem.eql(u8, you_parent, santa_parent)) {
                break :you;
            }
            santa_idx += 1;
        }
        you_idx += 1;
    }

    return std.fmt.allocPrint(allocator, "{}", .{you_idx + santa_idx});
}

fn parseInput(alloc: std.mem.Allocator, input: []const u8) !OrbitMap {
    var orbits = try OrbitMap.init(alloc);

    var line_iter = std.mem.splitSequence(u8, input, "\n");
    var idx: usize = 0;
    while (line_iter.next()) |line| {
        var parts = std.mem.splitScalar(u8, line, ')');
        const parent = parts.next() orelse unreachable;
        const child = parts.next() orelse unreachable;
        try orbits.addPair(parent, child);
        idx += 1;
    }

    return orbits;
}

const OrbitMap = struct {
    const ChildList = std.ArrayList([]const u8);
    const ParentMap = std.hash_map.StringHashMap([]const u8);
    const ChildMap = std.hash_map.StringHashMap(ChildList);
    const Self = @This();

    const ParentIter = struct {
        current: []const u8,
        map: *Self,

        pub fn next(self: *ParentIter) ?[]const u8 {
            const maybe_next = self.map.parent.get(self.current);
            if (maybe_next) |n| {
                self.current = n;
                return self.current;
            } else {
                return null;
            }
        }
    };

    parent: std.hash_map.StringHashMap([]const u8),
    children: std.hash_map.StringHashMap(ChildList),
    allocator: std.mem.Allocator,

    pub fn init(allocator: std.mem.Allocator) !Self {
        return Self{
            .allocator = allocator,
            .parent = ParentMap.init(allocator),
            .children = ChildMap.init(allocator),
        };
    }

    pub fn deinit(self: *Self) void {
        self.parent.deinit();
        var children_val_iter = self.children.valueIterator();
        while (children_val_iter.next()) |children| {
            children.deinit();
        }
        self.children.deinit();
    }

    pub fn addPair(self: *Self, parent: []const u8, child: []const u8) !void {
        try self.setParent(child, parent);
        try self.addChild(parent, child);
    }

    fn setParent(self: *Self, name: []const u8, parent: []const u8) !void {
        try self.parent.put(name, parent);
    }

    fn addChild(self: *Self, name: []const u8, child: []const u8) !void {
        var poa = try self.children.getOrPut(name);
        if (!poa.found_existing) {
            poa.value_ptr.* = ChildList.init(self.allocator);
        }
        try poa.value_ptr.append(child);
    }

    fn parentIter(self: *Self, start: []const u8) Self.ParentIter {
        return .{
            .current = start,
            .map = self,
        };
    }
};

test "part1Solver" {
    const input =
        \\COM)B
        \\B)C
        \\C)D
        \\D)E
        \\E)F
        \\B)G
        \\G)H
        \\D)I
        \\E)J
        \\J)K
        \\K)L
    ;
    const expected = "42";

    const allocator = std.testing.allocator;

    const actual = try part1solver(allocator, input);
    defer allocator.free(actual);
    try std.testing.expectEqualStrings(expected, actual);
}

test "part2Solver" {
    const input =
        \\COM)B
        \\B)C
        \\C)D
        \\D)E
        \\E)F
        \\B)G
        \\G)H
        \\D)I
        \\E)J
        \\J)K
        \\K)L
        \\K)YOU
        \\I)SAN
    ;
    const expected = "4";

    const allocator = std.testing.allocator;

    const actual = try part2solver(allocator, input);
    defer allocator.free(actual);
    try std.testing.expectEqualStrings(expected, actual);
}
