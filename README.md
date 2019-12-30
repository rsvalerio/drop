# example

## Name

*drop* - can be used to drop specified DNS query types e.g. AAAA, returning 'NOTIMP' as rcode.

## Description

This plugins drop DNS requests of specific types

## Syntax

~~~ txt
drop AAAA
~~~

## Metrics

If monitoring is enabled (via the *prometheus* directive) the following metric is exported:

* `coredns_drop_request_count_total{server}` - query count to the *drop* plugin.

The `server` label indicated which server handled the request, see the *metrics* plugin for details.

## Ready

This plugin reports readiness to the ready plugin. It will be immediately ready.

## Examples

In this configuration, all queries for *.example.com will be dropped

~~~ corefile
example.com {
  drop AAAA
}
~~~

## Also See

See the [manual](https://coredns.io/manual).
