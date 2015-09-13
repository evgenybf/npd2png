import glob
import os
import os.path
import sys

from PIL import Image

WIDTH = 320
HEIGHT = 190 #240

def main(argv):
    fileMask = "*.NPD"
    for inputFile in glob.glob(fileMask):
        outputFile = os.path.splitext(inputFile)[0] + ".png"
        print ("Convert %s to %s" % (inputFile, outputFile))
        img = Image.new('RGB', (WIDTH, HEIGHT), "white")
        pixels = img.load()
        with file(inputFile, "rb") as f:
            f.read(14) # Skip header
            row = 0
            while row < HEIGHT:
                data = f.read(WIDTH * 3)
                if not data or len(data) < WIDTH:
                    break
                col = 0
                for i in xrange(0, len(data), 3):
                    pixels[col, row] = tuple(ord(x) for x in data[i:i + 3])
                    col += 1
                row += 1
        img.save(outputFile)
        os.remove(inputFile)

if __name__ == "__main__":
    main(sys.argv[:])
