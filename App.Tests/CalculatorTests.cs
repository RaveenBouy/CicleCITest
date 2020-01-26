using NUnit.Framework;
using App;

namespace tests
{
    [TestFixture]
    public class CalculatorTests
    {
        [Test]
        public void Add_WhenCalled_ReturnTheTotalOfValue1AndValue2()
        {
            var calc = new Calculator(); 
            var result = calc.Add(12,23);
            Assert.That(result, Is.EqualTo(35));
        }
    }
}

