# Timewarp

Move images in time! (and by that I mean their EXIF date data)

## usage
```
timewarp -folder myfolder [-years 1] [-months 1] [-days 1]
```

## ...why?

I recently bought a new camera and took hundreds of pics with it, but they wouldn't show up in my gallery or google photos.
After days of wondering what is wrong, I realized I set my camera's date to 2021 instead of 2022 when setting it up and decided to make this to fix it!

## support

Initially this was written for my own use, so it only supported Macs with Xcode stuff installed, but now I updated it to work with the EXIF format instead, so it should hopefully work with any camera image .

You can still access the old branch [here](https://github.com/unickorn/timewarp/tree/master) if you for some reason want to change creation dates of files on a Mac with Xcode command line tools installed.