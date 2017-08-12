using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

using Dapper;
using Microsoft.Extensions.Configuration;
using System.Data.SqlClient;
using System.Data.Common;
// using dotnet.Startup;
// using Microsoft.AspNetCore.Mvc;

namespace dotnet.Controllers
{
    public class ListResult {
        public List<string> list;
    }

    [Route("api/[controller]")]
    public class ValuesController : Controller
    {
        // public static Func<DbConnection> ConnectionFactory = () => new SqlConnection(
        //     "User ID=sa;Password=p@ssw0rd;Server=localhost;Database=Words;Trusted_Connection=false;;Max Pool Size=1000;"
        // );

        // GET api/values
        [HttpGet]
        public JsonResult Get()
        {
            string sql = "select name from sysdatabases;";

            List<string> names = new List<string>();
            ListResult res = new ListResult();

            using (var connection = Startup.ConnectionFactory())
            {
                connection.Open();

                var rows = connection.Query(sql).ToList();
                foreach (var row in rows)
                {
                    names.Add(row.name);
                }

                res.list = names;
            }

            HttpContext.Response.Headers["Access-Control-Allow-Origin"] = "*";

            return Json(res);
        }

        // GET api/values/5
        [HttpGet("{id}")]
        public string Get(int id)
        {
            HttpContext.Response.Headers["Access-Control-Allow-Origin"] = "*";

            return "value" + id;
        }

        // POST api/values
        [HttpPost]
        public void Post([FromBody]string value)
        {
        }

        // PUT api/values/5
        [HttpPut("{id}")]
        public void Put(int id, [FromBody]string value)
        {
        }

        // DELETE api/values/5
        [HttpDelete("{id}")]
        public void Delete(int id)
        {
        }
    }
}
