using System;
using System.Drawing;
using System.IO;
using System.Linq;
using UsefulExtensions;

namespace Test_Console2
{
    class Program
    {


        class A
        {
            public int B { get; set; }
        }
        static void Main(string[] args)
        {
            var a = new A { B = 5 };
            var d = a.ToJsonDocument(a.GetType());

            var nums = new[] { 5, 2, 3, 1 };
            var result = QuickSortArray(nums);
            result.ToList().ForEach(Console.WriteLine);
        }

        public static int[] QuickSortArray(int[] nums)
        {
            QuickSort(nums, 0, nums.Length - 1);
            return nums;
        }

        public static void QuickSort(int[] nums, int start, int end)
        {
            if (start == end)
                return;

            var pivotIndex = end;

            for (var i = start; i < end; i++)
            {
                if (nums[i] > nums[pivotIndex])
                {
                    var temp = nums[pivotIndex];
                    nums[pivotIndex] = nums[i];
                    nums[i] = temp;
                    pivotIndex = i;
                }
            }

            QuickSort(nums, start, pivotIndex - 1);
            QuickSort(nums, pivotIndex + 1, end);
        }
    }
}
