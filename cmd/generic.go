package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/skybet/cali"
)

func init() {

	command := cli.NewCommand("generic [params]")
	command.SetShort("Very basic generic Container Runner")
	command.SetLong(`Runs containers as simply as possible.

This is basically the equivalent of:

docker run --rm -it -w /tmp/workdir -v $PWD:/tmp/workdir <image> "$@"

i.e. not particularly useful for much beyond simply testing an image.
Many images will require additional docker run parameters (and thus custom commands)
`)

	// Empty string for now. Image will be set during Init
	task := command.Task("")

	// Flag to specify which image to use
	command.Flags().StringP("docker-image", "i", "centos:7", "Image to use")
	command.Flags().StringP("entrypoint", "e", "sh", "entrypoint to use")
	command.BindFlags()

	task.SetInitFunc(func(t *cali.Task, args []string) {
		// Select the image to use
		image := cli.FlagValues().GetString("docker-image")
		log.Infof("Using image: %s", image)
		t.SetImage(image)

		entrypoint := cli.FlagValues().GetString("entrypoint")
		log.Infof("Using entrypoint: %s", entrypoint)
		t.Conf.Entrypoint = []string{entrypoint}
	})

}
