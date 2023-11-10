package sprites

import _ "embed"

var (
	//go:embed ground.png
	BgGround_png []byte

	//go:embed clouds.png
	BgClouds_png []byte

	//go:embed bushes.png
	BgBushes_png []byte

	//go:embed building.png
	BgBuilding_png []byte
)
