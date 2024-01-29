pipeline {
    agent any
    options {
        disableConcurrentBuilds()
    }
    stages {
        stage('Checkout'){
            steps {
                checkout scm
            }
        }
        stage('Build') {
            steps {
                script {
                    env.REAL_PWD = getDockerPWD();
                    sh 'docker run --rm -w /app -v $REAL_PWD:/app golang:alpine go build'
                }
            }
        }
        stage('Test') {
            steps {
                script {
                    env.REAL_PWD = getDockerPWD();
                    sh 'docker run --rm -w /app -v $REAL_PWD:/app golang:alpine go test ./...'
                }
            }
        }
        stage('Prep buildx') {
            when { branch 'master' }
            steps {
                script {
                    env.BUILDX_BUILDER = getBuildxBuilder();
                }
            }
        }
        stage('Dockerhub login') {
            when { branch 'master' }
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub', usernameVariable: 'DOCKERHUB_CREDENTIALS_USR', passwordVariable: 'DOCKERHUB_CREDENTIALS_PSW')]) {
                    sh 'docker login -u $DOCKERHUB_CREDENTIALS_USR -p "$DOCKERHUB_CREDENTIALS_PSW"'
                }
            }
        }
        stage('Build Fetcher Docker Image') {
            when { branch 'master' }
            steps {
                sh """
                    docker buildx build --pull --builder \$BUILDX_BUILDER --target fetcher --platform linux/arm64,linux/amd64 -t nbr23/atomic-banquet:latest -t nbr23/atomic-banquet:`git rev-parse --short HEAD` --push .
                    """
            }
        }
        stage('Build Server Docker Image') {
            when { branch 'master' }
            steps {
                sh """
                    docker buildx build --pull --builder \$BUILDX_BUILDER  --target server --platform linux/arm64,linux/amd64 -t nbr23/atomic-banquet:server-latest -t nbr23/atomic-banquet:server-`git rev-parse --short HEAD` --push .
                    """
            }
        }
        stage('Build Nginx Server Docker Image') {
            when { branch 'master' }
            steps {
                sh """
                    docker buildx build --pull --builder \$BUILDX_BUILDER  --target server --platform linux/arm64,linux/amd64 -t nbr23/atomic-banquet:server-nginx-latest -t nbr23/atomic-banquet:server-nginx-`git rev-parse --short HEAD` --push .
                    """
            }
        }
        stage('Sync github repos') {
            when { branch 'master' }
            steps {
                syncRemoteBranch('git@github.com:nbr23/atomic-banquet.git', 'master')
            }
        }
    }
    post {
        always {
            sh 'docker buildx stop $BUILDX_BUILDER || true'
            sh 'docker buildx rm $BUILDX_BUILDER || true'
        }
    }

}
