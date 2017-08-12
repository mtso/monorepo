using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

using Dapper;
using Microsoft.Extensions.Configuration;
using System.Data.SqlClient;
using System.Data.Common;
using dotnet.Models;

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
        // GET api/values
        [HttpGet("{id}")]
        public JsonResult Get(int id)
        {
            string sql = @"SELECT * FROM Words
            WHERE id=@WordID";

            ApiResult res = new ApiResult();

            using (var connection = Startup.ConnectionFactory())
            {
                connection.Open();
                try {
                    Word row = connection.QuerySingle<Word>(sql, new {WordID = id});

                    res.content = row;
                }
                catch(Exception err) {
                    res.message = "Invalid ID: " + id;
                }
            }

            res.is_success = res.message == null;
            HttpContext.Response.Headers["Access-Control-Allow-Origin"] = "*";
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
            string sql = @"select name from sysdatabases;";

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
