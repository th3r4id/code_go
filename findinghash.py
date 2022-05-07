import mmh3
import requests
import codecs
 
response = requests.get('https://kits-reverseproxy.kuhlmann-its.de/images/logo-text-vertical-grey.png')
favicon = codecs.encode(response.content,"base64")
hash = mmh3.hash(favicon)
print(hash)