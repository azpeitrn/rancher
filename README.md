# rancher

SUSE Quality Assurance Engineer – Technical Challenge
This is an exercise to help us get a sense of how you think. Don't worry about knowing too much about
the products in question but don’t shy away from reaching out to us for information.
Your submission will be reviewed by a member of our engineering team, who will check your response for
readability, simplicity and functional completeness.
There are 3 levels to this challenge. We know how busy life can get these days, so we hope you manage
to find time to complete the challenge.
Level 1: UI Automation using Cypress framework
Implement an automated e2e test for single node install of Rancher UI
Rancher single node Install: https://ranchermanager.docs.rancher.com/pages-for-
subheaders/installation-and-upgrade#single-node-kubernetes-install
Docker install: https://ranchermanager.docs.rancher.com/pages-for-subheaders/installation-and-
upgrade#docker-install
If you don’t want to use Docker Desktop, you can install Rancher Desktop and select moby as container
engine to use docker client (https://docs.rancherdesktop.io/tutorials/working-with-containers#running-
containers).
Cover these 3 cases in your automation script.
- Login into Rancher web page
- Check if the main web page opens up.
- Check if the main web page title is correct
  Please implement the solution in your own repository and share your repository link with us.
  Level 2: API Automation using Go lang (standard test framework or ginkgo)
  Implement e2e API test for single node install of Rancher UI
  Test to cover.
- Login into Rancher
  Please implement the solution in your own repository and share your repository link with us.
  Level 3: Deploy a VM on GCP
  Create a script using Terraform, Google API or Ansible to deploy a VM on GCP
  Google provides free credits which can be used to test the script.
  Please implement the solution in your own repository and share your repository link with us.