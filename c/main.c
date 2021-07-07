#include <stdio.h>

#define BUFFER_SIZE 512
#define MIN(a,b) (a < b ? a : b)

int main(int argc, char* argv[]) {
    if (argc != 3) {
        printf("Usage: %s [in_file] [out_file]", argv[0]);
        return -1;
    }

    const char* in_file_path = argv[1];
    const char* out_file_path = argv[2];

    FILE* in_file = fopen(in_file_path, "rb");
    if (in_file == NULL) {
        printf("Can't open file %s", in_file_path);
        return -2;
    }
    FILE* out_file = fopen(out_file_path, "wb");
    if (out_file == NULL) {
        printf("Can't create file %s", out_file_path);
        return -2;
    }

    unsigned char in_buf[BUFFER_SIZE];
    unsigned char out_buf[BUFFER_SIZE];
    fseek(in_file, 0, SEEK_END);
    size_t remaining = ftell(in_file);
    while(remaining > 0) {
        size_t to_read = MIN(BUFFER_SIZE, remaining);
        fseek(in_file, -to_read, SEEK_CUR);
        size_t read = fread(in_buf, sizeof(unsigned char), to_read, in_file);
        for (int i = 0; i < read; i++) {
            out_buf[i] = in_buf[read - i - 1];
        }
        fwrite(out_buf, sizeof(unsigned char), read, out_file);
        fseek(in_file, -read, SEEK_CUR);
        remaining -= read;
    }

    fclose(in_file);
    fclose(out_file);
}
