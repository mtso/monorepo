using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.DependencyInjection;

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
            IDemoService d = httpContext.RequestServices.GetService<IDemoService>();
            Console.WriteLine("FROM MIDDLEWARE GetService {0}", d.GetNum());
            
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
