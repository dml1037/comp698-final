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
comp698-final/static/image.png
 
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


