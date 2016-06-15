#!/usr/bin/env bash
# Try loading alfred-circleci cache. If the cache fails to load at first then
# retry 3 times with an exponetial backoff

function command { 
	/usr/local/opt/alfred-circleci/bin/alfred-circleci run loadcache
}

command && exit

n=1
until [ $n -ge 4 ]
do
	sleep $((10**$n))
	command && exit
	n=$[$n+1]
done
