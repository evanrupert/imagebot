import requests
import json
import os

bing_creds = os.environ['BING_API_KEY']

def bing_img_search(term, count, offset=0, out_file='image_script/urllist.txt'):
    params = {'q': term, 'count': count, 'offset': offset}
    with open(bing_creds, 'r') as creds:
        headers = json.loads(creds.read())
        r = requests.get('https://api.cognitive.microsoft.com/bing/v7.0/images/search',
                            params=params, headers=headers)

    out = open(out_file, 'w')
    for value in r.json()['value']:
        out.write(value['thumbnailUrl'] + '\n')
    out.close()
    subprocess.call(["./imagedownloader", out_file])
