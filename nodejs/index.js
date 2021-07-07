const fs = require('fs');

if (process.argv.length !== 4) {
    console.log('Usage:', process.argv[1], '[in_file] [out_file]');
    process.exit(-1);
}

const in_file_path = process.argv[2];
const out_file_path = process.argv[3];

const in_file = fs.openSync(in_file_path, 'r');
const out_file = fs.openSync(out_file_path, 'w');

const BUFFER_SIZE = 512;
const in_buf = Buffer.alloc(BUFFER_SIZE);
const out_buf = Buffer.alloc(BUFFER_SIZE);

const stats = fs.fstatSync(in_file);
let remaining = stats.size;

while (remaining > 0) {
    const pos = Math.max(remaining - BUFFER_SIZE, 0);
    const to_read = remaining < BUFFER_SIZE ? remaining : BUFFER_SIZE;
    const read = fs.readSync(in_file, in_buf, 0, to_read, pos);
    for (let i = 0; i < read; i++) {
        out_buf[i] = in_buf[read - i - 1];
    }
    fs.writeSync(out_file, out_buf, 0, read);
    remaining -= read;
}

fs.closeSync(in_file);
fs.closeSync(out_file);
