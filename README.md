# ImageBot

## What is it
ImageBot is a Discord bot that takes both an image and a keyword in order to create a collage out of images found by searching that keyword resembling the original image given to the bot

ImageBot also has a specific command to create any image out of minecraft block textures

But it doesn't stop there, ImageBot can also tell you the current weather as well as the predictions for up to the next four days.

## How we did it
The Discord interface was written in GO with a library called discordgo to create the message handlers and send messages through the Discord service.

The collage creration process was done in python with opencv and numpy.

The weather information process was also written in python and uses BeautifulSoup to collect data from html.

The images are found using the bing image search api to collect links and an asyncronous GO script to download a lot of images as quickly as possible.

The whole process is hosted on an instance of Compute on the Google Cloud Platform.