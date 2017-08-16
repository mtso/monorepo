using System;
using Microsoft.Extensions.Logging;

namespace middleware_dotnet
{
    public class MyLoggerProvider: ILoggerProvider
    {
        public ILogger CreateLogger(string CategoryName)
        {
            return new MyLogger();
        }
        
        public void Dispose()
        { }

        private class MyLogger: ILogger
        {
            public bool IsEnabled(LogLevel l)
            {
                return true;
            }

            public void Log<TState>(LogLevel lv, EventId eid, TState s, Exception e, Func<TState, Exception, string> fmt)
            {
                Console.WriteLine("MY LOGGER!!!!!", fmt(s, e));
            }

            public IDisposable BeginScope<TState>(TState s)
            {
                return null;
            }
        }
    }
}
