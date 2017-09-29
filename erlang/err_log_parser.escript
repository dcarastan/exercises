#!/usr/bin/env escript
%% -*- erlang -*-
%%! -smp enable -sname factorial -mnesia debug verbose

-define(BLOCK_MARKER, "| Block start").
-define(ERROR_MARKER, "| ERROR: FAILURE").

% expect only one argument
main([String]) ->
    {ok, F} = file:open(String, read),
    Line = io:get_line(F, ''),
    io:format("file ~p\n", [Line]),
    case Line(0, length(?BLOCK_MARKER)) ->
        ?BLOCK_MARKER ->
            acc += 1;
                

main(_) ->
    usage().

usage() ->
    io:format("usage: err_log_parser logfile\n"),
    halt(1).
