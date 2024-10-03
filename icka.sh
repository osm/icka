#!/usr/bin/env bash

#--------------#
# Instructions #
#--------------#
# Edit the script replacing any occurrence of `$user` with your shell username,
# your email, password and correct paths to the binary
##
# Make this file executable with:
#   chmod +x /home/$user/icka.sh
##
# Add the following to your crontab, to make it run every hour, so you never disconnect:
#   0 * * * * /home/$user/icka.sh >/dev/null 2>&1
##
# Enjoy!
#---------------------#
# End of instructions #
#---------------------#


##---------------##
## Configuration ##
##---------------##

#------------------------#
# IRCCloud account email #
#------------------------#
ickaEmail="email@example.com"

#---------------------------#
# IRCCloud account password #
#---------------------------#
ickaPassword="foobar"

#-------------------------#
# Path to the icka binary #
#-------------------------#
ickaBinary="/home/$user/go/bin/icka"

#-----------------------------------------------------------------------------------#
# Name of the file where we are going to write the successfailure message?          #
# Only one success/failure message will be kept every time this script is executed. #
#-----------------------------------------------------------------------------------#
ickaLog="icka.log"

##----------------------##
## End of configuration ##
##----------------------##

#--------------------------------------------------------------------------------------------------------------#
# DON'T EDIT ANYTHING BELOW UNLESS YOU KNOW EXACTLY WHAT YOU'RE DOING!                                         #
# If you touch the code below and then complain the script "suddenly stopped working" I'll touch you at night. #
#--------------------------------------------------------------------------------------------------------------#
"$ickaBinary" -email "$ickaEmail" -password "$ickaPassword"

if [ $? -eq 0 ]; then
	echo "$(date) - IRCCloud session kept alive." > "$HOME/$ickaLog"
else
	echo "$(date) - IRCCloud session couldn't be kept alive!" > "$HOME/$ickaLog"
fi
