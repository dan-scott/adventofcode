const std = @import("std");
const util = @import("../utils.zig");

pub const year = 2019;
pub const day = 8;

const COLS: usize = 25;
const ROWS: usize = 6;

const LAYER_LEN = COLS * ROWS;

pub fn part1solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    const layers = try splitLayers(allocator, input);
    defer allocator.free(layers);

    var min_0s = LAYER_LEN;
    var val: usize = 0;
    for (layers) |layer| {
        const zero_ct = std.mem.count(u8, layer, "0");
        if (zero_ct < min_0s) {
            const one_ct = std.mem.count(u8, layer, "1");
            const two_ct = std.mem.count(u8, layer, "2");
            val = one_ct * two_ct;
            min_0s = zero_ct;
        }
    }

    return std.fmt.allocPrint(allocator, "{}", .{val});
}

pub fn part2solver(allocator: std.mem.Allocator, input: []const u8) ![]const u8 {
    const layers = try splitLayers(allocator, input);
    defer allocator.free(layers);

    const depth = layers.len;

    var output = try allocator.alloc(u8, LAYER_LEN);
    defer allocator.free(output);
    for (0..ROWS) |row| {
        for (0..COLS) |col| {
            const idx = row * COLS + col;
            var layer: usize = 0;
            output[idx] = '2';
            while (output[idx] == '2' and layer < depth) {
                output[idx] = layers[layer][idx];
                layer += 1;
            }
        }
    }

    var printStr = try allocator.alloc(u8, LAYER_LEN + 1 + ROWS);
    var i: usize = 0;

    for (0..LAYER_LEN) |j| {
        if (j % COLS == 0) {
            printStr[i] = '\n';
            i += 1;
        }
        if (output[j] == '1') {
            printStr[i] = '0';
        } else {
            printStr[i] = ' ';
        }
        i += 1;
    }

    return @constCast(printStr);
}

fn splitLayers(allocator: std.mem.Allocator, input: []const u8) ![][]const u8 {
    const layer_ct = @divFloor(input.len, LAYER_LEN);
    var layers = try allocator.alloc([]const u8, layer_ct);
    var idx: usize = 0;
    var layer_iter = std.mem.window(u8, input, LAYER_LEN, LAYER_LEN);
    while (layer_iter.next()) |layer| {
        layers[idx] = layer;
        idx += 1;
    }
    return layers;
}
