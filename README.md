# Web server in HTML/javascript + Go (Golang) backend + supporting database
This project aims to provide a sample app to feed to the CI/CD pipeline created in a previous project.

The link to the CI/CD pipeline project is: https://github.com/j-furtado/azure-terraform

## Getting Started
These instructions will get you a copy of the project up and running on your CI/CD environment. It is assumed that you followed the previous guide in all aspects.

### Prerequisites
Things you need to have:
* **Git** - integrated with your SDK/Editor
* **CI/CD Pipeline** - created previously

### Forking the app
To be able to make your own changes, you should fork the code to your own repo/account.

### Creating the CI/CD pipeline
Assuming your Jenkins server is already installed and configured, let's jump straight into the pipeline.

On the main dashboard, click **New Item** on the left side menu. Choose your project name and select the **Pipeline** option. You will then be presented with the pipeline configuration.

Select the option **GitHub Project** and enter the URL of your project.

Then select **This project is parameterized**. Since the Jenkinsfile that supports the pipeline has some variables, you'll have to set them here. If you check the Jenkinsfile, it should state in the begining what variables you'll need. Here is the list:
* **ACR_CREDS**: Credentials for Azure Container Registry. Select the type **Credentials Parameter**
* **ACR_LOGINSERVER**: Azure Container Registry login server. Select the type **String Parameter**
* **AZ_CREDS**: Azure Service Principal Credentials. Select the type **Credentials Parameter**
* **KUBE_RSGRP**: Azure resource group name for Kubernetes. Select the type **String Parameter**. Check either the Terraform code or Azure.
* **KUBE_SERVICE**: Azure Kubernetes service name. Select the type **String Parameter**. Check either the Terraform code or Azure.
* **KUBE_CREDS**: SSH Credentials for Kubernetes Master. Select the type **Credentials Parameter**
* **KUBE_SECRET**: Kubernetes secret to access the Azure Container Registry. Select the type **String Parameter**. If not created already, you must create it in **Kubernetes**. You will set here the secret name.

The names must match those above. If you already created the credentials with the infrastructure, most of these you only need to select.

Then select **GitHub hook trigger for GITScm polling**. This will allow you to make commits and have the infrastructure update itself automatically.

Finally, you'll have to set which pipeline file to load (the Jenkinsfile). Since the Jenkinsfile is part of the source code of the app, you just need to point to the same GitHub repo. Select **Pipeline script from SCM**, then **Git** as your SCM. In the Repository URL, set the **project url**.

When you're ready, save the pipeline.

### Triggering the pipeline
The first build you'll have to trigger manually. To do that, on the Pipeline page, press the **Build with Parameters** option on the right. It will ask you to confirm the Parameters (which you can change if you want to point to other environments for a manual build). Press **Build** and you'll see it launch.

You probably will have to refresh the page if you want to see the graphical **Stage View**. It will display all the stages in the pipeline and the times it takes to build it.

After it deploys with success, you can check your Kubernetes cluster (with the dashboard + proxy method), to see what is the Cluster IP you have to access to see the App. Use the IP from the load balancer, with the port, and you should see your Website. The first build takes a bit longer, since Kubernetes has to create the Load Balancer.

Future builds will be on the same IP:Port, but will follow a rolling update method, preserving one replica alive all the time. So if you don't see your changes imediatly, keep refreshing the page. It should be ready in a few (10s-15s) seconds.
