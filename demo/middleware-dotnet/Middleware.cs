using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Http;

namespace middleware_dotnet
{
    public class Middleware
    {
        private readonly RequestDelegate _next;

        public Middleware(RequestDelegate next)
        {
            _next = next;
        }

        public async Task Invoke(HttpContext httpContext)
        {
            Console.WriteLine($"Request for {httpContext.Request.Path} received ({httpContext.Request.ContentLength ?? 0} bytes)");

            try {
                await _next.Invoke(httpContext);
            }
            catch (Exception e) {
                await httpContext.Response.WriteAsync(String.Format("CAUGHT EXCEPTION {0}", e));
            }
        }
    }
}
