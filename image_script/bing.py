import requests
import shutil
import json

creds_dir = 'key'

def bing_img_search(term, count, offset=0, outfolder='tmp/'):
    params = {'q': term, 'count': count, 'offset': offset}
    with open(creds_dir, 'r') as creds:
        headers = json.loads(creds.read())
        r = requests.get('https://api.cognitive.microsoft.com/bing/v7.0/images/search',
                            params=params, headers=headers)

    for value in r.json()['value']:
        i = requests.get(value['contentUrl'], stream=True)
        with open(outfolder + value['contentUrl'].split('/')[-1], 'wb') as out_file:
            shutil.copyfileobj(i.raw, out_file)
        del i


if __name__ == "__main__":
    bing_img_search("google logo", 20)
