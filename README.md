## mpd2dps & dps2mpd

mpd2dps converts millions of documents per day (mpd) to documents per second (dps)

dps2mpd converts documents per second (dps) to millions of documents per day (mpd)

mpd = dps / 1000000 * (24 * 60 * 60)
dps = mpd * 1000000 / (24 * 60 * 60)


### How to run

	mpd2dps 42
	486

	dps2mpd 1000
	86

### How to build

1. Install Go
2. go build mpd2dps.go
3. go build dps2mpd.go
