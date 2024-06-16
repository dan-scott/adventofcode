const std = @import("std");
const util = @import("../utils.zig");
const IntcodeVM = @import("./int_code_vm.zig").IntcodeVM;

pub const year = 2019;
pub const day = 5;

pub fn part1solver(alloc: std.mem.Allocator, input: []const u8) ![]const u8 {
    var vm = try IntcodeVM.init(alloc, input);
    defer vm.deinit();
    vm.run();
    vm.input(1);
    var output = vm.output();
    while (vm.state != .done) {
        if (output != 0) {
            return error.TestProgramFailed;
        }
        output = vm.output();
    }

    return std.fmt.allocPrint(alloc, "{}", .{output});
}

pub fn part2solver(alloc: std.mem.Allocator, input: []const u8) ![]const u8 {
    var vm = try IntcodeVM.init(alloc, input);
    defer vm.deinit();
    vm.run();
    vm.input(5);
    const output = vm.output();

    return std.fmt.allocPrint(alloc, "{}", .{output});
}
