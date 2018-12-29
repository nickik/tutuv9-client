#!/bin/bash

mkdir luxeriadoorbell
cd luxeriadoorbell
wget -q  https://github.com/nickik/tutuv9-client/blob/master/sound.mp3?raw=true 
mv -f sound.mp3?raw=true sound.mp3
mv -f sound.mp3 /usr/sbin/

wget -q https://github.com/nickik/tutuv9-client/blob/master/tutuv9-client?raw=true 
mv tutuv9-client?raw=true client
chmod +x client
mv -f ./client /usr/sbin/luxeriadoorbellclient

wget -q https://github.com/nickik/tutuv9-client/blob/master/tutu.service?raw=true
mv tutu.service?raw=true tutu.service
mv tutu.service /etc/systemd/system/
chmod 664 /etc/systemd/system/tutu.service

systemctl start tutu.service

cd ..
rm -rf luxeriadoorbell



