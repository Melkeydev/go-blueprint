package cicd

type JenkinsTemplate struct{}

func (j JenkinsTemplate) Dockerfile() []byte {
    return MakeDockerfile()
}

func (j JenkinsTemplate) Pipline() []byte {
	return []byte(`pipeline {
    agent any

    environment {
        GITHUB_USER = ''
        GITHUB_REPO = ''
        DOCKER_HUB_USERNAME = ''
        DOCKER_REPO_NAME = ''
        BRANCH = ''
        VERSION_PART = 'Patch' // Patch, Minor, Major
        DOCKER_JENKINS_CERDIDENTALS_ID = ''
        TAG = '' // Generated automatically
    }

    stages {
        stage('Checkout Code') {
            steps {
                script {
                    git(url: "https://github.com/${GITHUB_USER}/${GITHUB_REPO}/", branch: env.BRANCH_NAME)
                }
            }
        }

        stage('Test') {
            steps {
                script {
                    sh "go test -v ./..."
                }
            }
        }

        stage('Generate Docker Image Tag') {
            when {
                expression { env.BRANCH_NAME == env.BRANCH}
            }
            steps {
                script {
                    TAG = sh(script: "/home/jenkins/docker_tag.sh $DOCKER_HUB_USERNAME $DOCKER_REPO_NAME $VERSION_PART", returnStdout: true).trim()

                    if (TAG) {
                        echo "Docker image tag generated successfully: $TAG"
                    } else {
                        error "Failed to generate Docker image tag"
                    }

                    env.TAG = TAG
                }
            }
        }

        stage('Docker Login') {
            when {
                expression { env.BRANCH_NAME == env.BRANCH }
            }
            steps {
                script {

                    withCredentials([usernamePassword(credentialsId: env.DOCKER_JENKINS_CERDIDENTALS_ID, passwordVariable: 'DOCKER_PASSWORD', usernameVariable: 'DOCKER_USERNAME')]) {
                        sh "docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD"
                    }
                }
            }
        }

        stage('Build') {
            when {
                expression { env.BRANCH_NAME == env.BRANCH }
            }
            steps {
                script {
                    sh "docker build --no-cache -t ${DOCKER_HUB_USERNAME}/${DOCKER_REPO_NAME}:${TAG} ."
                }
            }
        }

        stage('Deploy') {
            when {
                expression { env.BRANCH_NAME == env.BRANCH }
            }
            steps {
                script {
                    sh "docker push ${DOCKER_HUB_USERNAME}/${DOCKER_REPO_NAME}:${TAG}"
                }
            }
        }

        stage('Environment Cleanup') {
            when {
                expression { env.BRANCH_NAME == env.BRANCH }
            }
            steps {
                script {
                    sh "docker rmi ${DOCKER_HUB_USERNAME}/${DOCKER_REPO_NAME}:${TAG}"
                }
            }
        }
    }

    post {
        success {
            echo "Pipeline completed successfully"
        }
    }
}
	`)
}