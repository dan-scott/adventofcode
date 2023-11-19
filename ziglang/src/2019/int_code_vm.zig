const std = @import("std");

pub const IntcodeVM = struct {
    pc: usize,
    mem: []usize,
    allocator: std.mem.Allocator,

    const Self = @This();

    pub fn new(allocator: std.mem.Allocator, program: []const u8) !Self {
        var tokens = std.mem.tokenizeScalar(u8, program, ',');
        var memSize: usize = 0;
        while (tokens.next()) |_| {
            memSize += 1;
        }
        var mem = try allocator.alloc(usize, memSize);
        tokens.reset();
        var pc: usize = 0;
        while (tokens.next()) |token| {
            mem[pc] = try std.fmt.parseInt(usize, token, 10);
            pc += 1;
        }
        return .{
            .pc = 0,
            .mem = mem,
            .allocator = allocator,
        };
    }

    pub fn deinit(self: *Self) void {
        self.allocator.free(self.mem);
    }

    pub fn run(self: *Self) void {
        while (self.mem[self.pc] != 99) {
            self.execCmd();
        }
    }

    fn execCmd(self: *Self) void {
        const a = self.mem[self.mem[self.pc + 1]];
        const b = self.mem[self.mem[self.pc + 2]];
        const store = self.mem[self.pc + 3];
        switch (self.mem[self.pc]) {
            1 => {
                self.mem[store] = a + b;
            },
            2 => {
                self.mem[store] = a * b;
            },
            else => {},
        }
        self.pc += 4;
    }
};

test "Intcode VM day 2 spec" {
    const cases = [_]struct { input: []const u8, expected: []const u8 }{
        .{
            .input = "1,9,10,3,2,3,11,0,99,30,40,50",
            .expected = "{ 3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50 }",
        },
        .{
            .input = "1,0,0,0,99",
            .expected = "{ 2, 0, 0, 0, 99 }",
        },
    };
    const alloc = std.testing.allocator;
    for (cases) |case| {
        var vm = try IntcodeVM.new(alloc, case.input);
        defer vm.deinit();
        vm.run();
        const actual = try std.fmt.allocPrint(alloc, "{any}", .{vm.mem});
        defer alloc.free(actual);
        try std.testing.expectEqualStrings(case.expected, actual);
    }
}
