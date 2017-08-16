using System;
using log4net;
using log4net.Core;
using log4net.Appender;

namespace logdemo
{
    public class DbAppender: AppenderSkeleton {
        protected override void Append(LoggingEvent loggingEvent)
        {
            Console.WriteLine(
                "!!!!!!!!!!!!!!!! CUSTOM APPENDER\n!!!!!!!!!!!!!!!!\n {0} {1}",
                loggingEvent.Level.Name,
                RenderLoggingEvent(loggingEvent)
            );
        }
    }
}
