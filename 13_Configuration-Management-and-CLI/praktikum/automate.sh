#!/bin/sh

NAME=$1
FB=$2
LINKEDIN=$3

# Parameter Input Barrier
if [ -z LINKEDIN ]; then
    echo "Arguments not enough. 3 Arguments Expected!"
    exit
fi

TREE_DIR[0]="$NAME $(date)"

TREE_DIR[1]="${TREE_DIR[0]}/about_me"
TREE_DIR[2]="${TREE_DIR[1]}/personal"
TREE_DIR[3]="${TREE_DIR[1]}/professional"

TREE_DIR[4]="${TREE_DIR[0]}/my_friends"

TREE_DIR[5]="${TREE_DIR[0]}/my_system_info"

for DIR in "${TREE_DIR[@]}"
do
    mkdir "$DIR"
done

echo "https://www.facebook.com/$FB" > "${TREE_DIR[2]}/facebook.txt"
echo "https://www.linkedin.com/in/$LINKEDIN" > "${TREE_DIR[3]}/linkedin.txt"

# about this laptop
echo "My username: nibras_alyassar" > "${TREE_DIR[5]}/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "${TREE_DIR[5]}/about_this_laptop.txt"

# get a list of friends
curl "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt" > "${TREE_DIR[4]}/list_of_my_friends.txt"

# check internet connection
echo "Connection to google:" > "${TREE_DIR[5]}/internet_connection.txt"
ping -n 3 forcesafesearch.google.com >> "${TREE_DIR[5]}/internet_connection.txt"






