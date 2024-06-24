const std = @import("std");
const util = @import("../utils.zig");

pub const year = 2019;
pub const day = 12;

pub fn part1solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    const energy = try calcTotalEnergy(input, 1000);

    return std.fmt.allocPrint(allocator, "{}", .{energy});
}

fn calcTotalEnergy(input: []const u8, steps: usize) !isize {
    var moons = try parseInput(input);

    for (0..steps) |_| {
        for (0..3) |a| {
            for ((a + 1)..4) |b| {
                var moon_a = &moons[a];
                var moon_b = &moons[b];

                const x_diff = deltas(moon_a.pos.x, moon_b.pos.x);
                const y_diff = deltas(moon_a.pos.y, moon_b.pos.y);
                const z_diff = deltas(moon_a.pos.z, moon_b.pos.z);

                moon_a.vel.x += x_diff[0];
                moon_a.vel.y += y_diff[0];
                moon_a.vel.z += z_diff[0];

                moon_b.vel.x += x_diff[1];
                moon_b.vel.y += y_diff[1];
                moon_b.vel.z += z_diff[1];
            }
        }

        for (0..4) |idx| {
            _ = moons[idx].pos.mutAdd(@constCast(&moons[idx].vel));
        }
    }

    var energy: isize = 0;
    for (moons) |moon| {
        energy += moon.energy();
    }

    return energy;
}

fn deltas(a: isize, b: isize) [2]isize {
    if (a > b) {
        return .{ -1, 1 };
    } else if (a < b) {
        return .{ 1, -1 };
    } else {
        return .{ 0, 0 };
    }
}

pub fn part2solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    const moons = try parseInput(input);

    const moons_x = [_]Moon1D{
        Moon1D.of(moons[0].pos.x, moons[0].vel.x),
        Moon1D.of(moons[1].pos.x, moons[1].vel.x),
        Moon1D.of(moons[2].pos.x, moons[2].vel.x),
        Moon1D.of(moons[3].pos.x, moons[3].vel.x),
    };

    const moons_y = [_]Moon1D{
        Moon1D.of(moons[0].pos.y, moons[0].vel.y),
        Moon1D.of(moons[1].pos.y, moons[1].vel.y),
        Moon1D.of(moons[2].pos.y, moons[2].vel.y),
        Moon1D.of(moons[3].pos.y, moons[3].vel.y),
    };

    const moons_z = [_]Moon1D{
        Moon1D.of(moons[0].pos.z, moons[0].vel.z),
        Moon1D.of(moons[1].pos.z, moons[1].vel.z),
        Moon1D.of(moons[2].pos.z, moons[2].vel.z),
        Moon1D.of(moons[3].pos.z, moons[3].vel.z),
    };

    const x_period = findPeriod(moons_x);
    const y_period = findPeriod(moons_y);
    const z_peroid = findPeriod(moons_z);

    const xy = @divTrunc(x_period * y_period, util.math.gcd(x_period, y_period) catch unreachable);

    const result = @divTrunc(xy * z_peroid, util.math.gcd(xy, z_peroid) catch unreachable);

    return std.fmt.allocPrint(allocator, "{}", .{result});
}

fn findPeriod(moons: [4]Moon1D) usize {
    var cpy = moons;
    var found = false;
    var ct: usize = 0;

    while (!found) {
        ct += 1;
        for (0..4) |a| {
            for (0..4) |b| {
                if (a == b) {
                    continue;
                } else if (cpy[a].x < cpy[b].x) {
                    cpy[a].y += 1;
                } else if (cpy[a].x > cpy[b].x) {
                    cpy[a].y -= 1;
                }
            }
        }
        for (0..4) |i| {
            cpy[i].x += cpy[i].y;
        }

        if (std.meta.eql(moons, cpy)) {
            found = true;
        }
    }

    return ct;
}

const Vec = util.vec.Vec3(isize);

const Moon1D = util.vec.Vec2(isize);

const Moon = struct {
    pos: Vec,
    vel: Vec,

    const Self = @This();

    pub fn init(line: []const u8) !Self {
        var parts = std.mem.splitSequence(u8, line[1..(line.len - 1)], ", ");
        const x = try getVal(parts.next() orelse unreachable);
        const y = try getVal(parts.next() orelse unreachable);
        const z = try getVal(parts.next() orelse unreachable);

        return .{
            .pos = Vec.of(x, y, z),
            .vel = Vec.zero(),
        };
    }

    fn getVal(str: []const u8) !isize {
        return try std.fmt.parseInt(isize, str[2..], 10);
    }

    fn energy(self: *const Self) isize {
        const pot = absSum(self.pos);
        const kin = absSum(self.vel);

        return pot * kin;
    }

    fn absSum(v: Vec) isize {
        const x: isize = @intCast(@abs(v.x));
        const y: isize = @intCast(@abs(v.y));
        const z: isize = @intCast(@abs(v.z));
        return x + y + z;
    }
};

pub fn parseInput(input: []const u8) ![4]Moon {
    var parts = std.mem.splitScalar(u8, input, '\n');
    return .{
        try Moon.init(parts.next() orelse unreachable),
        try Moon.init(parts.next() orelse unreachable),
        try Moon.init(parts.next() orelse unreachable),
        try Moon.init(parts.next() orelse unreachable),
    };
}

test "part1Solver" {
    const cases = [_]struct { input: []const u8, steps: usize, expected: isize }{
        .{
            .input =
            \\<x=-1, y=0, z=2>
            \\<x=2, y=-10, z=-7>
            \\<x=4, y=-8, z=8>
            \\<x=3, y=5, z=-1>
            ,
            .steps = 10,
            .expected = 179,
        },
        .{
            .input =
            \\<x=-8, y=-10, z=0>
            \\<x=5, y=5, z=10>
            \\<x=2, y=-7, z=3>
            \\<x=9, y=-8, z=-3>
            ,
            .steps = 100,
            .expected = 1940,
        },
    };
    for (cases) |case| {
        const actual = try calcTotalEnergy(case.input, case.steps);
        try std.testing.expectEqual(case.expected, actual);
    }
}

test "part2Solver" {
    const cases = [_]struct { input: []const u8, expected: []const u8 }{
        .{
            .input =
            \\<x=-1, y=0, z=2>
            \\<x=2, y=-10, z=-7>
            \\<x=4, y=-8, z=8>
            \\<x=3, y=5, z=-1>
            ,
            .expected = "2772",
        },
        .{
            .input =
            \\<x=-8, y=-10, z=0>
            \\<x=5, y=5, z=10>
            \\<x=2, y=-7, z=3>
            \\<x=9, y=-8, z=-3>
            ,
            .expected = "4686774924",
        },
    };
    const allocator = std.testing.allocator;
    for (cases) |case| {
        const actual = try part2solver(allocator, case.input);
        defer allocator.free(actual);
        try std.testing.expectEqualStrings(case.expected, actual);
    }
}
