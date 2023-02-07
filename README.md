# Crop image to square (1:1)

Drag and drop your png/jpg image to crop it to same width and height (square). The shorter axis is selected as the new dimentions. Only tested on windows.

# How to use
You can also use the application in the command line. Pass the image path as the first argument. You can use relative or absolute paths.


```
squarer.exe ./myImage.png
```

### Output format
You can also select the format for the new image (jpg or png). Pass the type as the second argument.

```
squarer.exe ./myImage.png jpg
```
Note: jpg to png is really slow and produces large output file.