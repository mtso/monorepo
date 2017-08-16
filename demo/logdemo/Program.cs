using System;
using System.Xml;
using System.IO;
using log4net;
using log4net.Config;
using log4net.Repository;
using System.Reflection;

namespace logdemo
{
    class Program
    {
        private static readonly ILog log = LogManager.GetLogger(typeof(Program));

        static void Main(string[] args)
        {
            XmlDocument config = new XmlDocument();
            config.Load(File.OpenRead("./log4net.config"));

            ILoggerRepository repo = log4net.LogManager.CreateRepository(
                Assembly.GetEntryAssembly(),
                typeof(log4net.Repository.Hierarchy.Hierarchy)
            );

            XmlConfigurator.Configure(repo, config["log4net"]);

            log.Info("Application - Main is invoked");
            log.InfoFormat("Application - Main is invoked {0}", 101);
            log.Debug("Application - Main is invoked");
            log.Warn("Application - Main is invoked");
            log.Error("Application - Main is invoked");
            log.Fatal("Application - Main is invoked");
        }
    }
}
