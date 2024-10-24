# Technical challenge solutions

I summarise here the steps followed in order to implement the different levels:

Level 1:
* Environment Setup:
    * Ubuntu EC2 Instance Creation: I Launched an Ubuntu EC2 instance on AWS to serve as the foundation for my setup.
    * Docker Installation: On the EC2 instance, I installed Docker to manage containerised applications.
* Rancher Setup:
    * Single-Node Rancher Installation: I deployed a single-node Rancher instance using Docker
    * User Creation for Rancher: Once Rancher was operational, I accessed the Rancher UI and created a new user to manage the instance
        * rancher-standard/WaczB8tSX0BppPyyYAtU
* Automation with Cypress:
    * Cypress Installation on my Project
    * Cypress Script Implementation: Login and verifications
        * To run cypress test case: npx cypress open --config baseUrl=https://44.214.93.44/  (being baseUrl the rancher url)
        * Script: home_page.cy.js
Evidence:
￼![Image1](https://github.com/user-attachments/assets/676b3cf8-28ab-4cc4-b3d1-905dc00ef192)

 Level 2:
* Go Installation on the project
* API Key Creation in Rancher UI: I created an API key in the Rancher UI to enable secure API interactions.
    * Automated E2E API Testing. The test included:
        * Login Automation: Using the API key, the script logged into Rancher.
            * Access key: token-j79bk
            * Secret key: j97wmb4nm54v88ff79lqlb2rhl9n97xvhl2zlpp2dw6vrpfznjjbxf
            * Bearer token: token-j79bk:j97wmb4nm54v88ff79lqlb2rhl9n97xvhl2zlpp2dw6vrpfznjjbxf
        * API Endpoint Verification
        * Script: login_test.go
Evidence:
![Image2](https://github.com/user-attachments/assets/fbbbf8c2-9294-4d5a-94ec-90a33acf3aa3)

 Level 3:
Not familiar with GCP, I would need more time to investigate and complete the exercise. These are the steps I run (I selected Google API - Python option):

* Install the Google Cloud Client Library
* Set Up Authentication
    * Create a service account with necessary permissions
    * Download the JSON key file
    * Set the GOOGLE_APPLICATION_CREDENTIALS environment variable to point to the key file.
* Create a Python Script: create_vm.py.
