# Dockerfile For building Golang Apps with Oracle InstantClient

Multistage Dockerfile for building Oracle Instantclient Apps

Download the -basic, -basiclite and -devel RPMs from http://www.oracle.com/technetwork/topics/linuxx86-64soft-092277.html

Make sure oci8.pc matches the InstantClient Version
Optionally setup tnsnames.ora and sqlnet.ora. You can still use ezconnect syntax without those files

Usage:
```
docker build .
```
