const std = @import("std");
const util = @import("../util.zig");
const VM = @import("./int_code_vm.zig").VM;

pub const year = 2019;
pub const day = 13;

pub fn part1solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var vm = try VM.init(allocator, input);
    defer vm.deinit();
    vm.run();

    var ct: usize = 0;
    while (vm.state != .done) {
        _ = vm.output();
        _ = vm.output();
        const tile: Tile = @enumFromInt(vm.output());
        if (tile == .block) {
            ct += 1;
        }
    }

    return std.fmt.allocPrint(allocator, "{}", .{ct});
}

const Vec = util.vec.Vec2(isize);

pub fn part2solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var vm = try VM.init(allocator, input);
    defer vm.deinit();
    vm.setMem(0, 2);
    vm.run();

    var ball_x: isize = 0;
    var paddle_x: isize = 0;

    var x: isize = 0;
    var y: isize = 0;
    var t: isize = 0;
    var score: isize = 0;

    while (vm.state != .done) {
        if (vm.state == .await_input) {
            if (ball_x < paddle_x) {
                vm.input(-1);
            } else if (ball_x > paddle_x) {
                vm.input(1);
            } else {
                vm.input(0);
            }
        }

        x = vm.output();
        y = vm.output();
        t = vm.output();

        if (x == -1 and y == 0) {
            score = t;
        } else if (t == 4) {
            ball_x = x;
        } else if (t == 3) {
            paddle_x = x;
        }
    }

    return std.fmt.allocPrint(allocator, "{}", .{score});
}

const Tile = enum {
    empty,
    wall,
    block,
    horiz_paddle,
    ball,
};
