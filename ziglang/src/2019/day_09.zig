const std = @import("std");
const VM = @import("./int_code_vm.zig").VM;

pub const year = 2019;
pub const day = 9;

pub fn part1solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var vm = try VM.init(allocator, input);
    defer vm.deinit();
    vm.run();
    vm.input(1);

    return std.fmt.allocPrint(allocator, "{any}", .{vm.output()});
}

pub fn part2solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var vm = try VM.init(allocator, input);
    defer vm.deinit();
    vm.run();
    vm.input(2);

    const output = vm.output();

    return std.fmt.allocPrint(allocator, "{}", .{output});
}
