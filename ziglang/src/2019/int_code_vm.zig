const std = @import("std");

pub const IntcodeVM = struct {
    pc: isize,
    mem: []isize,
    allocator: std.mem.Allocator,
    state: State,
    input_val: isize,
    output_val: isize,
    current_op: isize,
    param_mode: [3]isize,

    const Self = @This();

    const State = enum {
        run,
        await_input,
        process_input,
        pending_output,
        done,
    };

    pub fn new(allocator: std.mem.Allocator, program: []const u8) !Self {
        var tokens = std.mem.tokenizeScalar(u8, program, ',');
        const mem_size: usize = std.mem.count(u8, program, ",") + 1;
        var mem = try allocator.alloc(isize, mem_size * 10);
        tokens.reset();
        var pc: usize = 0;
        while (tokens.next()) |token| {
            mem[pc] = try std.fmt.parseInt(isize, token, 10);
            pc += 1;
        }
        return .{
            .pc = 0,
            .mem = mem,
            .allocator = allocator,
            .state = .run,
            .current_op = 0,
            .param_mode = .{ 0, 0, 0 },
            .input_val = 0,
            .output_val = 0,
        };
    }

    pub fn clone(self: *Self) !Self {
        const new_mem = try self.allocator.alloc(isize, self.mem.len);
        @memcpy(new_mem, self.mem);
        return .{
            .pc = self.pc,
            .mem = new_mem,
            .allocator = self.allocator,
            .state = self.state,
            .current_op = self.current_op,
            .param_mode = .{ self.param_mode[0], self.param_mode[1], self.param_mode[2] },
            .input_val = self.input_val,
            .output_val = self.output_val,
        };
    }

    pub fn input(self: *Self, val: isize) void {
        if (self.state == .await_input) {
            std.log.debug("Getting input\n", .{});
            self.state = .process_input;
            self.input_val = val;
            self.execOp();
        }

        self.run();
    }

    pub fn output(self: *Self) isize {
        defer self.run();
        if (self.state == .pending_output) {
            self.state = .run;
            return self.output_val;
        }
        return self.mem[0];
    }

    pub fn deinit(self: *Self) void {
        self.allocator.free(self.mem);
    }

    pub fn run(self: *Self) void {
        while (self.state == .run) {
            self.readOp();
            self.execOp();
        }
    }

    fn execOp(self: *Self) void {
        const inc = switch (self.current_op) {
            1 => self.add(),
            2 => self.mul(),
            3 => self.evalInput(),
            4 => self.evalOutput(),
            5 => self.jft(),
            6 => self.jff(),
            7 => self.lt(),
            8 => self.eq(),

            99 => self.done(),

            else => {
                std.log.debug("Unknown opcode {}\n", .{self.current_op});
                unreachable;
            },
        };

        self.pc += inc;
    }

    fn readOp(self: *Self) void {
        const code = self.getMemAt(self.pc);
        self.current_op = @mod(code, 100);
        self.param_mode[0] = @mod(@divTrunc(code, 100), 10);
        self.param_mode[1] = @mod(@divTrunc(code, 1_000), 10);
        self.param_mode[2] = @mod(@divTrunc(code, 10_000), 10);
    }

    fn add(self: *Self) isize {
        const a = self.getParam(1);
        const b = self.getParam(2);
        const addr = self.getMemAt(self.pc + 3);
        std.log.debug("Add: {} := {} + {}\n", .{ addr, a, b });
        self.setMem(addr, a + b);
        return 4;
    }

    fn mul(self: *Self) isize {
        const a = self.getParam(1);
        const b = self.getParam(2);
        const addr = self.getMemAt(self.pc + 3);
        std.log.debug("Mul: {} := {} * {}\n", .{ addr, a, b });
        self.setMem(addr, a * b);
        return 4;
    }

    fn evalInput(self: *Self) isize {
        if (self.state == .run) {
            std.log.debug("Await input\n", .{});
            self.state = .await_input;
            return 0;
        } else if (self.state == .process_input) {
            const addr = self.getMemAt(self.pc + 1);
            std.log.debug("Save input: {} := {}\n", .{ addr, self.input_val });
            self.setMem(addr, self.input_val);
            self.state = .run;
            return 2;
        } else {
            unreachable;
        }
    }

    fn evalOutput(self: *Self) isize {
        self.state = .pending_output;
        self.output_val = self.getParam(1);
        std.log.debug("Send output: {}\n", .{self.output_val});
        return 2;
    }

    fn jft(self: *Self) isize {
        const chk = self.getParam(1);
        const store = self.getParam(2);
        std.log.debug("IF {} != 0 THEN pc = {} ({})\n", .{ chk, store, chk != 0 });
        if (chk != 0) {
            return store - self.pc;
        }
        return 3;
    }

    fn jff(self: *Self) isize {
        const chk = self.getParam(1);
        const store = self.getParam(2);
        std.log.debug("IF {} == 0 THEN pc = {} ({})\n", .{ chk, store, chk == 0 });
        if (chk == 0) {
            return store - self.pc;
        }
        return 3;
    }

    fn lt(self: *Self) isize {
        const a = self.getParam(1);
        const b = self.getParam(2);
        const addr = self.getMemAt(self.pc + 3);
        std.log.debug("IF {0} < {1} THEN {2} := 1 ELSE {2} := 0 ({3})\n", .{ a, b, addr, a < b });
        if (a < b) {
            self.setMem(addr, 1);
        } else {
            self.setMem(addr, 0);
        }
        return 4;
    }

    fn eq(self: *Self) isize {
        const a = self.getParam(1);
        const b = self.getParam(2);
        const addr = self.getMemAt(self.pc + 3);
        std.log.debug("IF {0} == {1} THEN {2} := 1 ELSE {2} := 0 ({3})\n", .{ a, b, addr, a == b });
        if (a == b) {
            self.setMem(addr, 1);
        } else {
            self.setMem(addr, 0);
        }
        return 4;
    }

    fn done(self: *Self) isize {
        std.log.debug("Done\n", .{});
        self.state = .done;
        return 1;
    }

    fn getMemAt(self: *Self, addr: isize) isize {
        return self.mem[@intCast(addr)];
    }

    fn getParam(self: *Self, nr: isize) isize {
        return switch (self.param_mode[@intCast(nr - 1)]) {
            0 => self.getMemAt(self.getMemAt(self.pc + nr)),
            1 => self.getMemAt(self.pc + nr),
            else => unreachable,
        };
    }

    fn setMem(self: *Self, addr: isize, val: isize) void {
        self.mem[@intCast(addr)] = val;
    }
};

test "Intcode VM day 2 spec" {
    const cases = [_]struct { input: []const u8, expected: []const u8, len: usize }{
        .{
            .input = "1,9,10,3,2,3,11,0,99,30,40,50",
            .expected = "{ 3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50 }",
            .len = 12,
        },
        .{
            .input = "1,0,0,0,99",
            .expected = "{ 2, 0, 0, 0, 99 }",
            .len = 5,
        },
    };
    const alloc = std.testing.allocator;
    for (cases) |case| {
        var vm = try IntcodeVM.new(alloc, case.input);
        defer vm.deinit();
        vm.run();
        const actual = try std.fmt.allocPrint(alloc, "{any}", .{vm.mem[0..case.len]});
        defer alloc.free(actual);
        try std.testing.expectEqualStrings(case.expected, actual);
    }
}

test "Intcode VM day 5 spec: ops 3 and 4" {
    const input = "3,5,4,5,99,0";
    const alloc = std.testing.allocator;
    var vm = try IntcodeVM.new(alloc, input);
    defer vm.deinit();
    vm.run();
    vm.input(20);
    const output = vm.output();
    try std.testing.expectEqual(20, output);
}

test "Intcode VM day 5 spec: ops 5, 6, 7, 8" {
    const cases = [_]struct { code: []const u8, input: isize, output: isize }{

        // Equals, position mode
        .{ .code = "3,9,8,9,10,9,4,9,99,-1,8", .input = 5, .output = 0 },
        .{ .code = "3,9,8,9,10,9,4,9,99,-1,8", .input = 8, .output = 1 },
        // Equals, immediate mode
        .{ .code = "3,3,1108,-1,8,3,4,3,99", .input = 5, .output = 0 },
        .{ .code = "3,3,1108,-1,8,3,4,3,99", .input = 8, .output = 1 },

        // Less than, position mode
        .{ .code = "3,9,7,9,10,9,4,9,99,-1,8", .input = 9, .output = 0 },
        .{ .code = "3,9,7,9,10,9,4,9,99,-1,8", .input = 7, .output = 1 },
        // Less than, immediate mode
        .{ .code = "3,3,1107,-1,8,3,4,3,99", .input = 9, .output = 0 },
        .{ .code = "3,3,1107,-1,8,3,4,3,99", .input = 7, .output = 1 },

        // Jump, position mode
        .{ .code = "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", .input = 0, .output = 0 },
        .{ .code = "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", .input = 12, .output = 1 },
        // // Jump, immediate mode
        // .{ .code = "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", .input = 0, .output = 0 },
        // .{ .code = "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", .input = 12, .output = 1 },
    };

    const alloc = std.testing.allocator;
    for (cases) |case| {
        var vm = try IntcodeVM.new(alloc, case.code);
        defer vm.deinit();
        vm.run();
        vm.input(case.input);
        const actual = vm.output();
        try std.testing.expectEqual(case.output, actual);
    }
}
