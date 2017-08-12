using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;

using Dapper;
using System.Data.Common;
using System.Data.SqlClient;

namespace dotnet
{
    public class Startup
    {
        public Startup(IHostingEnvironment env)
        {
            var builder = new ConfigurationBuilder()
                .SetBasePath(env.ContentRootPath)
                .AddJsonFile("appsettings.json", optional: false, reloadOnChange: true)
                .AddJsonFile($"appsettings.{env.EnvironmentName}.json", optional: true)
                .AddEnvironmentVariables();
            
            Configuration = builder.Build();
            connectionString = Configuration.GetSection("ConnectionStrings").GetValue<string>("DefaultConnection");
            // SetUpDatabase();
        }

        private void SetUpDatabase() {
            string sql = @"USE Words
            GO

            IF NOT EXISTS ( SELECT * FROM Words )
            BEGIN
            CREATE TABLE words.[Words] (
                id bigint IDENTITY(1,1) PRIMARY KEY,
                value varchar(255) NOT NULL UNIQUE,
                level int NOT NULL,
                created_on DateTime2 NOT NULL DEFAULT GETDATE() -- CONSTRAINT DF_Words_created_on DEFAULT ( SYSDATETIME() ),

                -- CONSTRAINT UC_Words UNIQUE (id, value),
                INDEX IX_Words NONCLUSTERED (level)
            )
            END
            ELSE
            BEGIN
                SELECT * FROM Words
            END
            ";

            using (var connection = ConnectionFactory())
            {
                connection.Open();

                var result = connection.Query(sql);

                Console.WriteLine(result);
            }
        }

        public static IConfigurationRoot Configuration { get; set; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            // Add framework services.
            // services.Configure<AppSettings>(Configuration.GetConfigurationSection("ConnectionStrings"));
            services.AddMvc();
        }

        private static string connectionString { get; set; }

        public static Func<DbConnection> ConnectionFactory = () => new SqlConnection(connectionString);

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IHostingEnvironment env, ILoggerFactory loggerFactory)
        {
            loggerFactory.AddConsole(Configuration.GetSection("Logging"));
            loggerFactory.AddDebug();

            app.UseMvc();
        }
    }
}
