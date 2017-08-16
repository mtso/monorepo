using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

using Microsoft.Extensions.Logging;

namespace middleware_dotnet.Controllers
{
    [Route("api/[controller]")]
    public class ValuesController : Controller
    {
        private ILoggerFactory _Factory;
        private ILogger _Logger;

        public ValuesController(ILoggerFactory factory, ILogger<ValuesController> logger)
        {
            _Factory = factory;
            _Logger = logger;
        }

        // GET api/values
        [HttpGet]
        public IEnumerable<string> Get()
        {
            throw new Exception("arb exception");
            return new string[] { "value1", "value2" };
        }

        // GET api/values/5
        [HttpGet("{id}")]
        public string Get(int id)
        {
            var loggerFromDI = _Factory.CreateLogger("Values");
            // var loggerFactory = this.HttpContext.RequestServices.GetService<ValuesController>();
            // var loggerFromServices = loggerFactory.CreateLogger("Values");

            _Logger.LogDebug("From direct dependency injection");
            loggerFromDI.LogDebug("From dependency injection factory");
            // loggerFromServices.LogDebug("From services");

            return "value";
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
