# Guide: https://fabiobarbero.eu/posts/signalbot/

# install java
sudo apt update
sudo apt upgrade
sudo apt install openjdk-17-jdk-headless

# see version in https://github.com/AsamK/signal-cli/releases/latest
export VERSION="0.10.4.2"
wget https://github.com/AsamK/signal-cli/releases/download/v"${VERSION}"/signal-cli-"${VERSION}"-Linux.tar.gz
sudo tar xf signal-cli-"${VERSION}"-Linux.tar.gz -C /opt
sudo ln -sf /opt/signal-cli-"${VERSION}"/bin/signal-cli /usr/local/bin/


export SIGNALNUMBER="YOUR NUMBER"
# Tutorial: https://github.com/AsamK/signal-cli/wiki/Quickstart
signal-cli -u $SIGNALNUMBER register

# If ask for captcha
# Guide: https://github.com/AsamK/signal-cli/wiki/Registration-with-captcha
signal-cli -u $SIGNALNUMBER register --captcha {CAPTCHA}

# Then you will receive SMS in your NUMBER
signal-cli -u $SIGNALNUMBER verify {CODE}

# Test 
# Send
signal-cli -u $SIGNALNUMBER send -m "Hello World" {RECIPIENT}
# Receive
signal-cli -u $SIGNALNUMBER receive


# Then, Set up signal-cli-rest-api
# https://github.com/bbernhard/signal-cli-rest-api

# Edit ./methods/signal/signal.go Endpoint to signal-cli-rest-api endpoint
