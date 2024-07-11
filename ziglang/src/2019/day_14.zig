const std = @import("std");
const util = @import("../util.zig");

pub const year = 2019;
pub const day = 14;

pub fn part1solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var reactor = try Reactor.init(allocator, input);
    defer reactor.deinit();

    const fuel = reactor.reaction_map.get("FUEL") orelse unreachable;

    var inpt = InputTracker.init(allocator);
    defer inpt.deinit();

    try inpt.seed(fuel, 1);

    const ore = inpt.required.get("ORE") orelse unreachable;

    return std.fmt.allocPrint(allocator, "{}", .{ore});
}

pub fn part2solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    var reactor = try Reactor.init(allocator, input);
    defer reactor.deinit();

    const fuel = reactor.reaction_map.get("FUEL") orelse unreachable;

    var upper_bound: u128 = 1_000_000_000_000;
    var lower_bound: u128 = 0;
    var midpoint: u128 = upper_bound / 2;
    var prev = lower_bound;

    const ore = 1_000_000_000_000;

    var found = false;

    while (!found) {
        prev = midpoint;

        var inpt = InputTracker.init(allocator);
        defer inpt.deinit();

        try inpt.seed(fuel, midpoint);

        const required_ore = inpt.required.get("ORE") orelse unreachable;

        if (required_ore > ore) {
            if (midpoint == upper_bound or midpoint == lower_bound) {
                midpoint = prev;
                found = true;
            }
            upper_bound = midpoint;
            midpoint = (upper_bound - lower_bound) / 2 + lower_bound;
        }

        if (required_ore < ore) {
            if (midpoint == upper_bound or midpoint == lower_bound) {
                found = true;
            } else {
                lower_bound = midpoint;
                midpoint = (upper_bound - lower_bound) / 2 + lower_bound;
            }
        }
    }

    return std.fmt.allocPrint(allocator, "{}", .{midpoint});
}

const Reaction = struct {
    name: []const u8,
    amount: u128,
    inputs: InputList,

    const InputList = std.ArrayList(ReactionInput);
};

const ReactionInput = struct {
    amount: u128,
    reaction: *Reaction,
};

const Reactor = struct {
    arena: *std.heap.ArenaAllocator,
    allocator: std.mem.Allocator,
    reaction_map: ReactionMap,

    const Self = @This();
    const ReactionMap = std.StringHashMap(*Reaction);

    pub fn init(allocator: std.mem.Allocator, input: []const u8) !Self {
        const arena = try allocator.create(std.heap.ArenaAllocator);
        arena.* = std.heap.ArenaAllocator.init(allocator);
        const alloc = arena.allocator();
        var reactor = Self{
            .arena = arena,
            .allocator = alloc,
            .reaction_map = ReactionMap.init(alloc),
        };

        var lines = std.mem.tokenizeScalar(u8, input, '\n');
        while (lines.next()) |line| {
            try reactor.addReaction(line);
        }

        return reactor;
    }

    pub fn deinit(self: *Self) void {
        self.arena.deinit();
        self.arena.child_allocator.destroy(self.arena);
    }

    fn addReaction(self: *Self, line: []const u8) !void {
        var sides = std.mem.tokenizeSequence(u8, line, " => ");

        const lhs = sides.next() orelse unreachable;
        const rhs = sides.next() orelse unreachable;

        const rhs_chem = try parseChemical(rhs);

        const reaction: *Reaction = blk: {
            const reaction_entry = try self.reaction_map.getOrPut(rhs_chem.name);
            if (reaction_entry.found_existing) {
                reaction_entry.value_ptr.*.amount = rhs_chem.amount;
                break :blk reaction_entry.value_ptr.*;
            } else {
                const nr = try self.allocator.create(Reaction);
                nr.* = .{
                    .amount = rhs_chem.amount,
                    .name = rhs_chem.name,
                    .inputs = Reaction.InputList.init(self.allocator),
                };
                reaction_entry.value_ptr.* = nr;
                break :blk nr;
            }
        };

        var lhs_parts = std.mem.tokenizeSequence(u8, lhs, ", ");
        while (lhs_parts.next()) |part| {
            const chem = try parseChemical(part);

            const input: *Reaction = blk: {
                const input_reaction = try self.reaction_map.getOrPut(chem.name);
                if (input_reaction.found_existing) {
                    break :blk input_reaction.value_ptr.*;
                } else {
                    const nr = try self.allocator.create(Reaction);
                    nr.* = .{
                        .name = chem.name,
                        .amount = 1,
                        .inputs = Reaction.InputList.init(self.allocator),
                    };
                    input_reaction.value_ptr.* = nr;
                    break :blk nr;
                }
            };

            try reaction.inputs.append(.{
                .amount = chem.amount,
                .reaction = input,
            });
        }
    }

    fn parseChemical(part: []const u8) !struct { amount: u128, name: []const u8 } {
        var parts = std.mem.splitScalar(u8, part, ' ');

        const amount = try std.fmt.parseUnsigned(u128, parts.next() orelse unreachable, 10);
        const name = parts.next() orelse unreachable;

        return .{
            .amount = amount,
            .name = name,
        };
    }
};

const InputTracker = struct {
    required: Map,
    extra: Map,

    const Self = @This();
    const Map = std.StringHashMap(u128);

    pub fn init(allocator: std.mem.Allocator) Self {
        return .{
            .required = Map.init(allocator),
            .extra = Map.init(allocator),
        };
    }

    pub fn deinit(self: *Self) void {
        self.extra.deinit();
        self.required.deinit();
    }

    pub fn seed(self: *Self, reaction: *Reaction, required_output: u128) !void {
        var req = required_output;
        const maybe_stock = try self.extra.getOrPut(reaction.name);
        if (!maybe_stock.found_existing) {
            maybe_stock.value_ptr.* = 0;
        } else if (maybe_stock.value_ptr.* > required_output) {
            maybe_stock.value_ptr.* -= required_output;
            return;
        } else {
            req -= maybe_stock.value_ptr.*;
            maybe_stock.value_ptr.* = 0;
        }

        const f_req: f128 = @floatFromInt(req);
        const f_ct: f128 = @floatFromInt(reaction.amount);
        const ratio: u128 = @intFromFloat(@ceil(f_req / f_ct));
        const output = reaction.amount * ratio;

        maybe_stock.value_ptr.* += output - req;

        const maybe_req = try self.required.getOrPut(reaction.name);
        if (!maybe_req.found_existing) {
            maybe_req.value_ptr.* = req;
        } else {
            maybe_req.value_ptr.* += req;
        }

        for (reaction.inputs.items) |input| {
            try self.seed(input.reaction, input.amount * ratio);
        }
    }
};

test "part1Solver" {
    const cases = [_]struct { input: []const u8, expected: []const u8 }{
        .{
            .input =
            \\10 ORE => 10 A
            \\1 ORE => 1 B
            \\7 A, 1 B => 1 C
            \\7 A, 1 C => 1 D
            \\7 A, 1 D => 1 E
            \\7 A, 1 E => 1 FUEL
            ,
            .expected = "31",
        },
        .{
            .input =
            \\9 ORE => 2 A
            \\8 ORE => 3 B
            \\7 ORE => 5 C
            \\3 A, 4 B => 1 AB
            \\5 B, 7 C => 1 BC
            \\4 C, 1 A => 1 CA
            \\2 AB, 3 BC, 4 CA => 1 FUEL
            ,
            .expected = "165",
        },
        .{
            .input =
            \\157 ORE => 5 NZVS
            \\165 ORE => 6 DCFZ
            \\44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
            \\12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
            \\179 ORE => 7 PSHF
            \\177 ORE => 5 HKGWZ
            \\7 DCFZ, 7 PSHF => 2 XJWVT
            \\165 ORE => 2 GPVTF
            \\3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT
            ,
            .expected = "13312",
        },
        .{
            .input =
            \\2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG
            \\17 NVRVD, 3 JNWZP => 8 VPVL
            \\53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL
            \\22 VJHF, 37 MNCFX => 5 FWMGM
            \\139 ORE => 4 NVRVD
            \\144 ORE => 7 JNWZP
            \\5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC
            \\5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV
            \\145 ORE => 6 MNCFX
            \\1 NVRVD => 8 CXFTF
            \\1 VJHF, 6 MNCFX => 4 RFSQX
            \\176 ORE => 6 VJHF
            ,
            .expected = "180697",
        },
        .{
            .input =
            \\171 ORE => 8 CNZTR
            \\7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
            \\114 ORE => 4 BHXH
            \\14 VRPVC => 6 BMBT
            \\6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
            \\6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
            \\15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
            \\13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
            \\5 BMBT => 4 WPTQ
            \\189 ORE => 9 KTJDG
            \\1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
            \\12 VRPVC, 27 CNZTR => 2 XDBXC
            \\15 KTJDG, 12 BHXH => 5 XCVML
            \\3 BHXH, 2 VRPVC => 7 MZWV
            \\121 ORE => 7 VRPVC
            \\7 XCVML => 6 RJRHP
            \\5 BHXH, 4 VRPVC => 5 LTCX
            ,
            .expected = "2210736",
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
            .input =
            \\157 ORE => 5 NZVS
            \\165 ORE => 6 DCFZ
            \\44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
            \\12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
            \\179 ORE => 7 PSHF
            \\177 ORE => 5 HKGWZ
            \\7 DCFZ, 7 PSHF => 2 XJWVT
            \\165 ORE => 2 GPVTF
            \\3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT
            ,
            .expected = "82892753",
        },
        .{
            .input =
            \\2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG
            \\17 NVRVD, 3 JNWZP => 8 VPVL
            \\53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL
            \\22 VJHF, 37 MNCFX => 5 FWMGM
            \\139 ORE => 4 NVRVD
            \\144 ORE => 7 JNWZP
            \\5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC
            \\5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV
            \\145 ORE => 6 MNCFX
            \\1 NVRVD => 8 CXFTF
            \\1 VJHF, 6 MNCFX => 4 RFSQX
            \\176 ORE => 6 VJHF
            ,
            .expected = "5586022",
        },
        .{
            .input =
            \\171 ORE => 8 CNZTR
            \\7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
            \\114 ORE => 4 BHXH
            \\14 VRPVC => 6 BMBT
            \\6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
            \\6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
            \\15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
            \\13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
            \\5 BMBT => 4 WPTQ
            \\189 ORE => 9 KTJDG
            \\1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
            \\12 VRPVC, 27 CNZTR => 2 XDBXC
            \\15 KTJDG, 12 BHXH => 5 XCVML
            \\3 BHXH, 2 VRPVC => 7 MZWV
            \\121 ORE => 7 VRPVC
            \\7 XCVML => 6 RJRHP
            \\5 BHXH, 4 VRPVC => 5 LTCX
            ,
            .expected = "460664",
        },
    };
    const allocator = std.testing.allocator;
    for (cases) |case| {
        const actual = try part2solver(allocator, case.input);
        defer allocator.free(actual);
        try std.testing.expectEqualStrings(case.expected, actual);
    }
}
