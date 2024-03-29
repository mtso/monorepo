﻿using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

using Microsoft.Extensions.Logging;
using Microsoft.Extensions.DependencyInjection;

namespace middleware_dotnet.Controllers
{
    [Route("api/[controller]")]
    public class ValuesController : Controller
    {
        private ILoggerFactory _Factory;
        private ILogger _Logger;
        private readonly IDemoService _demoService;

        public ValuesController(ILoggerFactory factory, ILogger<ValuesController> logger, IDemoService demoService)
        {
            _Factory = factory;
            _Logger = logger;
            _demoService = demoService;
        }

        // GET api/values
        [HttpGet]
        public IEnumerable<string> Get()
        {
            throw new Exception("arb exception");
            return new string[] { "value1", "value2" };
        }

        [HttpGet("{a}/{b}")]
        public string Get(int a, int b)
        {
            return "another test";
        }

        // GET api/values/5
        [HttpGet("{id}")]
        public string Get(int id)
        {
            var svc = HttpContext.RequestServices.GetService<IDemoService>();
            Console.WriteLine("From services {0}", svc.GetNum());
            Console.WriteLine("From direct dependency injection {0}", _demoService.GetNum());

            var loggerFromDI = _Factory.CreateLogger("Values");
            var loggerFactory = this.HttpContext.RequestServices.GetService<ILoggerFactory>();
            var loggerFromServices = loggerFactory.CreateLogger("Values");

            _Logger.LogDebug("From direct dependency injection");
            loggerFromDI.LogDebug("From dependency injection factory");
            loggerFromServices.LogDebug("From services");

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
