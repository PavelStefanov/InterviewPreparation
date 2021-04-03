using System;
using System.Collections;
using System.Collections.Generic;

namespace QueueOnStacks
{
    class Program
    {
        static void Main(string[] args)
        {
            var a =11/2;
            Console.WriteLine(a);
            var queue = new Queue();

            queue.Add(1);
            queue.Add(2);
            queue.Add(3);
            queue.Add(4);
            queue.Add(5);

            Console.WriteLine(queue.Remove());
            Console.WriteLine(queue.Remove());

            queue.Add(6);
            queue.Add(7);

            Console.WriteLine(queue.Remove());
            Console.WriteLine(queue.Remove());
            Console.WriteLine(queue.Remove());
        }
    }

    public class Queue
    {
        private bool lastS1 { get; set; } = true;
        private Stack<int> s1 { get; set; } = new Stack<int>();
        private Stack<int> s2 { get; set; } = new Stack<int>();

        public void Add(int i)
        {

            while (s1.TryPop(out var last))
            {
                s2.Push(last);
            }
            s1.Push(i);
            while (s2.TryPop(out var last))
            {
                s1.Push(last);
            }
        }

        public int Remove()
        {
            return s1.Pop();
        }
    }
}
