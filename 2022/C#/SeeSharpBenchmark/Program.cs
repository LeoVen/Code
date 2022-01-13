using System.Text;
using BenchmarkDotNet.Attributes;
using BenchmarkDotNet.Running;

namespace SeeSharpBenchmark
{
    class Program
    {
        static void Main(string[] args)
        {
            var results = BenchmarkRunner.Run<Bench>();
        }
    }

    [MemoryDiagnoser]
    public class Bench
    {
        [Benchmark(Baseline = true)]
        public string BasicConcatenation()
        {
            string output = "";

            for (int i = 0; i < 100; i++)
            {
                output += i;
            }

            return output;
        }

        [Benchmark]
        public string UsingStringBuilder()
        {
            StringBuilder builder = new();

            for (int i = 0; i < 100; i++)
            {
                builder.Append(i);
            }

            return builder.ToString();
        }

        [Benchmark]
        public string StringInterpolation()
        {
            string result = "";

            for (int i = 0; i < 100; i++)
            {
                result = $"{result}{i}";
            }

            return result;
        }
    }
}
