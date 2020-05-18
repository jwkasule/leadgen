#!/bin/sh

echo "Please enter yourgoogle maps API key and hit ENTER:"
read API_KEY_INPUT
echo "#!/bin/sh\n\nexport API_KEY=$API_KEY_INPUT" > env/api.sh
chmod +x env/api.sh
echo "Success!"
