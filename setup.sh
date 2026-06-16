#!/bin/bash
# Diese Datei muss beim erstmaligen Einrichten des Servers ausgeführt werden.
#
# Ablauf des Skripts
# - Einrichten des Hotspots
# - Einrichten der Datenbank Zugangsdaten
# - Datenbank Migration
#
# Das Skript wird nach Namen und Passwörtern fragen. Diese sind frei wählbar.
# Nach manchen Fragen stehen eckige Klammern, das sind Standartwerte die bei leerer Eingabe angenommen werden.

# Absoluten Pfad des Ordners, in dem dieses Skript liegt, ermitteln
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)

echo "/===== SUNRISE LASERTAG SERVER - SETUP =====/"
echo "von realMorgon und Kuscheltiermafia"
echo "Unterstützt durch den Selbsthilfeclub der Götter (selbsthilfe.club)"
echo "Beep boop beep beep boop"
echo "----- ----- -----"
echo ""


# Hotspot
if systemctl is-active --quiet hostapd; then
    echo "hostapd ist bereits aktiv. Überspringe die Einrichtung."
    echo "Sollte das ein Fehler sein, deaktiviere hostapd und starte das Setup erneut."
    echo ""
else
    echo "===== HOTSPOT ====="
    read -p "SSID [LaserTag]: " SSID
    SSID=${SSID:-LaserTag}

    read -p "Passwort: " PASSWORD

    read -p "Server-IP [192.168.4.1]: " SERVER_IP
    SERVER_IP=${SERVER_IP:-192.168.4.1}

    read -p "IP-Range, untere Grenze [192.168.4.10]: " IP_RANGE_LOW
    IP_RANGE_LOW=${IP_RANGE_LOW:-192.168.4.10}

    read -p "IP-Range, obere Grenze [192.168.4.254]: " IP_RANGE_HIGH
    IP_RANGE_HIGH=${IP_RANGE_HIGH:-192.168.4.254}

    echo "Netzwerk wird eingerichtet..."

    # 1. Paketquellen aktualisieren und hostapd + dnsmasq (für IPs) installieren
    sudo apt update
    sudo apt install -y hostapd dnsmasq

    # 2. NetworkManager anweisen, wlan0 zu ignorieren (LAN bleibt aktiv!)
    sudo tee /etc/NetworkManager/conf.d/99-disable-wlan0.conf > /dev/null << 'EOF'
[keyfile]
unmanaged-devices=interface-name:wlan0
EOF
    sudo systemctl reload NetworkManager

    # 3. DNS-Konflikt auf Port 53 lösen (systemd-resolved deaktivieren)
    sudo systemctl stop systemd-resolved 2>/dev/null
    sudo systemctl disable systemd-resolved 2>/dev/null

    # 4. WLAN-Karte eine feste IP geben
    sudo tee /etc/udev/rules.d/99-wlan0-vars.rules > /dev/null << EOF
SUBSYSTEM=="net", ACTION=="add", KERNEL=="wlan0", RUN+="/usr/sbin/ip addr add $SERVER_IP/24 dev wlan0"
EOF

    sudo udevadm control --reload-rules
    sudo udevadm trigger


    # 5. hostapd Konfiguration schreiben (Das WLAN-Netz)
    sudo bash -c 'cat > /etc/hostapd/hostapd.conf' << EOF
interface=wlan0
driver=nl80211
ssid=$SSID
hw_mode=g
channel=1
country_code=DE
wmm_enabled=0
macaddr_acl=0
auth_algs=1
ignore_broadcast_ssid=0
wpa=2
wpa_passphrase=$PASSWORD
wpa_key_mgmt=WPA-PSK
wpa_pairwise=TKIP
rsn_pairwise=CCMP
EOF

    # 6. dnsmasq Konfiguration schreiben (Vergibt IPs an die ESP32-Westen)
    sudo bash -c 'cat > /etc/dnsmasq.conf' << EOF
interface=wlan0
dhcp-range=$IP_RANGE_LOW,$IP_RANGE_HIGH,255.255.255.0,12h
EOF

    # 7. softlock ausstellen
    echo 'SUBSYSTEM=="rfkill", ATTR{type}=="wlan", ATTR{state}="1"' | sudo tee /etc/udev/rules.d/60-ur-rfkill.rules > /dev/null

    # 8. Dienste aktivieren und starten
    sudo systemctl unmask hostapd
    sudo systemctl enable hostapd dnsmasq
    sudo systemctl restart hostapd
    sudo systemctl restart dnsmasq
fi

echo "===== DATENBANK ====="

# Checke, ob das Setup bereits durchgeführt wurde
if [ -f "$SCRIPT_DIR/.env" ]; then
    echo "Datenbank Setup wurde bereits durchgeführt. Überspringe die Einrichtung."
    echo "Sollte das nicht der Fall sein, lösche die .env Datei und starte das Setup erneut."
    echo "Stelle auch sicher, dass es keine PostgreSQL Datenbank und keinen Datenbank-Nutzer gibt."
    echo "Bestehende Daten werden nach dem Setup nicht mehr berücksichtigt."
else

    echo "#.env" > "$SCRIPT_DIR/.env"

    sudo apt update
    sudo apt install -y postgresql

    # Datenbank Zugangsdaten (.env) setup
    read -p "Datenbank-Nutzername [lasertag]: " DB_USER
    DB_USER=${DB_USER:-lasertag}

    read -s -p "Passwort für $DB_USER (Eingabe ist nicht sichtbar): " DB_PASSWORD
    echo ""

    read -p "Name der Datenbank [lasertag]: " DB_NAME
    DB_NAME=${DB_NAME:-lasertag}

    echo "DB_NAME=$DB_NAME" >> "$SCRIPT_DIR/.env"
    echo "DB_USER=$DB_USER" >> "$SCRIPT_DIR/.env"
    echo "DB_PASSWORD=$DB_PASSWORD" >> "$SCRIPT_DIR/.env"

    # PostgreSQL Datenbank und Benutzer setup
    sudo -u postgres psql -c "CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';"
    sudo -u postgres psql -c "CREATE DATABASE $DB_NAME OWNER $DB_USER;"
    sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;"
fi

sudo apt update
sudo apt install -y golang

echo ""
echo "/===== SETUP BEENDET =====/"
echo "System wird neu gestartet..."
sleep 5000
sudo reboot
