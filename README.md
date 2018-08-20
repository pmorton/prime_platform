This repository contains a golang web application that determines if a number is prime or not. Our software is code complete and needs to be deployed to a production environment. We have engaged you to create a build process for this software which should include building the application code and deploying it to a production environment. We don't have any infrastructure and we know that we want to host this application in a public IaaS cloud such as AWS or GCP. Further we abhor manual infrastructure build and code deployment cycles and wish to automate as much as possible. Our team is familiar with tools such as terraform, docker and packer however we are always looking to add quality tools our toolbox. 

At a minimum the solution should:

1. Build the application binary
2. Provision the infrastructure necessary to deploy this application
3. Deploy the application
4. Apply general best practices around scale and performance

Testing your deployment:

Once the application is deploy you should be able to test a prime number by submitting a get request with the `test` parameter. For example `curl -v "localhost:9090?test=2"` will return `true`
