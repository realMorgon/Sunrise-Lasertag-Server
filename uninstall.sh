#!/bin/bash
# Uninstall script for Sunrise-Lasertag-Server

echo "/===== UNINSTALL SUNRISE LASERTAG SERVER =====/"
echo ""
echo "Willst du den Server wirklich deinstallieren?"
read -p "Alle lokalen Daten werden dadurch gelöscht (J/n) " -n 1 -r
echo ""
if [[ ! $REPLY =~ ^[Jj]$ ]]; then
    exit 1
fi

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)

export $(grep -v '^#' "$SCRIPT_DIR/.env" | xargs)

# TODO: Kill running processes
sudo systemctl stop hostapd dnsmasq
sudo systemctl disable hostapd dnsmasq

sudo apt-get remove --purge -y hostapd dnsmasq
sudo apt-get autoremove -y

# Remove PostgreSQL database and user
sudo -u postgres psql -c "DROP DATABASE $DB_NAME;"
sudo -u postgres psql -c "DROP USER $DB_USER;"

sudo systemctl stop postgresql
sudo systemctl disable postgresql

sudo apt-get remove --purge -y postgresql
sudo apt-get autoremove -y

# Remove .env file
rm -f "$SCRIPT_DIR/.env"
