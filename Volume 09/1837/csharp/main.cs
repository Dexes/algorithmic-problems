using System;
using System.Collections.Generic;
using System.Linq;

namespace Timus
{
    public static class Program
    {
        private const string Root = "Isenbaev";

        private static void Main()
        {
            foreach (var pair in BreadthFirstSearch(ReadData()))
                Console.WriteLine(pair.Value < 0 ? pair.Key + " undefined" : pair.Key + " " + pair.Value);
        }

        private static IDictionary<string, int> BreadthFirstSearch(IDictionary<string, string[]> graph)
        {
            var visited = new SortedDictionary<string, int>();
            var queue = new Queue<string>();

            if (graph.ContainsKey(Root))
            {
                visited[Root] = 0;
                queue.Enqueue(Root);
            }

            while (queue.Count > 0)
            {
                var current = queue.Dequeue();
                foreach (var key in graph[current])
                {
                    if (visited.ContainsKey(key)) continue;

                    visited[key] = visited[current] + 1;
                    queue.Enqueue(key);
                }

                graph.Remove(current);
            }

            foreach (var pair in graph)
                visited[pair.Key] = -1;

            return visited;
        }

        private static IDictionary<string, string[]> ReadData()
        {
            var size = int.Parse(Console.ReadLine());
            var graph = new Dictionary<string, Dictionary<string, bool>>();

            for (var i = 0; i < size; i++)
            {
                var data = Console.ReadLine().Split(' ');
                AddToGraph(graph, data[0], data[1], data[2]);
                AddToGraph(graph, data[1], data[0], data[2]);
                AddToGraph(graph, data[2], data[0], data[1]);
            }

            var result = new Dictionary<string, string[]>();
            foreach (var pair in graph)
                result[pair.Key] = pair.Value.Keys.ToArray();

            return result;
        }

        private static void AddToGraph(
            IDictionary<string, Dictionary<string, bool>> graph,
            string key, string first, string second)
        {
            if (!graph.ContainsKey(key))
                graph[key] = new Dictionary<string, bool>();

            graph[key][first] = true;
            graph[key][second] = true;
        }
    }
}