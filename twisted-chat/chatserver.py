#!/usr/bin/env python

"""
chat_server.py -- Simple Twisted Python based chat server for chat_client.py

Usage: 'twistd -y chatserver.py' then connect with multiple telnet clients to port 5000

"""

from twisted.application import internet, service
from twisted.internet.defer import Deferred
from twisted.internet.protocol import ServerFactory
from twisted.protocols.basic import LineReceiver
from twisted.python import log


class ChatServerProtocol(LineReceiver):
    '''
    Implements the chat server protocol
    See: http://twistedmatrix.com/documents/current/core/howto/servers.html
    '''
    def __init__(self, protocol_factory):
        self.name = None
        self.factory = protocol_factory

    def _broadcast(self, message):
        for name, client in self.factory.clients.items():
            if client.transport.getPeer() != self.transport.getPeer():
                log.msg('Sending to {}: "{}"'.format(name, message))
                client.sendLine(message)

    def connectionMade(self):
        # Increment connection count
        self.factory.numProtocolInstances += 1
        self.name = "c{}".format(self.factory.numProtocolInstances)
        log.msg("{} connected".format(self.name))
        # Welcome client with MOTD message
        self.sendLine(self.factory.service.motd)
        # Inform all other clients about it
        self._broadcast("{} joined".format(self.name))
        # Register this client
        self.factory.clients[self.name] = self

    def connectionLost(self, reason):
        log.msg("Client {} disconnected. "
                "{} {}".format(self.name, self.transport.getPeer(), reason))
        # Unregister client
        self.factory.clients.pop(self.name, None)
        self.factory.numProtocolInstances -= 1

    def lineReceived(self, line):
        log.msg("{} says: {}\n".format(self.name, line))
        self._broadcast(line)

class ChatServerFactory(ServerFactory):
    """
    Chat server protocol factory
    """
    def __init__(self, _service):
        self.clients = {}
        self.done = Deferred()
        self.numProtocolInstances = 0
        self.protocol = ChatServerProtocol
        self.service = _service

    def buildProtocol(self, address):
        return ChatServerProtocol(self)

    def clientConnectionFailed(self, connector, reason):
        print('Connection failed:', reason.getErrorMessage())
        self.done.errback(reason)

    def clientConnectionLost(self, connector, reason):
        print('Connection lost:', reason.getErrorMessage())
        self.done.callback(None)

class ChatServerService(service.Service):
    """
    Defines the service even handlers
    """
    def __init__(self, _motd):
        self.motd = _motd
        self.setName("chat_server")

    def startService(self):
        """
        Start the service.
        """
        log.msg("Starting service {}".format(self.name))
        service.Service.startService(self)
        # TODO: Perform random MOTD reads from a file.
        #self.motd = open(self.motd_file).read()
        #log.msg("MOTD loaded from: %s" % (self.motd_file,))

    def stopService(self):
        """
        Stop the service.

        @rtype: L{Deferred}
        @return: a L{Deferred} which is triggered when the service has
            finished shutting down. If shutting down is immediate, a
            value can be returned (usually, C{None}).
        """
        log.msg("Stopping service {}".format(self.name))
        # TODO: Inform all connected clients that the server service is
        # going down.
        return None

#
#  M A I N
#

# TODO: Store all config values in a TAC file.
INTERFACE = "localhost"
PORT = 5000
MOTD = "Welcome to the chat server. Today is a fine day!"

# Create a top service node for the purpose of grouping the chat server
# service and its dependencies. There is also the option of subscribing
# both services to the application.
top_service = service.MultiService()

# The chat server service gets its MOTD when it is started.
chat_server_service = ChatServerService(MOTD)
chat_server_service.setServiceParent(top_service)

# The TCP service connects the factory to a listening socket.
# It creates the listening socket when it is started.
factory = ChatServerFactory(chat_server_service)
tcp_service = internet.TCPServer(PORT, factory, interface=INTERFACE)
tcp_service.setServiceParent(top_service)

# Tell twistd what application to start.
application = service.Application("chat_server")

# Hook the service tree we made to the application.
top_service.setServiceParent(application)

# The application is now ready to go. When started by twistd, it will start
# the child services and run the chat server.
