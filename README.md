# Luzifer / ed-fast-travel

`ed-fast-travel` is a small helper utility to plan long-distance travels inside [Elite: Dangerous](https://www.elitedangerous.com).

## Facts

- It uses the nightly database dump from [EDSM](https://www.edsm.net/)
- The database is cached locally in `~/.local/share/ed-fast-travel`
- It works better when calculating stops in 100 Ly or more distance on the route
- You can plot courses with more than 25000 Ly distance in ~3m

## Difficulties

- Maybe the system you want to travel to is not yet known in EDSM  
> As of today 0.001313% of the galaxy has been discovered on EDSM, it will take 132,444 years, 4 months and 19 days to discover it entirely.
- Calculations with ~25 Ly stop distance might work but maybe they contain jumps >30 Ly distance as of above reason
- Calculations with <20 Ly distance might even run into an endless loop as of above reason
- Inside is no real route-planning engine but only a simple vector calculation to find stops on a linear axis

## Usage

You can download pre-compiled binaries on [GoBuilder](https://gobuilder.me/github.com/Luzifer/ed-fast-travel) for your system. Afterwards you need to execute it using the console / cmd prompt:

[![asciicast](https://asciinema.org/a/7ea5fd8hexx9wy38bcge3er1j.png)](https://asciinema.org/a/7ea5fd8hexx9wy38bcge3er1j?t=12)

As you can see you only need to know from where you're starting and where you're going. The number (500 in above case) is the distance the stops should be calcualated. I used 500 Ly as that's a good distance for the ingame route engine to calculate a more detailed route.

To update the local database with a fresh nightly dump from EDSM just add the flag `--update` when executing the utility:

```bash
# ./ed-fast-travel --update
2016/09/11 17:21:56 No local EDSM dump found or update forced, fetching dump...
Usage: ed-fast-travel <start system> <target system> <distance between nav points>

Example: ed-fast-travel 'Sol' 'Dryooe Prou GL-Y d369' 500
  This will calculate stops on your route from Sol to Dryooe Prou GL-Y d369 every 500Ly
```

The usage explanation appears because no parameters for routing are passed in that execution but the database now is refreshed. Please be nice to the EDSM servers and refresh only if you are sure you need to get new data and that there is new data. Refreshing multiple times a day does not give you any advantages as the database is only updated once a day.

## Build from source

- To build this utility from source you need to have a working go 1.6 or 1.7 environment.
- Execute `go get github.com/Luzifer/ed-fast-travel` and you will get the `ed-fast-travel` binary inside your `$GOPATH/bin/` directory
