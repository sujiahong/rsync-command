import os, time

ret = os.system("rsync -vzrhtopg --progress --delete --bwlimit=2000 'rsync://10.0.1.210:8765/SteamLibrary/steamapps/common/Mafia III' /cygdrive/e/SteamLibrary/steamapps/common")
print("J=======  ", ret)



def doFileCopy(src, tag):
    if (os.path.isfile(src)):
        return
    for item in os.listdir(src):
        sourcePath = os.path.join(src, item)
        targetPath = os.path.join(tag, item)
        if (os.path.isdir(sourcePath)):
            if (not os.path.exists(targetPath)):
                os.makedirs(targetPath)
            doFileCopy(sourcePath, targetPath)
        elif(os.path.isfile(sourcePath)):
            ret = os.system('xcopy /e /s /f /y /i /c /d "{}" "{}"'.format(sourcePath, tag))
            #time.sleep(0.2)
            print("copy ret ", ret)
        else:
            print(r"什么情况，不是目录也不是文件：", sourcePath)

#doFileCopy(r"E:\\SteamLibrary\\steamapps\\common\\HatinTime", r"F:\\SteamLibrary\\steamapps\\common\\HatinTime")
    