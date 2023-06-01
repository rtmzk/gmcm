HOST={{ iterate . "ip" }}
PORT={{ (index . 0).SSHPort }}
NOPASS={{ (index . 0).NoPass }}
PASSWORD={{ iterate . "passwd" }}