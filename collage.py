import cv2
import numpy as np
import time
import os

def pickBestPattern(patt_imgs, x, y, image, kernelSize):
    pic_errors = []
    for pic in patt_imgs:
        error_array = abs(pic - image[y:y+kernelSize[1], x:x+kernelSize[0]])
        error_array = error_array.flatten()
        error = error_array.sum()
        pic_errors.append(error)
    min_index = pic_errors.index(min(pic_errors))
    pic = patt_imgs[min_index]
    image[y:y+kernelSize[1], x:x+kernelSize[0]] = pic

def createCollage(largeImagePath, imageRepoPath, kernelSize):
    large_image = cv2.imread(largeImagePath)
    large_image = cv2.resize(large_image, (800, 800))
    large_image = np.asarray(large_image, dtype=np.float)

    patt_imgs = []
    for patt_img in os.listdir(imageRepoPath):
        filepath = imageRepoPath + '/' + patt_img
        img = cv2.imread(filepath)
        img = cv2.resize(img, kernelSize)
        img = np.asarray(img, dtype=np.float)
        patt_imgs.append(img)

    start_time = time.time()
    print('Working...')
    for y in range(0, large_image.shape[1] - kernelSize[1] + 1, kernelSize[1]):
        for x in range(0, large_image.shape[0] - kernelSize[0] + 1, kernelSize[0]):
            pickBestPattern(patt_imgs, x, y, large_image, kernelSize)
    end_time = time.time()

    print(end_time - start_time)
    cv2.imshow('image', large_image / 255.0)
    cv2.waitKey(0)
    cv2.destroyAllWindows()