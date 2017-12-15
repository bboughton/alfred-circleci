#!/usr/bin/env bash
# Try loading alfred-circleci cache. If the cache fails to load at first then
# retry 3 times with an exponetial backoff

log() {
	echo $(date +%Y%m%dt%H%M%S) "$@"
}

clearcache() {
	/usr/local/opt/alfred-circleci/bin/alfred-circleci run clearcache
}

loadcache() {
	/usr/local/opt/alfred-circleci/bin/alfred-circleci run loadcache
}

log "clearing cache"
clearcache

log "retrieving projects from circle ci"
n=0
while (( $n < 4 ))
do
	loadcache && exit
	n=$[$n+1]
	log "failed to retrieve projects, waiting $((10**n)) seconds"
	sleep $((10**$n))
done

exit 1
