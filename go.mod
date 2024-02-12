module github.com/captncraig/streamdnddeck

go 1.21.4

require (
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/magicmonkey/go-streamdeck v0.0.6-0.20230123180902-478057861949
	golang.org/x/image v0.15.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/disintegration/gift v1.2.1 // indirect
	github.com/karalabe/hid v1.0.1-0.20190806082151-9c14560f9ee8 // indirect
)

replace github.com/magicmonkey/go-streamdeck => github.com/captncraig/go-streamdeck v0.0.0-20240212161916-e7cbc98bdc4c
