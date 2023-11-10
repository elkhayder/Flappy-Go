package sprites

import _ "embed"

var (
	/* Pipe */

	//go:embed pipe/red.png
	PipeRed_png []byte

	//go:embed pipe/green.png
	PipeGreen_png []byte

	/* Bird */

	//go:embed bird/downflap.png
	BirdDownflap_png []byte

	//go:embed bird/midflap.png
	BirdMidflap_png []byte

	//go:embed bird/upflap.png
	BirdUpflap_png []byte

	/* Background */

	//go:embed background/ground.png
	BgGround_png []byte

	//go:embed  background/clouds.png
	BgClouds_png []byte

	//go:embed  background/bushes.png
	BgBushes_png []byte

	//go:embed  background/building.png
	BgBuilding_png []byte
)
