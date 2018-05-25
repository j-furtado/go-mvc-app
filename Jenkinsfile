//
//  Params to set on the pipeline:
//    * ACR_CREDS: Credentials for Azure Container Registry
//    * ACR_LOGINSERVER: Azure Container Registry login server
//    * AZ_CREDS: Azure Service Principal Credentials
//    * KUBE_RSGRP: Azure resource group name for Kubernetes
//    * KUBE_SERVICE: Azure Kubernetes service name
//    * KUBE_CREDS: SSH Credentials for Kubernetes Master
//    * KUBE_SECRET: Kubernetes secret to access the Azure Container Registry
//    * POSTGRES_ACC: Postgres account credentials
//    * DB_DIR: directory where the postgres DB will live
//    * DB_NAME: name of the postgres DB
//
pipeline {
  agent any

  environment {
    // Sets the variables to be used in Kubernetes
    POSTGRES_ACC = credentials("${params.PG_ACC}")
    // will create a POSTGRES_ACC_USR and POSTGRES_ACC_PSW
    DB_DIR = "${params.PG_DIR}"
    DB_NAME = "${params.PG_BD}"
    MOUNT_PATH = "${params.MOUNT_DIR}"
    DEPLOY_DB = "${params.DEPLOY_DB}"
  }

  stages{
    stage('Clone Repo'){
      steps{
        // grab the source code from the repo
        checkout scm
      }
    }

    stage('Check if DB needs to deploy') {
      steps {
        // Builds the container image
        script {
          if(DEPLOY_DB) {
            stage('Deploy new DB image') {
              sh 'docker pull postgres:latest'
              withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: "${params.ACR_CREDS}",
                usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
                  sh "docker login -u $USERNAME -p $PASSWORD https://${params.ACR_LOGINSERVER}"
                  sh "docker tag postgres:latest ${params.ACR_LOGINSERVER}/postgres:latest"
                  sh "docker push ${params.ACR_LOGINSERVER}/postgres:latest"
                  sh "docker build -f 'Dockerfile-db' -t ${params.ACR_LOGINSERVER}/pgdb ."
                  sh "docker push ${params.ACR_LOGINSERVER}/pgdb"
              }
            } // Stage close
          } // If close
        } // script close
      } // steps close
    } // stage close

    stage('Build App Image') {
      steps{
        // Builds the container image
        sh 'docker pull golang:latest'
        withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: "${params.ACR_CREDS}",
          usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
            sh "docker login -u $USERNAME -p $PASSWORD https://${params.ACR_LOGINSERVER}"
            sh "docker tag golang:latest ${params.ACR_LOGINSERVER}/golang:latest"
            sh "docker push ${params.ACR_LOGINSERVER}/golang:latest"
            sh "docker build -f 'Dockerfile-app' -t ${params.ACR_LOGINSERVER}/gomvc:$BUILD_NUMBER ."
        }
      }
    }
    stage('Push App Image') {
      steps{
        // Pushes the image to the registry
        withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: "${params.ACR_CREDS}",
          usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
          sh "docker login -u $USERNAME -p $PASSWORD https://${params.ACR_LOGINSERVER}"
          sh "docker push ${params.ACR_LOGINSERVER}/gomvc:$BUILD_NUMBER"
        }
      }
    }
    stage('Deploy') {
      steps{
        // Deploys a container with the generated container image
        acsDeploy(azureCredentialsId: "${params.AZ_CREDS}",
            resourceGroupName: "${params.KUBE_RSGRP}",
            containerService: "${params.KUBE_SERVICE} | Kubernetes",
            sshCredentialsId: "${params.KUBE_CREDS}",
            configFilePaths: '*.yaml',
            enableConfigSubstitution: true,
            secretName: "${params.KUBE_SECRET}",
            secretNamespace: 'default',
            containerRegistryCredentials: [
                [credentialsId: "${params.ACR_CREDS}", url: "https://${params.ACR_LOGINSERVER}"]
        ])
      }
    }
  }
}
