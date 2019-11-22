node {
    def app

    stage('Clone repository') {
        /* Let's make sure we have the repository cloned to our workspace */

        checkout scm
    }

    stage('Build image') {
        /* This builds the actual image; synonymous to
         * docker build on the command line */

        app = docker.build("action-list", "./action-list")
    }

    stage('Push image') {
        /* Finally, we'll push the image with two tags:
         * First, the incremental build number from Jenkins
         * Second, the 'latest' tag.
         * Pushing multiple tags is cheap, as all the layers are reused. */
         app.push("latest")
//         docker.withRegistry('https://hub.docker.com', 'dockerhub') {
//             app.push("latest")
//         }
    }

     stage('Build image') {
            /* This builds the actual image; synonymous to
             * docker build on the command line */
            app = docker.build("create-project", "./create-project")
        }

        stage('Push image') {
            /* Finally, we'll push the image with two tags:
             * First, the incremental build number from Jenkins
             * Second, the 'latest' tag.
             * Pushing multiple tags is cheap, as all the layers are reused. */
            docker.withRegistry('https://hub.docker.com', 'dockerhub') {
                app.push("latest")
            }
        }
}