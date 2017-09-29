#!/usr/bin/env python
"""
chat_server.py -- Simple Twisted Python based chat server for chat_client.py

Usage: 'twistd -y chatclient.py localhost 5000'.

Note: chatserver.py must be running.
"""
from __future__ import print_function

import optparse, sys
from twisted.conch.manhole import ColoredManhole
from twisted.internet import error, protocol, task
from twisted.internet.defer import Deferred
from twisted.protocols.basic import LineReceiver
from twisted.python import failure, log, usage

CONNECTION_DONE = failure.Failure(error.ConnectionDone())


class ChatClient(LineReceiver):
    """
        Implements the chat client protocol
    """
    def __init__(self):
        # Client name set by the server
        self.name = None

    def __prompt(self):
        '''
        Prompt user for input
        '''
        sys.stdout.write('{}: '.format(self.name))
        sys.stdout.flush()
        message = sys.stdin.readline().rstrip()
        return message

    def connectionMade(self):
        # Nothing to do here. Server will send the welcome message a trigger
        # prompting for input.
        log.msg("Server connection established")
        return

    def lineReceived(self, msg):
        log.msg(msg)
        if not line:
            # Ignore blank lines
            return
        self.transport.write(msg)
        msg = self.__prompt()
        # Exit commands are included in the broadcast.
        self.transport.write(msg)
        if msg == '/quit' or msg == '.':
            self.transport.loseConnection()

    def connectionLost(self, reason=CONNECTION_DONE):
        print("Connection lost")


class ChatClientFactory(protocol.ClientFactory):
    """
    Implements the protocol factory
    """
    def __init__(self):
        self.done = Deferred()
        self.protocol = ChatClient

    def clientConnectionFailed(self, connector, reason):
        print("Server connection failed")
        self.done.errback(reason)

    def clientConnectionLost(self, connector, reason):
        print("Server connection lost")
        self.done.callback(None)


class ConsoleIO(ColoredManhole):
    """
    A manhole protocol specifically for use with L{stdio.StandardIO}.
    """
    def connectionLost(self, reason):
        """
        When the connection is lost, there is nothing more to do.  Stop the
        reactor so that the process can exit.
        """
        reactor.stop()


def parse_args():
    usage = """Usage: %prog [hostname]:port
A simple Twisted-based chat client.
"""
    parser = optparse.OptionParser(usage)
    _, args = parser.parse_args()

    if len(args) < 1:
        print(parser.format_help())
        parser.exit()

    address = args[0]
    if ':' not in address:
        host = '127.0.0.1'
        port = address
    else:
        host, port = address.split(':', 1)

    if not port.isdigit():
        parser.error('Port value must be an integer.')

    return host, int(port)


def main(reactor):
    """
    Twistd application main
    """
    log.startLogging(sys.stdout)

    host, port = parse_args()

    factory = ChatClientFactory()
    reactor.connectTCP(host, port, factory)
    return factory.done

#
#  M A I N
#
if __name__ == "__main__":
    sys.exit(task.react(main))
