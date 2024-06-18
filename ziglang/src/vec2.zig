const std = @import("std");

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
            // Euclid's algo to find GCD
            var n: isize = @intCast(@abs(grad.x));
            var d: isize = @intCast(@abs(grad.y));
            while (n != d) {
                if (n > d) {
                    n -= d;
                } else {
                    d -= n;
                }
            }

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
