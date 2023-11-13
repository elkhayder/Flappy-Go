package sounds

import _ "embed"

var (
	//go:embed background.mp3
	Background_mp3 []byte

	//go:embed die.wav
	Die_wav []byte

	//go:embed hit.wav
	Hit_wav []byte

	//go:embed point.wav
	Point_wav []byte

	//go:embed swoosh.wav
	Swoosh_wav []byte

	//go:embed wing.wav
	Wing_wav []byte
)
