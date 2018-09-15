from google.cloud import datastore
import requests
import shutil
import json
import os

proj_id = "docbottesting"
bing_creds = os.environ['BING_API_KEY']


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


def fetch_images(client, query_term):
    """Return list of images in datastore where entity.query == query_term.

    Pass output directly to download_images(), or filter first.
    """
    query = client.query(kind='ImageKind')
    query.add_filter('query', '=', query_term)
    return list(query.fetch())


def download_images(image_list, outfolder):
    """Download images contained in output of fetch_images()"""
    for image in image_list:
        i = requests.get(image['url'], stream=True)
        with open(outfolder + image['url'].split('/')[-1], 'wb') as out_file:
            shutil.copyfileobj(i.raw, out_file)
        del i


def bing_to_gcloud(client, term, count, offset=0, outfolder='tmp/'):
    """Add results of Bing image query to datastore."""
    params = {'q': term, 'count': count, 'offset': offset}
    with open(bing_creds, 'r') as creds:
        headers = json.loads(creds.read())
        r = requests.get(
            'https://api.cognitive.microsoft.com/bing/v7.0/images/search',
            params=params,
            headers=headers
            )

    for value in r.json()['value']:
        add_image(client, value['contentUrl'], value['accentColor'], term)


def main():
    ds_client = datastore.Client(proj_id)
    bing_to_gcloud(ds_client, "cats", 10)


# if __name__ == '__main__':
#     main()
