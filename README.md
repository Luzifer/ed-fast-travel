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

```bash
./ed-fast-travel 'Veil West Sector DL-Y d68' 'NGC 6820 Sector GW-W d1-0' 500
2016/09/11 16:40:11 Loading database...
2016/09/11 16:40:35 Searching your start / destination system...
2016/09/11 16:40:43 Found start system 'Veil West Sector DL-Y d68' at coordinates (-1398.25000 / -193.56250 / 418.90625)
2016/09/11 16:40:43 Found destination system 'NGC 6820 Sector GW-W d1-0' at coordinates (-5572.84375 / -11.65625 / 3342.68750)
2016/09/11 16:40:43 Linear distance between that systems is 5099.88 Ly
   1: 'Prooe Drye ME-K c11-15' (-1807.00000 / -183.87500 / 709.50000) with 501.61 Ly distance (total: 501.61 Ly)
   2: 'Prooe Drye EH-U e3-16' (-2216.15625 / -169.78125 / 987.34375) with 494.78 Ly distance (total: 996.39 Ly)
   3: 'Prae Drye XY-Y b47-4' (-2626.00000 / -145.87500 / 1252.78125) with 488.88 Ly distance (total: 1485.27 Ly)
   4: 'Aucofs GW-C d80' (-3042.09375 / -99.37500 / 1545.71875) with 510.99 Ly distance (total: 1996.25 Ly)
   5: 'Aucofs YK-H b16-0' (-3423.96875 / -85.90625 / 1846.21875) with 486.12 Ly distance (total: 2482.37 Ly)
   6: 'Aucofs FM-D c15-6' (-3831.46875 / -97.81250 / 2136.96875) with 500.73 Ly distance (total: 2983.11 Ly)
   7: 'Aucoths ZL-J d10-86' (-4258.78125 / -77.90625 / 2425.68750) with 516.09 Ly distance (total: 3499.20 Ly)
   8: 'NGC 6830 Sector GM-V d2-34' (-4662.75000 / -43.18750 / 2705.43750) with 492.60 Ly distance (total: 3991.80 Ly)
   9: 'Drojau BC-L a21-0' (-5086.93750 / -41.00000 / 2997.28125) with 514.89 Ly distance (total: 4506.69 Ly)
  10: 'NGC 6820 Sector DV-Y c1' (-5502.71875 / -16.90625 / 3289.50000) with 508.77 Ly distance (total: 5015.46 Ly)
  11: 'NGC 6820 Sector GW-W d1-0' (-5572.84375 / -11.65625 / 3342.68750) with 88.17 Ly distance (total: 5103.63 Ly)
2016/09/11 16:40:50 Calculation shows an overhead of 3.74 Ly in comparison to linear distance.
```

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
