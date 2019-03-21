import os
import sys
aws_uri = "need to know"
default_tag = "latest"
def script(version=default_tag):
    cwd = os.getcwd()
    print(cwd)
    default_image=aws_uri+":"+version
    os.system("sudo docker build -t "+default_image+" .")
    os.system("sudo docker push " + default_image)
script(sys.argv[1])
