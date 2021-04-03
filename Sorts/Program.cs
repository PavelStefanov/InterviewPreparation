using System;
using System.Linq;

namespace Sorts
{
    class Program
    {
        static void Main(string[] args)
        {
            var nums = new[] { 3, 5, 7, 2, 1, 6, 4 };

            var result = SortArray(nums);

            result.ToList().ForEach(Console.WriteLine);
        }

        public static int[] SortArray(int[] nums)
        {
            QuickSort(nums, 0, nums.Length - 1);
            return nums;
        }

        public static void QuickSort(int[] nums, int start, int end)
        {
            if (start >= end)
                return;

            var min = start;

            for (var i = start; i < end; i++)
            {
                var t = nums[i];
                if (nums[i] < nums[end])
                {
                    if (i != min)
                        Swap(nums, i, min);
                    min++;
                }
            }

            Swap(nums, min, end);

            QuickSort(nums, start, min - 1);
            QuickSort(nums, min + 1, end);
        }

        public static void Swap(int[] nums, int a, int b)
        {
            var temp = nums[a];
            nums[a] = nums[b];
            nums[b] = temp;
        }
    }
}