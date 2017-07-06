#include <stdio.h>
#include <sys/socket.h>
#include <arpa/inet.h>

int main(int argc, char **argv)
{
    int socket_desc, client_sock, c, read_size;
    struct sockaddr_in server, client;
    char client_msg[2049];

    socket_desc = socket(AF_INET, SOCK_STREAM, 0);
    if (socket_desc == -1)
    {
        printf("Could not create socket");
    }
    printf("Socket created");

    return 0;
}
