@ECHO off
ECHO "Se initializeaza imprimanta fiscala..."

set arg1=%1
Start "ExchangeV2" /B "%UserProfile%\Desktop\ExchangeV2\apiServer.exe" "%arg1%"

PAUSE