import os

ret = os.system("rsync -vzrhtopg --progress --delete --bwlimit=2000 'rsync://10.0.1.210:8765/test/steamapps/common/Farming Simulator 19' /cygdrive/e/SteamLibrary/steamapps/common")
print("J=======  ", ret)