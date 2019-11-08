package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/skybet/cali"

	"github.com/lmhd/lucli/lib"
)

func init() {

	command := cli.NewCommand("jess [params]")
	command.SetShort("Very basic generic Jessie Frazelle style GUI Container Runner")
	command.SetLong(`Runs containers as simply as possible.

This is basically the equivalent of:

docker run --rm -it \
	-e DISPLAY=unix$DISPLAY \
	--shm-size=2GB \
	jess/firefox "$@"

i.e. not particularly useful for much beyond simply testing an image.
Many images will require additional docker run parameters (and thus custom commands)
`)

	// Empty string for now. Image will be set during Init
	task := command.Task("")

	// More than the default SHM is often useful
	// --shm-size=2GB
	task.HostConf.ShmSize = 2147483648

	// Flag to specify which image to use
	command.Flags().StringP("image", "i", "docker.io/jess/firefox", "Image to use")
	command.BindFlags()

	task.SetInitFunc(func(t *cali.Task, args []string) {
		// Possibly move this call into lib.GetDisplay?
		err := lib.StartXQuartz()
		if err != nil {
			log.Fatalf("Unable to start X11 :(")
		}

		display, err := lib.GetDisplay()
		if err != nil {
			log.Fatalf("Unable to get DISPLAY :(")
		}

		t.AddEnv("DISPLAY", display)

		// Select the image to use
		image := cli.FlagValues().GetString("image")
		log.Infof("Using image: %s", image)
		t.SetImage(image)
	})

}
