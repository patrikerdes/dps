# dps

## mpd2dps & dps2mpd

mpd2dps converts millions of documents per day (mpd) to documents per second (dps)

dps2mpd converts documents per second (dps) to millions of documents per day (mpd)

mpd = dps / 1000000 * (24 * 60 * 60)

dps = mpd * 1000000 / (24 * 60 * 60)

### How to install

1. Download your choice of [release of dps](https://github.com/patrikerdes/dps/releases) for your operating system (Windows, OS X or Linux)
2. Unzip

### How to run

	mpd2dps 42
	486

	dps2mpd 1000
	86

### How to build

1. Install Go
2. go build mpd2dps.go
3. go build dps2mpd.go

#### Cross-compiling
cross-compile.sh will cross-compile the tools to Windows, OS X and Linux. (Requires Go 1.5+.)

#### Creating release files
create-release-files.sh will create zip files with the tools built for Windows, OS X and Linux. The version number must be supplied as an argument when running the tool.
