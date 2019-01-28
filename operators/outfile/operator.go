package outfile

import "os"

const outputFile = "./dist/body.txt"

func File(b []byte) error {
	out, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = out.Write(b)
	return err
}
