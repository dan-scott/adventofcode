const std = @import("std");
const VM = @import("./int_code_vm.zig").VM;

pub const year = 2019;
pub const day = 7;

pub fn part1solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var base_vm = try VM.init(allocator, input);
    defer base_vm.deinit();
    base_vm.run();

    var max: isize = 0;

    for (0..5) |a_phase| {
        var a = try base_vm.clone();
        defer a.deinit();
        a.input(@intCast(a_phase));
        a.input(0);

        const a_output = a.output();

        for (0..5) |b_phase| {
            if (b_phase == a_phase) {
                continue;
            }
            var b = try base_vm.clone();
            defer b.deinit();
            b.input(@intCast(b_phase));
            b.input(a_output);

            const b_output = b.output();

            for (0..5) |c_phase| {
                if (c_phase == a_phase or c_phase == b_phase) {
                    continue;
                }
                var c = try base_vm.clone();
                defer c.deinit();
                c.input(@intCast(c_phase));
                c.input(b_output);

                const c_output = c.output();

                for (0..5) |d_phase| {
                    if (d_phase == a_phase or d_phase == b_phase or d_phase == c_phase) {
                        continue;
                    }
                    var d = try base_vm.clone();
                    defer d.deinit();
                    d.input(@intCast(d_phase));
                    d.input(c_output);

                    const d_output = d.output();

                    for (0..5) |e_phase| {
                        if (e_phase == a_phase or e_phase == b_phase or e_phase == c_phase or e_phase == d_phase) {
                            continue;
                        }
                        var e = try base_vm.clone();
                        defer e.deinit();
                        e.input(@intCast(e_phase));
                        e.input(d_output);

                        max = @max(max, e.output());
                    }
                }
            }
        }
    }

    return std.fmt.allocPrint(allocator, "{}", .{max});
}

pub fn part2solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var base_vm = try VM.init(allocator, input);
    defer base_vm.deinit();
    base_vm.run();

    var max: isize = 0;

    for (0..3125) |val| {
        if (phasesFromInt(val)) |phases| {
            const output = try runFeedback(&base_vm, phases);
            max = @max(max, output);
        }
    }

    return std.fmt.allocPrint(allocator, "{}", .{max});
}

pub fn phasesFromInt(val: usize) ?[5]usize {
    const phases = [5]usize{
        (val / 625) % 5 + 5,
        (val / 125) % 5 + 5,
        (val / 25) % 5 + 5,
        (val / 5) % 5 + 5,
        val % 5 + 5,
    };
    inline for (1..5) |i| {
        if (phases[0] == phases[i]) {
            return null;
        }
    }
    inline for (2..5) |i| {
        if (phases[1] == phases[i]) {
            return null;
        }
    }
    inline for (3..5) |i| {
        if (phases[2] == phases[i]) {
            return null;
        }
    }
    if (phases[3] == phases[4]) {
        return null;
    }
    return phases;
}

fn runFeedback(vm: *VM, phases: [5]usize) !isize {
    var input: isize = 0;
    var amps = [5]VM{
        try vm.clone(),
        try vm.clone(),
        try vm.clone(),
        try vm.clone(),
        try vm.clone(),
    };

    defer amps[0].deinit();
    defer amps[1].deinit();
    defer amps[2].deinit();
    defer amps[3].deinit();
    defer amps[4].deinit();

    inline for (phases, 0..) |phase, idx| {
        amps[idx].input(@intCast(phase));
    }

    var idx: usize = 0;
    var done = false;

    while (!done) {
        if (amps[idx].state == .done) {
            done = true;
        } else {
            amps[idx].input(input);
            input = amps[idx].output();
            idx = (idx + 1) % 5;
        }
    }

    return input;
}

test "part1Solver" {
    const cases = [_]struct { input: []const u8, expected: []const u8 }{
        .{
            .input = "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0",
            .expected = "43210",
        },
        .{
            .input = "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0",
            .expected = "54321",
        },
        .{
            .input = "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0",
            .expected = "65210",
        },
    };
    const allocator = std.testing.allocator;
    for (cases) |case| {
        const actual = try part1solver(allocator, case.input);
        defer allocator.free(actual);
        try std.testing.expectEqualStrings(case.expected, actual);
    }
}

test "part2Solver" {
    const cases = [_]struct { input: []const u8, expected: []const u8 }{
        .{
            .input = "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5",
            .expected = "139629729",
        },
        .{
            .input = "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10",
            .expected = "18216",
        },
    };
    const allocator = std.testing.allocator;
    for (cases) |case| {
        const actual = try part2solver(allocator, case.input);
        defer allocator.free(actual);
        try std.testing.expectEqualStrings(case.expected, actual);
    }
}
