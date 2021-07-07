using System;
using System.IO;

namespace Reverso
{
    class Program
    {
        private const int BufferSize = 512;
        
        static void Main(string[] args)
        {
            if (args.Length != 2)
            {
                // There seems to be a bug in .net core where System.Environment.CommandLine is a dll even when running the produced .exe
                var exePath = System.Environment.CommandLine.Replace(".dll", ".exe");
                Console.WriteLine("Usage: {0} [in_file] [out_file]", exePath);
                System.Environment.ExitCode = -1;
                return;
            }

            var in_file_path = args[0];
            var out_file_path = args[1];

            var in_file = File.Open(in_file_path, FileMode.Open, FileAccess.Read);
            var out_file = File.Open(out_file_path, FileMode.OpenOrCreate, FileAccess.Write);

            var in_buf = new byte[BufferSize];
            var out_buf = new byte[BufferSize];
            var remaining = in_file.Length;
            in_file.Seek(0, SeekOrigin.End);
            while (remaining > 0)
            {
                var to_read = remaining < BufferSize ? (int)remaining : BufferSize;
                in_file.Seek(-to_read, SeekOrigin.Current);
                var read = in_file.Read(in_buf, 0, to_read);

                for (var i = 0; i < read; i++)
                {
                    out_buf[i] = in_buf[read - i - 1];
                }
                
                out_file.Write(out_buf, 0, read);
                in_file.Seek(-read, SeekOrigin.Current);
                remaining -= read;
            }
            
            in_file.Close();
            out_file.Close();
        }
    }
}