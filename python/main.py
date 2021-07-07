import os
import sys

BUFFER_SIZE = 512

def reverse_file(argv):
    if len(argv) != 3:
        print(f"Usage: {argv[0]} [in_file] [out_file]")
        sys.exit(-1)

    in_file_path = argv[1]
    out_file_path = argv[2]

    in_file = open(in_file_path, 'rb')
    if in_file == 0:
        print("Can't open file", in_file_path)
        sys.exit(-2)

    out_file = open(out_file_path, 'wb')
    if in_file == 0:
        print("Can't create file", in_file_path)
        sys.exit(-2)

    in_file.seek(0, os.SEEK_END)
    remaining = in_file.tell()
    while remaining > 0:
        to_read = min(BUFFER_SIZE, remaining)
        in_file.seek(-to_read, os.SEEK_CUR)
        in_buf = in_file.read(to_read)
        read = len(in_buf)
        out_buf = bytearray(read)
        for i in range(read):
            out_buf[i] = in_buf[read - i - 1]
        out_file.write(out_buf)
        in_file.seek(-read, os.SEEK_CUR)
        remaining -= read

    in_file.close()
    out_file.close()


if __name__ == '__main__':
    reverse_file(sys.argv)
