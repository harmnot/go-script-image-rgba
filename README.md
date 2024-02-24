# Before Start
- [x] Install go version 1.21.4
- [x] copy paste the image you want to use in the `images` folder

# How to use
Run `go run main.go {./images/{yourJPEGimages.jpeg}} {positiveOrNegativeNumber}` in the terminal

for warm (reddish) color use positive number 1 - 100 \
for cold (bluish) color use negative number -1 - -100

# Example
 Run `go run main.go ./images/yourJPEGimages.jpeg 1` in the terminal \
the output will be `./images/output.jpeg` in the `images` folder
