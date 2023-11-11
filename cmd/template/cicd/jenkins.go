package cicd

type JenkinsTemplate struct{}

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

func (j JenkinsTemplate) JenkinsSlave() []byte {
	return []byte(`FROM jenkins/ssh-agent

WORKDIR /home/jenkins

COPY . .

RUN apt update && \
	apt -y install jq ca-certificates gnupg software-properties-common wget curl git python3-pip python3.11 python3-venv python3.11-venv python3-dev python3.11-dev unzip zip libcurl4-openssl-dev libssl-dev

RUN wget https://golang.org/dl/go1.20.2.linux-amd64.tar.gz && \
	tar -C /usr/local -xzf go1.20.2.linux-amd64.tar.gz

RUN wget https://github.com/go-task/task/releases/download/v3.31.0/task_linux_amd64.tar.gz && \
	tar -xzf task_linux_amd64.tar.gz && \
	cp task /usr/bin/task && \
	chmod +x /usr/bin/task

ENV PATH="/usr/local/go/bin:${PATH}:/home/jenkins/bin"
ENV GOPATH="/home/jenkins/go"
ENV PATH="${PATH}:${GOPATH}/bin"

RUN go install github.com/jstemmer/go-junit-report/v2@latest 

RUN install -m 0755 -d /etc/apt/keyrings && \
	curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg && \
	chmod a+r /etc/apt/keyrings/docker.gpg && \
	echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian $(. /etc/os-release && echo $VERSION_CODENAME) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null && \
	apt update && \
	apt -y install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin docker-ce-cli

RUN chmod +x docker_tag.sh

CMD ["/bin/sh", "-c", "setup-sshd && dockerd --host=unix:///var/run/docker.sock"]

#docker run -it -v /var/run/docker.sock:/var/run/docker.sock <image> /bin/bash
	`)
}

func (j JenkinsTemplate) DockerTag() []byte {
	return []byte(`#!/bin/bash

DOCKER_HUB_USERNAME=$1
DOCKER_REPO_NAME=$2
VERSION_PART=$3

DOCKER_IMAGE="$DOCKER_HUB_USERNAME/$DOCKER_REPO_NAME"

TAGS=$(curl -s "https://hub.docker.com/v2/repositories/$DOCKER_IMAGE/tags/?page_size=100" | jq -r '.results[].name')

if [ -z "$TAGS" ]; then
	DEFAULT_TAG="1.0.0"
	NEW_TAG="$DEFAULT_TAG"
else
	LATEST_TAG=$(echo "$TAGS" | grep -E '^[0-9]+\.[0-9]+\.[0-9]+$' | sort -V | tail -n 1)

	if [ -z "$LATEST_TAG" ]; then
		LATEST_TAG="1.0.0"
	fi

	IFS='.' read -ra PARTS <<< "$LATEST_TAG"
	MAJOR=${PARTS[0]}
	MINOR=${PARTS[1]}
	PATCH=${PARTS[2]}

	if [[ "$VERSION_PART" == "Major" ]]; then
		NEW_TAG="$((MAJOR + 1)).0.0"
	elif [[ "$VERSION_PART" == "Minor" ]]; then
		NEW_TAG="$MAJOR.$((MINOR + 1)).0"
	elif [[ "$VERSION_PART" == "Patch" ]]; then
		NEW_TAG="$MAJOR.$MINOR.$((PATCH + 1))"
	else
		echo "Invalid version part specified. Usage: $0 [Major|Minor|Patch]"
		exit 1
	fi
fi

echo $NEW_TAG
	`)
}


func (j JenkinsTemplate) Dockerfile() []byte {
    return []byte(`FROM golang:1.21-alpine

WORKDIR /app
	
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/api/main.go

EXPOSE 8080

CMD ["./main"]
	`)
}


func (j JenkinsTemplate) JenkinsReadme() []byte {
    return []byte(`# Jenkins Configuration and Pipeline Setup

This guide offers short instructions for configuring Jenkins Controller and Agents to facilitate the execution of pipeline jobs. While Jenkins can be set up locally using Docker Compose, configuring a dedicated server is recommended. Numerous tutorials are available for controller installation, and you can refer to the official documentation.

## Step 1: Retrieve Jenkins Controller Initial Admin Password

After installation, whether as a local Docker container or on a server, locate the InitialAdminPassword in /var/jenkins_home/secrets/initialAdminPassword.

    cat /var/jenkins_home/secrets/initialAdminPassword

Copy the generated password.

## Step 2: Initial Setup

- Open Jenkins in a web browser.
- Paste the initial admin password.
- Follow the on-screen instructions to complete the setup.

## Step 3: Additional Configurations

Enhance your Jenkins setup with the following steps, consulting the Jenkins documentation for clarity on any ambiguous steps.

1. **Install Necessary Plugins:**
   - Navigate to "Manage Jenkins" > "Manage Plugins" > "Available."
   - Search and install the required plugins for your pipeline jobs.

2. **Create GitHub App Webhook:**
   - In your GitHub repository, go to "Settings" > "DeveloperSettings" > "GitHubApp."
   - This step might be complex; refer to the official documentation and accompanying video content.
   - If deploying on a local machine without a server and domain, consider using ngrok as a reverse proxy.

3. **Add GitHub and Docker Credentials:**
   - Navigate to "Manage Jenkins" > "Manage Credentials" to add credentials for GitHub, SSH Agent, and Docker (if needed).

4. **Create SSH Key in Jenkins Controller:**
   - Generate an SSH key in Jenkins for authenticating with version control systems.
   - Run ssh-keygen -t ed25519 -f ~/.ssh/jenkins_agent_key and grab the public key.

5. **Create SSH Agent for Your Job:**
   - Set up an SSH agent for your Jenkins job, avoiding running it on the controller for security reasons.
   - Details about configuration are in the SSHagent folder.
   - Run the following Docker command:

     docker run -v /var/run/docker.sock:/var/run/docker.sock -d --rm --name=agent1 -p 22:22 \
     -e "JENKINS_AGENT_SSH_PUBKEY=[your-public-key]" \
     <agent_image>

   - Follow official docs for configuration and pairing with the controller.

6. **Create Jenkins Multibranch Pipeline Job:**
   - The Jenkinsfile at the root of a project is a simple pipeline for generating Docker images and pushing them with tags into DockerHub.
   - The idea is to add a test stage for testing code in any branch and push the image into the repository if tests pass and the branch is the main one.
   - For other branches, the pipeline can be used for code testing.
   - More information about the pipeline is below.

7. **Deploy Your App:**
   - Once the pipeline is in place, every merge with passing tests results in a deployable image.
   - Application deployment can be achieved using Docker Compose and hosting on the cloud, self-hosting services like Collify, Kubernetes, etc.

## Note:
- Adjust configurations based on your specific requirements.
- Always consider security best practices, especially when handling credentials and sensitive information.
- Explore Jenkins documentation for detailed configuration options: Jenkins Documentation

# SSH-Go-Agent

In a Jenkins environment, agents play a crucial role in distributing workload and executing jobs in parallel. This guide illustrates how to set up Jenkins agents using Docker images with SSH. The Dockerfile uses the jenkins/ssh-agent as the base image and installs various tools and dependencies needed for Go development, testing, and containerization. Customize the Dockerfile to include any additional dependencies or tools your project may require.

Once all dependencies are satisfied, build and push the image to DockerHub:

    docker build -t <repoName>/<imageName>:<tagName> .
    docker push <repoName>/<imageName>:<tagName>

## Generating an SSH Key Pair

To set up an SSH key pair, follow these steps:

2. Generate the SSH key pair by running the following command:

    ssh-keygen -t ed25519 -f ~/.ssh/jenkins_agent_key

## Creating a Jenkins SSH Credential

1. Go to your Jenkins dashboard.

2. In the main menu, click on "Manage Jenkins" and select "Manage Credentials."

3. Click on the "Add Credentials" option from the global menu.

4. Fill in the following information:

   - Kind: SSH Username with private key
   - ID: jenkins
   - Description: The Jenkins SSH key
   - Username: jenkins
   - Private Key: Select "Enter directly" and paste the content of your private key file located at ~/.ssh/jenkins_agent_key
   - Passphrase: Fill in your passphrase used to generate the SSH key pair (leave empty if you didn't use one)

## Creating Your Docker Agent

Use the docker-ssh-agent image that you created and pushed into the DockerHub repo:

    docker run -v /var/run/docker.sock:/var/run/docker.sock -d --rm --name=agent1 -p 2222:22 \
    -e "JENKINS_AGENT_SSH_PUBKEY=[your-public-key]" \
    <repoName>/<imageName>:<tagName>

Replace [your-public-key] with your own SSH public key. You can find your public key value by running cat ~/.ssh/jenkins_agent_key.pub on the machine where you created it.

If your machine already has an SSH server running on port 22, consider using a different port for the Docker command, such as -p 2222:22.

### Registering the Agent in Jenkins

1. Go to your Jenkins dashboard.

2. Click on "Manage Jenkins" in the main menu.

3. Select "Manage Nodes and Clouds."

4. Click on "New Node" from the side menu.

5. Fill in the Node/agent name and select the type (e.g., Name: agent1, Type: Permanent Agent).

6. Fill in the following fields:

   - Remote root directory (e.g., /home/jenkins)
   - Label (e.g., agent1)
   - Usage (e.g., only build jobs with label expression)
   - Launch method (e.g., Launch agents by SSH)
     - Host (e.g., localhost or your IP address)
     - Credentials (e.g., jenkins)
     - Host Key Verification Strategy (e.g., Manually trusted key verification)
     - Change the port if needed; in my case, I need to use port 2222

# Jenkins Pipeline

The pipeline automates the process of checking out code from a GitHub repository, generating Docker image tags, building and pushing the Docker image to DockerHub, and performing environment cleanup.

## Configuration

To adapt the pipeline to your specific project, modify the following environment variables in the pipeline script:

- GITHUB_USER: GitHub username or organization name.
- GITHUB_REPO: Name of the GitHub repository.
- DOCKER_HUB_USERNAME: DockerHub username for image storage.
- DOCKER_REPO_NAME: Name of the Docker repository.
- BRANCH: Branch of the GitHub repository to be built and deployed.
- VERSION_PART: Versioning strategy (Patch, Minor, Major).
- DOCKER_JENKINS_CREDENTIALS_ID: Jenkins credentials ID for DockerHub login.

## Running the Pipeline

1. Create a new Jenkins job and select "Multibranch Pipeline" as the job type.

2. In the pipeline configuration add GitHub as source.

3. Configure the necessary parameters, such as GitHub and DockerHub credentials.

4. Save the pipeline configuration.

5. Run the Jenkins job to trigger the pipeline.

## Pipeline Overview

The Jenkins Pipeline is structured into several stages:

- Checkout Code: Checks out code from the specified GitHub repository and branch.

- Generate Docker Image Tag: Automatically generates a Docker image tag based on the specified versioning strategy (Patch, Minor, Major).

- Docker Login: Logs into DockerHub using provided credentials for image storage.

- Build: Builds the Docker image from checked-out code, incorporating project changes.

- Deploy: Pushes the built Docker image to DockerHub for deployment.

- Environment Cleanup: Removes the Docker image locally for resource management.

<br>

![](https://i.imgur.com/mkocrHE.png)
<br>

![](https://i.imgur.com/2zQaX2x.png)

<br>

![](https://i.imgur.com/hGnyEDH.png)
<br>

![](https://i.imgur.com/WnxCAGE.png)
<br>

![](https://i.imgur.com/4zukPXU.png)
`)
}
