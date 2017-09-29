#!/usr/bin/env python

'''
chat_server.py -- Simple chat server for chat_client.py
'''

import sys
import socket
import select

from twisted.internet import protocol, reactor, endpoints


# TODO: Store all config values in a YAML config file.
HOST = '127.0.0.1'
PORT = 5000
MAX_CLIENTS = 3
RECV_BUFFER = 4096
# Zero value makes it wait forever.
SERVER_TIMEOUT = 0

# Track open connections
CONNECTION_LIST = []

def send_message(server_socket, originating_sock, message):
    '''
    Message all connected clients
    '''
    for socket_item in CONNECTION_LIST:
        # Message everyone except the server and originating client.
        if socket_item != server_socket and socket_item != originating_sock:
            try:
                socket_item.send(message)
            except socket.error:
                socket_item.close()
                if socket_item in CONNECTION_LIST:
                    CONNECTION_LIST.remove(socket)

def chat_server():
    '''
    Chat server implementation
    '''

    # TODO: Extend support to include IPv6 protocol
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    # Avoid 'socket.error: [Errno 98] Address already in use' when socket is
    # in TIME_WAIT state.
    # https://docs.python.org/2/library/socket.html#example
    server_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    server_socket.bind((HOST, PORT))
    server_socket.listen(MAX_CLIENTS)

    # Start by listening to icomming connections to this server
    CONNECTION_LIST.append(server_socket)

    print "Chat server started on {}:{} ".format(HOST, PORT)

    while True:
        # Wait for server socket events
        ready_to_read, _, _ = select.select(CONNECTION_LIST, [], [],
                                            SERVER_TIMEOUT)
        for sock in ready_to_read:
            if sock == server_socket:
                # Incomming connection request
                sockfd, addr = server_socket.accept()
                CONNECTION_LIST.append(sockfd)
                print "Connected with client {}".format(addr)

                send_message(server_socket, sockfd,
                             "{} entered our chatting room" .format(addr))
            else:
                # Message received from a client
                try:
                    data = sock.recv(RECV_BUFFER)
                    if data:
                        # Relay the message
                        msg = '{}: {}'.format(sock.getpeername(), data)
                        send_message(server_socket, sock, msg)
                    else:
                        # De-list the client
                        if sock in CONNECTION_LIST:
                            CONNECTION_LIST.remove(sock)

                        print "DEBUG: client {} is offline".format(addr)
                        # No data means the connection has been broken
                        send_message(server_socket, sock,
                                     "Client {} is offline".format(addr))
                except socket.error as err:
                    send_message(server_socket, sock,
                                 "Client {} is offline".format(addr))
                    print "ERROR: Failed to comunicate with client. {}" \
                          "".format(err)
                    continue

    server_socket.close()

#
#  M A I N
#
if __name__ == "__main__":
    sys.exit(chat_server())
