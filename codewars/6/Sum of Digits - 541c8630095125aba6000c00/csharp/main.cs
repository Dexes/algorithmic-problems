public class Number
{
    public int DigitalRoot(long n)
    {
        if (n < 10) return (int) n;

        var sum = 0;
        while (n > 0) {
            sum += (int) (n % 10);
            n /= 10;
        }

        return DigitalRoot(sum);
    }
}