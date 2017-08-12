using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

using Dapper;
using Microsoft.Extensions.Configuration;
using System.Data.SqlClient;
using System.Data.Common;

namespace dotnet.Controllers
{
    public class ApiResult {
        public Boolean is_success;
        public string message;
        public object content;
    }

    public class TempWord {
        public int id;
        public string value;
    }

    [Route("api/[controller]")]
    public class WordController : Controller
    {
        // public static Func<DbConnection> ConnectionFactory = () => new SqlConnection(
        //     "User ID=sa;Password=p@ssw0rd;Server=localhost;Database=Words;Trusted_Connection=false;;Max Pool Size=1000;"
        // );

        // GET api/values
        [HttpGet]
        public JsonResult Get()
        {
            HttpContext.Response.Headers["Access-Control-Allow-Origin"] = "*";

            ApiResult res = new ApiResult();
            res.is_success = true;

            TempWord tmp = new TempWord();
            tmp.id = 123;
            tmp.value = "attention";
            res.content = tmp;
            // res.content = "{\"id\":123,\"value\":\"attention\"}";

            return Json(res);
        }
    }

    [Route("api/[controller]")]
    public class WordsController : Controller
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
    }
}