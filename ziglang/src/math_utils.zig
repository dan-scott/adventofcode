const std = @import("std");

pub fn gcd(a: anytype, b: @TypeOf(a)) error{InvalidZeroParam}!@TypeOf(a) {
    if (a == 0 or b == 0) {
        return error.InvalidZeroParam;
    }

    var n: @TypeOf(a) = @intCast(@abs(a));
    var d: @TypeOf(a) = @intCast(@abs(b));
    while (n != d) {
        if (n > d) {
            n -= d;
        } else {
            d -= n;
        }
    }

    return n;
}

test "gcd: isize" {
    const a: isize = -250;
    const b: isize = 325;

    const expected: isize = 25;
    const actual = try gcd(a, b);

    try std.testing.expectEqual(expected, actual);
}

test "gcd: usize" {
    const a: usize = 250;
    const b: usize = 325;

    const expected: usize = 25;
    const actual = try gcd(a, b);

    try std.testing.expectEqual(expected, actual);
}
