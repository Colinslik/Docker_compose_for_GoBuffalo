# Fix Docker DNS problem

1. Discover the address of your DNS server
   $ nmcli dev show | grep 'IP4.DNS'
   IP4.DNS[1]:                             10.0.0.2

2. Update the Docker daemon
   $  vi /etc/docker/daemon.json
   
   Add this line:
   {
    "dns": ["10.0.0.2", "8.8.8.8"]
   }
