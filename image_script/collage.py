import cv2
import numpy as np
import os
import sys
from google.cloud import datastore
import gcloud
import bing

print('Successfully started, lad')

KERNEL_SIZE = (20, 20)
NUM2DOWNLOAD = 150
IMAGE_REPO_PATH = 'image_script/images/image_repo/'

def pickBestPattern(patt_imgs, x, y, image):
    try:
        pic_errors = []
        for pic in patt_imgs:
            error_array = abs(pic - image[y:y+KERNEL_SIZE[1], x:x+KERNEL_SIZE[0]])
            error_array = error_array.flatten()
            error = error_array.sum()
            pic_errors.append(error)
        min_index = pic_errors.index(min(pic_errors))
        pic = patt_imgs[min_index]
        image[y:y+KERNEL_SIZE[1], x:x+KERNEL_SIZE[0]] = pic
    except:
        pass

def createCollage(largeImagePath, keyword):
    # Read in image to modify
    large_image = cv2.imread(largeImagePath)
    large_image = np.asarray(large_image, dtype=np.float)

    # # Add more images to DB if necessary
    # numDL = len(gcloud.fetch_images(ds_client, keyword))
    # numToDL = NUM2DOWNLOAD - numDL
    # if numToDL > 0:
    #     gcloud.bing_to_gcloud(ds_client, keyword, numToDL)
    #
    # # Download images from DB
    # image_list = gcloud.fetch_images(ds_client, keyword)
    # gcloud.download_images_async(image_list, 'image_script/urllist.txt')

    bing.bing_img_search(keyword, NUM2DOWNLOAD, out_file='image_script/urllist.txt')

    # Read in images in repo
    patt_imgs = []
    for patt_img in os.listdir(IMAGE_REPO_PATH):
        filepath = IMAGE_REPO_PATH + '/' + patt_img
        img = cv2.imread(filepath)
        img = cv2.resize(img, KERNEL_SIZE)
        img = np.asarray(img, dtype=np.float)
        patt_imgs.append(img)

    # Start modifying image
    for y in range(0, large_image.shape[0] - KERNEL_SIZE[1]+1, KERNEL_SIZE[1]):
        for x in range(0, large_image.shape[1] - KERNEL_SIZE[0]+1, KERNEL_SIZE[0]):
            pickBestPattern(patt_imgs, x, y, large_image)

    # Write out modified image
    cv2.imwrite('image_script/images/output.png', large_image)

input_file = sys.argv[1]
keyword = sys.argv[2]

try:
    ds_client = datastore.Client(gcloud.proj_id)
except:
    print('We messed up with ds_client.')

createCollage(input_file, keyword)
