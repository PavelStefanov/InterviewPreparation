using System;
using System.Collections.Generic;
using System.Linq;

namespace InvalidTransactions
{
    class Program
    {
        static void Main(string[] args)
        {

            var res = new Solution().InvalidTransactions(new[] { "alex,676,260,bangkok", "bob,656,1366,bangkok", "alex,393,616,bangkok", "bob,820,990,amsterdam", "alex,596,1390,amsterdam" });
            Console.WriteLine(res);
        }
    }

    public class Solution
    {

        public class Tr
        {
            public Tr(string name, int min, int money, string city)
            {
                Name = name;
                Min = min;
                Money = money;
                City = city;
            }

            public string Name { get; set; }
            public int Min { get; set; }
            public int Money { get; set; }
            public string City { get; set; }

            public bool IsAdded { get; set; }

            public override string ToString() => Name + "," + Min + "," + Money + "," + City;

            public static Tr Parse(string transaction)
            {
                var p = transaction.Split(",");
                return new Tr(p[0], int.Parse(p[1]), int.Parse(p[2]), p[3]);
            }
        }

        public IList<string> InvalidTransactions(string[] transactions)
        {
            var result = new List<string>();
            var dict = new Dictionary<string, List<Tr>>();

            foreach (var transaction in transactions)
            {
                var tr = Tr.Parse(transaction);

                if (!dict.ContainsKey(tr.Name))
                    dict[tr.Name] = new List<Tr>();

                dict[tr.Name].Add(tr);
            }

            foreach (var d in dict)
            {
                var trs = d.Value.OrderBy(t => t.Min).ToList();

                for (var i = 0; i < trs.Count; i++)
                {
                    if (i == trs.Count - 1)
                    {
                        if (trs[i].Money > 1000)
                            addIfNot(trs[i]);
                    }
                    else
                    {
                        if (trs[i].Money > 1000)
                            addIfNot(trs[i]);

                        if (trs[i + 1].Min - trs[i].Min <= 60)
                        {
                            var shouldAdd = false;
                            var j = i + 1;
                            for (; j < trs.Count; j++)
                            {
                                if (trs[j].Min - trs[i].Min <= 60)
                                {
                                    if (trs[i].City != trs[j].City)
                                        shouldAdd = true;
                                }
                                else
                                {
                                    break;
                                }
                            }

                            if (shouldAdd)
                            {
                                for (var k = i; k < j; k++)
                                    addIfNot(trs[k]);

                                i = k;
                            }
                        }
                    }
                }
            }

            return result;

            void addIfNot(Tr tr)
            {
                if (!tr.IsAdded)
                {
                    tr.IsAdded = true;
                    result.Add(tr.ToString());
                }
            }
        }
    }
}
