#!/usr/bin/env python

'''
chat_client.py -- Simple chat client for chat_server.py
'''
import logging
import select
import socket
import sys

LOG_FORMAT = '%(asctime)s %(name)s %(levelname)s %(message)s'
TIMED_ACTIVITY_FORMAT = '%(asctime)s %(message)s'
NO_FORMAT = '%(message)s'

ACTIVITY_LOG = 'activity.log'
TIMED_ACTIVITY_LOG = 'timed_activity.log'

def prompt(prompt_message):
    '''
    Prompt user for input
    '''
    sys.stdout.write('{}: '.format(prompt_message))
    sys.stdout.flush()

def chat_client(host, port):
    '''
    Chat client implementation

    Arguments:

    host -- hostname or IP address to connect to (IPv4 only)
    port -- port to connect to. chat_server.py must be running there.
    '''

    # Setup specialized loggers.
    a_log = setup_logger('a_log', ACTIVITY_LOG, log_level=logging.INFO,
                         log_format=NO_FORMAT)
    t_log = setup_logger('t_log', TIMED_ACTIVITY_LOG, log_level=logging.INFO,
                         log_format=TIMED_ACTIVITY_FORMAT)
    c_log = setup_logger('log', log_console=True, log_level=logging.DEBUG,
                         log_format=LOG_FORMAT)

    c_log.debug('Client started')

    prompt_message = "{}".format((host, port))

    # TODO: Extend support to include IPv6 protocol
    client_sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client_sock.settimeout(2)

    try:
        client_sock.connect((host, port))
    except socket.error as err:
        error_msg = 'Failed to connect to {}:{}. {}'.format(host, port, err)
        c_log.error(error_msg)
        return 1

    print("Welcome to the chat server {}:{}. Type '/quit' to exit"
          "".format(host, port))
    prompt(prompt_message)

    done = False
    while not done:
        input_list = [sys.stdin, client_sock]
        # Block while reading from the input list
        read_sockets, _, _ = select.select(input_list, [], [])
        for sock in read_sockets:
            if sock == client_sock:
                # Incomming message from server
                data = sock.recv(4096)
                if not data:
                    print '\nDisconnected from chat server'
                    done = True
                    break
                else:
                    # Display the message
                    sys.stdout.write("\n{}".format(data))
                    sys.stdout.flush()
                    a_log.info(data)
                    t_log.info(data)
            else:
                # Read from console and send message.
                msg = sys.stdin.readline().rstrip()
                a_log.info('%s: %s', prompt_message, msg)
                t_log.info('%s: %s', prompt_message, msg)
                if msg == '/quit' or msg == '.':
                    # User wants to quit
                    done = True
                    break
                if msg:
                    client_sock.send(msg)
                prompt(prompt_message)

    # We are done communicating. Now display the conversation log.
    print "{} Session:".format(prompt_message)
    print "----------------------------"
    with open(ACTIVITY_LOG) as log_handle:
        for line in log_handle:
            print line.rstrip()


def setup_logger(logger_name, log_file=None, log_console=False,
                 log_level=logging.INFO, log_format=NO_FORMAT):
    '''
    Creates a distinct logger

    Arguments:

    logger_name -- Name of logger to be created.
    log_file    -- Output filename for logger
    level       -- Log level threshold
    '''
    xlog = logging.getLogger(logger_name)
    xlog.setLevel(log_level)

    formatter = logging.Formatter(log_format)

    if log_file:
        file_handler = logging.FileHandler(log_file, mode='w')
        file_handler.setFormatter(formatter)
        xlog.addHandler(file_handler)

    if log_console:
        stream_handler = logging.StreamHandler()
        stream_handler.setFormatter(formatter)
        xlog.addHandler(stream_handler)

    return xlog

#
#  M A I N
#
if __name__ == "__main__":

    if len(sys.argv) < 3:
        print 'Usage: chat_client.py hostname port'
        sys.exit(1)

    # TODO: Implement complete CLI support using argparse()
    chat_server = sys.argv[1]
    chat_port = int(sys.argv[2])

    sys.exit(chat_client(chat_server, chat_port))
