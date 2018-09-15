from google.cloud import datastore
import requests
import shutil
import json

proj_id = "docbottesting"
bing_creds = 'key'


def add_image(client, url, color, query):
    """Add image to client."""
    # Ask client for key for this entity
    key = client.key('ImageKind')
    # Create new entity
    image = datastore.Entity(key)
    # Add properties to entity
    image.update({
        'url': url,
        'query': query,
        'hex': color
    })
    # Add entity to datastore
    client.put(image)

    return image.key


def bing_to_gcloud(client, term, count, offset=0, outfolder='tmp/'):
    """Add results of Bing image query to datastore."""
    params = {'q': term, 'count': count, 'offset': offset}
    with open(bing_creds, 'r') as creds:
        headers = json.loads(creds.read())
        r = requests.get('https://api.cognitive.microsoft.com/bing/v7.0/images/search',
                            params=params, headers=headers)

    for value in r.json()['value']:
        add_image(client, value['contentUrl'], value['accentColor'], term)

def main():
    ds_client = datastore.Client(proj_id)
    bing_to_gcloud(ds_client, "cats", 10)


if __name__ == '__main__':
    main()
