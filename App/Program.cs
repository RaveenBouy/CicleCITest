using System;

namespace App
{
    class Program
    {
        static void Main(string[] args)
        {
            Calculator calc = new Calculator();

            Console.WriteLine("Enter 1 if you would like to perform addition");
            Console.WriteLine("Action: ");
            var input1 = Convert.ToInt32(Console.ReadLine());

            if(input1 == 1)
            {
                Console.WriteLine("Number 1:");
                var num1 = Convert.ToInt32(Console.ReadLine());
                Console.WriteLine("Number 2:");
                var num2 = Convert.ToInt32(Console.ReadLine());
                Console.WriteLine($"Result: {calc.Add(num1, num2)}");
            }
        }
    }
}
