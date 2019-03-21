import os
import sys
aws_uri = "need to know"
default_tag = "latest"
dev_context = "dev-context"
cronjob = "go-cli-cronjob"
containter_tag = "go-cli-prompt"
def script(version=default_tag, environment =dev_context ):
    cwd = os.getcwd()
    print(cwd)
    os.system("kubectl config use-context "+environment)
    os.system("kubectl set image cronjob/"+cronjob+" "+containter_tag+"="+aws_uri+":"+version)
    os.system("kubectl patch cronjob "+cronjob+" -p \"{\\\"spec\\\":{\\\"template\\\":{\\\"metadata\\\":{\\\"labels\\\":{\\\"date\\\":\\\"`date +'%s'`\\\"}}}}}\" ")

script(sys.argv[1],sys.argv[2])
