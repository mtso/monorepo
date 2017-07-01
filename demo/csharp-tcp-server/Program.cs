// using Simple;
// using System;

// class Program {
//   static void Main(string[] args) {
//     WebServer ws = new WebServer(SendResponse, "http://localhost:3750/test/");
//     ws.Run();

//     // block thread
//     Console.WriteLine("A simple webserver. Press a key to quit.");
//     Console.ReadKey();
//     ws.Stop();
//   }

//   public static string SendResponse(HttpListenerRequest request) {
//     return string.Format("<html><body>hello~<br>{0}</body></html>", DateTime.Now);
//   }
// }
