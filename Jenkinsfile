node {
    def app

    triggers {
        pollSCM '' //empty quotes tells it to build on a push
    }

    stage('Clone repository') {
        checkout scm
    }

    stage('Build image for action-list') {
        app = docker.build("atifdockerventure/action-list", "./action-list")
    }

    stage('Push image for action-list') {
        docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
            app.push("latest")
        }
    }

     stage('Build image for create-project') {
            app = docker.build("atifdockerventure/create-project", "./create-project")
        }

     stage('Push image for create-project') {
         docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
                app.push("latest")
         }
     }

     stage('Build image for identifier-repository') {
         app = docker.build("atifdockerventure/identifier-repository", "./identifier-repository")
     }

     stage('Push image for identifier-repository') {
         docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
             app.push("latest")
         }
     }

      stage('Build image for test-repository') {
             app = docker.build("atifdockerventure/test-repository", "./test-repository")
         }

      stage('Push image for test-repository') {
          docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
                 app.push("latest")
          }
      }

      stage('Build image for test-config') {
          app = docker.build("atifdockerventure/test-config", "./test-config")
      }

      stage('Push image for test-config') {
          docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
              app.push("latest")
          }
      }
      stage('Build image for data-setup') {
          app = docker.build("atifdockerventure/data-setup", "./data-setup")
      }

      stage('Push image for data-setup') {
          docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
              app.push("latest")
          }
      }
}