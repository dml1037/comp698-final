# comp698-final
1) Why use git?
One of the biggest advantages of Git is its branching capabilities. Unlike centralized version control systems, Git branches are cheap and easy to merge. This facilitates the feature branch workflow popular with many Git users. It is a distributed version control system. Instead of a working copy, each developer gets their own local repository, complete with a full history of commits. Many source code management tools such as Bitbucket enhance core Git functionality with pull requests. A pull request is a way to ask another developer to merge one of your branches into their repository. 
Reference: https://www.atlassian.com/git/tutorials/why-git

2) Why write in golang? 
It is fast. And not only fast in the sense that programs written in it run fast when compared to other common languages; but also fast in the sense that its compiler can compile projects in the blink of an eye. It is a garbage-collected language. It has built-in concurrency, which allows parallelism in an easier way than is possible in other languages. Go has documentation as a standard feature. That makes it easier for developers to document their code and generate human-readable data out of source code comments. Go has a rich standard library which covers a lot of areas. In fact, Go is probably the only language that can claim to have a fully working Web server as part of its standard library.
Reference: https://altabel.wordpress.com/2015/11/10/golang-pros-and-cons/

3) What purpose does Travis CI serve?
“Travis CI is a hosted, distributed continuous integration service used to build and test projects hosted at GitHub. Travis CI automatically detects when a commit has been made and pushed to a GitHub repository that is using Travis CI, and each time this happens, it will try to build the project and run tests. This includes commits to all branches, not just to the master branch.” 
Travis CI runs your programs test when you add one to GitHub and verifies that everything builds correctly. It double checks that there are no bugs in your code and makes it easier to discover if your commit broke something and fixes it before it becomes a problem. 

4) What is the purpose of terraform?
Terraform is a tool for building, changing, and versioning infrastructure safely and efficiently. Terraform can manage existing and popular service providers as well as custom in-house solutions. Configuration files describe to Terraform the components needed to run a single application or your entire datacenter. Terraform generates an execution plan describing what it will do to reach the desired state, and then executes it to build the described infrastructure. As the configuration changes, Terraform is able to determine what changed and create incremental execution plans which can be applied. The infrastructure Terraform can manage includes low-level components such as compute instances, storage, and networking, as well as high-level components such as DNS entries, SaaS features, etc.
Reference: https://www.terraform.io/intro/index.html

5) Why use virtualized resources?
Virtualization lets companies reduce their IT costs by requiring fewer hardware servers and related resources to achieve the same level of computing performance, availability and scalability. IT can greatly reduce the ongoing administration and management of manual, time-consuming processes by automating operations, thus resulting in lower operational expenses. As companies reduce the size of their hardware and server footprint, they lower their energy consumption, cooling power and data center square footage, thus resulting in lower costs.

6) Diagram of the architecture from laptop to GCP
![Architecture image](/static/image.png)
 
7) Why use bash commands vs clicking through UI?
There are many reasons why you should use bash commands over UI. First, you can gain greater control over system functions. You can use Node Package Manager for package installs. You can utilize Git version control. You need it to use preprocessors and task runners. It is used in local backend development. And finally, anyone who wants to work in the software industry, it is imperative that you know command line. Why did you think you could be a programmer and not use it?!

8) Why docker vs installing application directly on the host?
With installing application directly, you are more likely to run into configuration issues when there are two applications. The runtime isolation has configurable resource limits and the ports reconfig even by third party or legacy software. Integrate deployment of third-party or legacy software in your standard Docker deployment. It is easier to participate in cloud. As soon you package to standard container and deploy to cloud, you profit from cloud features you have (e.g. hot migration, automatic backup, auto scaling).
Reference: http://alexander.holbreich.org/package-manager-vs-docker/

9) What protections are there against accidentally introducing bad code into production?
* Work on bugs and features on separate branches
* Do not use deployments for rapidly changing environments; run software locally
* Test, test, test!
* Do not test on master

DOCUMENTATED STEPS TAKEN FOR THIS PROJECT
Documentation for Comp 698 Final
1) Create a new GitHub repo
* Created a new folder on my local desktop under UNH/comp698/comp698-final
* Created a GitHub repository online
* Initialized it with a Readme file
* Chose MIT license 
* Copied clipboard link, cloned it to my laptop “git clone https://github.com/dml1037/comp698-final.git” in my comp698 folder.
* Went to settings, branches, then branch protection rules. Chose Master, protect this branch, checked require status checks to pass before merging, select include administrators, initialized Travis-ci.

2) Initial hello world application
* Created two code files in sublime on my local; main.go and main_test.go.
* Performed git add main.go and git add main_test.go, git commit –m “adding main.go and main_test.go” to master, git push origin master to add both files to GitHub repo. Got a pull request, waited for approval, merged. The code for those files was found at https://github.com/mathyourlife/comp698-final/tree/helloworld-source

3) Add Dockerfile
* Created a local branch called addDock (git branch addDock)
* Created a new code file on sublime called dockerfile with the contents from   https://github.com/mathyourlife/comp698final/blob/helloworld-dockerfile/Dockerfile
* added to branch/committed/pushed (git add dockerfile, git commit –m “adding dockerfile for hello world application”, git push origin addDock)
* Got pull request on GitHub, merged successful.
* Everything was working, except I noticed when I copied and pasted the code, there were “ + ”’s at the beginning of each line. So, I locally fixed the file. I did the same process as above. (git add dockerfile, git commit –m “adding dockerfile for hello world application”, git push origin addDock). However, I got a newline error. Travis failed it as well because it was not set up yet. I fixed it locally, then did the same process and it worked.

4) Add Travis ci PR validation
5) Update Travis ci for PR validation and functional testing
* Created a .travis.yml and functional-test.sh file locally using sublime. The contents were from https://github.com/mathyourlife/comp698-final/tree/helloworld-travisci
* git add .travis.yml, git add functional-test.sh, git commit –m “adding travis ci for pull request validation”, git push origin addDock).
* When trying to merge the .travis.yml pull request, it failed multiple times because I had not yet added the functional-test.sh file. Once I merged that file, both were merged and added to GitHub successfully.

6) Add GCP build trigger
* Went to GCP then Build Trigger
* Created trigger with settings the module on Canvas (GitHub, comp698 final, branch master, image name gcr.io/$PROJECT_ID/$REPO_NAME:$COMMIT_SHA)
* Ran trigger, waited for successful message

7) Create terrform to deploy hello world version to 1 staging server
* Created a subfolder in my local comp698 folder called “terraform”. Created main.tf file based on “my-web-server” project.
* git branch terraformFile, add main.tf, git commit –m “adding main.tf to terraform”, git push origin terraformFile).
* Got pull request, validated Travis Ci, merged successful
* sshed into terraform-configuration server (gcloud compute ssh terraform-configuration) 
* git clone comp698-final 
* cd into UNH/comp698/comp698-final/terraform
* ran “terraform apply”
* got an error with the bucket name
* updated my sublime code to change it to a more unique name
* added/commited/pushed/pull request/merge
* git pull on ssh window to make sure new file was up to date
* ran terraform apply, it threw an error, ran terraform destroy, then terraform apply again. Ran successful. Deployed “hello world”.

8) Update application code to full
* Updated every single file found here: https://github.com/mathyourlife/comp698-final/tree/bootstrap-source and here: https://github.com/mathyourlife/comp698-final/tree/bootstrap-source
* added/committed/push/merge
* I did get some errors in this process. The reason being because some needed others to run successful. I basically did them out of order. Once everything was in there it was successful.
* added final touches to main.tf. Added updated image name (path to GCP)
* Did the same terraform steps as in step 7. Sshed, updated main.tf, ran terraform apply, image build successful.

9) Update terraform to deploy full app to 1 staging server and hello world to 1 prod server
* Updated my main.tf file one last time to input the image names (long paths to GCP)
* When testing both external IP’s, I noticed that they were displaying backwards.
* Went in to my local files, switched the paths, added/committed/pushed/pull request/merged
* updated terraform using terraform apply in ssh window
* problem solved, when you click each external IP they both go to the correct windows.
* deleted branches “addDock” and “terraformFile” locally and remotely



