const std = @import("std");
const util = @import("../utils.zig");
const IntcodeVM = @import("./int_code_vm.zig").IntcodeVM;

pub const year = 2019;
pub const day = 2;

pub fn part1solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var vm = try IntcodeVM.new(allocator, input);
    vm.mem[1] = 12;
    vm.mem[2] = 2;
    vm.run();
    return std.fmt.allocPrint(allocator, "{}", .{vm.mem[0]});
}

pub fn part2solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    for (0..100) |noun| {
        for (0..100) |verb| {
            var vm = try IntcodeVM.new(allocator, input);
            vm.mem[1] = @intCast(noun);
            vm.mem[2] = @intCast(verb);
            vm.run();
            if (vm.mem[0] == 19690720) {
                return std.fmt.allocPrint(allocator, "{}", .{100 * noun + verb});
            }
        }
    }
    return "Pair not found";
}
