package audio

import (
	"log"
	"os/exec"
)

func Mp32Pcm(mp3_file, pcm_file string) error {
	//mp3转pcm格式
	cmdStr := "ffmpeg -y  -i " + mp3_file + "  -acodec pcm_s16le -f s16le -ac 1 -ar 16000 " + pcm_file
	_, err := exec.Command("bash", "-c", cmdStr).Output()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
