using System;
using System.IO;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Net.Http;

class MyTcpListener {
  public static void Main() {
    TcpListener server = null;

    try {
      Int32 port = 3750;
      IPAddress localAddr = IPAddress.Parse("127.0.0.1");

      server = new TcpListener(localAddr, port);

      server.Start();

      Byte[] bytes = new Byte[1024];
      String data = null;

      while(true) {
        Console.Write("Waiting for a connection...");

        var task = server.AcceptTcpClientAsync();
        var client = task.Result;

        // var client = server.AcceptTcpClientAsync();
        // server.AcceptTcpClientAsync();
        // TcpClient client = server.AcceptTcpClient();
        Console.WriteLine("Connected!");

        data = null;

        NetworkStream stream = client.GetStream();

        int i;

        while ((i = stream.Read(bytes, 0, bytes.Length)) != 0) {
          data = System.Text.Encoding.ASCII.GetString(bytes, 0, i);
          Console.WriteLine("Received: {0}", data);

          data = data.ToUpper();
          byte[] msg = System.Text.Encoding.ASCII.GetBytes(data);

          var markup = "<html><body>hi</body></html>";
          byte[] raw = System.Text.Encoding.ASCII.GetBytes(markup);

          stream.Write(raw, 0, raw.Length);
          Console.WriteLine("Sent: {0}", data);
        }

        client.Dispose();
        // client.Close();
      }
    } catch(SocketException e) {
      Console.WriteLine("SocketException: {0}", e);
    } finally {
      server.Stop();
    }

    Console.WriteLine("\nHit enter to continue...");
    Console.Read();
  }
}

/*
namespace Simple {
  public class WebServer {
    private readonly HttpListener _listener = new HttpListener();
    private readonly Func<HttpListenerRequest, string> _handler;

    public WebServer(Func<HttpListenerRequest, string> handler, string[] prefixes) {
      if (!HttpListener.IsSupported) {
        throw new NotSupportedException("Needs Windows XP SP2, Server 2003 or later.");
      }

      // URI prefixes required
      // e.x.: http://localhost:3750/index/
      if (prefixes == null || prefixes.Length < 1) {
        throw new ArgumentException("prefixes");
      } 

      if (method == null) {
        throw new ArgumentException("handler required");
      }

      foreach (string s in prefixes) {
        _listener.Prefixes.Add(s);
      }

      _handler = handler;
      _listener.Start();
    }

    public void Run() {
      ThreadPool.QueueUserWorkItem((o) => {
        try {
          while (_listener.IsListening) {
            ThreadPool.QueueUserWorkItem((c) => {
              var ctx = c as HttpListenerContext;

              try {
                string responseString = _handler(ctx.Request);
                byte[] buf = Encoding.UTF8.GetBytes(responseString);
                ctx.Response.ContentLength64 = buf.Length;
                ctx.Response.OutputStream.Write(buf, 0, buf.Length);
              } catch {
                // suppress exceptions
              } finally {
                // always close stream
                ctx.Response.OutputStream.Close();
              }
            }, _listener.GetContext());
          }
        } catch {
          
        }
      });
    }

    public void Stop() {
      _listener.Stop();
      _listener.Close();
    }
  }
}
*/
