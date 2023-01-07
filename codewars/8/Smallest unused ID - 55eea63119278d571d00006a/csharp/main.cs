using System;

public class Kata
{
	public static int NextId(int[] ids)
	{
		Array.Sort(ids);

		var result = 0;
		foreach (var id in ids) {
			if (result != id)
				return result;

			result++;
		}

		return result;
	}
}