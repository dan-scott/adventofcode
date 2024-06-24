const std = @import("std");
const m = @import("./math.zig");

pub fn Vec2(comptime T: type) type {
    return struct {
        x: T,
        y: T,

        const Self = @This();

        pub fn of(x: T, y: T) Self {
            return .{ .x = x, .y = y };
        }

        pub fn add(self: *const Self, other: *const Self) Self {
            return .{
                .x = self.x + other.x,
                .y = self.y + other.y,
            };
        }

        pub fn sub(self: *const Self, other: *const Self) Self {
            return .{
                .x = self.x - other.x,
                .y = self.y - other.y,
            };
        }

        pub fn mutAdd(self: *Self, other: *const Self) *Self {
            self.x += other.x;
            self.y += other.y;
            return self;
        }

        pub fn gradient(self: *const Self, other: *const Self) Self {
            var grad = other.sub(self);
            if (grad.x == 0) {
                if (grad.y > 0) {
                    return Self.of(0, 1);
                } else {
                    return Self.of(0, -1);
                }
            }
            if (grad.y == 0) {
                if (grad.x > 0) {
                    return Self.of(1, 0);
                } else {
                    return Self.of(-1, 0);
                }
            }

            // Simplify the gradient
            const n: isize = m.gcd(grad.x, grad.y) catch unreachable;

            grad.x = @divTrunc(grad.x, n);
            grad.y = @divTrunc(grad.y, n);

            return grad;
        }

        pub fn angle(self: *const Self) f64 {
            const x: f64 = @floatFromInt(self.x);
            const y: f64 = @floatFromInt(self.y);
            var result: f64 = std.math.atan2(y, x) + (std.math.pi / @as(f64, 2));
            if (result < 0) {
                result += std.math.pi * 2;
            }
            return result;
        }

        pub fn dist(self: *const Self, other: *const Self) f64 {
            const x2: f64 = @floatFromInt(std.math.pow(T, other.x - self.x, 2));
            const y2: f64 = @floatFromInt(std.math.pow(T, other.y - self.y, 2));
            return @sqrt(x2 + y2);
        }
    };
}

pub fn Vec3(comptime T: type) type {
    return struct {
        x: T,
        y: T,
        z: T,

        const Self = @This();

        pub fn of(x: T, y: T, z: T) Self {
            return .{ .x = x, .y = y, .z = z };
        }

        pub fn zero() Self {
            return .{ .x = 0, .y = 0, .z = 0 };
        }

        pub fn add(self: *const Self, other: *const Self) Self {
            return .{
                .x = self.x + other.x,
                .y = self.y + other.y,
                .z = self.z + other.z,
            };
        }

        pub fn mutAdd(self: *Self, other: *const Self) void {
            self.x += other.x;
            self.y += other.y;
            self.z += other.z;
        }
    };
}
